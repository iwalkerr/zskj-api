package dao

import (
	"xframe/frontend/common/db"
)

const (
	insert           = ``
	deleteByKey      = ``
	updateByKey      = ``
	selectByKey      = ``
	selectAll        = ``
	getPwdByUsername = `SELECT user_id,password FROM api_users WHERE real_name=?`
)

func (p *Entity) GetPwdByUsername(realname string) (userId, pwd string) {
	_ = db.Conn(1).QueryRow(getPwdByUsername, realname).Scan(&userId, &pwd)
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
