package constant

import "tnb-technical-test/models"

var (
	GlobalLanguage = []models.ProgrammingLanguage{
		{
			Language:       "C",
			Appeared:       1972,
			Created:        []string{"Dennis Ritchie"},
			Functional:     false,
			ObjectOriented: false,
			Relation: models.LanguageRelation{
				InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
				Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
			},
		},
	}
)
