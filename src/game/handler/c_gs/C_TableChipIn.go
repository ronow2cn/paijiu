package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableChipIn(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableChipIn)
	plr := ctx.(*app.Player)
	log.Info("C_TableChipIn", plr.GetId(), req.Id, req.Pos, req.Score)
	res := &msg.GS_TableChipIn_R{}
	res.ErrorCode = func() int {
		er := room.Room.ChipIn(plr.GetId(), req.Id, req.Pos, req.Score)
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
