package commands

type CommandType int

const (
	Start CommandType = iota
	CreatePlan
	GetPlans
)

func (commandType CommandType) Strting() string {
	return [...]string{"start", "createPlan", "getPlans"}[commandType]
}
