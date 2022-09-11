package bill

import (
	"net/http"

	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/labstack/echo/v4"
)

type settleBillRequest struct {
	BillId int64
}

type settleGroupBillRequest struct {
	GroupId int64
}

type settleUserBillRequest struct {
	LendUserID   int64
	BorrowUserID int64
}

func SettleBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var bill settleBillRequest
		if err := c.Bind(&bill); err != nil {
			return err
		}

		amount_settled, err := s.BillService.SettleBillByBillId(ctx, bill.BillId)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]int64{"Amount Settled": amount_settled})
	}
}

func SettleGroupBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var bill settleGroupBillRequest
		if err := c.Bind(&bill); err != nil {
			return err
		}

		amount_settled, err := s.BillService.SettleBillByBillId(ctx, bill.GroupId)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, map[string]int64{"Amount Settled": amount_settled})
	}
}

func SettleUserBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var userIds store.SettleUserBillsByUserIdParams
		if err := c.Bind(&userIds); err != nil {
			return err
		}

		err := s.BillService.SettleBillByBillUserId(ctx, userIds)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, err)
	}
}
