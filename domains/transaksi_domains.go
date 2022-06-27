package domains

import "WallE/models"

type TransaksiDomain interface {
	TransaksiBaru(transaksi models.Transaksi) error
	UpdateTransaksi(orderid string, transkasi models.Transaksi) error
	GetListTransactionByUserId(userid uint) []models.Transaksi
}

type TransaksiService interface {
	TransaksiBaru(transaksi models.Transaksi) error
	UpdateTransaksi(orderid string, transkasi models.Transaksi) error
	GetListTransactionByUserId(userid uint) []models.Transaksi
}
