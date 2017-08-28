package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableSeatDown(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableSeatDown)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableSeatDown_R{}
	res.ErrorCode = func() int {
		er := room.Room.SeatDownTable(plr.GetId(), req.Id, req.Pos)
		if er != Err.OK {
			return er
		}

		return Err.OK
	}()

	plr.SendMsg(res)

	table := room.Room.GetTableById(req.Id)
	if table == nil {
		return
	}

	table.NotifyTableInfoToAll()
}
