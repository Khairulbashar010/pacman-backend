package game

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

// WebSocket upgrader to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow requests from any origin
    },
}

// Connected clients (key: WebSocket connection, value: true)
var clients = make(map[*websocket.Conn]bool)

// Broadcast channel to send messages to all clients
var broadcast = make(chan Message)

// Message defines the structure of WebSocket messages
type Message struct {
    PlayerID string `json:"playerId"`
    GameID   string `json:"gameId"`
    X        int    `json:"x"`
    Y        int    `json:"y"`
    Action   string `json:"action"`
}

// HandleConnections manages WebSocket connections and listens for messages
func HandleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade the HTTP request to a WebSocket connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("WebSocket upgrade error: %v", err)
        return
    }
    defer ws.Close()

    // Register the new client
    clients[ws] = true

    // Listen for incoming messages from this WebSocket connection
    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("WebSocket error: %v", err)
            delete(clients, ws)
            break
        }
        // Send the received message to the broadcast channel
        broadcast <- msg
    }
}

// HandleMessages listens on the broadcast channel and sends messages to all connected clients
func HandleMessages() {
    for {
        // Grab the next message from the broadcast channel
        msg := <-broadcast

        // Send the message to all connected clients
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Error sending WebSocket message: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}
