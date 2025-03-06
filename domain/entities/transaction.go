package entities

import (
	"fmt"
	"transactions/domain/vo"
)

type Transaction struct {
	id                   vo.TransactionID
	relatedTransactionID vo.RelatedTransactionID
	accountOwnerDocument vo.AccountOwnerDocument
	transactionType      vo.TransactionType
	amount               int
}

func (t Transaction) ID() vo.TransactionID {
	return t.id
}

func (t Transaction) RelatedTransactionID() vo.RelatedTransactionID {
	return t.relatedTransactionID
}

func (t Transaction) AccountOwnerDocument() vo.AccountOwnerDocument {
	return t.accountOwnerDocument
}

func (t Transaction) TransactionType() vo.TransactionType {
	return t.transactionType
}

func (t Transaction) Amount() int {
	return t.amount
}

func (t Transaction) ParseTransaction(
	id string,
	relatedTransactionID string,
	accountOwnerDocument string,
	transactionType string,
	amount int,
) error {
	definitiveID, err := vo.ParseTransactionID(id)
	if err != nil {
		return fmt.Errorf("could not parse transaction id %s", id)
	}

	definitiveRelatedTransactionID, err := vo.ParseRelatedTransactionID(relatedTransactionID)
	if err != nil {
		return fmt.Errorf("could not parse related transaction id %s", relatedTransactionID)
	}

	definitiveAccountOwnerDocument, err := vo.ParseAccountOwnerDocument(accountOwnerDocument)
	if err != nil {
		return fmt.Errorf("could not parse account owner document %s", accountOwnerDocument)
	}

	definitiveTransactionType, err := vo.ParseTransactionType(transactionType)
	if err != nil {
		return fmt.Errorf("could not parse transaction type %s", transactionType)
	}

	t.id = definitiveID
	t.relatedTransactionID = definitiveRelatedTransactionID
	t.accountOwnerDocument = definitiveAccountOwnerDocument
	t.transactionType = definitiveTransactionType
	t.amount = amount

	return nil
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
