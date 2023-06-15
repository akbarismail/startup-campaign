package campaign

import "strings"

type campaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func CampaignFormat(campaign Campaign) campaignFormatter {
	formatter := campaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func CampaignsFormat(campaigns []Campaign) []campaignFormatter {
	formatCampaigns := []campaignFormatter{}

	for _, campaign := range campaigns {
		formatCampaign := CampaignFormat(campaign)
		formatCampaigns = append(formatCampaigns, formatCampaign)
	}

	return formatCampaigns
}

type campaignDetailFormatter struct {
	ID               int                      `json:"id"`
	UserId           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CountBacker      int                      `json:"count_backer"`
	CurrentAmount    int                      `json:"current_amount"`
	Description      string                   `json:"description"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             campaignUserFormatter    `json:"user"`
	Images           []campaignImageFormatter `json:"images"`
}

type campaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type campaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignDetailFormat(campaign Campaign) campaignDetailFormatter {
	var perks []string
	perkSplit := strings.Split(campaign.Perks, ",")
	for _, perk := range perkSplit {
		perks = append(perks, strings.TrimSpace(perk))
	}

	formatter := campaignDetailFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Description:      campaign.Description,
		Slug:             campaign.Slug,
		Perks:            perks,
		CountBacker:      campaign.CountBacker,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	user := campaign.User

	formatterUser := campaignUserFormatter{
		Name:     user.Name,
		ImageUrl: user.AvatarFileName,
	}

	formatter.User = formatterUser

	images := []campaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		formatterImage := campaignImageFormatter{
			ImageUrl:  image.FileName,
			IsPrimary: isPrimary,
		}

		images = append(images, formatterImage)
	}

	formatter.Images = images

	return formatter
}
