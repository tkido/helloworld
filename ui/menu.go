package ui

// Menu is menu
type Menu struct {
	Box
	Choices
	Choiced
}

type Choice interface{}
type Choices []Choice
type Choiced func(c Choice)

func NewMenu(w, h int, cs Choices, cd Choiced) *Menu {
	b := NewBox(w, h, nil)
	m := &Menu{*b, cs, cd}
	m.Sub = m

	return m
}
