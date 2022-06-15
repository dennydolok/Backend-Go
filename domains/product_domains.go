package domains

import "WallE/models"

type ProductDomain interface {
	AmbilKategori() []models.Kategory
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
}

type ProductService interface {
	GetKategori() []models.Kategory
	IncreaseBalance(amount int, categoryid uint) error
}
