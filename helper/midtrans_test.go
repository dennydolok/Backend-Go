package helper

import (
	"testing"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/assert"
)

var chargeBank = coreapi.ChargeResponse{
	TransactionTime: "2022-07-11 14:34:07",
	TransactionID:   "e64903d5-415d-4671-b499-7134824a0b19",
	OrderID:         "INV/11072022/PLS/111",
	VaNumbers: []coreapi.VANumber{
		{
			Bank:     "bca",
			VANumber: "12313213123",
		},
	},
}

var chargeBankPermata = coreapi.ChargeResponse{
	TransactionTime: "2022-07-11 14:34:07",
	TransactionID:   "e64903d5-415d-4671-b499-7134824a0b19",
	OrderID:         "INV/11072022/PLS/111",
	Bank:            "peramata",
	PermataVaNumber: "1231421321",
}

var chargeGopay = coreapi.ChargeResponse{
	TransactionTime: "2022-07-11 14:34:07",
	TransactionID:   "e64903d5-415d-4671-b499-7134824a0b19",
	OrderID:         "INV/11072022/PLS/111",
	Actions: []coreapi.Action{
		{
			Name:   "generate-qr-code",
			Method: "GET",
			URL:    "https://api.sandbox.veritrans.co.id/v2/gopay/231c79c5-e39e-4993-86da-cadcaee56c1d/qr-code",
		},
		{
			Name:   "deeplink-redirect",
			Method: "GET",
			URL:    "https://api.sandbox.veritrans.co.id/v2/gopay/231c79c5-e39e-4993-86da-cadcaee56c1d/qr-code",
		},
		{
			Name:   "get-status",
			Method: "GET",
			URL:    "https://api.sandbox.veritrans.co.id/v2/gopay/231c79c5-e39e-4993-86da-cadcaee56c1d/qr-code",
		},
		{
			Name:   "cancel",
			Method: "POST",
			URL:    "https://api.sandbox.veritrans.co.id/v2/231c79c5-e39e-4993-86da-cadcaee56c1d/cancel",
		},
	},
}

func TestFromMidBank(t *testing.T) {
	t.Run("bank non permata", func(t *testing.T) {
		result := FromMidBank(chargeBank, uint(1), "pulsa axis", "083811992244", "bca", int64(12000), int64(10000))
		assert.NotEmpty(t, result)
	})
	t.Run("bank permata", func(t *testing.T) {
		result := FromMidBank(chargeBankPermata, uint(1), "pulsa axis", "083811992244", "permata", int64(12000), int64(10000))
		assert.NotEmpty(t, result)
	})
}

func TestFromWallet(t *testing.T) {
	result := FromMidEWallet(chargeGopay, uint(1), "pulsa axis", "083811991222", int64(12000), int64(10000))
	assert.NotEmpty(t, result)
}
