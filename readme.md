# Snake and Ladder Game API

A RESTful API implementation of the classic Snake and Ladder game using Go and Echo framework.

## Features

- Create game lobbies
- Join existing game lobbies
- Roll dice and move players
- Implement snake and ladder rules
- Track game state and player positions
- Handle multiple concurrent games

## Prerequisites

- Go 1.16 or higher
- Echo framework v4

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd snake-ladder-game
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Create Lobby
Creates a new game lobby and returns the lobby ID.

- **URL**: `/lobby/create`
- **Method**: `POST`
- **Request Body**:
```json
{
    "playerName": "string"
}
```
- **Success Response**:
  - **Code**: 201 CREATED
  - **Content**:
```json
{
    "game": {
        "id": "string",
        "players": [{
            "id": "string",
            "name": "string",
            "position": 1
        }],
        "currentTurn": 0,
        "gameStatus": "waiting"
    },
    "message": "Game lobby created successfully"
}
```

### Join Lobby
Allows a player to join an existing lobby.

- **URL**: `/lobby/join`
- **Method**: `POST`
- **Request Body**:
```json
{
    "lobbyId": "string",
    "playerName": "string"
}
```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
```json
{
    "game": {
        "id": "string",
        "players": [
            {
                "id": "string",
                "name": "string",
                "position": 1
            }
        ],
        "currentTurn": 0,
        "gameStatus": "active"
    },
    "message": "Successfully joined the game"
}
```

### Roll Dice
Rolls a dice for the current player.

- **URL**: `/game/roll`
- **Method**: `POST`
- **Request Body**:
```json
{
    "gameId": "string",
    "playerId": "string"
}
```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
```json
{
    "game": {
        "id": "string",
        "players": [
            {
                "id": "string",
                "name": "string",
                "position": 1
            }
        ],
        "currentTurn": 1,
        "gameStatus": "active"
    },
    "message": "Dice rolled successfully"
}
```

### Get Game Status
Retrieves the current state of a game.

- **URL**: `/game/:gameId`
- **Method**: `GET`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
```json
{
    "game": {
        "id": "string",
        "players": [
            {
                "id": "string",
                "name": "string",
                "position": 1
            }
        ],
        "currentTurn": 0,
        "gameStatus": "active"
    },
    "message": "Game retrieved successfully"
}
```

## Error Responses

- **Code**: 400 BAD REQUEST
```json
{
    "error": "Error message describing the issue"
}
```

- **Code**: 404 NOT FOUND
```json
{
    "error": "Game not found"
}
```


## Board Configuration

The game board is configured with the following snakes and ladders:

### Snakes (Head → Tail)
- 16 → 6
- 47 → 26
- 49 → 11
- 56 → 53
- 62 → 19
- 64 → 60
- 87 → 24
- 93 → 73
- 95 → 75
- 98 → 78

### Ladders (Bottom → Top)
- 1 → 38
- 4 → 14
- 9 → 31
- 21 → 42
- 28 → 84
- 36 → 44
- 51 → 67
- 71 → 91
- 80 → 100

## Game Rules

1. Players start at position 1
2. Players take turns rolling a dice (1-6)
3. Players move forward according to dice roll
4. Landing on a snake head moves player down to the tail
5. Landing on a ladder bottom moves player up to the top
6. First player to reach position 100 wins
7. Exact roll needed to win (must roll exact number to reach 100)

