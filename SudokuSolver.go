package main

import "fmt"

type sudokuBoard struct {
  board[9][9] int
  //row, col
}

func testRow(row int, board sudokuBoard) bool {
  for col := 0; col < 9; col++{
    for j := 0; j < col; j++{
      if board.board[row][col] == board.board[row][j]{
        return false
      }
    }
  }
  return true
}

func testCol(col int, board sudokuBoard) bool {
  for row := 0; row < 9; row++{
    for j := 0; j < row; j++{
      if board.board[row][col] == board.board[j][col]{
        return false
      }
    }
  }
  return true
}

func testBox(box int, board sudokuBoard) bool {
  //set of rows = (box)/3
  //set of cols = (box-1)%3
  //box 1 is 0,0 | box 7 is 2,0 | box 9 is 2,2
  rowSet := (box)/3 //0-2
  colSet := (box)%3 //0-2
  var test[9] int
  for row := 0; row < 3; row++{
    for col := 0; col < 3; col++{
      test[(row*3)+col] = board.board[row+(rowSet*3)][col+(colSet*3)]
      for k := 0; k < ((row*3)+col); k++{
        if test[k] == test[(row*3)+col]{
          return false
        }
      }
    }
  }
  return true
}

func testBoard(board sudokuBoard) bool {
  for i := 0; i < 9; i++{
    a := testRow(i,board)
    b := testCol(i,board)
    c := testBox(i,board)
    if !a || !b || !c {
      return false
    }
  }
  return true
}

func newBoard(board[9][9] int) sudokuBoard {
  newBoard := sudokuBoard{board: board}
  return newBoard
}

func main(){
  boardToBe := [9][9]int {
    {2,9,6,3,1,8,5,7,4},
    {5,8,4,9,7,2,6,1,3},
    {7,1,3,6,4,5,2,8,9},
    {6,2,5,8,9,7,3,4,1},
    {9,3,1,4,2,6,8,5,7},
    {4,7,8,5,3,1,9,2,6},
    {1,6,7,2,5,3,4,9,8},
    {8,5,9,7,6,4,1,3,2},
    {3,4,2,1,8,9,7,6,5}}
  board := newBoard(boardToBe)
  trueFalse := testBoard(board)
  fmt.Printf("%t",trueFalse)
}
