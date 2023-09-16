package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/scortier/gopherbank/db/sqlc"
)

type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

// createTransfer is a handler function that creates a new transfer
func (server *Server) createTransfer(ctx *gin.Context) {
	var req TransferRequest

	// BindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
	if err := ctx.ShouldBindJSON(&req); err != nil { // bind request to req
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error response
		return
	}

	// check if accounts currencies are valid
	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	// CreateTransfer creates a new transfer in the database
	transfer, err := server.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil { // if error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // return error response
		return
	}

	// return transfer as response
	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil { // if error
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err)) // return error response
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // return error response
		}
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error response
		return false
	}

	return true
}
