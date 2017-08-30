package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableNextPlay(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableNextPlay)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableNextPlay_R{}
	res.ErrorCode = func() int {
		er := room.Room.NextPlay(plr.GetId(), req.Id)
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
