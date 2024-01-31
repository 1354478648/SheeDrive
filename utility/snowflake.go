package utility

import (
	"SheeDrive/internal/consts"
	"time"
)

const (
	timestampBits  = 41
	machineIDBits  = 10
	sequenceBits   = 12
	maxMachineID   = -1 ^ (-1 << machineIDBits)
	maxSequenceNum = -1 ^ (-1 << sequenceBits)
)

type snowflake struct {
	timestamp   int64
	machineID   int64
	sequenceNum int64
}

func newSnowflake(machineID int64) *snowflake {
	if machineID < 0 || machineID > maxMachineID {
		panic("Invalid machine ID")
	}

	return &snowflake{
		timestamp:   time.Now().UnixNano() / 1e6,
		machineID:   machineID,
		sequenceNum: 0,
	}
}

func (s *snowflake) generateID() int64 {
	currentTimestamp := time.Now().UnixNano() / 1e6

	if currentTimestamp == s.timestamp {
		s.sequenceNum = (s.sequenceNum + 1) & maxSequenceNum
		if s.sequenceNum == 0 {
			currentTimestamp = s.waitNextMillis()
		}
	} else {
		s.sequenceNum = 0
	}

	s.timestamp = currentTimestamp

	id := (currentTimestamp << (machineIDBits + sequenceBits)) |
		(s.machineID << sequenceBits) |
		s.sequenceNum

	return id
}

func (s *snowflake) waitNextMillis() int64 {
	currentTimestamp := time.Now().UnixNano() / 1e6
	for currentTimestamp <= s.timestamp {
		currentTimestamp = time.Now().UnixNano() / 1e6
	}
	return currentTimestamp
}

func GenSnowFlakeId() int64 {
	return newSnowflake(int64(consts.MachineID)).generateID()
}
