package spide

import (
	"bytes"
	"demo/configure"
	"demo/exception"
	"demo/logger"
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"image/gif"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetImageCode() string {
	var ret string
	req, err := http.NewRequest("GET", "https://www.zhihu.com/api/v4/anticrawl/captcha_appeal", nil)
	if err == nil {
		req.Header.Set("cookie", configure.GetCookie())
		req.Header.Set("user-agent", configure.GetUserAgent())
		resp, err := (&http.Client{}).Do(req)
		if err == nil {
			defer resp.Body.Close()
			bys, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				capId := configure.GetCookieValue(resp.Header.Get("Set-Cookie"), "anc_cap_id")
				configure.UpdateCapId(capId)
				str := string(bys)
				ret = gjson.Get(str, "img_base64").String()
			} else {
				panic(exception.ZhiError{
					Code:     exception.ReadBodyError,
					FuncName: "GetImageCode",
				})
			}
		} else {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "GetImageCode",
			})
		}
	} else {
		panic(exception.ZhiError{
			Code:     exception.HttpGetFail,
			FuncName: "GetImageCode",
		})
	}
	logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "get image len:"+strconv.Itoa(len(ret)))
	return ret
}
func GetBaiDuAccessToken() string {
	var ret string
	v := url.Values{}

	v.Set("grant_type", "client_credentials")
	v.Add("client_id", configure.GetBaiDuApiKey())
	v.Add("client_secret", configure.GetBaiDuSecretKey())
	resp, err := http.PostForm("https://aip.baidubce.com/oauth/2.0/token", v)
	if err == nil {
		defer resp.Body.Close()
		byes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			str := string(byes)
			ret = gjson.Get(str, "access_token").String()
		} else {
			panic(exception.ZhiError{
				Code:     exception.ReadBodyError,
				FuncName: "GetBaiDuAccessToken",
			})
		}

	} else {
		panic(exception.ZhiError{
			Code:     exception.RespBodyNil,
			FuncName: "GetBaiDuAccessToken",
		})
	}
	return ret
}
func convertToJpeg222(base64Image string) string {
	var ret string
	ddd, _ := base64.StdEncoding.DecodeString(base64Image)
	ioutil.WriteFile("./in", ddd, 0666)
	cmd := exec.Command("lissajous.exe")
	cmd.Output()
	time.Sleep(time.Second * 1)
	bytes, _ := ioutil.ReadFile("./jpeg")
	ret = base64.StdEncoding.EncodeToString(bytes)
	//os.Remove("./in")
	os.Remove("./jpeg")
	//ret22:=convertToJpeg2(base64Image)
	//fmt.Println(ret22)
	return ret
}
func convertToJpeg(base64Image string) string {
	var ret string
	debases, err := base64.StdEncoding.DecodeString(base64Image)
	if err == nil {
		var bwr io.Reader
		bwr = strings.NewReader(string(debases))
		img, err := gif.Decode(bwr)
		if err == nil {
			var bsb bytes.Buffer
			jpeg.Encode(&bsb, img, nil)
			file, _ := os.Create("code.jpeg")
			file.Write(bsb.Bytes())
			file.Close()
			ret = base64.StdEncoding.EncodeToString(bsb.Bytes())
			time.Sleep(time.Second * 2)
			return ret
		}
	}
	panic(exception.ZhiError{
		Code:     exception.ImageTransError,
		FuncName: "convertToJpeg",
	})
}
func PutCode(code string) {
	apiUrl := "https://www.zhihu.com/api/v4/anticrawl/captcha_appeal"
	respStruct := struct {
		Captcha string `json:"captcha"`
	}{Captcha: code}
	bytes, err := json.Marshal(respStruct)
	if err != nil {
		return
	}
	respJson := string(bytes)
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(respJson))
	req.Header.Set("cookie", configure.GetCookie())
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("User-Agent", configure.GetUserAgent())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	if err == nil {
		resp, err := (&http.Client{}).Do(req)
		if err == nil {
			defer resp.Body.Close()
			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(exception.ZhiError{
					Code:     exception.RespBodyNil,
					FuncName: "PutCode",
				})
			}
			logger.DBLog(logrus.Fields{}, logrus.InfoLevel, "PostCode "+code)
		} else {
			panic(exception.ZhiError{
				Code:     exception.RespBodyNil,
				FuncName: "PutCode",
			})
		}
	} else {
		panic(exception.ZhiError{
			Code:     exception.HttpGetFail,
			FuncName: "PutCode",
		})
	}
}
func PostVerify(base64Image string) {
	base64Image = convertToJpeg(base64Image)
	api := "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	api += "?"
	api += "access_token=" + GetBaiDuAccessToken()
	v := url.Values{}
	v.Set("image", base64Image)
	req, err := http.NewRequest("POST", api, strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := (&http.Client{}).Do(req)
	if err == nil {
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			str := string(bytes)
			code := gjson.Get(str, "words_result.0.words").String()
			PutCode(code)
		} else {
			panic(exception.ZhiError{
				Code:     exception.ReadBodyError,
				FuncName: "PostVerify",
			})
		}
	} else {
		panic(exception.ZhiError{
			Code:     exception.HttpGetFail,
			FuncName: "PostVerify",
		})
	}
}
