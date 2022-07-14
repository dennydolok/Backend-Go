package rest

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"net/http"
	"strconv"

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
			"kode":  http.StatusUnauthorized,
			"pesan": checkCustomer.Error(),
		})
	}
	transaksi := models.Transaksi{}
	c.Bind(&transaksi)
	err, res := cont.services.NewTransactionEWallet(transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
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
			"kode":  http.StatusInternalServerError,
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
	data := helper.ToArrayJsonBody(transactions)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"transaksi": data,
	})
}

func (cont *transaksiController) GetAllTransaction(c echo.Context) error {
	role := helper.GetClaim(c.Request().Header.Get("Authorization"))
	checkAdmin := helper.CheckAdmin(role)
	filter := c.QueryParam("filter")
	if checkAdmin != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"kode":  http.StatusUnauthorized,
			"pesan": checkAdmin.Error(),
		})
	}
	transactions := cont.services.GetAllTransaction(filter)
	data := helper.ToArrayJsonBody(transactions)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"transaksi": data,
	})
}

func (cont *transaksiController) GetTransactionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "error",
		})
	}
	transactions := cont.services.GetTransactionById(uint(id))
	data := helper.ToJsonBody(transactions)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"transaksi": data,
	})
}

func (cont *transaksiController) GetTotalIncome(c echo.Context) error {
	income := cont.services.GetTotalIncome()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":      http.StatusOK,
		"pemasukan": income,
	})
}
