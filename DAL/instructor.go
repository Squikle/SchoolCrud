package DAL

type Instructor struct {
	Id       int `db:"Id"`
	PersonId int `db:"PersonId"`
	SchoolId int `db:"SchoolId"`
}

func (i *Instructor) GetId() int {
	return i.Id
}
