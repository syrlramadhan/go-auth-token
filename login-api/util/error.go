package util

func SentPanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
