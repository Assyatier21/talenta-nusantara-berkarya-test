package helper

import (
	"net/http"

	"tnb-technical-test/models"
	"tnb-technical-test/utils/constant"

	"github.com/labstack/echo/v4"
)

func WriteResponse(c echo.Context, req models.StandardResponseReq) error {
	var status = constant.SUCCESS

	if req.Code > 299 && req.Message != constant.SUCCESS {
		status = constant.FAILED
	} else if req.Code == 404 && req.Message == constant.SUCCESS {
		status = constant.SUCCESS
		req.Message = constant.ERROR_NO_DATA
	}

	var errResp interface{}
	if req.Error != nil {
		errResp = req.Error.Error()
	}

	if req.Message == "" {
		req.Message = http.StatusText(req.Code)
	}

	return c.JSON(req.Code, models.StandardResponse{
		Code:    req.Code,
		Status:  status,
		Message: req.Message,
		Data:    req.Data,
		Error:   errResp,
	})
}
