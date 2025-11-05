// 数据库层
package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/miver02/learn-program/go/webook/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

// User 直接对应数据库表
type User struct {
	Id           int64 `gorm:"primaryKey,autoIncrement"`
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Birthday     string
	Introduction string

	// 创建时间
	Ctime int64
	// 更新时间
	Utime int64
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) InsertUserInfo(ctx context.Context, new_ud domain.User) error {
	updates := map[string]interface{}{
		"name":         new_ud.Name,
		"birthday":     new_ud.Birthday,
		"introduction": new_ud.Introduction,
	}
	err := dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", new_ud.Id).
		Updates(updates).Error
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var ud User
	err := dao.db.WithContext(ctx).Where("email = ?", email).Find(&ud).Error
	return ud, err
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var ud User
	err := dao.db.WithContext(ctx).Where("id = ?", id).Find(&ud).Error
	return ud, err
}
