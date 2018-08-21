package utils

import "github.com/bxcodec/faker"

type Packet struct {
	Oid       string `faker:"month_name" json:"_oid"`
	UserName  string `faker:"username" json:"username"`
	UnixTime  int64  `faker:"unix_time" json:"unixtime"`
	Date      string `faker:"date" json:"date"`
	MonthName string `faker:"month_name" json:"monthName"`
	Year      string `faker:"year" json:"year"`
	DayOfWeek string `faker:"day_of_week" json:"dayOfWeek"`
	Timestamp string `faker:"timestamp" json:"timestamp"`
	TimeZone  string `faker:"timezone"  json:"timezone"`
	IPV4      string `faker:"ipv4" json:"IPv4"`
	Bool      bool
	Int       int
}

func GetRandomPacket() (interface{}, error) {
	a := Packet{}
	err := faker.FakeData(&a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
