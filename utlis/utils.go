package utlis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//unix 时间格式化

func FormatAsDate(unix int64) string {
	t := time.Unix(unix, 0)
	var t1 = t.Format("2006-01-02")
	return t1
}

type IPInfo struct {
	IP          string `json:"ip"`
	Pro         string `json:"pro"`
	ProCode     string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

// convertGBKToUTF8 将 GBK 编码的字节切片转换为 UTF-8 编码
func convertGBKToUTF8(input []byte) ([]byte, error) {
	// 创建 GBK 解码器
	utf8Bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(input), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return nil, err
	}
	return utf8Bytes, nil
}

// 根据ip 获取城市
func Get_ip_city(ip string) string {
	api_url := "https://whois.pconline.com.cn/ipJson.jsp?ip=101.34.54.23"
	res, err := http.Get(api_url)
	if err != nil {
		return "0.0.0.0"
	}
	defer res.Body.Close()
	fmt.Println(res.Header.Get("Content-Type"))
	if res.StatusCode != http.StatusOK {
		return "0.0.0.0"
	}
	re := regexp.MustCompile(`\{[^{^}]*\}`)
	body, _ := ioutil.ReadAll(res.Body)
	utf8Body, err := convertGBKToUTF8(body)
	matches := re.FindString(string(utf8Body))
	if matches == "" {
		return "0.0.0.0"
	}
	var ipInfo IPInfo
	er := json.Unmarshal([]byte(matches), &ipInfo)
	if er != nil {
		return "0.0.0.0"
	}
	fmt.Println(ipInfo)
	return ipInfo.City
}

// 获取客户端ip
func GetIPFromRequest(c *gin.Context) string {
	// 获取客户端IP的函数
	getIP := func(req *http.Request) string {
		xForwardedFor := req.Header.Get("X-Forwarded-For")
		if xForwardedFor != "" {
			ipList := strings.Split(xForwardedFor, ",")
			if len(ipList) > 0 {
				return strings.TrimSpace(ipList[0])
			}
		}

		xRealIP := req.Header.Get("X-Real-IP")
		if xRealIP != "" {
			return xRealIP
		}

		remoteAddr := req.RemoteAddr
		ip, _, err := net.SplitHostPort(remoteAddr)
		if err != nil {
			return remoteAddr
		}
		return ip
	}

	clientIP := getIP(c.Request)
	return clientIP
}
