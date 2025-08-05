package entity

type TaskStatus string

const (
	TaskOverDue   TaskStatus = "overdue"
	TaskSchedule  TaskStatus = "schedule"
	TaskToBeDone  TaskStatus = "wait to be done"
	TaskCompleted TaskStatus = "completed"
)

func (t TaskStatus) String() string {
	return string(t)
}

type Task struct {
	ID int64

	Content  string
	Priority string
	TaskTyp  TaskStatus // overdue | schedule | wait to be done | completed
	Date     string
}
