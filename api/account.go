package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/scortier/gopherbank/db/sqlc"
)

// CreateAccountRequest contains the input parameters for account creation
type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// createAccount is a handler function that creates a new account
func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest

	// BindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
	if err := ctx.BindJSON(&req); err != nil { // bind request to req
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error response
		return
	}

	// CreateAccount creates a new account in the database
	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
	})
	if err != nil { // if error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // return error response
		return
	}

	// return account as response
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAccount is a handler function that returns a specific account
func (server *Server) getAccount(ctx *gin.Context) {
	// get account id from url
	var req getAccountRequest
	// BindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
	if err := ctx.ShouldBindUri(&req); err != nil { // bind request to req
		ctx.JSON(http.StatusNotFound, errorResponse(err)) // return error response
		return
	}

	// GetAccount returns a specific account from the database
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil { // if error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // return error response
		return
	}

	// return account as response
	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// listAccounts is a handler function that returns a list of accounts
func (server *Server) listAccounts(ctx *gin.Context) {
	// get page_id and page_size from url
	var req listAccountsRequest
	// BindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
	if err := ctx.ShouldBindQuery(&req); err != nil { // bind request to req
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // return error response
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,                    // no of records that db should return
		Offset: (req.PageID - 1) * req.PageSize, // no of records that db should skip
	}
	fmt.Println(arg)

	// ListAccounts returns a list of accounts from the database
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil { // if error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err)) // return error response
		return
	}
	fmt.Println(accounts)
	// return accounts as response
	ctx.JSON(http.StatusOK, accounts)
}
