package spide

import (
	"demo/configure"
	"demo/exception"
	"demo/logger"
	"demo/orm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	NORMAL = iota
	VERIFY = iota
	OTHER  = iota
)

func GetSinglePage(url string) []string {
	var ret []string
	url = strings.Replace(url, "https://www.zhihu.com/question/", "https://www.zhihu.com/api/v4/questions/", 1)
	par := `/answers?include=data[*].is_normal,admin_closed_comment,reward_info,is_collapsed,annotation_action,annotation_detail,collapse_reason,is_sticky,collapsed_by,suggest_edit,comment_count,can_comment,content,editable_content,voteup_count,reshipment_settings,comment_permission,created_time,updated_time,review_info,relevant_info,question,excerpt,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp,is_labeled,is_recognized,paid_info,paid_info_content;data[*].mark_infos[*].url;data[*].author.follower_count,badge[*].topics&`
	url = url + par
	url += `limit=`
	url += string(strconv.Itoa(configure.GetSingleMax()))
	src := url
	pos := 0
	for {
		offset := pos
		url += "&offset="
		url += string(strconv.Itoa(offset))
		url += "&platform=desktop&sort_by=default"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Cookie", configure.GetCookie())
		resp, err := (&http.Client{}).Do(req)
		if err == nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				var nameJson = gjson.Get(string(data), "data.#.author.url_token")
				nameJson.ForEach(func(key, value gjson.Result) bool {
					if value.String() != "" {
						ret = append(ret, value.String())
					}
					return true
				})
				fmt.Println(nameJson)
				var eleNums int = len(nameJson.Array())
				if eleNums != configure.GetSingleMax() {
					break
				}
				pos += configure.GetSingleMax()
				url = src
			} else {
				panic(exception.ZhiError{
					Code:     exception.ReadBodyError,
					FuncName: "GetSinglePage",
				})
			}
		} else {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "GetSinglePage",
			})
		}
	}
	return ret
}

func GetHotPages() []string {
	var ret []string
	req, err := http.NewRequest("GET", "https://www.zhihu.com/hot", nil)

	req.Header.Set("user-agent", configure.GetUserAgent())
	req.Header.Set("Cookie", configure.GetCookie())
	resp, err := (&http.Client{}).Do(req)
	if err == nil {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err == nil {
			rest := doc.Find("div.HotItem-content")
			rest.Each(func(i int, selection *goquery.Selection) {
				resultStr, ok := selection.Find("a").Attr("href")
				if ok {
					ret = append(ret, resultStr)
				}
			})
		}

	}
	return ret
}

func AddNewUser(userToken string) {
	GetUserInformation("https://www.zhihu.com/people/" + userToken)
}

func FreshDayHot() {
	//获取知乎热门50篇个话题
	r := GetHotPages()
	//并发50个协程
	var wg sync.WaitGroup
	wg.Add(len(r))
	for v := 0; v < len(r); v++ {
		go func(url string) {
			username := GetSinglePage(url)
			for _, r := range username {
				AddNewUser(r)
			}
			wg.Done()
		}(r[v])
	}
	wg.Wait()
	//
}

func GetUserInformation(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		req.Header.Set("user-agent", configure.GetUserAgent())
		req.Header.Set("Cookie", configure.GetCookie())
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "GetUserInformation",
			})
		}
		if err == nil {
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err == nil {
				sec := doc.Find("body").Find("#js-initialData")
				location := gjson.Get(sec.Text(), "initialState.entities.users.*.locations.#.name").String()
				school := gjson.Get(sec.Text(), "initialState.entities.users.*.educations.#.school.name").String()
				major := gjson.Get(sec.Text(), "initialState.entities.users.*.educations.#.major.name").String()
				followerCount := gjson.Get(sec.Text(), "initialState.entities.users.*.followerCount").Int()
				followingCount := gjson.Get(sec.Text(), "initialState.entities.users.*.followingCount").Int()
				workIn := gjson.Get(sec.Text(), "initialState.entities.users.*.business.name").String()
				userToken := gjson.Get(sec.Text(), "initialState.entities.users.*.urlToken").String()
				p1 := orm.People{
					UserToken:      userToken,
					Locations:      location,
					School:         school,
					FollowerCount:  int(followerCount),
					FollowingCount: int(followingCount),
					Workin:         workIn,
					Major:          major,
				}
				orm.InsertPeople(p1)
			} else {
				panic(exception.ZhiError{
					Code:     exception.ReadBodyError,
					FuncName: "GetUserInformation",
				})
			}
		}
	}
}

func getRandomQuestion() []string {
	var ret []string
	var url = "https://www.zhihu.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		req.Header.Set("user-agent", configure.GetUserAgent())
		req.Header.Set("Cookie", configure.GetCookie())
		resp, err := (&http.Client{}).Do(req)
		if err == nil {
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err == nil {
				sec := doc.Find("body").Find("#js-initialData")
				as := gjson.Get(sec.Text(), "initialState.entities.answers")
				as.ForEach(func(key, value gjson.Result) bool {
					ret = append(ret, value.Get("question.url").String())
					return true
				})
			} else {
				panic(exception.ZhiError{
					Code:     exception.ReadBodyError,
					FuncName: "getRandomQuestion",
				})
			}
		} else {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "getRandomQuestion",
			})
		}
	} else {
		panic(exception.ZhiError{
			Code:     exception.HttpGetFail,
			FuncName: "getRandomQuestion",
		})
	}
	//预处理
	for i := 0; i < len(ret); i++ {
		s := ret[i]
		pos := strings.LastIndex(s, "/")
		ret[i] = "https://www.zhihu.com/question/" + s[pos+1:]
	}
	return ret
}

func FreshRandom() {
	for {
		var wg sync.WaitGroup
		for i := 0; i < configure.GetMaxThreadNums(); i++ {
			wg.Add(1)
			//
			go func() {
				//处理协程的异常
				defer func() {
					wg.Done()
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()

				r := getRandomQuestion()
				for _, v := range r {
					username := GetSinglePage(v)
					for _, r := range username {
						if len(r) != 0 {
							AddNewUser(r)
						}
					}
				}
			}()
		}
		wg.Wait()
		logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "50 goroutime finished")
		//Check Status
		func() {
			defer func() {
				//处理会抛出的异常
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()

			//N次验证码校验
			success := false
			for i := 0; i < configure.GetFailedRetry(); i++ {
				status := CheckStatus()
				msg := "Try to pass the code:" + strconv.Itoa(i)
				switch status {
				case NORMAL:
					success = true
				case VERIFY:
					PostVerify(GetImageCode())
					logger.DBLog(logrus.Fields{}, logrus.InfoLevel, msg)
				case OTHER:
					logger.DBLog(logrus.Fields{}, logrus.WarnLevel, "Other Problem")
					logrus.Exit(OTHER)
				}
				if success {
					break
				}
				//最大重试次数达到了  还是没通过验证码  退出
				if i == configure.GetFailedRetry()-1 {
					logrus.Exit(OTHER)
				}
			}

		}()
		logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "Sleep")
		time.Sleep(100 * time.Second)
	}
}

func CheckStatus() int {
	req, err := http.NewRequest("GET", "https://www.zhihu.com", nil)
	req.Header.Set("cookie", configure.GetCookie())
	if err == nil {
		resp, err := (&http.Client{}).Do(req)
		if err == nil {
			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				str := string(bytes)
				pos := strings.Index(str, "安全验证")
				if pos != -1 {
					return VERIFY
				}
				pos = strings.Index(str, "我的收藏")
				if pos != -1 {
					return NORMAL
				}
				return OTHER
			} else {
				panic(exception.ZhiError{
					Code:     exception.ReadBodyError,
					FuncName: "CheckStatus",
				})
			}
		} else {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "CheckStatus",
			})
		}
	} else {
		panic(exception.ZhiError{
			Code:     exception.HttpGetFail,
			FuncName: "CheckStatus",
		})
	}
	return OTHER
}
