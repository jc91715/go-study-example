package model

import (
	"fmt"
)

type Model struct {
	Attributes map[string]string
	table      string
}

func (m *Model) Find(id int) int {
	query := fmt.Sprintf("select * from %s where uid = '%d' limit 1", m.table, id)
	fmt.Println(query)
	rows, err := Mgr.db.Query(query)
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
