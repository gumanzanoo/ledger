package transactions

import (
	"transactions/domain/entities"
	"transactions/domain/vo"
)

type transactionRepository interface {
	InsertTransaction(transaction entities.Transaction) error
	GetTransactionsByDocument(document vo.AccountOwnerDocument) ([]entities.Transaction, error)
}

type accountRepository interface {
	GetAccountByDocument(document vo.AccountOwnerDocument) (entities.Account, error)
}
