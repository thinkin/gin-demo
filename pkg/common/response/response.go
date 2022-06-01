package response

const (
	DefaultErrno = 1
)

// Data indicate response data struct
type Data struct {
	Errno  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// Success return successful response
func Success(data interface{}) Data {
	return Data{
		Errno:  0,
		ErrMsg: "",
		Data:   data,
	}
}

// Error return failed response
// eg.
// 	Error(err)
// 	Error(err, data)
// 	Error(err, data, errno)
func Error(err error, args ...interface{}) Data {
	var errno int
	var data interface{}
	switch len(args) {
	case 1:
		data = args[0]
		errno = DefaultErrno
	case 2:
		data = args[0]
		errno = args[1].(int)
	default:
		data = nil
		errno = DefaultErrno
	}

	return Data{
		Errno:  errno,
		ErrMsg: err.Error(),
		Data:   data,
	}
}
