package methods

type Student struct {
	Name   string
	Grades []int
}

func (stu *Student) Average() float64 {
	var sum float64
	for _, grade := range stu.Grades {
		sum += float64(grade)
	}
	return sum / float64(len(stu.Grades))
}
