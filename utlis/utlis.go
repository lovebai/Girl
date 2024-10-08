package utlis

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix() * 1000
}

// unix 时间格式化
func FormatAsDate(unix int64) string {
	t := time.Unix(unix/1000, 0).In(time.Local)
	var t1 = t.Format("2006-01-02")
	return t1
}

// 时间戳格式化为什么什么之前
func FormatAsTimeAgo(timestamp int64) string {
	now := time.Now()
	t := time.Unix(timestamp/1000, 0)
	diff := now.Sub(t)

	years := int(diff.Hours() / 24 / 365)
	months := int(diff.Hours() / 24 / 30)
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes())
	seconds := int(diff.Seconds())

	switch {
	case years > 0:
		return fmt.Sprintf("%d年之前", years)
	case months > 0:
		return fmt.Sprintf("%d个月之前", months)
	case days > 0:
		return fmt.Sprintf("%d天之前", days)
	case hours > 0:
		return fmt.Sprintf("%d小时前", hours)
	case minutes > 0:
		return fmt.Sprintf("%d分钟前", minutes)
	default:
		return fmt.Sprintf("%d秒之前", seconds)
	}
}

// ConvertTimestampToString 将毫秒时间戳转换为指定格式的时间字符串
func ConvertTimestampToString(timestampMillis int64) string {
	timestamp := timestampMillis / 1000
	t := time.Unix(timestamp, 0).In(time.Local)
	layout := "2006-01-02T15:04"
	return t.Format(layout)
}

// ConvertToTimestamp 将时间字符串转换为时间戳
func ConvertToTimestamp(timeStr string) int64 {
	layout := "2006-01-02T15:04"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0
	}
	return t.Unix() * 1000
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
	utf8Bytes, err := io.ReadAll(transform.NewReader(bytes.NewReader(input), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return nil, err
	}
	return utf8Bytes, nil
}

// 根据ip 获取城市
func Get_ip_city(ip string) string {
	api_url := "https://whois.pconline.com.cn/ipJson.jsp?ip=" + ip
	res, err := http.Get(api_url)
	if err != nil {
		return "0.0.0.0"
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "0.0.0.0"
	}
	re := regexp.MustCompile(`\{[^{^}]*\}`)
	body, _ := io.ReadAll(res.Body)
	utf8Body, _ := convertGBKToUTF8(body)
	matches := re.FindString(string(utf8Body))
	if matches == "" {
		return "0.0.0.0"
	}
	var ipInfo IPInfo
	er := json.Unmarshal([]byte(matches), &ipInfo)
	if er != nil {
		return "0.0.0.0"
	}
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

// 随机返回数组中的一个字符串元素
func GetRandomElement(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(arr))
	return arr[index]
}

// MD5Encrypt 函数接受一个字符串并返回其 MD5 加密后的字符串
func MD5Encrypt(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str + GetConfBody().Salt))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// 输出html
func ToHtml(str string) template.HTML {
	content := blackfriday.Run([]byte(str))
	return template.HTML(content)
}

// 输出html
func ToHtmlTable(str string) template.HTML {
	return template.HTML(str)
}
