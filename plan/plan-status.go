package plan

type PlanStatus int

const (
	Exist PlanStatus = iota
	InProgress
	Finished
)

func (status PlanStatus) String() string {
	return [...]string{"exist", "inProgress", "finished"}[status]
}
