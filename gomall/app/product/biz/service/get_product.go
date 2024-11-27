package service

import (
	"context"
	
	"github.com/hcdog/pycode/gomall/app/product/biz/model"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hcdog/pycode/gomall/app/product/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/product/biz/dal/redis"
	product "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product id is required")
	}

	p, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
		},
	}, err
}
