package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
	"xframe/pkg/utils/ip"

	"github.com/gin-gonic/gin"
)

// 本地IP
var localHost = ""

/// 数量控制接口服务器内网IP，getone Ip
var GetOneIp = "127.0.0.1"
var GetOnePort = "8084"

// 服务器间隔时间
var interval = 20

// 创建全局变量
var AccessControll = &AccessControl{sourceArray: make(map[int]time.Time)}

var hashContanst *Constant

// var RabbitMqValidate *rabbitmq.RabbitMQ

// 用来存放控制信息
type AccessControl struct {
	// 用来存放用户想要存放的信息
	sourceArray map[int]time.Time
	sync.RWMutex
}

// 黑名单定义
type BlackList struct {
	listArray map[int]bool
	sync.RWMutex
}

var blackList = &BlackList{listArray: make(map[int]bool)}

// 获取黑名单
func (m *BlackList) GetBlackListById(uid int) bool {
	m.RLock()
	defer m.RUnlock()

	return m.listArray[uid]
}

// 添加黑名单
func (m *BlackList) SetBlackListById(uid int) bool {
	m.Lock()
	defer m.Unlock()

	m.listArray[uid] = true
	return true
}

// 获取指定的数据
func (m *AccessControl) GetNewRecord(uid int) time.Time {
	m.RLock()
	defer m.RUnlock()
	data := m.sourceArray[uid]
	return data
}

// 设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.Lock()
	m.sourceArray[uid] = time.Now()
	m.Unlock()
}

func (m *AccessControl) GetDdistributedRight(c *gin.Context) bool {
	// 获取用户uid
	uid := "1"

	// 采用一致性hash算法，根据用户ID，判断获取具体机器
	hostRequest, err := hashContanst.Get(uid)
	if err != nil {
		return false
	}

	if hostRequest == localHost {
		// 执行本机数据读取与校验
		return m.GetDataFromMap(uid)
	} else {
		// 不是本机
		return GetDataFromOtherMap(hostRequest, c)
	}
}

// 获取本机map，并且处理业务逻辑，返回结果类型bool
func (m *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}

	// 添加黑名单
	if blackList.GetBlackListById(uidInt) {
		return false
	}

	data := m.GetNewRecord(uidInt)
	if !data.IsZero() {
		if data.Add(time.Duration(interval) * time.Second).After(time.Now()) {
			return false
		}
	}
	m.SetNewRecord(uidInt)

	// 测试
	return true
}

func GetDataFromOtherMap(host string, c *gin.Context) bool {
	// http://127.0.0.1:8084/getone/v1/product
	hosturl := "http://" + host + ":8082" + "/validate/v1/product/checkRight"
	response, body, err := GetCurl(hosturl, c)
	if err != nil {
		return false
	}

	// 判断状态
	if response.StatusCode == 200 {
		return string(body) == "true"
	}

	return false
}

func GetCurl(hosturl string, c *gin.Context) (response *http.Response, body []byte, err error) {
	// 获取用户uid
	uid := "2"
	// 验证用户sign
	uidSign := "45rdd434343ddfd4"

	// 模拟接口访问
	client := &http.Client{}
	req, err := http.NewRequest("GET", hosturl, nil)
	if err != nil {
		return
	}

	// 手动指定cookies
	cookieUid := &http.Cookie{Name: "uid", Value: uid, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign, Path: "/"}
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)

	response, err = client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	return
}

// 存储服务器信息，分布到hash环上
func AddHostInfo(hostArray []string) {
	hashContanst = NewConstant()
	for _, v := range hostArray {
		hashContanst.Add(v)
	}

	localHost = ip.GetLocalIP()
	fmt.Println(localHost)
}
