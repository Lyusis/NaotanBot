package bilibili

import (
	"reflect"
	"testing"

	"github.com/Lyusis/NaotanBot/scheduler/engine"
	"github.com/Lyusis/NaotanBot/scheduler/fetcher"
)

func TestSendLiveUrl(t *testing.T) {
	type args struct {
		contents []byte
	}
	content, _ := fetcher.GetFetcher("https://api.live.bilibili.com/room/v1/Room/room_init?id=33942")
	arg := args{}
	arg.contents = content
	tests := []struct {
		name string
		args args
		want engine.ResultItems
	}{
		{"1", arg, engine.ResultItems{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendLiveUrl(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendLiveUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
