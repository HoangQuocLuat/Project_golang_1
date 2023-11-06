package models

type UserTbl struct {
	ID 			int64 	`db:"id"`
	Username 	string 	`db:"username"`
	Password 	string 	`db:"password"`
	Fullname 	string 	`db:"fullname"`
	Email 		string 	`db:"email"`
}