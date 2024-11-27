package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/hcdog/pycode/gomall/app/payment/biz/dal/model"
	"github.com/hcdog/pycode/gomall/app/payment/biz/dal/mysql"
	payment "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	//构造卡信息
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400401, err.Error())
	}

	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(400501, err.Error())
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: translationId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
}