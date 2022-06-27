package domains

import "WallE/models"

type ProductDomain interface {
	GetKategori() []models.Kategori
	AddSaldo(saldobaru int, kategoriid uint) error
	GetSaldo() []models.Saldo
	GetProdukById(id uint) models.Produk
	GetProdukByKategori(kategoriid uint) []models.Produk
	GetProdukByKategoriProvider(kategoriid, providerid uint) []models.Produk
	GetPurchaseableProduct(kategoriid, providerid uint) interface{}
	AddProduct(produk models.Produk) error
	GetProviderByKategori(kategoriid uint) interface{}
	UpdateProductById(id uint, produk models.Produk) error
}

type ProductService interface {
	GetKategori() []models.Kategori
	AddSaldo(saldobaru int, kategoriid uint) error
	GetProdukById(id uint) models.Produk
	GetProdukByKategori(kategoriid uint) []models.Produk
	GetProdukByKategoriProvider(kategoriid, providerid uint) []models.Produk
	AddProduct(produk models.Produk) error
	GetPurchaseableProduct(kategoriid, providerid uint) interface{}
	GetSaldo() []models.Saldo
	GetProviderByKategori(kategoriid uint) interface{}
	UpdateProductById(id uint, produk models.Produk) error
}
