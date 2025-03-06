package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"transactions/domain/transactions"

	"github.com/go-chi/render"
)

type MakeDepositUC interface {
	ExecuteDeposit(input transactions.ExecuteDepositInput) error
}

func MakeDepositHandler(uc MakeDepositUC) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input transactions.ExecuteDepositInput
		input.UserAccountOwnerDocument = r.FormValue("destination_document")
		amountStr := r.FormValue("amount")
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, fmt.Errorf("error while making deposit: %w", err))
			return
		}

		input.Amount = amount

		err = uc.ExecuteDeposit(input)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, fmt.Errorf("error while making deposit: %w", err))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, "deposit made")
	}
}
