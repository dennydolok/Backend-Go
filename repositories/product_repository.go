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

func (r *repositoriProduk) AmbilKategori() []models.Kategori {
	kategori := []models.Kategori{}
	r.DB.Find(&kategori).Association("Balance")
	return kategori
}

func (r *repositoriProduk) TambahSaldo(saldobaru int, kategoriid uint) error {
	saldo := models.Saldo{}
	r.DB.Where("kategori_id = ?", kategoriid).Find(&saldo)
	saldobaru += saldo.Saldo
	err := r.DB.Model(&saldo).Update("saldo", saldobaru).Where("kategori_id = ?", kategoriid).Error
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func (r *repositoriProduk) AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk {
	produk := []models.Produk{}
	r.DB.Where("kategori_id = ?", kategoriid).Preload(clause.Associations).Preload("Provider").Preload("Kategori").Find(&produk)
	return produk
}

func (r *repositoriProduk) AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk {
	produk := []models.Produk{}
	r.DB.Where("kategori_id = ?", kategoriid).Where("provider_id = ?", providerid).Preload("Provider").Preload("Kategori").Preload(clause.Associations).Find(&produk)
	return produk
}

func (r *repositoriProduk) TambahProduk(produk models.Produk) error {
	err := r.DB.Create(&produk).Error
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func (r *repositoriProduk) AmbilSaldo() []models.Saldo {
	saldo := []models.Saldo{}
	r.DB.Preload(clause.Associations).Preload("Kategori").Find(&saldo)
	return saldo
}

func (r *repositoriProduk) AmbilProviderBerdasarkanKategori(kategoriid uint) interface{} {
	type Result struct {
		ID   string
		Nama string
	}
	var result []Result
	r.DB.Raw("SELECT pr.nama AS Nama, provider_id AS ID, kategori_id FROM produks AS p JOIN providers as pr ON p.provider_id = pr.id JOIN kategoris c ON p.kategori_id = c.id WHERE kategori_id = ? GROUP BY id", kategoriid).Scan(&result)
	return result
}

func NewProductRepository(db *gorm.DB) domains.ProductDomain {
	return &repositoriProduk{
		DB: db,
	}
}
