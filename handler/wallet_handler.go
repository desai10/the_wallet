package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"wallet/the_wallet/handler/model"
	"wallet/the_wallet/utils"
)

type WalletHandler struct {
}

func NewHandler() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) DepositMoney(ctx *gin.Context) {
	slog.InfoContext(ctx, "DepositMoney")
	req := model.DepositMoneyRequest{}
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		handleResponse(ctx, nil, err)
		return
	}
	slog.InfoContext(ctx, "Req", "req", req)
}

func (h *WalletHandler) WithdrawMoney(ctx *gin.Context) {
	slog.InfoContext(ctx, "WithdrawMoney")
	req := model.WithdrawMoneyRequest{}
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		handleResponse(ctx, nil, err)
		return
	}
	slog.InfoContext(ctx, "Req", "req", req)
}

func (h *WalletHandler) SendMoney(ctx *gin.Context) {
	slog.InfoContext(ctx, "SendMoney")
	req := model.SendMoneyRequest{}
	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		handleResponse(ctx, nil, err)
		return
	}
	slog.InfoContext(ctx, "Req", "req", req)
}

func (h *WalletHandler) GetWalletBalance(ctx *gin.Context) {
	slog.InfoContext(ctx, "GetWalletBalance", "userID", ctx.Param("user_id"))
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		handleResponse(ctx, nil, errors.New("userID not valid"))
		return
	}
	handleResponse(ctx, model.UserWallet{
		UserID:  userID,
		Balance: "1000",
	}, nil)
}

func (h *WalletHandler) GetTransactionHistory(ctx *gin.Context) {
	slog.InfoContext(ctx, "GetTransactionHistory", "userID", ctx.Param("user_id"))

	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		handleResponse(ctx, nil, errors.New("userID not valid"))
		return
	}
	handleResponse(ctx, model.UserTransactions{
		UserID: userID,
		Transactions: []model.Transaction{
			{
				OpType: "Deposit",
				Amount: "1000",
			},
		},
	}, nil)
}

func handleResponse(ctx *gin.Context, resp interface{}, err error) {
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, model.APIResponse{
			ErrorMessage: utils.ToPtr(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.APIResponse{
		Response: resp,
	})
}
