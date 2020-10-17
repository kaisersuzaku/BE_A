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
)

func GetInvalidPayloadResp() RespError {
	return invalidPayload
}
