package api

import (

    "github.com/gorilla/mux"
    "packman-backend/internal/api/handlers"
    "packman-backend/internal/api/middleware"
)

func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    // Authentication routes
    router.HandleFunc("/api/auth/signup", handlers.SignupHandler).Methods("POST")
    router.HandleFunc("/api/auth/login", handlers.LoginHandler).Methods("POST")

    // Game routes (protected by JWT middleware)
    gameRouter := router.PathPrefix("/api/game").Subrouter()
    gameRouter.Use(middleware.JwtAuthMiddleware)

    gameRouter.HandleFunc("/join", handlers.JoinGameHandler).Methods("POST")
    gameRouter.HandleFunc("/leave", handlers.LeaveGameHandler).Methods("POST")
    gameRouter.HandleFunc("/move", handlers.MovePlayerHandler).Methods("POST")

    // Leaderboard routes
    router.HandleFunc("/api/leaderboard", handlers.GetGlobalLeaderboardHandler).Methods("GET")
    router.HandleFunc("/api/leaderboard/{gameId}", handlers.GetGameLeaderboardHandler).Methods("GET")

	// Leaderboard routes
    router.HandleFunc("/api/leaderboard", handlers.GetGlobalLeaderboardHandler).Methods("GET")
    router.HandleFunc("/api/leaderboard/{gameId}", handlers.GetGameLeaderboardHandler).Methods("GET")
    router.HandleFunc("/api/leaderboard/store", handlers.StoreScoreHandler).Methods("POST")

    return router
}
