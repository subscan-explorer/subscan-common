package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/subscan-explorer/subscan-common/core/ecode"
)

// Resp .
type Resp struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// JSONResponse .
func JSONResponse(c *gin.Context, httpStatus int, ecode int, msg string, data interface{}) {
	resp := &Resp{
		Code:    ecode,
		Data:    data,
		Message: msg,
	}
	c.JSON(httpStatus, resp)
}

// JSONSuccess .
func JSONSuccess(c *gin.Context, data interface{}) {
	JSONResponse(c, http.StatusOK, 0, "success", data)
}

// JSONFail .
func JSONFail(c *gin.Context, err error) {
	var (
		ec ecode.Codes
		ok bool
	)
	if err != nil {
		ec, ok = errors.Cause(err).(ecode.Codes)
		if ok {
			JSONResponse(c, http.StatusOK, ec.Code(), ec.Message(), nil)
		} else {
			JSONResponse(c, http.StatusOK, ecode.ServerErr.Code(), err.Error(), nil)
		}
	}
}

// AbortWithJSONResponse .
func AbortWithJSONResponse(c *gin.Context, httpStatus int, ecode int, msg string, data interface{}) {
	resp := &Resp{
		Code:    ecode,
		Data:    data,
		Message: msg,
	}
	c.Abort()
	c.JSON(httpStatus, resp)
}

// AbortWithJSONSuccess .
func AbortWithJSONSuccess(c *gin.Context, data interface{}) {
	AbortWithJSONResponse(c, http.StatusOK, 0, "success", data)
}

// AbortWithJSONFail .
func AbortWithJSONFail(c *gin.Context, ecode int, msg string) {
	AbortWithJSONResponse(c, http.StatusOK, ecode, msg, nil)
}
