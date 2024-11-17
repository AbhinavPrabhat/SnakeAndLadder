package models

// Game represents the state of a snake and ladder game
type Game struct {
    ID           string     `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
    Players      []Player   `json:"players"`
    CurrentTurn  int        `json:"currentTurn" example:"0"`
    Board        Board      `json:"board"`
    GameStatus   GameStatus `json:"gameStatus" example:"active"`
    Winner       *Player    `json:"winner,omitempty"`
}

// Player represents a player in the game
type Player struct {
    ID       string `json:"id" example:"player-123"`
    Name     string `json:"name" example:"John"`
    Position int    `json:"position" example:"1"`
}

type GameStatus string

const (
    StatusWaiting   GameStatus = "waiting"
    StatusActive    GameStatus = "active"
    StatusCompleted GameStatus = "completed"
)

// Board represents the snake and ladder game board
type Board struct {
    Snakes  map[int]int `json:"snakes"`
    Ladders map[int]int `json:"ladders"`
    Size    int         `json:"size" example:"100"`
}

// CreateLobbyRequest represents the request to create a new game lobby
type CreateLobbyRequest struct {
    PlayerName string `json:"playerName" example:"John" binding:"required"`
}

// JoinLobbyRequest represents the request to join an existing game lobby
type JoinLobbyRequest struct {
    LobbyID    string `json:"lobbyId" example:"123e4567-e89b-12d3-a456-426614174000" binding:"required"`
    PlayerName string `json:"playerName" example:"Jane" binding:"required"`
}

// RollDiceRequest represents the request to roll dice
type RollDiceRequest struct {
    GameID   string `json:"gameId" example:"123e4567-e89b-12d3-a456-426614174000" binding:"required"`
    PlayerID string `json:"playerId" example:"player-123" binding:"required"`
}

// GameResponse represents the response structure for game-related endpoints
type GameResponse struct {
    Game    *Game  `json:"game,omitempty"`
    Message string `json:"message,omitempty" example:"Operation successful"`
    Error   string `json:"error,omitempty" example:"Invalid request format"`
}