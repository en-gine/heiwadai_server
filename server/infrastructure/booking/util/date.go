package util

import "time"

type YYYYMMDD string

func (s YYYYMMDD) ToDate() (time.Time, error) {
	return time.Parse("2006-01-02", s.ToString())
}

func (s YYYYMMDD) ToString() string {
	return string(s)
}

func DateToYYYYMMDD(date time.Time) YYYYMMDD {
	// date型を受け取ってYYYY-MM-DDの形式の文字列に変換する
	// time.Localを使用してシステムのタイムゾーン（JST）でフォーマットする
	return YYYYMMDD(date.In(time.Local).Format("2006-01-02"))
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
	// date型を受け取ってYYYY-MM-DDの形式の文字列に変換する
	// time.Localを使用してシステムのタイムゾーン（JST）でフォーマットする
	return StrDate(date.In(time.Local).Format("2006-01-02"))
}
