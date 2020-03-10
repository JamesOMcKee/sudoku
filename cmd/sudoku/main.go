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

//Essentially the following functions create a list of possibilities for each
//spot on the board, 1-9. Pruning happens at each spot on the board which has
//no value (value set to 0). Each value in the possibility array gets set to
//0 if the number exists in the same row, or column, or box. A spot that can
//have a 2, a 4, or a 5 after that would then have a poss array
//of {0,2,0,4,5,0,0,0,0}

func initializePossibilities(board sudokuBoard) [9][9][9]int {
  var poss[9][9][9] int
  for row := 0; row < 9; row++{
    for col := 0; col < 9; col++{
      for p := 0; p < 9; p++{
        poss[row][col][p] = p+1
      }
    }
  }
  return poss
}

func pruneByRow(toPrune[9][9][9] int, board sudokuBoard, row int, col int) [9][9][9]int {
  prune := toPrune
  //make list of items in row
  var list[9] int
  for index := 0; index < 9; index++{
    if board.board[row][index] != 0{
      list[board.board[row][index]-1] = board.board[row][index]
    }
  }

  //if item exists both in col and in target's poss array, remove it from poss array (set 0)
  for comp := 0; comp < 9; comp++{
    if(list[comp] == prune[row][col][comp]){
      prune[row][col][comp] = 0
    }
  }

  return prune
}

func pruneByCol(toPrune[9][9][9] int, board sudokuBoard, row int, col int) [9][9][9]int {
  prune := toPrune
  //make list of items in col
  var list[9] int
  for index := 0; index < 9; index++{
    if board.board[index][col] != 0{
      list[board.board[index][col]-1] = board.board[index][col]
    }
  }

  //if item exists both in col and in target's poss array, remove it from poss array (set 0)
  for comp := 0; comp < 9; comp++{
    if(list[comp] == prune[row][col][comp]){
      prune[row][col][comp] = 0
    }
  }

  return prune
}

func pruneByBox(toPrune[9][9][9] int, board sudokuBoard, row int, col int) [9][9][9]int {
  prune := toPrune
  //what box are we in?
  rowSet := row/3
  colSet := col/3
  //box := ((rowSet+1)*3) + (colSet+1)

  //make list of items in box
  var list[9] int
  for row := 0; row < 3; row++{
    for col := 0; col < 3; col++{
      if board.board[row+(rowSet*3)][col+(colSet*3)] != 0{
        list[board.board[row+(rowSet*3)][col+(colSet*3)]-1] = board.board[row+(rowSet*3)][col+(colSet*3)]
      }
    }
  }
  //if item exists in box and in target's poss array, remove it from poss array (set 0)
  for comp := 0; comp < 9; comp++{
    if list[comp] == prune[row][col][comp]{
      prune[row][col][comp] = 0
    }
  }
  return prune
}

func basicPrunePossibilities(poss[9][9][9] int, board sudokuBoard) [9][9][9]int {
  var pruned[9][9][9] int
  for row := 0; row < 9; row++{
    for col := 0; col < 9; col++{
      if board.board[row][col] == 0{
        pruned = pruneByRow(pruned, board, row, col)
        pruned = pruneByCol(pruned, board, row, col)
        pruned = pruneByBox(pruned, board, row, col)
      }
    }
  }
  return pruned
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
