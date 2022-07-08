package repositories

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetByEmail(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
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
	_, res := repo.GetByEmail("Client02@gmail.com")
	assert.NoError(t, res)
}

func TestGetByEmailError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("Error"))
	_, res := repo.GetByEmail("err")
	assert.Error(t, res)
}

func TestGetById(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
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
	_, res := repo.GetUserDataById(1)
	assert.NoError(t, res)
}

func TestGetByIdError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("Error"))
	_, res := repo.GetUserDataById(0)
	assert.Error(t, res)
}

func TestGetResetPassword(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"1", "012345", "test@gmail.com", "4", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000"}).
			AddRow("1", "012345", "test@gmail.com", "4", "1", "2022-07-01 20:06:34.000", "2022-07-01 20:06:34.000"))
	_, res := repo.GetResetPassword("test@gmail.com")
	assert.NoError(t, res)
}

func TestGetResetPasswordError(t *testing.T) {
	var dbmock, mock, _ = sqlmock.New()
	var db, _ = gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbmock,
		SkipInitializeWithVersion: true,
	},
	})
	var repo = NewUserRepository(db)
	defer dbmock.Close()
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
