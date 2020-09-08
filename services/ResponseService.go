package services

import (
	"encoding/json"
	"github.com/trello-analog/backend/customerrors"
	"github.com/trello-analog/backend/entity"
	"net/http"
)

type ResponseService struct {
	Data   interface{}         `json:"data"`
	Writer http.ResponseWriter `json:"writer"`
}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

func (rs *ResponseService) SetData(data interface{}) *ResponseService {
	rs.Data = data
	return rs
}

func (rs *ResponseService) SetWriter(writer http.ResponseWriter) *ResponseService {
	rs.Writer = writer
	return rs
}

func (rs *ResponseService) Error() {
	d := rs.Data.(*customerrors.APIError)
	rs.Writer.WriteHeader(d.Status)
	json.NewEncoder(rs.Writer).Encode(entity.NewErrorResponse(d))
}

func (rs *ResponseService) Success() {
	json.NewEncoder(rs.Writer).Encode(entity.NewSuccessResponse(rs.Data))

}
