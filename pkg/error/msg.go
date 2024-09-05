package errpkg

var MsgFlags = map[Code]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Request parameter error",

	ERROR_AUTH_GENERATE_TOKEN_FAIL: "Failed to generate token",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Failed to check token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token expired",
	ERROR_AUTH_LOGIN_FAIL:          "Login failed",
	ERROR_AUTH_USERNAME_EXIST:      "Username already exists",
	ERROR_AUTH_USER_NOT_EXIST:      "User does not exist",
	ERROR_AUTH_PASSWORD_INCORRECT:  "Password is incorrect",
	ERROR_AUTH_REGISTER_FAIL:       "Register failed",
}

// GetMsg get error information based on Code
func GetMsg(code Code) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
