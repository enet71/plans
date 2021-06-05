package plan

var planList = PlanList{}

type PlanList struct {
	Plans []Plan
}

func (list *PlanList) AddPlan(plan Plan) {
	(*list).Plans = append(list.Plans, plan)
}

func GetList() *PlanList {
	return &planList
}
