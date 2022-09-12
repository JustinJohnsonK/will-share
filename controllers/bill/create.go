package bill

import (
	"context"
	"math"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/internal/utils"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type user struct {
	UserId int64 `json:"user_id"`
	Amount int32 `json:"amount"`
}

type lenderValues struct {
	LenderId       int64
	LendAmount     int32
	LendProportion float32
}

type createBillRequest struct {
	BillDescription string `json:"bill_description"`
	BillTitle       string `json:"bill_title"`
	GroupId         int64  `json:"group_id"`
	Amount          int32  `json:"amount"`
	Borrowers       []user `json:"borrowers"`
	Lenders         []user `json:"lenders"`
}

func Create(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var bill createBillRequest
		if err := c.Bind(&bill); err != nil {
			return response.InternalError(c, nil)
		}

		// Validate the bill before creating entries
		if !validateBillData(s, ctx, bill.Amount, bill.GroupId, bill.Borrowers, bill.Lenders) {
			return response.BadRequest(c)
		}

		// Add the bill to bill table
		newBill := store.AddBillParams{
			BillDescription: utils.ToNullString(bill.BillDescription),
			BillTitle:       utils.ToNullString(bill.BillTitle),
			GroupID:         bill.GroupId,
			Amount:          bill.Amount,
		}

		// Create Bill entry in bills table
		createdBill, err := s.BillService.Create(ctx, newBill)
		if err != nil {
			return response.BadRequest(c)
		}

		// Process the bill data for user_bills table
		bill_users_amount := generateBillForUserBills(bill, newBill, createdBill.BillID)

		// Add the generated bills to bill_users table
		for _, bill := range bill_users_amount {
			_, err := s.BillService.CreateUserBill(ctx, bill)
			if err != nil {
				return response.InternalError(c, nil)
			}
		}

		return response.Created(c, createdBill)
	}
}

// Check if the amount given is valid and
// The users are valid and belongs to the same group
func validateBillData(s services.APIService, ctx context.Context, amount int32, groupId int64, borrowers, lenders []user) bool {
	total_borrow_amount := 0
	total_lend_amount := 0
	bill_users := []int64{}

	for _, borrower := range borrowers {
		total_borrow_amount += int(borrower.Amount)
		bill_users = append(bill_users, borrower.UserId)
	}

	for _, lender := range lenders {
		total_lend_amount += int(lender.Amount)
		bill_users = append(bill_users, lender.UserId)
	}

	// Check for amount mismatches in the bill
	if total_borrow_amount != total_lend_amount || total_lend_amount != int(amount) {
		return false
	}

	// Check all users are part of the same group id
	group_users, err := s.GroupService.GetGroupUserIds(ctx, groupId)
	if err != nil {
		return false
	}

	for _, id := range bill_users {
		if !utils.Int64InSlice(id, group_users) {
			return false
		}
	}

	return true
}

func generateBillForUserBills(bill createBillRequest, newBill store.AddBillParams, billId int64) []store.AddUserBillParams {
	lender_map := map[int64]int32{}
	borrower_map := map[int64]int32{}
	lender_proportions := map[int64]float32{}

	for _, lender := range bill.Lenders {
		lender_map[lender.UserId] = lender_map[lender.UserId] + lender.Amount
	}

	for _, borrower := range bill.Borrowers {
		borrower_map[borrower.UserId] = borrower_map[borrower.UserId] + borrower.Amount
	}

	for borrwer, borrow_amount := range borrower_map {
		if lender_map[borrwer] > 0 {
			lend_amount := lender_map[borrwer]

			if lend_amount == borrow_amount {
				lender_map[borrwer] = 0
				borrower_map[borrwer] = 0
			} else if lend_amount > borrow_amount {
				amount := int32(math.Abs(float64(lend_amount) - float64(borrow_amount)))
				lender_map[borrwer] = amount
				borrower_map[borrwer] = 0
			} else {
				borrower_map[borrwer] = borrow_amount - lend_amount
				lender_map[borrwer] = 0
			}
		}
	}

	var total_lend_amount int32
	for _, lend_amount := range lender_map {
		total_lend_amount += lend_amount
	}

	// Add the proportion of cash lended by the lenders
	for lender, amount := range lender_map {
		if amount > 0 {
			lender_proportions[lender] = float32(amount) / float32(total_lend_amount)
		}
	}

	// Calulate amount each borrower should pay to the lender based on the proportion of lending
	bill_users_amount := []store.AddUserBillParams{}
	for borrower, amount := range borrower_map {
		if amount == 0 {
			continue
		}

		for lender, proportion := range lender_proportions {
			if lender == borrower || amount == 0 || proportion == 0 {
				continue
			}
			bill_users_amount = append(bill_users_amount, store.AddUserBillParams{
				BillID:       billId,
				GroupID:      newBill.GroupID,
				LendUserID:   lender,
				BorrowUserID: borrower,
				Amount:       int32(float32(amount) * proportion),
			})
		}
	}

	return bill_users_amount
}
