package repositories

import (
	"WallE/domains"
	"WallE/models"

	"gorm.io/gorm"
)

type RepositoryTransaksi struct {
	DB *gorm.DB
}

func (r *RepositoryTransaksi) TransaksiBaru(transaksi models.Transaksi) error {
	err := r.DB.Create(&transaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryTransaksi) UpdateTransaksi(orderid string, transkasi models.Transaksi) error {
	err := r.DB.Where("order_id = ?", orderid).Save(&orderid).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryTransaksi) GetListTransactionByUserId(userid uint) []models.Transaksi {
	transactions := []models.Transaksi{}
	r.DB.Where("user_id = ?", userid).Find(&transactions)
	return transactions
}

func NewTransaksiRepository(db *gorm.DB) domains.TransaksiDomain {
	return &RepositoryTransaksi{
		DB: db,
	}
}
