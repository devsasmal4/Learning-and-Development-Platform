package error

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleNoRouteError() Error {
	return Error{
		Code:    "PAGE_NOT_FOUND",
		Message: "Page not found",
	}
}

func HandleUnauthorizedError(err error) Error {
	return Error{
		Code:    "ERR_UNAUTHORIZED",
		Message: err.Error(),
	}
}
