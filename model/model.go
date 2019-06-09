package model

import (
	// "encoding/json"
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

func init() {

	orm.RegisterModel(new(Userinfo), new(RainlabBlogPosts))
}

type Model struct {
	Attributes map[string]string
	table      string
}

func (m *Model) Find(id int) *Model {
	query := fmt.Sprintf("select * from %s where uid = '%d' limit 1", m.table, id)
	fmt.Println(query)
	rows, err := Mgr.db.Query(query)
	CheckErr(err)
	fmt.Println(reflect.ValueOf(m))
	for rows.Next() {
		var uid string
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &created, &department)
		CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		m.Attributes["uid"] = uid
		m.Attributes["username"] = username
		m.Attributes["department"] = department
		m.Attributes["created"] = created
		// s := department
		// in := []byte(s)
		// var raw map[string]interface{}
		// json.Unmarshal(in, &raw)
		// raw["count"] = 1
		// out, _ := json.Marshal(raw)
		// println(string(out))

	}
	// fmt.Println(res)
	// stmt, err := Mgr.db.Prepare("INSERT " + m.table + " SET username=?,department=?,created=?")
	// _, err = stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// CheckErr(err)
	return m
}
