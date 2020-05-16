package types

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
	ExBaseId   int
}

type ExCost struct {
	Id       int
	Usecase  string
	Price    int
	ExBaseId int
}
