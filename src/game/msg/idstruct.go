package msg

var MsgCreators = map[uint32]func() Message{
    2000: func() Message {
        return &GW_RegisterGate{}
    },
    2001: func() Message {
        return &GS_RegisterGate_R{}
    },
    2002: func() Message {
        return &GW_UserOnline{}
    },
    2003: func() Message {
        return &GW_LogoutPlayer{}
    },
    2004: func() Message {
        return &GS_Kick{}
    },
    3000: func() Message {
        return &GS_LoginError{}
    },
    3001: func() Message {
        return &GS_UserInfo{}
    },
    4000: func() Message {
        return &C_TableCreate{}
    },
    4001: func() Message {
        return &GS_TableCreate_R{}
    },
    4002: func() Message {
        return &C_TableEnter{}
    },
    4003: func() Message {
        return &GS_TableEnter_R{}
    },
    4004: func() Message {
        return &GS_TableInfoNotify{}
    },
    4006: func() Message {
        return &C_TableLeave{}
    },
    4007: func() Message {
        return &GS_TableLeave_R{}
    },
    4008: func() Message {
        return &C_TableSeatDown{}
    },
    4009: func() Message {
        return &GS_TableSeatDown_R{}
    },
    4010: func() Message {
        return &C_TableStandUp{}
    },
    4011: func() Message {
        return &GS_TableStandUp_R{}
    },
    4012: func() Message {
        return &C_TableDice{}
    },
    4013: func() Message {
        return &GS_TableDice_R{}
    },
    4014: func() Message {
        return &C_TableDisMiss{}
    },
    4015: func() Message {
        return &GS_TableDisMiss_R{}
    },
    4016: func() Message {
        return &C_TableChipIn{}
    },
    4017: func() Message {
        return &GS_TableChipIn_R{}
    },
    4018: func() Message {
        return &C_TableBeginFight{}
    },
    4019: func() Message {
        return &GS_TableBeginFight_R{}
    },
    4020: func() Message {
        return &C_TableNextPlay{}
    },
    4021: func() Message {
        return &GS_TableNextPlay_R{}
    },
    4022: func() Message {
        return &C_TableGetRecord{}
    },
    4023: func() Message {
        return &GS_TableGetRecord_R{}
    },
    100: func() Message {
        return &C_Test{}
    },
    101: func() Message {
        return &GS_Test_R{}
    },
}

func (self *GW_RegisterGate) MsgId() uint32 {
    return 2000
}

func (self *GS_RegisterGate_R) MsgId() uint32 {
    return 2001
}

func (self *GW_UserOnline) MsgId() uint32 {
    return 2002
}

func (self *GW_LogoutPlayer) MsgId() uint32 {
    return 2003
}

func (self *GS_Kick) MsgId() uint32 {
    return 2004
}

func (self *GS_LoginError) MsgId() uint32 {
    return 3000
}

func (self *GS_UserInfo) MsgId() uint32 {
    return 3001
}

func (self *C_TableCreate) MsgId() uint32 {
    return 4000
}

func (self *GS_TableCreate_R) MsgId() uint32 {
    return 4001
}

func (self *C_TableEnter) MsgId() uint32 {
    return 4002
}

func (self *GS_TableEnter_R) MsgId() uint32 {
    return 4003
}

func (self *GS_TableInfoNotify) MsgId() uint32 {
    return 4004
}

func (self *C_TableLeave) MsgId() uint32 {
    return 4006
}

func (self *GS_TableLeave_R) MsgId() uint32 {
    return 4007
}

func (self *C_TableSeatDown) MsgId() uint32 {
    return 4008
}

func (self *GS_TableSeatDown_R) MsgId() uint32 {
    return 4009
}

func (self *C_TableStandUp) MsgId() uint32 {
    return 4010
}

func (self *GS_TableStandUp_R) MsgId() uint32 {
    return 4011
}

func (self *C_TableDice) MsgId() uint32 {
    return 4012
}

func (self *GS_TableDice_R) MsgId() uint32 {
    return 4013
}

func (self *C_TableDisMiss) MsgId() uint32 {
    return 4014
}

func (self *GS_TableDisMiss_R) MsgId() uint32 {
    return 4015
}

func (self *C_TableChipIn) MsgId() uint32 {
    return 4016
}

func (self *GS_TableChipIn_R) MsgId() uint32 {
    return 4017
}

func (self *C_TableBeginFight) MsgId() uint32 {
    return 4018
}

func (self *GS_TableBeginFight_R) MsgId() uint32 {
    return 4019
}

func (self *C_TableNextPlay) MsgId() uint32 {
    return 4020
}

func (self *GS_TableNextPlay_R) MsgId() uint32 {
    return 4021
}

func (self *C_TableGetRecord) MsgId() uint32 {
    return 4022
}

func (self *GS_TableGetRecord_R) MsgId() uint32 {
    return 4023
}

func (self *C_Test) MsgId() uint32 {
    return 100
}

func (self *GS_Test_R) MsgId() uint32 {
    return 101
}
