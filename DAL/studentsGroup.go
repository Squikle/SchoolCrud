package DAL

type StudentsGroup struct {
	Id        int `db:"Id"`
	StudentId int `db:"StudentId"`
	GroupId   int `db:"GroupId"`
}

func (sg *StudentsGroup) GetId() int {
	return sg.Id
}
