package language

import (
	"context"
	m "tnb-technical-test/models"
)

type Repository interface {
	HelloDev(ctx context.Context) string
	GetLanguages(ctx context.Context) *[]m.ProgrammingLanguage
	AddLanguage(ctx context.Context, lang m.ProgrammingLanguage) *[]m.ProgrammingLanguage
	GetLanguageByID(ctx context.Context, id int) m.ProgrammingLanguage
	UpdateLanguage(ctx context.Context, id int, lang m.ProgrammingLanguage) m.ProgrammingLanguage
	DeleteLanguageByID(ctx context.Context, id int) *[]m.ProgrammingLanguage
}

type repository struct {
	db *[]m.ProgrammingLanguage
}

func NewRepository(db *[]m.ProgrammingLanguage) Repository {
	return &repository{
		db: db,
	}
}
