package transaction

import (
	"time"
)

type campaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func CampaignTransactionFormat(transaction Transaction) campaignTransactionFormatter {
	formatter := campaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func CampaignTransactionsFormat(transactions []Transaction) []campaignTransactionFormatter {
	if len(transactions) == 0 {
		return []campaignTransactionFormatter{}
	}

	formatCampaignTransactions := []campaignTransactionFormatter{}
	for _, campaignTransaction := range transactions {
		formatCampaignTransaction := CampaignTransactionFormat(campaignTransaction)
		formatCampaignTransactions = append(formatCampaignTransactions, formatCampaignTransaction)
	}

	return formatCampaignTransactions
}

type userTransactionFormatter struct {
	ID        int                              `json:"id"`
	Amount    int                              `json:"amount"`
	Status    string                           `json:"status"`
	CreatedAt time.Time                        `json:"created_at"`
	Campaign  campaignUserTransactionFormatter `json:"campaign"`
}

type campaignUserTransactionFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func UserTransactionFormat(transaction Transaction) userTransactionFormatter {
	formatter := userTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	formatterCampaign := campaignUserTransactionFormatter{}
	formatterCampaign.Name = transaction.Campaign.Name
	formatterCampaign.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		formatterCampaign.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = formatterCampaign

	return formatter
}

func UserTransactionsFormat(transactions []Transaction) []userTransactionFormatter {
	if len(transactions) == 0 {
		return []userTransactionFormatter{}
	}

	formatUserTransactions := []userTransactionFormatter{}
	for _, transaction := range transactions {
		formatUserTransaction := UserTransactionFormat(transaction)
		formatUserTransactions = append(formatUserTransactions, formatUserTransaction)
	}

	return formatUserTransactions
}

type transactionFormatter struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id"`
	CampaignId int    `json:"campaign_id"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	Amount     int    `json:"amount"`
	PaymentUrl string `json:"payment_url"`
}

func TransactionFormat(transaction Transaction) transactionFormatter {
	formatter := transactionFormatter{
		ID:         transaction.ID,
		UserId:     transaction.UserId,
		CampaignId: transaction.CampaignId,
		Status:     transaction.Status,
		Code:       transaction.Code,
		Amount:     transaction.Amount,
		PaymentUrl: transaction.PaymentUrl,
	}

	return formatter
}
