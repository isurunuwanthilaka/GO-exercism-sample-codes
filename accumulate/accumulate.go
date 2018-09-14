package accumulate

func Accumulate(list []string, function func(string) string) []string {
	for index, element := range list {
		list[index] = function(element)
	}
	return list
}
