package dao

const (
	insert      = ``
	deleteByKey = ``
	updateByKey = ``
	selectByKey = ``
	selectAll   = ``
)

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
