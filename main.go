package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/shiki-tak/notori/types"
	"github.com/tealeg/xlsx"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	json := ExcelToJson(input)
	OutputToFile(json, output)
}

func OutputToFile(input, output string) {
	file, err := os.Create(fmt.Sprintf(`./json/%s`, output))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(input))
}

func ExcelToJson(excelFileName string) string {
	exBases, exRoutes, exCosts := GenerateExData(excelFileName)

	plans := []types.Plan{}
	for _, exb := range exBases {
		plan := types.Plan{}

		plan.Base = exb
		// for route
		for _, exr := range exRoutes {
			route := types.Route{}
			if exr.ExBaseId == exb.Id {
				route.CourseTime = exr.CourseTime
				route.Points = exr.Points

				plan.Routes = append(plan.Routes, route)
			}
		}

		// for cost
		for _, exc := range exCosts {
			cost := types.Cost{}
			id := exc.ExBaseId
			if id == exb.Id {
				cost.Usecase = exc.Usecase
				cost.Price = exc.Price

				plan.Costs = append(plan.Costs, cost)
			}
		}

		plans = append(plans, plan)
	}

	return Json(plans)
}

func GenerateExData(excelFileName string) (exBases []types.ExBase, exRoutes []types.ExRoute, exCosts []types.ExCost) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	for i, sheet := range xlFile.Sheets {
		switch i {
		case 0:
			exBases = GenerateExBase(sheet)
		case 1:
			exRoutes = GenerateExRoute(sheet)
		case 2:
			exCosts = GenerateExCost(sheet)
		default:
			panic("invalid sheet")
		}
	}

	return
}

func GenerateExBase(sheet *xlsx.Sheet) []types.ExBase {
	data := GetCellData(sheet)

	exBases := make([]types.ExBase, 0)

	for _, elm := range data {
		exBase := types.ExBase{}

		i, _ := strconv.Atoi(elm[0])
		exBase.Id = i
		exBase.Title = elm[1]
		exBase.Description = elm[2]
		p, _ := strconv.Atoi(elm[3])
		exBase.Period = p

		exBases = append(exBases, exBase)
	}

	return exBases
}

func GenerateExRoute(sheet *xlsx.Sheet) []types.ExRoute {
	data := GetCellData(sheet)

	exRoutes := make([]types.ExRoute, 0)
	for _, elm := range data {
		exRoute := types.ExRoute{}

		i, _ := strconv.Atoi(elm[0])
		exRoute.Id = i
		c, _ := strconv.Atoi(elm[1])
		exRoute.CourseTime = c
		bi, _ := strconv.Atoi(elm[2])
		exRoute.ExBaseId = bi

		points := []string{}
		for _, p := range elm[3:] {
			points = append(points, p)
		}
		exRoute.Points = points

		exRoutes = append(exRoutes, exRoute)
	}

	return exRoutes
}

func GenerateExCost(sheet *xlsx.Sheet) []types.ExCost {
	data := GetCellData(sheet)

	exCosts := make([]types.ExCost, 0)
	for _, elm := range data {
		exCost := types.ExCost{}

		i, _ := strconv.Atoi(elm[0])
		exCost.Id = i
		exCost.Usecase = elm[1]
		p, _ := strconv.Atoi(elm[2])
		exCost.Price = p
		bi, _ := strconv.Atoi(elm[3])
		exCost.ExBaseId = bi

		exCosts = append(exCosts, exCost)
	}

	return exCosts
}

func GetCellData(sheet *xlsx.Sheet) [][]string {
	data := make([][]string, 0)
	for i, row := range sheet.Rows {
		if i != 0 {
			element := []string{}
			for _, cell := range row.Cells {
				text := cell.String()
				element = append(element, text)
			}
			data = append(data, element)
		}
	}
	return data
}

func Json(plans []types.Plan) string {
	jsonBytes, err := json.Marshal(plans)
	if err != nil {
		panic(err)
	}
	jsonStr := string(jsonBytes)

	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonStr), "", "	")
	if err != nil {
		panic(err)
	}
	return buf.String()

}

// func CreateSamplePlans() ([]ExBase, []ExRoute, []ExCost) {

// 	// for sample plans1
// 	b1 := ExBase{0, "sample_title1", "sample_description1", 2}
// 	r1 := ExRoute{0, 10, []string{"sp1", "sp2", "sp3"}, 0}
// 	r2 := ExRoute{1, 9, []string{"sp4", "sp5"}, 0}
// 	c1 := ExCost{0, "sample_usecase1", 1000, 0}
// 	c2 := ExCost{1, "sample_usecase2", 2000, 0}

// 	// for sample plans2
// 	b2 := ExBase{1, "sample_title2", "sample_description2", 1}
// 	r3 := ExRoute{2, 8, []string{"sp6", "sp7", "sp8", "sp9"}, 1}
// 	c3 := ExCost{2, "sample_usecase3", 3000, 1}
// 	c4 := ExCost{3, "sample_usecase4", 4000, 1}
// 	c5 := ExCost{4, "sample_usecase5", 5000, 1}

// 	return []ExBase{b1, b2}, []ExRoute{r1, r2, r3}, []ExCost{c1, c2, c3, c4, c5}
// }
