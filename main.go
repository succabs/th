package main

//importing the important stuff
import (
	"errors"
	"fmt"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Create our empty vars
var (
	err                error
	spaceShip          *ebiten.Image
	playerOne          player
	regularTermination = errors.New("regular termination")
	actualNames        [10]string
	fstName            [5]string
	sndName            [5]string
	sideBarImg         *ebiten.Image
	downBarImg         *ebiten.Image
)

//game struct
type Game struct {
	count int
}

// Create the player struct
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Our game constants
const (
	screenWidth, screenHeight = 640, 480
	startX, startY            = 100, 100
	sideBarX, sideBarY        = 1100, 0
	downBarX, downBarY        = 0, 600
)

//call once when the program starts
func init() {

	sideBarImg, _, err = ebitenutil.NewImageFromFile("assets/sivupalkki.png")
	if err != nil {
		log.Fatal(err)
	}

	downBarImg, _, err = ebitenutil.NewImageFromFile("assets/alapalkki.png")
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
	playerOne = player{spaceShip, startX, startY, 4}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.count++

	movePlayer()

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return regularTermination
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	scale := ebiten.DeviceScaleFactor()

	ebitenutil.DebugPrint(screen, "Arrows to move, q to quit")
	randomizeNames()

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	playerOp.GeoM.Scale(scale, scale)
	screen.DrawImage(playerOne.image, playerOp)

	sideBarOp := &ebiten.DrawImageOptions{}
	sideBarOp.GeoM.Translate(sideBarX, sideBarY)
	sideBarOp.GeoM.Scale(scale, scale*1.5)
	screen.DrawImage(sideBarImg, sideBarOp)

	downBarOp := &ebiten.DrawImageOptions{}
	downBarOp.GeoM.Translate(downBarX, downBarY)
	downBarOp.GeoM.Scale(scale*2.2, scale)
	screen.DrawImage(downBarImg, downBarOp)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Taistelun hurmos")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

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
