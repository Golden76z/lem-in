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
	ant      int
	path     int
	step     int
	progress float64
}

func NewVisualizer(roomStruct *functions.RoomStruct, paths [][]string, antDistribution [][]int) *Visualizer {
	occupiedRooms := make(map[string]bool)
	// Initialiser toutes les salles comme non-occupées sauf la salle de départ
	for _, room := range roomStruct.AllRooms {
		occupiedRooms[room.Name] = false
	}
	// La salle de départ peut contenir plusieurs fourmis
	occupiedRooms[roomStruct.StartingRoom.Name] = true // Toujours libre pour de nouvelles fourmis

	return &Visualizer{
		roomStruct:      roomStruct,
		paths:           paths,
		antDistribution: antDistribution,
		paused:          true,
		moveDelay:       time.Second,
	}
}

func (v *Visualizer) initAntPositions() {
	for pathIndex, ants := range v.antDistribution {
		for _, ant := range ants {
			v.antPositions = append(v.antPositions, AntPosition{ant, pathIndex, 0, moveSpeed})
		}
	}
}

const (
	moveInterval = 800 * time.Millisecond // Intervalle de temps entre les mouvements des fourmis
)

const (
	moveSpeed = 0.01 // Vitesse de déplacement des fourmis
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
		// Mettre à jour la progression des fourmis
		v.updateAntsProgress()

		// Si l'intervalle de temps est atteint, passer à la prochaine étape
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

func (v *Visualizer) updateAntsProgress() {
	// Mettre à jour la progression des fourmis entre les salles
	for i := range v.antPositions {
		// Si la fourmi est toujours en déplacement (avant d'atteindre la dernière étape)
		if v.antPositions[i].step < len(v.paths[v.antPositions[i].path])-1 {
			v.antPositions[i].progress += moveSpeed
			// Limiter la progression à 1.0
			if v.antPositions[i].progress > 1.0 {
				v.antPositions[i].progress = 1.0
			}
		}
	}
}

func (v *Visualizer) drawAnts(screen *ebiten.Image) {
	for _, pos := range v.antPositions {
		if pos.step < len(v.paths[pos.path])-1 {
			currentRoom := v.paths[pos.path][pos.step]
			nextRoom := v.paths[pos.path][pos.step+1]

			var currentX, currentY, nextX, nextY int

			// Obtenir les coordonnées des salles actuelle et suivante
			for _, room := range v.roomStruct.AllRooms {
				if room.Name == currentRoom {
					currentX, currentY = v.mapCoordinates(room.X_value, room.Y_value)
				}
				if room.Name == nextRoom {
					nextX, nextY = v.mapCoordinates(room.X_value, room.Y_value)
				}
			}

			// Interpoler les positions en fonction de la progression
			interpolatedX := int(float64(currentX)*(1.0-pos.progress) + float64(nextX)*pos.progress)
			interpolatedY := int(float64(currentY)*(1.0-pos.progress) + float64(nextY)*pos.progress)

			// Dessiner un cercle pour représenter la fourmi à la position interpolée
			ebitenutil.DrawCircle(screen, float64(interpolatedX), float64(interpolatedY), 5, color.RGBA{255, 0, 0, 255})
			// Afficher l'identifiant de la fourmi
			text.Draw(screen, fmt.Sprintf("L%d", pos.ant), gameFont, interpolatedX-5, interpolatedY-10, color.White)
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
		if pos.step < len(v.paths[pos.path])-1 {
			currentRoom := v.paths[pos.path][pos.step]
			nextRoom := v.paths[pos.path][pos.step+1]

			// Si la fourmi est dans la salle de départ, elle ne peut avancer que si la première salle de son chemin est libre
			if currentRoom == v.roomStruct.StartingRoom.Name {
				// Vérifie si la salle suivante (la première du chemin) est libre
				if !v.occupiedRooms[nextRoom] {
					// Si la salle suivante est libre, la fourmi peut commencer à avancer
					pos.step++
					pos.progress = 0.0
					v.occupiedRooms[nextRoom] = true // Occupe la première salle du chemin
				}
				// Sinon, la fourmi reste dans la salle de départ
				newPositions = append(newPositions, pos)
				continue
			}

			// Si la fourmi est déjà en déplacement (et pas dans la salle de départ)
			isEndRoom := nextRoom == v.roomStruct.EndingRoom.Name
			if pos.progress >= 1.0 {
				// Si la salle suivante est libre ou que c'est la salle de fin
				if isEndRoom || !v.occupiedRooms[nextRoom] {
					// La fourmi avance d'une salle
					pos.step++
					pos.progress = 0.0
					v.occupiedRooms[currentRoom] = false // Libère la salle actuelle
					if !isEndRoom {
						v.occupiedRooms[nextRoom] = true // Occupe la salle suivante
					}
				}
			}
			newPositions = append(newPositions, pos)
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
