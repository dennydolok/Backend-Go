package repositories

import (
	"WallE/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var kategori = models.Kategori{
	ID:           1,
	Nama:         "Pulsa",
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
}

var produk = models.Produk{
	Nama:         "Pulsa Axis 10k",
	Nominal:      10000,
	Harga:        12000,
	Deskripsi:    "Pulsa Axis 10000",
	KategoriID:   1,
	ProviderID:   1,
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
	Dihapus:      gorm.DeletedAt{},
}

var saldo = models.Saldo{
	ID:           1,
	Saldo:        10000,
	KategoriID:   1,
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
}

var provider = models.Provider{
	ID:           1,
	Nama:         "Axis",
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
}

func TestGetKategori(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", kategori.Nama, kategori.DibuatPada.String(), kategori.DiupdatePada.String()}).
			AddRow(kategori.ID, kategori.Nama, kategori.DibuatPada, kategori.DiupdatePada))
	res := repo.GetKategori()
	assert.NotEmpty(t, res)
}

func TestGetProdukById(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "harga", "deskripsi", "kategori_id", "provider_id", "dibuat_pada", "diupdate_pada", "nominal", "dihapus"}).
			AddRow(produk.ID, produk.Nama, produk.Harga, produk.Deskripsi, produk.KategoriID, produk.ProviderID, produk.DibuatPada, produk.DiupdatePada, produk.Nominal, produk.Dihapus))
	res := repo.GetProdukById(1)
	assert.NotEmpty(t, res)
}

func TestGetProductByProviderId(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "harga", "deskripsi", "kategori_id", "provider_id", "dibuat_pada", "diupdate_pada", "nominal", "dihapus"}).
			AddRow(produk.ID, produk.Nama, produk.Harga, produk.Deskripsi, produk.KategoriID, produk.ProviderID, produk.DibuatPada, produk.DiupdatePada, produk.Nominal, produk.Dihapus))
	res := repo.GetProdukByKategori(1)
	assert.NotEmpty(t, res)
}

func TestGetProductByProviderKategoriId(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "harga", "deskripsi", "kategori_id", "provider_id", "dibuat_pada", "diupdate_pada", "nominal", "dihapus"}).
			AddRow(produk.ID, produk.Nama, produk.Harga, produk.Deskripsi, produk.KategoriID, produk.ProviderID, produk.DibuatPada, produk.DiupdatePada, produk.Nominal, produk.Dihapus))
	res := repo.GetProdukByKategoriProvider(1, 1)
	assert.NotEmpty(t, res)
}

func TestAddProduct(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewProductRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(produk.Nama, produk.Nominal, produk.Harga, produk.Deskripsi, produk.KategoriID, produk.ProviderID, produk.DiupdatePada, produk.DibuatPada, produk.Dihapus).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.AddProduct(produk)
	assert.NoError(t, err)
}

func TestAddProductError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewProductRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(produk.Nama, produk.Harga, produk.Deskripsi, produk.ProviderID, produk.DibuatPada, produk.DiupdatePada, produk.Nominal, produk.Dihapus).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.AddProduct(produk)
	assert.Error(t, err)
}

func TestGetSaldo(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"id", "saldo", "kategori_id", "dibuat_pada", "diupdate_pada"}).
			AddRow(saldo.ID, saldo.Saldo, saldo.KategoriID, saldo.DibuatPada, saldo.DiupdatePada))
	res := repo.GetSaldo()
	assert.NotEmpty(t, res)
}

func TestGetProviderByKategori(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "dibuat_pada", "diupdate_pada"}).
			AddRow(provider.ID, provider.Nama, provider.DibuatPada, provider.DiupdatePada))
	res := repo.GetProviderByKategori(1)
	assert.NotEmpty(t, res)
}

func TestPurchaseableProduct(t *testing.T) {
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
	var res = Result{
		ID:            1,
		Nama:          "Axis 10k",
		Nominal:       10000,
		Harga:         12000,
		Deskripsi:     "Pulsa Axis 10k",
		Nama_kategori: "Pulsa",
		Nama_provider: "Axis",
		Tersedia:      true,
	}
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewProductRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "nominal", "harga", "deskripsi", "nama_kategori", "nama_provider", "tersedia"}).
			AddRow(res.ID, res.Nama, res.Nominal, res.Harga, res.Deskripsi, res.Nama_kategori, res.Nama_provider, res.Tersedia))
	result := repo.GetPurchaseableProduct(1, 1)
	assert.NotEmpty(t, result)
}
