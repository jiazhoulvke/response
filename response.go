package response

//Errors define errors
var Errors = map[int]string{
	418: "I'm a teapot", //example
}

//Fail fail
func Fail(errCode int, errMsg ...string) map[string]interface{} {
	if errCode == 0 {
		errCode = 400
	}
	resp := map[string]interface{}{
		"status":   "fail",
		"err_code": errCode,
	}
	if len(errMsg) > 0 {
		resp["err_msg"] = errMsg[0]
	} else {
		if msg, ok := Errors[errCode]; ok {
			resp["err_msg"] = msg
		}
	}
	return resp
}

//Succ success
func Succ(data map[string]interface{}) map[string]interface{} {
	resp := map[string]interface{}{
		"status": "succ",
	}
	if data != nil {
		for k, v := range data {
			resp[k] = v
		}
	}
	return resp
}
