package DAL

type Group struct {
	Id           int `db:"Id"`
	InstructorId int `db:"InstructorId"`
	SchoolId     int `db:"SchoolId"`
}

func (g *Group) GetId() int {
	return g.Id
}
