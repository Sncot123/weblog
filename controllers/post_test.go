package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

// 跑测试 如果不用goland自带的  执行代码 go test   /-v
func TestCreatePostHandler(t *testing.T) {
	//r := routes.SetUp()  //使用自己的路由  注意闭环调用！！！！！！！！！！！

	//为避免形成闭环，重新注册一个
	router := gin.Default()
	url := "/api/v1/post"
	router.POST(url, CreatePostHandler)
	body := `{
	"community_id":1,
	"title":"test",
	"content":"just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//判断相应的内容是不是按预期返回了需要登录的错误

	//方法一：判断响应内容是都包含指定字符串
	assert.Contains(t, w.Body.String(), "请登录")

	//方法二：将相应的内容反序列化到ResponseData,然后判断与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed,err:%v\n")
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
