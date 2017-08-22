package app

import (
	"comm/logger"
	"gopkg.in/mgo.v2/bson"
)

// ============================================================================

var log = logger.DefaultLogger

// ============================================================================

func sid2gateid(sid uint64) int32 {
	return int32(sid >> 42)
}

// ============================================================================

func CloneBsonObject(obj interface{}) interface{} {
	buf, err := bson.Marshal(obj)
	if err != nil {
		log.Error("marshal object data failed:", err)
		return nil
	}

	out := make(map[string]interface{})
	err = bson.Unmarshal(buf, &out)
	if err != nil {
		log.Error("unmarshal object data failed:", err)
		return nil
	}

	return out
}
