package domains

import "WallE/models"

type ProductDomain interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilSaldo() []models.Saldo
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	AmbilProdukBisaDibeli(kategoriid, providerid uint) interface{}
	TambahProduk(produk models.Produk) error
	AmbilProviderBerdasarkanKategori(kategoriid uint) interface{}
}

type ProductService interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	TambahProduk(produk models.Produk) error
	AmbilProdukBisaDibeli(kategoriid, providerid uint) interface{}
	AmbilSaldo() []models.Saldo
	AmbilProviderBerdasarkanKategori(kategoriid uint) interface{}
}
