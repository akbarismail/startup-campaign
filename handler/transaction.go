package handler

import (
	"net/http"
	"startup-campaign/helper"
	"startup-campaign/transaction"
	"startup-campaign/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	// parameter di uri
	// tangkap paramater kemudian mapping ke input struct
	// input struct akan di passing ke service
	// service akan memanggil repo
	// repo akan mencari data sesuai dengan id=x

	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignId(input)
	if err != nil {
		response := helper.APIResponse("Failed to campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's Transactions", http.StatusOK, "success", transaction.CampaignTransactionsFormat(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransactionsByUser(c *gin.Context) {
	// GetUserTransactions
	// handler
	// ambil nilai user dari jwt/middleware
	// service
	// repo => ambil data transactions (preload campaign)

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserId(userId)
	if err != nil {
		response := helper.APIResponse("Failed to user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Transactions", http.StatusOK, "success", transaction.UserTransactionsFormat(transactions))
	c.JSON(http.StatusOK, response)
}
