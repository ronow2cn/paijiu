package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableBeginFight(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableBeginFight)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableBeginFight_R{}
	res.ErrorCode = func() int {
		er := room.Room.BeginFight(plr.GetId(), req.Id)
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
