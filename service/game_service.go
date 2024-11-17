package service

import (
    "errors"
    "sync"
    "task/models"
    "task/util"
    "github.com/google/uuid"
)

type GameService struct {
    games     map[string]*models.Game
    mutex     sync.RWMutex
    dice      util.DiceRoller
}

func NewGameService() *GameService {
    return &GameService{
        games: make(map[string]*models.Game),
        dice:  util.NewStandardDice(),
    }
}

func (s *GameService) CreateLobby(playerName string) (*models.Game, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    gameID := uuid.New().String()
    playerID := uuid.New().String()

    // Initialize board 
    board := models.Board{
        Size:    100,
        Snakes:  map[int]int{
            16: 6,
            47: 26,
            49: 11,
            56: 53,
            62: 19,
            64: 60,
            87: 24,
            93: 73,
            95: 75,
            98: 78,
        },
        Ladders: map[int]int{
            1:  38,
            4:  14,
            9:  31,
            21: 42,
            28: 84,
            36: 44,
            51: 67,
            71: 91,
            80: 100,
        },
    }

    game := &models.Game{
        ID:          gameID,
        Players:     []models.Player{{ID: playerID, Name: playerName, Position: 1}},
        CurrentTurn: 0,
        Board:       board,
        GameStatus:  models.StatusWaiting,
    }

    s.games[gameID] = game
    return game, nil
}

func (s *GameService) JoinLobby(gameID, playerName string) (*models.Game, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    game, exists := s.games[gameID]
    if !exists {
        return nil, errors.New("game not found")
    }

    if game.GameStatus != models.StatusWaiting {
        return nil, errors.New("game already started")
    }

    if len(game.Players) >= 4 {
        return nil, errors.New("game lobby is full")
    }

    playerID := uuid.New().String()
    game.Players = append(game.Players, models.Player{
        ID:       playerID,
        Name:     playerName,
        Position: 1,
    })

    if len(game.Players) >= 2 {
        game.GameStatus = models.StatusActive
    }

    return game, nil
}

func (s *GameService) RollDice(gameID, playerID string) (*models.Game, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    game, exists := s.games[gameID]
    if !exists {
        return nil, errors.New("game not found")
    }

    if game.GameStatus != models.StatusActive {
        return nil, errors.New("game is not active")
    }

    currentPlayer := &game.Players[game.CurrentTurn]
    if currentPlayer.ID != playerID {
        return nil, errors.New("not your turn")
    }

    // Roll the dice
    roll := s.dice.Roll()

    // Calculate new position
    newPosition := currentPlayer.Position + roll
    
    // Check if the position is within bounds
    if newPosition > game.Board.Size {
        return game, nil // Player needs exact number to win
    }

    // Update position and check for snakes or ladders
    currentPlayer.Position = newPosition
    
    // Check for snakes
    if newPos, exists := game.Board.Snakes[newPosition]; exists {
        currentPlayer.Position = newPos
    }
    
    // Check for ladders
    if newPos, exists := game.Board.Ladders[newPosition]; exists {
        currentPlayer.Position = newPos
    }

    // Check for win condition
    if currentPlayer.Position == game.Board.Size {
        game.GameStatus = models.StatusCompleted
        game.Winner = currentPlayer
        return game, nil
    }

    // Move to next player's turn
    game.CurrentTurn = (game.CurrentTurn + 1) % len(game.Players)

    return game, nil
}

func (s *GameService) GetGame(gameID string) (*models.Game, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    game, exists := s.games[gameID]
    if !exists {
        return nil, errors.New("game not found")
    }

    return game, nil
}