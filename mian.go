package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Shooter")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("First Window :)", 190, 200, 20, rl.LightGray)
		rl.DrawRectangle(100 , 100 , 100 , 100, rl.Blue)
		rl.EndDrawing()
	}
}