func solution(w int, h int) int64 {
	var answer int64 = 0
	if h > w {
		temp := h
		h = w
		w = temp
	}
	var a, b int
	a = w
	b = h
	for {
		if b == 0 {
			break
		}
		r := a % b
		a = b
		b = r
	}
	w = int(w / a)
	h = int(h / a)

	var k float64
	k = float64(h) / float64(w)
	var i, j int
	var cnt int64
	for {
		if j >= h {
			break
		}
		cnt++
		if k*float64(i+1) > float64(j+1) {
			j = j + 1
		} else if k*float64(i+1) == float64(int(k*float64(i+1))) {
			i++
			j++
		} else {
			i++
		}
	}
	cnt = cnt * int64(a)
	answer = (int64(w*h)*int64(a*a) - cnt)
	return answer
}