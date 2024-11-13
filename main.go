package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  = 1000
	SCREEN_HEIGHT = 800
)

var (
	running  = true
	BG_COLOR = rl.Black
)

func drawScene() {
}

func input() {

}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(BG_COLOR)

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
}

func init() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Shooter")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func quit() {
	rl.CloseWindow()
}
func main() {
	for running {
		render()
		input()
		update()
	}
	quit()
}
