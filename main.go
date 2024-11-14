package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	X          float32 // X position of the bullet
	Y          float32 // Y position of the bullet
	Velocity   float32 // Speed of the bullet
	Angle      float32 // Direction of the bullet in radians
	bulletSrc  rl.Rectangle
	bulletDest rl.Rectangle
}

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
	bulletSprite rl.Texture2D

	playerSrc   rl.Rectangle
	playerDest  rl.Rectangle
	playerAngle = 0
	playerSpeed = 0.0
	bullets     []Bullet
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

func updateBullets() {
	for i := range bullets { // Use index to modify the original slice
		bullets[i].bulletDest.X += float32(bullets[i].Velocity) * float32(math.Sin(float64(degToRad(float32(bullets[i].Angle)))))
		bullets[i].bulletDest.Y += float32(bullets[i].Velocity) * float32(-math.Cos(float64(degToRad(float32(bullets[i].Angle)))))
	}
}

func drawScene() {
	rl.DrawTexture(background, 0, 0, rl.NewColor(255, 255, 255, 130))
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), float32(playerAngle), rl.White)
	updateBullets()
	for _, bullet := range bullets {
		rl.DrawTexturePro(bulletSprite, bullet.bulletSrc, bullet.bulletDest, rl.NewVector2(bullet.bulletDest.Width/2, bullet.bulletDest.Height/2), float32(bullet.Angle-90), rl.White)
		// rl.Draw(int32(bullet.X), int32(bullet.Y), 5, rl.Blue)
	}
	movePlayer()
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {

		if playerSpeed < 5 {
			playerSpeed += playerAcc
		}
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerAngle -= 5
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerAngle += 5
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		bullets = append(bullets, Bullet{
			X:          float32(playerDest.X),
			Y:          float32(playerDest.Y),
			Velocity:   12,
			Angle:      float32(playerAngle),
			bulletSrc:  rl.NewRectangle(0, 48, 24, 24),
			bulletDest: rl.NewRectangle(float32(playerDest.X), float32(playerDest.Y), 24, 24),
		})
		fmt.Printf("bullets: %v\n", bullets)
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
	bulletSprite = rl.LoadTexture("res/projectiles/bullets.png")
	playerSrc = rl.NewRectangle(0, 0, 44, 48)
	playerDest = rl.NewRectangle(100, 100, 50, 50)
	bullets = []Bullet{}
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
