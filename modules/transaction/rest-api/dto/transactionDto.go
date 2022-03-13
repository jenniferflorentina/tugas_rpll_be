package dto

import (
	"HarapanBangsaMarket/base"
	member "HarapanBangsaMarket/modules/member/rest-api/dto"
)

type TransactionDTO struct {
	base.DTO
	Member                 member.MemberDTO       `json:"member"`
	Details                []TransactionDetailDTO `json:"details"`
	AmountToPayTransaction float32                `json:"amountToPay"`
}
