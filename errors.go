package errors

func check(err error) {
	if err != nil {
		panic(err)
	}
}
