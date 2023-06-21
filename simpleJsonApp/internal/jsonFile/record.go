package jsonFile

type Record struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Designation string   `json:"designation"`
	Location    string   `json:"location"`
	Todos       []string `json:"todos"`
}

func ValidateRecord(fileContents []Record, record Record) bool {
	for _, rec := range fileContents {
		if rec.Name == record.Name {
			return false
		}
	}
	return true
}
