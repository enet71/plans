package plan

import (
	"main/utils"
	"time"
)

type Plan struct {
	Uuid      string
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    PlanStatus
	Tasks     []Task
}

func CreatePlan(name string, startDate time.Time, endDate time.Time) (plan Plan) {
	plan = Plan{
		Uuid:      utils.GenUuid(),
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
	}
	return
}
