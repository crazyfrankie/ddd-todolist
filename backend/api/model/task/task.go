package task

type CreateTaskRequest struct {
	Content  string  `json:"content,omitempty" binding:"required"`
	Date     *int64  `json:"date,omitempty"`
	Priority *string `json:"priority,omitempty"`
}

type UpdateTaskRequest struct {
	TaskID      int64   `json:"task_id"`
	Content     *string `json:"content,omitempty"`
	Priority    *string `json:"priority"`
	Date        *int64  `json:"date"`
	IsCompleted *bool   `json:"isCompleted"`
}

type Task struct {
	ID int64 `json:"id"`

	Content string `json:"content"`
	Date    string `json:"date"`
	TaskTyp string `json:"taskTyp"`
}

type TaskItem struct {
	Content string `json:"content"`
	TaskTyp string `json:"taskTyp"`
}
