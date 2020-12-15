package time

import "time"

const (
	ConstTime = "2006-01-02 15:04:05"
)

//获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err1 := time.ParseInLocation(ConstTime, startTime, time.Local)
	t2, err2 := time.ParseInLocation(ConstTime, endTime, time.Local)
	if err1 == nil && err2 == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
	}
	return hour
}

// 获取当前时间
func NowTime() string {
	return time.Now().Format(ConstTime)
}

// 获取当前时间之前的时间
// min:  time.Minute * 10 10分钟之前
func PreNowTime(min time.Duration) string {
	return time.Now().Add(-min).Format(ConstTime)
}

func DateStr2Year(dataStr, formate string) int {
	timeStr, err := time.Parse(formate, dataStr)
	if err != nil {
		return 0
	}
	str2year := timeStr.Year()

	return time.Now().Year() - str2year
}
