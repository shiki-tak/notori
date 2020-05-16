package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// For json data
type Plan struct {
	Base   ExBase  `json:"base"`
	Routes []Route `json:"Routes"`
	Costs  []Cost  `json:"Costs"`
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
	Id          int    `josn:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Period      int    `json:"period"`
}

type ExRoute struct {
	Id         int `json:"id"`
	CourseTime int `json:"course_time"`
	BaseId     int `json:"base_id"`
}

type ExPoint struct {
	RouteId int    `json:"route_id"`
	Name    string `json:"point"`
}

type ExCost struct {
	ExBaseId int    `json:"id"`
	Usecase  string `json:"usecase"`
	Price    int    `json:"price"`
}

func main() {
	exBase, exRoutes, exPoins, exCosts := CreateSamplePlans()

	plans := []Plan{}
	for _, exb := range exBase {
		plan := Plan{}

		plan.Base = exb
		// for route
		for _, exr := range exRoutes {
			route := Route{}
			id := exr.BaseId
			if id == exb.Id {
				route.CourseTime = exr.CourseTime
				for _, exp := range exPoins {
					if exr.Id == exp.RouteId {
						route.Points = append(route.Points, exp.Name)
					}
				}
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

func CreateSamplePlans() ([]ExBase, []ExRoute, []ExPoint, []ExCost) {

	// for sample plans1
	b1 := ExBase{0, "sample_title1", "sample_description1", 2}
	r1 := ExRoute{0, 10, 0}
	r2 := ExRoute{1, 9, 0}
	p1 := ExPoint{0, "sp1"}
	p2 := ExPoint{0, "sp2"}
	p3 := ExPoint{0, "sp3"}
	p4 := ExPoint{1, "sp4"}
	p5 := ExPoint{1, "sp5"}
	c1 := ExCost{0, "sample_usecase1", 1000}
	c2 := ExCost{0, "sample_usecase2", 2000}

	// for sample plans2
	b2 := ExBase{1, "sample_title2", "sample_description2", 1}
	r3 := ExRoute{2, 8, 1}
	p6 := ExPoint{2, "sp6"}
	p7 := ExPoint{2, "sp7"}
	p8 := ExPoint{2, "sp8"}
	p9 := ExPoint{2, "sp9"}
	c3 := ExCost{1, "sample_usecase3", 3000}
	c4 := ExCost{1, "sample_usecase4", 4000}
	c5 := ExCost{1, "sample_usecase5", 5000}

	return []ExBase{b1, b2}, []ExRoute{r1, r2, r3}, []ExPoint{p1, p2, p3, p4, p5, p6, p7, p8, p9}, []ExCost{c1, c2, c3, c4, c5}
}
