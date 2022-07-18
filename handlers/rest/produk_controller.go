package rest

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"net/http"
	"strconv"
	"strings"

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
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	type data struct {
		Saldo      int `json:"saldo" form:"saldo"`
		KategoriId int `json:"kategori_id" form:"kategori_id"`
	}
	dataSaldo := data{}
	c.Bind(&dataSaldo)
	if dataSaldo.Saldo < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "Saldo harus diisi",
		})
	}
	err := cont.services.AddSaldo(dataSaldo.Saldo, uint(dataSaldo.KategoriId))
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
func (cont *productController) GetProdukByKategori(c echo.Context) error {
	kategoriid, _ := strconv.Atoi(c.Param("id"))
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
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	produk := models.Produk{}
	c.Bind(&produk)
	nama := strings.TrimSpace(produk.Nama)
	if len(nama) == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "nama produk harus diisi",
		})
	}
	err := cont.services.AddProduct(produk)
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

func (cont *productController) GetSaldo(c echo.Context) error {
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	saldo := cont.services.GetSaldo()
	count := c.QueryParam("hitung")
	if count == "total" {
		total := 0
		for _, v := range saldo {
			total += v.Saldo
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"kode":  http.StatusOK,
			"saldo": total,
		})
	}
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
	produkid, _ := strconv.Atoi(c.Param("id"))
	if produkid == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "error",
		})
	}
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	produk := models.Produk{}
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

func (cont *productController) DeleteProdukById(c echo.Context) error {
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	produkid, _ := strconv.Atoi(c.Param("id"))
	err := cont.services.DeleteProdukById(uint(produkid))
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
