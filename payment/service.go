package payment

import (
	"startup-campaign/user"
	"startup-campaign/utils"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()

	serverKey, _ := utils.EnvVariabel("SERVER_KEY")
	clientKey, _ := utils.EnvVariabel("CLIENT_KEY")

	midclient.ServerKey = serverKey
	midclient.ClientKey = clientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},

		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
