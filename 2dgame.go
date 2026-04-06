package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth = 800
	screenHeight = 300
	fontSize = 36
)

func main() {
  rl.InitWindow(800, 450, "raylib [core] example - basic window")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()

    rl.ClearBackground(rl.White)
    rl.DrawText("Congrats! You created your first window!", int32(screenWidth), int32(screenHeight), int32(fontSize), rl.LightGray)


    rl.EndDrawing()

  }
}
