package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"transactions/domain/transactions"

	"github.com/go-chi/render"
)

type MakeTransactionUC interface {
	ExecuteTransaction(input transactions.ExecuteTransactionInput) error
}

func MakeTransactionHandler(uc MakeTransactionUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input transactions.ExecuteTransactionInput
		input.UserOriginDocument = r.FormValue("origin_document")
		input.UserDestinationDocument = r.FormValue("destination_document")
		amountStr := r.FormValue("amount")
		amount, err := strconv.Atoi(amountStr)

		if err != nil {
			http.Error(w, "invalid amount", http.StatusBadRequest)
			return
		}

		input.Amount = amount

		err = uc.ExecuteTransaction(input)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, fmt.Errorf("error making transaction: %w", err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, "transaction made")
	}
}
