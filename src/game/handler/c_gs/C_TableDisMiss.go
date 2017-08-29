package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableDisMiss(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableDisMiss)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableDisMiss_R{}
	res.ErrorCode = func() int {
		er := room.Room.DisMissTable(plr.GetId(), req.Id)
		if er != Err.OK {
			return er
		}

		res.Id = req.Id
		return Err.OK
	}()

	if res.ErrorCode != Err.OK {
		plr.SendMsg(res)
	}
}
