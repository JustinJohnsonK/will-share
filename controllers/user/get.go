package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type Transcations struct {
	UserId int64
	Amount int32
}

type userStatus struct {
	Borrowings []store.GetBorrowingsByUserIdRow `json:"borrowings"`
	Lendings   []store.GetLendingsByUserIdRow   `json:"lendings"`
}

func Get(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		user_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		i, err := s.UserService.Get(ctx, user_id)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.JSON(http.StatusOK, i)
	}
}

// Get the total borrowings and lendings of this user
func GetStatus(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		user_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.Unprocessable(c, err)
		}

		// Borrowings by this user
		userBorrowings, err := s.BillService.GetBorrowingsByUserID(ctx, user_id)

		// Lending by this user
		userLendings, err := s.BillService.GetLendingsByUserID(ctx, user_id)

		// Settle the balances between borrowings and lendings
		cleanTransactions(userBorrowings, userLendings)

		// Remove zero amounts from status
		for i, borrow := range userBorrowings {
			if borrow.Amount == 0 {
				userBorrowings = append(userBorrowings[:i], userBorrowings[i+1:]...)
			}
		}

		for i, lend := range userLendings {
			if lend.Amount == 0 {
				userLendings = append(userLendings[:i], userLendings[i+1:]...)
			}
		}

		// Create the responce
		_response := userStatus{
			Borrowings: userBorrowings,
			Lendings:   userLendings,
		}

		return response.Ok(c, _response)
	}
}

func cleanTransactions(borrowings []store.GetBorrowingsByUserIdRow, lendings []store.GetLendingsByUserIdRow) {
	// Find the duplicated in the lendings and borrowings
	// and settle them up.
	for i, borrow := range borrowings {
		_borrower := &borrowings[i]

		for j, lend := range lendings {
			_lender := &lendings[j]
			if borrow.LendUserID == lend.BorrowUserID {
				if borrow.Amount >= lend.Amount {
					_borrower.Amount = borrow.Amount - lend.Amount
					_lender.Amount = 0
				} else {
					_lender.Amount = lend.Amount - borrow.Amount
					_borrower.Amount = 0
				}
			}
		}
	}
}
