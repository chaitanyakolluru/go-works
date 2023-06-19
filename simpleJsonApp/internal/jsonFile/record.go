package jsonFile

type Record struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Designation string   `json:"designation"`
	Location    string   `json:"location"`
	Todos       []string `json:"todos"`
}
