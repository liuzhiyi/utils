package levelcache

import (
	"fmt"
	"testing"
)

type SessionInfo struct {
	AppKey          string //该会员登陆使用的AppKey
	AppId           int    //应用id
	AppType         int    //应用类型
	Loginname       string //会员账号
	Mid             int    //会员ID
	Nickname        string //昵称
	LoginNum        int    //登录密码错误次数
	RangKey         string //短信校验成功时返回的密钥
	RangEndtime     int64  //短信密钥过期时间
	CaptchaCode     string //图形验证码
	CaptchaKey      string //图形验证码成功返回的密钥
	CaptchaEndntime int64  //图形验证密钥过期时间
}

func Test_Action(t *testing.T) {
	cdb := NewSessionDB()
	sessionId := "625a6bf3-27b1-462a-a865-8ed1ee8cd4ad"
	// sinfo := SessionInfo{}
	// sinfo.AppKey = "123456"
	// sinfo.AppType = 1
	// cdb.Set(sessionId, sinfo, 24*60*60)
	var rel SessionInfo
	cdb.Get(sessionId, &rel)
	//fmt.Println(sinfo)
	fmt.Println("sffdgg", rel)
}

func NewSessionDB() *Levelcache {
	levelcache := NewLevecache(24*60*60, 24*60*60, "D:/goWork/src/ucapi/db/session")
	return levelcache
}
