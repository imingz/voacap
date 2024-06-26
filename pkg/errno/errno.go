package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

const (
	SuccessCode    = 200
	ServiceErrCode = iota + 10000
	ParamErrCode
)

const (
	SuccessMsg   = "Success"
	ServerErrMsg = "Service is unable to start successfully"
	ParamErrMsg  = "Wrong Parameter has been given"
)

var (
	Success    = ErrNo{SuccessCode, SuccessMsg}
	ServiceErr = ErrNo{ServiceErrCode, ServerErrMsg}
	ParamErr   = ErrNo{ParamErrCode, ParamErrMsg}
)
