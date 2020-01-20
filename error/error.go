package error

// Check -- Simple Check function with Panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
