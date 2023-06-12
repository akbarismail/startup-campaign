package transaction

import (
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
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
}
