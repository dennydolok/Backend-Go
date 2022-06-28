package services

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"fmt"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type serviceTransaksi struct {
	repo domains.TransaksiDomain
}

func (s *serviceTransaksi) NewTransactionBank(transaksi models.Transaksi) (error, interface{}) {
	midtrans.ServerKey = "SB-Mid-server-oiWZtAAp4u5TaWKDm5AiRr1R"
	midtrans.Environment = midtrans.Sandbox
	produk := s.repo.GetProdukById(transaksi.ProdukID)
	user := s.repo.GetUserById(transaksi.UserID)
	item := midtrans.ItemDetails{
		ID:       strconv.FormatUint(uint64(produk.ID), 10),
		Name:     produk.Nama,
		Price:    int64(produk.Harga),
		Qty:      1,
		Brand:    produk.Provider.Nama,
		Category: produk.Kategory.Nama,
	}
	id := strconv.FormatUint(uint64(s.repo.GetLastId()), 10)
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(transaksi.Bank),
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  helper.GenerateOrderId(id, produk.Kategory.Nama),
			GrossAmt: int64(produk.Harga),
		},
		Items: &[]midtrans.ItemDetails{item},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.Name,
			Phone: user.PhoneNumber,
			Email: user.Email,
		},
	}
	Response, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		return err, ""
	}
	transaksi.TipePembayaran = Response.PaymentType
	transaksi.OrderID = Response.OrderID
	transaksi.TotalHarga = Response.GrossAmount
	transaksi.WaktuTransaksi = Response.TransactionTime
	transaksi.TransaksiID = Response.TransactionID
	transaksi.Status = Response.TransactionStatus
	errDB := s.repo.TransaksiBaru(transaksi)
	if err != nil {
		return errDB, ""
	}
	return nil, Response
}

func (s *serviceTransaksi) NewTransactionEWallet(transaksi models.Transaksi) (error, interface{}) {
	midtrans.ServerKey = "SB-Mid-server-oiWZtAAp4u5TaWKDm5AiRr1R"
	midtrans.Environment = midtrans.Sandbox
	produk := s.repo.GetProdukById(transaksi.ProdukID)
	user := s.repo.GetUserById(transaksi.UserID)
	item := midtrans.ItemDetails{
		ID:       strconv.FormatUint(uint64(produk.ID), 10),
		Name:     produk.Nama,
		Price:    int64(produk.Harga),
		Qty:      1,
		Brand:    produk.Provider.Nama,
		Category: produk.Kategory.Nama,
	}
	id := strconv.FormatUint(uint64(s.repo.GetLastId()), 10)
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  helper.GenerateOrderId(id, produk.Kategory.Nama),
			GrossAmt: int64(produk.Harga),
		},
		Items: &[]midtrans.ItemDetails{item},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.Name,
			Phone: user.PhoneNumber,
			Email: user.Email,
		},
	}
	Response, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		return err, ""
	}
	transaksi.TipePembayaran = Response.PaymentType
	transaksi.OrderID = Response.OrderID
	transaksi.TotalHarga = Response.GrossAmount
	transaksi.WaktuTransaksi = Response.TransactionTime
	transaksi.TransaksiID = Response.TransactionID
	transaksi.Status = Response.TransactionStatus
	errDB := s.repo.TransaksiBaru(transaksi)
	if err != nil {
		return errDB, ""
	}
	return nil, Response
}

func (s *serviceTransaksi) UpdateTransaksi(orderid string, transkasi models.Transaksi) error {
	fmt.Println(orderid, transkasi)
	
	return s.repo.UpdateTransaksi(orderid, transkasi)
}

func (s *serviceTransaksi) GetListTransactionByUserId(userid uint) []models.Transaksi {
	transaksi := []models.Transaksi{}
	return transaksi
}

func NewTransaksiService(repo domains.TransaksiDomain) domains.TransaksiService {
	return &serviceTransaksi{
		repo: repo,
	}
}
