package rest

import (
	"WallE/domains"
	"WallE/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transaksiController struct {
	services domains.TransaksiService
}

func (cont *transaksiController) NewTransaksiWallet(c echo.Context) error {
	transaksi := models.Transaksi{}
	c.Bind(&transaksi)
	err, res := cont.services.NewTransactionEWallet(transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusOK,
			"error": res,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode": http.StatusOK,
		"data": res,
	})
}

func (cont *transaksiController) NewTransactionBank(c echo.Context) error {
	transaksi := models.Transaksi{}
	c.Bind(&transaksi)
	err, res := cont.services.NewTransactionBank(transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusOK,
			"error": res,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode": http.StatusOK,
		"data": res,
	})
}

func (cont *transaksiController) UpdateTransaksi(c echo.Context) error {
	transaksi := models.Transaksi{}
	c.Bind(&transaksi)
	err := cont.services.UpdateTransaksi(transaksi.OrderID, transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusOK,
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode": http.StatusOK,
	})
}
