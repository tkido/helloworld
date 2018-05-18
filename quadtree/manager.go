package quadtree

// Collisioner is collisioner
type Collisioner interface {
	Check(Collisioner) bool
	GetCellNum() int
	SetCellNum(int)
}

// Manager is collision manager
type Manager struct {
	Width, Height float64
	Cells         []map[Collisioner]struct{}
	Stack         []Collisioner
}

// NewManager is
func NewManager(w, h float64) *Manager {
	cells := make([]map[Collisioner]struct{}, 1365)
	for i := range cells {
		cells[i] = make(map[Collisioner]struct{})
	}
	stack := make([]Collisioner, 0, 128)
	return &Manager{w, h, cells, stack}
}

func sanitize(f, max float64) int {
	switch {
	case f < 0:
		return 0
	case f >= max:
		return 31
	default:
		return int(f * 32 / max)
	}
}

// Update updates collisioner
func (m *Manager) Update(c Collisioner, x1, y1, x2, y2 float64) {
	c0 := c.GetCellNum()
	tl, br := morton(sanitize(x1, m.Width), sanitize(y1, m.Width)), morton(sanitize(x2, m.Width), sanitize(y2, m.Width))
	c1 := cellNum(tl, br)
	if c0 != c1 {
		// fmt.Printf("ボールのセルを%dから%dへ移動する。\n", c0, c1)
		if c0 != -1 {
			delete(m.Cells[c0], c)
		}
		m.Cells[c1][c] = struct{}{}
		c.SetCellNum(c1)
	}
}

// Check check all collisioners
func (m *Manager) Check(i int) {
	list := m.Cells[i]
	// fmt.Printf("%d番セルのチェック\n", i)
	// fmt.Printf("このセルには%d個のボールが所属する\n", len(list))
	// fmt.Printf("現在のスタックサイズは%dである\n", len(m.Stack))
	for c := range list {
		for o := range list {
			c.Check(o)
		}
		for _, o := range m.Stack {
			c.Check(o)
			o.Check(c)
		}
		m.Stack = append(m.Stack, c)
	}
	base := i*4 + 1
	if base < 1365 {
		for j := base; j < base+4; j++ {
			m.Check(j)
		}
	}
	m.Stack = m.Stack[:len(m.Stack)-len(list)]
	// fmt.Printf("%d番セルのチェックを終了する\n", i)
	// fmt.Printf("スタックから%d個取り除いた\n", len(list))
	// fmt.Printf("スタックサイズは%dになった\n", len(m.Stack))

}
