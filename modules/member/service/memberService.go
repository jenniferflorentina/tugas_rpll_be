package service

import (
	"HarapanBangsaMarket/modules/member/domain/model"
	"HarapanBangsaMarket/modules/member/repository"
	"HarapanBangsaMarket/modules/member/rest-api/dto"
)

func FindAllMember() (*[]model.Member, error) {
	return repository.FindAllMember()
}

func FindOneMember(id int64) (*model.Member, error) {
	return repository.FindOneMember(id)
}

func CreateMember(Member *model.Member) error {
	return repository.CreateMember(Member)
}

func UpdateMember(updateDto *dto.UpdateMemberDTO, id int64, userId int64) (*model.Member, error) {
	member, err := repository.FindOneMember(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Name != "" {
		member.Name = updateDto.Name
	}

	if updateDto.Telephone != "" {
		member.Telephone = updateDto.Telephone
	}

	if updateDto.Point >= 0 {
		member.Point = updateDto.Point
	}

	member.UpdatedBy = userId

	member, err = repository.UpdateMember(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func DeleteMember(id int64, userId int64) (*model.Member, error) {
	member, err := repository.FindOneMember(id)
	if err != nil {
		return nil, err
	}

	member.DeletedBy = userId
	member, err = repository.UpdateMember(member)
	if err != nil {
		return nil, err
	}

	member, err = repository.DeleteMember(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}
