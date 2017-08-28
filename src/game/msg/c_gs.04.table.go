package msg

// msgid ragne for C <-> GS: [4000, 4100)

type PlayerInfo struct {
	Name  string
	Head  string
	Score int32 //当前拥有筹码
}

//牌局下注情况
type Chip struct {
	Bets map[string]int32
}

//当前牌局信息
type CurPlay struct {
	Chips map[int32]*Chip
}

//桌面信息
type TableData struct {
	Id       int32
	Plrs     map[string]*PlayerInfo
	Pos      map[int32]string
	PlayIdx  int32
	DiceNum  int32
	CurPlay  *CurPlay
	CreateTs int64
}

type C_TableCreate struct { // msgid: 4000
	Score int32
}

type GS_TableCreate_R struct { // msgid:4001
	ErrorCode int
	TableData *TableData `msgpack:",omitempty"`
}

type C_TableEnter struct { // msgid: 4002
	Id int32
}

type GS_TableEnter struct { // msgid: 4003
	ErrorCode int
	TableData *TableData `msgpack:",omitempty"`
}

type GS_TableInfoNotify struct { // msgid: 4004
	TableData *TableData
}

type C_TableLeave struct { // msgid: 4006
	Id int32
}

type GS_TableLeave_R struct { // msgid: 4007
	ErrorCode int
}

type C_TableSeatDown struct { // msgid: 4008
	Id int32
}

type GS_TableSeatDown_R struct { // msgid:4009
	ErrorCode int
}

type C_TableStandUp struct { // msgid: 4010
	Id int32
}

type GS_TableStandUp struct { // msgid: 4011
	ErrorCode int
}

type C_TableDice struct { // msgid: 4012
	Id int32
}

type GS_TableDice struct { //msgid: 4013
	ErrorCode int
}
