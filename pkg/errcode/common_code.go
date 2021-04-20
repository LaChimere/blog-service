package errcode

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "server internal error")
	InvalidParams             = NewError(10000001, "invalid parameters")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "unauthorized, AppKey and AppSecret not found")
	UnauthorizedTokenError    = NewError(10000004, "unauthorized, token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "unauthorized, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "unauthorized, token generation failed")
	TooManyRequests           = NewError(10000007, "too many requests")
)
