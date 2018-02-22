package main

import (
)

type Game struct{
  board [9][9] int
  nextPlayer int
}

type Move struct{
  player int
  startx, starty, endx, endy int
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
  player := mov.player
  endx := mov.endx
  endy := mov.endy
  startx := mov.startx
  starty := mov.starty
  // checks to make sure destination is nonempty
  if (gam.board[endx][endy]) != 0 {
    return false
  }
  // checks that the starting position has that players piece in it
  if (gam.board[startx][starty] != player) {
    if (gam.board[startx][starty] != 10 || player == -1){
      return false
    }
  }
  // checks that it is submitting player's turn
  if (gam.nextPlayer != player) {
    return false
  }
  // checks that the move isn't a null move, start space different than end space
  if (startx == endx && starty == endy) {
    return false
  }
  // checks that the move is in either the same row or the same column
  if (startx != endx && starty != endy) {
    return false
  }
  // checks that that spaces in between the start and end are empty
  minx := min(startx, endx)
  maxx := max(startx, endx)
  miny := min(starty, endy)
  maxy := max(starty, endy)
  if (maxx - minx != 0 ) {
    for xtester := minx + 1; xtester < maxx; xtester++ {
      if (gam.board[xtester][miny] != 0){
        return false
      }
    }
  }
  if (maxy - miny != 0) {
    for ytester := miny + 1; ytester < maxy; ytester++ {
      if (gam.board[minx][ytester] != 0){
        return false
      }
    }
  }
  // checks that pieces aren't landing on the throne
  if(endy == 4 && endx == 4){
    return false
  }

  return true



}

func main(){

}
