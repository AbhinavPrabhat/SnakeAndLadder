package controllers

import (
    "net/http"
    "task/models"
    "task/service"
    "github.com/labstack/echo/v4"
)

type GameController struct {
    gameService *service.GameService
}

func NewGameController(gameService *service.GameService) *GameController {
    return &GameController{
        gameService: gameService,
    }
}

// CreateLobby godoc
// @Summary Create a new game lobby
// @Description Creates a new game lobby and returns the lobby ID
// @Tags lobby
// @Accept json
// @Produce json
// @Param request body models.CreateLobbyRequest true "Create Lobby Request"
// @Success 201 {object} models.GameResponse
// @Failure 400 {object} models.GameResponse
// @Router /lobby/create [post]
func (gc *GameController) CreateLobby(c echo.Context) error {
    var req models.CreateLobbyRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, models.GameResponse{
            Error: "Invalid request format",
        })
    }

    game, err := gc.gameService.CreateLobby(req.PlayerName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, models.GameResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(http.StatusCreated, models.GameResponse{
        Game:    game,
        Message: "Game lobby created successfully",
    })
}

// JoinLobby godoc
// @Summary Join an existing game lobby
// @Description Allows a player to join an existing lobby using the lobby ID
// @Tags lobby
// @Accept json
// @Produce json
// @Param request body models.JoinLobbyRequest true "Join Lobby Request"
// @Success 200 {object} models.GameResponse
// @Failure 400 {object} models.GameResponse
// @Router /lobby/join [post]
func (gc *GameController) JoinLobby(c echo.Context) error {
    var req models.JoinLobbyRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, models.GameResponse{
            Error: "Invalid request format",
        })
    }

    game, err := gc.gameService.JoinLobby(req.LobbyID, req.PlayerName)
    if err != nil {
        return c.JSON(http.StatusBadRequest, models.GameResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, models.GameResponse{
        Game:    game,
        Message: "Successfully joined the game",
    })
}

// RollDice godoc
// @Summary Roll dice for the current player
// @Description Rolls a dice and updates the current player's position
// @Tags game
// @Accept json
// @Produce json
// @Param request body models.RollDiceRequest true "Roll Dice Request"
// @Success 200 {object} models.GameResponse
// @Failure 400 {object} models.GameResponse
// @Router /game/roll [post]
func (gc *GameController) RollDice(c echo.Context) error {
    var req models.RollDiceRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, models.GameResponse{
            Error: "Invalid request format",
        })
    }

    game, err := gc.gameService.RollDice(req.GameID, req.PlayerID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, models.GameResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, models.GameResponse{
        Game:    game,
        Message: "Dice rolled successfully",
    })
}

// GetGame godoc
// @Summary Get game status
// @Description Retrieves the current state of a game
// @Tags game
// @Accept json
// @Produce json
// @Param gameId path string true "Game ID"
// @Success 200 {object} models.GameResponse
// @Failure 404 {object} models.GameResponse
// @Router /game/{gameId} [get]
func (gc *GameController) GetGame(c echo.Context) error {
    gameID := c.Param("gameId")
    game, err := gc.gameService.GetGame(gameID)
    if err != nil {
        return c.JSON(http.StatusNotFound, models.GameResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, models.GameResponse{
        Game:    game,
        Message: "Game retrieved successfully",
    })
}