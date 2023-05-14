package language

import (
	"context"
	m "tnb-technical-test/models"
	"tnb-technical-test/utils/constant"
)

func (r *repository) HelloDev(ctx context.Context) string {
	var (
		res string = "Hello Go Developers"
	)

	return res
}

func (r *repository) GetLanguages(ctx context.Context) *[]m.ProgrammingLanguage {
	return &constant.GlobalLanguage
}

func (r *repository) AddLanguage(ctx context.Context, lang m.ProgrammingLanguage) *[]m.ProgrammingLanguage {
	constant.GlobalLanguage = append(constant.GlobalLanguage, lang)
	return &constant.GlobalLanguage
}

func (r *repository) GetLanguageByID(ctx context.Context, id int) m.ProgrammingLanguage {
	return constant.GlobalLanguage[id]
}

func (r *repository) UpdateLanguage(ctx context.Context, id int, lang m.ProgrammingLanguage) m.ProgrammingLanguage {
	constant.GlobalLanguage[id] = lang
	return constant.GlobalLanguage[id]
}

func (r *repository) DeleteLanguageByID(ctx context.Context, id int) *[]m.ProgrammingLanguage {
	constant.GlobalLanguage = append(constant.GlobalLanguage[:id], constant.GlobalLanguage[id+1:]...)
	return &constant.GlobalLanguage
}
