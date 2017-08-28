package msg

type PlayerInfo struct {
	Name  string
	Head  string
	Score int32 //当前拥有筹码
}

type Chip struct {
	Bets map[string]int32
}

type CurPlay struct {
	Chips map[int32]*Chip
}

type TableData struct {
	Id       int32
	Plrs     map[string]*PlayerInfo
	Pos      map[int32]string
	PlayIdx  int32
	DiceNum  int32
	CurPlay  *CurPlay
	CreateTs int64
}
