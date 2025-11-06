// 数据库层
package dao

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/miver02/learn-program/go/webook/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("邮箱或者手机号冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

// User 直接对应数据库表
type User struct {
	Id           int64 `gorm:"primaryKey,autoIncrement"`
	Name         string
	Email        sql.NullString `gorm:"unique"`
	Password     string
	Birthday     string
	Introduction string
	Phone        sql.NullString `gorm:"unique"`

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

func (dao *UserDAO) InsertUserInfo(ctx context.Context, domain_user domain.User) error {
	updates := map[string]interface{}{
		"name":         domain_user.Name,
		"birthday":     domain_user.Birthday,
		"introduction": domain_user.Introduction,
		"phone":        domain_user.Phone,
	}
	err := dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", domain_user.Id).
		Updates(updates).Error
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	return user, err
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
			return ErrUserDuplicate
		}
	}
	return err
}

func (dao *UserDAO) FindEmailById(ctx context.Context, id int64) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	return user, err
}

func (dao *UserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).Find(&user).Error
	return user, err
}
