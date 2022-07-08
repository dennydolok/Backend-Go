package models

type Transaksi struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Status         string `json:"transaction_status" form:"transaction_status"`
	TotalHarga     string `json:"gross_amount"`
	OrderID        string `json:"order_id"`
	TipePembayaran string `json:"payment_type" form:"payment_type"`
	Bank           string `json:"bank" form:"bank"`
	NomorHP        string `json:"nomor_handphone" form:"nomor_handphone"`
	WaktuTransaksi string `json:"transaction_time"`
	TransaksiID    string `json:"transaction_id"`
	WaktuBayar     string `json:"settlement_time"`
	UserID         uint   `json:"user_id" form:"user_id"`
	ProdukID       uint   `json:"produk_id" form:"produk_id"`
	User           User   `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
	Produk         Produk `json:"produk" gorm:"foreignKey:ProdukID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
