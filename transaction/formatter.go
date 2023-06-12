package transaction

import "time"

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
