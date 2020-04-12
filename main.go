package main

import (
	"demo/httpserver"
	"demo/spide"
)

type respcode struct {
	Captcha string `json:"captcha"`
}

func main() {
	//spide.GetImageCode()
	//spide.PostVerify(spide.GetImageCode())
	//spide.FreshDayHot()
	//fmt.Println(spide.GetSinglePage("https://www.zhihu.com/question/384802353"))
	//status:=spide.CheckStatus()
	//fmt.Println(status)
	//spide.AddNewUser("zhi-li-gong-jue");
	//fmt.Println(spide.GetRandomQuestion())
	//fmt.Println(spide.GetSinglePage("https://www.zhihu.com/question/386668818"))
	//return
	go spide.FreshRandom()
	//orm.SelectUserSizeByConditional("locations like '%北京%'")
	//return
	httpserver.StartWebServer()
	//orm.GetLocationsContain("北京")
	//spide.GetSinglePage();
	//spide.FreshRandom()
	//spide.CheckStatus()
}
