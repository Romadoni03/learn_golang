package helper

func IfPanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicWithMessage(err error, message string) {
	if err != nil {
		panic(message)
	}
}
