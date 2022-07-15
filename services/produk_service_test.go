package services

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var produkRepository m.ProductDomain

var produk = models.Produk{
	ID:           1,
	Nama:         "Axis 10k",
	Nominal:      10000,
	Harga:        12000,
	Deskripsi:    "Pulsa Axis 10k",
	KategoriID:   1,
	ProviderID:   1,
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
}

var kategori = models.Kategori{
	ID:   1,
	Nama: "Pulsa",
}

var saldo = models.Saldo{
	ID:           1,
	Saldo:        1000,
	KategoriID:   1,
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
}

func TestAddProduct(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success add product", func(t *testing.T) {
		produkRepository.On("AddProduct", mock.Anything).Return(nil).Once()
		err := produkService.AddProduct(produk)
		assert.NoError(t, err)
	})
}

func TestGetProdukByKategoriProvider(t *testing.T) {
	products := []models.Produk{
		produk,
	}
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success get produk", func(t *testing.T) {
		produkRepository.On("GetProdukByKategoriProvider", mock.Anything, mock.Anything).Return(products).Once()
		products := produkService.GetProdukByKategoriProvider(1, 1)
		assert.NotEmpty(t, products)
	})
}

func TestGetProdukByKategori(t *testing.T) {
	products := []models.Produk{
		produk,
	}
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success get produk", func(t *testing.T) {
		produkRepository.On("GetProdukByKategori", mock.Anything).Return(products).Once()
		products := produkService.GetProdukByKategori(1)
		assert.NotEmpty(t, products)
	})
}

func TestAddSaldo(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success add saldo", func(t *testing.T) {
		produkRepository.On("AddSaldo", mock.Anything, mock.Anything).Return(nil).Once()
		err := produkService.AddSaldo(10000, 1)
		assert.NoError(t, err)
	})
	t.Run("failed add saldo", func(t *testing.T) {
		produkRepository.On("AddSaldo", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()
		err := produkService.AddSaldo(10000, 1)
		assert.Error(t, err)
	})
}

func TestGetKategori(t *testing.T) {
	kategori := []models.Kategori{
		kategori,
	}
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success get kategori", func(t *testing.T) {
		produkRepository.On("GetKategori").Return(kategori).Once()
		kategori := produkService.GetKategori()
		assert.NotEmpty(t, kategori)
	})
}

func TestGetSaldo(t *testing.T) {
	saldo := []models.Saldo{
		saldo,
	}
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success get saldo", func(t *testing.T) {
		produkRepository.On("GetSaldo").Return(saldo).Once()
		saldo := produkService.GetSaldo()
		assert.NotEmpty(t, saldo)
	})
}

func TestGetProdukById(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success get produk", func(t *testing.T) {
		produkRepository.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
		produk = produkService.GetProdukById(uint(1))
		assert.NotEmpty(t, produk)
	})
}
func TestGetProviderByKategori(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	type Result struct {
		ID   uint   `json:"id"`
		Nama string `json:"nama"`
	}
	res := Result{
		ID:   1,
		Nama: "Axis",
	}
	result := []Result{
		res,
	}
	t.Run("success get provider", func(t *testing.T) {
		produkRepository.On("GetProviderByKategori", mock.AnythingOfType("uint")).Return(result).Once()
		provider := produkService.GetProviderByKategori(uint(1))
		assert.NotEmpty(t, provider)
	})
}

func TestGetPurchaseableProduct(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	type Result struct {
		ID            uint   `json:"id"`
		Nama          string `json:"nama"`
		Nominal       int    `json:"nominal"`
		Harga         int    `json:"harga"`
		Deskripsi     string `json:"deskripsi"`
		Nama_kategori string `json:"nama_kategori"`
		Nama_provider string `json:"nama_provider"`
		Tersedia      bool   `json:"tersedia"`
	}
	produk := Result{
		ID:            1,
		Nama:          "Axis 10k",
		Nominal:       10000,
		Harga:         12000,
		Deskripsi:     "Pulsa Axis 10k",
		Nama_kategori: "Pulsa",
		Nama_provider: "Axis",
		Tersedia:      true,
	}
	produks := []Result{
		produk,
	}
	t.Run("success get products", func(t *testing.T) {
		produkRepository.On("GetPurchaseableProduct", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(produks).Once()
		provider := produkService.GetPurchaseableProduct(uint(1), uint(1))
		assert.NotEmpty(t, provider)
	})
}

func TestUpdateProductById(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success update products", func(t *testing.T) {
		produkRepository.On("UpdateProductById", mock.AnythingOfType("uint"), produk).Return(nil).Once()
		err := produkService.UpdateProductById(uint(1), produk)
		assert.NoError(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	produkService := serviceProduk{
		repo: &produkRepository,
	}
	t.Run("success delete products", func(t *testing.T) {
		produkRepository.On("DeleteProdukById", mock.AnythingOfType("uint")).Return(nil).Once()
		err := produkService.DeleteProdukById(uint(1))
		assert.NoError(t, err)
	})
}

func TestNewProdukServi(t *testing.T) {
	result := NewProdukService(&produkRepository)
	assert.NotEmpty(t, result)
}
