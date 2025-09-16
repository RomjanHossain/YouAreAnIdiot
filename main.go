package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/color"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed sound.wav Roboto.ttf
var _fs embed.FS

const (
	screenW       = 300
	screenH       = 300
	flashInterval = 300 * time.Millisecond // flash screen (black/white) speed
	audioFilename = "sound.wav"
	fontFileName  = "Roboto.ttf"
)

type Game struct {
	lastFlash   time.Time
	startTime   time.Time
	white       bool
	started     bool
	stopped     bool
	audioPlayer *audio.Player
	fontFace    font.Face
	windowX     float64
	windowY     float64
	vx          float64
	vy          float64
}

// Update is called on every tick of the game.
func (g *Game) Update() error {
	if !g.started {
		g.started = true
		g.startTime = time.Now()
		if g.audioPlayer != nil {
			g.audioPlayer.Play()
		}
	}

	// Toggle flash state each interval
	if g.started && !g.stopped && time.Since(g.lastFlash) >= flashInterval {
		g.white = !g.white
		g.lastFlash = time.Now()
	}
	// if audio stops then play again
	if !g.audioPlayer.IsPlaying() {
		println("audio stopped")
		g.audioPlayer.SetPosition(time.Duration(0))
		g.audioPlayer.Play()
	}

	// Check if the game is running and not stopped
	if g.started && !g.stopped {
		// Update window position based on velocity
		g.windowX += g.vx
		g.windowY += g.vy

		// Get the monitor dimensions to check for collisions
		monitorW, monitorH := ebiten.Monitor().Size()

		// Check for horizontal collision
		if g.windowX <= 0 || g.windowX >= float64(monitorW-screenW) {
			g.vx = -g.vx // Reverse x velocity
		}

		// Check for vertical collision
		if g.windowY <= 0 || g.windowY >= float64(monitorH-screenH) {
			g.vy = -g.vy // Reverse y velocity
		}

		// Update the actual window position
		ebiten.SetWindowPosition(int(g.windowX), int(g.windowY))
	}

	return nil
}

// Draw is called on every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	// Only draw the flashing background if the image isn't supposed to be shown yet
	if g.white {
		screen.Fill(color.Color(color.RGBA{255, 255, 255, 255}))
	} else {
		screen.Fill(color.Color(color.RGBA{0, 0, 0, 255}))
	}
	// a text on the center
	msg := "YOU ARE AN IDIOT"
	text.Draw(screen, msg, g.fontFace, screenW/8, screenH/2, getTextColor(g))
}

// get color
func getTextColor(g *Game) color.Color {
	if g.white {
		return (color.Color(color.RGBA{0, 0, 0, 255}))
	} else {
		return (color.Color(color.RGBA{255, 255, 255, 255}))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--child" {
		// let's embed the assets
		audioData, err := fs.ReadFile(_fs, audioFilename)
		if err != nil {
			log.Fatal("Failed to read embedded audio:", err)
		}

		fontData, err := fs.ReadFile(_fs, fontFileName)
		if err != nil {
			log.Fatal("Failed to read embedded font:", err)
		}
		// random window placement
		rand.Seed(time.Now().UnixNano())
		// get monitor size
		monitorW, monitorH := ebiten.Monitor().Size()
		x := rand.Intn(monitorW - screenW)
		y := rand.Intn(monitorH - screenH)
		// Initial random velocity
		vx := (rand.Float64() - 0.5) * 4 // a random value between -2 and 2
		vy := (rand.Float64() - 0.5) * 4 // a random value between -2 and 2

		// Console prompt for explicit consent
		fmt.Println("=== YouAreAnIdiot (Go) ===")
		fmt.Println("This program will show *rapid flashing* and play sound.")
		fmt.Println("Photosensitive people may be harmed. Do not run this around others.")

		// Initialize Ebiten's audio context
		audioContext := audio.NewContext(44100) // Standard sample rate

		// Load the audio file

		// f, err := os.Open(audioFilename)
		// if err != nil {
		// 	log.Fatal("Failed to open audio file:", err)
		// }
		// defer f.Close()
		// stream, err := wav.DecodeWithSampleRate(44100, f)
		stream, err := wav.DecodeWithSampleRate(44100, bytes.NewReader(audioData))
		if err != nil {
			log.Fatal("Failed to decode WAV file:", err)
		}

		audioPlayer, err := audioContext.NewPlayer(stream)
		if err != nil {
			log.Fatal("Failed to create audio player:", err)
		}
		// load the font
		// fontData, err := os.ReadFile(fontFileName)
		if err != nil {
			log.Fatal("Error reading font file")
		}
		// parse the font
		tt, err := opentype.Parse(fontData)
		if err != nil {
			log.Fatal("Error Parsing fontData")
		}
		// make it font fact
		fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    25,
			DPI:     80,
			Hinting: font.HintingFull,
		})

		game := &Game{
			lastFlash:   time.Now(),
			white:       false,
			audioPlayer: audioPlayer,
			fontFace:    fontFace,
			windowX:     float64(x),
			windowY:     float64(y),
			vx:          vx,
			vy:          vy,
		}

		// set up window
		ebiten.SetWindowDecorated(false)
		ebiten.SetWindowSize(screenW, screenH)
		ebiten.SetWindowPosition(x, y)
		ebiten.SetWindowTitle("YouAreAnIdiot")

		// open window
		if err := ebiten.RunGame(game); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
		fmt.Println("Exiting.")
	} else {
		// This is the fix. Get the path to the current executable.
		exePath, err := os.Executable()
		if err != nil {
			log.Fatal("Couldn't load the exec")
		}
		for i := range 50 {
			cmd := exec.Command(exePath, "--child")
			go func() {
				err := cmd.Start()
				if err != nil {
					fmt.Printf("Error launching window %d: %v\n", i+1, err)
				}
			}()
			time.Sleep(2 * time.Second)

		}
		time.Sleep(5 * time.Second)
		fmt.Println("Launched finished. exiting..")
		os.Exit(0)
	}
}
