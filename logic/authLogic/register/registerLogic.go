package logic

import (
	"Project/config/db"
	customRes "Project/types/Res_types/CustomRes"
	"Project/types/req_types/authReq_types"
	"log"
)


func RegisterLogic(req authReq_types.UserRegister) (*customRes.UserInfor, error) {
	var count int 
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users_tbl WHERE username = ?", req.Username).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		
	}
}
