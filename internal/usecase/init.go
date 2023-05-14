package usecase

import (
	"context"
	"tnb-technical-test/internal/repository/language"
	m "tnb-technical-test/models"
)

type UsecaseHandler interface {
	HelloDev(ctx context.Context) m.StandardResponseReq
	IsPalindrome(ctx context.Context, text string) m.StandardResponseReq
	GetLanguages(ctx context.Context) m.StandardResponseReq
	AddLanguage(ctx context.Context, lang m.ProgrammingLanguage) m.StandardResponseReq
	GetLanguageByID(ctx context.Context, id int) m.StandardResponseReq
	UpdateLanguage(ctx context.Context, id int, req m.ProgrammingLanguage) m.StandardResponseReq
	DeleteLanguageByID(ctx context.Context, id int) m.StandardResponseReq
}

type usecase struct {
	repository language.Repository
}

func NewUsecase(repository language.Repository) UsecaseHandler {
	return &usecase{
		repository: repository,
	}
}
