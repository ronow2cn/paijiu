package c_gs

import (
	"client/msg"
)

func GS_TableDisMiss_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableDisMiss_R)
	log.Info("Res:", req.ErrorCode, req.Id)
}
