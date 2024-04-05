package models

type Lesson struct {
	Id         string `json:"id"`
	ScheduleId string `json:"schedule_id"`
	GroupId    string `json:"group_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Theme      string `json:"theme"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateLesson struct {
	ScheduleId string `json:"schedule_id"`
	GroupId    string `json:"group_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Theme      string `json:"theme"`
}

type UpdateLesson struct {
	ScheduleId string `json:"schedule_id"`
	GroupId    string `json:"group_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Theme      string `json:"theme"`
}

type GetLesson struct {
	Id         string `json:"id"`
	ScheduleId string `json:"schedule_id"`
	GroupId    string `json:"group_id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Theme      string `json:"theme"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}


type GetAllLessonsResponse struct {
	Lessons []Lesson `json:"lessons"`
	Count   int16    `json:"count"`
}

type GetAllLessonsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}
