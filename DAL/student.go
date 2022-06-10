package DAL

type Student struct {
	Id       int `db:"Id"`
	PersonId int `db:"PersonId"`
	SchoolId int `db:"SchoolId"`
}

func (s *Student) GetId() int {
	return s.Id
}
