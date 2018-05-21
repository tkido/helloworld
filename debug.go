package main

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func debugPrint(screen *ebiten.Image) (err error) {
	mx, my := ebiten.CursorPosition()
	buttons := []string{}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		buttons = append(buttons, "LEFT")
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		buttons = append(buttons, "RIGHT")
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		buttons = append(buttons, "MIDDLE")
	}

	pressed := []ebiten.Key{}
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
	}
	keyStrs := []string{}
	for _, p := range pressed {
		keyStrs = append(keyStrs, p.String())
	}

	sx, sy := ebiten.MonitorSize()

	const format = `CurrentFPS: %0.2f
CursorPosition: (%d, %d)
IsMouseButtonPressed: %v
IsKeyPressed: %v
%sIsCursorVisible: %v
DeviceScaleFactor: %v
IsFullscreen: %v
IsRunnableInBackground: %v
IsRunningSlowly: %v
IsWindowDecorated: %v
MonitorSize: (%d, %d)
ScreenScale: %0.2f
`
	msg := fmt.Sprintf(format,
		ebiten.CurrentFPS(),
		mx, my, buttons,
		keyStrs,
		gamePadInfo(),
		ebiten.IsCursorVisible(),
		ebiten.DeviceScaleFactor(),
		ebiten.IsFullscreen(),
		ebiten.IsRunnableInBackground(),
		ebiten.IsRunningSlowly(),
		ebiten.IsWindowDecorated(),
		sx, sy,
		ebiten.ScreenScale(),
	)
	ebitenutil.DebugPrint(screen, msg)

	return
}

func gamePadInfo() string {
	ids := ebiten.GamepadIDs()
	if len(ids) == 0 {
		return "GamepadIDs: []\n"
	}
	buf := bytes.Buffer{}
	for _, id := range ids {
		buf.WriteString(fmt.Sprintf("GamepadIDs[%d]:\n", id))

		axisNum := ebiten.GamepadAxisNum(id)
		buf.WriteString(fmt.Sprintf(" GamepadAxisNum: %d\n  ", axisNum))
		for i := 0; i < axisNum; i++ {
			buf.WriteString(fmt.Sprintf("GamepadAxis[%d]: %f ", i, ebiten.GamepadAxis(id, i)))
		}
		buf.WriteString("\n")

		buttonNum := ebiten.GamepadButtonNum(id)
		buf.WriteString(fmt.Sprintf(" GamepadButtonNum: %d\n  ", buttonNum))
		pressed := []int{}
		for i := 0; i < buttonNum; i++ {
			if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton(i)) {
				pressed = append(pressed, i)
			}
		}
		buf.WriteString(fmt.Sprintf("IsGamepadButtonPressed: %v", pressed))
		buf.WriteString("\n")
	}
	return buf.String()
}
