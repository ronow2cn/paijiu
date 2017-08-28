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

type Chip struct {
	Bets map[string]int32 `bson:"num"` //[玩家]玩家下注数量
}

type Play struct {
	Chips chips `bson:"chips"` //下注情况
}

type chips map[int32]*Chip //[位置]位置上筹码
type pos map[int32]string  //[位置]位置上玩家

type Table struct {
	Id       int32                  `bson:"id"`
	Plrs     map[string]*playerInfo `bson:"plrs"`
	Pos      pos                    `bson:"Pos"`
	PlayIdx  int32                  `bson:"play_idx"`
	DiceNum  int32                  `bson:"dice_num"`
	CurPlay  *Play                  `bson:"cur_play"`
	CreateTs time.Time              `bson:"create_ts"`
}

// ============================================================================
// marshalling

func (self chips) GetBSON() (interface{}, error) {
	type chip_t struct {
		Id  int32
		Pos *Chip
	}
	var arr []*chip_t

	for id, val := range self {
		arr = append(arr, &chip_t{id, val})
	}

	return arr, nil
}

func (self *chips) SetBSON(raw bson.Raw) error {
	type chip_t struct {
		Id  int32
		Pos *Chip
	}
	var arr []*chip_t

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(chips)
	for _, v := range arr {
		(*self)[v.Id] = v.Pos
	}

	return nil
}

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
		PlayIdx:  1,
	}
}

func (self *Table) Init(id int32, plrid string, score int32) int {
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
	if ok {
		return Err.Table_NotInTable
	}

	if self.IsBanker(plrid) {
		return Err.Table_IsBanker
	}

	self.DelPos(plrid)

	if self.Plrs[plrid].Score == 0 {
		delete(self.Plrs, plrid)
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
			Chips: make(map[int32]*msg.Chip),
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

	for k, v := range self.CurPlay.Chips {
		cp := &msg.Chip{
			Bets: make(map[string]int32),
		}

		for kk, vv := range v.Bets {
			cp.Bets[kk] = vv
		}

		ret.CurPlay.Chips[k] = cp
	}

	return ret
}
