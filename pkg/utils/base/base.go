package base

import (
	"regexp"
	"strconv"
	"strings"
	itime "xframe/pkg/utils/time"

	"github.com/Lofanmi/pinyin-golang/pinyin"
)

const (
	phoneRegular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	emailRegular = "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
)

// 验证电子邮件
func ValidateEmail(email string) bool {
	ok, _ := regexp.MatchString(emailRegular, email)
	return ok
}

// 验证手机号码的合法性
func ValidatePhone(mobileNum string) bool {
	reg := regexp.MustCompile(phoneRegular)
	return reg.MatchString(mobileNum)
}

// 动态添加SQL ?号
func Placeholders(n int) string {
	var b strings.Builder
	for i := 0; i < n-1; i++ {
		b.WriteString("?,")
	}
	if n > 0 {
		b.WriteString("?")
	}
	return b.String()
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 汉字转拼音
func China2PinYin(name string) string {
	return pinyin.NewDict().Convert(name, "").None()
}

// 根据身份证字符串获取年龄和性别
func Identity2Data(identity string) (int, int) {
	lenn := len(identity)

	// sex 1为男，0为女
	var age, sex int
	var sexStr string

	// 330521 820721 052
	if lenn == 15 {
		date := identity[6:12]
		sexStr = identity[14:]
		age = itime.DateStr2Year(date, "060102")
	} else if lenn == 18 {
		date := identity[6:14]
		sexStr = identity[16:17]
		age = itime.DateStr2Year(date, "20060102")
	}
	sexInt, _ := strconv.Atoi(sexStr)
	if sexInt%2 == 0 {
		sex = 0
	} else {
		sex = 1
	}
	return age, sex
}
