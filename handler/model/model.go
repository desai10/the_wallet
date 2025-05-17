package model

type DepositMoneyRequest struct {
	UserID        int64  `json:"user_id"`
	DepositAmount string `json:"deposit_amount"`
}

type WithdrawMoneyRequest struct {
	UserID         int64  `json:"user_id"`
	WithdrawAmount string `json:"withdraw_amount"`
}

type SendMoneyRequest struct {
	UserID         int64  `json:"user_id"`
	ReceiverUserID int64  `json:"receiver_user_id"`
	Amount         string `json:"amount"`
}

type Transaction struct {
	OpType    string `json:"op_type"`
	Amount    string `json:"amounts"`
	Recipient *int64 `json:"recipient"`
}

type UserTransactions struct {
	UserID       int64         `json:"user_id"`
	Transactions []Transaction `json:"transactions"`
}

type UserWallet struct {
	UserID  int64  `json:"user_id"`
	Balance string `json:"balance"`
}

type APIResponse struct {
	Response     interface{} `json:"response"`
	ErrorMessage *string     `json:"error_message"`
}
