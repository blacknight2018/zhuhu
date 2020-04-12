package configure

var cookie string = `_zap=75ffb5d5-f934-4ade-9869-2da897a81be7; d_c0="AFBV59ODFhGPTl7OzLKnKlNoySKnDr2PL7w=|1586348728"; _ga=GA1.2.957551617.1586348729; _gid=GA1.2.916158996.1586348729; capsion_ticket="2|1:0|10:1586423952|14:capsion_ticket|44:MjRkMjJjNWI3YzY3NDFkMGJkMTA4ZDE5YjFhM2FjYTA=|c53f3925798c57dd6ef14245397d0a9177632822b96cb24c24419294e6dfec56"; r_cap_id="NDU0YjJlYmZlNWJjNDU1NDllY2ZhNTI1ZWY0NDI0NjE=|1586423955|7dee5f86efa13a0332b8c032feeaab8bd7b6edc5"; cap_id="ZWM3MmFiOGExYjA0NDNhNmI2NzM0NjY2MDZhMjdjNTE=|1586423955|3f315d30182eca5aeaeaeab82409764e96a00a1c"; l_cap_id="NWY3MWRlMWE3M2U1NGY4Y2JhY2MwYTYxODM0ODIxNWQ=|1586423955|ea7a8cc6c58c9c2629dc1f248f7a68800342ff81"; z_c0=Mi4xTE9BVEF3QUFBQUFBVUZYbjA0TVdFUmNBQUFCaEFsVk5salo4WHdEOWw5Q19Fc1BLZzE5alNpbFdBcU9rT09tWmFR|1586423958|3218bec9bef7d854d961ae2b71b13467be9755f7; q_c1=8546fdec8b47423497c9c9c18dc95531|1586425734000|1586425734000; tshl=; _xsrf=174234a8-1dce-4c59-b414-5a6724cc369d; SESSIONID=BvgV1f5Tpfa7rXbNxLmHlCoYBghPUIaZCCbb7uehfYm; JOID=UFgTB0_z6RmCzVKOY_VdS05hyI92s5l81Po2vTWFjX63vwzuK9-witrJVophYmraCyyXPg7kGzGnqy4oIXLs1zo=; osd=VFwQCk737RqPzFaKYPhcT0pixY5yt5px1f4yvjiEiXq0sg3qL9y9i97NVYdgZm7ZBi2TOg3pGjWjqCMpJXbv2js=; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1586571336,1586580145,1586580233,1586589296; anc_cap_id=6f3f94b632114630a1d05f219a953985; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1586591172; KLBRSID=4efa8d1879cb42f8c5b48fe9f8d37c16|1586591393|1586580143; tst=r`

func GetCookie() string {
	return cookie
}

func UpdateCapId(capId string) {
	cookie = SetCookieValue(cookie, "anc_cap_id", capId)
}

func GetUserAgent() string {
	return `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36`
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
	return 40
}

func GetHttpServerAddressPort() string {
	return ":8888"
}

func GetCheckInterval() int {
	return 100
}

func GetSinglePageLimits() int {
	return 2000
}
