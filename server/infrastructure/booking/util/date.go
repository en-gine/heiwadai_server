package util

import (
	"time"
	_ "time/tzdata"
)

var JST, _ = time.LoadLocation("Asia/Tokyo")

type YYYYMMDD string

func (s YYYYMMDD) ToDate() (time.Time, error) {
	return time.Parse("2006-01-02", s.ToString())
}

func (s YYYYMMDD) ToString() string {
	return string(s)
}

func DateToYYYYMMDD(date time.Time) YYYYMMDD {
	// JST固定でフォーマット（main.goのtime.Local設定に依存させない）
	return YYYYMMDD(date.In(JST).Format("2006-01-02"))
}

func StringToYYYYMMDD(dateStr string) (*YYYYMMDD, error) {
	// YYYY-MM-DDの形式の文字列を受け取ってdate型に変換する
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}
	formated := DateToYYYYMMDD(date)
	return &formated, nil
}

func YYYYMMDDToDate(str YYYYMMDD) (time.Time, error) {
	return time.Parse("2006-01-02", str.ToString())
}

// --------------------------------------------------------

type StrDate string // 'YYYY-MM-DD形式'

func (s StrDate) ToDate() (time.Time, error) {
	return time.Parse("2006-01-02", s.ToString())
}

func (s StrDate) ToString() string {
	return string(s)
}

func DateToStrDate(date time.Time) StrDate {
	// JST固定でフォーマット（main.goのtime.Local設定に依存させない）
	return StrDate(date.In(JST).Format("2006-01-02"))
}
