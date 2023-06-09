package campaign

import "time"

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	Description      string
	ShortDescription string
	GoalAmout        int
	CurrentAmount    int
	Perks            string
	CountBacker      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
