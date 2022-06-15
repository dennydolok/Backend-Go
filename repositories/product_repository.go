package repositories

import (
	"WallE/domains"
	"WallE/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repositoriProduk struct {
	DB *gorm.DB
}

func (r *repositoriProduk) AmbilKategori() []models.Kategory {
	kategori := []models.Kategory{}
	r.DB.Find(&kategori).Association("Balance")
	return kategori
}

func (r *repositoriProduk) TambahSaldo(saldobaru int, kategoriid uint) error {
	saldo := models.Saldo{}
	r.DB.Find(&saldo).Where("kategori_id = ?", kategoriid)
	err := r.DB.Model(&saldo).Update("saldo", saldo.Saldo+saldobaru).Error
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func (r *repositoriProduk) AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk {
	produk := []models.Produk{}
	r.DB.Find(&produk).Preload(clause.Associations).Where("ketogori_id = ?", kategoriid)
	return produk
}

func NewProductRepository(db *gorm.DB) domains.ProductDomain {
	return &repositoriProduk{
		DB: db,
	}
}
