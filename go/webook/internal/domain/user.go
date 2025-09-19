package domain

import "time"

type User struct {
	Id 				int64 	
	Name			string
	Email 			string	
	Password 		string
	Birthday  		string
	Introduction 	string

	// 创建时间
	Ctime 		time.Time
	// 更新时间
	Utime 		time.Time
}