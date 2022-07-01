package services

import (
	"WallE/domains"
	"WallE/models"
	"time"
)

type serviceProduk struct {
	repo domains.ProductDomain
}

func (s *serviceProduk) AddProduct(produk models.Produk) error {
	produk.DibuatPada = time.Now()
	produk.DiupdatePada = time.Now()
	return s.repo.AddProduct(produk)
}

func (s *serviceProduk) GetProdukByKategoriProvider(kategoriid, providerid uint) []models.Produk {
	produk := []models.Produk{}
	produk = s.repo.GetProdukByKategoriProvider(kategoriid, providerid)
	return produk
}

func (s *serviceProduk) GetProdukByKategori(kategoriid uint) []models.Produk {
	produk := []models.Produk{}
	produk = s.repo.GetProdukByKategori(kategoriid)
	return produk
}

func (s *serviceProduk) AddSaldo(saldobaru int, kategoriid uint) error {
	err := s.repo.AddSaldo(saldobaru, kategoriid)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceProduk) GetKategori() []models.Kategori {
	kategori := []models.Kategori{}
	kategori = s.repo.GetKategori()
	return kategori
}

func (s *serviceProduk) GetSaldo() []models.Saldo {
	return s.repo.GetSaldo()
}

func (s *serviceProduk) GetProdukById(id uint) models.Produk {
	return s.repo.GetProdukById(id)
}

func (s *serviceProduk) GetProviderByKategori(kategoriid uint) interface{} {
	return s.repo.GetProviderByKategori(kategoriid)
}

func (s *serviceProduk) GetPurchaseableProduct(kategoriid, providerid uint) interface{} {
	return s.repo.GetPurchaseableProduct(kategoriid, providerid)
}

func (s *serviceProduk) UpdateProductById(id uint, produk models.Produk) error {
	produk.KategoriID = s.repo.GetProdukById(uint(id)).KategoriID
	produk.ProviderID = s.repo.GetProdukById(uint(id)).ProviderID
	return s.repo.UpdateProductById(id, produk)
}

func (s *serviceProduk) DeleteProdukById(id uint) error {
	return s.repo.DeleteProdukById(id)
}

func NewProdukService(repo domains.ProductDomain) domains.ProductService {
	return &serviceProduk{
		repo: repo,
	}
}
