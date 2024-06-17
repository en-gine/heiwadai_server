package main

import (
	"server/infrastructure/booking/util"
	"testing"
)

func TestHiraToHalfKana(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"あいうえお", "ｱｲｳｴｵ"},
		{"かきくけこ", "ｶｷｸｹｺ"},
		{"さしすせそ", "ｻｼｽｾｿ"},
		{"たちつてと", "ﾀﾁﾂﾃﾄ"},
		{"なにぬねの", "ﾅﾆﾇﾈﾉ"},
		{"はひふへほ", "ﾊﾋﾌﾍﾎ"},
		{"まみむめも", "ﾏﾐﾑﾒﾓ"},
		{"やゆよ", "ﾔﾕﾖ"},
		{"らりるれろ", "ﾗﾘﾙﾚﾛ"},
		{"わをん", "ﾜｦﾝ"},
		{"がぎぐげご", "ｶﾞｷﾞｸﾞｹﾞｺﾞ"},
		{"ざじずぜぞ", "ｻﾞｼﾞｽﾞｾﾞｿﾞ"},
		{"だぢづでど", "ﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞ"},
		{"ばびぶべぼ", "ﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞ"},
		{"ぱぴぷぺぽ", "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ"},
		{"ゔっゃゅょ", "ｳﾞｯｬｭｮ"},
		{"こんにちは、せかい！", "ｺﾝﾆﾁﾊ､ｾｶｲ!"},
	}
	for _, test := range tests {
		result := util.HiraToHalfKana(test.input)
		if result != test.output {
			t.Errorf("HiraToHalfKana(%q) = %q; want %q", test.input, result, test.output)
		}
	}
}
