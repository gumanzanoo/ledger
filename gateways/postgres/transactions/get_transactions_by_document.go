package transactions

import (
	"context"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

func (r Repository) GetTransactionsByDocument(document vo.AccountOwnerDocument) ([]entities.Transaction, error) {
	query := `SELECT id, related_transaction_id, account_owner_document, transaction_type, amount FROM transactions WHERE account_owner_document = $1`

	rows, err := r.Conn.Query(context.Background(), query, document.Value())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entities.Transaction

	for rows.Next() {
		var (
			id                   string
			relatedTransactionID string
			accountOwnerDocument string
			transactionType      string
			amount               int
		)

		err := rows.Scan(&id, &relatedTransactionID, &accountOwnerDocument, &transactionType, &amount)
		if err != nil {
			return nil, err
		}

		var transaction entities.Transaction
		if err := transaction.ParseTransaction(
			id, relatedTransactionID, accountOwnerDocument, transactionType, amount); err != nil {
			return []entities.Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
