export interface GameMessage {
  playerId: string;
  gameId:   string;
}

export interface MoveMessage extends GameMessage {
  type: "move";
  player: string;
  move: number;
}

export interface EventMessage extends GameMessage {
  type: string;
  payload?: any;
}

export const EventTypes = {
  START_GAME: 'start_game',
  WIN:        'win',
  LOSE:       'lose',
  DRAW:       'draw',
  MOVE:       'move',
}

export type EventType = typeof EventTypes[keyof typeof EventTypes] 
