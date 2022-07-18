package api

type transferRequest struct {
	FromAccountID int64  `json:"from_account" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=1"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

// func (server *Server) createTransfer(ctx *gin, Context) {
// 	var req transferRequest
// 	if err:= ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	}

// }
