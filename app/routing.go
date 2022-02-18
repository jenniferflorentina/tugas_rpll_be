package app

import (
	r "HarapanBangsaMarket/modules/router"
)

func mapURLs() {
	r.RouteUsers(router)
	r.RoutePromotions(router)
}
