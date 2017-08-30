package handler

import (
    "game/handler/c_gs"
    "game/handler/gw_gs"
    "game/msg"
)

func Init() {
    msg.Handler(2000, gw_gs.GW_RegisterGate)
    msg.Handler(2002, gw_gs.GW_UserOnline)
    msg.Handler(2003, gw_gs.GW_LogoutPlayer)
    msg.Handler(4000, c_gs.C_TableCreate)
    msg.Handler(4002, c_gs.C_TableEnter)
    msg.Handler(4006, c_gs.C_TableLeave)
    msg.Handler(4008, c_gs.C_TableSeatDown)
    msg.Handler(4010, c_gs.C_TableStandUp)
    msg.Handler(4012, c_gs.C_TableDice)
    msg.Handler(4014, c_gs.C_TableDisMiss)
    msg.Handler(4016, c_gs.C_TableChipIn)
    msg.Handler(4018, c_gs.C_TableBeginFight)
    msg.Handler(4020, c_gs.C_TableNextPlay)
    msg.Handler(100, c_gs.C_Test)
}
