package c_gs

import (
	"game/app"
	"game/app/modules/room"
	"game/msg"
	"proto/errorcode"
)

func C_TableCreate(message msg.Message, ctx interface{}) {
	req := message.(*msg.C_TableCreate)
	plr := ctx.(*app.Player)

	res := &msg.GS_TableCreate_R{}
	res.ErrorCode = func() int {
		if req.Score <= 0 && req.Score >= 1000000 {
			return Err.Table_ScoreError
		}

		id := room.Room.CreateTable(plr.GetId(), req.Score)

		table := room.Room.GetTableById(id)
		if table == nil {
			return Err.Table_NotExist
		}

		res.TableData = table.ToMsg()

		return Err.OK
	}()

	plr.SendMsg(res)
}
