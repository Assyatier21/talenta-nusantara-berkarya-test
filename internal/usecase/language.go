package usecase

import (
	"context"
	"net/http"
	m "tnb-technical-test/models"
	"tnb-technical-test/utils/constant"
	"tnb-technical-test/utils/helper"
)

func (u *usecase) HelloDev(ctx context.Context) m.StandardResponseReq {
	message := u.repository.HelloDev(ctx)
	return m.StandardResponseReq{Code: http.StatusOK, Message: message, Data: nil, Error: nil}
}

func (u *usecase) IsPalindrome(ctx context.Context, text string) m.StandardResponseReq {
	if helper.IsPalindrome(text) {
		return m.StandardResponseReq{Code: http.StatusOK, Message: constant.PALINDROME}
	}
	return m.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.NOT_PALINDROME}
}

func (u *usecase) GetLanguages(ctx context.Context) m.StandardResponseReq {
	var (
		language = &[]m.ProgrammingLanguage{}
	)
	language = u.repository.GetLanguages(ctx)

	return m.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: language}
}

func (u *usecase) AddLanguage(ctx context.Context, req m.ProgrammingLanguage) m.StandardResponseReq {
	var (
		language = &[]m.ProgrammingLanguage{}
	)

	language = u.repository.AddLanguage(ctx, req)
	return m.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: language}
}

func (u *usecase) GetLanguageByID(ctx context.Context, id int) m.StandardResponseReq {
	var (
		language = m.ProgrammingLanguage{}
	)

	if id >= len(constant.GlobalLanguage) || id < 0 {
		return m.StandardResponseReq{Code: http.StatusNotFound, Message: constant.SUCCESS, Data: nil}
	}

	language = u.repository.GetLanguageByID(ctx, id)
	return m.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: language}
}

func (u *usecase) UpdateLanguage(ctx context.Context, id int, lang m.ProgrammingLanguage) m.StandardResponseReq {
	if id >= len(constant.GlobalLanguage) || id < 0 {
		return m.StandardResponseReq{Code: http.StatusNotFound, Message: constant.SUCCESS, Data: nil}
	}

	updatedLang := u.repository.UpdateLanguage(ctx, id, lang)
	return m.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: updatedLang}
}

func (u *usecase) DeleteLanguageByID(ctx context.Context, id int) m.StandardResponseReq {
	if id >= len(constant.GlobalLanguage) || id < 0 {
		return m.StandardResponseReq{Code: http.StatusNotFound, Message: constant.SUCCESS, Data: nil}
	}

	lang := u.repository.DeleteLanguageByID(ctx, id)
	return m.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: lang}
}
