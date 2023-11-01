package util

import "github.com/ktnyt/go-moji"

type Kana string

func (kana Kana) ToString() string {
	// 全角カタカナを半角カタカナに変換
	return string(kana)
}

// ToNarrowKana 全角カタカナを半角カタカナに変換
func (kana Kana) ToNarrowKana() string {
	return moji.Convert(kana.ToString(), moji.ZK, moji.HK)
}
