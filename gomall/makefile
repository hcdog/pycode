
gen-frontend:cd frontend
	cwgo server --type HTTP --idl  ..\..\idl\frontend\home.proto --service frontend --module app/frontend -I ..\..\idl
gen-auth:
	cwgo server --type HTTP --idl  ..\..\idl\frontend\auth_page.proto --service frontend --module app/frontend -I ..\..\idl
gen- client-user:
	cwgo client --type RPC --service user --module github.com/hcdog/pycode/gomall/rpc_gen -I ../idl --idl  ../idl/user.proto
gen- service-user:	
	cwgo server --type RPC --service user --module github.com/hcdog/pycode/gomall/app/user --pass "-use github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto
.PHONY: gen_product
gen- client-product:
	@cd rpc_gen&&cwgo client --type RPC --service product --module github.com/hcdog/pycode/gomall/rpc_gen -I ../idl --idl  ../idl/product.proto
gen- service-product:	
	@cd app/product&&cwgo server --type RPC --service product --module github.com/hcdog/pycode/gomall/app/product --pass "-use github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto
.PHONY: gen_product_frontend
gen- frontend-product:
	@cd app/frontend&&cwgo server --type HTTP --idl  ..\..\idl\frontend\product_page.proto --service frontend --module github.com/hcdog/pycode/gomall/app/frontend -I ..\..\idl
.PHONY: gen_category
gen- frontend-category:
	@cd app/frontend&&cwgo server --type HTTP --idl  ..\..\idl\frontend\category_page.proto --service frontend --module github.com/hcdog/pycode/gomall/app/frontend -I ..\..\idl
gen- client-cart:
	@cd rpc_gen&&cwgo client --type RPC --service cart --module github.com/hcdog/pycode/gomall/rpc_gen -I ../idl --idl  ../idl/cart.proto
gen- service-cart:	
	@cd app/cart&&cwgo server --type RPC --service cart --module github.com/hcdog/pycode/gomall/app/cart --pass "-use github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto
gen- frontend-cart:
	@cd app/frontend&&cwgo server --type HTTP --idl  ..\..\idl\frontend\cart_page.proto --service frontend --module github.com/hcdog/pycode/gomall/app/frontend -I ..\..\idl
gen- client-payment:
	@cd rpc_gen&&cwgo client --type RPC --service payment --module github.com/hcdog/pycode/gomall/rpc_gen -I ../idl --idl  ../idl/payment.proto
gen- service-payment:	
	@cd app/payment&&cwgo server --type RPC --service payment --module github.com/hcdog/pycode/gomall/app/payment --pass "-use github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto
gen- client-checkout:
	@cd rpc_gen&&cwgo client --type RPC --service checkout --module github.com/hcdog/pycode/gomall/rpc_gen -I ../idl --idl  ../idl/checkout.proto
gen- service-checkout:	
	@cd app/checkout&&cwgo server --type RPC --service checkout --module github.com/hcdog/pycode/gomall/app/checkout --pass "-use github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
