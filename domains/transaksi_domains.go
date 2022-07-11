package domains

import "WallE/models"

type TransaksiDomain interface {
	TransaksiBaru(transaksi models.Transaksi) error
	UpdateTransaksi(orderid string, transkasi models.Transaksi) error
	GetListTransactionByUserId(userid uint) []models.Transaksi
	GetProdukById(id uint) models.Produk
	GetUserById(id uint) models.User
	GetUserTransactions(id uint, filter string) []models.Transaksi
	ReduceBalance(id uint, balance int) error
	RefundBalance(id uint, balance int) error
	GetTransactionByOrderId(orderid string) models.Transaksi
	GetAllTransaction(filter string) []models.Transaksi
	GetLastId() uint
}

type TransaksiService interface {
	NewTransactionEWallet(transaksi models.Transaksi) (error, interface{})
	NewTransactionBank(transaksi models.Transaksi) (error, interface{})
	UpdateTransaksi(orderid string, transkasi models.Transaksi) error
	GetListTransactionByUserId(userid uint) []models.Transaksi
	GetUserTransactions(id uint, filter string) []models.Transaksi
	GetAllTransaction(filter string) []models.Transaksi
}
