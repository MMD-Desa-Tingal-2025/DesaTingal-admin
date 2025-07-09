package model

type Region struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	MalePopulation    int    `json:"male_population"`
	FemalePopulation  int    `json:"female_population"`
	TotalPopulation   int    `json:"total_population"`
	PopulationDensity float64 `json:"population_density"`
	Area             float64 `json:"area"`
	AverageAge       float64 `json:"average_age"`
	RegionCode       string `json:"region_code"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func (r *Region) CalculateTotalPopulation() {
	r.TotalPopulation = r.MalePopulation + r.FemalePopulation
}