package transactions

import (
	"context"
	"fmt"
	"transactions/domain/entities"
)

func (r Repository) InsertTransaction(transaction entities.Transaction) error {
	query := `INSERT INTO transactions (id, related_transaction_id, account_owner_document, transaction_type, amount) VALUES ($1, $2, $3, $4, $5)`

	row, err := r.Conn.Exec(
		context.Background(),
		query,
		transaction.ID,
		transaction.RelatedTransactionID,
		transaction.AccountOwnerDocument,
		transaction.TransactionType,
		transaction.Amount,
	)

	if err != nil {
		return err
	}

	if row.RowsAffected() < 1 {
		return fmt.Errorf("could not insert transaction %s", transaction.ID())
	}

	return nil
}
