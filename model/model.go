package model

import (
	"fmt"
)

type Model struct {
	Attributes map[string]string
	table      string
}

func (m *Model) Find(id int) int {
	fmt.Println(id)

	rows, err := Mgr.db.Query("select * from userinfo where uid = ? limit 1", id)
	CheckErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	// fmt.Println(res)
	// stmt, err := Mgr.db.Prepare("INSERT " + m.table + " SET username=?,department=?,created=?")
	// _, err = stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// CheckErr(err)
	return 1
}
