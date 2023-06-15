package transaction

import (
	"errors"
	"startup-campaign/campaign"
	"startup-campaign/payment"
	"strconv"
)

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository         Repository
	repositoryCampaign campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, repositoryCampaign campaign.Repository, payment payment.Service) *service {
	return &service{repository, repositoryCampaign, payment}
}

func (s *service) GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.repositoryCampaign.FindById(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.FindByCampaignId(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignId: input.CampaignId,
		Amount:     input.Amount,
		UserId:     input.User.ID,
		Status:     "pending",
	}

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, err := strconv.Atoi(input.OrderId)
	if err != nil {
		return err
	}

	transaction, err := s.repository.FindById(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}

	updateTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.repositoryCampaign.FindById(updateTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid" {
		campaign.CountBacker = campaign.CountBacker + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updateTransaction.Amount

		_, err := s.repositoryCampaign.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
