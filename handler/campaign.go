package handler

import (
	"fmt"
	"net/http"
	"startup-campaign/campaign"
	"startup-campaign/helper"
	"startup-campaign/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan repository mana yang di call
// repository: GetAll, GetByUserId
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.CampaignsFormat(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// handler: mapping id yang ada di url ke struct input => service, call formatter
	// service: struct input => menangkap id di url, manggil repository
	// repository: get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign.CampaignDetailFormat(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	// tangkap parameter dari user ke input struct
	// ambil current user dari jwt/handler
	// panggil service, parameternya input struct (dan juga dibuatkan slug)
	// panggil repository untuk simpan data campaign baru

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", campaign.CampaignFormat(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// user masukan input
	// handler: input user dan input uri passing ke service
	// service: kedua input dari user mapping ke input struct passing ke repo
	// repository: update data campaign

	var inputId campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.APIResponse("Failed to updated campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to updated campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User.ID = currentUser.ID

	updatedCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to updated campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated campaign", http.StatusOK, "success", campaign.CampaignFormat(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	// handler:
	// tangkap input dan ubah ke struct input
	// save image campaign ke suatu folder
	// service (kondisi manggil point 2 di repo, manggil point 1)
	// repository:
	// 1. create image/save image di table campaign_images
	// 2. ubah is_primary = true ke false

	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("campaign_image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Image successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
