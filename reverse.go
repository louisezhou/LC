func reverse(x int) int {
    res := 0
	for x != 0 {
		temp := x % 10
		res = res*10 + temp
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}
