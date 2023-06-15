package transaction

import (
	"startup-campaign/campaign"
	"startup-campaign/user"
	"time"
)

type Transaction struct {
	ID         int
	UserId     int
	CampaignId int
	Status     string
	Code       string
	Amount     int
	PaymentUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
	Campaign   campaign.Campaign
}
