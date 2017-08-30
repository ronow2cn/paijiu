package room

import (
	"game/app/modules/table"
	"game/msg"
	"math/rand"
	"proto/errorcode"
	"time"
)

// ============================================================================

var (
	RANDTABLEID = rand.New(rand.NewSource(time.Now().Unix()))
)

var Room = &room{
	Table: make(map[int32]*table.Table),
}

// ============================================================================

type room struct {
	Table map[int32]*table.Table
}

// ============================================================================
func Init() {

}

func (self *room) GetRoomCnt() int32 {
	return int32(len(self.Table))
}

func (self *room) genTableId() int32 {
	for i := 0; i < 50; i++ {
		id := RANDTABLEID.Int31n(899999) + 100000

		if _, ok := self.Table[id]; !ok {
			return id
		}
	}

	for i := 100000; i <= 999999; i++ {
		if _, ok := self.Table[int32(i)]; !ok {
			return int32(i)
		}
	}

	return 0
}

func (self *room) GetTableById(id int32) *table.Table {
	table, ok := self.Table[id]
	if !ok {
		return nil
	}

	return table
}

//检查玩家身上的tableid 对应的桌子是否还在
func (self *room) CheckPlrTableId(id int32, ts time.Time) int32 {
	t, ok := self.Table[id]
	if ok {
		if ts.Equal(t.CreateTs) {
			return id
		}
	}

	return 0
}

// ============================================================================
//开桌人的id，开房总分
func (self *room) CreateTable(plrid string, score int32) int32 {
	id := self.genTableId()
	if id == 0 {
		return id
	}

	self.Table[id] = table.NewTable()
	self.Table[id].InitNewTable(id, plrid, score)

	return id
}

//解散桌子，只有桌子创建者才能解散
func (self *room) DisMissTable(plrid string, id int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	if !table.IsBanker(plrid) {
		return Err.Table_IsNotBanker
	}

	table.SetPlrsTableId(0, time.Now())

	table.BroadcastMsg(&msg.GS_TableDisMiss_R{
		ErrorCode: Err.OK,
		Id:        id,
	})

	delete(self.Table, id)

	return Err.OK
}

//进入桌子
func (self *room) EnterTable(plrid string, id int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	er := table.Enter(plrid)
	if er != Err.OK {
		return er
	}

	return Err.OK
}

//离开桌子
func (self *room) LeaveTable(plrid string, id int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	er := table.Leave(plrid)
	if er != Err.OK {
		return er
	}

	return Err.OK
}

//坐上桌位
func (self *room) SeatDownTable(plrid string, id int32, pos int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	er := table.SeatDown(plrid, pos)
	if er != Err.OK {
		return er
	}

	return Err.OK
}

//站起来
func (self *room) StandUpTable(plrid string, id int32, pos int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	er := table.StandUp(plrid, pos)
	if er != Err.OK {
		return er
	}

	return Err.OK
}

//掷骰子
func (self *room) DiceTable(plrid string, id int32) (int32, int) {
	table, ok := self.Table[id]
	if !ok {
		return 0, Err.Table_NotExist
	}

	if !table.IsBanker(plrid) {
		return 0, Err.Table_IsNotBanker
	}

	d := table.Dice()

	return d, Err.OK
}

func (self *room) ChipIn(plrid string, id int32, pos int32, score int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	if !table.IsCurPlay() {
		return Err.Table_NotCurPlay
	}

	if table.IsBanker(plrid) {
		return Err.Table_IsBanker
	}

	if table.GetBankerScore() < score {
		return Err.Table_BankerScoreNotEnough
	}

	er := table.ChipIn(plrid, pos, score)

	return er
}

func (self *room) BeginFight(plrid string, id int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	if !table.IsCurPlay() {
		return Err.Table_NotCurPlay
	}

	if !table.IsBanker(plrid) {
		return Err.Table_IsNotBanker
	}

	er := table.BeginFight(plrid)

	return er
}

func (self *room) NextPlay(plrid string, id int32) int {
	table, ok := self.Table[id]
	if !ok {
		return Err.Table_NotExist
	}

	if table.IsCurPlay() {
		return Err.Table_CurPlayIng
	}

	er := table.NextPlay(plrid)

	return er
}
