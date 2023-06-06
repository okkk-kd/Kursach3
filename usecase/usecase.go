package usecase

import (
	"dbKursach/internal/models"
	"dbKursach/internal/repository"
)

type Usecase struct {
	r repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return Usecase{
		r: r,
	}
}

func (uc *Usecase) CreateUser(params models.CreateUser) (err error) {
	err = uc.r.CreateUser(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) UpdateUserPassword(params models.UpdateUserPassword) (err error) {
	err = uc.r.UpdateUserPassword(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) Authorization(params models.Authorization) (result models.AuthorizationResponse, err error) {
	result, err = uc.r.Authorization(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) Withdraw(params models.WithdrawReq) (result models.WithdrawRes, err error) {
	result, err = uc.r.Withdraw(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) CreateCard(params models.CreateCardReq) (result models.CreateCardRes, err error) {
	result, err = uc.r.CreateCard(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetCurrencies() (result models.GetCurrenciesRes, err error) {
	result, err = uc.r.GetCurrencies()
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetCards(params models.GetCardsReq) (result models.GetCardsRes, err error) {
	result, err = uc.r.GetCards(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetBlockedCards(params models.GetBlockedCardsReq) (result models.GetBlockedCardsRes, err error) {
	result, err = uc.r.GetBlockedCards(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetUserInfo(params models.GetUserInfoReq) (result models.GetUserInfoRes, err error) {
	result, err = uc.r.GetUserInfo(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetOrders(params models.GetOrdersReq) (result models.GetOrdersRes, err error) {
	result, err = uc.r.GetOrders(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetOrderInfo(params models.GetOrderInfoReq) (result models.GetOrderInfoRes, err error) {
	result, err = uc.r.GetOrderInfo(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) GetCardOrders(params models.GetCardOrdersReq) (result models.GetCardOrdersRes, err error) {
	result, err = uc.r.GetCardOrders(params)
	if err != nil {
		return
	}
	return
}

//--------------ADM----------------

func (uc *Usecase) AdmGetCardValidation(params models.AdmGetCardsValidationReq) (result models.AdmGetCardsValidationRes, err error) {
	result, err = uc.r.AdmGetCardValidation(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmValidateCard(params models.AdmValidateCardReq) (result models.AdmValidateCardRes, err error) {
	result, err = uc.r.AdmValidateCard(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmDenyValidationCard(params models.AdmDenyCardValidationReq) (result models.AdmDenyCardValidationRes, err error) {
	result, err = uc.r.AdmDenyValidationCard(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmGetUserCards(params models.AdmGetUserCardsReq) (result models.AdmGetUserCardsRes, err error) {
	result, err = uc.r.AdmGetUserCards(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmGetUserCardInfo(params models.AdmGetUserCardInfoReq) (result models.AdmGetUserCardInfoRes, err error) {
	result, err = uc.r.AdmGetUserCardInfo(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmGetUserOrders(params models.AdmGetUserOrdersReq) (result models.AdmGetUserOrdersRes, err error) {
	result, err = uc.r.AdmGetUserOrders(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmGetUserOrderInfo(params models.AdmGetUserOrderInfoReq) (result models.AdmGetUserOrderInfoRes, err error) {
	result, err = uc.r.AdmGetUserOrderInfo(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmBlockUserCard(params models.AdmBlockUserCardReq) (result models.AdmBlockUserCardRes, err error) {
	result, err = uc.r.AdmBlockUserCard(params)
	if err != nil {
		return
	}
	return
}

func (uc *Usecase) AdmUnblockUserCard(params models.AdmUnblockUserCardReq) (result models.AdmUnblockUserCardRes, err error) {
	result, err = uc.r.AdmUnblockUserCard(params)
	if err != nil {
		return
	}
	return
}
