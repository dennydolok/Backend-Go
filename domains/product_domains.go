package domains

import "WallE/models"

type ProductDomain interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilSaldo() []models.Saldo
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	TambahProduk(produk models.Produk) error
	AmbilProviderBerdasarkanKategori(kategoriid uint) interface{}
}

type ProductService interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	TambahProduk(produk models.Produk) error
	AmbilSaldo() []models.Saldo
	AmbilProviderBerdasarkanKategori(kategoriid uint) interface{}
}
