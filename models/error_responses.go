package models

type RespError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

var (
	invalidPayload RespError = RespError{
		ErrorCode:    "ER-FBC-01",
		ErrorMessage: "Invalid Payload",
	}
	unhandledRequest RespError = RespError{
		ErrorCode:    "ER-FBC-02",
		ErrorMessage: "Unhandled Request, can only process request with ball container size 3 or 4",
	}
)

func GetInvalidPayloadResp() RespError {
	return invalidPayload
}

func GetUnhandledRequest() RespError {
	return unhandledRequest
}
