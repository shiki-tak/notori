package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// For json data
type Plan struct {
	Base   ExBase  `json:"base"`
	Routes []Route `json:"routes"`
	Costs  []Cost  `json:"costs"`
}

type Route struct {
	Points     []string `json:"points"`
	CourseTime int      `json:"course_time"`
}

type Cost struct {
	Usecase string `json:"usecase"`
	Price   int    `json:"price"`
}

// For excel data
type ExBase struct {
	Id          int
	Title       string
	Description string
	Period      int
}

type ExRoute struct {
	Id         int
	CourseTime int
	Points     []string
	BaseId     int
}

type ExCost struct {
	ExBaseId int
	Usecase  string
	Price    int
}

func main() {
	exBase, exRoutes, exCosts := CreateSamplePlans()

	plans := []Plan{}
	for _, exb := range exBase {
		plan := Plan{}

		plan.Base = exb
		// for route
		for _, exr := range exRoutes {
			route := Route{}
			if exr.BaseId == exb.Id {
				route.CourseTime = exr.CourseTime
				route.Points = exr.Points

				plan.Routes = append(plan.Routes, route)
			}
		}

		// for cost
		for _, exc := range exCosts {
			cost := Cost{}
			id := exc.ExBaseId
			if id == exb.Id {
				cost.Usecase = exc.Usecase
				cost.Price = exc.Price

				plan.Costs = append(plan.Costs, cost)
			}
		}

		plans = append(plans, plan)
	}

	jsonBytes, err := json.Marshal(plans)
	if err != nil {
		fmt.Print(err)
		return
	}
	jsonStr := string(jsonBytes)

	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonStr), "", "  ")
	if err != nil {
		panic(err)
	}
	indentJSON := buf.String()

	fmt.Println(indentJSON)
}

func CreateSamplePlans() ([]ExBase, []ExRoute, []ExCost) {

	// for sample plans1
	b1 := ExBase{0, "sample_title1", "sample_description1", 2}
	r1 := ExRoute{0, 10, []string{"sp1", "sp2", "sp3"}, 0}
	r2 := ExRoute{1, 9, []string{"sp4", "sp5"}, 0}
	c1 := ExCost{0, "sample_usecase1", 1000}
	c2 := ExCost{0, "sample_usecase2", 2000}

	// for sample plans2
	b2 := ExBase{1, "sample_title2", "sample_description2", 1}
	r3 := ExRoute{2, 8, []string{"sp6", "sp7", "sp8", "sp9"}, 1}
	c3 := ExCost{1, "sample_usecase3", 3000}
	c4 := ExCost{1, "sample_usecase4", 4000}
	c5 := ExCost{1, "sample_usecase5", 5000}

	return []ExBase{b1, b2}, []ExRoute{r1, r2, r3}, []ExCost{c1, c2, c3, c4, c5}
}
