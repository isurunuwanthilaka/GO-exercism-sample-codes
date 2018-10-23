package summultiples

//SumMultiples for total of multiples
func SumMultiples(limit int, devisors ...int) int {
	var tot int
	if len(devisors) == 0 {
		return tot
	}
	for i := 1; i < limit; i++ {
		for _, j := range devisors {
			if i%j == 0 {
				tot += i
				break
			}
		}
	}
	return tot
}
