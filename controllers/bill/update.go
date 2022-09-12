package bill

import (
	"github.com/JustinJohnsonK/will-share/internal/services"
	"github.com/JustinJohnsonK/will-share/internal/store"
	"github.com/JustinJohnsonK/will-share/pkg/response"
	"github.com/labstack/echo/v4"
)

type settleBillRequest struct {
	BillId int64 `json:"bill_id"`
}

type settleGroupBillRequest struct {
	GroupId int64 `json:"group_id"`
}

func SettleBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var bill settleBillRequest
		if err := c.Bind(&bill); err != nil {
			return response.InternalError(c, nil)
		}

		amount_settled, err := s.BillService.SettleBillByBillId(ctx, bill.BillId)
		if err != nil {
			return response.BadRequest(c)
		}

		return response.Ok(c, map[string]int64{"amount-settled": int64(amount_settled)})
	}
}

func SettleGroupBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var bill settleGroupBillRequest
		if err := c.Bind(&bill); err != nil {
			return response.InternalError(c, nil)
		}

		err := s.BillService.SettleBillByBillGroupId(ctx, bill.GroupId)
		if err != nil {
			return response.BadRequest(c)
		}

		return response.Ok(c, nil)
	}
}

func SettleUserBill(s services.APIService) func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var userIds store.SettleUserBillsByUserIdParams
		if err := c.Bind(&userIds); err != nil {
			return response.InternalError(c, nil)
		}

		err := s.BillService.SettleBillByBillUserId(ctx, userIds)
		if err != nil {
			return response.BadRequest(c)
		}

		return response.Ok(c, nil)
	}
}
