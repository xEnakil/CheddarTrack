package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/xenakil/cheddartrack/internal/model"
    "github.com/xenakil/cheddartrack/internal/service"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func NewTransactionHanlder(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{s}
}

func (h *TransactionHandler) RegisterRouter(r *gin.Engine) {
	txns := r.Group("/transactions") 
	{
		txns.POST("", h.CreateTransaction)
		txns.GET("/:user_id", h.GetTransactionByUser)
	}
}

// CreateTransaction godoc
// @Summary Create a transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param input body model.CreateTransactionInput true "Transaction Input"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input model.CreateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message":"transaction created"})
}

// GetTransactionByUser godoc
// @Summary Get transactions by user
// @Tags Transactions
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} model.TransactionResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{user_id} [get]
func (h *TransactionHandler) GetTransactionByUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid user_id"})
		return
	}

	txns, err := h.svc.GetAll(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, txns)
}

