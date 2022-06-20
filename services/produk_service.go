package services

import (
	"WallE/domains"
	"WallE/models"
	"time"
)

type serviceProduk struct {
	repo domains.ProductDomain
}

func (s *serviceProduk) TambahProduk(produk models.Produk) error {
	produk.DibuatPada = time.Now()
	produk.DiupdatePada = time.Now()
	return s.repo.TambahProduk(produk)
}

func (s *serviceProduk) AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid uint) []models.Produk {
	produk := []models.Produk{}
	produk = s.repo.AmbilProdukBerdasarkanProviderKategori(kategoriid, providerid)
	return produk
}

func (s *serviceProduk) AmbilProdukBerdasarkanKategori(kategoriid uint) []models.Produk {
	produk := []models.Produk{}
	produk = s.repo.AmbilProdukBerdasarkanKategori(kategoriid)
	return produk
}

func (s *serviceProduk) TambahSaldo(saldobaru int, kategoriid uint) error {
	err := s.repo.TambahSaldo(saldobaru, kategoriid)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceProduk) AmbilKategori() []models.Kategori {
	kategori := []models.Kategori{}
	kategori = s.repo.AmbilKategori()
	return kategori
}

func (s *serviceProduk) AmbilSaldo() []models.Saldo {
	return s.repo.AmbilSaldo()
}

func (s *serviceProduk) AmbilProviderBerdasarkanKategori(kategoriid uint) interface{} {
	return s.repo.AmbilProviderBerdasarkanKategori(kategoriid)
}

func NewProdukService(repo domains.ProductDomain) domains.ProductService {
	return &serviceProduk{
		repo: repo,
	}
}
