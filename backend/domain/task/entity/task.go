package entity

type Task struct {
	ID int64

	Content  string
	Priority string
	TaskTyp  string // overdue | schedule | wait to be done | completed
	Date     string
}
