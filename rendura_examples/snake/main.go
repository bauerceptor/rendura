// A minimal snake (worm) game implementation.
package main

import (
	_ "embed"
	"math/rand"
	"slices"
	"strconv"

	"github.com/bauerceptor/rendura"
	"github.com/bauerceptor/rendura/rendura_cofont"
	"github.com/bauerceptor/rendura/rendura_ebiten"
	"github.com/bauerceptor/rendura/rendura_key"
	"github.com/bauerceptor/rendura/rendura_pad"
)

var snake []rendura.Position           // snake body segments
var fruit rendura.Position             // fruit location
var direction rendura.Position         // current snake heading
var possibleDirection rendura.Position // next possible snake direction

var frame = 0
var speed int
var gameOver = false

const gridSize = 8
const width = 16
const height = 16

var leftDirection = rendura.Position{X: -1}
var rightDirection = rendura.Position{X: 1}
var upDirection = rendura.Position{Y: -1}
var downDirection = rendura.Position{Y: 1}

func startNewGame() {
	gameOver = false

	speed = 5
	direction = rendura.Position{X: 1, Y: 0}
	possibleDirection = direction
	fruit = rendura.Position{X: 8, Y: 8}
	snake = []rendura.Position{
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
	}
}

func handleUserInput() {
	if (rendura_key.Duration(rendura_key.Left) > 0 || rendura_pad.Duration(rendura_pad.Left) > 0) && direction.X == 0 {
		possibleDirection = leftDirection
	}
	if (rendura_key.Duration(rendura_key.Right) > 0 || rendura_pad.Duration(rendura_pad.Right) > 0) && direction.X == 0 {
		possibleDirection = rightDirection
	}
	if (rendura_key.Duration(rendura_key.Up) > 0 || rendura_pad.Duration(rendura_pad.Top) > 0) && direction.Y == 0 {
		possibleDirection = upDirection
	}
	if (rendura_key.Duration(rendura_key.Down) > 0 || rendura_pad.Duration(rendura_pad.Bottom) > 0) && direction.Y == 0 {
		possibleDirection = downDirection
	}
}

func spawnFruit() {
	fruit.X = rand.Intn(width)
	fruit.Y = rand.Intn(height)
}

func update() {
	if gameOver {
		if rendura_key.Duration(rendura_key.Enter) > 0 || rendura_pad.Duration(rendura_pad.A) > 0 {
			startNewGame()
		}
		return
	}

	handleUserInput()

	frame += 1
	if frame%speed == 0 {
		direction = possibleDirection
		// create new head position
		newPos := snake[0].Add(direction)

		// collisions
		// check collision with wall
		if newPos.X < 0 || newPos.X >= width || newPos.Y < 0 || newPos.Y >= height {
			gameOver = true
			return
		}
		// check collision with the snake itself
		for i := 0; i < len(snake); i++ {
			if snake[i] == newPos {
				gameOver = true
				return
			}
		}

		// move the snake body
		snake = slices.Insert(snake, 0, newPos)
		// check if it eats the apple
		if newPos == fruit {
			spawnFruit()
			if len(snake)%10 == 0 && speed > 0 {
				speed -= 1 // increase speed
			}
		} else {
			snake = snake[:len(snake)-1] // remove tail
		}
	}
}

func draw() {
	rendura.Screen().Clear(0)

	drawGrid()
	drawFruit()
	drawSnake()

	if gameOver {
		score := "SCORE: " + strconv.Itoa(len(snake)-3)
		rendura_cofont.Sheet.PrintStroked(score, 54, 58, 7, 5)
		rendura.SetColor(7)
		rendura_cofont.Sheet.Print("HIT ENTER TO START", 33, 74)
	}
}

func drawGrid() {
	rendura.SetColor(1)
	for i := 0; i < width; i++ {
		rendura.Line(i*gridSize, 0, i*gridSize, height*gridSize)
		rendura.Line(0, i*gridSize, width*gridSize, i*gridSize)
	}
}

func drawFruit() {
	verticalShift := frame % 10 / 5 // simple animation
	rendura.DrawSprite(fruitSprite, fruit.X*gridSize, fruit.Y*gridSize+verticalShift)
}

func drawSnake() {
	var headSprite rendura.Sprite
	switch direction {
	case leftDirection:
		headSprite = headHorizontal.WithFlipX(true) // reuse sprite
	case rightDirection:
		headSprite = headHorizontal
	case upDirection:
		headSprite = headVertical
	case downDirection:
		headSprite = headVertical.WithFlipY(true) // reuse sprite
	}
	rendura.DrawSprite(headSprite, snake[0].X*gridSize, snake[0].Y*gridSize)
	for i := 1; i < len(snake); i++ {
		bodySegment := snake[i]
		rendura.DrawSprite(bodySprite, bodySegment.X*gridSize, bodySegment.Y*gridSize)
	}
}

//go:embed "sprites.png"
var spritesPNG []byte

var fruitSprite, headVertical, headHorizontal, bodySprite rendura.Sprite

func main() {
	rendura.Palette = rendura.DecodePalette(spritesPNG)
	sprites := rendura.DecodeCanvas(spritesPNG)
	fruitSprite = rendura.SpriteFrom(sprites, 0, 0, 8, 8)
	headVertical = rendura.SpriteFrom(sprites, 8, 0, 8, 8)
	headHorizontal = rendura.SpriteFrom(sprites, 16, 0, 8, 8)
	bodySprite = rendura.SpriteFrom(sprites, 24, 0, 8, 8)

	rendura.SetTPS(30) // 60 is for hardcore players!
	rendura.SetScreenSize(gridSize*width, gridSize*height)
	rendura.Update = update
	rendura.Draw = draw

	startNewGame()

	rendura_ebiten.Run()
}
