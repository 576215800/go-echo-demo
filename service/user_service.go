package service

import (
	logger "echodemo/logs"
	"echodemo/model"
	"time"
)

/*
	插入用户数据
*/
func InsertUser(user model.User) (int64, error) {
	insertSql := "INSERT INTO User(username,account,password,phone_num,createTime,id) values(?,?,?,?,?,?)"
	stmt, err := db.Prepare(insertSql)
	if err != nil {
		logger.Error("用户插入sql语句错误....")
		return -1, err
	}
	user.Password = model.GetMD5Hash(user.Password)
	user.Id, err = model.GetNewID()
	//time包中的time对应于mysql中的datetime(精确到时分秒)和date(只精确到年月日)
	user.CreateTime = time.Now()
	if err != nil {
		logger.Error("用户ID生成出错....")
		return -1, err
	}
	_, err = stmt.Exec(user.Username, user.Account, user.Password, user.PhoneNum, user.CreateTime, user.Id)
	if err != nil {
		return -1, err
	}
	return user.Id, err
}

/*
	根据用户名	查询用户信息
*/
func GetUserByAccount(id string) (*model.User, error) {
	user := model.User{}
	selectsql := "SELECT username,Account,Password,PhoneNum,createTime,Id,FROM User where id=?"
	row := db.QueryRow(selectsql, id)
	err := row.Scan(&user.Username, &user.Account, &user.Password, &user.PhoneNum, &user.CreateTime, &user.Id)
	if err != nil {
		logger.Error("用户信息查询失败....")
	}
	return &user, err
}

/*
	根据用户id 删除用户信息
*/
func DeleteUser(id int64) error {
	deletesql := "delete from User where Id=?"
	stmt, err := db.Prepare(deletesql)
	if err != nil {
		logger.Error("用户删除sql语句出错.....")
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Error("删除用户失败.....")
		return err
	}
	return nil
}
func UpdateUser(id, newPassword string) error {
	updateSql := "update from User set Password=? where id=?"
	stmt, err := db.Prepare(updateSql)
	if err != nil {
		logger.Error("用户密码更新sql语句出错.....")
		return err
	}
	_, err = stmt.Exec(id, newPassword)
	if err!=nil{
		logger.Error("用户更新密码失败....")
		return err
	}
	return nil
}
