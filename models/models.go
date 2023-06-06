package models

type User struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorization struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorizationResponse struct {
	Token string `json:"token"`
}

type UpdateUserPassword struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type WithdrawReq struct {
	UserID       int     `json:"userID"`
	FromWalletID int     `json:"fromWalletID"`
	ToWalletID   int     `json:"toWalletID"`
	CurrencyID   int     `json:"currencyID"`
	Quontity     float64 `json:"quontity"`
	CardID       int     `json:"cardID"`
}

type WithdrawRes struct {
	Success bool `json:"success"`
}

type CreateCardReq struct {
	UserID int `json:"userID"`
}

type CreateCardRes struct {
	Success bool `json:"success"`
}

type CreateWalletReq struct {
	UserID     int `json:"userID"`
	CardID     int `json:"cardID"`
	CurrencyID int `json:"currencyID"`
}

type CreateWalletRes struct {
	Success bool `json:"success"`
}

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetCurrenciesRes struct {
	Currencies []Currency `json:"currencies"`
}

type Wallet struct {
	ID         int      `json:"id"`
	CurrencyID int      `json:"currencyID"`
	Quontity   float64  `json:"quontity"`
	Block      bool     `json:"block"`
	Currency   Currency `json:"currency"`
}

type Card struct {
	ID        int      `json:"id"`
	Wallets   []Wallet `json:"wallets"`
	Number    string   `json:"number"`
	Validated bool     `json:"validated"`
}

type Order struct {
	ID         int      `json:"id"`
	FromWallet Wallet   `json:"fromWallet"`
	ToWallet   Wallet   `json:"toWallet"`
	Quontity   float64  `json:"quontity"`
	Currency   Currency `json:"currency"`
}

type GetCardsReq struct {
	UserID int `json:"userID"`
}

type GetCardsRes struct {
	Cards []Card `json:"cards"`
}

type GetWalletInfoReq struct {
	UserID     int `json:"userID"`
	CardID     int `json:"cardID"`
	CurrencyID int `json:"currencyID"`
	WalletID   int `json:"WalletID"`
}

type GetWalletInfoRes struct {
	Wallet Wallet `json:"wallet"`
}

type GetWalletsReq struct {
	UserID     int `json:"userID"`
	CardID     int `json:"cardID"`
	CurrencyID int `json:"currencyID"`
}

type GetWalletsRes struct {
	Wallets []Wallet `json:"wallets"`
}

type GetCardInfoReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type GetCardInfoRes struct {
	Card Card `json:"card"`
}

type GetUserInfoReq struct {
	UserID int `json:"userID"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetUserInfoRes struct {
	Username string `json:"username"`
	Active   bool   `json:"active"`
	Role     Role   `json:"role"`
}

type GetOrdersReq struct {
	UserID int `json:"userID"`
}

type GetOrdersRes struct {
	Orders []Order `json:"orders"`
}

type GetOrderInfoReq struct {
	UserID  int `json:"userID"`
	OrderID int `json:"orderID"`
}

type GetOrderInfoRes struct {
	Order Order `json:"order"`
}

type GetCardOrdersReq struct {
	UserID     int `json:"userID"`
	CardID     int `json:"cardID"`
	CurrencyID int `json:"currencyID"`
}

type GetCardOrdersRes struct {
	Orders []Order `json:"orders"`
}

type GetBlockedCardsReq struct {
	UserID int `json:"userID"`
}

type GetBlockedCardsRes struct {
	Cards []Card `json:"cards"`
}

//--------------ADMIN----------------

type AdmGetCardsValidationReq struct {
}

type AdmGetCardsValidationRes struct {
	Cards []Card `json:"cards"`
}

type AdmDenyCardValidationReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type AdmDenyCardValidationRes struct {
	Success bool `json:"success"`
}

type AdmValidateCardReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type AdmValidateCardRes struct {
	Success bool `json:"success"`
}

type AdmGetUserCardsReq struct {
	UserID int `json:"userID"`
}

type AdmGetUserCardsRes struct {
	Cards []Card `json:"cards"`
}

type AdmGetUserCardInfoReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type AdmGetUserCardInfoRes struct {
	Card Card `json:"card"`
}

type AdmGetUserOrdersReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type AdmGetUserOrdersRes struct {
	Orders []Order `json:"orders"`
}

type AdmGetUserOrderInfoReq struct {
	UserID  int `json:"userID"`
	CardID  int `json:"cardID"`
	OrderID int `json:"orderID"`
}

type AdmGetUserOrderInfoRes struct {
	Order Order `json:"order"`
}

type AdmBlockUserWalletReq struct {
	UserID   int `json:"userID"`
	CardID   int `json:"cardID"`
	WalletID int `json:"walletID"`
}

type AdmBlockUserWalletRes struct {
	Success bool `json:"success"`
}

type AdmBlockUserCardReq struct {
	UserID   int `json:"userID"`
	CardID   int `json:"cardID"`
	WalletID int `json:"walletID"`
}

type AdmBlockUserCardRes struct {
	Success bool `json:"success"`
}

type AdmUnblockUserWalletReq struct {
	UserID   int `json:"userID"`
	CardID   int `json:"cardID"`
	WalletID int `json:"walletID"`
}

type AdmUnblockUserWalletRes struct {
	Success bool `json:"success"`
}

type AdmUnblockUserCardReq struct {
	UserID int `json:"userID"`
	CardID int `json:"cardID"`
}

type AdmUnblockUserCardRes struct {
	Success bool `json:"success"`
}
