package transaction

import "time"

type Transaction struct {
	ID         int
	UserId     int
	CampaignId int
	Status     string
	Code       int
	Amount     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
