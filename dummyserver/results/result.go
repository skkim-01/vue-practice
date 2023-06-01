package results

import (
	JsonMapper "github.com/skkim-01/json-mapper/src"
)

var KEY_ROOT string = ""
var KEY_RESULT_BASE string = "result"
var KEY_ERROR_CODE string = "code"
var KEY_ERROR_MESSAGE string = "msg"

var ERROR_UNHANDLED = -1
var ERROR_SUCCESS = 0

// ApiResult struct
type ApiResult struct {
	jsonError *JsonMapper.JsonMap
}

// NewError: create api error object
func NewResult() *ApiResult {
	e := &ApiResult{}
	e.jsonError, _ = JsonMapper.NewBytes([]byte(`{"code":0,"result":{}}`))
	return e
}

// Parse: parse return string
func Parse(result string) (int, map[string]interface{}) {
	m, err := JsonMapper.NewBytes([]byte(result))
	if nil != err {
		return -1, nil
	}
	return int(m.Find(KEY_ERROR_CODE).(float64)), m.Find(KEY_RESULT_BASE).(map[string]interface{})
}

// Success: make result: success json object
//   - params : - k, v, k, v
func (e *ApiResult) Success(msg ...interface{}) {
	e.jsonError.Insert(KEY_ROOT, KEY_ERROR_CODE, ERROR_SUCCESS)
	var k string = ""
	var v interface{} = nil

	for i := range msg {
		if i%2 == 0 { // key
			k = msg[i].(string)
		} else { // value
			v = msg[i]
			e.jsonError.Insert(KEY_RESULT_BASE, k, v)
			k = ""
			v = nil
		}
	}
}

// Error: make result: error json object
func (e *ApiResult) Error(code int, msg interface{}) {
	e.jsonError.Insert(KEY_ROOT, KEY_ERROR_CODE, code)
	e.jsonError.Insert(KEY_RESULT_BASE, KEY_ERROR_MESSAGE, msg)
}

// Build: json object to string
func (e *ApiResult) Build() string {
	return e.jsonError.Print()
}
