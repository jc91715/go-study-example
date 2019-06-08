package model

// import (
// 	"fmt"
// )

type UserModel struct {
	Uid        int
	Username   string
	Department string
	Model
	// Attributes map[string]string
	// Table      string
}

func NewUserModel() *UserModel {
	m := UserModel{}
	m.table = "userinfo"
	m.Attributes = make(map[string]string)
	return &m
}

// func (m *UserModel) Find(id int) *UserModel {
// 	query := fmt.Sprintf("select * from %s where uid = '%d' limit 1", m.Table, id)
// 	fmt.Println(query)
// 	rows, err := Mgr.db.Query(query)
// 	CheckErr(err)

// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		CheckErr(err)
// 		fmt.Println(uid)
// 		fmt.Println(username)
// 		fmt.Println(department)
// 		fmt.Println(created)
// 		fmt.Println(m.Table)
// 		m.Attributes["username"] = username

// 	}

// 	// stmt, err := Mgr.db.Prepare("INSERT " + m.table + " SET username=?,department=?,created=?")
// 	// _, err = stmt.Exec("astaxie", "研发部门", "2012-12-09")
// 	// CheckErr(err)
// 	return m
// }
