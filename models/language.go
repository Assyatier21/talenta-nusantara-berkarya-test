package models

type LanguageRelation struct {
	InfluencedBy []string `json:"influenced_by"`
	Influences   []string `json:"influences"`
}

type ProgrammingLanguage struct {
	Language       string           `json:"language"`
	Appeared       int              `json:"appeared"`
	Created        []string         `json:"created"`
	Functional     bool             `json:"functional"`
	ObjectOriented bool             `json:"object_oriented"`
	Relation       LanguageRelation `json:"relation"`
}
