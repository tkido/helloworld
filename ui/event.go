package ui

import "log"

// EventType is type of all UI event
type EventType int

// MouseEvents
const (
	LeftDown EventType = iota
	RightDown
	MiddleDown
	LeftUp
	RightUp
	MiddleUp
	LeftClick
	RightClick
	MiddleClick
	LeftDoubleClick
	RightDoubleClick
	MiddleDoubleClick
	MouseOn
	MouseOver
	MouseOut
	MouseIn
	MouseEnter
	MouseLeave
)

// String() for fmt.Stringer interface
func (et EventType) String() string {
	switch et {
	case LeftDown:
		return "LeftDown"
	case RightDown:
		return "RightDown"
	case MiddleDown:
		return "MiddleDown"
	case LeftUp:
		return "LeftUp"
	case RightUp:
		return "RightUp"
	case MiddleUp:
		return "MiddleUp"
	case LeftClick:
		return "LeftClick"
	case RightClick:
		return "RightClick"
	case MiddleClick:
		return "MiddleClick"
	case LeftDoubleClick:
		return "LeftDoubleClick"
	case RightDoubleClick:
		return "RightDoubleClick"
	case MiddleDoubleClick:
		return "MiddleDoubleClick"
	case MouseOn:
		return "MouseOn"
	case MouseOver:
		return "MouseOver"
	case MouseOut:
		return "MouseOut"
	case MouseIn:
		return "MouseIn"
	case MouseEnter:
		return "MouseEnter"
	case MouseLeave:
		return "MouseLeave"
	default:
		log.Panicf("unknown EventType %d", et)
		return ""
	}
}
