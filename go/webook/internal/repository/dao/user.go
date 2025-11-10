// 数据库层
package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/miver02/learn-program/go/webook/internal/consts"
	"gorm.io/gorm"
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

func (dao *UserDAO) InsertUserInfo(ctx context.Context, user User) error {
	updates := make(map[string]interface{})

	if user.Password != "" {
		updates["password"] = user.Password
	}
	if user.Phone.Valid && user.Phone.String != "" {
		updates["phone"] = user.Phone
	}
	if user.Name != "" {
		updates["name"] = user.Name
	}
	if user.Email.Valid && user.Email.String != "" {
		updates["email"] = user.Email
	}
	if user.Birthday != "" {
		updates["birthday"] = user.Birthday
	}
	if user.Introduction != "" {
		updates["introduction"] = user.Introduction
	}
	// fmt.Printf("%v", updates)
	err := dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", user.Id).
		Updates(updates).Error
	// 处理唯一性冲突错误
	err = dao.uniqueCheck(err)
	return err
}

func (dao *UserDAO) uniqueCheck(err error) error {
	// 处理唯一性冲突错误
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突或者手机好冲突
			return consts.ErrUserDuplicate
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	return user, err
}

func (dao *UserDAO) Insert(ctx context.Context, user User) (User, error) {
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Ctime = now
	err := dao.db.WithContext(ctx).Create(&user).Error
	// 处理唯一性冲突错误
	err = dao.uniqueCheck(err)
	return user, err
}

func (dao *UserDAO) FindEmailById(ctx context.Context, id int64) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	return user, err
}

func (dao *UserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return user, err
}
