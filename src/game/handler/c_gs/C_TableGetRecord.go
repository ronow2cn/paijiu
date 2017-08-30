package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableGetRecord(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableGetRecord)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableGetRecord_R{}
	res.ErrorCode = func() int {
		table := room.Room.GetTableById(req.Id)
		if table == nil {
			return Err.Table_NotExist
		}

		if !table.IsInTable(plr.GetId()) {
			return Err.Table_PlayerNotInTable
		}

		res.Records = table.GetRecordToMsg()

		return Err.OK
	}()

	plr.SendMsg(res)
}
