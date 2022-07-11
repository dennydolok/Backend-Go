package helper

import (
	"WallE/models"
	"fmt"
	"strings"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

func GenerateOrderId(id, kategori string) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	time := fmt.Sprintf("%02d%02d%02d", now.Day(), now.Month(), now.Year())
	orderid := "INV/" + time + "/" + kategori + "/" + id
	return orderid
}

func GetShortCategory(id uint) string {
	if id == 1 {
		return "PLS"
	} else if id == 2 {
		return "KTA"
	}
	return "VCG"
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

type RespBody struct {
	ID             uint          `json:"id"`
	Status         string        `json:"status_transaksi"`
	TotalHarga     string        `json:"total_harga"`
	OrderID        string        `json:"order_id"`
	TipePembayaran string        `json:"metode_pembayaran"`
	Bank           string        `json:"bank" form:"bank"`
	NomorHP        string        `json:"nomor_handphone" form:"nomor_handphone"`
	WaktuTransaksi string        `json:"waktu_transaksi"`
	TransaksiID    string        `json:"transaksi_id"`
	WaktuBayar     string        `json:"waktu_pembayaran"`
	UserID         uint          `json:"user_id" form:"user_id"`
	ProdukID       uint          `json:"produk_id" form:"produk_id"`
	Produk         models.Produk `json:"produk"`
	User           models.User   `json:"user"`
}

func FromMidBank(res coreapi.ChargeResponse, produk, NoHp, bank string, nominal, harga int64) RespBank {
	WaktuTanggal := strings.Split(res.TransactionTime, " ")
	var vanumber string

	fmt.Println(res)
	if bank == "permata" {
		vanumber = res.PermataVaNumber
	} else {
		vanumber = res.VaNumbers[0].VANumber
	}

	return RespBank{
		TransaksiId:      res.TransactionID,
		OrderId:          res.OrderID,
		TanggalTransaksi: WaktuTanggal[0],
		NamaProduk:       produk,
		Nominal:          nominal,
		Harga:            harga,
		NoHP:             NoHp,
		Jam:              WaktuTanggal[1],
		MetodePembayaran: "Bank Transfer",
		NomorVA:          vanumber,
	}
}

func FromMidEWallet(res coreapi.ChargeResponse, produk, NoHp string, nominal, harga int64) RespEWallet {
	WaktuTanggal := strings.Split(res.TransactionTime, " ")
	return RespEWallet{
		TransaksiId:       res.TransactionID,
		OrderId:           res.OrderID,
		TanggalTransaksi:  WaktuTanggal[0],
		NamaProduk:        produk,
		Nominal:           nominal,
		Harga:             harga,
		NoHP:              NoHp,
		Jam:               WaktuTanggal[1],
		MetodePembayaran:  "GoPay",
		QRCode:            res.Actions[0].URL,
		CancelPaymentLink: res.Actions[3].URL,
	}
}

func ToJsonBody(transaksi models.Transaksi) RespBody {
	return RespBody{
		ID:             transaksi.ID,
		Status:         transaksi.Status,
		TotalHarga:     transaksi.TotalHarga,
		OrderID:        transaksi.OrderID,
		TipePembayaran: transaksi.TipePembayaran,
		Bank:           transaksi.Bank,
		NomorHP:        transaksi.NomorHP,
		WaktuTransaksi: transaksi.WaktuTransaksi,
		TransaksiID:    transaksi.TransaksiID,
		WaktuBayar:     transaksi.WaktuBayar,
		UserID:         transaksi.UserID,
		ProdukID:       transaksi.ProdukID,
		Produk:         transaksi.Produk,
		User:           transaksi.User,
	}
}

func ToArrayJsonBody(transaksi []models.Transaksi) []RespBody{
	Arr := []RespBody{}
	for _ , j := range transaksi{
		trans :=  RespBody{
			ID: j.ID,
			Status: j.Status, 
			TotalHarga: j.TotalHarga,
			OrderID: j.OrderID,
			TipePembayaran: j.TipePembayaran,
			Bank: j.Bank,
			NomorHP: j.NomorHP,
			WaktuTransaksi: j.WaktuTransaksi,
			TransaksiID: j.TransaksiID,
			WaktuBayar: j.WaktuBayar,
			UserID: j.UserID,
			ProdukID: j.ProdukID,
			Produk: j.Produk,
			User: j.User,
		}
		Arr = append(Arr, trans)	
	}
	return Arr
}