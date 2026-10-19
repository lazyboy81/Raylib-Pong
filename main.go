package main

import (
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	WIDTH     = 1200
	HEIGHT    = 580
	BarWidth  = 20
	BarHeight = 100
	BarSpeed  = 200.0
)

var DashLineSize = rl.NewVector2(15, 40)

var DashPositions = []rl.Vector2{
	rl.NewVector2(WIDTH/2-DashLineSize.X, 0),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 60),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 120),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 180),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 240),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 300),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 360),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 420),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 480),
	rl.NewVector2(WIDTH/2-DashLineSize.X, 540),
}

var BarSize = rl.NewVector2(BarWidth, BarHeight)
var LeftBar = rl.NewVector2(20.0, 0.0)
var RightBar = rl.NewVector2(WIDTH-BarWidth-20.0, 0.0)

var BallSize = float32(10.0)
var Ball = rl.NewVector2((WIDTH/2)-(BallSize/2), (HEIGHT/2)-(BallSize/2))
var BallSpeed = rl.NewVector2(250.0, 250.0)

var LeftScore = 0
var RightScore = 0

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "Pong")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rand.Seed(uint64(time.Now().Unix()))
	ballRandomMovement()

	for !rl.WindowShouldClose() {
		checkBoundaries()

		Ball.X = Ball.X + (BallSpeed.X * rl.GetFrameTime())
		Ball.Y = Ball.Y + (BallSpeed.Y * rl.GetFrameTime())

		handleKeyPress()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		drawCentralDashedLines()

		rl.DrawRectangleV(LeftBar, BarSize, rl.White)
		rl.DrawText(strconv.Itoa(LeftScore), 425, 50, 50, rl.White)
		rl.DrawCircleV(Ball, float32(BallSize), rl.White)
		rl.DrawText(strconv.Itoa(RightScore), 725, 50, 50, rl.White)
		rl.DrawRectangleV(RightBar, BarSize, rl.White)

		rl.EndDrawing()
	}
}

func ballRandomMovement() {
	if rl.GetRandomValue(1, 2) == 1 {
		BallSpeed.X = -150
	} else {
		BallSpeed.X = 150
	}

	if rl.GetRandomValue(1, 2) == 1 {
		BallSpeed.Y = float32(rl.GetRandomValue(-100, -150))
	} else {
		BallSpeed.Y = float32(rl.GetRandomValue(100, 150))
	}
}

func handleKeyPress() {
	if rl.IsKeyDown(rl.KeyS) {
		LeftBar.Y = LeftBar.Y + (BarSpeed * rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyW) {
		LeftBar.Y = LeftBar.Y - (BarSpeed * rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyDown) {
		RightBar.Y = RightBar.Y + (BarSpeed * rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyUp) {
		RightBar.Y = RightBar.Y - (BarSpeed * rl.GetFrameTime())
	}
}

func checkBoundaries() {

	// ball and bar collision
	if Ball.X+BallSize >= RightBar.X && (Ball.Y >= RightBar.Y && Ball.Y <= RightBar.Y+BarHeight) {
		BallSpeed.X *= -1.1
		Ball.X = RightBar.X - BallSize - 1
	}

	// right side scoring
	if Ball.X+BallSize >= WIDTH {
		BallSpeed.X *= -1
		Ball.X = (WIDTH / 2) - (BallSize / 2)
		LeftScore += 1
	}

	// ball and bar collision
	if Ball.X-BallSize <= LeftBar.X+BarWidth && (Ball.Y >= LeftBar.Y && Ball.Y <= LeftBar.Y+BarHeight) {
		BallSpeed.X *= -1.1
		Ball.X = LeftBar.X + BarWidth + BallSize + 1
	}

	// left side scoring
	if Ball.X-BallSize <= 0 {
		BallSpeed.X *= -1
		Ball.X = (WIDTH / 2) - (BallSize / 2)
		RightScore += 1
	}

	if Ball.Y-BallSize <= 0 {
		BallSpeed.Y *= -1
	}

	if Ball.Y+BallSize >= HEIGHT {
		BallSpeed.Y *= -1
	}

	if LeftBar.Y <= 0 {
		LeftBar.Y = 0
	}

	if LeftBar.Y+BarHeight >= HEIGHT {
		LeftBar.Y = HEIGHT - BarHeight
	}

	if RightBar.Y <= 0 {
		RightBar.Y = 0
	}

	if RightBar.Y+BarHeight >= HEIGHT {
		RightBar.Y = HEIGHT - BarHeight
	}
}

func drawCentralDashedLines() {
	for _, pos := range DashPositions {
		rl.DrawRectangleV(pos, DashLineSize, rl.White)
	}
}
