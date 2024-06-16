package payment

import (
	"github.com/qingwg/payjs"
	model "github.com/yyyyymmmmm/Test/models"
	"github.com/yyyyymmmmm/Test/pkg/serializer"
)

// PayJSClient PayJS支付处理
type PayJSClient struct {
	Client *payjs.PayJS
}

// Create 创建订单
func (pay *PayJSClient) Create(order *model.Order, pack *serializer.PackProduct, group *serializer.GroupProducts, user *model.User) (*OrderCreateRes, error) {
	if _, err := order.Create(); err != nil {
		return nil, ErrInsertOrder.WithError(err)
	}

	PayNative := pay.Client.GetNative()
	res, err := PayNative.Create(int64(order.Price*order.Num), order.Name, order.OrderNo, "", "")
	if err != nil {
		return nil, ErrIssueOrder.WithError(err)
	}

	return &OrderCreateRes{
		Payment: true,
		QRCode:  res.CodeUrl,
		ID:      order.OrderNo,
	}, nil
}
