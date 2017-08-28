package c_gs

import (
	"client/msg"
)

func GS_TableSeatDown_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableSeatDown_R)
	log.Info("Res:", req.ErrorCode)
}
