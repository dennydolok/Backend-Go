package services

import (
	"WallE/domains"
	"WallE/helper"
	"WallE/models"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type serviceTransaksi struct {
	repo domains.TransaksiDomain
}

func (s *serviceTransaksi) NewTransactionBank(transaksi models.Transaksi) (error, interface{}) {
	midtrans.ServerKey = "SB-Mid-server-EU1Q1faz7h8T1eL51zvGViIC"
	midtrans.Environment = midtrans.Sandbox
	produk := s.repo.GetProdukById(transaksi.ProdukID)
	user := s.repo.GetUserById(transaksi.UserID)
	if len(transaksi.NomorHP) == 0 {
		transaksi.NomorHP = user.NomorHP
	}
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
			OrderID:  helper.GenerateOrderId(id, helper.GetShortCategory(produk.KategoriID)),
			GrossAmt: int64(produk.Harga),
		},
		Items: &[]midtrans.ItemDetails{item},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.Nama,
			Phone: user.NomorHP,
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
	data, errDB := s.repo.TransaksiBaru(transaksi)
	Result := helper.FromMidBank(*Response, data.ID, produk.Nama, transaksi.NomorHP, transaksi.Bank, int64(produk.Nominal), int64(produk.Harga))
	if errDB != nil {
		return errDB, ""
	}
	return nil, Result
}

func (s *serviceTransaksi) NewTransactionEWallet(transaksi models.Transaksi) (error, interface{}) {
	midtrans.ServerKey = "SB-Mid-server-EU1Q1faz7h8T1eL51zvGViIC"
	midtrans.Environment = midtrans.Sandbox
	produk := s.repo.GetProdukById(transaksi.ProdukID)
	user := s.repo.GetUserById(transaksi.UserID)
	if len(transaksi.NomorHP) == 0 {
		transaksi.NomorHP = user.NomorHP
	}
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
			OrderID:  helper.GenerateOrderId(id, helper.GetShortCategory(produk.KategoriID)),
			GrossAmt: int64(produk.Harga),
		},
		Items: &[]midtrans.ItemDetails{item},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: user.Nama,
			Phone: user.NomorHP,
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
	data, errDB := s.repo.TransaksiBaru(transaksi)
	res := helper.FromMidEWallet(*Response, data.ID, produk.Nama, transaksi.NomorHP, int64(produk.Nominal), int64(produk.Harga))
	if err != nil {
		return errDB, ""
	}
	errorBalance := s.repo.ReduceBalance(produk.KategoriID, produk.Nominal)
	if errorBalance != nil {
		return errorBalance, ""
	}
	return nil, res
}

func (s *serviceTransaksi) UpdateTransaksi(orderid string, transaksi models.Transaksi) error {
	err := s.repo.UpdateTransaksi(orderid, transaksi)
	if err != nil {
		return err
	}
	if transaksi.Status == "expire" || transaksi.Status == "cancel" {
		transaksi = s.repo.GetTransactionByOrderId(transaksi.OrderID)
		err = s.repo.RefundBalance(transaksi.Produk.KategoriID, transaksi.Produk.Nominal)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *serviceTransaksi) GetUserTransactions(id uint, filter string) []models.Transaksi {
	transaksi := s.repo.GetUserTransactions(id, filter)
	return transaksi
}

func (s *serviceTransaksi) GetListTransactionByUserId(userid uint) []models.Transaksi {
	return s.repo.GetListTransactionByUserId(userid)
}

func (s *serviceTransaksi) GetAllTransaction(filter string) []models.Transaksi {
	return s.repo.GetAllTransaction(filter)
}

func (s *serviceTransaksi) GetTransactionById(id uint) models.Transaksi {
	transaksi := s.repo.GetTransactionById(id)
	return transaksi
}

func (s *serviceTransaksi) GetTotalIncome() int {
	return s.repo.GetTotalIncome()
}

func NewTransaksiService(repo domains.TransaksiDomain) domains.TransaksiService {
	return &serviceTransaksi{
		repo: repo,
	}
}
