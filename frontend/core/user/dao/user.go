package dao

import (
	"fmt"
	"xframe/frontend/common/db"
)

const (
	deleteByKey      = ``
	updateByKey      = ``
	selectByKey      = ``
	selectAll        = ``
	getPwdByUsername = `SELECT name,head_picture,real_name,signature,sex,birthday,password,user_id,phone FROM api_users WHERE real_name=? or phone=?`
	existPhone       = `SELECT COUNT(user_id) count FROM api_users WHERE phone=?`
	insert           = `insert into api_users(name,head_picture,real_name,password,phone,created_time,updated_time) values(?,?,?,?,?,?,?)`
)

func (p *Entity) ExistPhone(phone string) bool {
	var count int
	_ = db.Conn(1).QueryRow(existPhone, phone).Scan(&count)
	return count != 0
}

func (p *Entity) GetPwdByUsername(realname string) (e LoginResp) {
	err := db.Conn(1).QueryRowx(getPwdByUsername, realname, realname).StructScan(&e)
	fmt.Println(err)
	return
}

func (p *Entity) Insert() (int, error) {
	res, err := db.Conn().Exec(insert, &p.Name, &p.HeadPicture, &p.Username, &p.Password, &p.Phone, &p.CreateTime, &p.UpdateTime)
	if err != nil {
		return 0, err
	}
	if id, err := res.LastInsertId(); err != nil {
		return 0, err
	} else {
		return int(id), nil
	}
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
