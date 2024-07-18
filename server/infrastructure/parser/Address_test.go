package parser_test

import (
	"reflect"
	"server/infrastructure/parser"
	"testing"
)

func TestParseAddress(t *testing.T) {
	testCases := []struct {
		input    string
		expected [3]string
	}{
		{"東京都千代田区千代田1-1-1", [3]string{"東京都", "千代田区", "千代田1-1-1"}},
		{"福岡県福岡市中央区舞鶴1-5-6", [3]string{"福岡県", "福岡市中央区", "舞鶴1-5-6"}},
		{"北海道札幌市中央区北1条西2丁目", [3]string{"北海道", "札幌市中央区", "北1条西2丁目"}},
		{"京都府京都市上京区今出川通烏丸東入", [3]string{"京都", "府京都市上京区", "今出川通烏丸東入"}},
		{"福岡県うきは市浮羽町浅田555コーポ135", [3]string{"福岡県", "うきは市", "浮羽町浅田555コーポ135"}},
	}

	for _, tc := range testCases {
		parsed := parser.ParseAddress(tc.input)
		if !reflect.DeepEqual(parsed, tc.expected) {
			t.Errorf("For input %s, expected %v, but got %v", tc.input, tc.expected, parsed)
		}
	}
}
