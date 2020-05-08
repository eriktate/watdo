package service

import "github.com/eriktate/wrkhub"

// A Service aggregates all of the service implementations.
type Service struct {
	AccountService
	ProjectService
	UserService
	TaskService
}

func NewService(accountStore wrkhub.AccountStore, projectStore wrkhub.ProjectStore, userStore wrkhub.UserStore, taskStore wrkhub.TaskStore) Service {
	return Service{
		AccountService: NewAccountService(accountStore),
		ProjectService: NewProjectService(projectStore),
		UserService:    NewUserService(userStore),
		TaskService:    NewTaskService(taskStore),
	}
}
