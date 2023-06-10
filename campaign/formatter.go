package campaign

type campaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmout        int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func campaignFormat(campaign Campaign) campaignFormatter {
	formatter := campaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmout:        campaign.GoalAmout,
		CurrentAmount:    campaign.CurrentAmount,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func CampaignsFormat(campaigns []Campaign) []campaignFormatter {
	formatCampaigns := []campaignFormatter{}

	for _, campaign := range campaigns {
		formatCampaign := campaignFormat(campaign)
		formatCampaigns = append(formatCampaigns, formatCampaign)
	}

	return formatCampaigns
}
