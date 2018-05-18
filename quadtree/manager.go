package quadtree

// Collisioner is collisioner
type Collisioner interface {
	Check(Collisioner)
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
	stack := make([]Collisioner, 0, 128)
	return &Manager{w, h, cells, stack}
}

func (m *Manager) update(c Collisioner, x1, y1, x2, y2 float64) {
	c0 := c.GetCellNum()
	tl, br := morton(int(x1/32), int(y1/32)), morton(int(x2/32), int(y2/32))
	c1 := cellNum(tl, br)
	if c0 != c1 {
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
}
