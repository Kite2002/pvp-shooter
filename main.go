package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 1000
	SCREEN_HEIGHT = 800
	playerAcc     = 0.2
)

var (
	running      = true
	BG_COLOR     = rl.Black
	background   rl.Texture2D
	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerAngle  = 0
	playerSpeed  = 0.0
)

func degToRad(degrees float32) float32 {
	return degrees * (math.Pi / 180.0)
}

func movePlayer() {
	playerDest.X += float32(playerSpeed) * float32(math.Sin(float64(degToRad(float32(playerAngle)))))
	playerDest.Y += float32(playerSpeed) * float32(-math.Cos(float64(degToRad(float32(playerAngle)))))
	if playerSpeed > 0 {
		playerSpeed -= 0.1
	}
	if playerSpeed < 0 {
		playerSpeed = 0
	}
}

func drawScene() {
	rl.DrawTexture(background, 0, 0, rl.NewColor(255, 255, 255, 130))
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), float32(playerAngle), rl.White)
	movePlayer()
}

func input() {
	println(playerAngle)
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {

		if playerSpeed < 5 {
			playerSpeed += playerAcc
		}
	}
	// if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
	// 	playerDest.X -= playerSpeed * float32(math.Sin(float64(degToRad(float32(playerAngle)))))
	// 	playerDest.Y -= playerSpeed * float32(-math.Cos(float64(degToRad(float32(playerAngle)))))
	// }
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerAngle -= 5
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerAngle += 5
	}
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(BG_COLOR)
	drawScene()
	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
}

func init() {

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Shooter")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	background = rl.LoadTexture("res/bg/bl.png")
	playerSprite = rl.LoadTexture("res/tiny-spaceships/2X/tiny_ship1.png")
	playerSrc = rl.NewRectangle(0, 0, 44, 48)
	playerDest = rl.NewRectangle(100, 100, 50, 50)
}

func quit() {
	rl.UnloadTexture(background)
	rl.UnloadTexture(playerSprite)
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
