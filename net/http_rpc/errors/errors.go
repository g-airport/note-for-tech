package errors

import (
	"encoding/json"
	errmicro "github.com/micro/go-micro/errors"
	"github.com/prometheus/common/log"
	"strings"
)

type Error struct {
	Code     int         `json:"code"`
	Status   int         `json:"status"`
	Detail   string      `json:"detail"`
	Internal string      `json:"internal,omitempty"`
	Content  interface{} `json:"content,omitempty"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func Parse(err string) *Error {
	e := new(Error)
	error := json.Unmarshal([]byte(err), e)
	if error != nil {
		e.Detail = err
	}
	return e
}

func ParseFromRPCError(err error) error {
	if err == nil {
		return nil
	}
	errtxt := err.Error()

	merr := errmicro.Parse(errtxt)
	errtxt = merr.Detail

	if idx := strings.Index(errtxt, "{"); idx != -1 {
		errtxt = errtxt[idx:]
	}

	return Parse(errtxt)
}

var Errors = map[int]*Error{}

func addError(err *Error) *Error {
	e, ok := Errors[err.Code]
	if ok {
		log.Fatalf("duplate error code: %v, %v", e, err)
	}
	Errors[err.Code] = err
	return err
}

//常用错误
func BadRequest(code int, detail string) error {
	return addError(&Error{
		Code:code,
		Status:400,
		Detail:detail,
	})
}
