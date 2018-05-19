package quadtree

// Collisioner is collisioner
type Collisioner interface {
	Check(Collisioner) bool
	GetCellNum() int
	SetCellNum(int)
}

// Manager is collision manager
type Manager struct {
	w, h    float64
	level   int
	offsets []int
	cells   []map[Collisioner]struct{}
	stack   []Collisioner
	length  int
}

// NewManager is
func NewManager(w, h float64, level int) *Manager {
	offsets := []int{0}
	for i := 0; i <= level; i++ {
		offsets = append(offsets, offsets[i]+1<<(uint(2*i)))
	}
	cells := make([]map[Collisioner]struct{}, offsets[level+1])
	for i := range cells {
		cells[i] = make(map[Collisioner]struct{})
	}
	stack := make([]Collisioner, 0, 128)
	length := 1 << uint(level)
	return &Manager{w, h, level, offsets, cells, stack, length}
}

func (m *Manager) cellNum(topLeft, bottomRight int) int {
	n := (msb(topLeft^bottomRight) + 2) / 2
	return bottomRight>>uint(n*2) + m.offsets[m.level-n]
}

func (m *Manager) sanitize(f, max float64) int {
	switch {
	case f < 0:
		return 0
	case f >= max:
		return m.length - 1
	default:
		return int(f * float64(m.length) / max)
	}
}

// Update updates collisioner
func (m *Manager) Update(c Collisioner, x1, y1, x2, y2 float64) {
	c0 := c.GetCellNum()
	tl, br := morton(m.sanitize(x1, m.w), m.sanitize(y1, m.w)), morton(m.sanitize(x2, m.w), m.sanitize(y2, m.w))
	c1 := m.cellNum(tl, br)
	if c0 != c1 {
		if c0 != -1 {
			delete(m.cells[c0], c)
		}
		m.cells[c1][c] = struct{}{}
		c.SetCellNum(c1)
	}
}

// Check check all collisioners
func (m *Manager) Check() {
	m.check(0)
}

func (m *Manager) check(i int) {
	list := m.cells[i]
	// fmt.Printf("%d番セルのチェック\n", i)
	// fmt.Printf("このセルには%d個のボールが所属する\n", len(list))
	// fmt.Printf("現在のスタックサイズは%dである\n", len(m.Stack))
	for c := range list {
		for o := range list {
			c.Check(o)
		}
		for _, o := range m.stack {
			c.Check(o)
			o.Check(c)
		}
		m.stack = append(m.stack, c)
	}
	base := i*4 + 1
	if base < m.offsets[m.level+1] {
		for j := base; j < base+4; j++ {
			m.check(j)
		}
	}
	m.stack = m.stack[:len(m.stack)-len(list)]
	// fmt.Printf("%d番セルのチェックを終了する\n", i)
	// fmt.Printf("スタックから%d個取り除いた\n", len(list))
	// fmt.Printf("スタックサイズは%dになった\n", len(m.Stack))

}
