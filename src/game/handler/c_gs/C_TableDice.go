package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableDice(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableDice)
	plr := ctx.(*app.Player)
	res := &msg.GS_TableDice_R{}

	res.ErrorCode = func() int {
		n, er := room.Room.DiceTable(req.Id)
		if er != Err.OK {
			return er
		}

		res.DiceNum = n

		return Err.OK
	}()

	if res.ErrorCode != Err.OK {
		plr.SendMsg(res)
		return
	}

	table := room.Room.GetTableById(req.Id)
	if table == nil {
		return
	}

	table.BroadcastMsg(res)
}
