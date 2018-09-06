package produce

type Produce struct {
	Name        string  `json:"Name"`
	ProduceCode string  `json:"ProduceCode"`
	UnitPrice   float64 `json:"UnitPrice"`
}


type Inventory []Produce

//type ProduceDB struct {
//	Data  []*Produce
//	Cache map[string]bool
//	Lock  sync.RWMutex
//}

var Array = Inventory{
	Produce{
		"Lettuce",
		"A12T-4GH7-QPL9-3N4M",
		3.46,
	},
	Produce{
		"Peach",
		"E5T6-9UI3-TH15-QR88",
		2.99,
	},
	Produce{
		"Green Pepper",
		"YRT6-72AS-K736-L4AR",
		0.79,
	},
	Produce{
		"Gala Apple",
		"TQ4C-VV6T-75ZX-1RMR",
		3.59,
	},
}
//TODO Mutex locking
//var newArray = ProduceDB{
//	Data: []*Produce{
//		&Produce{
//			Name: "test",
//		},
//	},
//}