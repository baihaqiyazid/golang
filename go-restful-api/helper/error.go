package helper

func Panic(err error)  {
	if err != nil {
		Panic(err)
	}
}