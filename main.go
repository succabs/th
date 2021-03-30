package main

//importing the important stuff
import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

// Our game constants
const (
	screenWidth, screenHeight = 640, 480
)

//call once when the program starts
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{spaceShip, screenWidth / 2, screenHeight / 2, 4}
}

//Function for updating the screen
func update(screen *ebiten.Image) error {
	movePlayer()
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
	ebitenutil.DebugPrint(screen, "Tervetuloa!")

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

// Move the player depending on which key is pressed
func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

//arrays for character names
var actualNames [10]string
var fstName [5]string
var sndName [5]string

//generating the names, takes first and last part of the name from lists randomly
//and combines them in a for-loop to an array called actualNames
func randomizeNames() {

	fstName[0] = "Mahti" // Assign a value to the first element
	fstName[1] = "Tora"  // Assign a value to the first element
	fstName[2] = "Kina"  // Assign a value to the first element
	fstName[3] = "Kuolo" // Assign a value to the first element
	fstName[4] = "Pyhä"  // Assign a value to the first element

	sndName[0] = "hammas" // Assign a value to the first element
	sndName[1] = "käsi"   // Assign a value to the first element
	sndName[2] = "pappa"  // Assign a value to the first element
	sndName[3] = "koira"  // Assign a value to the first element
	sndName[4] = "tonttu" // Assign a value to the first element

	for i := 0; i < len(actualNames); i++ {
		randomIndex := rand.Intn(len(fstName))
		pick1 := fstName[randomIndex]
		randomIndex2 := rand.Intn(len(sndName))
		pick2 := sndName[randomIndex2]

		actualNames[i] = pick1 + pick2 + " \n"
		fmt.Print(actualNames[i])
	}
}

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

//the main function
func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
