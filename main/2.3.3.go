package main

//这部分代码定义了"main"包，并导入了两个包："fmt" 用于格式化输入/输出，"net/http" 用于处理HTTP请求和响应。
import (
	"fmt"
	"net/http"
)

// 这是一个函数定义，它接受两个参数：w 和 r。这是一个典型的模式，其中 w 是用于编写HTTP响应的对象，而 r 是包含HTTP请求信息的对象。
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

// 这是两个处理函数，它们将在访问特定路由时执行。
// helloWorldHandler 在请求根路径("/") 时返回字符串"Hello World!"，
// 而 myGetHandler 从URL中提取两个查询参数"arg1"和"arg2"，
// 并在访问"/my_get"路径时生成包含这些值的响应。
func myGetHandler(w http.ResponseWriter, r *http.Request) {
	arg1 := r.URL.Query().Get("arg1")                                    //这一行代码从HTTP请求的URL参数中提取名为 "arg1" 的值，并将其存储在变量 arg1 中。
	arg2 := r.URL.Query().Get("arg2")                                    //这一行代码从URL参数中提取名为 "arg2" 的值，并将其存储在变量 arg2 中。
	response := fmt.Sprintf("接收到参数，分别为: arg1: %s, arg2: %s", arg1, arg2) //这一行代码创建了一个包含参数值的响应字符串。它使用 fmt.Sprintf 函数，将 "arg1" 和 "arg2" 的值插入到字符串中，以便将其包含在响应中。
	fmt.Fprint(w, response)                                              //这行代码将生成的响应字符串 response 写入到 w（http.ResponseWriter 对象），这将作为HTTP响应的主体内容返回给客户端。
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/my_get", myGetHandler)
	http.ListenAndServe(":9090", nil)
}

//这是主函数，它配置路由处理程序并启动HTTP服务器。它通过 http.HandleFunc 来将处理函数与特定的路由路径关联，

//然后通过 http.ListenAndServe 启动服务器监听端口8080。如果出现错误，它将在控制台打印错误消息。

//from flask import Flask
//from flask import request	#引入此依赖，可以接收HTTP请求信息
//app = Flask(__name__)
//
//@app.route('/')
//def hello_world():
//    return 'Hello World!'
//
//@app.route('/my_get', methods=['GET'])
//def my_get():
//    arg1 = request.args.get("arg1") 	# 使用request.args可以接收GET请求的参数
//    arg2 = request.args.get("arg2")
//    return f"接收到参数，分别为: arg1: {arg1} , arg2: {arg2}"  #HTTP请求对应的响应内容
//
//if __name__ == '__main__':
//    app.run()
