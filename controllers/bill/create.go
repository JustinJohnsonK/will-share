package bill

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/internal/utils"
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
			return err
		}

		// Validate the bill before creating entries
		if !validateBillData(s, ctx, bill.Amount, bill.GroupId, bill.Borrowers, bill.Lenders) {
			return errors.New("Invalid Bill")
		}

		// Add the bill to bill table
		newBill := store.AddBillParams{
			BillDescription: utils.ToNullString(bill.BillDescription),
			BillTitle:       utils.ToNullString(bill.BillTitle),
			GroupID:         bill.GroupId,
			Amount:          bill.Amount,
		}

		fmt.Printf("Bill input data = %+v\n", newBill)

		// Create Bill entry in bills table
		createdBill, err := s.BillService.Create(ctx, newBill)
		if err != nil {
			return err
		}

		// Process the bill data for user_bills table
		bill_users_amount := generateBillForUserBills(bill, newBill, createdBill.BillID)
		fmt.Printf("User Bills generated = %+v\n", bill_users_amount)

		// Add the generated bills to bill_users table
		for _, bill := range bill_users_amount {
			_, err := s.BillService.CreateUserBill(ctx, bill)
			if err != nil {
				return err
			}
		}

		return c.JSON(http.StatusCreated, createdBill)
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
	lender_amounts := []lenderValues{}
	total_lend_amount := 0

	for i, borrower := range bill.Borrowers {
		amount := borrower.Amount

		for _, lender := range bill.Lenders {
			if lender.UserId == borrower.UserId {
				_borrower := &bill.Borrowers[i]
				if lender.Amount == borrower.Amount {
					// bill.Borrowers = append(bill.Borrowers[:i], bill.Borrowers[i+1:]...)
					_borrower.Amount = 0
					continue
				} else if lender.Amount > borrower.Amount {
					amount = int32(math.Abs(float64(lender.Amount) - float64(borrower.Amount)))
					total_lend_amount += int(amount)
					_borrower.Amount = 0

					lender_amounts = append(lender_amounts, lenderValues{
						LenderId:   lender.UserId,
						LendAmount: amount,
					})
				} else {
					_borrower.Amount = borrower.Amount - lender.Amount
				}
			}
		}
	}

	fmt.Printf("Updated Borrower bills = %+v\n", bill.Borrowers)
	fmt.Println("Total lendings = ", total_lend_amount)

	// Add the proportion of cash lended by the lenders
	for i, lender := range lender_amounts {
		fmt.Println(lender)
		_lender := &lender_amounts[i]
		_lender.LendProportion = float32(lender.LendAmount) / float32(total_lend_amount)
		fmt.Println("Lend proportion = ", lender.LendAmount, total_lend_amount, float32(lender.LendAmount)/float32(total_lend_amount), lender)
	}

	fmt.Printf("Processed Lendings = %+v\n", lender_amounts)

	// Calulate amount each borrower should pay to the lender based on the proportion of lending
	bill_users_amount := []store.AddUserBillParams{}
	for _, borrower := range bill.Borrowers {

		for _, lender := range lender_amounts {
			if lender.LenderId == borrower.UserId || borrower.Amount == 0 {
				continue
			}

			bill_users_amount = append(bill_users_amount, store.AddUserBillParams{
				BillID:       billId,
				GroupID:      newBill.GroupID,
				LendUserID:   lender.LenderId,
				BorrowUserID: borrower.UserId,
				Amount:       int32(float32(borrower.Amount) * lender.LendProportion),
			})
		}
	}

	return bill_users_amount
}
