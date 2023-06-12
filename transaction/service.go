package transaction

import (
	"errors"
	"startup-campaign/campaign"
)

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	repositoryCampaign campaign.Repository
}

func NewService(repository Repository, repositoryCampaign campaign.Repository) *service {
	return &service{repository, repositoryCampaign}
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
