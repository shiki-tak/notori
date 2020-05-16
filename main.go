package main

import (
	"fmt"
	"os"

	"github.com/shiki-tak/notori/handler"
)

const OUTPUT_PATH = "./json"

func main() {
	if err := Run(os.Args[1], os.Args[2]); err != nil {
		fmt.Println(err)
	}
}

func Run(input, output string) error {
	json, err := handler.ExcelToJson(input)
	if err != nil {
		return err
	}

	err = OutputToFile(json, output)
	if err != nil {
		return err
	}

	return nil
}

func OutputToFile(input, output string) error {
	file, err := os.Create(fmt.Sprintf(`%s/%s`, OUTPUT_PATH, output))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(([]byte)(input))
	if err != nil {
		return err
	}

	return nil

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
