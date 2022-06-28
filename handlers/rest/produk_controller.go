package rest

import (
	"WallE/domains"
	"WallE/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productController struct {
	services domains.ProductService
}

func (cont *productController) GetKategori(c echo.Context) error {
	kategori := cont.services.GetKategori()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":     http.StatusOK,
		"kategori": kategori,
	})
}

func (cont *productController) AddSaldo(c echo.Context) error {
	saldobaru, _ := strconv.Atoi(c.FormValue("saldo"))
	kategoriid, _ := strconv.Atoi(c.FormValue("kategori_id"))
	fmt.Println(saldobaru, kategoriid)
	err := cont.services.AddSaldo(saldobaru, uint(kategoriid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":    http.StatusOK,
		"message": "sukses",
	})
}
func (cont *productController) GetProdukByKategori(c echo.Context) error {
	// kategoriid, _ := strconv.Atoi(c.FormValue("kategori_id"))
	kategoriid, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(kategoriid)
	produk := cont.services.GetProdukByKategori(uint(kategoriid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) GetProdukById(c echo.Context) error {
	produkid, _ := strconv.Atoi(c.Param("id"))
	produk := cont.services.GetProdukById(uint(produkid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) GetProdukByKategoriProvider(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.QueryParam("kategori"))
	providerid, _ := strconv.Atoi(c.QueryParam("provider"))
	produk := cont.services.GetProdukByKategoriProvider(uint(kategoriid), uint(providerid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) AddProduct(c echo.Context) error {
	produk := models.Produk{}
	c.Bind(&produk)
	err := cont.services.AddProduct(produk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":    http.StatusOK,
		"message": "sukses",
	})
}

func (cont *productController) GetSaldo(c echo.Context) error {
	saldo := cont.services.GetSaldo()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"saldo": saldo,
	})
}

func (cont *productController) GetProviderByKategori(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.Param("kategori_id"))
	provider := cont.services.GetProviderByKategori(uint(kategoriid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":     http.StatusOK,
		"provider": provider,
	})
}

func (cont *productController) GetPurchaseableProduct(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.QueryParam("kategori"))
	providerid, _ := strconv.Atoi(c.QueryParam("provider"))
	produk := cont.services.GetPurchaseableProduct(uint(kategoriid), uint(providerid))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":   http.StatusOK,
		"produk": produk,
	})
}

func (cont *productController) UpdateProductById(c echo.Context) error {
	produk := models.Produk{}
	produkid, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&produk)
	err := cont.services.UpdateProductById(uint(produkid), produk)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "sukses",
	})
}
