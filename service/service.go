package service

import "github.com/eriktate/wrkhub"

// A Service aggregates all of the service implementations.
type Service struct {
	AccountService
	ProjectService
}

func NewService(accountStore wrkhub.AccountStore, projectStore wrkhub.ProjectStore) Service {
	return Service{
		AccountService: NewAccountService(accountStore),
		ProjectService: NewProjectService(projectStore),
	}
}
