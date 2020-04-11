package configure

var cookie string = `_xsrf=RfQyRc0p7Fy6slGhmeZCSSAWkuas2vbH; _zap=a00721ee-0b5e-455e-a819-04d6c99b7449; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1586175814,1586175843,1586175981,1586176103; d_c0="ANCUFn3QEhGPTlBk65yr6X_jixdXwfCCmLc=|1586100389"; _ga=GA1.2.1435564676.1586100390; _gid=GA1.2.1640474393.1586100390; capsion_ticket="2|1:0|10:1586176650|14:capsion_ticket|44:MjI3OTE0YmY2OGZmNGNjODgxZDVkN2RmMWNkZWU1MmQ=|7bf4e57cd4f50f1e1598030b2c322f4784eaa047c1234d70d1fb3703c36dd0dc"; r_cap_id="NWM4MWM1YzBkMjYxNGJmYjhhMGM1YWUwYWFmMzY2MGM=|1586176275|d807432d0ea6c343f6063ae8c69f66e7565f15c7"; cap_id="MjgzMzE0ODc0NDFkNDEwZWI5NmE2MTUyZGFhY2Q3NzM=|1586176275|075eb98254ef87e5414e195461714a42fbeb62f9"; l_cap_id="MzE2OWE5ODE3MWY2NGQ1OTkzNjQ1ZWM2YTRmM2IxMmY=|1586176275|baf70892e82a36f3e8c5dfd79f1a633e37b680ff"; tst=r; q_c1=77ac187b17d640cf9f1018ec21addb81|1586146307000|1586146307000; KLBRSID=b33d76655747159914ef8c32323d16fd|1586176715|1586161459; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1586176716; anc_cap_id=5b2d754c33494f1293a67ae9a2e2256a; SESSIONID=JbB5PuR2a2kJa77Vftaye7zJty3zAqHjXDzVzv9unLi; JOID=WlwVBUv5LjpiKymxNfca6K2iu7Asv0h-UXNQ5wGQEFUuTBaKZ9cNsTIvIbM1nhD8g49UkQdbxne3KqG8LHfMGGI=; osd=V10WB0j0LzlgKCSwNvUZ5ayhubMhvkt8Un5R5AOTHVQtThWHZtQPsj8uIrE2kxH_gYxZkARZxXq2KaO_IXbPGmE=; l_n_c=1; n_c=1; z_c0=Mi4xTE9BVEF3QUFBQUFBMEpRV2ZkQVNFUmNBQUFCaEFsVk5qM0I0WHdDbnR6b0swSnRlWlA2WmdZb2M1QjVqN1Z0b3pB|1586176655|7527ef434a666277d9940947a15d351381881048; _gat_gtag_UA_149949619_1=1`

func GetCookie() string {
	return cookie
}

func UpdateCapId(capId string) {
	cookie = SetCookieValue(cookie, "anc_cap_id", capId)
}

func GetUserAgent() string {
	return `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36`
}

func GetSingleMax() int {
	return 20
}
func GetSingleReplyMax() int {
	return 20
}
func GetDSN() string {
	return `root:root@tcp(localhost:3306)/zhihu?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`
}

func GetBaiDuApiKey() string {
	return `3Q6Ys54h8Fqa4NQsAqN3RtbN`
}

func GetBaiDuSecretKey() string {
	return `msv69CZ0oXqE2FgbtkbswqVGVdB6m3u1`
}

func GetFailedRetry() int {
	return 100
}

func GetMaxThreadNums() int {
	return 300
}

func GetHttpServerAddressPort() string {
	return ":8888"
}

func GetCheckInterval() int {
	return 100
}
