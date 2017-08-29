package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableCreate(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableCreate)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableCreate_R{}
	id := int32(0)
	res.ErrorCode = func() int {
		if req.Score <= 0 && req.Score >= 1000000 {
			return Err.Table_ScoreError
		}

		tableid, ts := plr.GetPlrTable().GetTableId(), plr.GetPlrTable().GetTCreateTime()
		if tableid != 0 {
			return Err.Table_PlayerHaveTable
		}

		check := room.Room.CheckPlrTableId(tableid, ts)
		if check != 0 {
			return Err.Table_IsInOtherTable
		}

		id = room.Room.CreateTable(plr.GetId(), req.Score)

		return Err.OK
	}()

	plr.SendMsg(res)

	if res.ErrorCode != Err.OK {
		return
	}

	table := room.Room.GetTableById(id)
	if table == nil {
		return
	}

	table.NotifyTableInfoToAll()
}
