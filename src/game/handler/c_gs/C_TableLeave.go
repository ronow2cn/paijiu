package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableLeave(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableLeave)

	plr := ctx.(*app.Player)
	res := &msg.GS_TableLeave_R{}

	res.ErrorCode = func() int {
		er := room.Room.LeaveTable(plr.GetId(), req.Id)
		if er != Err.OK {
			return er
		}

		table := room.Room.GetTableById(req.Id)
		if table == nil {
			return Err.Table_NotExist
		}

		table.NotifyTableInfoToAll()

		return Err.OK
	}()

	plr.SendMsg(res)
}
