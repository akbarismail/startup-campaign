package campaign

import (
	"startup-campaign/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	Description      string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	CountBacker      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
