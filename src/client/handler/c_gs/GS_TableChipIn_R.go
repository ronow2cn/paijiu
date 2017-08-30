package c_gs

import (
	"client/msg"
)

func GS_TableChipIn_R(message msg.Message, ctx interface{}) {
	req := message.(*msg.GS_TableChipIn_R)
	log.Info("Res:", req.ErrorCode)
}
