package plan

import "time"

type Plan struct {
	name      string
	startDate time.Time
	endDate   time.Time
	status    PlanStatus
	tasks     []Task
}

func CreatePlan() (plan Plan) {
	plan = Plan{
		startDate: time.Now(),
		endDate:   time.Now(),
	}
	return
}
