package rest

import (
	m "WallE/domains/mocks"
	"WallE/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transaksiService m.TransaksiService

var transaksi = models.Transaksi{
	ID:             1,
	Status:         "pending",
	TotalHarga:     "12000",
	OrderID:        "INV/02072022/VCG/2",
	TipePembayaran: "bank_transfer",
	Bank:           "bca",
	NomorHP:        "083811991642",
	WaktuTransaksi: time.Now().String(),
	TransaksiID:    "1661865e-7ce8-4e23-898b-9174b01c89f5",
	WaktuBayar:     time.Now().String(),
	UserID:         1,
	ProdukID:       1,
}

type RespBank struct {
	TransaksiId      string `json:"transaksi_id"`
	OrderId          string `json:"order_id"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	NamaProduk       string `json:"nama_produk"`
	Nominal          int64  `json:"nominal"`
	Harga            int64  `json:"harga"`
	NoHP             string `json:"no_handphone"`
	Jam              string `json:"jam"`
	MetodePembayaran string `json:"metode_pembayaran"`
	NomorVA          string `json:"nomor_va"`
}

var responseBank = RespBank{
	TransaksiId:      "1661865e-7ce8-4e23-898b-9174b01c89f5",
	OrderId:          "INV/02072022/VCG/2",
	TanggalTransaksi: time.Now().String(),
	NamaProduk:       "Pulsa Axis 10k",
	Nominal:          10000,
	Harga:            120000,
	NoHP:             "083811991642",
	Jam:              "16:16",
	MetodePembayaran: "bank_transfer",
	NomorVA:          "2131482132",
}

type RespEWallet struct {
	TransaksiId       string `json:"transaksi_id"`
	OrderId           string `json:"order_id"`
	TanggalTransaksi  string `json:"tanggal_transaksi"`
	NamaProduk        string `json:"nama_produk"`
	Nominal           int64  `json:"nominal"`
	Harga             int64  `json:"harga"`
	NoHP              string `json:"no_handphone"`
	Jam               string `json:"jam"`
	MetodePembayaran  string `json:"metode_pembayaran"`
	QRCode            string `json:"qr_kode_link"`
	CancelPaymentLink string `json:"batal_transaksi_link"`
}

var responseEWallet = RespEWallet{
	TransaksiId:       "1661865e-7ce8-4e23-898b-9174b01c89f5",
	OrderId:           "INV/02072022/VCG/2",
	TanggalTransaksi:  time.Now().String(),
	NamaProduk:        "Pulsa Axis 10k",
	Nominal:           10000,
	Harga:             120000,
	NoHP:              "083811991642",
	Jam:               "16:16",
	MetodePembayaran:  "gopay",
	QRCode:            "thisisqrcodelink.com",
	CancelPaymentLink: "thisisCancelcodelink.com",
}

func TestTransaksiBank(t *testing.T) {
	transaksiService.On("NewTransactionBank", mock.Anything).Return(nil, responseBank).Once()
	transaksiService.On("NewTransactionBank", mock.Anything).Return(errors.New("error bukan admin"), responseBank).Once()
	transaksiService.On("NewTransactionBank", mock.Anything).Return(errors.New("error service"), responseBank).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("transaksi bank", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("bank", "bca")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransactionBank(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("transaksi bank error admin", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("bank", "bca")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransactionBank(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("transaksi bank error service", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("bank", "bca")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransactionBank(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestTransaksiEWallet(t *testing.T) {
	transaksiService.On("NewTransactionEWallet", mock.Anything).Return(nil, responseEWallet).Once()
	transaksiService.On("NewTransactionEWallet", mock.Anything).Return(errors.New("error bukan admin"), responseEWallet).Once()
	transaksiService.On("NewTransactionEWallet", mock.Anything).Return(errors.New("error service"), responseEWallet).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("transaksi bank", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransaksiWallet(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("transaksi bank error admin", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + token
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransaksiWallet(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
	t.Run("transaksi bank error service", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "1")
		form.Add("produk_id", "1")
		form.Add("nomor_handphone", "083811991642")
		form.Encode()
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.NewTransaksiWallet(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestUpdateTransaksi(t *testing.T) {
	transaksiService.On("UpdateTransaksi", mock.Anything, mock.Anything).Return(nil).Once()
	transaksiService.On("UpdateTransaksi", mock.Anything, mock.Anything).Return(errors.New("error service")).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	postBody := map[string]interface{}{
		"transactions_id":    "1661865e-7ce8-4e23-898b-9174b01c89f5",
		"order_id":           "INV/02072022/VCG/2",
		"transaction_status": "settlement",
	}
	body, _ := json.Marshal(postBody)
	t.Run("update transaksi", func(t *testing.T) {
		r := httptest.NewRequest("PUT", "/", bytes.NewBuffer(body))
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.UpdateTransaksi(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("update transaksi error service", func(t *testing.T) {
		r := httptest.NewRequest("PUT", "/", bytes.NewBuffer(body))
		r.Header.Add("Accept", "application/json")
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Body.Close()
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.UpdateTransaksi(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetUserTransactions(t *testing.T) {
	var transactions []models.Transaksi
	transaksiService.On("GetUserTransactions", mock.AnythingOfType("uint"), mock.Anything).Return(transactions).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("get user data", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc2ODY5NzUsImlkIjo0LCJyb2xlIjoxfQ.Le1-pu7wUwRAHC6qhx4ScwSmzIMd9AqzbTEHY4S8WhA"
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.GetUserTransactions(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllTransactions(t *testing.T) {
	var transactions []models.Transaksi
	transaksiService.On("GetAllTransaction", mock.Anything).Return(transactions).Once()
	transaksiService.On("GetAllTransaction", mock.Anything).Return(transactions).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("get user data", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.GetAllTransaction(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("get user data error admin", func(t *testing.T) {
		bearer := "Bearer " + tokenCustomer
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.GetAllTransaction(eContext)
		assert.Equal(t, 401, w.Result().StatusCode)
	})
}

func TestGetTransactionById(t *testing.T) {
	transaksiService.On("GetTransactionById", mock.AnythingOfType("uint")).Return(transaksi).Once()
	transaksiService.On("GetTransactionById", mock.AnythingOfType("uint")).Return(transaksi).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("get transaksi by id", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		eContext.SetPath("/:id")
		eContext.SetParamNames("id")
		eContext.SetParamValues("1")
		controllerTransaksi.GetTransactionById(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
	t.Run("get transaksi by id error", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.GetTransactionById(eContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetTotalIncome(t *testing.T) {
	transaksiService.On("GetTotalIncome").Return(10000).Once()
	controllerTransaksi := transaksiController{
		services: &transaksiService,
	}
	e := echo.New()
	t.Run("get total income", func(t *testing.T) {
		bearer := "Bearer " + token
		r := httptest.NewRequest("GET", "/", bytes.NewBuffer(nil))
		r.Header.Set("Authorization", bearer)
		r.Header.Add("Accept", "application/json")
		w := httptest.NewRecorder()
		eContext := e.NewContext(r, w)
		controllerTransaksi.GetTotalIncome(eContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}
