package main

import (
	"demo/spide"
)

type respcode struct {
	Captcha string `json:"captcha"`
}

func main() {
	//spide.GetImageCode()
	//spide.PostVerify(spide.GetImageCode())
	//spide.FreshDayHot()
	//fmt.Println(spide.GetSinglePage("https://www.zhihu.com/question/308829198"))
	//status:=spide.CheckStatus()
	//fmt.Println(status)

	spide.FreshRandom()
	//spide.CheckStatus()
}
