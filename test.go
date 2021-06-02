package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var html_text = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test—-{{.code}}</title>
</head>
<body>
  <h1 style="color:red;">Welcome to come here!</h1>
  <h4>{{.status}}</h4>
  <h4>{{.message}}</h4>
</body>
</html>
`))
func main(){
	var outport int
	flag.IntVar(&outport,"o",8044,"输出端口")
	if !flag.Parsed(){
		flag.Parse()
	}

	engine := gin.Default()
	engine.SetHTMLTemplate(html_text)

	engine.Handle("GET","/", func(context *gin.Context) {
		context.HTML(200,"https",gin.H{
			"code":200,
			"status":"OK",
			"message":"请求成功。一般用于GET与POST请求",
		})
	})

	engine.Handle("GET","/100", func(context *gin.Context) {
		context.HTML(100,"https",gin.H{
			"code":100,
			"status":"Continue",
			"message":"继续。客户端应继续其请求",
		})
	})

	engine.Handle("GET","/101", func(context *gin.Context) {
		context.HTML(101,"https",gin.H{
			"code":101,
			"status":"Switching Protocals",
			"message":"切换协议。服务器根据客户端的请求切换协议。只能切换到更高级的协议，例如，切换到HTTP的新版本协议",
		})
	})

	engine.Handle("GET","/201", func(context *gin.Context) {
		context.HTML(201,"https",gin.H{
			"code":201,
			"status":"Created",
			"message":"已创建。成功请求并创建了新的资源",
		})
	})

	engine.Handle("GET","/202", func(context *gin.Context) {
		context.HTML(202,"https",gin.H{
			"code":202,
			"status":"Accepted",
			"message":"已接受。已经接受请求，但未处理完成",
		})
	})

	engine.Handle("GET","/203", func(context *gin.Context) {
		context.HTML(203,"https",gin.H{
			"code":203,
			"status":"Non-Authoritative Infomation",
			"message":"非授权信息。请求成功。但返回的meta信息不在原始的服务器，而是一个副本",
		})
	})

	engine.Handle("GET","/204", func(context *gin.Context) {
		context.HTML(204,"https",gin.H{
			"code":204,
			"status":"No Content",
			"message":"无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档",
		})
	})

	engine.Handle("GET","/205", func(context *gin.Context) {
		context.HTML(205,"https",gin.H{
			"code":205,
			"status":"Reset Content",
			"message":"重置内容。服务器处理成功，用户终端（例如：浏览器）应重置文档视图。可通过此返回码清除浏览器的表单域",
		})
	})

	engine.Handle("GET","/206", func(context *gin.Context) {
		context.HTML(206,"https",gin.H{
			"code":206,
			"status":"Partial Content",
			"message":"部分内容。服务器成功处理了部分GET请求",
		})
	})

	engine.Handle("GET","/300", func(context *gin.Context) {
		context.HTML(300,"https",gin.H{
			"code":300,
			"status":"Multiple Choices",
			"message":"多种选择。请求的资源可包括多个位置，相应可返回一个资源特征与地址的列表用于用户终端（例如：浏览器）选择",
		})
	})

	engine.Handle("GET","/301", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	engine.Handle("GET","/302", func(context *gin.Context) {
		context.Redirect(302,"https://www.baidu.com")
	})

	engine.Handle("GET","/303", func(context *gin.Context) {
		context.HTML(303,"https",gin.H{
			"code":303,
			"status":"See Other",
			"message":"查看其它地址。与301类似。使用GET和POST请求查看",
		})
	})

	engine.Handle("GET","/304", func(context *gin.Context) {
		context.HTML(304,"https",gin.H{
			"code":304,
			"status":"Not Modeified",
			"message":"未修改。所请求的资源未修改，服务器返回此状态码时，不会返回任何资源。客户端通常会缓存访问过的资源，通过提供一个头信息指出客户端希望只返回在指定日期之后修改的资源",
		})
	})

	engine.Handle("GET","/305", func(context *gin.Context) {
		context.HTML(305,"https",gin.H{
			"code":305,
			"status":"User Proxy",
			"message":"使用代理。所请求的资源必须通过代理访问",
		})
	})

	engine.Handle("GET","/306", func(context *gin.Context) {
		context.HTML(306,"https",gin.H{
			"code":306,
			"status":"Unused",
			"message":"已经被废弃的HTTP状态码",
		})
	})

	engine.Handle("GET","/307", func(context *gin.Context) {
		context.HTML(307,"https",gin.H{
			"code":307,
			"status":"Temporary Redirect",
			"message":"临时重定向。与302类似。使用GET请求重定向",
		})
	})

	engine.Handle("GET","/400", func(context *gin.Context) {
		context.HTML(400,"https",gin.H{
			"code":400,
			"status":"Bad Request",
			"message":"客户端请求的语法错误，服务器无法理解",
		})
	})

	engine.Handle("GET","/401", func(context *gin.Context) {
		context.HTML(401,"https",gin.H{
			"code":401,
			"status":"Unauthorized",
			"message":"请求要求用户的身份认证",
		})
	})

	engine.Handle("GET","/402", func(context *gin.Context) {
		context.HTML(402,"https",gin.H{
			"code":402,
			"status":"Payment Required",
			"message":"保留，将来使用",
		})
	})

	engine.Handle("GET","/403", func(context *gin.Context) {
		context.HTML(403,"https",gin.H{
			"code":403,
			"status":"Forbidden",
			"message":"服务器理解请求客户端的请求，但是拒绝执行此请求",
		})
	})

	engine.Handle("GET","/404", func(context *gin.Context) {
		context.HTML(404,"https",gin.H{
			"code":404,
			"status":"Not Found",
			"message":"服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置\"您所请求的资源无法找到\"的个性页面",
		})
	})

	engine.Handle("GET","/405", func(context *gin.Context) {
		context.HTML(405,"https",gin.H{
			"code":405,
			"status":"Method Not Allowed",
			"message":"客户端请求中的方法被禁止",
		})
	})

	engine.Handle("GET","/406", func(context *gin.Context) {
		context.HTML(406,"https",gin.H{
			"code":406,
			"status":"Not Acceptable",
			"message":"服务器无法根据客户端请求的内容特性完成请求",
		})
	})

	engine.Handle("GET","/407", func(context *gin.Context) {
		context.HTML(407,"https",gin.H{
			"code":407,
			"status":"Proxy Authentication Required",
			"message":"请求要求代理的身份认证，与401类似，但请求者应当使用代理进行授权",
		})
	})

	engine.Handle("GET","/408", func(context *gin.Context) {
		context.HTML(408,"https",gin.H{
			"code":408,
			"status":"Request Time-out",
			"message":"服务器等待客户端发送的请求时间过长，超时",
		})
	})

	engine.Handle("GET","/409", func(context *gin.Context) {
		context.HTML(409,"https",gin.H{
			"code":409,
			"status":"Conflict",
			"message":"服务器完成客户端的 PUT 请求时可能返回此代码，服务器处理请求时发生了冲突",
		})
	})

	engine.Handle("GET","/410", func(context *gin.Context) {
		context.HTML(410,"https",gin.H{
			"code":410,
			"status":"Gone",
			"message":"客户端请求的资源已经不存在。410不同于404，如果资源以前有现在被永久删除了可使用410代码，网站设计人员可通过301代码指定资源的新位置",
		})
	})

	engine.Handle("GET","/411", func(context *gin.Context) {
		context.HTML(411,"https",gin.H{
			"code":411,
			"status":"Length Required",
			"message":"服务器无法处理客户端发送的不带Content-Length的请求信息",
		})
	})

	engine.Handle("GET","/412", func(context *gin.Context) {
		context.HTML(412,"https",gin.H{
			"code":412,
			"status":"Precondition Failed",
			"message":"客户端请求信息的先决条件错误",
		})
	})

	engine.Handle("GET","/413", func(context *gin.Context) {
		context.HTML(413,"https",gin.H{
			"code":413,
			"status":"Request Entity Too Large",
			"message":"由于请求的实体过大，服务器无法处理，因此拒绝请求。为防止客户端的连续请求，服务器可能会关闭连接。如果只是服务器暂时无法处理，则会包含一个Retry-After的响应信息",
		})
	})

	engine.Handle("GET","/414", func(context *gin.Context) {
		context.HTML(414,"https",gin.H{
			"code":414,
			"status":"Request-URI Too Large",
			"message":"请求的URI过长（URI通常为网址），服务器无法处理",
		})
	})

	engine.Handle("GET","/415", func(context *gin.Context) {
		context.HTML(415,"https",gin.H{
			"code":415,
			"status":"Unsupported Media Type",
			"message":"服务器无法处理请求附带的媒体格式",
		})
	})

	engine.Handle("GET","/416", func(context *gin.Context) {
		context.HTML(416,"https",gin.H{
			"code":416,
			"status":"Requested range not Satisfiable",
			"message":"客户端请求的范围无效",
		})
	})

	engine.Handle("GET","/417", func(context *gin.Context) {
		context.HTML(417,"https",gin.H{
			"code":417,
			"status":"Expectation Failed",
			"message":"服务器无法满足Expect的请求头信息",
		})
	})

	engine.Handle("GET","/500", func(context *gin.Context) {
		context.HTML(500,"https",gin.H{
			"code":500,
			"status":"Internal Server Error",
			"message":"服务器内部错误，无法完成请求",
		})
	})

	engine.Handle("GET","/501", func(context *gin.Context) {
		context.HTML(501,"https",gin.H{
			"code":501,
			"status":"Not Implemented",
			"message":"服务器不支持请求的功能，无法完成请求",
		})
	})

	engine.Handle("GET","/502", func(context *gin.Context) {
		context.HTML(502,"https",gin.H{
			"code":502,
			"status":"Bad Gateway",
			"message":"作为网关或者代理工作的服务器尝试执行请求时，从远程服务器接收到了一个无效的响应",
		})
	})

	engine.Handle("GET","/503", func(context *gin.Context) {
		context.HTML(503,"https",gin.H{
			"code":503,
			"status":"Service Unavailable",
			"message":"由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中",
		})
	})

	engine.Handle("GET","/504", func(context *gin.Context) {
		context.HTML(504,"https",gin.H{
			"code":504,
			"status":"Gateway Time-out",
			"message":"充当网关或代理的服务器，未及时从远端服务器获取请求",
		})
	})

	engine.Handle("GET","/505", func(context *gin.Context) {
		context.HTML(505,"https",gin.H{
			"code":505,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
		})
	})

	engine.Handle("GET","/json", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code":200,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
		})
	})

	engine.Handle("GET","/xml", func(context *gin.Context) {
		context.XML(200,gin.H{
			"code":505,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
		})
	})

	engine.Handle("GET","/get", func(context *gin.Context) {
		query := context.Request.URL.Query()
		context.JSON(200,gin.H{
			"code":505,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
			"data":query,
		})
	})

	engine.Handle("POST","/post", func(context *gin.Context) {
		json := make(map[string]interface{})
		contentType := strings.ToLower(context.Request.Header.Get("content-type"))
		fmt.Println(contentType)
		if strings.Contains(contentType,"multipart/form-data"){
			err := context.Request.ParseMultipartForm(128)
			if err != nil {
				fmt.Println(err.Error())
			}
			form := context.Request.Form
			for k,v :=range form{
				if len(v) == 1 {
					json[k] = v[0]
				}else{
					json[k] = strings.Join(v,";")
				}
			}
		}else if  strings.Contains(contentType,"x-www-form-urlencoded") {
			err := context.Request.ParseForm()
			if err != nil {
				fmt.Println(err.Error())
			}
			form := context.Request.Form
			for k,v :=range form{
				fmt.Println(len(v))
				if len(v) == 1 {
					json[k] = v[0]
				}else{
					json[k] = strings.Join(v,";")
				}
			}
		}else{
			context.BindJSON(&json)
		}

		context.JSON(200,gin.H{
			"code":505,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
			"data":json,
		})
	})

	engine.Handle("POST","/api", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code":505,
			"status":"HTTP Version not supported",
			"message":"服务器不支持请求的HTTP协议的版本，无法完成处理",
		})
	})



	outputStr := strconv.Itoa(outport)
	fmt.Println("打开浏览器：http://localhost:"+outputStr)
	engine.Run(":"+ outputStr)
}

