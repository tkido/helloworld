package main

import (
	"bytes"
	"fmt"
	"strings"

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

	const format = `FPS: %0.2f
mouse: (%d, %d) %v
keys: %s
IsCursorVisible: %v
DeviceScaleFactor: %v
IsFullscreen: %v
IsRunnableInBackground: %v
IsRunningSlowly: %v
IsWindowDecorated: %v
MonitorSize: (%d, %d)
ScreenScale: %0.2f
%s
`
	msg := fmt.Sprintf(format,
		ebiten.CurrentFPS(),
		mx, my, buttons,
		strings.Join(keyStrs, ", "),
		ebiten.IsCursorVisible(),
		ebiten.DeviceScaleFactor(),
		ebiten.IsFullscreen(),
		ebiten.IsRunnableInBackground(),
		ebiten.IsRunningSlowly(),
		ebiten.IsWindowDecorated(),
		sx, sy,
		ebiten.ScreenScale(),
		gamePadInfo(),
	)
	ebitenutil.DebugPrint(screen, msg)

	return
}

func gamePadInfo() string {
	ids := ebiten.GamepadIDs()
	buf := bytes.Buffer{}
	for _, id := range ids {
		buf.WriteString(fmt.Sprintf("gamepads[%d]: ", id))
		axisNum := ebiten.GamepadAxisNum(id)
		for i := 0; i < axisNum; i++ {
			buf.WriteString(fmt.Sprintf(" axis[%d]: %f", i, ebiten.GamepadAxis(id, i)))
		}
		buttonNum := ebiten.GamepadButtonNum(id)
		pressed := []int{}
		for i := 0; i < buttonNum; i++ {
			if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton(i)) {
				pressed = append(pressed, i)
			}
		}
		buf.WriteString(fmt.Sprintf(" %dbuttons: %v", buttonNum, pressed))
		buf.WriteString("\n")
	}
	return buf.String()
}
