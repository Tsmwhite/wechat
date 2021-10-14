package error

const TokenMissing = 4000
const TokenError = 4001
const TokenExpired = 4002

var errorString = map[int]string{
	TokenMissing: "Missing token",
	TokenError:   "Incorrect token",
	TokenExpired: "Token has expired",
}

func Output(code int) string {
	return errorString[code]
}
