package services

import (
    "packman-backend/internal/models"
    "log"
)

// GetGlobalLeaderboard retrieves the top scores from the entire game system
func GetGlobalLeaderboard(limit int) ([]models.Score, error) {
    scores, err := models.GetTopScores("", limit)  // "" for global scores
    if err != nil {
        log.Printf("Error retrieving global leaderboard: %v", err)
        return nil, err
    }

    return scores, nil
}

// GetGameLeaderboard retrieves the top scores for a specific game
func GetGameLeaderboard(gameID string, limit int) ([]models.Score, error) {
    scores, err := models.GetTopScores(gameID, limit)
    if err != nil {
        log.Printf("Error retrieving leaderboard for game %s: %v", gameID, err)
        return nil, err
    }

    return scores, nil
}

// StoreScore records a score for a user in a game
func StoreScore(userID int, gameID string, score int) error {
    newScore := models.Score{
        UserID: userID,
        GameID: gameID,
        Score:  score,
    }

    err := newScore.CreateScore()
    if err != nil {
        log.Printf("Error storing score for user %d in game %s: %v", userID, gameID, err)
        return err
    }

    return nil
}
