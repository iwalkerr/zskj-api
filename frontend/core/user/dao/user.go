package dao

import (
	"fmt"
	"xframe/frontend/common/db"
)

const (
	insert           = ``
	deleteByKey      = ``
	updateByKey      = ``
	selectByKey      = ``
	selectAll        = ``
	getPwdByUsername = `SELECT name,head_picture,real_name,signature,sex,birthday,password,user_id,phone FROM api_users WHERE real_name=? or phone=?`
)

func (p *Entity) GetPwdByUsername(realname string) (e LoginResp) {
	err := db.Conn(1).QueryRowx(getPwdByUsername, realname, realname).StructScan(&e)
	fmt.Println(err)
	return
}

func (p *Entity) Insert() (int, error) {

	return 0, nil
}

func (p *Entity) DeleteByKey(id int) error {

	return nil
}

func (p *Entity) UpdateByKey(id int) error {

	return nil
}

func (p *Entity) SelectByKey(id int) *Entity {

	return nil
}

func (p *Entity) SelectAll() []*Entity {

	return nil
}
