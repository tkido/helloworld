package quadtree

func count(n int) int {
	n = n&0x5555 + n>>1&0x5555
	n = n&0x3333 + n>>2&0x3333
	n = n&0x0f0f + n>>4&0x0f0f
	n = n&0x00ff + n>>8&0x00ff
	return n
}

func msb(n int) int {
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
