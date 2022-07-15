package rest

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"errors"
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var kategori = models.Kategori{
	ID:   1,
	Nama: "Pulsa",
}

var produk = models.Produk{
	ID:           1,
	Nama:         "Pulsa Axis",
	Nominal:      10000,
	Harga:        12000,
	Deskripsi:    "Pulsa Axis 10000",
	KategoriID:   1,
	ProviderID:   1,
	DibuatPada:   time.Now(),
	DiupdatePada: time.Now(),
	Dihapus:      gorm.DeletedAt{},
}

var produkService m.ProductService

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgyOTU5ODEsImlkIjo0LCJyb2xlIjoxfQ.yxhWlBYvH5J_DJu5Yks1byOuCJUuVXsTvxYiyRBsplo"
var tokenCustomer = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgyOTQ1MjksImlkIjo2LCJyb2xlIjoyfQ.eTBg5QPyxVC53Z9RhiPdr14bp7f-CZoo_12hEz2GY1c"

func TestGetKategori(t *testing.T) {
	kategoris := []models.Kategori{}
	produkService.On("GetKategori").Return(kategoris).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get kategori", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetKategori(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestAddSaldo(t *testing.T) {
	produkService.On("AddSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("uint")).Return(nil).Once()
	produkService.On("AddSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("uint")).Return(errors.New("saldo must be above 1")).Once()
	produkService.On("AddSaldo", mock.AnythingOfType("int"), 0).Return(errors.New("error kategori id")).Once()
	produkService.On("AddSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("uint")).Return(errors.New("error not an admin")).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("add saldo", func(t *testing.T) {
		form := url.Values{}
		form.Add("saldo", "1000")
		form.Add("kategori_id", "1")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddSaldo(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("add saldo error amount", func(t *testing.T) {
		form := url.Values{}
		form.Add("saldo", "0")
		form.Add("kategori_id", "1")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddSaldo(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
	t.Run("add saldo error kategori id", func(t *testing.T) {
		form := url.Values{}
		form.Add("saldo", "1000")
		form.Add("kategori_id", "0")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddSaldo(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
	t.Run("add saldo not an admin", func(t *testing.T) {
		form := url.Values{}
		form.Add("saldo", "1000")
		form.Add("kategori_id", "1")
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddSaldo(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestGetProductByKategori(t *testing.T) {
	produkService.On("GetProdukByKategori", mock.AnythingOfType("uint")).Return(nil).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get product by kategori", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		query := r.URL.Query()
		query.Add("id", "1")
		r.URL.RawQuery = query.Encode()
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetProdukByKategori(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetProductById(t *testing.T) {
	produkService.On("GetProdukById", mock.AnythingOfType("uint")).Return(produk).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get product by kategori", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		query := r.URL.Query()
		query.Add("id", "1")
		r.URL.RawQuery = query.Encode()
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetProdukById(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetProductByKategoriProvider(t *testing.T) {
	produks := []models.Produk{}
	produkService.On("GetProdukByKategoriProvider", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(produks).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get product by kategori", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		query := r.URL.Query()
		query.Add("kategori", "1")
		query.Add("provider", "1")
		r.URL.RawQuery = query.Encode()
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetProdukByKategoriProvider(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestAddProduct(t *testing.T) {
	produkService.On("AddProduct", mock.Anything).Return(nil).Once()
	produkService.On("AddProduct", mock.Anything).Return(errors.New("empty product name")).Once()
	produkService.On("AddProduct", mock.Anything).Return(errors.New("error")).Once()
	produkService.On("AddProduct", mock.Anything).Return(errors.New("error not an admin")).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("add product", func(t *testing.T) {
		form := url.Values{}
		form.Add("nama", "axis")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddProduct(eContext)
		fmt.Println(w.Body.String())
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("add product error nama", func(t *testing.T) {
		form := url.Values{}
		form.Add("nama", "")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddProduct(eContext)
		fmt.Println(w.Body.String())
		assert.Equal(t, 500, w.Result().StatusCode)
	})
	t.Run("add product error insert", func(t *testing.T) {
		form := url.Values{}
		form.Add("nama", "axis")
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddProduct(eContext)
		fmt.Println(w.Body.String())
		assert.Equal(t, 500, w.Result().StatusCode)
	})
	t.Run("add product not an admin", func(t *testing.T) {
		form := url.Values{}
		form.Add("nama", "axis")
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.AddProduct(eContext)
		fmt.Println(w.Body.String())
		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestGetSaldo(t *testing.T) {
	saldo := []models.Saldo{}
	produkService.On("GetSaldo").Return(saldo).Once()
	produkService.On("GetSaldo").Return(errors.New("error")).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get saldo", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetSaldo(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("Get saldo not admin", func(t *testing.T) {
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetSaldo(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestGetProviderByKategori(t *testing.T) {
	type Result struct {
		ID   uint   `json:"id"`
		Nama string `json:"nama"`
	}
	var res []Result
	produkService.On("GetProviderByKategori", mock.AnythingOfType("uint")).Return(res).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get provider by kategori", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		query := r.URL.Query()
		query.Add("id", "1")
		r.URL.RawQuery = query.Encode()
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetProviderByKategori(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetPurchaseableProduct(t *testing.T) {
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
	var result []Result
	produkService.On("GetPurchaseableProduct", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(result).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("get purchaseable product", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", nil)
		query := r.URL.Query()
		query.Add("kategori", "1")
		query.Add("provider", "1")
		r.URL.RawQuery = query.Encode()
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.GetPurchaseableProduct(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateProduct(t *testing.T) {
	produk = models.Produk{
		Nama:       "Pulsa Axis",
		KategoriID: 1,
		ProviderID: 1,
	}
	produkService.On("UpdateProductById", uint(1), produk).Return(nil).Once()
	produkService.On("UpdateProductById", uint(1), produk).Return(errors.New("param empty")).Once()
	produkService.On("UpdateProductById", uint(1), produk).Return(errors.New("error not admin")).Once()
	produkService.On("UpdateProductById", uint(1), produk).Return(errors.New("error service")).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("update product", func(t *testing.T) {
		bearer := "Bearer " + token
		form := url.Values{}
		form.Add("nama", "Pulsa Axis")
		form.Add("kategori_id", "1")
		form.Add("provider_id", "1")
		form.Encode()
		r := httptest.NewRequest("PUT", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerProduct.UpdateProductById(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("update product error param empty", func(t *testing.T) {
		bearer := "Bearer " + token
		form := url.Values{}
		form.Add("nama", "Pulsa Axis")
		form.Add("kategori_id", "1")
		form.Add("provider_id", "1")
		form.Encode()
		r := httptest.NewRequest("PUT", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.UpdateProductById(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
	t.Run("update product error not admin", func(t *testing.T) {
		bearer := "Bearer " + tokenCustomer
		form := url.Values{}
		form.Add("nama", "Pulsa Axis")
		form.Add("kategori_id", "1")
		form.Add("provider_id", "1")
		form.Encode()
		r := httptest.NewRequest("PUT", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerProduct.UpdateProductById(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("update product error", func(t *testing.T) {
		bearer := "Bearer " + token
		form := url.Values{}
		form.Add("nama", "Pulsa Axis")
		form.Add("kategori_id", "1")
		form.Add("provider_id", "1")
		form.Encode()
		r := httptest.NewRequest("PUT", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerProduct.UpdateProductById(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestDeleteProduct(t *testing.T) {
	produkService.On("DeleteProdukById", mock.AnythingOfType("uint")).Return(nil).Once()
	produkService.On("DeleteProdukById", mock.AnythingOfType("uint")).Return(errors.New("error not admin")).Once()
	produkService.On("DeleteProdukById", mock.AnythingOfType("uint")).Return(errors.New("error service")).Once()
	controllerProduct := productController{
		services: &produkService,
	}
	e := echo.New()
	t.Run("delete product", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("PUT", "/", nil)
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerProduct.DeleteProdukById(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("error delete product not admin", func(t *testing.T) {
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("PUT", "/", nil)
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerProduct.DeleteProdukById(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("error delete product not admin", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("PUT", "/", nil)
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerProduct.DeleteProdukById(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
