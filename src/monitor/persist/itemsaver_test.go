package persist

import (
	"fmt"
	"model"
	"testing"
)

func TestSave(t *testing.T) {
	item := model.LiveDataResponseData{
		RoomId:      21672023,
		ShortId:     0,
		Uid:         477317922,
		NeedP2p:     0,
		IsHidden:    false,
		IsLocked:    false,
		IsPortrait:  false,
		LiveStatus:  0,
		HiddenTill:  0,
		LockTill:    0,
		Encrypted:   false,
		PwdVerified: false,
		LiveTime:    -62170012800,
		RoomShield:  1,
		IsSp:        0,
		SpecialType: 0,
	}
	fmt.Println(item)
}
