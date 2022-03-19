package service

import (
	memberRepository "HarapanBangsaMarket/modules/member/repository"
	model "HarapanBangsaMarket/modules/payment/domain/model"
	repository "HarapanBangsaMarket/modules/payment/repository"
	transactionRepository "HarapanBangsaMarket/modules/transaction/repository"
	"errors"
)

func FindPayment() (*[]model.Payment, error) {
	return repository.FindPayment()
}

func FindOnePayment(id int64) (*model.Payment, error) {
	return repository.FindOnePayment(id)
}

func FindOnePaymentByTransactionId(transactionId int64) (*model.Payment, error) {
	return repository.FindOnePaymentByTransactionId(transactionId)
}

func CreatePayment(payment *model.Payment) error {

	var transaction, err = transactionRepository.FindOneTransaction(payment.TransactionId)
	if err != nil {
		return err
	}

	if transaction.MemberId != nil {
		var member, err = memberRepository.FindOneMember(*transaction.MemberId)
		if err != nil {
			return err
		}
		if payment.Point > 0 {
			member.Point -= int64(payment.Point)
			if member.Point < 0 {
				return errors.New("point is not enough")
			}
			memberRepository.UpdateMember(member)
		}
	}

	payment.Status = "Complete"
	return repository.CreatePayment(payment)
}
