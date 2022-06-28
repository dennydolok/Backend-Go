package database

import (
	"WallE/config"
	"WallE/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf config.Config) *gorm.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		fmt.Println("Cannot connect to database : ", err)
	}
	DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.ResetPassword{},
		&models.Provider{},
		&models.Kategori{},
		&models.Saldo{},
		&models.Produk{},
		&models.Transaksi{},
	)
	return DB
}
