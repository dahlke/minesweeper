# Minesweeper

Go(lang) implementation of Minesweeper game API

More about the game: https://en.wikipedia.org/wiki/Minesweeper_(video_game)

# Build

```
  $ go build -o build/minesweeper ./cmd
```

# Run

```
  $ ./build/minesweeper
```

# Create a New Game

```
  $ curl -i -X POST '127.0.0.1:3000/game' -d '{"name": "teste", "rows": 10, "cols": 8, "mines": 20}'
```

## Start the game

```
  $ curl -i -X POST '127.0.0.1:3000/game/teste/start'
```

# Run tests

```
  $ go clean  $(go list ./... | grep -v /vendor/)
  $ go test  $(go list ./... | grep -v /vendor/) -v
```

