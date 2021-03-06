package table

import (
	"game/app"
	"game/app/gconst"
	"game/msg"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"proto/errorcode"
	"time"
)

// ============================================================================
var (
	RANDTABLEDICE = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	CTABLEMAXPLAYER = 16
)

// ============================================================================

type playerInfo struct {
	Name  string
	Head  string
	Score int32 //当前拥有筹码
}

type pos map[int32]string //[位置]位置上玩家

type recOne struct {
	Id      int32
	PosCard poscard `bson:"poscard"` //位置上的牌
}
type record []*recOne

type Table struct {
	Id       int32                  `bson:"id"`
	Plrs     map[string]*playerInfo `bson:"plrs"`
	Pos      pos                    `bson:"pos"`
	PlayIdx  int32                  `bson:"play_idx"`
	DiceNum  int32                  `bson:"dice_num"`
	CurPlay  *Play                  `bson:"cur_play"`
	CreateTs time.Time              `bson:"create_ts"`
	Record   record                 `bson:"record"`
}

// ============================================================================
func (self pos) GetBSON() (interface{}, error) {
	type pos_t struct {
		Id    int32
		PlrId string
	}
	var arr []*pos_t

	for id, val := range self {
		arr = append(arr, &pos_t{id, val})
	}

	return arr, nil
}

func (self *pos) SetBSON(raw bson.Raw) error {
	type pos_t struct {
		Id    int32
		PlrId string
	}
	var arr []*pos_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(pos)
	for _, v := range arr {
		(*self)[v.Id] = v.PlrId
	}

	return nil
}

// ============================================================================

func NewTable() *Table {
	return &Table{
		Id:       0,
		Plrs:     make(map[string]*playerInfo),
		CreateTs: time.Now(),
		Pos:      make(pos),
		PlayIdx:  1,
		DiceNum:  1,
		CurPlay:  NewPlay(),
	}
}

func (self *Table) InitNewTable(id int32, plrid string, score int32) int {
	self.Id = id

	plr := app.PlayerMgr.LoadPlayer(plrid)
	if plr == nil {
		log.Error("player not found")
		return Err.Failed
	}

	self.Pos[gconst.TablePosBanker] = plrid

	self.Plrs[plrid] = &playerInfo{
		Name:  plr.GetName(),
		Head:  plr.GetHead(),
		Score: score,
	}

	self.CurPlay.CreateInit(self.PlayIdx)

	plr.GetPlrTable().Set(self.Id, self.CreateTs)

	return Err.OK
}

// ============================================================================

func (self *Table) IsBanker(plrid string) bool {
	return plrid == self.Pos[gconst.TablePosBanker]
}

func (self *Table) DelPos(plrid string) {
	for k, v := range self.Pos {
		if v == plrid {
			delete(self.Pos, k)
		}
	}
}

func (self *Table) GetBankerScore() int32 {
	plrid := self.Pos[gconst.TablePosBanker]
	return self.Plrs[plrid].Score
}

func (self *Table) IsCurPlay() bool {
	return self.PlayIdx == self.CurPlay.Id
}

//设置桌上玩家身上的tableid
func (self *Table) SetPlrsTableId(id int32, ts time.Time) {
	for plrid, _ := range self.Plrs {
		plr := app.PlayerMgr.LoadPlayer(plrid)
		if plr == nil {
			continue
		}

		plr.GetPlrTable().Set(id, ts)
	}
}

func (self *Table) IsInTable(plrid string) bool {
	_, ok := self.Plrs[plrid]
	return ok
}

// ============================================================================

func (self *Table) Enter(plrid string) int {
	if len(self.Plrs) >= CTABLEMAXPLAYER {
		return Err.Table_Full
	}

	_, ok := self.Plrs[plrid]
	if !ok {
		plr := app.PlayerMgr.LoadPlayer(plrid)
		if plr == nil {
			log.Error("player not found")
			return Err.Failed
		}

		if plr.GetPlrTable().GetTableId() == self.Id {
			if !plr.GetPlrTable().GetTCreateTime().Equal(self.CreateTs) {
				return Err.Table_IdIsOver
			}
		}

		plr.GetPlrTable().Set(self.Id, self.CreateTs)

		self.Plrs[plrid] = &playerInfo{
			Score: 0,
			Name:  plr.GetName(),
			Head:  plr.GetHead(),
		}
	}

	return Err.OK
}

func (self *Table) Leave(plrid string) int {
	_, ok := self.Plrs[plrid]
	if !ok {
		return Err.Table_NotInTable
	}

	if self.IsBanker(plrid) {
		return Err.Table_IsBanker
	}

	self.DelPos(plrid)

	if self.Plrs[plrid].Score == 0 {
		delete(self.Plrs, plrid)
	}

	plr := app.PlayerMgr.LoadPlayer(plrid)
	if plr != nil {
		plr.GetPlrTable().Set(0, time.Now())
	}

	return Err.OK
}

func (self *Table) StandUp(plrid string, pos int32) int {
	if self.IsBanker(plrid) {
		return Err.Table_IsBanker
	}

	_, ok := self.Pos[pos]
	if !ok {
		return Err.Table_ErrorPos
	}

	if self.Pos[pos] != plrid {
		return Err.Table_PosPlrError
	}

	self.DelPos(plrid)

	return Err.OK
}

func (self *Table) SeatDown(plrid string, pos int32) int {
	if self.IsBanker(plrid) {
		return Err.Table_IsBanker
	}

	_, ok := self.Pos[pos]
	if ok {
		return Err.Table_PosOccupy
	}

	self.Pos[pos] = plrid

	return Err.OK
}

func (self *Table) Dice() int32 {
	self.DiceNum = RANDTABLEDICE.Int31n(6) + RANDTABLEDICE.Int31n(6) + 2
	return self.DiceNum
}

func (self *Table) ChipIn(plrid string, pos int32, score int32) int {
	self.CurPlay.ChipIn(plrid, pos, score)

	banker, ok1 := self.Pos[gconst.TablePosBanker]
	if !ok1 {
		log.Warning("ChipIn Pos not found", banker, pos)
		return Err.Table_PlayerNotFound
	}

	_, ok2 := self.Plrs[banker]
	_, ok3 := self.Plrs[plrid]
	if !ok2 || !ok3 {
		log.Warning("ChipIn plrs not found", banker, pos)
		return Err.Table_PlayerNotFound
	}

	self.Plrs[banker].Score -= score
	self.Plrs[plrid].Score -= score

	return Err.OK
}

func (self *Table) BeginFight(plrid string) int {
	self.CurPlay.Init(self.PlayIdx)
	res1, res2, res3 := self.CurPlay.BeginFight()

	for pid, score := range self.CurPlay.Chips[gconst.TablePosPlayer1].Bets {
		if res1 == int32(1) {
			self.Plrs[pid].Score += (2 * score)
		} else {
			self.Plrs[plrid].Score += (2 * score)
		}
	}

	for pid, score := range self.CurPlay.Chips[gconst.TablePosPlayer2].Bets {
		if res2 == int32(1) {
			self.Plrs[pid].Score += (2 * score)
		} else {
			self.Plrs[plrid].Score += (2 * score)
		}
	}

	for pid, score := range self.CurPlay.Chips[gconst.TablePosPlayer3].Bets {
		if res3 == int32(1) {
			self.Plrs[pid].Score += (2 * score)
		} else {
			self.Plrs[plrid].Score += (2 * score)
		}
	}

	self.AddRecord()

	self.PlayIdx++

	return Err.OK
}

func (self *Table) NextPlay(plrid string) int {
	self.CurPlay.Reset(self.PlayIdx)

	return Err.OK
}

func (self *Table) AddRecord() {
	rec := &recOne{
		Id:      self.CurPlay.Id,
		PosCard: self.CurPlay.PosCard,
	}
	self.Record = append(self.Record, rec)
}

func (self *Table) GetRecordToMsg() []*msg.RecOne {
	return self.Record.ToMsg()
}

// ============================================================================
//广播消息
func (self *Table) BroadcastMsg(message msg.Message) {
	for id, _ := range self.Plrs {
		plr := app.PlayerMgr.LoadPlayer(id)
		if plr == nil {
			continue
		}

		plr.SendMsg(message)
	}
}

func (self *Table) ToMsg() *msg.TableData {
	ret := &msg.TableData{
		Plrs: make(map[string]*msg.PlayerInfo),
		Pos:  make(map[int32]string),
		CurPlay: &msg.CurPlay{
			Chips:   make(map[int32]*msg.Chip),
			PosCard: make(map[int32][]*msg.Card),
		},
	}

	ret.Id = self.Id

	for k, v := range self.Plrs {
		ret.Plrs[k] = &msg.PlayerInfo{
			Name:  v.Name,
			Head:  v.Head,
			Score: v.Score,
		}
	}

	for k, v := range self.Pos {
		ret.Pos[k] = v
	}

	ret.PlayIdx = self.PlayIdx
	ret.DiceNum = self.DiceNum
	ret.CreateTs = self.CreateTs.Unix()

	ret.CurPlay.Id = self.CurPlay.Id
	for k, v := range self.CurPlay.Chips {
		cp := &msg.Chip{
			Bets: make(map[string]int32),
		}

		for kk, vv := range v.Bets {
			cp.Bets[kk] = vv
		}

		ret.CurPlay.Chips[k] = cp
	}

	for k, v := range self.CurPlay.PosCard {
		if len(v) < 2 {
			log.Warning("pos card miss", k)
			continue
		}

		ret.CurPlay.PosCard[k] = append(ret.CurPlay.PosCard[k], &msg.Card{T: v[0].T, N: v[0].N})
		ret.CurPlay.PosCard[k] = append(ret.CurPlay.PosCard[k], &msg.Card{T: v[1].T, N: v[1].N})
	}

	return ret
}

func (self *Table) NotifyTableInfoToAll() {
	res := &msg.GS_TableInfoNotify{}

	res.TableData = self.ToMsg()
	self.BroadcastMsg(res)
}

func (self record) ToMsg() []*msg.RecOne {
	ret := []*msg.RecOne{}

	for _, v := range self {
		poscard := map[int32][]*msg.Card{}
		for k, vv := range v.PosCard {
			for _, vvv := range vv {
				poscard[k] = append(poscard[k], &msg.Card{T: vvv.T, N: vvv.N, H: vvv.H})
			}
		}

		ret = append(ret, &msg.RecOne{
			Id:      v.Id,
			PosCard: poscard,
		})
	}

	return ret
}

// ============================================================================
