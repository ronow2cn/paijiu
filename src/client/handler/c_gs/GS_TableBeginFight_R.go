package c_gs

import (
	"client/msg"
)

func GS_TableBeginFight_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableBeginFight_R)
	log.Info("Res:", req.ErrorCode)
}
