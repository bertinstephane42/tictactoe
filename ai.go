package main

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func minimax(board []*canvas.Rectangle, depth int, isMaximizing bool, maxDepth int) int {
	if depth >= maxDepth {
		return 0
	}
	if isWinningState(board, color.RGBA{255, 0, 0, 255}) { // Rouge (IA)
		return 10 - depth
	}
	if isWinningState(board, color.RGBA{0, 0, 255, 255}) { // Bleu (joueur)
		return depth - 10
	}
	if isGameOver(board) {
		return 0 // Match nul
	}

	if isMaximizing {
		bestScore := -1000
		for _, cell := range board {
			if cell.FillColor == color.White {
				cell.FillColor = color.RGBA{255, 0, 0, 255}
				score := minimax(board, depth+1, false, maxDepth)
				cell.FillColor = color.White
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for _, cell := range board {
			if cell.FillColor == color.White {
				cell.FillColor = color.RGBA{0, 0, 255, 255}
				score := minimax(board, depth+1, true, maxDepth)
				cell.FillColor = color.White
				if score < bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
}

func isWinningState(board []*canvas.Rectangle, colorToCheck color.Color) bool {
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, pattern := range winPatterns {
		if board[pattern[0]].FillColor == colorToCheck &&
			board[pattern[1]].FillColor == colorToCheck &&
			board[pattern[2]].FillColor == colorToCheck {
			return true
		}
	}
	return false
}
