package main

//importing the important stuff
import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Create our empty vars
var (
	err                error
	spaceShip          *ebiten.Image
	playerOne          player
	regularTermination = errors.New("regular termination")
	actualNames        [100]string
	fstName            [12]string
	sndName            [12]string
	sideBarImg         *ebiten.Image
	downBarImg         *ebiten.Image
	laattaImg          *ebiten.Image
)

//game struct
type Game struct {
	count  int
	layers [][]int
}

// Create the player struct
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Our game constants
const (
	screenWidth, screenHeight = 1366, 768
	startX, startY            = 512, 512
	sideBarX, sideBarY        = 1088, 0
	downBarX, downBarY        = 0, 576
	nimiX, nimiY              = 100, 300
	tileSize                  = 64
	tileXNum                  = 8
)

// Font var
var (
	mplusNormalFont font.Face
)

//set the fonts
func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

//call once when the program starts
//load images and use the randomizeNames() to get random names to things
func init() {
	randomizeNames()

	sideBarImg, _, err = ebitenutil.NewImageFromFile("assets/sivupalkki.png")
	if err != nil {
		log.Fatal(err)
	}

	laattaImg, _, err = ebitenutil.NewImageFromFile("assets/laatta.png")
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
	playerOne = player{spaceShip, startX, startY, 64}
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

// placeholder consts for the bar info
const level = "Level: 1"
const xpBar = "XP BAR"
const hpBar = "HP bar"

const head = "Helmet"
const body = "Body \narmor"
const rHand = "Weapon"
const lHand = "Shield"
const legs = "Pants"
const boots = "Boots"

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	scale := ebiten.DeviceScaleFactor()

	// Draw the tiles with a for loop
	const xNum = 1100 / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(laattaImg.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	// Print the instructions
	ebitenutil.DebugPrint(screen, "\n Arrows to move, q to quit")

	// Draw player and set its options
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	playerOp.GeoM.Scale(scale, scale)
	screen.DrawImage(playerOne.image, playerOp)

	// Draw the red placeholder bar to divide tilemap from the right bar
	sideBarOp := &ebiten.DrawImageOptions{}
	sideBarOp.GeoM.Translate(sideBarX, sideBarY)
	sideBarOp.GeoM.Scale(scale, scale*1.45)
	screen.DrawImage(sideBarImg, sideBarOp)

	// Draw the red placeholder bar to divide tilemap from the bottom bar
	downBarOp := &ebiten.DrawImageOptions{}
	downBarOp.GeoM.Translate(downBarX, downBarY)
	downBarOp.GeoM.Scale(scale*2.18, scale)
	screen.DrawImage(downBarImg, downBarOp)

	// Draw everything to the right, names are placeholders for now
	text.Draw(screen, actualNames[98], mplusNormalFont, 1100, 80, color.White)
	text.Draw(screen, level, mplusNormalFont, 1100, 100, color.White)
	text.Draw(screen, xpBar, mplusNormalFont, 1100, 120, color.White)
	text.Draw(screen, hpBar, mplusNormalFont, 1100, 140, color.White)

	text.Draw(screen, head, mplusNormalFont, 1185, 240, color.White)
	text.Draw(screen, body, mplusNormalFont, 1190, 280, color.White)
	text.Draw(screen, rHand, mplusNormalFont, 1100, 280, color.White)
	text.Draw(screen, lHand, mplusNormalFont, 1280, 280, color.White)
	text.Draw(screen, legs, mplusNormalFont, 1185, 360, color.White)
	text.Draw(screen, boots, mplusNormalFont, 1185, 400, color.White)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

// Main func, set where and what tiles to draw and run the actual game
func main() {
	game := &Game{
		layers: [][]int{
			{
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243,
			},
			{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
	}
	// Window size will be set to fullscreen
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Taistelun hurmos")
	// starts the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

//generating the names, takes first and last part of the name from lists randomly
//and combines them in a for-loop to an array called actualNames
func randomizeNames() {

	fstName[0] = "Mahti"   // Assign a value to the first element
	fstName[1] = "Tora"    // Assign a value to the first element
	fstName[2] = "Kina"    // Assign a value to the first element
	fstName[3] = "Kuolo"   // Assign a value to the first element
	fstName[4] = "Pyhä"    // Assign a value to the first element
	fstName[5] = "Kauhu"   // Assign a value to the first element
	fstName[6] = "Hilpari" // Assign a value to the first element
	fstName[7] = "Huurre"  // Assign a value to the first element
	fstName[8] = "Pyörre"  // Assign a value to the first element
	fstName[9] = "Karu"    // Assign a value to the first element
	fstName[10] = "Kurja"  // Assign a value to the first element
	fstName[11] = "Pauhu"  // Assign a value to the first element

	sndName[0] = "hammas"     // Assign a value to the first element
	sndName[1] = "käsi"       // Assign a value to the first element
	sndName[2] = "pappa"      // Assign a value to the first element
	sndName[3] = "koira"      // Assign a value to the first element
	sndName[4] = "tonttu"     // Assign a value to the first element
	sndName[5] = "rakki"      // Assign a value to the first element
	sndName[6] = "turjake"    // Assign a value to the first element
	sndName[7] = "huithapeli" // Assign a value to the first element
	sndName[8] = "karju"      // Assign a value to the first element
	sndName[9] = "pentele"    // Assign a value to the first element
	sndName[10] = "valtias"   // Assign a value to the first element
	sndName[11] = "setä"      // Assign a value to the first element

	for i := 0; i < len(actualNames); i++ {
		randomIndex := rand.Intn(len(fstName))
		pick1 := fstName[randomIndex]
		randomIndex2 := rand.Intn(len(sndName))
		pick2 := sndName[randomIndex2]

		actualNames[i] = pick1 + pick2 + "\n"
		fmt.Print(actualNames[i])
	}
}

// Move the player depending on which key is pressed
func movePlayer() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}
