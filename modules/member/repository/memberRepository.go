package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/member/domain/model"
	"errors"
)

func FindAllMember() (*[]model.Member, error) {
	var members []model.Member
	result := db.Orm.Where("deleted_at is null").Find(&members)
	if result.Error != nil {
		return nil, result.Error
	}

	return &members, nil
}

func FindOneMember(id int64) (*model.Member, error) {
	var member model.Member
	result := db.Orm.First(&member, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no member category found")
	}

	return &member, nil
}

func CreateMember(member *model.Member) error {
	result := db.Orm.Create(member)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateMember(member *model.Member) (*model.Member, error) {
	result := db.Orm.Save(&member)
	if result.Error != nil {
		return nil, result.Error
	}
	member, _ = FindOneMember(member.Id)
	return member, nil
}

func DeleteMember(member *model.Member) (*model.Member, error) {
	result := db.Orm.Delete(&member)
	if result.Error != nil {
		return nil, result.Error
	}
	return member, nil
}
