package models

type Task struct {
	Id        string `json:"id"`
	LessonId  string `json:"lesson_id"`
	GroupId   string `json:"task"`
	Score     string `json:"score"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateTask struct {
	LessonId  string `json:"lesson_id"`
	GroupId   string `json:"task"`
	Score     string `json:"score"`
}

type UpdateTask struct {
	LessonId  string `json:"lesson_id"`
	GroupId   string `json:"task"`
	Score     string `json:"score"`
}

type GetTask struct {
	Id        string `json:"id"`
	LessonId  string `json:"lesson_id"`
	GroupId   string `json:"task"`
	Score     string `json:"score"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllTasksResponse struct {
	Tasks []Task `json:"tasks"`
	Count int16  `json:"count"`
}

type GetAllTasksRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

