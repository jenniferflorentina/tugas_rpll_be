package app

import (
	r "HarapanBangsaMarket/modules/router"
)

func mapURLs() {
	r.RouteUsers(router)
	r.RouteProductCategories(router)
	r.RouteProducts(router)
	r.RoutePromotions(router)
	r.RouteMembers(router)
}
