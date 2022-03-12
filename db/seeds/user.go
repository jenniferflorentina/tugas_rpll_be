package seeds

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/hashing"
	"HarapanBangsaMarket/modules/user/domain/model"
)

func (s *Seed) SeedUser() {
	var user []model.User
	var count int64
	db.Orm.Model(&user).Count(&count)

	if count > 0 {
		return
	}

	users := make([]model.User, 4)
	users[0] = model.User{
		Username: "Jojoan",
		Password: hashing.HashAndSalt([]byte("joan123")),
		Name:     "Eirenika Joanna Grace",
		Role:     0,
	}
	users[1] = model.User{
		Username: "JeniJenJen",
		Password: hashing.HashAndSalt([]byte("jeni123")),
		Name:     "Jenifer Florentina",
		Role:     1,
	}
	users[2] = model.User{
		Username: "Marimar",
		Password: hashing.HashAndSalt([]byte("maria123")),
		Name:     "Maria Vabiolla",
		Role:     1,
	}
	users[3] = model.User{
		Username: "Enjeljel",
		Password: hashing.HashAndSalt([]byte("enjel123")),
		Name:     "Tridia Enjeliani",
		Role:     1,
	}

	_ = db.Orm.Create(&users)
}
