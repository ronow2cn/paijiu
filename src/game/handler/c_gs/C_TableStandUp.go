package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableStandUp(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableStandUp)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableStandUp_R{}
	res.ErrorCode = func() int {
		er := room.Room.StandUpTable(plr.GetId(), req.Id, req.Pos)
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
