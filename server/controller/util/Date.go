package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeStampPtrToTimePtr(time *timestamppb.Timestamp) *time.Time {
	if time == nil {
		return nil
	}
	t := time.AsTime()
	return &t
}

func TimePtrToTimeStampPtr(time *time.Time) *timestamppb.Timestamp {
	if time == nil {
		return nil
	}
	t := timestamppb.New(*time)
	return t
}
