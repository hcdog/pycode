package service

import (
	"context"
	
	"github.com/hcdog/pycode/gomall/app/product/biz/model"
	"github.com/hcdog/pycode/gomall/app/product/biz/dal/mysql"
	product "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	p, err := model.SearchProduct(mysql.DB, s.ctx, req.Query)
	var results []*product.Product
	for _, v := range p {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.SearchProductsResp{Results: results}, err
}
