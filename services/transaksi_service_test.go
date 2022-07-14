package services

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"errors"
	"testing"

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
