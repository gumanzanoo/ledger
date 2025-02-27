package entities

import (
	"transactions/domain/vo"
)

type Transaction struct {
	id                   vo.TransactionID
	relatedTransactionID vo.RelatedTransactionID
	accountOwnerDocument vo.AccountOwnerDocument
	transactionType      vo.TransactionType
	amount               int
}

func NewTransaction(
	relatedTransactionID vo.RelatedTransactionID,
	accountOwnerDocument vo.AccountOwnerDocument,
	transactionType vo.TransactionType,
	amount int,
) Transaction {
	definitiveRelatedTransactionID := relatedTransactionID
	return Transaction{
		id:                   vo.NewTransactionID(),
		relatedTransactionID: definitiveRelatedTransactionID,
		accountOwnerDocument: accountOwnerDocument,
		transactionType:      transactionType,
		amount:               amount,
	}
}
