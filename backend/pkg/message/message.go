package message

const (
	EventStartGame		= "start_game"
	EventWin					= "win"
	EventLose					= "lose"
	EventDraw					= "draw"
)

type GameMessage interface {
		GetPlayerId() string
		GetGameId()		string
}

type MoveMessage struct {
	Type			string	`json:"type"`
	PlayerId	string	`json:"playerId"`
	Player		string	`json:"player"`
	GameId		string	`json:"gameId"`
	Move			int			`json:"move"`
}

type EventMessage struct {
	Type			string			`json:"type"`
	PlayerId	string			`json:"playerId"`
	GameId		string			`json:"gameId"`
	Payload		interface{}	`json:"payload,omitempty"`
}

func (msg MoveMessage) GetGameId() string {
		return msg.GameId
}

func (msg MoveMessage) GetPlayerId() string {
		return msg.PlayerId
}

func NewStartGame(player, game, playerId string) EventMessage {
	return EventMessage{
		Type:			EventStartGame,
		PlayerId: playerId,
		GameId:		game,
		Payload:	player,
	}
}

func NewWinEvent(playerId string) EventMessage {
	return EventMessage{
		Type: EventWin,
		PlayerId: playerId,
	}
}

func NewLoseEvent(playerId string) EventMessage {
	return EventMessage{
		Type: EventLose,
		PlayerId: playerId,
	}
}

func NewDrawEvent(playerId string) EventMessage {
	return EventMessage{
		Type: EventDraw,
		PlayerId: playerId,
	}
}


