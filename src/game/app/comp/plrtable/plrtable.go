package plrtable

import (
	"time"
)

type PlrTable struct {
	Id        int32     `bson:"id"` //加入桌子id
	TCreateTs time.Time `bson:"tcreate_ts"`
	plr       IPlayer
}

// ============================================================================

func NewPlrTable() *PlrTable {
	return &PlrTable{}
}

func (self *PlrTable) Init(plr IPlayer) {
	self.plr = plr
}

func (self *PlrTable) Set(id int32, t time.Time) {
	self.Id = id
	self.TCreateTs = t
}

func (self *PlrTable) GetTableId() int32 {
	return self.Id
}

func (self *PlrTable) GetTCreateTime() time.Time {
	return self.TCreateTs
}
