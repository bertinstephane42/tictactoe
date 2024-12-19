package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showScoresWindow(db *sql.DB, parent fyne.Window) {
	scores, err := fetchScores(db)
	if err != nil {
		widget.NewLabel("Erreur lors du chargement des scores")
		return
	}

	scoreList := widget.NewList(
		func() int { return len(scores) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Score")
		},
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(scores[i])
		},
	)

	scoresWindow := fyne.CurrentApp().NewWindow("Scores enregistrés")
	scoresWindow.SetContent(container.NewVBox(
		widget.NewLabel("Scores enregistrés :"),
		scoreList,
	))
	scoresWindow.Resize(fyne.NewSize(400, 300))
	scoresWindow.Show()
}
