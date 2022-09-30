package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	tests := []struct {
		name    string
		cbxType string
		cash    float32
		want    float32
	}{
		{
			name:    "TestStrategy 正常收費",
			cbxType: "正常收費",
			cash:    10000,
			want:    10000,
		},
		{
			name:    "TestStrategy 滿300送100",
			cbxType: "滿300送100",
			cash:    10000,
			want:    6700,
		},
		{
			name:    "TestStrategy 打8折",
			cbxType: "打8折",
			cash:    10000,
			want:    8000,
		},
	}

	for _, tt := range tests {
		var context CashContext
		err := context.CashContext(tt.cbxType)
		if err != nil {
			t.Error(err)
		} else {
			if val := context.GetResult(tt.cash); val != tt.want {
				t.Error("Add Error", val)
			} else {
				// 測試所有優惠都上
				// context.CashContext("滿300送100")
				// val = context.GetResult(val)
				// fmt.Println(val)
				// context.CashContext("打8折")
				// val = context.GetResult(val)
				// fmt.Println(val)
				// fmt.Println("------")
			}
		}
	}
}
