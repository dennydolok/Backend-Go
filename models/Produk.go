package models

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama         string         `json:"nama" form:"nama"`
	Nominal      int            `json:"nominal" form:"nominal"`
	Harga        int            `json:"harga" form:"harga"`
	Deskripsi    string         `json:"deskripsi" form:"deskripsi"`
	KategoriID   uint           `json:"kategori_id"`
	ProviderID   uint           `json:"provider_id"  form:"provider_id"`
	DibuatPada   time.Time      `json:"dibuat_pada"`
	DiupdatePada time.Time      `json:"diupdate_pada"`
	Dihapus      gorm.DeletedAt `json:",omitempty"`
	Kategory     Kategori       `json:"kategori" gorm:"foreignKey:KategoriID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
	Provider     Provider       `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
