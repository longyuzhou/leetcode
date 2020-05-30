// Given a 2D board containing 'X' and 'O' (the letter O),
// capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// Example:
//   X X X X
//   X O O X
//   X X O X
//   X O X X
// After running your function, the board should be:
//   X X X X
//   X X X X
//   X X X X
//   X O X X

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solve = function (board) {
  const h = board.length;
  if (h <= 1) {
    return;
  }
  const w = board[0].length;

  const dfs = (x, y) => {
    board[y][x] = '*';
    if (y > 0 && board[y - 1][x] === 'O') {
      dfs(x, y - 1);
    }
    if (x < w - 1 && board[y][x + 1] === 'O') {
      dfs(x + 1, y);
    }
    if (y < h - 1 && board[y + 1][x] === 'O') {
      dfs(x, y + 1);
    }
    if (x > 0 && board[y][x - 1] === 'O') {
      dfs(x - 1, y);
    }
  };

  board.forEach((row, y) => {
    row.forEach((val, x) => {
      if (val === 'O' && (y === 0 || y === h - 1 || x === 0 || x === w - 1)) {
        dfs(x, y);
      }
    });
  });

  board.forEach((row, y) => {
    row.forEach((val, x) => {
      if (val === '*') {
        board[y][x] = 'O';
      } else if (val === 'O') {
        board[y][x] = 'X';
      }
    });
  });
};

const test = (board, expect) => {
  solve(board);
  let pass = true;
  for (let i = 0; i < expect.length; i++) {
    for (let j = 0; j < expect[i].length; j++) {
      if (expect[i][j] !== board[i][j]) {
        pass = false;
        break;
      }
    }
  }
  if (!pass) {
    console.log(`fail! expect ${expect}, got ${board}`);
  } else {
    console.log('pass');
  }
};

test(
  [
    ['O', 'O'],
    ['O', 'O'],
  ],
  [
    ['O', 'O'],
    ['O', 'O'],
  ]
);

test(
  [
    ['O', 'O', 'O'],
    ['O', 'O', 'O'],
    ['O', 'O', 'O'],
  ],
  [
    ['O', 'O', 'O'],
    ['O', 'O', 'O'],
    ['O', 'O', 'O'],
  ]
);

test(
  [
    ['X', 'X', 'X', 'X'],
    ['X', 'O', 'O', 'X'],
    ['X', 'X', 'O', 'X'],
    ['X', 'O', 'X', 'X'],
  ],
  [
    ['X', 'X', 'X', 'X'],
    ['X', 'X', 'X', 'X'],
    ['X', 'X', 'X', 'X'],
    ['X', 'O', 'X', 'X'],
  ]
);

solve([
  ['O', 'X', 'X', 'O', 'X'],
  ['X', 'O', 'O', 'X', 'O'],
  ['X', 'O', 'X', 'O', 'X'],
  ['O', 'X', 'O', 'O', 'O'],
  ['X', 'X', 'O', 'X', 'O'],
]);
