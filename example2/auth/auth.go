package auth

func CheckLogin(
	username string,
	password string,
	tokenKey string,
) (validation string) {

	chkName := username
	chkPassword := password
	chkToken := tokenKey

	if chkName == "admin" && chkPassword == "1234" && chkToken == "12345678" {
		validation = showMessage("ยินดีต้อนรับ : " + chkName)
	} else {
		validation = showMessage("username or password ไม่ถูกต้อง")
	}

	return
}

func showMessage(msg string) string {
	return msg
}
