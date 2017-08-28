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
		er := room.Room.EnterTable(plr.GetId(), req.Id)
		if er != Err.OK {
			return er
		}

		table := room.Room.GetTableById(req.Id)
		if table == nil {
			return Err.Table_NotExist
		}

		res.TableData = table.ToMsg()

		return Err.OK

	}()

	plr.SendMsg(res)
}
