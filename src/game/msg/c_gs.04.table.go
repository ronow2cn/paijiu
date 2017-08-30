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

type Card struct {
	T int32 //类型
	N int32 //数值
}

//当前牌局信息
type CurPlay struct {
	Id      int32
	Chips   map[int32]*Chip
	PosCard map[int32][]*Card
}

type RecOne struct {
	Id      int32
	PosCard map[int32][]*Card
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
}

type C_TableEnter struct { // msgid: 4002
	Id int32
}

type GS_TableEnter_R struct { // msgid: 4003
	ErrorCode int
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
	Id  int32
	Pos int32
}

type GS_TableSeatDown_R struct { // msgid:4009
	ErrorCode int
}

type C_TableStandUp struct { // msgid: 4010
	Id  int32
	Pos int32
}

type GS_TableStandUp_R struct { // msgid: 4011
	ErrorCode int
}

type C_TableDice struct { // msgid: 4012
	Id int32
}

type GS_TableDice_R struct { //msgid: 4013
	ErrorCode int
	DiceNum   int32
}

type C_TableDisMiss struct { //msgid: 4014
	Id int32
}

type GS_TableDisMiss_R struct { //msgid: 4015
	ErrorCode int
	Id        int32
}

type C_TableChipIn struct { //msgid: 4016
	Id    int32
	Pos   int32
	Score int32
}

type GS_TableChipIn_R struct { // msgid: 4017
	ErrorCode int
}

type C_TableBeginFight struct { //msgid: 4018
	Id int32
}

type GS_TableBeginFight_R struct { //msgid: 4019
	ErrorCode int
}

type C_TableNextPlay struct { // msgid:4020
	Id int32
}

type GS_TableNextPlay_R struct { // msgid:4021
	ErrorCode int
}

type C_TableGetRecord struct { // msgid: 4022
	Id int32
}

type GS_TableGetRecord_R struct { //msgid: 4023
	ErrorCode int
	Records   []*RecOne `msgpack:",omitempty"`
}
