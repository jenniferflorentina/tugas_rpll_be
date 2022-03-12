package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/member/domain/model"
)

func (s *Seed) SeedMember() {
	var members []model.Member
	var count int64
	db.Orm.Model(&members).Count(&count)

	if count > 0 {
		return
	}

	members = make([]model.Member, 4)
	members[0] = model.Member{
		Point:     10,
		Telephone: "0239283492493",
	}
	members[1] = model.Member{
		Point:     100000,
		Telephone: "0239283422495",
	}
	members[2] = model.Member{
		Point:     1000,
		Telephone: "0838273857879",
	}
	members[3] = model.Member{
		Point:     2340,
		Telephone: "0838273823170",
	}
	_ = db.Orm.Create(&members)
}
