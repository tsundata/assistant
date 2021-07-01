package util

import "time"

const HourLayout = "2006-01-02 15:04"

func IsApartDay(check, last, now string, day int) (bool, error) {
	nowTime, err := time.ParseInLocation(HourLayout, now, time.Local)
	if err != nil {
		return false, err
	}
	checkTime, err := time.ParseInLocation(HourLayout, check, time.Local)
	if err != nil {
		return false, err
	}

	if nowTime.Before(checkTime) {
		return false, nil
	}

	if last == "" { // todo 1 day
		return checkTime.YearDay() == nowTime.YearDay(), nil
	}

	lastTime, err := time.ParseInLocation(HourLayout, last, time.Local)
	if err != nil {
		return false, err
	}

	if nowTime.Sub(lastTime).Hours() >= float64(day)*24 {
		return true, nil
	}
	return false, nil
}

func IsDaily(check, last, now string) (bool, error) {
	return IsApartDay(check, last, now, 1)
}

func IsWeekly(check, last, now string) (bool, error) {
	return IsApartDay(check, last, now, 7)
}

func IsMonthly(check, last, now string) (bool, error) {
	return IsApartDay(check, last, now, 30)
}

func IsAnnually(check, last, now string) (bool, error) {
	return IsApartDay(check, last, now, 356)
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
