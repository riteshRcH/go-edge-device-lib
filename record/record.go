package record

import (
	pb "github.com/riteshRcH/go-edge-device-lib/record/pb"
)

// MakePutRecord creates a dht record for the given key/value pair
func MakePutRecord(key string, value []byte) *pb.Record {
	record := new(pb.Record)
	record.Key = []byte(key)
	record.Value = value
	return record
}
