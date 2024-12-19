package main

import (
	"database/sql"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"golang.org/x/exp/rand"
)

func resetBoard(cells []*canvas.Rectangle) {
	for _, cell := range cells {
		cell.FillColor = color.White
		cell.Refresh()
	}
}

func checkWin(cells []*canvas.Rectangle, colorToCheck color.Color) ([]int, bool) {
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, pattern := range winPatterns {
		if cells[pattern[0]].FillColor == colorToCheck &&
			cells[pattern[1]].FillColor == colorToCheck &&
			cells[pattern[2]].FillColor == colorToCheck {
			return pattern, true
		}
	}
	return nil, false
}

func isGameOver(cells []*canvas.Rectangle) bool {
	for _, cell := range cells {
		if cell.FillColor == color.White {
			return false
		}
	}
	return true
}

func highlightWinningCells(cells []*canvas.Rectangle, pattern []int) {
	for _, index := range pattern {
		cells[index].FillColor = color.RGBA{0, 255, 0, 255} // Vert pour les cases gagnantes
		cells[index].Refresh()
	}
}

// Fonction modifiée pour trouver le meilleur coup en fonction de la difficulté
func findBestMoveWithDifficulty(board []*canvas.Rectangle) int {
	if difficulty == "Facile" {
		// Coup aléatoire pour le niveau Facile
		rand.Seed(uint64(time.Now().UnixNano()))
		for {
			randomIndex := rand.Intn(9)
			if board[randomIndex].FillColor == color.White {
				return randomIndex
			}
		}
	} else {
		bestScore := -1000
		bestMove := -1
		maxDepth := getMaxDepth()

		for i, cell := range board {
			if cell.FillColor == color.White {
				cell.FillColor = color.RGBA{255, 0, 0, 255} // Ordinateur joue
				score := minimax(board, 0, false, maxDepth)
				cell.FillColor = color.White // Annulation du coup
				if score > bestScore {
					bestScore = score
					bestMove = i
				}
			}
		}
		return bestMove
	}
}

// Fonction pour choisir la profondeur maximale en fonction de la difficulté
func getMaxDepth() int {
	switch difficulty {
	case "Facile":
		return 0 // L'ordinateur joue aléatoirement
	case "Moyen":
		return 2
	case "Difficile":
		return 4
	case "Hardcore":
		return 10 // Analyse complète
	default:
		return 2
	}
}

func handleCellClick(cells []*canvas.Rectangle, index int, isPlayerTurn *bool, db *sql.DB, window fyne.Window) {
	if cells[index].FillColor == color.White && *isPlayerTurn {
		// Joueur joue
		cells[index].FillColor = color.RGBA{0, 0, 255, 255} // Bleu
		cells[index].Refresh()
		playSound("assets/place_piece.mp3")

		if pattern, win := checkWin(cells, color.RGBA{0, 0, 255, 255}); win {
			highlightWinningCells(cells, pattern)
			playSound("assets/win.mp3")
			saveScores(db, 1, 0)
			return
		}

		if isGameOver(cells) {
			playSound("assets/fail.mp3")
			return
		}

		*isPlayerTurn = false

		// IA joue
		go func() {
			time.Sleep(1 * time.Second)

			bestMove := findBestMoveWithDifficulty(cells)
			if bestMove != -1 {
				cells[bestMove].FillColor = color.RGBA{255, 0, 0, 255} // Rouge
				cells[bestMove].Refresh()
				playSound("assets/place_piece.mp3")
			}

			if pattern, win := checkWin(cells, color.RGBA{255, 0, 0, 255}); win {
				highlightWinningCells(cells, pattern)
				playSound("assets/fail.mp3")
				saveScores(db, 0, 1)
				return
			}

			if isGameOver(cells) {
				playSound("assets/fail.mp3")
				return
			}

			*isPlayerTurn = true
		}()
	} else {
		playSound("assets/fail.mp3")
	}
}
