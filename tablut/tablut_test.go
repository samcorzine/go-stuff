package tests

import(
  "practice/tablut/main"
  "testing"
  "fmt"
)

func TestMoveChecker(t *testing.T) {
    testGame := Game{board: main.getStartingBoard(), nextPlayer: -1}

    // should return true
    testMove1 := Move{startx : 0, starty : 3, endx: 0, endy: 2, player: -1}
    if testGame.moveIsValid(testMove1) != true {
      fmt.Println("valid move was rejected")
    }
    // should return false, jumps a piece
    testMove2 := Move{startx : 0, starty : 5, endx: 7, endy: 5, player: -1}
    if testGame.moveIsValid(testMove2) != false {
      fmt.Println("piece jumped other piece and was accepted")
    }
    // should return false, lands on a piece
    testMove3 := Move{startx : 0, starty : 3, endx: 0, endy: 4, player: -1}
    if testGame.moveIsValid(testMove3) != false {
      fmt.Println("piece landed on other piece and was accepted")
    }
    // should return false, not blacks piece being moved
    testMove4 := Move{startx : 2, starty : 4, endx: 2, endy: 2, player: -1}
    if testGame.moveIsValid(testMove4) != false {
      fmt.Println("player tried to move piece that wasn't there's and was accepted")
    }
    // should return false, not white's turn
    testMove5 := Move{startx : 2, starty : 4, endx: 2, endy: 2, player: 1}
    if testGame.moveIsValid(testMove5) != false {
      fmt.Println("wrong player made move, and was accepted")
    }
}
