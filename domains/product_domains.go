package domains

import "WallE/models"

type ProductDomain interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	TambahProduk(produk models.Produk) error
}

type ProductService interface {
	AmbilKategori() []models.Kategori
	TambahSaldo(saldobaru int, kategoriid uint) error
	AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk
	AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk
	TambahProduk(produk models.Produk) error
}
