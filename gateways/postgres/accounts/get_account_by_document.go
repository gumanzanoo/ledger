package accounts

import (
	"context"
	"fmt"
	"transactions/domain/entities"
	"transactions/domain/vo"
)

func (r Repository) GetAccountByDocument(document vo.AccountOwnerDocument) (entities.Account, error) {
	query := `SELECT owner_document FROM accounts WHERE owner_document = $1`

	var (
		id            string
		ownerDocument string
	)

	if err := r.Conn.QueryRow(context.Background(), query, document.Value()).Scan(&id, &ownerDocument); err != nil {
		return entities.Account{}, fmt.Errorf("could not find account with owner document %s", document)
	}

	var account entities.Account

	if err := account.ParseAccount(id, ownerDocument); err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
