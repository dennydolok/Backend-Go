package repositories

import (
	"WallE/config"
	"WallE/models"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
	config.InitConfig().DB_USERNAME,
	config.InitConfig().DB_PASSWORD,
	config.InitConfig().DB_HOST,
	config.InitConfig().DB_PORT,
	"db_walle_test",
)

func TestGetByEmail(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"}).
			AddRow("1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"))
	_, res := repo.GetByEmail("Client02@gmail.com")
	assert.NoError(t, res)
}

func TestGetByEmailError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("Error"))
	_, res := repo.GetByEmail("err")
	assert.Error(t, res)
}

func TestGetById(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"}).
			AddRow("1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"))
	_, res := repo.GetUserDataById(1)
	assert.NoError(t, res)
}

func TestGetByIdError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("Error"))
	_, res := repo.GetUserDataById(0)
	assert.Error(t, res)
}

func TestGetResetPassword(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", "012345", "test@gmail.com", "4", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000"}).
			AddRow("1", "012345", "test@gmail.com", "4", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000"))
	_, res := repo.GetResetPassword("test@gmail.com")
	assert.NoError(t, res)
}

func TestGetResetPasswordError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("Error"))
	_, res := repo.GetResetPassword("")
	assert.Error(t, res)
}

// func TestUpdatePassword(t *testing.T) {
// 	var dbmock, mock, _ = sqlmock.New()
// 	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
// 		Conn:                      dbmock,
// 		SkipInitializeWithVersion: true,
// 	},
// 	})
// 	var repo = NewUserRepository(db)
// 	defer dbmock.Close()
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()
// 	err := repo.UpdatePassword("test@gmail.com", "MTIz")
// 	assert.NoError(t, err)
// }

func TestVerifikasi(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.Verifikasi(1)
	assert.NoError(t, err)
}

func TestVerifikasiError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs().
		WillReturnError(fmt.Errorf("error"))
	mock.ExpectCommit()
	err := repo.Verifikasi(0)
	assert.Error(t, err)
}

func TestGetUserByEmailError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"}).
			AddRow("1", "test", "test@gmail.com", "MTIz", "083811223312", "518929", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000", "1"))
	_, res := repo.GetUserByEmail("dennydolok12@gmail.com")
	assert.Error(t, res)
}

func TestRegister(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	user := models.User{
		Nama:         "test",
		Email:        "test@gmail.com",
		Password:     "1234",
		NomorHP:      "01231231",
		Kode:         "123132",
		Verifikasi:   false,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
		RoleID:       1,
	}
	repo := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(user.Nama, user.Email, user.Password, user.NomorHP, user.Kode, user.Verifikasi, user.DiBuatPada, user.DiUpdatePada, user.RoleID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.Register(user)
	assert.NoError(t, err)
}

func TestRegisterError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	user := models.User{
		Nama:         "test",
		Email:        "test@gmail.com",
		Password:     "1234",
		NomorHP:      "01231231",
		Kode:         "123132",
		Verifikasi:   false,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
		RoleID:       1,
	}
	repo := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(user.Nama, user.Password, user.NomorHP, user.Kode, user.Verifikasi, user.DiBuatPada, user.DiUpdatePada, user.RoleID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.Register(user)
	assert.Error(t, err)
}

func TestCreateResetPassword(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	reset := models.ResetPassword{
		Kode:         "123456",
		Email:        "testing@gmail.com",
		UserID:       1,
		Selesai:      false,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
	}
	repo := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(reset.Kode, reset.Email, reset.UserID, reset.Selesai, reset.DiBuatPada, reset.DiUpdatePada).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.CreateResetPassword(reset)
	assert.NoError(t, err)
}

func TestCreateResetPasswod(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()

	dbmock.Begin()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	reset := models.ResetPassword{
		Kode:         "123456",
		Email:        "testing@gmail.com",
		UserID:       1,
		Selesai:      false,
		DiBuatPada:   time.Now(),
		DiUpdatePada: time.Now(),
	}
	repo := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(reset.Kode, reset.Email, reset.Selesai, reset.DiBuatPada, reset.DiUpdatePada).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.CreateResetPassword(reset)
	assert.Error(t, err)
}

// func TestUpdatePassword(t *testing.T) {
// 	var dbmock, mock, _ = sqlmock.New()
// 	dbmock.Begin()
// 	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
// 		Conn:                      dbmock,
// 		SkipInitializeWithVersion: true,
// 		DSN:                       Connection,
// 	},
// 	})
// 	user := models.User{
// 		Email:    "email@gmail.com",
// 		Password: "MTIz",
// 	}
// 	var repo = NewUserRepository(db)
// 	defer dbmock.Close()
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(user.Password, user.Email).
// 		WillReturnResult(sqlmock.NewResult(0, 0))
// 	mock.ExpectCommit()
// 	err := repo.UpdatePassword(user.Email, user.Password)
// 	assert.NoError(t, err)
// }
