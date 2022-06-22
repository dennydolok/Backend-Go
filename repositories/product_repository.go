package repositories

import (
	"WallE/domains"
	"WallE/models"
	"errors"
	"fmt"

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
		ID   uint   `json:"id"`
		Nama string `json:"nama"`
	}
	var result []Result
	r.DB.Raw("SELECT pr.nama AS Nama, provider_id AS ID FROM produks AS p JOIN providers as pr ON p.provider_id = pr.id JOIN kategoris c ON p.kategori_id = c.id WHERE kategori_id = ? GROUP BY id", kategoriid).Scan(&result)
	return result
}

func (r *repositoriProduk) AmbilProdukBisaDibeli(kategoriid, providerid uint) interface{} {
	type Result struct {
		ID            uint   `json:"id"`
		Nama          string `json:"nama"`
		Nominal       int    `json:"nominal"`
		Harga         int    `json:"harga"`
		Nama_kategori string `json:"nama_kategori"`
		Nama_provider string `json:"nama_provider"`
		Tersedia      bool   `json:"tersedia"`
	}
	var result []Result
	r.DB.Raw("SELECT p.id AS id, p.nama AS nama, p.nominal AS nominal, p.harga AS harga, pr.nama AS nama_provider, c.nama AS nama_kategori, IF(p.nominal <= s.saldo, 1, 0) AS tersedia FROM produks AS p JOIN providers AS pr ON p.provider_id = pr.id JOIN kategoris c ON p.kategori_id = c.id JOIN saldos s ON s.kategori_id = c.id WHERE p.kategori_id = ? AND p.provider_id = ? GROUP BY p.id;", kategoriid, providerid).Scan(&result)
	fmt.Println(result, providerid, kategoriid)
	return result
}

func NewProductRepository(db *gorm.DB) domains.ProductDomain {
	return &repositoriProduk{
		DB: db,
	}
}
