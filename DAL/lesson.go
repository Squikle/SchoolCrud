package DAL

type Lessons struct {
	Id           int `db:"Id"`
	InstructorId int `db:"InstructorId"`
	TopicId      int `db:"TopicId"`
}

func (l *Lessons) GetId() int {
	return l.Id
}
