package input

import "github.com/go-gl/glfw"

var (
	mouseX, mouseY int
)

func Init() {
	glfw.Disable(glfw.MouseCursor)
	glfw.SetMousePosCallback(OnMouseMove)
}

func OnMouseMove(x, y int) {
	mouseX, mouseY = x, y
}

func MouseX() float32 {
	return float32(mouseX)
}

func MouseY() float32 {
	return float32(mouseY)
}

func IsKeyDown(key int) bool {
	return glfw.Key(key) == glfw.KeyPress
}
