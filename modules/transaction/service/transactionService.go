package service

import (
	"HarapanBangsaMarket/modules/transaction/domain/model"
	"HarapanBangsaMarket/modules/transaction/repository"
	"errors"
	"strconv"
	"strings"

	memberRepo "HarapanBangsaMarket/modules/member/repository"
	productRepo "HarapanBangsaMarket/modules/product/repository"
	promotionRepo "HarapanBangsaMarket/modules/promotion/repository"
)

func FindTransaction() (*[]model.Transaction, error) {
	return repository.FindTransaction()
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	return repository.FindOneTransaction(id)
}

func FindAllTransactionDetailByTransaction(transactionId int64) (*[]model.TransactionDetail, error) {
	transactions, err := repository.FindOneTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	return repository.FindTransactionDetailByTransactionId(transactions.Id)
}

func CreateTransaction(transactions *model.Transaction, transactionDetails *[]model.TransactionDetail) error {
	if *transactions.MemberId > 0 {
		member, err := memberRepo.FindOneMember(*transactions.MemberId)
		if err != nil {
			return err
		}
		transactions.MemberId = &member.Id
	} else {
		transactions.MemberId = nil
	}

	transactions.Status = 2
	transactions.Details = nil
	repository.CreateTransaction(transactions)

	for i := 0; i < len(*transactionDetails); i++ {
		var transactionDetail = &(*transactionDetails)[i]
		(*transactionDetails)[i].TransactionId = transactions.Id
		product, err := productRepo.FindOneProduct(transactionDetail.ProductId)
		if err != nil {
			return err
		}
		productCategory, err := productRepo.FindOneProductCategory(product.ProductCategoryId)
		if err != nil {
			return err
		}
		if productCategory.Id > 1 {
			if product.Stock < 0 {
				return errors.New("product empty")
			}
			product.Stock = product.Stock - transactionDetail.Quantity
			productRepo.UpdateProduct(product)
		}
		(*transactionDetails)[i].ProductId = product.Id
		if *transactionDetail.PromotionId > 0 {
			promotion, err := promotionRepo.FindOnePromotion(*transactionDetail.PromotionId)
			if err != nil {
				return err
			}
			(*transactionDetails)[i].PromotionId = func() *int64 { i := int64(promotion.Id); return &i }()
		} else {
			(*transactionDetails)[i].PromotionId = nil
		}
		repository.CreateTransactionDetail(transactionDetail)
	}

	return nil
}

func AmountToPayTransaction(transactionDetails *[]model.TransactionDetail) (float32, error) {
	total := float32(0)
	for i := 0; i < len(*transactionDetails); i++ {
		transactionDetail := &(*transactionDetails)[i]
		product, err := productRepo.FindOneProduct(transactionDetail.ProductId)
		if err != nil {
			return 0, err
		}
		productCategory, err := productRepo.FindOneProductCategory(product.ProductCategoryId)
		if err != nil {
			return 0, err
		}
		if productCategory.Id > 1 {
			total += product.Price * float32(transactionDetail.Quantity)
		} else {
			total += product.Price + float32(transactionDetail.Quantity)
		}

		if transactionDetail.PromotionId != nil {
			if *transactionDetail.PromotionId > 0 {
				promotion, err := promotionRepo.FindPromotionDetailByPromotionIdAndProductId(*transactionDetail.PromotionId, transactionDetail.ProductId)
				if err != nil {
					return 0, err
				}

				if strings.Contains(strings.ToLower(promotion.Promotion.Type), "buy") && strings.Contains(strings.ToLower(promotion.Promotion.Type), "get") {
					split := strings.Split(promotion.Promotion.Type, " ")
					promo, err := strconv.ParseFloat(split[len(split)-1], 32)
					if err != nil {
						promo = 0
					}
					total -= float32(promo) * product.Price
				} else {
					total -= promotion.Discount
				}
			}
		}
	}

	return total, nil
}
