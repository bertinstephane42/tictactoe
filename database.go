package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Initialise la base de données SQLite
func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./scores.db")
	if err != nil {
		panic(err)
	}

	// Création de la table pour les scores
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			player_score INTEGER,
			computer_score INTEGER,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	return db
}

// Enregistre les scores dans la base de données
func saveScores(db *sql.DB, playerScore, computerScore int) {
	_, err := db.Exec("INSERT INTO scores (player_score, computer_score) VALUES (?, ?)", playerScore, computerScore)
	if err != nil {
		fmt.Println("Erreur lors de l'enregistrement des scores :", err)
	}
}

// Récupère les scores depuis la base de données
func fetchScores(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT player_score, computer_score, date FROM scores ORDER BY date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []string
	for rows.Next() {
		var playerScore, computerScore int
		var date string
		err := rows.Scan(&playerScore, &computerScore, &date)
		if err != nil {
			return nil, err
		}
		scores = append(scores, fmt.Sprintf("Date: %s | Joueur: %d | Ordinateur: %d", date, playerScore, computerScore))
	}
	return scores, nil
}
