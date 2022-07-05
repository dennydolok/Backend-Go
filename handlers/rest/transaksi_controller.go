package rest

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transaksiController struct {
	services domains.TransaksiService
}

func (cont *transaksiController) NewTransaksiWallet(c echo.Context) error {
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkCustomer := helper.CheckCustomer(role)
	if checkCustomer != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": checkCustomer.Error(),
		})
	}
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
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkCustomer := helper.CheckCustomer(role)
	if checkCustomer != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": checkCustomer.Error(),
		})
	}
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

func (cont *transaksiController) GetUserTransactions(c echo.Context) error {
	filter := c.QueryParam("filter")
	userId := helper.GetUserId(c.Request().Header.Get("Authorization"))
	transactions := cont.services.GetUserTransactions(uint(userId), filter)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"transaksi": transactions,
	})
}

func (cont *transaksiController) GetAllTransaction(c echo.Context) error {
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": checkAdmin.Error(),
		})
	}
	transactions := cont.services.GetAllTransaction()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"transaksi": transactions,
	})
}
