package main

import (
	"encoding/json"

	"bufio"

	"fmt"
	"math/rand"
	"path/filepath"

	"strings"
	"syscall"
	"time"
	"net"

	"io/ioutil"
	"os"

	"net/http"

	"log"

	"os/exec"

	"strconv"

	"sync"
	"runtime"
)

type V1V struct {
	Q18Q	string
	Q11Q	string
	Q17Q	string
	Q6Q	int
	Q3Q	string

	Q1Q	string
	Q14Q	string

	Q16Q	[]string
	Q19Q		string
	Q5Q	int
	Q8Q	string
	Q15Q	int
	Q2Q	string
	Q4Q	string
	Q7Q	int

	Q9Q	string
	Q10Q	map[string]string
	Q21Q		string
	Q20Q		string
	Q13Q	string
	Q12Q	string

	Q22Q	sync.Mutex
}

var V2V = &V1V{}

type V3V struct {
	A0_str	string
	L1_str	string
	L2_str	string
	T1_str	string
	T2_str	string
	T3_str	string

	M3_str	string

	C2_str	string
	C3_str	string
	C5_str	string
	C6_str	string

	M9_str	string
}

var V4V = &V3V{}

func main() {

	F0F()

	go F9F()

	select {}

}

func F0F() {

	V2V.Q11Q, _ = os.Getwd()

	V2V.Q10Q = make(map[string]string)

	V2V.Q18Q = "汉语"
	V2V.Q18Q = "英语English"

	if V2V.Q18Q == "汉语" {

		V2V.Q12Q = "关"
		V2V.Q17Q = "PC-gui框架演示-V1.1"
	} else {

		V2V.Q12Q = "off"
		V2V.Q17Q = "PC-gui framework demonstration -V1.1"
	}

	V2V.Q1Q = fmt.Sprintf("%s/%s.html", V2V.Q11Q, V2V.Q17Q)

	return

}

func F1F(V53V string) {

	V53V = strings.TrimSpace(V53V)

	var V8V = F23F(V53V)

	var V57V string

	defer func() {

	}()

	if strings.HasPrefix(V53V, "初始:") {
		V2V.Q6Q++

		V4V.M3_str = F22F(V8V["初始欢迎"])
		V4V.L2_str = ""
		V2V.Q3Q = "前端"
	} else if strings.HasPrefix(V53V, "输入内容:") {
		_, V53V, _ = strings.Cut(V53V, ":")

		var V13V = "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "]"

		V13V = "<br />" + V13V

		V57V = V8V["输入内容提示"] + V53V + V13V

		F20F("成功", V57V, strings.Replace(V57V, "<br />", "\n", -1))

	} else if strings.HasPrefix(V53V, "流式输出:") {
		_, V53V, _ = strings.Cut(V53V, ":")

		V2V.Q16Q = V2V.Q16Q[:0]
		V2V.Q15Q = 10
		V2V.Q4Q = ""
		V2V.Q8Q = ""

		V2V.Q22Q.Lock()
		V2V.Q19Q = "123"
		V2V.Q7Q++

		V2V.Q22Q.Unlock()

		go func() {

			F24F(V2V.Q15Q, V2V.Q7Q)

		}()

		V4V.M9_str = "流式:"

	} else if strings.HasPrefix(V53V, "停止流式的输出:") {

		V2V.Q22Q.Lock()
		V2V.Q19Q = "流式结束"
		V2V.Q22Q.Unlock()

	} else if strings.HasPrefix(V53V, "对话前的多次弹窗配置:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q16Q = V2V.Q16Q[:0]
		V2V.Q15Q = 0
		V2V.Q4Q = ""
		V2V.Q8Q = ""

		V2V.Q22Q.Lock()
		V2V.Q19Q = "123"
		V2V.Q7Q++
		V2V.Q22Q.Unlock()

		go func() {

			F25F()

			F24F(V2V.Q15Q, V2V.Q7Q)

		}()

		V4V.M9_str = "多次等待弹窗选择判断:"
		V2V.Q2Q = "多次等待弹窗选择判断:"
	} else if strings.HasPrefix(V53V, "流式输出一多项选择:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V53V = strings.TrimSpace(V53V)

		if V53V == "对话配置完成" {
			V2V.Q4Q = V53V
		} else if strings.Contains(V53V, "次") {
			V53V = strings.Replace(V53V, "次", "", -1)

			var V22V, V21V = strconv.Atoi(strings.TrimSpace(V53V))
			if V21V != nil {
				V2V.Q15Q = 5
			} else {
				V2V.Q15Q = V22V
			}

		} else {

		}

		V2V.Q2Q = "流式输出一多项选择:" + V53V

	} else if strings.HasPrefix(V53V, "多次等待弹窗选择判断:") {

		V4V.M9_str = ""
		var V24V = V2V.Q7Q
		for {
			if V2V.Q19Q == "流式终断" {

				V2V.Q8Q = "流式终断"
				V4V.M9_str = "流式:"
				return
			} else if V24V != V2V.Q7Q {
				return
			}

			if len(V2V.Q2Q) > 2 {

				V2V.Q2Q = ""
				break
			}
			time.Sleep(300 * time.Millisecond)

		}

		if V2V.Q15Q == 0 {

			V4V.C3_str = F27F()
			V4V.M9_str = "多次等待弹窗选择判断:"
			return
		}

		if V2V.Q4Q == "" {
			V4V.C3_str = F28F()
			V4V.M9_str = "多次等待弹窗选择判断:"
			return
		}

		V2V.Q8Q = "等待多选择完成"
		V4V.M9_str = "流式:"

	} else if strings.HasPrefix(V53V, "流式:") {
		V4V.M9_str = ""
		var V26V = V2V.Q7Q
		var V28V = V8V["流式结束"]
		for {
			if V26V != V2V.Q7Q {
				return

			} else if V2V.Q19Q == "流式结束" {

				V4V.L1_str = V28V
				return
			}

			V29V := len(V2V.Q16Q)

			if V29V == 0 || V29V == V2V.Q5Q {
				continue
			}
			V2V.Q5Q = V29V
			break

			time.Sleep(300 * time.Millisecond)

		}
		V4V.M9_str = "流式:"
		V4V.L1_str = strings.Join(V2V.Q16Q, "\n")

	} else if strings.HasPrefix(V53V, "单选选择框示例:") {

		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q9Q = ""

		V4V.L2_str = F30F()
		V57V = V8V["展示选择框示例"]
		F20F("成功", V57V, V57V)
	} else if strings.HasPrefix(V53V, "复选选择框示例:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q10Q = make(map[string]string)

		V4V.L2_str = F29F()
		V57V = V8V["展示选择框示例"]
		F20F("成功", V57V, V57V)
	} else if strings.HasPrefix(V53V, "单选值:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q9Q = strings.TrimSpace(V53V)

		V4V.L2_str = F30F()
		V57V = V8V["单选框的选值"] + V53V
		F20F("成功", V57V, V57V)

	} else if strings.HasPrefix(V53V, "复选值:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		var V39V, V38V, _ = strings.Cut(V53V, ":")

		if V39V == "对" {
			V57V = V8V["复选框的选值"] + V38V
			V2V.Q22Q.Lock()
			V2V.Q10Q[V38V] = "checked"
			V2V.Q22Q.Unlock()
		} else {
			V57V = V8V["复选框的取消值"] + V38V
			V2V.Q22Q.Lock()
			V2V.Q10Q[V38V] = ""

			V2V.Q22Q.Unlock()
		}
		V4V.L2_str = F29F()

		F20F("成功", V57V, V57V)

	} else if strings.HasPrefix(V53V, "下拉列表示例:") {

		V4V.L2_str = F31F()
		V57V = V8V["下拉表单示例"]
		F20F("成功", V57V, V57V)
	} else if strings.HasPrefix(V53V, "推理等级:") {

		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q13Q, V2V.Q12Q, _ = strings.Cut(V53V, ":")

		V4V.L2_str = F31F()
		V57V = V8V["选择推理等级"] + V2V.Q12Q
		F20F("成功", V57V, V57V)
	} else if strings.HasPrefix(V53V, "多项输入表单示例:") {
		V2V.Q21Q, V2V.Q20Q = "", ""

		V4V.L2_str = F32F()
		V57V = V8V["多项输入表单示例"]
		F20F("成功", V57V, V57V)
	} else if strings.HasPrefix(V53V, "多项输入表单提交:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V2V.Q21Q, V2V.Q20Q, _ = strings.Cut(V53V, "||")
		V57V = fmt.Sprintf(V8V["多项输入表单的值"], V2V.Q21Q, V2V.Q20Q)

		V4V.L2_str = F32F()

		F20F("成功", V57V, strings.Replace(V57V, "<br />", "\n", -1))

	} else if strings.HasPrefix(V53V, "停止流式的输出:") {
		V2V.Q22Q.Lock()
		V2V.Q19Q = "流式结束"
		V2V.Q22Q.Unlock()

	} else if strings.HasPrefix(V53V, "打开目录:") {

		_, V53V, _ = strings.Cut(V53V, ":")

		var V52V = strings.TrimSpace(V53V)

		V57V = V8V["打开目录"] + V52V
		F20F("成功", V57V, V57V)

		V52V = strings.Replace(V52V, "/", `\`, -1)

		F4F(V52V)
	} else if strings.HasPrefix(V53V, "PC-Gui框架说明:") {
		_, V53V, _ = strings.Cut(V53V, ":")
		V4V.L2_str = F33F()

	} else if strings.HasPrefix(V53V, "赞赏:") {

		var V56V string

		V56V = `<br /><h2 class="tishi-model-item">&nbsp;&nbsp;赞赏&nbsp;&nbsp;<i class="fas fa-mug-hot"></i></h2><br />%s`
		V56V = fmt.Sprintf(V56V, F22F(`<img src="https://pic1.imgdb.cn/item/69c5222753701b6d63d8f497.jpg" width="230" height="300">`))

		V4V.L2_str = V56V

	} else if strings.HasPrefix(V53V, "退出:") {

		V57V = F22F(V8V["程序已经退出提示"])
		F20F("成功", V57V, V8V["程序已经退出提示"])

		go func() {
			time.Sleep(1 * time.Second)
			log.Fatal("退出程序.")
			os.Exit(0)

		}()

	} else {

	}

	return

}

func F2F(V58V string) {

	var V61V = `<html> <head><meta http-equiv="Content-Type" content="text/html; charset=utf8" /><title></title><meta http-equiv="refresh" content="0;URL=替换网址"></head><body></body></html>`

	V61V = strings.Replace(V61V, "替换网址", V58V, -1)
	F7F(V2V.Q1Q, V61V, "保存")

	return

}

func F3F(V62V string) {
	time.Sleep(300 * time.Millisecond)

	V63V := exec.Command("rundll32", "url.dll,FileProtocolHandler", V62V).Run()
	if V63V == nil {
		return
	}

	V64V := exec.Command("cmd", "/c", "start", V62V)

	V64V.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	var V66V = V64V.Start()
	if V66V != nil {

		_ = V66V
	}
	return
}

func F4F(V71V string) {
	V71V = strings.Replace(V71V, `/`, `\`, -1)
	V71V = strings.Replace(V71V, `\\`, ``, -1)

	V71V = strings.TrimSpace(V71V)
	V71V = strings.Trim(V71V, `\`)

	V72V := exec.Command("cmd", "/c", "start", "", V71V)

	V74V := V72V.Run()

	if V74V == nil {

		return
	}

	V74V = exec.Command(`cmd`, `/c`, `explorer`, V71V).Start()

	if V74V != nil {

		_ = V74V
	}

	return

}

func F5F() {
	var V75V runtime.MemStats
	runtime.ReadMemStats(&V75V)

	runtime.GC()
	runtime.ReadMemStats(&V75V)

}

func F6F(V76V string) string {
	V77V, V78V := ioutil.ReadFile(V76V)
	if V78V != nil {

		log.Fatal("读取本地主页网页文件并返回源码 ：错误,Read the local homepage webpage file and return the source code: error.", V78V)
		return ""
	}

	return string(V77V)
}

func F7F(V79V, V80V, V81V string) {
	var V83V = filepath.Dir(V79V)
	var _, V89V = os.Stat(V83V)
	if V89V != nil {

		_ = os.MkdirAll(V83V, os.ModePerm)

	}

	var V88V *os.File
	if V81V == "追加" {
		V88V, V89V = os.OpenFile(V79V, os.O_APPEND|os.O_CREATE, 0666)
	} else {
		V88V, V89V = os.OpenFile(V79V, os.O_TRUNC|os.O_CREATE, 0666)

	}

	if V89V != nil {

		return
	}
	defer V88V.Close()
	var V91V = bufio.NewWriter(V88V)
	V91V.WriteString(V80V)
	V91V.Flush()

}

func F8F() {
	fmt.Println("")
	time.Sleep(100 * time.Millisecond)
	log.Fatal("")
}

type V92V struct {
	Message string `json:"message"`
}

func F9F() {

	var V94V int
	rand.Seed(time.Now().UnixNano())
	V94V = 35658 + rand.Intn(2300)

	for {
		if F11F(V94V) {

			break
		} else {

			V94V++
		}

	}

	var V96V = fmt.Sprintf(":%d", V94V)

	http.HandleFunc("/", F10F)

	http.HandleFunc("/chat", F12F)

	http.HandleFunc("/bc", F13F)

	V2V.Q14Q = fmt.Sprintf("http://localhost%s", V96V)

	go func() {

		F2F(V2V.Q14Q)

		F3F(V2V.Q14Q)

	}()

	if V97V := http.ListenAndServe(V96V, nil); V97V != nil {

		log.Fatal(V97V)
	}

}

func F10F(V98V http.ResponseWriter, V99V *http.Request) {

	if V99V.URL.Path != "/" {
		http.NotFound(V98V, V99V)
		return
	}

	var V113V string

	V113V = F14F()

	V113V = strings.Replace(V113V, "|替换换行|", "\n", -1)

	V113V = F17F(V113V)
	V113V = F16F(V113V)

	V113V = strings.Replace(V113V, "|替换程序名|", V2V.Q17Q, -1)
	V113V = strings.Replace(V113V, "|替换程序目录|", V2V.Q11Q, -1)

	V113V = strings.Replace(V113V, "|替换CSS样式|", F19F(), -1)

	var V109V = make(map[string]string)

	if V2V.Q18Q == "汉语" {

		V109V["IE浏览器过时"] = "您的IE浏览器过时，推荐使用 Chrome/Firefox/Edge 等现代浏览器."
	} else {

		V109V["IE浏览器过时"] = "Your IE browser is outdated. We recommend using modern browsers like Chrome, Firefox, or Edge."

	}

	V113V = strings.Replace(V113V, "|替换IE浏览器过时|", V109V["IE浏览器过时"], -1)
	V113V = strings.Replace(V113V, "<!--替换各类AI模型接入-->", V109V["AI模型接入按钮"], -1)
	V113V = strings.Replace(V113V, "<!--替换各国语言-->", F18F(), -1)

	V113V = strings.Replace(V113V, "|替换程序名|", V2V.Q17Q, -1)

	V98V.Header().Set("Content-Type", "text/html; charset=utf-8")

	V98V.Write([]byte(V113V))

	return
}

func F11F(V114V int) bool {

	V115V, V116V := net.Listen("tcp", fmt.Sprintf(":%d", V114V))
	if V116V != nil {

		return false
	}

	_ = V115V.Close()
	return true
}

func F12F(V117V http.ResponseWriter, V118V *http.Request) {
	defer func() {

	}()

	if V118V.Method != http.MethodPost {
		http.Error(V117V, "只支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	var V119V V92V
	V120V := json.NewDecoder(V118V.Body).Decode(&V119V)
	if V120V != nil {
		http.Error(V117V, "解析 JSON 数据失败", http.StatusBadRequest)
		return
	}
	defer V118V.Body.Close()

	V2V.Q22Q.Lock()
	*V4V = V3V{}
	V2V.Q22Q.Unlock()

	F1F(V119V.Message)

	V117V.Header().Set("Content-Type", "application/json")

	json.NewEncoder(V117V).Encode(V4V)

	time.Sleep(200 * time.Millisecond)

	V118V.Body.Close()
	
	if strings.HasPrefix(V119V.Message, "多次等待弹窗选择判断:")||strings.HasPrefix(V119V.Message, "流式:") { 

		} else { 
			F5F()
		} 
	

}

func F13F(V121V http.ResponseWriter, V122V *http.Request) {
	V121V.Header().Set("Content-Type", "text/event-stream")
	V121V.Header().Set("Cache-Control", "no-cache")
	V121V.Header().Set("Connection", "keep-alive")
	V121V.Header().Set("Access-Control-Allow-Origin", "*")

	V123V, V124V := V121V.(http.Flusher)
	if !V124V {
		http.Error(V121V, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	var V126V = V2V.Q6Q
	for {
		for {
			if V2V.Q6Q != V126V {

				return
			}

			if V2V.Q3Q != "" {

				V2V.Q3Q = ""
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		V127V, _ := json.Marshal(V4V)

		fmt.Fprintf(V121V, "data: %s\n\n", V127V)
		V123V.Flush()

	}

}

func F14F() (V130V string) {

	V130V = `

	<!DOCTYPE html>
	<html lang="|替换语言|">
	<head>
				<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge"> 
		<!--[if IE]><script>alert('|替换IE浏览器过时|');</script><![endif]-->
			<title>|替换程序名|</title>
			<link rel="shortcut icon" href="./app/20B.ico" type="image/x-icon">
			|替换CSS样式|  
			
	</head>

	<body>
	      
	<div class="container">
    <div id="C5_str">
            <div id="C6_str">
            </div>
        </div>
          <div id="C2_str" class="Dataset-Text">
            <div id="C3_str">
            </div>
        </div>
    <div class="history-panel">
        <div id="toast"></div>
         <!--替换各国语言--> <br />
		 <div class="input-group">
			<input type="text" class="input-field" placeholder="|替换输入框说明|">
			<button class="btn btn-primary Dataset-Text"  data-text="输入内容:" title="&nbsp;|替换输入按钮提示|">►|替换确认|</button>

		</div>

		<textarea class="textarea-panel"  id="L1_str"  placeholder="|替换只读多行输入框|" readonly></textarea>

		<br />

			<div id="L2_str"> </div>

			<div class="button-row" style="grid-template-columns: 70% 30%;"> 

			<button class="stream-button Dataset-Text s1"  data-text="流式输出:"  title="&nbsp;|替换Ai流式输出按钮提示|">💬|替换Ai流式输出按钮|</button>
			<button class="stream-button Dataset-Text s2"  data-text="停止流式的输出:" title="&nbsp;|替换停止输出按钮提示|">⏹️|替换停止流式输出|</span>

		</div>


		<div class="button-row" style="grid-template-columns: 70% 30%;"> 

			<span class="stream-button Dataset-Text s1"  data-text="对话前的多次弹窗配置:" title="&nbsp;|替换多次弹窗提示|"> ⚙️|替换多次弹窗配置| </span>
			<span class="stream-button Dataset-Text s2"  data-text="停止流式的输出:" title="&nbsp;|替换停止输出按钮提示|">⏹️|替换停止流式输出|</span>
		</div>



		<hr>

		<div class="button-row" style="grid-template-columns: 36% 33% 30%;">
			<span class="stream-button Dataset-Text s1"  data-text="单选选择框示例:" title="&nbsp;|替换单选选择框提示|">🔘|替换单选选择框示例|</span>
			<span class="stream-button Dataset-Text s2"  data-text="复选选择框示例:" title="&nbsp;|替换复选选择框提示|"> ☑️|替换复选选择框示例|</span>
			<span class="stream-button Dataset-Text s3"  data-text="下拉列表示例:" title="&nbsp;|替换下拉列表提示|">🔽|替换下拉列表示例|</span>

		</div>

		<div class="button-row" style="grid-template-columns: 36% 33% 30%;">
			<span class="stream-button Dataset-Text s1"  data-text="多项输入表单示例:" title="&nbsp;|替换输入表单提示|">📝|替换多项输入表单示例|</span>
			<span class="stream-button Dataset-Text s2"  data-text="打开目录:|替换程序目录|"  title="&nbsp;|替换打开目录提示|">📁 |替换打开程序目录|</span>
			<span class="stream-button Dataset-Text s3"  data-text="赞赏:"  title="&nbsp;|替换赞赏提示|">❤️|赞赏|</span>

		</div>
		
		<hr>

		<div class="button-row" style="grid-template-columns: 36% 33% 30%;">
			<span class="stream-button Dataset-Text s3"  data-text="PC-Gui框架说明:"  title="&nbsp;|替换PC-Gui框架说明|">💻|替换PC-Gui框架说明|</span>
			<span class="stream-button Dataset-Text s2"  data-text="退出:" title="&nbsp;|替换退出程序提示|">🚫|替换退出程序|</span>

			</div>
		

		</div>
		<script>
	
	
		|替换JS|

		</script>
    
	</body>
	</html>
	`
	V130V = strings.Replace(V130V, "|替换JS|", F15F(), -1)

	return

}

func F15F() (V133V string) {

	V133V = `   const L1_str = document.getElementById('L1_str');
        const L2_str = document.getElementById('L2_str');
        const T2_str = document.getElementById('T2_str');
        const C2_str = document.getElementById('C2_str');
           const C3_str = document.getElementById('C3_str');
           const C5_str = document.getElementById('C5_str');
           const C6_str = document.getElementById('C6_str');
           const M3 = document.getElementById('toast');
       let eventSource = null;
        window.onload = function() {   
          PostB("初始:").then((reTxt) => {
          });
          connect()
       };
       document.querySelector('.input-field').addEventListener('keypress', function(e) {
           if(e.key === 'Enter') {
               document.querySelector('.btn-primary').click();
           }
       });
		document.body.addEventListener('click', function(event) {
		const target = event.target.closest('.Dataset-Text');
		if (!target) return;
		let  dataText =''
		if (target) {
			dataText = target.dataset.text
		}
		if (target.matches('button')) { 
			const inputElement = target.previousElementSibling;
			if (inputElement) {
				if (inputElement.tagName === 'INPUT' || inputElement.tagName === 'TEXTAREA') {
					const inputText = inputElement.value;
					dataText += inputText;  
				}
			}
		} 
			if (dataText.length < 2) {
				return
			}  
			PostB(dataText).then((data) => {
			});
		});
		document.body.addEventListener('change', function(event) {
		let  dataText =''
		if (event.target.classList.contains('change-value') && event.target.tagName === 'SELECT') {
			const selectedOption = event.target.options[event.target.selectedIndex];
			dataText= event.target.value;
				if (dataText.includes("归0:")) {
					dataText = dataText.replace("归0:", ""); 
				}
				PostB(dataText).then((data) => {
				if (event.target.value.includes("归0:")) {
					event.target.value = ""
				}
			});
		}
		});  
		document.body.addEventListener('change', function(event) {
		const target = event.target;
			if (target.type === 'checkbox' && target.classList.contains('toggle-item')) {
			const checked = target.checked;
			let checked_str='';
			if (checked == false){ 
				checked_str="复选值:否:"+target.value;
			} else {
				checked_str="复选值:对:"+target.value;
			}
			PostB(checked_str).then((reTxt) => {
				});
			} else if (target.type === 'radio' && target.name === 'choice') {
			PostB("单选值:"+target.value).then((reTxt) => {
				});
			}
		});
		document.body.addEventListener('submit', function(event) {
		const form = event.target;
			if (form.classList.contains('ajax-form')) {
				event.preventDefault();
				const formData = new FormData(form);
			const name = form.querySelector('[name="fullName"]').value;
			const interest = form.querySelector('[name="interest"]').value;
			console.log('姓名:', name);
			console.log('兴趣:', interest);
					PostB("多项输入表单提交:"+name+'||'+interest).then((reTxt) => {
				});
			return;
		}
		});
		function connect() {
		if (eventSource) {
			return;
		}
		eventSource = new EventSource('/bc');
		eventSource.onopen = function() {
		};
		eventSource.onmessage = function(event) {
			try {
				UPpage(JSON.parse(event.data)); 
			} catch (e) {
			}
		};
		}
		function UPpage(data) {
			if (data.A0_str !== "") {
				console.log(data.A0_str); 
					return ; 
				}
			console.log("data数据"); 
			console.log(data); 
				if (data.L1_str !== "") {
					L1_str.innerHTML = data.L1_str; 
					L1_str.scrollTop = L1_str.scrollHeight;
				}
					if (data.L2_str !== "") {
						L2_str.innerHTML = data.L2_str; 
					} else { 
						L2_str.innerHTML =  ""
				}
				if (data.M3_str !== "") {
						showTemporaryAlert(data.M3_str)
					}
					if (data.C3_str !== "") {
					C2_str.style.display = 'flex';
					C3_str.innerHTML = data.C3_str;
				} else { 
					C2_str.style.display = 'none';
				}
				if (data.C6_str !== "") {
					C5_str.style.display = 'flex';
					C6_str.innerHTML = data.C6_str;
				} else { 
					C5_str.style.display = 'none';
				}
				if (data.T1_str !== "") {
					showTemporaryAlert(data.T1_str)
				}
				if (data.M9_str !== "") {
					PostB(data.M9_str).then((data) => {
					});
				}
				if (data.T2_str !== "") {
					T2_str.innerHTML = data.T2_str; 
				}
		}
		async function PostB(reTxt) {
		const message = reTxt.trim();
		try {
			const response = await fetch('/chat', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({message})
			});
			if (!response.ok) {
				throw new Error(|替换点符号|HTTP error! status: ${response.status}|替换点符号|);
			}
			const data = await response.json();
			UPpage(data) 
			return null;
		} catch (error) {
			console.error('err:', error);
			return null;
		}
		}
		function showTemporaryAlert(reTxt) {
				M3.innerHTML = reTxt;
				M3.classList.add('active');
				setTimeout(() => {
					M3.innerHTML =''
					M3.classList.remove('active');              
				}, 3500);
			}


			`

	V133V = strings.Replace(V133V, "|替换点符号|", "`", -1)
	return
}

func F16F(V152V string) string {
	var V136V = make(map[string]string)
	if V2V.Q18Q == "汉语" {

		V136V["网页语言"] = "zh-CN"
		V136V["请输入内容"] = "请输入内容..."
		V136V["确认"] = "确认"
		V136V["Ai流式输出变更按钮"] = "Ai流式输出-变更按钮"
		V136V["停止流式输出"] = "停止流式输出"
		V136V["对话前的多次弹窗配置"] = "对话前的多次弹窗配置"
		V136V["单选选择框示例"] = "单选选择框示例"
		V136V["复选选择框示例流式"] = "复选选择框示例"
		V136V["下拉列表示例"] = "下拉列表示例"
		V136V["多项输入表单示例"] = "多项输入表单示例"
		V136V["打开程序目录"] = "打开程序目录"
		V136V["退出程序"] = "退出程序"

		V136V["只读多行输入框"] = "只读多行输入框. 鼠标悬停按钮时，自动显示提示信息。"

		V136V["赞赏"] = "赞赏"
		V136V["PC-Gui框架说明"] = "PC-Gui框架说明"
	} else {

		V136V["网页语言"] = "en-US"
		V136V["请输入内容"] = "Please enter content..."
		V136V["确认"] = "Confirm"
		V136V["只读多行输入框"] = "Read-only multiline input box.A tooltip appears when you hover over the button."

		V136V["Ai流式输出变更按钮"] = "AI Streaming Output - Change Button"
		V136V["停止流式输出"] = "Stop Streaming Output"
		V136V["对话前的多次弹窗配置"] = "Multiple Popup Configuration Before Conversation"
		V136V["单选选择框示例"] = "Example of radio selection box"
		V136V["复选选择框示例流式"] = "Checkbox Example"
		V136V["下拉列表示例"] = "Dropdown List Example"
		V136V["多项输入表单示例"] = "Multi-input Form Example"
		V136V["打开程序目录"] = "Open Program Directory"
		V136V["退出程序"] = "Exit Program"
		V136V["赞赏"] = "admire"
		V136V["PC-Gui框架说明"] = "PC-Gui explain"

	}

	V152V = strings.Replace(V152V, "|替换程序名|", V2V.Q17Q, -1)

	V152V = strings.Replace(V152V, "|替换语言|", V136V["网页语言"], -1)
	V152V = strings.Replace(V152V, "|替换输入框说明|", V136V["请输入内容"], -1)
	V152V = strings.Replace(V152V, "|替换只读多行输入框|", V136V["只读多行输入框"], -1)

	V152V = strings.Replace(V152V, "|替换确认|", V136V["确认"], -1)
	V152V = strings.Replace(V152V, "|替换Ai流式输出按钮|", V136V["Ai流式输出变更按钮"], -1)
	V152V = strings.Replace(V152V, "|替换停止流式输出|", V136V["停止流式输出"], -1)
	V152V = strings.Replace(V152V, "|替换多次弹窗配置|", V136V["对话前的多次弹窗配置"], -1)
	V152V = strings.Replace(V152V, "|替换单选选择框示例|", V136V["单选选择框示例"], -1)
	V152V = strings.Replace(V152V, "|替换复选选择框示例|", V136V["复选选择框示例流式"], -1)
	V152V = strings.Replace(V152V, "|替换下拉列表示例|", V136V["下拉列表示例"], -1)
	V152V = strings.Replace(V152V, "|替换多项输入表单示例|", V136V["多项输入表单示例"], -1)
	V152V = strings.Replace(V152V, "|替换打开程序目录|", V136V["打开程序目录"], -1)
	V152V = strings.Replace(V152V, "|替换退出程序|", V136V["退出程序"], -1)
	V152V = strings.Replace(V152V, "|替换PC-Gui框架说明|", V136V["PC-Gui框架说明"], -1)

	V152V = strings.Replace(V152V, "|赞赏|", V136V["赞赏"], -1)

	return V152V
}

func F17F(V166V string) string {
	var V155V = make(map[string]string)
	if V2V.Q18Q == "汉语" {

		V155V["输入按钮"] = `前端html: class='Dataset-Text' data-text='输入内容:'&#10;JS:统一响应点击click元素'.Dataset-Text',同时检查 INPUT 和 TEXTAREA 元素 &#10;后端接收并响应:strings.HasPrefix('', '输入内容:')`
		V155V["流式输出按钮"] = `前端html: 异步不断更新前端网页内容,减少AI返回内容等待时间,增加动态感.&#10;<textarea id='L1_str'> </textarea>&#10;JS:接收后端数据, if (data.L1_str !== '') {&#10;L1_str.innerHTML = data.L1_str;&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;后端接收并响应: (struct).L1_str='AI返回内容' &#10; (struct).M9_str='流式:'&#10; strings.HasPrefix('', '流式:')`
		V155V["停止输出"] = `前端html: class='Dataset-Text' data-text='停止流式的输出:'&#10;JS:统一响应点击click元素'.Dataset-Text' &#10;后端接收并响应:strings.HasPrefix('', '停止流式的输出:')`

		V155V["多次弹窗"] = `本套框架最繁杂部分,为此放弃不稳定的'SSE' 和 'WebSocket'.&#10;前端html: 多次更新弹出弹窗网页内容,等待选择.&#10;<div id='C3_str'> </div>&#10;JS:接收后端数据, if (data.C3_str !== '') {&#10;C3_str.innerHTML = data.C3_str;&#10;C2_str.style.display = 'flex';//显示&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;后端接收并响应:(struct).C3_str='等待选择的弹窗源码' &#10; (struct).M9_str='多次等待弹窗选择判断:'&#10; strings.HasPrefix('', '多次等待弹窗选择判断:')`
		V155V["单选选择框"] = `前端html: <input type='radio' name='choice'  value='A'>&#10;JS:统一响应【Change事件】统一处理 单选与复选元素'change',&#10;处理单选框 else if (target.type === 'radio' && target.name === 'choice')&#10;后端接收并响应:strings.HasPrefix('', '单选值:')`
		V155V["复选选择框"] = `前端html: <input type='checkbox' class='toggle-item'  value='A' >&#10;JS:统一响应【Change事件】统一处理 单选与复选元素'change',&#10;处理复选框  if (target.type === 'checkbox' && target.classList.contains('toggle-item'))&#10;后端接收并响应:strings.HasPrefix('', '复选值:')`
		V155V["下拉列表"] = `前端html: <select name='format' class='form-select  change-value' >&#10;<option value='推理等级:关'></option>&#10;JS:统一响应下拉表元素 change 元素'change',检查是否是目标下拉列表&#10; if (event.target.classList.contains('change-value') && event.target.tagName === 'SELECT')&#10;后端接收并响应:strings.HasPrefix('', '推理等级:')`
		V155V["输入表单"] = `前端html: <form class='ajax-form' method='POST' action='#'>&#10;<input name='fullName'>&#10;JS:统一响应【Submit事件】表单提交 元素'submit',使用 .value 获取各个字段的值&#10; const name = form.querySelector('[name='fullName']').value; &#10;后端接收并响应:strings.HasPrefix('', '多项输入表单提交:')`
		V155V["打开目录提示"] = `前端html: class='Dataset-Text' data-text='打开目录:'&#10;后端接收并响应:strings.HasPrefix('', '打开目录:')&#10;测试安全无误报,打开本地电脑的目录或文件`
		V155V["赞赏提示"] = "前端html: class='Dataset-Text' data-text='赞赏:'&#10;后端接收并响应:strings.HasPrefix('', '赞赏:')&#10;测试展示,打包文件的图片或网络图片"
		V155V["退出程序"] = "前端html: class='Dataset-Text' data-text='退出:'&#10;后端接收并响应:strings.HasPrefix('', '退出:')&#10;先展示'退出程序'提示,程序最后才关闭退出."

	} else {

		V155V["输入按钮"] = `Frontend HTML: class='Dataset-Text' data-text='输入内容:'&#10;JS: Unified response to click on element '.Dataset-Text', while checking INPUT and TEXTAREA elements &#10;Backend receives and responds: strings.HasPrefix('', '输入内容:')`
		V155V["流式输出按钮"] = `Frontend HTML: Asynchronously updates frontend web content continuously, reducing waiting time for AI responses, adding dynamic feel.&#10;<textarea id='L1_str'> </textarea>&#10;JS: Receives backend data, if (data.L1_str !== '') {&#10;L1_str.innerHTML = data.L1_str;&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;Backend receives and responds: (struct).L1_str='AI response content' &#10; (struct).M9_str='流式:'&#10; strings.HasPrefix('', '流式:')`
		V155V["停止输出"] = `Frontend HTML: class='Dataset-Text' data-text='停止流式的输出:'&#10;JS: Unified response to click on element '.Dataset-Text' &#10;Backend receives and responds: strings.HasPrefix('', '停止流式的输出:')`
		V155V["多次弹窗"] = `The most complex part of this framework, therefore abandoning unstable 'SSE' and 'WebSocket'.&#10;Frontend HTML: Multiple updates to popup modal content, waiting for selection.&#10;<div id='C3_str'> </div>&#10;JS: Receives backend data, if (data.C3_str !== '') {&#10;C3_str.innerHTML = data.C3_str;&#10;C2_str.style.display = 'flex';//display&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;Backend receives and responds: (struct).C3_str='popup source code waiting for selection' &#10; (struct).M9_str='多次等待弹窗选择判断:'&#10; strings.HasPrefix('', '多次等待弹窗选择判断:')`
		V155V["单选选择框"] = `Frontend HTML: <input type='radio' name='choice'  value='A'>&#10;JS: Unified response to [Change event] uniformly handling radio and checkbox elements 'change',&#10;Handle radio button else if (target.type === 'radio' && target.name === 'choice')&#10;Backend receives and responds: strings.HasPrefix('', '单选值:')`
		V155V["复选选择框"] = `Frontend HTML: <input type='checkbox' class='toggle-item'  value='A' >&#10;JS: Unified response to [Change event] uniformly handling radio and checkbox elements 'change',&#10;Handle checkbox if (target.type === 'checkbox' && target.classList.contains('toggle-item'))&#10;Backend receives and responds: strings.HasPrefix('', '复选值:')`
		V155V["下拉列表"] = `Frontend HTML: <select name='format' class='form-select  change-value' >&#10;<option value='推理等级:关'></option>&#10;JS: Unified response to select element change event 'change', check if it is the target dropdown&#10; if (event.target.classList.contains('change-value') && event.target.tagName === 'SELECT')&#10;Backend receives and responds: strings.HasPrefix('', '推理等级:')`
		V155V["输入表单"] = `Frontend HTML: <form class='ajax-form' method='POST' action='#'>&#10;<input name='fullName'>&#10;JS: Unified response to [Submit event] form submission element 'submit', use .value to get the values of each field&#10; const name = form.querySelector('[name='fullName']').value; &#10;Backend receives and responds: strings.HasPrefix('', '多项输入表单提交:')`
		V155V["打开目录提示"] = `Frontend HTML: class='Dataset-Text' data-text='打开目录:'&#10;Backend receives and responds: strings.HasPrefix('', '打开目录:')&#10;Tested safe without false positives, open a directory or file on the local computer`
		V155V["赞赏提示"] = `Frontend HTML: class='Dataset-Text' data-text='赞赏:'&#10;Backend receives and responds: strings.HasPrefix('', '赞赏:')&#10;Test display, images from packaged files or network images`
		V155V["退出程序"] = `Frontend HTML: class='Dataset-Text' data-text='退出:'&#10;Backend receives and responds: strings.HasPrefix('', '退出:')&#10;First display 'Exit program' prompt, then the program finally closes and exits.`

	}

	V166V = strings.Replace(V166V, "|替换输入按钮提示|", V155V["输入按钮"], -1)
	V166V = strings.Replace(V166V, "|替换Ai流式输出按钮提示|", V155V["流式输出按钮"], -1)
	V166V = strings.Replace(V166V, "|替换停止输出按钮提示|", V155V["停止输出"], -1)

	V166V = strings.Replace(V166V, "|替换多次弹窗提示|", V155V["多次弹窗"], -1)
	V166V = strings.Replace(V166V, "|替换单选选择框提示|", V155V["单选选择框"], -1)
	V166V = strings.Replace(V166V, "|替换复选选择框提示|", V155V["复选选择框"], -1)
	V166V = strings.Replace(V166V, "|替换下拉列表提示|", V155V["下拉列表"], -1)
	V166V = strings.Replace(V166V, "|替换输入表单提示|", V155V["输入表单"], -1)
	V166V = strings.Replace(V166V, "|替换打开目录提示|", V155V["打开目录提示"], -1)
	V166V = strings.Replace(V166V, "|替换赞赏提示|", V155V["赞赏提示"], -1)
	V166V = strings.Replace(V166V, "|替换退出程序提示|", V155V["退出程序"], -1)

	return V166V
}

func F18F() (V168V string) {

	V168V = ` 	   <div class="dropdown-container">
		<span class="dropdown-L0" > English - Please use the browser's built-in translation <i class="fas fa-chevron-down"></i></span>
		<div class="dropdown-menu" style="top:100%;min-height:550px;">
		<span  class="chat-item-C">English - Please use the browser's built-in translation</span>
		<span  class="chat-item-C" >简体.请用浏览器内置的翻译功能来翻译程序的网页界面。</span>
		<span  class="chat-item-C">Español - Utilice la traducción del navegador</span>
		<span  class="chat-item-C" >Français - Utilisez la traduction du navigateur</span>
		<span  class="chat-item-C" >Deutsch - Browserübersetzung verwenden</span>
		<span  class="chat-item-C" >日本語 - ブラウザの翻訳機能をご利用ください</span>
		<span  class="chat-item-C" >繁體.請用瀏覽器自帶的翻譯，翻譯程序網頁界面。</span>
		<span  class="chat-item-C">한국어 - 브라우저 번역 기능을 사용하세요</span>
		<span  class="chat-item-C" >Русский - Используйте встроенный перевод</span>
		<span  class="chat-item-C">العربية - استخدم ترجمة المتصفح المدمجة</span>
		<span  class="chat-item-C" >Português - Use a tradução do navegador</span>
		<span  class="chat-item-C" >Italiano - Usa la traduzione del browser</span>
		<span  class="chat-item-C">Nederlands - Gebruik browsertranslatie</span>
		<span  class="chat-item-C">Svenska - Använd webbläsaröversättning</span>
		<span  class="chat-item-C">Türkçe - Tarayıcı çevirisini kullanın</span>
		<span  class="chat-item-C" >हिन्दी - ब्राउज़र अनुवाद का उपयोग करें</span>
		<span  class="chat-item-C">ไทย - ใช้การแปลของเบราว์เซอร์</span>
		<span  class="chat-item-C" >Tiếng Việt - Dùng tính năng dịch trình duyệt</span>
		<span  class="chat-item-C">Polski - Użyj tłumaczenia przeglądarki</span>
		<span  class="chat-item-C">Українська - Використовуйте переклад браузера</span>
		<span  class="chat-item-C">Ελληνικά - Χρησιμοποιήστε μετάφραση προγράμματος περιήγησης</span>
		<span  class="chat-item-C">עברית - השתמש בתרגום דפדפן מובנה</span>
		<span  class="chat-item-C">فارسی - از ترجمه مرورگر استفاده کنید</span>
		<span  class="chat-item-C">اردو - براؤزر ترجمہ کا استعمال کریں</span>
		<span  class="chat-item-C">Čeština - Použijte překlad prohlížeče</span>
		<span  class="chat-item-C" >Magyar - Használja a böngészőfordítót</span>
		<span  class="chat-item-C">Suomi - Käytä selaimen käännöstä</span>
		<span  class="chat-item-C">Norsk - Bruk nettleseroversettelse</span>
		<span  class="chat-item-C">Dansk - Brug browseroversættelse</span>
		<span  class="chat-item-C" >Română - Utilizați traducerea browserului</span>
		<span  class="chat-item-C">Bahasa Indonesia - Gunakan terjemahan browser</span>
		<span  class="chat-item-C">Bahasa Melayu - Gunakan terjemahan pelayar</span>
		<span  class="chat-item-C" >Filipino - Gamitin ang browser translation</span>
		<span  class="chat-item-C">বাংলা - ব্রাউজার অনুবাদ ব্যবহার করুন</span>
		<span  class="chat-item-C">தமிழ் - உலாவி மொழிபெயர்ப்பைப் பயன்படுத்தவும்</span>
		<span  class="chat-item-C">తెలుగు - బ్రౌజర్ అనువాదాన్ని ఉపయోగించండి</span>
		<span  class="chat-item-C" >Kiswahili - Tumia tafsiri ya kivinjari</span>
		<span  class="chat-item-C" >नेपाली - ब्राउजर अनुवाद प्रयोग गर्नुहोस्</span>
		<span  class="chat-item-C">සිංහල - බ්රවුසර් පරිවර්තනය භාවිතා කරන්න</span>
		<span  class="chat-item-C" >ខ្មែរ - ប្រើការបកប្រែរបស់កម្មវិធីរុករក</span>
		<span  class="chat-item-C">မြန်မာ - ဘရောက်ဆာ၏ဘာသာပြန်ချက်ကိုသုံးပါ</span>
		<span  class="chat-item-C" >አማርኛ - የማሰሻገሪያ አሞሌ ይጠቀሙ</span>
		<span  class="chat-item-C">Zulu - Sebenzisa ukuhumusha kwebrowser</span>
		<span  class="chat-item-C" >Xhosa - Sebenzisa uthelekiso lwebrowser</span>
		<span  class="chat-item-C">Chichewa - Gwiritsani ntchito kumasulira browser</span>
		<span  class="chat-item-C">Shona - Shandisa browser translation</span>
		<span  class="chat-item-C" >Sesotho - Sebedisa phetolelo ya browser</span>
		
	</div>
		</div>

	`

	return
}

func F19F() (V170V string) {
	V170V = `			  <style type='text/css'>
	body {
		font-family: 'Microsoft YaHei', 'PingFang SC', sans-serif;
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
		min-height: 100vh;
		margin: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 16px;
		box-sizing: border-box;
	}
	.container {
		background: rgba(240, 236, 236, 0.95);
		border-radius: 24px;
		padding: 30px;
		max-width: 68%;   
		width: 100%;
		box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
		backdrop-filter: blur(10px);
	}
	.input-group {
		display: flex;
		gap: 15px;
		margin-bottom: 30px;
	}
	.input-field {
		flex: 1;
		padding: 15px 20px;
		border: 2px solid #e0e0e0;
		border-radius: 12px;
		font-size: 16px;
		transition: all 0.3s ease;
		outline: none;
		box-sizing: border-box;
	}
	.input-field:focus {
		border-color: #667eea;
		box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
	}
	.btn {
		padding: 15px 35px;
		border: none;
		border-radius: 12px;
		font-size: 16px;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.3s ease;
		box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
		white-space: nowrap; 
	}
	.btn-primary {
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
		color: white;
	}
	.btn-primary:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
		background: linear-gradient(135deg, #fa709a 0%, #fc8d94 100%);
	}
	.btn-secondary {
		background: white;
		color: #667eea;
		border: 2px solid #667eea;
	}
	.btn-secondary:hover {
		background: #667eea;
		color: white;
	}
	.radio-group-title {
		font-size: 18px;
		font-weight: 600;
		color: #334155;
		margin-bottom: 16px;
		display: block;
	}
	.radio-options {
		display: flex;
		gap: 20px;
		flex-wrap: wrap;
	}
	.radio-item {
		display: flex;
		align-items: center;
		cursor: pointer;
		position: relative;
	}
	.radio-item input[type="radio"] {
		appearance: none;
		width: 24px;
		height: 24px;
		border: 2px solid #cbd5e1;
		border-radius: 50%;
		margin-right: 10px;
		cursor: pointer;
		transition: all 0.3s ease;
		position: relative;
	}
	.radio-item input[type="radio"]:checked {
		border-color: #667eea;
		background: #667eea;
	}
	.radio-item input[type="radio"]:checked::after {
		content: '';
		position: absolute;
		width: 10px;
		height: 10px;
		background: white;
		border-radius: 50%;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
	}
	.radio-item label {
		font-size: 16px;
		color: #475569;
		cursor: pointer;
		user-select: none;
	}
	.action-section {
		background: linear-gradient(135deg, #f0f4ff 0%, #fef3f4 100%);
		padding: 24px;
		border-radius: 16px;
		margin-bottom: 30px;
		border: 2px solid #e0e7ff;
	}
	.action-section h3 {
		font-size: 18px;
		color: #334155;
		margin-bottom: 16px;
		font-weight: 600;
	}
	.button-row {
		display: grid;
		grid-template-columns: 1fr 1fr; 
		gap: 10px;
		margin-bottom: 10px;
	}
	.stream-button {
		color: white;
		padding: 18px;
		border-radius: 12px;
		text-align: center;
		font-size: 16px;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.3s ease;
		margin-bottom: 30px;
		box-shadow: 0 4px 15px rgba(48, 207, 208, 0.3);
		word-break: keep-all;
	}
	.stream-button:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(48, 207, 208, 0.4);
		background: linear-gradient(135deg, #fa709a 0%, #fc8d94 100%);
	}
	 .s1 {
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	}
	 .s2 {
		background: linear-gradient(135deg, #fa709a 0%, #9b8922 100%);
	}
	.s3 {
		background: linear-gradient(135deg, #30cfd0 0%, #330867 100%);
	}
	.textarea-panel {
		background: #f8f9fa;
		border: 2px solid #e0e0e0;
		border-radius: 12px;
		padding: 20px;
		min-height: 100px;
		font-size: 16px;
		color: #666;
		line-height: 1.6;
		resize: vertical;
		width: 100%;
		outline: none;
		transition: all 0.3s ease;
		box-sizing: border-box;
	}
	.textarea-panel:focus {
		border-color: #667eea;
		box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
	}
	@media (max-width: 600px) {
		.container {
			padding: 20px;
		}
		.btn {
			padding: 12px 20px;
			font-size: 15px;
		}
		.stream-button {
			padding: 16px;
			font-size: 16px;
		}
		.button-row {
			gap: 15px;
		}
		.radio-options {
			gap: 20px;
		}
	}
	@media (max-width: 480px) {
		.input-group {
			flex-direction: column;
		}
		.btn {
			width: 100%;
		}
		.button-row {
			grid-template-columns: 1fr;  
		}
		.radio-options {
			flex-direction: column;
			gap: 12px;
		}
		.stream-button {
			font-size: 15px;
			padding: 14px;
		}
	}
	  #toast {
		position: fixed;
		top: 20px;
		left: 50%;
		transform: translateX(-50%);
		background: rgba(230, 245, 225, 0.9);
		padding: 12px 24px;
		border-radius: 30px;
		opacity: 0;
		transition: opacity 0.3s ease;
		z-index: 1000;
		color: #555;
		font-weight: 500;
		border: 1px solid rgba(150, 180, 145, 0.3);
		box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
		background: linear-gradient(135deg, #1a2a6c, #2a4365, #0f4c75);
		color: white;
	}
	#toast.active {
		opacity: 1;
	}
	 .tishi-note-box {
		background: rgba(255, 255, 255, 0.1);
		border-radius: 8px;
		padding: 15px;
		margin: 20px 0;
		border: 1px solid rgba(76, 201, 240, 0.3);
	}
	.tishi-note-text1 {
		margin-top: 0;
		color: #ffd166;
		font-size: 15px;
	}
	.tishi-note-text2 {
		margin-top: 0;
		color: #00b42a;
		font-size: 15px;
	}
	.tishi-icon {
		display: inline-block;
		width: 38px;
		height: 38px;
		margin-right: 10px;
		vertical-align: middle;
		background-color: #ffd166;
		color: white;
		border-radius: 50%;
		text-align: center;
		line-height: 32px;
		font-size: 0.9rem;
	}
	 .tishi-success-icon {
		width: 38px;
		height: 38px;
		border-radius: 50%;
		background-color: #4caf50; 
		color: white; 
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 32px;
		font-weight: bold;
	}
	.dynamic-form {
		margin-bottom: 7px;
	}
	.form-select {
		width: 100%;
		padding: 12px 15px;
		border-radius: 12px;
		color: #333;
		font-size: 15px;
		appearance: none;
		background-repeat: no-repeat;
		background-position: right 15px center;
		background-size: 16px;
		background: rgba(238, 241, 237, 0.6);
		border: 1px solid rgba(160, 200, 155, 0.8);
		backdrop-filter: blur(5px);
		-webkit-backdrop-filter: blur(5px);
	}
	.form-select:focus {
		outline: none;
		border-color: #8ab661;
		box-shadow: 0 0 0 2px rgba(138, 182, 97, 0.2);
	}     
	.form-select:hover {
		box-shadow: 0 0 0 2px #254e8b;
		border-color: #506abe;
	}
 #C2_str {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.7);
		display: none;
		align-items: center;
		justify-content: center;
		z-index: 9999;
	}
	#C3_str {
	background: linear-gradient(135deg, #667eea 0%, #66569e 100%);
	padding: 20px 30px;
	border-radius: 6px;
	box-shadow: 0 4px 12px rgba(0,0,0,.3);
	text-align: center;
	}
		 #C5_str {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.7);
		display: none;
		align-items: center;
		justify-content: center;
		z-index: 9999;
	}
	#C6_str {
		background: linear-gradient(135deg, #1a2a6c, #2a4365, #0f4c75);
		padding: 20px 30px;
		border-radius: 6px;
		box-shadow: 0 4px 12px rgba(0,0,0,.3);
		text-align: center;
		pointer-events: auto;
	}
		  .tishi-divider {
		height: 1px;
		background: rgba(144, 224, 239, 0.3);
		margin: 15px 0;
	}
.dropdown-container {
	position: relative;
}
.dropdown-container:hover .dropdown-menu{
	opacity: 1;
	visibility: visible;
	transform: translateY(0);
}
 .dropdown-container:hover .dropdown-menu2{
	opacity: 1;
	visibility: visible;
	transform: translateY(0);
}
.dropdown-menu {
	bottom: 100%;
	min-Height : 250px;  
	min-width: 510px;
	position : absolute;  
	 left: 0;
	border-radius: 12px;
	opacity: 0;
	visibility: hidden;
	transform: translateY(10px);
	transition: all 0.3s ease;
	box-shadow: 0 10px 30px rgba(0,0,0,0.15);
	z-index: 10;
	max-height: 550px;
	overflow-y: auto;
	 background: rgba(219, 231, 217, 0.8);
border: 1px solid rgba(160, 200, 155, 0.8);
backdrop-filter: blur(5px);
-webkit-backdrop-filter: blur(5px);
}
 .dropdown-L0 {
	width: 100%;
	padding: 5px 8px;
	border: 1px solid rgba(141, 136, 136, 0.2);
	border-radius: 8px;
	color: rgb(53, 47, 47);
	font-size: 15px;
	font-weight: 300;
	cursor: pointer;
	display: flex;
	justify-content: space-between;
	align-items: center;
	transition: all 0.3s ease;
	backdrop-filter: blur(5px);
   background: linear-gradient(135deg, #8897da 0%, #ad94c5 100%);
color: #2d4d2a; 
border-bottom: 1px solid rgba(160, 200, 155, 0.8);
}
.dropdown-L0:hover {
	background: rgba(255, 255, 255, 0.25);
	transform: translateY(-2px);
	box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}
.dropdown-L0 i {
	transition: transform 0.3s ease;
}
.dropdown-container:hover .dropdown-L0 i {
	transform: rotate(180deg);
}
.chat-item-C {
width: 150% !important; 
box-sizing: border-box; 
padding: 1px 3px;
font-size: 13px;
margin: 5px 0;
border-radius: 8px;
min-width: 250px; 
cursor: pointer;
transition: all 0.3s ease;
border: 1px solid rgba(150, 180, 145, 0.1);
color: #333;
display: block; 
max-width: 100%; 
}
.chat-item-C:hover {
background: rgba(5, 88, 243, 0.87);
color: rgb(235, 227, 227);
transform: translateX(3px);
}
	.multi-select-section {
	font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
	max-width: 520px;
	margin: 24px auto;
	background: #e6e1e1;
	border-radius: 28px;
	box-shadow: 0 20px 35px -8px rgba(0, 0, 0, 0.08), 0 8px 18px -6px rgba(0, 0, 0, 0.02), 0 0 0 1px rgba(0, 0, 0, 0.02);
	padding: 28px 24px;
	transition: box-shadow 0.25s ease, transform 0.2s ease;
	}
	.multi-select-section:hover {
	box-shadow: 0 24px 42px -12px rgba(0, 0, 0, 0.12), 0 0 0 1px rgba(0, 0, 0, 0.03);
	}
	.checkbox-group-title {
	display: flex;
	align-items: center;
	font-size: 1.3rem;
	font-weight: 600;
	color: #0a1a2f;
	margin-bottom: 1.5rem;
	letter-spacing: -0.02em;
	border-bottom: 2px solid #bac4d4;
	padding-bottom: 14px;
	}
	.checkbox-group-title::after {
	content: '';
	flex: 1;
	height: 2px;
	background: linear-gradient(90deg, #c3cbd6, transparent);
	margin-left: 16px;
	}
	.checkbox-options {
	display: flex;
	flex-direction: column;
	gap: 14px;
	}
	.checkbox-item {
	display: flex;
	align-items: center;
	transition: background-color 0.2s;
	padding: 4px 0;
	border-radius: 12px;
	}
	.checkbox-item:hover {
	background-color: #d1dae4;
	}
  .card {
		background: white;
		border-radius: 20px;
		box-shadow: 0 4px 12px rgba(0,0,0,0.05);
		padding: 1.8rem 2rem;
		margin: 2rem 0;
	}
	.note {
		background: #e6f0ff;
		border-left: 5px solid #2563eb;
		padding: 1rem 1.5rem;
		border-radius: 12px;
		margin: 1.5rem 0;
		font-size: 0.95rem;
	}
	.inputB {
		width: 100%;
		padding: 0.7rem 0.9rem;
		margin-top: 0.3rem;
		border: 1px solid #cbd5e1;
		border-radius: 16px;
		font-size: 1rem;
		box-sizing: border-box;
		transition: 0.15s;
	}
	.inputB:focus {
		outline: none;
		border-color: #2563eb;
		box-shadow: 0 0 0 3px rgba(37,99,235,0.2);
	}
	hr {
		border: none; height: 2px;
		background: linear-gradient(90deg, transparent, #667eea, transparent); 
		margin: 2em 0;
	}
	
	.footnote {
		font-size: 0.9rem;
		color: #64748b;
		text-align: center;
		margin-top: 2.5rem;
	}
	.code-snip {
		background: #1e293b;
		color: #bbd7fb;
		padding: 0.2rem 0.6rem;
		border-radius: 8px;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.9rem;
	}
	
	tr:hover {
		background-color: #edf2f7 !important;
	}
	
		
	</style>	`

	return
}

func F20F(V171V, V172V, V173V string) (V174V string) {

	if strings.Contains(V172V, "class=") {

		V4V.M3_str = V172V
	} else if V171V == "成功" {
		V4V.M3_str = F22F(V172V)
	} else if V171V == "错误" {

		V4V.M3_str = F21F(V172V)
	} else {
		V4V.M3_str = F22F(V172V)
	}

	V4V.L1_str = strings.Replace(V173V, "|替换提示窗口|", "", -1)

	return

}

func F21F(V175V string) (V178V string) {
	V178V = `<div >
	<div class="tishi-note-box">
	<h3 class="tishi-note-text2">
	<h3 class="tishi-icon">✗ </h3> 
            <p  class="prompt-content">|替换错误提示|</p>
		</h3>
	</div>	
	</div>	
	`
	V178V = strings.Replace(V178V, "|替换错误提示|", V175V, -1)

	return

}

func F22F(V179V string) (V182V string) {
	V182V = `<div>
	<div class="tishi-note-box">
	<h3 class="tishi-note-text2">
	<h3 class="tishi-success-icon">✔</h3> 
            <p class="prompt-content">|替换提示|</p>
		</h3>
	</div>
	</div>
	`

	V182V = strings.Replace(V182V, "|替换提示|", V179V, -1)

	return

}

func F23F(V183V string) (V185V map[string]string) {
	V185V = make(map[string]string)
	if V2V.Q18Q == "汉语" {

		V185V["打开目录"] = "正在打开目录:"
		V185V["初始欢迎"] = "程序初始,欢迎测试演示。"

		V185V["程序已经退出提示"] = "程序已经退出."
		V185V["展示选择框示例"] = "展示选择框示例."
		V185V["输入内容提示"] = "输入内容提示:<br />"

		V185V["下拉表单示例"] = "下拉表单示例."
		V185V["多项输入表单示例"] = "多项输入表单示例."

		V185V["选择推理等级"] = "选择推理等级:"
		V185V["单选框的选值"] = "单选框的选值:"
		V185V["复选框的选值"] = "复选框的选值:"
		V185V["复选框的取消值"] = "复选框的取消值:"
		V185V["多项输入表单的值"] = "多项输入表单:<br />姓名:%s<br />兴趣:%s"
		V185V["流式结束"] = "流式结束."

	} else {
		V185V["打开目录"] = "Opening directory:"
		V185V["初始欢迎"] = "Program initialized, welcome to the test demo."
		V185V["程序已经退出提示"] = "Program has exited."
		V185V["展示选择框示例"] = "Displaying selection box example."
		V185V["输入内容提示"] = "Input prompt:<br />"
		V185V["下拉表单示例"] = "Dropdown form example."
		V185V["多项输入表单示例"] = "Multi-input form example."
		V185V["选择推理等级"] = "Inference grade:"
		V185V["单选框的选值"] = "Radio button selected value:"
		V185V["复选框的选值"] = "Checkbox selected value:"
		V185V["复选框的取消值"] = "Checkbox deselected value:"
		V185V["多项输入表单的值"] = "Multi-input form:<br />Name:%s<br />Interests:%s"
		V185V["流式结束"] = "Streaming end."

	}

	return

}

func F24F(V186V, V187V int) {
	var V188V int
	var V191V, V192V string

	for {
		V188V++
		if V188V == V186V+1 {

			V2V.Q22Q.Lock()
			V2V.Q19Q = "流式结束"
			V2V.Q22Q.Unlock()

			return
		} else if V187V != V2V.Q7Q {
			return
		}

		V191V = "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "]"

		V192V = fmt.Sprintf("%d ,%s", V188V, V191V)

		V2V.Q16Q = append(V2V.Q16Q, V192V)
		time.Sleep(2 * time.Second)
		if V2V.Q19Q == "流式结束" {

			return
		}

	}

	return

}

func F25F() {
	V2V.Q8Q = ""
	for {

		time.Sleep(300 * time.Millisecond)

		if len(V2V.Q8Q) > 2 {
			return
		}

	}

	return

}

func F26F() {
	if V2V.Q15Q == 0 {
		V4V.C3_str = F27F()
		V4V.M9_str = "多次等待弹窗选择判断:"
		return
	}

	if V2V.Q4Q == "" {
		V4V.C3_str = F28F()
		V4V.M9_str = "多次等待弹窗选择判断:"
		return
	}

	V2V.Q8Q = "等待多选择完成"
	V4V.M9_str = "流式:"

	return

}

func F27F() (V205V string) {

	var V195V = make(map[string]string)
	var V196V, V199V string

	if V2V.Q18Q == "汉语" {

		V195V["纪录栏源码"] =
			` 
		<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> <span class="tishi-highlight">  流式输出</span>  请选择:</h2>
			<hr>
				<div class="dropdown-container">
			<span class="dropdown-L0" style="background: linear-gradient(135deg, #a6bed1 0%, #7d97b4 100%);color: #133025;" >流式输出次数▼ <i class="fas fa-chevron-down"></i></span>
					<div class="dropdown-menu" style="top:100%;min-height:190px;">
							|替换流式输出次数表项|
						
					</div>
				</div>
		</aside>
		`
		V199V = `<span class="chat-item-C Dataset-Text" data-text="流式输出一多项选择:|替换次数|次">|替换次数| 次</span>`

	} else {

		V195V["纪录栏源码"] = `<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> <span class="tishi-highlight"> Streaming output </span > please select:</h2>
			<hr>
				<div class="dropdown-container">
			<span class="dropdown-L0" style="background: linear-gradient(135deg, #a6bed1 0%, #7d97b4 100%);color: #133025;" >Number of streaming outputs ▼ <i class="fas fa-chevron-down"></i></span>
					<div class="dropdown-menu" style="top:100%;min-height:190px;">
							|替换流式输出次数表项|
						
					</div>
				</div>
		</aside>`
		V199V = `<span class="chat-item-C Dataset-Text" data-text="流式输出一多项选择:|替换次数|次">|替换次数| number</span>`

	}

	V205V = V195V["纪录栏源码"]

	var V202V = []string{"5", "10", "15"}

	var V204V []string

	for _, V196V = range V202V {

		V204V = append(V204V, strings.Replace(V199V, "|替换次数|", V196V, -1))
	}

	V205V = strings.Replace(V205V, "|替换流式输出次数表项|", strings.Join(V204V, "\n"), -1)

	return

}

func F28F() (V209V string) {

	var V208V = make(map[string]string)
	if V2V.Q18Q == "汉语" {

		V208V["纪录栏源码"] =
			` 
		<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> 流式输出<span style="color:#133025;"> 对话配置完成 </span> :</h2>
			<hr>
			
			<button data-text="流式输出一多项选择:对话配置完成"  class="stream-button Dataset-Text s3" style="color: #133025; width: 100%;"><span class="tishi-highlight">对话配置完成</span></button>

			
		</aside>
		`

	} else {

		V208V["纪录栏源码"] = `		<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> Streaming output <span style="color:#133025;"> Dialog configuration completed </span> :</h2>
			<hr>
			
			<button data-text="流式输出一多项选择:对话配置完成"  class="stream-button Dataset-Text s3" style="color: #133025; width: 100%;"><span class="tishi-highlight">Dialog configuration completed</span></button>

			
			</aside>`

	}

	V209V = V208V["纪录栏源码"]

	return

}

func F29F() (V224V string) {
	var V215V string

	if V2V.Q18Q == "汉语" {

		V224V = `<div class="action-section">
			<div class="multi-select-section">
			<div class="checkbox-group">
				<span class="checkbox-group-title">多选（点击后查看控制台输出）</span>
				<div class="checkbox-options">
				|替换选项列表|
					
				</div>
			</div>
		</div>
	 
	 </div>`
		V215V = `<div class="checkbox-item"> <input type="checkbox" class="toggle-item"  value="|替换选项|" |替换已勾选|>
	 <label>选项 |替换选项|</label></div>`
	} else {

		V224V = `<div class="action-section">
			<div class="multi-select-section">
			<div class="checkbox-group">
				<span class="checkbox-group-title">Multi-selection (click to view console output)</span>
				<div class="checkbox-options">
				|替换选项列表|
					
				</div>
			</div>
		</div>
	 
	 </div>`
		V215V = `<div class="checkbox-item"> <input type="checkbox" class="toggle-item"  value="|替换选项|" |替换已勾选|>
	 <label>option |替换选项|</label></div>`

	}

	var V217V = []string{"A", "B", "C"}
	var V218V, V222V string
	var V223V []string

	if len(V2V.Q10Q) == 0 {
		for _, V218V = range V217V {
			V2V.Q22Q.Lock()
			V2V.Q10Q[V218V] = ""
			V2V.Q22Q.Unlock()
		}
	}

	for _, V218V = range V217V {

		V222V = strings.Replace(V215V, "|替换选项|", V218V, -1)

		V222V = strings.Replace(V222V, "|替换已勾选|", V2V.Q10Q[V218V], -1)
		V223V = append(V223V, V222V)

	}

	V224V = strings.Replace(V224V, "|替换选项列表|", strings.Join(V223V, "\n"), -1)

	return

}

func F30F() (V240V string) {
	var V230V string

	if V2V.Q18Q == "汉语" {

		V240V = `<!-- 单选项 -->
        <div class="action-section">
            <div class="radio-group">
                <span class="radio-group-title">单选</span>
				<div class="radio-options">
				|替换选项列表|
                    
                </div>
            </div>
		 </div>`
		V230V = `<div class="radio-item"><input type="radio" name="choice"  value="|替换选项|" |替换已勾选|>
		 <label>选项|替换选项|</label></div>`

	} else {

		V240V = `<div class="action-section">
		<div class="radio-group">
			<span class="radio-group-title">Radio</span>
			<div class="radio-options">
			|替换选项列表|
			</div>
		</div>
	 </div>`
		V230V = `<div class="radio-item"><input type="radio" name="choice"  value="|替换选项|" |替换已勾选|>
	<label>option |替换选项|</label></div>`

	}

	var V232V = []string{"A", "B", "C", "D"}
	var V233V, V238V string
	var V239V []string

	for _, V233V = range V232V {

		V238V = strings.Replace(V230V, "|替换选项|", V233V, -1)

		if V233V == V2V.Q9Q {
			V238V = strings.Replace(V238V, "|替换已勾选|", "checked", -1)
		} else {
			V238V = strings.Replace(V238V, "|替换已勾选|", "", -1)
		}
		V239V = append(V239V, V238V)

	}

	V240V = strings.Replace(V240V, "|替换选项列表|", strings.Join(V239V, "\n"), -1)

	return

}

func F31F() (V244V string) {

	if V2V.Q18Q == "汉语" {

		V244V = `<div class="action-section">
		下拉表单示例: 选择推理等级。
		<br /><br /><form class="dynamic-form" action="#">
			<div style="display: flex; align-items: center; gap: 5px;">
				<select name="format" class="form-select  change-value"  style="background-color: #f0f7f0;color: #388e3c; flex: 1; padding: 6px; border-radius: 8px; border: 1px solid #c8e6c9;">
				<option value="" selected disabled>调整推理等级▼:</option>
				<option value="" disabled>当前推理等级: |替换当前推理等级|</option>
				<option value="推理等级:关:关">推理等级:关</option>
				<option value="推理等级:低:低">推理等级:低</option>
				<option value="推理等级:中:中">推理等级:中</option>
				<option value="推理等级:高:高">推理等级:高</option>

				</select> </div> </form> </div>`
	} else {

		V244V = `<div class="action-section"> Dropdown example: Select reasoning program. <br /><br /><form class="dynamic-form" action="#"> 
		<div style="display: flex; align-items: center; gap: 5px;"> 
		<select name="format" class="form-select change-value" style="background-color: #f0f7f0; color: #388e3c; flex: 1; padding: 6px; border-radius: 8px; border: 1px solid #c8e6c9;"> <option value="" selected disabled>Adjust reasoning program ▼:</option> 
		<option value="" disabled>Current reasoning program: |替换当前推理等级|</option>
		 <option value="推理等级:关:Off">Reasoning program: Off</option>
		  <option value="推理等级:低:Low">Reasoning program: Low</option>
		   <option value="推理等级:中:Medium">Reasoning program: Medium</option>
			<option value="推理等级:高:High">Reasoning program: High</option>
			 </select> </div> </form> </div>`

	}

	V244V = strings.Replace(V244V, "|替换当前推理等级|", V2V.Q12Q, -1)
	return

}

func F32F() (V251V string) {

	if V2V.Q18Q == "汉语" {

		V251V = ` <div class="action-section">
		<form class="ajax-form" method="POST" action="#">
			<label for="name">姓名</label>
			<input type="text"  class="inputB" name="fullName" placeholder="例如：李小萌" |替换姓名值|>
	
			<label for="interest">感兴趣的领域</label>
			<input type="text" class="inputB" name="interest" placeholder="前端 / 交互 / 设计"  |替换兴趣值|>
			<br />
			<button  class="btn btn-primary"  type="submit">📨 提交 (ajax-form) </button>
		</form></div> `
	} else {

		V251V = `<div class="action-section">
		<form class="ajax-form" method="POST" action="#">
			<label for="name">Full Name</label>
			<input type="text" class="inputB" name="fullName" placeholder="e.g., Li Xiaomeng" |替换姓名值|>
	
			<label for="interest">Areas of Interest</label>
			<input type="text" class="inputB"  name="interest" placeholder="Frontend / Interaction / Design"  |替换兴趣值|>
			<br />
			<button  class="btn btn-primary"  type="submit">📨 Submit (ajax-form) </button>
		</form>
		</div>`

	}

	if V2V.Q21Q == "" {
		V251V = strings.Replace(V251V, "|替换姓名值|", "", -1)
		V251V = strings.Replace(V251V, "|替换兴趣值|", "", -1)

	} else {
		V251V = strings.Replace(V251V, "|替换姓名值|", "value="+V2V.Q21Q, -1)
		V251V = strings.Replace(V251V, "|替换兴趣值|", "value="+V2V.Q20Q, -1)

	}

	return

}

func F33F() (V254V string) {

	if V2V.Q18Q == "汉语" {

		V254V = `				<div style="max-width: 110%;">
		<h2 style="font-size: 2.5em; font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: 一款 MIT 开源的轻量级桌面 GUI 框架 🎉</h2>
		
		
		
		<blockquote style="background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%); border-left: 5px solid #667eea; padding: 1em 1.5em; margin: 1.5em 0; border-radius: 0 8px 8px 0; color: #4a5568; font-style: italic;"><p style="margin: 0; line-height: 1.6;">💡 <strong style="color: #2d3748; font-weight: 700;">核心理念：极速开发 · 极致体积 · 原生性能 · 助力打造用户愿意付费的优质工具</strong></p></blockquote>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		在桌面端，用户对高效、实用工具的需求从未减弱，并且拥有强烈的付费意愿。<br />		 
		PC-Gui 旨在帮助开发者快速响应这一市场需求，用最简单、最稳定的技术，构建出小巧而强大的商业级桌面应用。
		</p>
		
		<hr>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">核心技术栈</h3>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		摒弃了复杂的依赖，臃肿的第三方 GUI 库，回归编程的本质：<strong style="color: #2d3748; font-weight: 700;">用后端思维构建桌面应用</strong>。
		<br />通过一个稳定的 Go 后端提供 Web 服务，驱动一个灵活的 Web 前端界面，实现了无与伦比的轻量化与性能。
		</p>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			<tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">组件</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">技术详情</th>
			</tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">后端服务</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 语言，基于标准库 <code style="background: #edf2f7; color: #e53e3e; padding: 0.2em 0.4em; border-radius: 4px; font-family: \'Consolas\', monospace; font-size: 0.9em;">net/http</code> 提供本地 Web 服务。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">前端界面</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">HTML, JavaScript, CSS 标准 Web 技术。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">数据存储</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">本地加密的 SQLite 数据库，轻量、可靠。</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">核心优势 & 多方案对比</h3>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			<tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">类别</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">✅ PC-Gui 优势</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">⚠️ 其他方案对比</th>
			</tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🚀 零依赖运行</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">后端Go 语言</strong>极速开发，强类型易于维护；交叉编译，生成单一可执行文件，无需用户安装任何运行时或依赖库，双击即可运行。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️需要用户预装并配置  WebView2, Python、C++, Node.js 等复杂的环境和依赖。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🎨界面技术 (HTML)</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">HTML</strong> 前端界面可借助海量模板与 AI 工具快速生成，不仅效率极高，还能轻松打造出精美、现代的视觉风格。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">传统 GUI 库界面通常较为陈旧，自定义难度高。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">AI 流式输出</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">仅需简单的异步处理，即可实现 AI 内容的流式输出，提升用户体验。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">实现流式输出通常需要处理复杂的回调或多线程。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Markdown 渲染</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">完美渲染 AI 返回的 Markdown 格式，并支持各类语言的语法高亮。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Chatbox、Cherry等对 Markdown 渲染及代码高亮效果较为朴素。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">单文件部署</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">利用 Go 标准库中的 embed，可以将所有静态资源（如图片、CSS 等）直接打包到单一可执行文件中，并复用 HTML 服务进行直接访问。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">依赖臃肿:需借助外部工具打包，产物体积庞大或文件零散，部署复杂。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">📦 极致体积</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">应用打包后体积仅 <strong style="color: #2d3748; font-weight: 700;">10-25MB</strong>，分发和下载毫无压力。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ 基于 Electron / WebView2 的应用体积普遍在 <strong style="color: #2d3748; font-weight: 700;">100MB</strong> 以上。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🧠 超低内存占用</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">运行时内存占用仅约 <strong style="color: #2d3748; font-weight: 700;">8MB</strong>，CPU 开销近乎为零，轻快如飞。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Electron / WebView2 应用内存占用轻松达到 <strong style="color: #2d3748; font-weight: 700;">500MB</strong> 以上。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">代码安全性</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 编译后的二进制文件,结合 garble 混淆技术，有效防止逻辑被反编译。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">易泄露:Python、Node.js 脚本语言极易被反编译、扒光，毫无商业机密。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💻跨平台兼容</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 语言原生支持 Windows 7/10/11, Linux, macOS，覆盖最广泛的用户群体。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Webview2 等方案不支持 Windows 7 等旧版系统。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">运行稳定性</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">核心仅依赖 Go 官方标准库，可长期稳定运行不崩溃。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">依赖大量第三方库，版本兼容性和稳定性存在风险。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💯 完全掌控</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">核心代码仅依赖 Go 官方标准库，<strong style="color: #2d3748; font-weight: 700;">无任何第三方 GUI 框架黑盒</strong>，代码完全自主可控，便于长期维护与排查问题。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ 依赖大型第三方框架，代码冗余，遇到疑难杂症时排查困难。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🌐 全球化支持</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">界面基于标准网页，可直接利用浏览器的<strong style="color: #2d3748; font-weight: 700;">内置翻译功能</strong>，轻松支持全球数百种语言。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">需要内置多语言文本库，工作量巨大。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💡跨语言复用</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">框架思路清晰，任何支持 HTTP 服务的语言（如 C#, Python, Rust）均可借鉴。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">框架与特定语言或平台深度绑定，难以迁移。</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">致开发者</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		在 AI 浪潮席卷全球、就业市场面临挑战的今天，掌握一门能够快速创造价值的技能至关重要。
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		希望 PC-Gui 这套轻量、高效的框架，能成为您手中的利器，帮助您快速将创意变为现实，开发出用户愿意付费的桌面实用工具，最终实现实现个人价值与商业创收。
		</p>
		
		<hr>
		
		
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 MIT 开源授权许可</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		本项目采用 <strong style="color: #2d3748; font-weight: 700;">MIT 许可证</strong>。这意味着您可以完全自由地使用框架源码。
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		您可以：(通俗解释 )<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">商业使用</strong>：允许将本作品及其衍生品用于商业目的，并进行销售。<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">修改分发</strong>：允许修改代码，并以开源或闭源的形式重新分发。<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">私人使用</strong>：允许在私人场合使用和修改。<br />
		</p>
		
		<hr>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<strong style="color: #2d3748; font-weight: 700;">唯一要求</strong>：在您的软件副本或重要部分中，包含原始的版权和许可声明。
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">许可原文 (Official License Text)</h3>
		
		<p>
		<pre><code>		 
		MIT License
	
		Copyright (c) [2026] [github/jiqi136]
	
		Permission is hereby granted, free of charge, to any person obtaining a copy
		of this software and associated documentation files (the "Software"), to deal
		in the Software without restriction, including without limitation the rights
		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
		copies of the Software, and to permit persons to whom the Software is
		furnished to do so, subject to the following conditions:
		The above copyright notice and this permission notice shall be included in all
		copies or substantial portions of the Software.
		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
		SOFTWARE.
		</code></pre>
		</p>
			</div>
			<hr>
			`

	} else {

		V254V = `									<div style="max-width: 110%;">
		<h2 style="font-size: 2.5em; font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: An MIT-licensed open-source lightweight desktop GUI framework 🎉</h2>
		
		
		
		<blockquote style="background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%); border-left: 5px solid #667eea; padding: 1em 1.5em; margin: 1.5em 0; border-radius: 0 8px 8px 0; color: #4a5568; font-style: italic;"><p style="margin: 0; line-height: 1.6;">💡 <strong style="color: #2d3748; font-weight: 700;">Core Philosophy: Rapid Development · Minimal Size · Native Performance · Empowering Developers to Build Premium Tools Users Are Willing to Pay For</strong></p></blockquote>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		On the desktop, users' demand for efficient, practical tools has never waned, and they have a strong willingness to pay.<br />		 
		PC-Gui aims to help developers quickly respond to this market need, using the simplest and most stable technologies to build compact yet powerful commercial-grade desktop applications.
		</p>
		
		<hr>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">Core Tech Stack</h3>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		Abandoning complex dependencies and bloated third-party GUI libraries, it returns to the essence of programming: <strong style="color: #2d3748; font-weight: 700;">Building desktop applications with a backend mindset</strong>.
		<br />By providing web services through a stable Go backend that drives a flexible web frontend, it achieves unparalleled lightweight and performance.
		</p>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			 <tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Component</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Technical Details</th>
			 </tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Backend Service</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go language, providing local web services based on the standard library <code style="background: #edf2f7; color: #e53e3e; padding: 0.2em 0.4em; border-radius: 4px; font-family: \'Consolas\', monospace; font-size: 0.9em;">net/http</code>.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Frontend Interface</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Standard web technologies: HTML, JavaScript, CSS.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Data Storage</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Locally encrypted SQLite database, lightweight and reliable.</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">Core Advantages & Multi-Solution Comparison</h3>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			 <tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Category</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">✅ PC-Gui Advantages</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">⚠️ Other Solutions Comparison</th>
			 </tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🚀 Zero-Dependency Runtime</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Backend Go language</strong> enables rapid development, strong typing for easy maintenance; cross-compilation generates a single executable file, requiring no runtime or dependency installation from users—just double-click to run.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Requires users to pre-install and configure complex environments and dependencies such as  WebView2, Python, C++, Node.js, etc.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🎨 Interface Technology (HTML)</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">HTML</strong> frontend interface can be rapidly generated using numerous templates and AI tools, offering extremely high efficiency and making it easy to create beautiful, modern visual styles.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Traditional GUI library interfaces are often outdated and difficult to customize.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">AI Streaming Output</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">With simple asynchronous handling, AI content streaming can be achieved, enhancing user experience.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Implementing streaming output typically requires dealing with complex callbacks or multithreading.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Markdown Rendering</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Perfectly renders Markdown format returned by AI and supports syntax highlighting for various languages.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Chatbox, Cherry, etc., have relatively basic Markdown rendering and code highlighting.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Single-File Deployment</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">By using the embed package in the GoLang standard library, all static resources (such as images, CSS, etc.) can be directly packaged into a single executable file, while reusing the HTML service for direct access.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Bloated dependencies: requires external tools for packaging, resulting in large output size or scattered files, making deployment complex.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">📦 Minimal Size</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The packaged application size is only <strong style="color: #2d3748; font-weight: 700;">10-25MB</strong>, making distribution and download effortless.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Applications based on Electron / WebView2 typically exceed <strong style="color: #2d3748; font-weight: 700;">100MB</strong>.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🧠 Ultra-Low Memory Usage</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Runtime memory usage is only about <strong style="color: #2d3748; font-weight: 700;">8MB</strong>, with near-zero CPU overhead, fast and light.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Electron / WebView2 applications easily exceed <strong style="color: #2d3748; font-weight: 700;">500MB</strong> of memory usage.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Code Security</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go compiled binaries, combined with garble obfuscation technology, effectively prevent logic from being decompiled.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Vulnerable to leaks: scripting languages like Python and Node.js are easily decompiled and exposed, with no trade secret protection.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💻 Cross-Platform Compatibility</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go language natively supports Windows 7/10/11, Linux, macOS, covering the widest range of users.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Solutions like Webview2 do not support older systems such as Windows 7.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">Runtime Stability</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The core only depends on the Go official standard library, ensuring long-term stable operation without crashes.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Relies on a large number of third-party libraries, posing risks in version compatibility and stability.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💯 Full Control</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Core code only depends on the Go official standard library, <strong style="color: #2d3748; font-weight: 700;">with no third-party GUI framework black boxes</strong>, making the code fully autonomous and controllable, facilitating long-term maintenance and troubleshooting.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Relies on large third-party frameworks, leading to code redundancy and difficulty in troubleshooting complex issues.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">🌐 Globalization Support</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The interface is based on standard web pages and can directly leverage the browser's <strong style="color: #2d3748; font-weight: 700;">built-in translation feature</strong>, easily supporting hundreds of languages worldwide.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Requires built-in multilingual text libraries, a massive amount of work.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #2d3748; font-weight: 700;">💡 Cross-Language Reusability</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The framework concept is clear and can be adopted by any language that supports HTTP services (e.g., C#, Python, Rust).</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Frameworks deeply tied to specific languages or platforms are difficult to migrate.</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">To Developers</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		In today's world, where the AI wave is sweeping the globe and the job market faces challenges, it is crucial to master a skill that can quickly create value.
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		We hope that PC-Gui, this lightweight and efficient framework, will become a powerful tool in your hands, helping you quickly turn ideas into reality, develop desktop utilities that users are willing to pay for, and ultimately achieve personal value and commercial revenue.
		</p>
		
		<hr>
		
		
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 MIT Open Source License</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		This project is licensed under the <strong style="color: #2d3748; font-weight: 700;">MIT License</strong>. This means you are completely free to use the framework's source code.
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		You can: (plain English explanation)<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">Commercial Use</strong>: Use the work and its derivatives for commercial purposes and sell them.<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">Modify and Distribute</strong>: Modify the code and redistribute it under open-source or closed-source licenses.<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">Private Use</strong>: Use and modify the code for personal purposes.<br />
		</p>
		
		<hr>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		<strong style="color: #2d3748; font-weight: 700;">Only Requirement</strong>: Include the original copyright and license notice in your software copies or substantial portions.
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">Official License Text</h3>
		
		<p>
		<pre><code>		  
		MIT License
		
		Copyright (c) [2026] [github/jiqi136]
		
		Permission is hereby granted, free of charge, to any person obtaining a copy
		of this software and associated documentation files (the "Software"), to deal
		in the Software without restriction, including without limitation the rights
		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
		copies of the Software, and to permit persons to whom the Software is
		furnished to do so, subject to the following conditions:
		
		The above copyright notice and this permission notice shall be included in all
		copies or substantial portions of the Software.
		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
		SOFTWARE.
		</code></pre>
		</p>
			</div>
			<hr>`

	}

	return

}
