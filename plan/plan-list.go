package plan

type PlanList struct {
	Plans []Plan
}

func (list *PlanList) AddPlan(plan Plan) {
	(*list).Plans = append(list.Plans, plan)
}

func (list *PlanList) GetLength() int {
	return len(list.Plans)
}
