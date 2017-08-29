package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableEnter(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableEnter)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableEnter_R{}
	res.ErrorCode = func() int {
		check := room.Room.CheckPlrTableId(plr.GetPlrTable().GetTableId(), plr.GetPlrTable().GetTCreateTime())
		if check != 0 && check != req.Id {
			return Err.Table_IsInOtherTable
		}

		er := room.Room.EnterTable(plr.GetId(), req.Id)
		if er != Err.OK {
			return er
		}

		return Err.OK
	}()

	plr.SendMsg(res)

	if res.ErrorCode != Err.OK {
		return
	}

	table := room.Room.GetTableById(req.Id)
	if table == nil {
		return
	}

	table.NotifyTableInfoToAll()
}
