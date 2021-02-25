package main

//importing the important stuff
import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Function for updating the screen
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
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

		actualNames[i] = pick1 + pick2
		fmt.Print(actualNames[i])
	}
}

//the main function
func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
