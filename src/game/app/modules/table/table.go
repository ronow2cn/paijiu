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
	TableMaxPlayer = 16
)

// ============================================================================

type playerinfo struct {
	Pos   int32 //位置
	Score int32 //当前拥有筹码
}

type chips map[int32]*Chip //[位置]位置上筹码

type Chip struct {
	Bets map[string]int32 `bson:"num"` //[玩家]玩家下注数量
}

type Play struct {
	Chips chips `bson:"chips"` //下注情况
}

type Table struct {
	Id        int32                  `bson:"id"`
	Plrs      map[string]*playerinfo `bson:"plrs"`
	Banker    string                 `bson:"banker"`
	PlayIdx   int32                  `bson:"play_idx"`
	DiceNum   int32                  `bson:"dice_num"`
	CurPlay   *Play                  `bson:"cur_play"`
	CreatetTs time.Time              `bson:"create_ts"`
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

// ============================================================================

func NewTable() *Table {
	return &Table{
		Id:        0,
		Plrs:      make(map[string]*playerinfo),
		CreatetTs: time.Now(),
		PlayIdx:   1,
	}
}

func (self *Table) Init(id int32, plrid string, score int32) {
	self.Id = id
	self.Banker = plrid

	self.Plrs[plrid] = &playerinfo{
		Pos:   gconst.TablePosBanker,
		Score: score,
	}
}

// ============================================================================

func (self *Table) checkPos(pos int32) bool {
	return pos == gconst.TablePosBanker || pos == gconst.TablePosPlayer1 || pos == gconst.TablePosPlayer2 ||
		pos == gconst.TablePosPlayer3 || pos == gconst.TablePosPlayerWatch

}

func (self *Table) FindPosPlrId(pos int32) (string, int) {
	if !self.checkPos(pos) {
		return "", Err.Table_ErrorPos
	}

	for id, v := range self.Plrs {
		if v.Pos == pos {
			return id, Err.OK
		}
	}

	return "", Err.OK
}

func (self *Table) IsBanker(plrid string) bool {
	return plrid == self.Banker
}

// ============================================================================

func (self *Table) Enter(plrid string) int {
	if len(self.Plrs) >= TableMaxPlayer {
		return Err.Table_Full
	}

	_, ok := self.Plrs[plrid]
	if !ok {
		self.Plrs[plrid] = &playerinfo{
			Pos:   gconst.TablePosPlayerWatch,
			Score: 0,
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

	self.Plrs[plrid].Pos = gconst.TablePosPlayerWatch

	if self.Plrs[plrid].Score == 0 {
		delete(self.Plrs, plrid)
	}

	return Err.OK
}

func (self *Table) StandUp(plrid string) int {
	if self.Banker == plrid {
		return Err.Table_IsBanker
	}

	_, ok := self.Plrs[plrid]
	if ok {
		self.Plrs[plrid].Pos = gconst.TablePosPlayerWatch
	}

	return Err.OK
}

func (self *Table) SeatDown(plrid string, pos int32) int {
	if self.Banker == plrid {
		return Err.Table_IsBanker
	}

	id, err := self.FindPosPlrId(pos)
	if err != Err.OK {
		return err
	}

	if id == "" {
		_, ok := self.Plrs[plrid]
		if ok {
			self.Plrs[plrid].Pos = pos
		}
	} else {
		return Err.Table_PosOccupy
	}

	return Err.OK
}

func (self *Table) Dice() int32 {
	return RANDTABLEDICE.Int31n(6) + RANDTABLEDICE.Int31n(6) + 2
}

func (self *Table) BroadcastMsg(message msg.Message) {
	for id, _ := range self.Plrs {
		plr := app.PlayerMgr.LoadPlayer(id)
		if plr == nil {
			continue
		}

		plr.SendMsg(message)
	}
}
