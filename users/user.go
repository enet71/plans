package users

import "main/plan"

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	ChatID    int64
	PlansList *plan.PlanList
}
