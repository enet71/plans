package main

import (
	"fmt"
	"main/plan"
)

func main() {
	plan.GetList().AddPlan(plan.CreatePlan())
	fmt.Println(plan.GetList())
}
