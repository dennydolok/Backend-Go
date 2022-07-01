package repositories

import (
	"WallE/domains"
	"WallE/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err := r.DB.Where("order_id = ?", orderid).Updates(&transkasi).Error
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

func (r *RepositoryTransaksi) GetUserTransactions(id uint, filter string) []models.Transaksi {
	transactions := []models.Transaksi{}
	if filter == "berhasil" {
		r.DB.Where("user_id = ?", id).Where("status = ?", "settlement").Preload(clause.Associations).Find(&transactions)
	} else if filter == "tertunda" {
		r.DB.Where("user_id = ?", id).Where("status = ?", "pending").Preload(clause.Associations).Find(&transactions)
	} else if filter == "gagal" {
		r.DB.Where("user_id = ?", id).Where("status = ?", "cancel").Preload(clause.Associations).Find(&transactions)
	} else {
		r.DB.Preload(clause.Associations).Preload("Produk."+clause.Associations).Where("user_id = ?", id).Find(&transactions)
	}
	return transactions
}

func (r *RepositoryTransaksi) GetProdukById(id uint) models.Produk {
	produk := models.Produk{}
	r.DB.Where("id = ?", id).Preload(clause.Associations).Find(&produk)
	return produk
}

func (r *RepositoryTransaksi) GetUserById(id uint) models.User {
	user := models.User{}
	r.DB.Where("id = ? ", id).Preload(clause.Associations).Find(&user)
	return user
}

func (r *RepositoryTransaksi) GetLastId() uint {
	transaksi := models.Transaksi{}
	r.DB.Last(&transaksi)
	if transaksi.ID == 0 {
		return 1
	}
	return (transaksi.ID + 1)
}

func NewTransaksiRepository(db *gorm.DB) domains.TransaksiDomain {
	return &RepositoryTransaksi{
		DB: db,
	}
}
