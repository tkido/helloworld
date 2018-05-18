package quadtree

var cellMaxs = [7]int{0, 1, 5, 21, 85, 341, 1365}

func cellNum(upperLeft, lowerRight int) int {
	n := upperLeft ^ lowerRight
	m := msb(n)
	l := m / 2
	if m == -1 {
		l = -1
	}
	shift := (l + 1) * 2
	return lowerRight>>uint(shift) + cellMaxs[4-l]
}

func cellNum2(upperLeft, lowerRight int) int {
	n := upperLeft ^ lowerRight
	switch {
	case n>>8&0x3 != 0:
		return 0
	case n>>6&0x3 != 0:
		return lowerRight>>8 + 1
	case n>>4&0x3 != 0:
		return lowerRight>>6 + 5
	case n>>2&0x3 != 0:
		return lowerRight>>4 + 21
	case n&0x3 != 0:
		return lowerRight>>2 + 85
	default:
		return lowerRight + 341
	}
}

func count(n int) int {
	n = n&0x5555 + n>>1&0x5555
	n = n&0x3333 + n>>2&0x3333
	n = n&0x0f0f + n>>4&0x0f0f
	n = n&0x00ff + n>>8&0x00ff
	return n
}

func msb(n int) int {
	if n == 0 {
		return -1
	}
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	return count(n) - 1
}

func separate(n int) int {
	n = (n | (n << 8)) & 0x00ff00ff
	n = (n | (n << 4)) & 0x0f0f0f0f
	n = (n | (n << 2)) & 0x33333333
	n = (n | (n << 1)) & 0x55555555
	return n
}

func morton(x, y int) int {
	return separate(x) | separate(y)<<1
}
