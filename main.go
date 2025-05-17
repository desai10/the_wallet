package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"wallet/the_wallet/handler"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	router := gin.Default()

	wh := handler.NewHandler()

	router.POST("/wallet/deposit", wh.DepositMoney)
	router.POST("/wallet/withdraw", wh.WithdrawMoney)
	router.POST("/wallet/send", wh.SendMoney)
	router.GET("/wallet/balance/:user_id", wh.GetWalletBalance)
	router.GET("/wallet/history/:user_id", wh.GetTransactionHistory)

	slog.Info("starting server...")
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
