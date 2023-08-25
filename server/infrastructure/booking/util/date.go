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
	return YYYYMMDD(date.Format("2006-01-02"))
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
