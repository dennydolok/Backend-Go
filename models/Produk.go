package models

import "time"

type Produk struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama         string    `json:"nama"`
	Harga        string    `json:"price"`
	KategoriID   uint      `json:"kategori_id"`
	ProviderID   uint      `json:"provider_id"`
	DibuatPada   time.Time `json:"dibuat_pada"`
	DiupdatePada time.Time `json:"diupdate_pada"`
	Kategory     Kategory  `json:"kategory" gorm:"foreignKey:KategoriID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
	Provider     Provider  `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE,OnUpdate:Cascade"`
}
