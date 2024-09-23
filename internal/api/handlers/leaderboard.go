package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "packman-backend/internal/services"
    "github.com/gorilla/mux"
)

// GetGlobalLeaderboardHandler handles requests for the global leaderboard
func GetGlobalLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
    limitParam := r.URL.Query().Get("limit")
    limit, _ := strconv.Atoi(limitParam)
    if limit == 0 {
        limit = 10 // Default limit
    }

    scores, err := services.GetGlobalLeaderboard(limit)
    if err != nil {
        http.Error(w, "Failed to retrieve global leaderboard", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "leaderboard": scores,
    })
}

// GetGameLeaderboardHandler handles requests for a game's leaderboard
func GetGameLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gameID := vars["gameId"]

    limitParam := r.URL.Query().Get("limit")
    limit, _ := strconv.Atoi(limitParam)
    if limit == 0 {
        limit = 10 // Default limit
    }

    scores, err := services.GetGameLeaderboard(gameID, limit)
    if err != nil {
        http.Error(w, "Failed to retrieve game leaderboard", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "leaderboard": scores,
    })
}

// StoreScoreHandler handles requests for storing a new score
func StoreScoreHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID int    `json:"user_id"`
        GameID string `json:"game_id"`
        Score  int    `json:"score"`
    }

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err = services.StoreScore(req.UserID, req.GameID, req.Score)
    if err != nil {
        http.Error(w, "Failed to store score", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Score stored successfully"})
}
