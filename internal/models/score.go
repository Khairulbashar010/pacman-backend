package models

import (
	"database/sql"
	"packman-backend/internal/db"
	"time"
)

type Score struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    GameID    string    `json:"game_id"`
    Score     int       `json:"score"`
    CreatedAt time.Time `json:"created_at"`
}

// CreateScore inserts a new score into the database
func (s *Score) CreateScore() error {
    query := "INSERT INTO scores (user_id, game_id, score, created_at) VALUES (?, ?, ?, ?)"
    _, err := db.DB.Exec(query, s.UserID, s.GameID, s.Score, time.Now())
    if err != nil {
        return err
    }
    return nil
}

// GetTopScores retrieves the top scores for a game or globally if gameID is empty
func GetTopScores(gameID string, limit int) ([]Score, error) {
    var rows *sql.Rows
    var err error

    if gameID == "" {
        // Global leaderboard: No specific gameID
        query := "SELECT id, user_id, game_id, score, created_at FROM scores ORDER BY score DESC LIMIT ?"
        rows, err = db.DB.Query(query, limit)
    } else {
        // Leaderboard for a specific game
        query := "SELECT id, user_id, game_id, score, created_at FROM scores WHERE game_id = ? ORDER BY score DESC LIMIT ?"
        rows, err = db.DB.Query(query, gameID, limit)
    }

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var scores []Score
    for rows.Next() {
        var score Score
        if err := rows.Scan(&score.ID, &score.UserID, &score.GameID, &score.Score, &score.CreatedAt); err != nil {
            return nil, err
        }
        scores = append(scores, score)
    }

    return scores, nil
}
