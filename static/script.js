function collectBoard() {
    let board = [];
    for (let i = 0; i < 9; i++) {
        let row = [];
        for (let j = 0; j < 9; j++) {
            let cellValue = document.getElementById(`cell-${i}-${j}`).value;
            row.push(cellValue === "" ? 0 : parseInt(cellValue));
        }
        board.push(row);
    }
    return board;
}


function solveSudoku() {
    const board = collectBoard();
    fetch('/solve', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ board: board })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            updateBoard(data.board); 
        }
    })
    .catch(error => console.error('Error:', error));
}

function updateBoard(board) {
    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            document.getElementById(`cell-${i}-${j}`).value = board[i][j];
        }
    }
}
