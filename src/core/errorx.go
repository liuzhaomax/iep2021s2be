/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/8 13:53
 * @version     v1.0
 * @filename    errorx.go
 * @description
 ***************************************************************************/
package core

var ErrorMsg = map[int]string{
	100: "Duplicate user email.",
	101: "Invalid user email and password.",
	102: "Invalid old password.",
	103: "Duplicate nick name",
	104: "No searched user",
	105: "Email and password are not matched.",
	106: "Token generating failed.",
	107: "Encrypted error.",
	108: "Decrypted error.",
	109: "Token expired.",
	110: "Invalid token.",

	200: "File open failed.",
	201: "File reading failed.",

	300: "Create blog failed.",
	301: "Update blog failed.",
	302: "Delete blog failed.",

	900: "Unknown error.",
	901: "Request data parsing failed",
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Err     error  `json:"error"`
}

func (err *Error) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func NewError(errorCode int, err error) *Error {
	var errObj = new(Error)
	errObj.Code = errorCode
	errObj.Message = ErrorMsg[errorCode]
	if err != nil {
		errObj.Err = err
	}
	return errObj
}
