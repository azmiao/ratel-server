package consts

import (
	"fmt"
	"github.com/ratel-online/core/consts"
	"time"
)

type StateID int

const (
	_ StateID = iota
	StateWelcome
	StateHome
	StateJoin
	StateNew
	StateSetting
	StateWaiting
	StateClassics
)

const (
	IS = consts.IS

	MinPlayers = 3
	MaxPlayers = 6

	RoomStateWaiting = 1
	RoomStateRunning = 2

	GameTypeClassic = 1
	GameTypeLaiZi   = 2
	GameTypeRunFast = 3

	ClassicsRobTimeout  = 10 * time.Second
	ClassicsPlayTimeout = 20 * time.Second
)

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}

func NewErr(msg string) Error {
	return Error{Msg: msg}
}

var (
	ErrorsTimeout                = NewErr("Timeout. ")
	ErrorsExist                  = NewErr("Exist. ")
	ErrorsInputInvalid           = NewErr("Input invalid. ")
	ErrorsAuthFail               = NewErr("Auth fail. ")
	ErrorsRoomInvalid            = NewErr("Room invalid. ")
	ErrorsPlayersInvalid         = NewErr(fmt.Sprintf("Invalid players, must %d-%d", MinPlayers, MaxPlayers))
	ErrorsGameTypeInvalid        = NewErr("Game type invalid. ")
	ErrorsRoomPlayersIsFull      = NewErr("Room players is fill. ")
	ErrorsJoinFailForRoomRunning = NewErr("Join fail, room is running. ")
	ErrorsGamePlayersInvalid     = NewErr("Game players invalid. ")

	GameTypes = map[int]string{
		GameTypeClassic: "Classic",
		GameTypeLaiZi:   "LaiZi",
		GameTypeRunFast: "RunFast",
	}
	GameTypesIds = []int{GameTypeClassic, GameTypeLaiZi, GameTypeRunFast}
	RoomStates   = map[int]string{
		RoomStateWaiting: "Waiting",
		RoomStateRunning: "Running",
	}
)
