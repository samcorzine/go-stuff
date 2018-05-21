package main


import (
  "net/http"
  "encoding/json"
  "fmt"
  "math/rand"
  "bytes"
  // "log"
  // "strconv"
  // "log"
  // "os"
  "time"
  "os"
  "strconv"
)


func min(a, b int) int {
  if a < b {
      return a
  }
  return b
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}


type Game struct{
  // board has entries -1 -> black, 0 -> empty tile, 1 -> white, 10 -> king
  // nextPlayer is either -1 or 1 for black or white
  Board [9][9] int `json:"game"`
  NextPlayer int `json:"nextPlayer"`
  Winner int `json:winner`
}

type Coordinate struct{
  X int `json:"x"`
  Y int `json:"y"`
}

type Move struct{
  Player *int `json:"player,omitempty"`
  Startx *int `json:"startx,omitempty"`
  Starty *int `json:"starty,omitempty"`
  Endx *int  `json:"endx,omitempty"`
  Endy *int `json:"endy,omitempty"`
}


func (gam Game) moveIsValid(mov Move) bool {
  player := *mov.Player
  endx := *mov.Endx
  endy := *mov.Endy
  startx := *mov.Startx
  starty := *mov.Starty

  // checks to make sure destination is empty
  if (gam.Board[endx][endy]) != 0 {
    return false
  }
  // fmt.Println("tile wasn't empty")
  // checks that the starting position has that players piece in it
  if (gam.Board[startx][starty] != player) {
    if (gam.Board[startx][starty] != 10 || player == -1){
      return false
    }
  }
  // fmt.Println("Player didn't try to move opponents piece")

  // checks that it is submitting player's turn
  if (gam.NextPlayer != player) {
    return false
  }
  // fmt.Println("The correct player went")
  // checks that the move isn't a null move, start space different than end space
  if (startx == endx && starty == endy) {
    return false
  }
  // fmt.Println("Move was non-null")

  // checks that the move is in either the same row or the same column
  if (startx != endx && starty != endy) {
    return false
  }
  // fmt.Println("Move was in the same row or in the same column")
  // checks that that spaces in between the start and end are empty
  minx := min(startx, endx)
  maxx := max(startx, endx)
  miny := min(starty, endy)
  maxy := max(starty, endy)
  if (maxx - minx != 0 ) {
    for xtester := minx + 1; xtester < maxx; xtester++ {
      if (gam.Board[xtester][miny] != 0){
        return false
      }
    }
  }
  if (maxy - miny != 0) {
    for ytester := miny + 1; ytester < maxy; ytester++ {
      if (gam.Board[minx][ytester] != 0){
        return false
      }
    }
  }
  // fmt.Println("No pieces were in the way")
  // checks that pieces aren't landing on the throne
  if(endy == 4 && endx == 4){
    return false
  }
  // fmt.Println("Piece did not land on the throne")
  return true
}

func (gam Game) myPieces(playerNum int) []Coordinate {
  var coordinateHolder []Coordinate
  for xCounter := 0; xCounter < 9; xCounter++ {
    for yCounter := 0; yCounter < 9; yCounter++{
      if gam.Board[xCounter][yCounter] == playerNum || (playerNum == 1 && gam.Board[xCounter][yCounter] == 10){
        coor := Coordinate{X:xCounter , Y:yCounter}
        coordinateHolder = append(coordinateHolder, coor)
      }
    }
  }
  return coordinateHolder
}


func defenseMoveMaker(gameState Game) Move {

}


func moveMaker(gameState Game, clientPlayerNum int) Move {
  var foundValidMove bool
  availablePieces := gameState.myPieces(clientPlayerNum)
  for foundValidMove != true {
    coorIndex := rand.Intn(len(availablePieces))
    startingCoordinate := availablePieces[coorIndex]
    startx := startingCoordinate.X
    starty := startingCoordinate.Y
    endx := rand.Intn(9)
    endy := rand.Intn(9)
    move := Move{Player:&clientPlayerNum, Startx:&startx, Starty:&starty, Endx:&endx, Endy:&endy}
    if gameState.moveIsValid(move){
      return move
    }
  }
  ten := 10
  return Move{&ten, &ten, &ten, &ten, &ten}
}


func main(){
  clientPlayerNum, _ := strconv.Atoi(os.Args[1])
  fmt.Println(clientPlayerNum)
  gameIsOn := true
  for gameIsOn {
    time.Sleep(time.Duration(1)*time.Second)
    resp, err := http.Get("http://localhost:8765/game")
    if err != nil {
      panic(err)
    }
    defer resp.Body.Close()
    var gameState Game
    _ = json.NewDecoder(resp.Body).Decode(&gameState)
    if gameState.Winner != 0 {
      break
    }
    if gameState.NextPlayer == clientPlayerNum{
      fmt.Println("Making a Move")
      move := moveMaker(gameState, clientPlayerNum)
      jsonValue, _ := json.Marshal(move)
      resp, err := http.Post("http://localhost:8765/submit", "application/json", bytes.NewBuffer(jsonValue))
      fmt.Println("Response:")
      fmt.Println(resp.Body)
      fmt.Println("Error:")
      fmt.Println(err)
    }

  }
  // http.HandleFunc("/", altMoveHandler)
  // if err := http.ListenAndServe(":8765", nil); err != nil {
  //   panic(err)
  // }
}
