package plan

import (
	"main/files"
	"os"
	"strconv"
)

const PATH = "plans-lists"

type PlansList struct {
	UserId int
	Plans  []Plan
}

func CreatePlansList(UserId int) PlansList {
	plansList := PlansList{
		UserId: UserId,
	}

	files.WriteFile(PATH, strconv.Itoa(UserId), plansList)

	return plansList
}

func GetPlansList(userId int) PlansList {
	plansList := &PlansList{}
	files.ReadFile(plansList, PATH, strconv.Itoa(userId))

	return *plansList
}

func UpdatePlansList(plansList PlansList) {
	files.WriteFile(PATH, strconv.Itoa(plansList.UserId), plansList)
}

func GetPlansLists() (plansLists []*PlansList) {
	files.ReadAllFiles(PATH, func(file os.FileInfo) interface{} {
		plansList := &PlansList{}
		plansLists = append(plansLists, plansList)
		return plansList
	})

	return plansLists
}

func (list *PlansList) AddPlan(plan Plan) {
	(*list).Plans = append(list.Plans, plan)
	UpdatePlansList(*list)
}

func (list *PlansList) GetLength() int {
	return len(list.Plans)
}
