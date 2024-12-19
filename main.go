package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var difficulty = "Facile"

func main() {
	// Initialisation de la base de données
	db := initDatabase()
	defer db.Close()

	// Création de l'application
	myApp := app.New()
	myWindow := myApp.NewWindow("Jeu de Morpion")
	myWindow.Resize(fyne.NewSize(600, 400))

	// Grille de jeu
	grid := container.NewGridWithColumns(3)
	cells := make([]*canvas.Rectangle, 9)
	isPlayerTurn := true

	for i := 0; i < 9; i++ {
		cell := canvas.NewRectangle(color.White)
		cell.SetMinSize(fyne.NewSize(100, 100))
		cells[i] = cell
		cellIndex := i

		tappableCell := container.NewMax(cell, widget.NewButton("", func() {
			handleCellClick(cells, cellIndex, &isPlayerTurn, db, myWindow)
		}))
		grid.Add(tappableCell)
	}

	// Interface utilisateur
	scoreLabel := widget.NewLabel("Score : Joueur 0 - Ordinateur 0")
	playerScore, computerScore := 0, 0

	difficultySelect := widget.NewSelect([]string{"Facile", "Moyen", "Difficile"}, func(selected string) {
		difficulty = selected
	})

	startButton := widget.NewButton("Démarrer", func() {
		playSound("assets/start.mp3")
		resetBoard(cells)
		scoreLabel.SetText(fmt.Sprintf("Score : Joueur %d - Ordinateur %d", playerScore, computerScore))
		isPlayerTurn = true
	})

	showScoresButton := widget.NewButton("Afficher les scores", func() {
		showScoresWindow(db, myWindow)
	})

	content := container.NewVBox(grid, scoreLabel, difficultySelect, startButton, showScoresButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
