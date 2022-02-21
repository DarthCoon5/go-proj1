package time

import "time"

//UTC
func UTC(date time.Time) time.Time {
	return date.In(time.UTC)
}

//Moscow
func Moscow(date time.Time) time.Time {
	secondsEastOfUTC := int((3 * time.Hour).Seconds())
	Moscow := time.FixedZone("Moscow Time", secondsEastOfUTC)
	return date.In(Moscow)
}

//NowUTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

//AnotherLocation
func AnotherLocation(date time.Time, offset int) time.Time {
	location := time.FixedZone("local Time", offset)
	return date.In(location)
}
