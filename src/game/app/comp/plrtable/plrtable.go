package plrtable

import ()

type PlrTable struct {
	Id  int32 `bson:"id"` //加入桌子id
	plr IPlayer
}

// ============================================================================

func NewPlrTable() *PlrTable {
	return &PlrTable{}
}

func (self *PlrTable) Init(plr IPlayer) {
	self.plr = plr
}
