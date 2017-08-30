package handler

import (
    "client/handler/c_gs"
    "client/handler/c_gw"
    "client/msg"
)

func Init() {
    msg.Handler(1001, c_gw.GW_Login_R)
    msg.Handler(3000, c_gs.GS_LoginError)
    msg.Handler(3001, c_gs.GS_UserInfo)
    msg.Handler(4001, c_gs.GS_TableCreate_R)
    msg.Handler(4003, c_gs.GS_TableEnter_R)
    msg.Handler(4004, c_gs.GS_TableInfoNotify)
    msg.Handler(4007, c_gs.GS_TableLeave_R)
    msg.Handler(4009, c_gs.GS_TableSeatDown_R)
    msg.Handler(4011, c_gs.GS_TableStandUp_R)
    msg.Handler(4013, c_gs.GS_TableDice_R)
    msg.Handler(4015, c_gs.GS_TableDisMiss_R)
    msg.Handler(4017, c_gs.GS_TableChipIn_R)
    msg.Handler(4019, c_gs.GS_TableBeginFight_R)
    msg.Handler(4021, c_gs.GS_TableNextPlay_R)
    msg.Handler(101, c_gs.GS_Test_R)
}
