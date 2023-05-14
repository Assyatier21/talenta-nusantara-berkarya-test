package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tnb-technical-test/models"
	"tnb-technical-test/utils/constant"
	"tnb-technical-test/utils/helper"

	"github.com/labstack/echo/v4"
)

func (h *handler) HelloDev(e echo.Context) error {
	var (
		res models.StandardResponseReq
		err error
	)
	res = models.StandardResponseReq{
		Code:    http.StatusOK,
		Message: "Hello Go Developers",
		Error:   err,
	}
	return helper.WriteResponse(e, res)
}

func (h *handler) IsPalindrome(e echo.Context) error {
	var (
		res  models.StandardResponseReq
		text string
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	text = e.FormValue("text")
	if text == "" {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   errors.New(constant.ERROR_VALIDATION_TEXT),
		}
		return helper.WriteResponse(e, res)
	}

	res = h.usecase.IsPalindrome(ctx, text)
	return helper.WriteResponse(e, res)
}

func (h *handler) GetLanguages(e echo.Context) error {
	var (
		res models.StandardResponseReq
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	res = h.usecase.GetLanguages(ctx)
	return helper.WriteResponse(e, res)
}

func (h *handler) AddLanguage(e echo.Context) error {
	var (
		req models.ProgrammingLanguage
		res models.StandardResponseReq
		err error
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	err = e.Bind(&req)
	if err != nil {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err,
		}
		return helper.WriteResponse(e, res)
	}

	res = h.usecase.AddLanguage(ctx, req)
	return helper.WriteResponse(e, res)
}

func (h *handler) GetLanguageByID(e echo.Context) error {
	var (
		res models.StandardResponseReq
		id  int
		err error
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	id, err = strconv.Atoi(e.Param("id"))
	if err != nil {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: constant.ERROR_VALIDATION_ID,
			Error:   err,
		}
		return helper.WriteResponse(e, res)
	}

	fmt.Println("e.Param(id): ", e.Param("id"))
	res = h.usecase.GetLanguageByID(ctx, id)
	return helper.WriteResponse(e, res)
}

func (h *handler) UpdateLanguage(e echo.Context) error {
	var (
		req models.ProgrammingLanguage
		res models.StandardResponseReq
		err error
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: constant.ERROR_VALIDATION_ID,
			Error:   err,
		}
		return helper.WriteResponse(e, res)
	}

	err = e.Bind(&req)
	if err != nil {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err,
		}
		return helper.WriteResponse(e, res)
	}

	res = h.usecase.UpdateLanguage(ctx, id, req)
	return helper.WriteResponse(e, res)
}

func (h *handler) DeleteLanguageByID(e echo.Context) error {
	var (
		res models.StandardResponseReq
		id  int
		err error
	)

	ctx, cancel := helper.GetContext()
	defer cancel()

	id, err = strconv.Atoi(e.Param("id"))
	if err != nil {
		res = models.StandardResponseReq{
			Code:    http.StatusBadRequest,
			Message: constant.ERROR_VALIDATION_ID,
			Error:   err,
		}
		return helper.WriteResponse(e, res)
	}

	res = h.usecase.DeleteLanguageByID(ctx, id)
	return helper.WriteResponse(e, res)
}
