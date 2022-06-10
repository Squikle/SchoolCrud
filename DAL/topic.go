package DAL

type Topic struct {
	Id          int     `db:"Id"`
	Name        *string `db:"Name"`
	Description *string `db:"Description"`
}

func (t *Topic) GetId() int {
	return t.Id
}
