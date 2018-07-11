if v, ok := varI.(T); ok { // checked type assertion
	Process(v)
	return
}
// varI is not of type 
