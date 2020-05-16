package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/shiki-tak/notori/types"
	"github.com/tealeg/xlsx"
)

func ExcelToJson(excelFileName string) (string, error) {
	exBases, exRoutes, exCosts, err := GenerateExData(excelFileName)
	if err != nil {
		return "", err
	}

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

func GenerateExData(excelFileName string) ([]types.ExBase, []types.ExRoute, []types.ExCost, error) {
	exBases := []types.ExBase{}
	exRoutes := []types.ExRoute{}
	exCosts := []types.ExCost{}

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return exBases, exRoutes, exCosts, err
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
			return exBases, exRoutes, exCosts, errors.New("invalid sheet")
		}
	}

	return exBases, exRoutes, exCosts, nil
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

func Json(plans []types.Plan) (string, error) {
	jsonBytes, err := json.Marshal(plans)
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)

	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonStr), "", "	")
	if err != nil {
		return "", nil
	}
	return buf.String(), nil

}
