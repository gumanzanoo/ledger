package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	v1 "transactions/app/service/api/v1"
	"transactions/domain/transactions"
	"transactions/gateways/postgres"
	accountsRepository "transactions/gateways/postgres/accounts"
	transactionsRepository "transactions/gateways/postgres/transactions"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()
	conn := postgres.New()
	defer conn.Close(ctx)

	accountRepository := accountsRepository.Repository{
		Conn: conn,
	}

	transactionsRepository := transactionsRepository.Repository{
		Conn: conn,
	}

	makeTransactionUC := transactions.MakeTransactionUC{
		TransactionRepository: transactionsRepository,
		AccountRepository:     accountRepository,
	}

	apiV1 := v1.API{
		MakeTransaction: v1.MakeTransactionHandler(makeTransactionUC),
	}

	router := chi.NewRouter()
	apiV1.Routes(router)

	server := http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("server error:", err)
	}
}
