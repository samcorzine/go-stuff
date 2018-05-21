package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "fmt"
  // "strconv"
  "log"
  // "strings"
  // "os"
)

type Game struct{
  // board has entries -1 -> black, 0 -> empty tile, 1 -> white, 10 -> king
  // nextPlayer is either -1 or 1 for black or white
  Board [9][9] int `json:"game"`
  NextPlayer int `json:"nextPlayer"`
  Winner int `json:winner`
}

type Move struct{
  Player *int `json:"player,omitempty"`
  Startx *int `json:"startx,omitempty"`
  Starty *int `json:"starty,omitempty"`
  Endx *int  `json:"endx,omitempty"`
  Endy *int `json:"endy,omitempty"`
}

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


func getStartingBoard() [9][9]int {
  var board = [9][9]int{}
  var blackspaces = [16][2]int{{0,3}, {0,4}, {0,5}, {8,3}, {8,4}, {8,5}, {3,0}, {4,0}, {5,0}, {3,8}, {4,8}, {5,8}, {1,4}, {7,4}, {4,1}, {4,7} }
  var whitespaces = [8][2]int{{2,4}, {3,4}, {5,4}, {6,4}, {4,2}, {4,3}, {4,5}, {4,6}}
  for _, space := range blackspaces {
    board[space[0]][space[1]] = -1
  }
  for _, space := range whitespaces {
    board[space[0]][space[1]] = 1
  }
  board[4][4] = 10
  return board
}

func (gam Game) moveIsValid(mov Move) bool {
  player := *mov.Player
  endx := *mov.Endx
  endy := *mov.Endy
  startx := *mov.Startx
  starty := *mov.Starty
  // checks to make sure destination is nonempty
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

func (gam *Game) isOtherPlayer(myPlayerNum int, theirX int, theirY int) bool{
  if myPlayerNum == -1 {
    if gam.Board[theirX][theirY] == 10 || gam.Board[theirX][theirY] == 1{
      return true
    }
  }
  if myPlayerNum == 1 {
    if gam.Board[theirX][theirY] == -1{
      return true
    }
  }
  return false
}

func (gam *Game) isSamePlayer(myPlayerNum int, theirX int, theirY int) bool {
  if gam.isOtherPlayer(myPlayerNum, theirX, theirY) || gam.Board[theirX][theirY] == 0 {
    return false
  }
  return true
}

func (gam Game) victoryCheck() int {
  foundKing := false
  kingX := -1
  kingY := -1
  for xCounter := 0; xCounter < 9; xCounter++ {
    for yCounter := 0; yCounter < 9; yCounter++ {
      if gam.Board[xCounter][yCounter] == 10 {
        foundKing = true
        kingX = xCounter
        kingY = yCounter
      }
    }
  }
  if !foundKing {
    return -1
  }
  if kingX == 0 || kingX == 8 || kingY == 0 || kingY == 8{
    return 1
  }
  return 0
}

func (gam *Game) checkForCapture(targetX int, targetY int, myX int, myY int, curPlayer int) bool {
  // Doesn't handle throne effects for non-king pieces

<<<<<<< HEAD
=======

  // Handles potential King Captures
>>>>>>> c0b649bdfe588c556fdf544c52ba04373b59d7f6
  if gam.Board[targetX][targetY] == 10  && curPlayer == -1 {
    // On throne
    if targetX == 4 && targetY == 4 {
      if gam.Board[5][4] == curPlayer && gam.Board[3][4] == curPlayer && gam.Board[4][5] == curPlayer && gam.Board[4][3] == curPlayer {
        return true
      }
    } else if targetX == 5 && targetY == 4{
      if gam.Board[6][4] == curPlayer && gam.Board[5][5] == curPlayer && gam.Board[5][3] == curPlayer {
        return true
      }
    } else if targetX == 3 && targetY == 4{
      if gam.Board[2][4] == curPlayer && gam.Board[3][5] == curPlayer && gam.Board[3][3] == curPlayer {
        return true
      }
    } else if targetX == 4 && targetY == 5 {
      if gam.Board[4][6] == curPlayer && gam.Board[3][5] == curPlayer && gam.Board[5][5] == curPlayer {
        return true
      }
    } else if targetX == 4 && targetY == 3{
      if gam.Board[4][2] == curPlayer && gam.Board[3][3] == curPlayer && gam.Board[5][3] == curPlayer {
        return true
      }
    } else if myX != targetX {
      if myX > targetX && targetX != 0{
        if gam.isSamePlayer(curPlayer, targetX - 1, targetY){
          return true
        }
      }
      if myX < targetX && targetX != 8{
        if gam.isSamePlayer(curPlayer, targetX + 1, targetY){
          return true
        }
      }
    } else {
      if myY > targetY && targetY != 0{
        if gam.isSamePlayer(curPlayer, targetX, targetY - 1){
          return true
        }
      }
      if myY < targetY && targetY != 8{
        if gam.isSamePlayer(curPlayer, targetX, targetY + 1){
          return true
        }
      }
    }
  } else {
    if myX != targetX {
      if myX > targetX && targetX != 0{
        if gam.isSamePlayer(curPlayer, targetX - 1, targetY) || (targetX - 1 == 4 && targetY == 4 && gam.Board[4][4] != 10) {
          return true
        }
      }
      if myX < targetX && targetX != 8{
        if gam.isSamePlayer(curPlayer, targetX + 1, targetY) || (targetX + 1 == 4 && targetY == 4 && gam.Board[4][4] != 10) {
          return true
        }
      }
    } else {
      if myY > targetY && targetY != 0{
        if gam.isSamePlayer(curPlayer, targetX, targetY - 1) || (targetX == 4 && targetY - 1 == 4 && gam.Board[4][4] != 10) {
          return true
        }
      }
      if myY < targetY && targetY != 8{
        if gam.isSamePlayer(curPlayer, targetX, targetY + 1) || (targetX == 4 && targetY + 1 == 4 && gam.Board[4][4] != 10) {
          return true
        }
      }
    }
  }
  return false
}


func (gam *Game) doMove(mov Move) {
  player := *mov.Player
  endx := *mov.Endx
  endy := *mov.Endy
  startx := *mov.Startx
  starty := *mov.Starty
  gam.Board[endx][endy] = gam.Board[startx][starty]
  gam.Board[startx][starty] = 0

  if endx < 8 && gam.checkForCapture(endx + 1, endy, endx, endy, player){
    gam.Board[endx + 1][endy] = 0
  }
  if endx > 0 && gam.checkForCapture(endx - 1, endy, endx, endy, player){
    gam.Board[endx - 1][endy] = 0
  }
  if endy < 8 && gam.checkForCapture(endx, endy + 1, endx, endy, player){
    gam.Board[endx][endy + 1] = 0
  }
  if endy > 0 && gam.checkForCapture(endx, endy - 1, endx, endy, player){
    gam.Board[endx][endy - 1] = 0
  }


  if gam.NextPlayer == -1 {
    fmt.Println("Setting next player to 1")
    gam.NextPlayer = 1
  } else {
    fmt.Println("Setting next player to -1")
    gam.NextPlayer = -1
  }
}

func altMoveHandler(w http.ResponseWriter, r *http.Request) {
    // params := mux.Vars(r)
    // fmt.Println("Params")
    // fmt.Println(params)
    var move Move
    // err := json.NewDecoder(r.Body).Decode(&move)
    // var moveArr [1]Movestr
    // textbytes = []byte(text)
    var _ = json.NewDecoder(r.Body).Decode(&move)
    // fmt.Println(r.FormValue("player"))
    // fmt.Println("Error:")
    // fmt.Println(err)
    // fmt.Println("Request:")
    // fmt.Println(*r)
    // fmt.Println("Request body:")
    // fmt.Println(r.Body)
    // fmt.Println("Decoded move:")
    // fmt.Println(move)
    if game.moveIsValid(move) {
      fmt.Println("Move was valid")
      (&game).doMove(move)
    } else {
      fmt.Println("Move was invalid")
    }
    for xCounter := 0 ; xCounter < 9; xCounter++ {
      fmt.Println(game.Board[xCounter])
    }
    fmt.Println("*****")
    // fmt.Println(game.NextPlayer)
    // fmt.Println(move.Player)
    if game.victoryCheck() == 1{
      fmt.Println("White Won")
      game.Winner = 1
    } else if game.victoryCheck() == -1 {
      fmt.Println("Black Won")
      game.Winner = -1
    }
    json.NewEncoder(w).Encode(game)

    // move.player, _ = strconv.Atoi(params["player"])
    // people = append(people, person)
    // var move2 = Move{1,1,1,1,1}
    // fmt.Println("All 1's move:")
    // fmt.Println(move2)
    // json.NewEncoder(w).Encode(move)
}
// // func moveHandler(w http.ResponseWriter, r *http.Request){
// //   params := mux.Vars(r)
// //   fmt.Println(params["player"])
// //   player := strconv.Atoi(params["player"])
// //   startx := strconv.Atoi(params["startx"])
// //   starty := strconv.Atoi(params["starty"])
// //   endx := strconv.Atoi(params["endx"])
// //   endy := strconv.Atoi(params["endy"])
// //   var move = Move{player, startx , starty, endx, endy}
// //
// //   // err := json.NewDecoder(r.Body).Decode(&move)
// //   // fmt.Println(err)
  // if game.moveIsValid(move) {
  //   game.doMove(move)
  // }
  // fmt.Println(move.player)
  // json.NewEncoder(w).Encode(game)
// //
// // }
//

func gameGetter(w http.ResponseWriter, r *http.Request){
  _ = json.NewEncoder(w).Encode(game)
}
var game = Game{getStartingBoard(), -1, 0}



func main(){
  router := mux.NewRouter()
  router.HandleFunc("/submit", altMoveHandler).Methods("POST")
  router.HandleFunc("/game", gameGetter).Methods("GET")
  log.Fatal(http.ListenAndServe(":8765", router))
  // http.HandleFunc("/", altMoveHandler)
  // if err := http.ListenAndServe(":8765", nil); err != nil {
  //   panic(err)
  // }
}
