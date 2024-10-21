import { useEffect, useState } from "react";
import './Board.css';
import { MoveMessage, EventTypes } from "../../types/message";
import { checkWin } from "../../game/checkWin";

type Player = 'X' | 'O' | '';

interface BoardProps {
  socket: WebSocket | null;
}

export const Board = ({ socket }: BoardProps) => {
  const emptyBoard = Array(9).fill('');
  const [board, setBoard] = useState<string[]>(emptyBoard);
  const [player, setPlayer] = useState<Player>('');
  const [playerId, setPlayerId] = useState<string>('');
  const [gameId, setGameId] = useState<string>('');
  const [winner, setWinner] = useState<Player | null>(null);

  const [turn, setTurn] = useState(1)

  useEffect(() => {
    if (socket) {
      const handleMessage = (e: MessageEvent) => {
        const message = JSON.parse(e.data);
        const { playerId, gameId, payload } = message;

        if (message.type === EventTypes.START_GAME) {
          setBoard(emptyBoard);
          setPlayerId(playerId);
          setGameId(gameId);
          setPlayer(payload); // Set the current player (X or O)
        }

        if (message.type === 'move') {
          handleMove(message.move, message.player); // Update board with the received move
        }
      };

      socket.addEventListener('message', handleMessage);

      return () => {
        socket.removeEventListener('message', handleMessage);
      };
    }
  }, [socket, emptyBoard]); // Added emptyBoard to dependencies

  const handleMove = (move: number, player: string) => {
    if (board[move] || winner) return; // Prevent move if cell is occupied or there's a winner
    setTurn((turn) => turn + 1)

    const newBoard = [...board];
    newBoard[move] = player; // Mark the opponent's move
    setBoard(newBoard);

    // Check for winner after the move
    const isWinner = checkWin(newBoard);

    if (isWinner) {
      handleWinner(isWinner);
    }
  };

  const handleClick = (index: number) => {
    if (board[index] || winner) return; // Prevent move if cell is occupied or there's a winner
    setTurn((turn) => turn + 1)

    const newBoard = [...board];
    newBoard[index] = player; // Mark the move for the current player
    setBoard(newBoard);

    if (socket) {
      sendMove(playerId, index);
    }

    // Check for winner after the move
    const isWinner = checkWin(newBoard);
    if isWinner = turn == 9 ? "It's a tie" : ''
    if (isWinner) {
      handleWinner(isWinner);
    }
  };

  const sendMove = (playerId: string, move: number) => {
    const msg: MoveMessage = {
      type: "move",
      player,
      playerId,
      gameId,
      move,
    };

    if (socket) {
      socket.send(JSON.stringify(msg));
    }
  };

  const handleWinner = (winner: Player) => {
    setWinner(winner);
    console.log(`${winner} Wins`);
  };

  return (
    <>
      <div className="board">
        {board.map((cell, index) => (
          <div 
            key={index} 
            className={`cell ${winner ? 'disabled' : ''}`}
            onClick={() => handleClick(index)}
          >
            {cell}
          </div>
        ))}
      </div>
      {winner && <div className="winner-message">{`${winner} Wins!`}</div>}
    </>
  );
};
