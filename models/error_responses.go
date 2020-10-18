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
	productUnavailable RespError = RespError{
		ErrorCode:    "ER-STK-01",
		ErrorMessage: "Product is Unavailable (stock is 0)",
	}
	stockLessThanRequest RespError = RespError{
		ErrorCode:    "ER-STK-02",
		ErrorMessage: "Requested item quantity more than available stock",
	}
	requestFailed RespError = RespError{
		ErrorCode:    "ER-GNR-00",
		ErrorMessage: "Oops, Something went wrong, please try again",
	}
	productNotFound RespError = RespError{
		ErrorCode:    "ER-STK-03",
		ErrorMessage: "Product not found",
	}
)

func GetInvalidPayloadResp() RespError {
	return invalidPayload
}

func GetUnhandledRequest() RespError {
	return unhandledRequest
}

func GetProductUnavailable() RespError {
	return productUnavailable
}

func GetStockLessThanRequest() RespError {
	return stockLessThanRequest
}

func GetRequstFailed() RespError {
	return requestFailed
}

func GetProductNotFound() RespError {
	return productNotFound
}
