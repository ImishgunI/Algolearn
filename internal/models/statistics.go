package models

type AdminStatistics struct {
	UsersCount    int `json:"usersCount"`
	LessonsCount  int `json:"lessonsCount"`
	CommentsCount int `json:"commentsCount"`
	TasksCount    int `json:"tasksCount"`
}
