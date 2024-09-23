package handlers

import (
    "encoding/json"
    "net/http"
)

type GameSession struct {
    GameID   string `json:"gameId"`
    PlayerID string `json:"playerId"`
}

type PlayerMove struct {
    GameID   string `json:"gameId"`
    PlayerID string `json:"playerId"`
    X        int    `json:"x"`
    Y        int    `json:"y"`
}

var gameSessions = make(map[string][]string) // Stores players in each game

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
    var session GameSession
    err := json.NewDecoder(r.Body).Decode(&session)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Add player to game session (create new if necessary)
    gameSessions[session.GameID] = append(gameSessions[session.GameID], session.PlayerID)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "gameId":   session.GameID,
        "playerId": session.PlayerID,
        "status":   "joined",
    })
}

func LeaveGameHandler(w http.ResponseWriter, r *http.Request) {
    var session GameSession
    err := json.NewDecoder(r.Body).Decode(&session)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    players := gameSessions[session.GameID]
    for i, playerID := range players {
        if playerID == session.PlayerID {
            gameSessions[session.GameID] = append(players[:i], players[i+1:]...)
            break
        }
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Player left the game",
    })
}

func MovePlayerHandler(w http.ResponseWriter, r *http.Request) {
    var move PlayerMove
    err := json.NewDecoder(r.Body).Decode(&move)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Broadcast the player's move to other players in the session (Stubbed here)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Player moved",
        "x":       string(rune(move.X)),
        "y":       string(rune(move.Y)),
    })
}
