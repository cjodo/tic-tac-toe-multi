
const buildBoard = (board:Array<string>) => {
  const newBoard = [[], [], []];
  let index = 0;
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      newBoard[i][j] = board[index];
      index++;
    }
  }
  return newBoard;
}

export const checkWin = (board: Array<string>) => {
  const newBoard = buildBoard(board);
  let winner = '';

  for (let row = 0; row < 3; row++) {
    // check for horizontal win
    let left = newBoard[row][0];
    let middle = newBoard[row][1];
    let right = newBoard[row][2];

    if (left || middle || right) {
      // if theres an undefined spot there can't be a win
      if (left == middle && middle == right) {
        winner = left; // Returns who won
        return winner;
      }
    }
  }
  for (let col = 0; col < 3; col++) {
    // Check for vertical win
    let left = newBoard[0][col];
    let middle = newBoard[1][col];
    let right = newBoard[2][col];
    if (left || middle || right) {
      if (left == middle && middle == right) {
        winner = left; // Returns who won
        return winner;
      }
    }
  }
  // Checks for diagonal win
  let UL = newBoard[0][0];
  let M = newBoard[1][1];
  let BR = newBoard[2][2];
  // Check the first diagonal
  if (UL || M || BR) {
    if (UL == M && M == BR) {
      winner = UL; // Returns who won
      return winner;
    }
  }

  let UR = newBoard[0][2];
  let BL = newBoard[2][0];
  if (UR || M || BL) {
    if (UR == M && M == BL) {
      winner = UR;
      return winner;
    }
  }
  return winner;
}
