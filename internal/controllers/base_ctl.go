package controllers

import (
	"encoding/json"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"os/exec"
	"strings"
	"time"
	"todo_list/doreamon"
	"todo_list/internal/logic"
	"todo_list/internal/models"
	"todo_list/internal/parameters"
)

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	arr := strings.Split(string(output), "\n")
	for _, s := range arr {
		log.Println(s)
	}
	//log.Println(arr[2])
	//arr1 := strings.Split(arr[2], "\n")

	return arr[2], err
}
func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}

type Result struct {
	code    int
	message string
	data    any
}

func RespOk(ctx *gin.Context, data interface{}) {
	//ctx.JSON(200, &Result{code: 200, message: "success", data: data})
	ctx.JSON(200, gin.H{"code": 200, "message": "success", "data": data})
	ctx.Done()
}
func RespError(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "message": "fail", "data": data})
	ctx.Done()
}

func RespErrorWithMsg(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(200, gin.H{"code": code, "message": message, "data": data})
	ctx.Done()
}

// 生成令牌
func generateToken(c *gin.Context, user *models.User) {
	j := &doreamon.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	claims := doreamon.CustomClaims{
		ID:   user.Id,
		Name: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	log.Println(token)
	userData, _ := json.Marshal(user)
	mp := make(map[string]any, 0)
	json.Unmarshal(userData, &mp)
	mp["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功！",
		"data":    mp,
	})
	key := fmt.Sprintf("user:%d:token", user.Id)
	logic.Client.Set(key, token, 3600*time.Second)
	return
}

func generateNewToken(c *gin.Context, user *models.UserModel) {
	j := &doreamon.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	claims := doreamon.CustomClaims{
		ID:   int(user.Time.Id),
		Name: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	log.Println(token)
	userData, _ := json.Marshal(user)
	mp := make(map[string]any, 0)
	json.Unmarshal(userData, &mp)
	mp["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功！",
		"data":    mp,
	})
	key := fmt.Sprintf("user:%d:token", user.Id)
	logic.Client.Set(key, token, 3600*time.Second)
	return
}

const (
	TokenErrorCode          = 1000 // token err
	InsertDBErrorCode       = 1001 // token err
	UpdateDBErrorCode       = 1002 // token err
	ParameterErrorCode      = 1003 // param err
	DeleteDBErrorCode       = 1004 // param err
	QueryDBErrorCode        = 1005 // param err
	UpdatePasswordErrorCode = 1006 // param err
	SendOtpErrorCode        = 1007 // param err
	OtpErrorCode            = 1008 // param err
	LoginPasswordErrorCode  = 2000 // param err
	UnknownErrorCode        = 9999 // token err
)

type Types interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func ParserReqParameters[T parameters.Request](req *T, ctx *gin.Context) {
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		log.Printf("解析参数异常--->>>%s\n ----->>> %#v", err.Error(), req)
		RespError(ctx, ParameterErrorCode, "解析参数异常")
	} else {
		ctx.Next()
	}
}

func ParserReq[T parameters.Request](req *T, ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("解析参数异常--->>>%s\n ----->>> %#v", err.Error(), req)
		//RespError(ctx, ParameterErrorCode, "解析参数异常")
	}
	return err
}

// GetRequestIP 获取ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
func SendEmail(to string, otp string) {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "xx <xxx@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{to}

	// 设置主题
	em.Subject = "小魔童给你发邮件了"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("hello world， 咱们用 golang 发个邮件！！")

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "自己的邮箱账号", "自己邮箱的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
