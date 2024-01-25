package life

type Work struct {
	name       string
	enjoyment  float64
	happiness  float64
	fulfilment float64
}

type Relationship struct {
	dates      float64
	enjoyment  float64
	sex        float64
	fulfilment float64
	longterm   float64
}

type Life struct {
	name          string
	age           int
	gender        string
	work          Work
	Relationships []Relationship
}

func CreateWork(name string, enjoyment float64, happiness float64, fulfilment float64) *Work {
	return &Work{name: name, enjoyment: enjoyment, happiness: happiness, fulfilment: fulfilment}
}

func CreateRelationship(dates float64, enjoyment float64, sex float64, fulfilment float64, longterm float64) *Relationship {
	return &Relationship{dates: dates, enjoyment: enjoyment, sex: sex, fulfilment: fulfilment, longterm: longterm}
}

func CreateLife(name string, age int, kind string, nameWork string, enjoyment float64, happiness float64, fulfilmentWork float64, dates float64, enjoymentRelationship float64, sex float64, fulfilmentRelationship float64, longterm float64) *Life {
	return &Life{name, age, kind, *CreateWork(nameWork, enjoyment, happiness, fulfilmentWork), []Relationship{*CreateRelationship(dates, enjoymentRelationship, sex, fulfilmentRelationship, longterm)}}
}
