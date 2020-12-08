package ip

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// 腾讯查询地理位置API, 用于根据IP查询用户所在地
	TengXunUrl = "https://apis.map.qq.com/ws/location/v1/ip?ip=%s&key=%s"
	TengXunKey = "JFBBZ-PDJCW-6OYRG-ONNB2-DJDB6-PGBRC"
)

func GetCityByIP(ip string) (addr string) {
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
	}

	url := fmt.Sprintf(TengXunUrl, ip, TengXunKey)
	resp, err := http.Get(url)
	if resp == nil || err != nil {
		return
	}
	defer resp.Body.Close()

	var r DataJson
	_ = json.NewDecoder(resp.Body).Decode(&r)

	if r.Status == 0 {
		addr = r.Result.AdInfo.Province + " " + r.Result.AdInfo.City + " " + r.Result.AdInfo.District
	}
	return addr
}

// 定义数据结构
type DataJson struct {
	Status int `json:"status"`
	Result struct {
		AdInfo struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
		} `json:"ad_info"`
	} `json:"result"`
}
