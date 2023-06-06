package repository

import (
	"dbKursach/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(params models.CreateUser) (err error) {
	_, err = r.db.Exec(queryCreateUser, params.Username, params.Password)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUserPassword(params models.UpdateUserPassword) (err error) {
	var ok bool
	err = r.db.Get(&ok, queryUpdatePassword, params.NewPassword, params.Username, params.CurrentPassword)
	if err != nil {
		return errors.New("Current password isn't correct")
	}
	if !ok {
		return errors.New("Current password isn't correct")
	}
	return
}

func (r *Repository) Authorization(params models.Authorization) (result models.AuthorizationResponse, err error) {
	var authed bool
	tx, err := r.db.Begin()
	if err != nil {
		return
	}
	err = r.db.Get(&authed, queryCheckUserParams, params.Username, params.Password)
	if err != nil {
		return
	}
	_, err = tx.Exec(queryUpdateSessionKey, params.Username)
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return
		}
		return
	}
	token, err := uuid.NewUUID()
	if err != nil {
		return
	}
	_, err = tx.Exec(queryAuthorizationInsertSessionKey, token.String(), true, params.Username)
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return
		}
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	result.Token = token.String()
	return
}

func (r *Repository) Withdraw(params models.WithdrawReq) (models.WithdrawRes, error) {
	_, err := r.db.Exec(queryWithdraw, params.FromWalletID, params.ToWalletID, params.Quontity, params.CurrencyID, params.CardID)
	if err != nil {
		return models.WithdrawRes{}, err
	}
	return models.WithdrawRes{
		Success: true,
	}, nil
}

func (r *Repository) CreateCard(params models.CreateCardReq) (models.CreateCardRes, error) {
	_, err := r.db.Exec(queryCreateCard, params.UserID)
	if err != nil {
		return models.CreateCardRes{}, err
	}
	return models.CreateCardRes{Success: true}, nil
}

func (r *Repository) GetCurrencies() (models.GetCurrenciesRes, error) {
	var data models.GetCurrenciesRes
	if err := r.db.Select(&data.Currencies, queryGetCurrencies); err != nil {
		return models.GetCurrenciesRes{}, err
	}
	return data, nil
}

func (r *Repository) GetCards(params models.GetCardsReq) (models.GetCardsRes, error) {
	var data models.GetCardsRes
	if err := r.db.Select(&data.Cards, queryGetCards, params.UserID); err != nil {
		return models.GetCardsRes{}, err
	}
	return data, nil
}

func (r *Repository) GetBlockedCards(params models.GetBlockedCardsReq) (models.GetBlockedCardsRes, error) {
	var data models.GetBlockedCardsRes
	if err := r.db.Select(&data.Cards, queryGetBlockedCards, params.UserID); err != nil {
		return models.GetBlockedCardsRes{}, err
	}
	return data, nil
}

func (r *Repository) GetUserInfo(params models.GetUserInfoReq) (models.GetUserInfoRes, error) {
	var data models.GetUserInfoRes
	if err := r.db.Get(&data, queryGetUserInfo, params.UserID); err != nil {
		return models.GetUserInfoRes{}, err
	}
	return data, nil
}
func (r *Repository) GetOrders(params models.GetOrdersReq) (models.GetOrdersRes, error) {
	var data models.GetOrdersRes
	if err := r.db.Select(&data.Orders, queryGetOrders, params.UserID); err != nil {
		return models.GetOrdersRes{}, err
	}
	return data, nil
}
func (r *Repository) GetOrderInfo(params models.GetOrderInfoReq) (models.GetOrderInfoRes, error) {
	var data models.GetOrderInfoRes
	if err := r.db.Get(&data, queryGetOrderInfo, params.UserID, params.OrderID); err != nil {
		return models.GetOrderInfoRes{}, nil
	}
	return data, nil
}
func (r *Repository) GetCardOrders(params models.GetCardOrdersReq) (models.GetCardOrdersRes, error) {
	var data models.GetCardOrdersRes
	if err := r.db.Select(&data, queryGetCardOrders, params.UserID, params.CardID); err != nil {
		return models.GetCardOrdersRes{}, err
	}
	return data, nil
}

//--------------ADM----------------

func (r *Repository) AdmGetCardValidation(params models.AdmGetCardsValidationReq) (models.AdmGetCardsValidationRes, error) {
	var data models.AdmGetCardsValidationRes
	return data, nil
}
func (r *Repository) AdmValidateCard(params models.AdmValidateCardReq) (models.AdmValidateCardRes, error) {
	var data models.AdmValidateCardRes
	return data, nil
}
func (r *Repository) AdmDenyValidationCard(params models.AdmDenyCardValidationReq) (models.AdmDenyCardValidationRes, error) {
	var data models.AdmDenyCardValidationRes
	return data, nil
}
func (r *Repository) AdmGetUserCards(params models.AdmGetUserCardsReq) (models.AdmGetUserCardsRes, error) {
	var data models.AdmGetUserCardsRes
	return data, nil
}
func (r *Repository) AdmGetUserCardInfo(params models.AdmGetUserCardInfoReq) (models.AdmGetUserCardInfoRes, error) {
	var data models.AdmGetUserCardInfoRes
	return data, nil
}
func (r *Repository) AdmGetUserOrders(params models.AdmGetUserOrdersReq) (models.AdmGetUserOrdersRes, error) {
	var data models.AdmGetUserOrdersRes
	return data, nil
}
func (r *Repository) AdmGetUserOrderInfo(params models.AdmGetUserOrderInfoReq) (models.AdmGetUserOrderInfoRes, error) {
	var data models.AdmGetUserOrderInfoRes
	return data, nil
}
func (r *Repository) AdmBlockUserCard(params models.AdmBlockUserCardReq) (models.AdmBlockUserCardRes, error) {
	var data models.AdmBlockUserCardRes
	return data, nil
}
func (r *Repository) AdmUnblockUserCard(params models.AdmUnblockUserCardReq) (models.AdmUnblockUserCardRes, error) {
	var data models.AdmUnblockUserCardRes
	return data, nil
}
