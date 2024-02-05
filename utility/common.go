package utility

import (
	"strconv"

	"github.com/gogf/gf/v2/os/gtime"
)

func GetGender(idNumber string) string {
	// 身份证号码的倒数第二位表示性别，偶数为女性，奇数为男性
	genderDigit, _ := strconv.Atoi(string(idNumber[len(idNumber)-2]))

	if genderDigit%2 == 0 {
		return "女"
	} else {
		return "男"
	}
}

func GetBirthDay(idNumber string) *gtime.Time {
	// 身份证号码中倒数第七位到第十四位表示出生日期
	birthYear := idNumber[6:10]
	birthMonth := idNumber[10:12]
	birthDay := idNumber[12:14]

	birthdayStr := birthYear + "-" + birthMonth + "-" + birthDay + " 00:00:00"

	birthday, _ := gtime.StrToTimeFormat(birthdayStr, "Y-m-d H:i:s")

	return birthday
}
