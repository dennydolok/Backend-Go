package services

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"errors"
	"testing"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transaksiRepositori m.TransaksiDomain

var transaksi = models.Transaksi{
	ID:         1,
	Status:     "pending",
	TotalHarga: "12000",
	OrderID:    "INV/14072022/PLS/160",
	Bank:       "bca",
	UserID:     1,
	ProdukID:   1,
}

var TransaksiWallet = models.Transaksi{
	ID:             1,
	Status:         "pending",
	TotalHarga:     "12000",
	OrderID:        "INV/14072022/PLS/160",
	TipePembayaran: "gopay",
	UserID:         1,
	ProdukID:       1,
}

func TestTransaksiBank(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success transaction", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "bank_transfer",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
				VaNumbers: []coreapi.VANumber{
					{
						Bank:     "bca",
						VANumber: "12313213123",
					},
				},
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(transaksi, nil).Once()
		transaksiRepositori.On("ReduceBalance", mock.AnythingOfType("uint"), mock.Anything).Return(nil).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.NoError(t, err)
	})
	t.Run("error no product", func(t *testing.T) {
		emptyProduk := models.Produk{
			ID: 0,
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(emptyProduk).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.Error(t, err)
	})
	t.Run("error no user", func(t *testing.T) {
		emptyUser := models.User{
			ID: 0,
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(emptyUser).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.Error(t, err)
	})
	t.Run("error midtrans", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
					TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
					PaymentType:       "bank_transfer",
					OrderID:           "INV/02072022/VCG/2",
					GrossAmount:       "12000",
					TransactionTime:   "2022-07-02 19:19:57",
					TransactionStatus: "pending",
					VaNumbers: []coreapi.VANumber{
						{
							Bank:     "bca",
							VANumber: "12313213123",
						},
					},
				}, &midtrans.Error{
					Message:    "error",
					StatusCode: 500,
					RawError:   errors.New("error midtrans"),
				}
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.Error(t, err)
	})
	t.Run("error transaction database", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "bank_transfer",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
				VaNumbers: []coreapi.VANumber{
					{
						Bank:     "bca",
						VANumber: "12313213123",
					},
				},
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(transaksi, errors.New("error database")).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.Error(t, err)
	})
	t.Run("error reduce balance", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "bank_transfer",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
				VaNumbers: []coreapi.VANumber{
					{
						Bank:     "bca",
						VANumber: "12313213123",
					},
				},
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(transaksi, nil).Once()
		transaksiRepositori.On("ReduceBalance", mock.AnythingOfType("uint"), mock.Anything).Return(errors.New("error database")).Once()
		err, _ := transaksiService.NewTransactionBank(transaksi)
		assert.Error(t, err)
	})
}

func TestTransaksiEWallet(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success transaction", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "gopay",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
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
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(TransaksiWallet, nil).Once()
		transaksiRepositori.On("ReduceBalance", mock.AnythingOfType("uint"), mock.Anything).Return(nil).Once()
		err, _ := transaksiService.NewTransactionEWallet(TransaksiWallet)
		assert.NoError(t, err)
	})
	t.Run("error no product", func(t *testing.T) {
		emptyProduk := models.Produk{
			ID: 0,
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(emptyProduk).Once()
		err, _ := transaksiService.NewTransactionEWallet(TransaksiWallet)
		assert.Error(t, err)
	})
	t.Run("error no user", func(t *testing.T) {
		emptyUser := models.User{
			ID: 0,
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(emptyUser).Once()
		err, _ := transaksiService.NewTransactionEWallet(transaksi)
		assert.Error(t, err)
	})
	t.Run("error midtrans", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
					TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
					PaymentType:       "gopay",
					OrderID:           "INV/02072022/VCG/2",
					GrossAmount:       "12000",
					TransactionTime:   "2022-07-02 19:19:57",
					TransactionStatus: "pending",
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
				}, &midtrans.Error{
					Message:    "error",
					StatusCode: 500,
					RawError:   errors.New("error midtrans"),
				}
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		err, _ := transaksiService.NewTransactionEWallet(TransaksiWallet)
		assert.Error(t, err)
	})
	t.Run("error reduce balance", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "gopay",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
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
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(TransaksiWallet, nil).Once()
		transaksiRepositori.On("ReduceBalance", mock.AnythingOfType("uint"), mock.Anything).Return(errors.New("error database")).Once()
		err, _ := transaksiService.NewTransactionEWallet(TransaksiWallet)
		assert.Error(t, err)
	})
	t.Run("error transaction database", func(t *testing.T) {
		defer func() { ChargeTransaction = oldChargeTransaction }()
		ChargeTransaction = func(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
			return &coreapi.ChargeResponse{
				TransactionID:     "1661865e-7ce8-4e23-898b-9174b01c89f2",
				PaymentType:       "gopay",
				OrderID:           "INV/02072022/VCG/2",
				GrossAmount:       "12000",
				TransactionTime:   "2022-07-02 19:19:57",
				TransactionStatus: "pending",
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
			}, nil
		}
		transaksiRepositori.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		transaksiRepositori.On("GetUserById", mock.AnythingOfType("uint")).Return(user).Once()
		transaksiRepositori.On("GetLastId").Return(uint(16)).Once()
		transaksiRepositori.On("TransaksiBaru", mock.Anything).Return(TransaksiWallet, errors.New("error database")).Once()
		err, _ := transaksiService.NewTransactionEWallet(TransaksiWallet)
		assert.Error(t, err)
	})
}

func TestUpdateTransaksi(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success update", func(t *testing.T) {
		transaksiRepositori.On("UpdateTransaksi", mock.Anything, transaksi).Return(nil).Once()
		err := transaksiService.UpdateTransaksi(transaksi.OrderID, transaksi)
		assert.NoError(t, err)
	})
	t.Run("error update", func(t *testing.T) {
		transaksiRepositori.On("UpdateTransaksi", mock.Anything, transaksi).Return(errors.New("error")).Once()
		err := transaksiService.UpdateTransaksi(transaksi.OrderID, transaksi)
		assert.Error(t, err)
	})
	t.Run("error update refund", func(t *testing.T) {
		transaksi.Status = "expire"
		transaksiRepositori.On("UpdateTransaksi", mock.Anything, transaksi).Return(nil).Once()
		transaksiRepositori.On("GetTransactionByOrderId", mock.Anything).Return(transaksi).Once()
		transaksiRepositori.On("RefundBalance", mock.AnythingOfType("uint"), mock.Anything).Return(nil).Once()
		err := transaksiService.UpdateTransaksi(transaksi.OrderID, transaksi)
		assert.NoError(t, err)
	})
	t.Run("error update refund", func(t *testing.T) {
		transaksi.Status = "expire"
		transaksiRepositori.On("UpdateTransaksi", mock.Anything, transaksi).Return(nil).Once()
		transaksiRepositori.On("GetTransactionByOrderId", mock.Anything).Return(transaksi).Once()
		transaksiRepositori.On("RefundBalance", mock.AnythingOfType("uint"), mock.Anything).Return(errors.New("error")).Once()
		err := transaksiService.UpdateTransaksi(transaksi.OrderID, transaksi)
		assert.Error(t, err)
	})
}

func TestGetUserTransactions(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success get transactions", func(t *testing.T) {
		transactions := []models.Transaksi{
			transaksi,
		}
		transaksiRepositori.On("GetUserTransactions", mock.AnythingOfType("uint"), mock.Anything).Return(transactions).Once()
		transactions = transaksiService.GetUserTransactions(1, "")
		assert.NotEmpty(t, transactions)
	})
}
func TestGetListTransactionByUserId(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success get transactions", func(t *testing.T) {
		transactions := []models.Transaksi{
			transaksi,
		}
		transaksiRepositori.On("GetListTransactionByUserId", mock.AnythingOfType("uint")).Return(transactions).Once()
		transactions = transaksiService.GetListTransactionByUserId(1)
		assert.NotEmpty(t, transactions)
	})
}

func TestGetAllTransactions(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success get transactions", func(t *testing.T) {
		transactions := []models.Transaksi{
			transaksi,
		}
		transaksiRepositori.On("GetAllTransaction", mock.Anything).Return(transactions).Once()
		transactions = transaksiService.GetAllTransaction("")
		assert.NotEmpty(t, transactions)
	})
}

func TestGetTransactionsById(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success get transactions", func(t *testing.T) {
		transaksiRepositori.On("GetTransactionById", mock.AnythingOfType("uint")).Return(transaksi).Once()
		result := transaksiService.GetTransactionById(1)
		assert.NotEmpty(t, result)
	})
}

func TestGetTotalIncome(t *testing.T) {
	transaksiService := serviceTransaksi{
		repo: &transaksiRepositori,
	}
	t.Run("success get transactions", func(t *testing.T) {
		transaksiRepositori.On("GetTotalIncome").Return(10000).Once()
		result := transaksiService.GetTotalIncome()
		assert.NotEmpty(t, result)
	})
}

func TestNewTransactionsService(t *testing.T) {
	result := NewTransaksiService(&transaksiRepositori)
	assert.NotEmpty(t, result)
}
