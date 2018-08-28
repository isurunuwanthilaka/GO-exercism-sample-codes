package twofer

func ShareWith(name string) string {
	/*
		Generating some text for ShareWith someone
	*/
	var s string

	if name != "" {
		s = "One for " + name + ", one for me."
	} else {
		s = "One for you, one for me."
	}

	return s
}
