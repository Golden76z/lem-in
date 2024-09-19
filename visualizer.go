package main

import (
	"fmt"
	"image/color"
	"lemin/functions"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 800
	screenHeight = 600
	fontSize     = 12
)

var (
	gameFont font.Face
)

type Visualizer struct {
	moveDelay       time.Duration
	roomStruct      *functions.RoomStruct
	paths           [][]string
	antDistribution [][]int
	antPositions    []AntPosition
	step            int
	paused          bool
	lastMoveTime    time.Time
	occupiedRooms   map[string]bool
}

type AntPosition struct {
	ant  int
	path int
	step int
}

func NewVisualizer(roomStruct *functions.RoomStruct, paths [][]string, antDistribution [][]int) *Visualizer {
	occupiedRooms := make(map[string]bool)
	// Initialiser toutes les salles comme non-occupées sauf la salle de départ
	for _, room := range roomStruct.AllRooms {
		occupiedRooms[room.Name] = false
	}
	// La salle de départ peut contenir plusieurs fourmis
	occupiedRooms[roomStruct.StartingRoom.Name] = false // Toujours libre pour de nouvelles fourmis

	return &Visualizer{
		roomStruct:      roomStruct,
		paths:           paths,
		antDistribution: antDistribution,
		paused:          false,
		moveDelay:       time.Second,
	}
}

func (v *Visualizer) initAntPositions() {
	for pathIndex, ants := range v.antDistribution {
		for _, ant := range ants {
			v.antPositions = append(v.antPositions, AntPosition{ant, pathIndex, 0})
		}
	}
}

const (
	moveInterval = 500 * time.Millisecond // Intervalle de temps entre les mouvements des fourmis
)

func (v *Visualizer) Update() error {
	// Permet de mettre en pause ou de reprendre avec la touche espace
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		v.paused = !v.paused
	}

	// Ajuste la vitesse avec les touches fléchées UP et DOWN
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		v.moveDelay -= 50 * time.Millisecond
		if v.moveDelay < 50*time.Millisecond {
			v.moveDelay = 50 * time.Millisecond
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		v.moveDelay += 50 * time.Millisecond
		if v.moveDelay > 5*time.Second {
			v.moveDelay = 5 * time.Second
		}
	}

	// Déplace les fourmis automatiquement si le jeu n'est pas en pause
	if !v.paused {
		if time.Since(v.lastMoveTime) >= v.moveDelay {
			v.step++
			v.moveAnts()
			v.lastMoveTime = time.Now()
		}
	}

	return nil
}

func (v *Visualizer) Draw(screen *ebiten.Image) {
	v.drawRooms(screen)
	v.drawTunnels(screen)
	v.drawAnts(screen)
	v.drawInfo(screen)
}

func (v *Visualizer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (v *Visualizer) drawRooms(screen *ebiten.Image) {
	for _, room := range v.roomStruct.AllRooms {
		x, y := v.mapCoordinates(room.X_value, room.Y_value)
		ebitenutil.DrawRect(screen, float64(x-5), float64(y-5), 10, 10, color.White)
		text.Draw(screen, room.Name, gameFont, x+10, y+5, color.White)
	}

	startRoom := v.roomStruct.StartingRoom
	endRoom := v.roomStruct.EndingRoom
	startX, startY := v.mapCoordinates(startRoom.X_value, startRoom.Y_value)
	endX, endY := v.mapCoordinates(endRoom.X_value, endRoom.Y_value)

	ebitenutil.DrawRect(screen, float64(startX-7), float64(startY-7), 14, 14, color.RGBA{0, 255, 0, 255})
	ebitenutil.DrawRect(screen, float64(endX-7), float64(endY-7), 14, 14, color.RGBA{255, 0, 0, 255})
}

func (v *Visualizer) drawTunnels(screen *ebiten.Image) {
	for _, room := range v.roomStruct.AllRooms {
		x1, y1 := v.mapCoordinates(room.X_value, room.Y_value)
		for _, link := range room.Links {
			for _, linkedRoom := range v.roomStruct.AllRooms {
				if linkedRoom.Name == link {
					x2, y2 := v.mapCoordinates(linkedRoom.X_value, linkedRoom.Y_value)
					ebitenutil.DrawLine(screen, float64(x1), float64(y1), float64(x2), float64(y2), color.RGBA{100, 100, 100, 255})
					break
				}
			}
		}
	}
}

func (v *Visualizer) drawAnts(screen *ebiten.Image) {
	for _, pos := range v.antPositions {
		if pos.step < len(v.paths[pos.path]) {
			roomName := v.paths[pos.path][pos.step]
			for _, room := range v.roomStruct.AllRooms {
				if room.Name == roomName {
					x, y := v.mapCoordinates(room.X_value, room.Y_value)
					// Dessiner un cercle pour représenter la fourmi
					ebitenutil.DrawCircle(screen, float64(x), float64(y), 5, color.RGBA{255, 0, 0, 255})
					// Afficher l'identifiant de la fourmi
					text.Draw(screen, fmt.Sprintf("L%d", pos.ant), gameFont, x-5, y-10, color.White)
					break
				}
			}
		}
	}
}

func (v *Visualizer) drawInfo(screen *ebiten.Image) {
	info := fmt.Sprintf("Step: %d\nAnts: %d\nPaths: %d\nPress SPACE to pause/resume\nMove delay: %.1f seconds\nUse UP/DOWN arrows to adjust speed",
		v.step, len(v.antPositions), len(v.paths), v.moveDelay.Seconds())
	text.Draw(screen, info, gameFont, 10, screenHeight-100, color.White)
}

const (
	antMoveInterval = 2 // Nombre d'étapes avant qu'une fourmi se déplace
)

func (v *Visualizer) moveAnts() {
	var newPositions []AntPosition

	// Initialiser la map occupiedRooms si nécessaire
	if v.occupiedRooms == nil {
		v.occupiedRooms = make(map[string]bool)
	}

	// Boucle sur toutes les fourmis
	for _, pos := range v.antPositions {
		// Vérifie que la fourmi n'a pas atteint la dernière étape de son chemin
		if pos.step < len(v.paths[pos.path])-1 {
			currentRoom := v.paths[pos.path][pos.step]
			nextRoom := v.paths[pos.path][pos.step+1]

			// Vérifie si la salle "nextRoom" est la salle "end"
			isEndRoom := nextRoom == v.roomStruct.EndingRoom.Name

			// Si la salle suivante est la salle "end" ou n'est pas occupée, la fourmi peut avancer
			if isEndRoom || !v.occupiedRooms[nextRoom] {
				// Déplace la fourmi
				v.occupiedRooms[currentRoom] = false // Libère la salle actuelle
				if !isEndRoom {                      // Si ce n'est pas la salle "end", occupe la prochaine salle
					v.occupiedRooms[nextRoom] = true
				}
				// Ajoute la nouvelle position de la fourmi (incrémente step)
				newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
			} else {
				// Si la salle suivante est occupée, la fourmi reste sur place
				newPositions = append(newPositions, pos)
			}
		} else {
			// Si la fourmi est déjà à la dernière étape, elle ne bouge pas
			newPositions = append(newPositions, pos)
		}
	}

	// Met à jour les positions des fourmis
	v.antPositions = newPositions
}

func (v *Visualizer) mapCoordinates(x, y int) (int, int) {
	// Ajuste la mise à l'échelle des coordonnées si nécessaire
	return x*50 + 50, y*50 + 50
}

func initFont() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	gameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func RunVisualizer(roomStruct *functions.RoomStruct, paths [][]string, antDistribution [][]int) {
	initFont()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Lem-in Visualizer")
	v := NewVisualizer(roomStruct, paths, antDistribution)
	v.initAntPositions() // Initialize the positions of the ants
	if err := ebiten.RunGame(v); err != nil {
		log.Fatal(err)
	}
}
