package c_gs

import (
	"client/msg"
)

func GS_TableStandUp_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableStandUp_R)
	log.Info("Res:", req.ErrorCode)
}
