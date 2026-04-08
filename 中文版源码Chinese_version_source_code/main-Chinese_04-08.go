//================================ 版本 version  1.2 ===================================================
// GOlang V1.20 版本,安装与配置 ，请提问Ai
//  GOlang  V1.20 version, installation and configuration, please ask AI.

//==================================  设置与目录 Configuration and Directories    ===================
// 程序各项配置 ： = "英语English"
// Program configuration ： = "英语English"
// 程序响应的功能目录 ： "初始:")
//Program Response Function Directory  ： "初始:")

//===================================================================================

package main

import (
	"encoding/json"

	"bufio"

	"fmt"
	"math/rand"
	"path/filepath"

	"net"
	"strings"
	"syscall"
	"time"

	"io/ioutil"
	"os"

	"net/http"

	"log"

	"os/exec"

	"strconv"

	"runtime"
	"sync"
)

type 定义一全局变量一一变量结构体 struct {
	程序英汉版一文本    string
	当前目录一路径文本   string
	程序名与版本一文本   string
	双向通信计数器一整数  int
	双向再次通信序号一文本 string

	程序界面跳转网址一文件路径 string
	本地网页网址一文本     string

	流式输出内容一列表     []string
	流式终断一文本       string
	上次流式列表数一整数    int
	等待多选择完成一文本    string
	流式输入次数一整数     int
	输入对话多选择的类型一文本 string
	流式输出对话完成一文本   string
	流式输出序号数一整数    int

	单选择框选项一文本 string
	多选择框选项一字典 map[string]string
	输入姓名一文本   string
	输入兴趣一文本   string
	推理等级设置一文本 string
	推理等级介绍一文本 string

	安全锁 sync.Mutex
}

var 全局 = &定义一全局变量一一变量结构体{}



type 定义一前端一结构体 struct {
	A0_str string //测试心跳
	L1_str string //只读多行公告提示
	L2_str string //展示选择框或者下拉表单
	T1_str string // 弹出窗口提示
	T2_str string // 显示窗口提示

	M3_str string //2秒显示的置顶提示框

	C2_str string // 隐藏的  遮罩 + 输入网址选择:提取网址内容 选择弹窗
	C3_str string // 隐藏的  遮罩下的 输入网址选择:提取网址内容 选择弹窗
	C5_str string // 隐藏的  输入的+ 弹窗 头
	C6_str string // 隐藏的  输入 的+ 弹窗 正文
	M9_str string //  流式 或者 等待 选择 弹出窗口选项

}

var 返回前端一结构体 = &定义一前端一结构体{}

func main() {

	A0模具一一前期准备()

	go A1模具一一开启网页服务器()

	select {}

}

func A0模具一一前期准备() {

	全局.当前目录一路径文本, _ = os.Getwd()

	全局.多选择框选项一字典 = make(map[string]string)

	全局.程序英汉版一文本 = "英语English"
	全局.程序英汉版一文本 = "汉语"

	var 防火墙提示语一文本 string
	if 全局.程序英汉版一文本 == "汉语" {

		全局.推理等级介绍一文本 = "关"
		全局.程序名与版本一文本 = "PC-gui框架演示-V1.1"
		防火墙提示语一文本 = "防火墙放行-"
	} else {

		全局.推理等级介绍一文本 = "off"
		全局.程序名与版本一文本 = "PC-gui framework demonstration -V1.1"
		防火墙提示语一文本 = "Allow through firewall-"
	}

	全局.程序界面跳转网址一文件路径 = fmt.Sprintf("%s/%s%s.html", 全局.当前目录一路径文本, 防火墙提示语一文本, 全局.程序名与版本一文本)

	return

}

func A2模具一一传递前端信息各种判断(传入传递的前端信息一文本 string) {

	传入传递的前端信息一文本 = strings.TrimSpace(传入传递的前端信息一文本)

	var 返回当前英简体双语一字典 = A2模具一一各种判断信息并返回一简体与英文民字典(传入传递的前端信息一文本)

	var 提示窗口一文本 string

	defer func() {

	}()

	if strings.HasPrefix(传入传递的前端信息一文本, "初始:") {
		全局.双向通信计数器一整数++

		返回前端一结构体.M3_str = A2模具一一返回一般提示一纪录栏源码(返回当前英简体双语一字典["初始欢迎"])
		返回前端一结构体.L2_str = ""
		全局.双向再次通信序号一文本 = "前端"
	} else if strings.HasPrefix(传入传递的前端信息一文本, "输入内容:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")

		var 时间一文本 = "<br />[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "]"

		提示窗口一文本 = 返回当前英简体双语一字典["输入内容提示"] + 传入传递的前端信息一文本 + 时间一文本

		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, strings.Replace(提示窗口一文本, "<br />", "\n", -1))

	} else if strings.HasPrefix(传入传递的前端信息一文本, "流式输出:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")

		全局.流式输出内容一列表 = 全局.流式输出内容一列表[:0]
		全局.流式输入次数一整数 = 10
		全局.流式输出对话完成一文本 = ""
		全局.等待多选择完成一文本 = ""

		全局.安全锁.Lock()
		全局.流式终断一文本 = "123"
		全局.流式输出序号数一整数++

		全局.安全锁.Unlock()

		go func() {

			A2模具一一流式输出(全局.流式输入次数一整数, 全局.流式输出序号数一整数)

		}()

		返回前端一结构体.M9_str = "流式:"

	} else if strings.HasPrefix(传入传递的前端信息一文本, "停止流式的输出:") {

		全局.安全锁.Lock()
		全局.流式终断一文本 = "流式结束"
		全局.安全锁.Unlock()

	} else if strings.HasPrefix(传入传递的前端信息一文本, "对话前的多次弹窗配置:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.流式输出内容一列表 = 全局.流式输出内容一列表[:0]
		全局.流式输入次数一整数 = 0
		全局.流式输出对话完成一文本 = ""
		全局.等待多选择完成一文本 = ""

		全局.安全锁.Lock()
		全局.流式终断一文本 = "123"
		全局.流式输出序号数一整数++
		全局.安全锁.Unlock()

		go func() {

			A2模具一一弹出窗口一等待多选择完成()

			A2模具一一流式输出(全局.流式输入次数一整数, 全局.流式输出序号数一整数)

		}()

		返回前端一结构体.M9_str = "多次等待弹窗选择判断:"
		全局.输入对话多选择的类型一文本 = "多次等待弹窗选择判断:"
	} else if strings.HasPrefix(传入传递的前端信息一文本, "流式输出一多项选择:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		传入传递的前端信息一文本 = strings.TrimSpace(传入传递的前端信息一文本)

		if 传入传递的前端信息一文本 == "对话配置完成" {
			全局.流式输出对话完成一文本 = 传入传递的前端信息一文本
		} else if strings.Contains(传入传递的前端信息一文本, "次") {
			传入传递的前端信息一文本 = strings.Replace(传入传递的前端信息一文本, "次", "", -1)

			var 原文本值一整数, 错误一空值 = strconv.Atoi(strings.TrimSpace(传入传递的前端信息一文本))
			if 错误一空值 != nil {
				全局.流式输入次数一整数 = 5
			} else {
				全局.流式输入次数一整数 = 原文本值一整数
			}

		} else {

		}

		全局.输入对话多选择的类型一文本 = "流式输出一多项选择:" + 传入传递的前端信息一文本

	} else if strings.HasPrefix(传入传递的前端信息一文本, "多次等待弹窗选择判断:") {

		返回前端一结构体.M9_str = ""
		var 当前流式输出序号数一整数 = 全局.流式输出序号数一整数
		for {
			if 全局.流式终断一文本 == "流式结束" {

				全局.等待多选择完成一文本 = "流式终断"
				返回前端一结构体.M9_str = "流式:"
				return
			} else if 当前流式输出序号数一整数 != 全局.流式输出序号数一整数 {
				return
			}

			if len(全局.输入对话多选择的类型一文本) > 2 {

				全局.输入对话多选择的类型一文本 = ""
				break
			}
			time.Sleep(300 * time.Millisecond)

		}

		if 全局.流式输入次数一整数 == 0 {

			返回前端一结构体.C3_str = A2模具一一返回遮罩整个网页界面一询问流式输出次数一弹窗源码()
			返回前端一结构体.M9_str = "多次等待弹窗选择判断:"
			return
		}

		if 全局.流式输出对话完成一文本 == "" {
			返回前端一结构体.C3_str = A2模具一一返回遮罩整个网页界面一询问流式输出完成对话一弹窗源码()
			返回前端一结构体.M9_str = "多次等待弹窗选择判断:"
			return
		}

		全局.等待多选择完成一文本 = "等待多选择完成"
		返回前端一结构体.M9_str = "流式:"

	} else if strings.HasPrefix(传入传递的前端信息一文本, "流式:") {
		返回前端一结构体.M9_str = ""
		var 当前流式输出序号数一整数 = 全局.流式输出序号数一整数
		var 流式结束说明一文本 = 返回当前英简体双语一字典["流式结束"]
		for {
			if 当前流式输出序号数一整数 != 全局.流式输出序号数一整数 {
				return

			} else if 全局.流式终断一文本 == "流式结束" {

				返回前端一结构体.L1_str = 流式结束说明一文本
				return
			}

			当前流式列表数一整数 := len(全局.流式输出内容一列表)

			if 当前流式列表数一整数 == 0 || 当前流式列表数一整数 == 全局.上次流式列表数一整数 {
				continue
			}
			全局.上次流式列表数一整数 = 当前流式列表数一整数
			break

			time.Sleep(300 * time.Millisecond)

		}
		返回前端一结构体.M9_str = "流式:"
		返回前端一结构体.L1_str = strings.Join(全局.流式输出内容一列表, "\n")

	} else if strings.HasPrefix(传入传递的前端信息一文本, "单选选择框示例:") {

		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.单选择框选项一文本 = ""

		返回前端一结构体.L2_str = A2模具一一返回单选选择框示例源码()
		提示窗口一文本 = 返回当前英简体双语一字典["展示选择框示例"]
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "复选选择框示例:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.多选择框选项一字典 = make(map[string]string)

		返回前端一结构体.L2_str = A2模具一一返回复选选择框示例源码()
		提示窗口一文本 = 返回当前英简体双语一字典["展示选择框示例"]
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "单选值:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.单选择框选项一文本 = strings.TrimSpace(传入传递的前端信息一文本)

		返回前端一结构体.L2_str = A2模具一一返回单选选择框示例源码()
		提示窗口一文本 = 返回当前英简体双语一字典["单选框的选值"] + 传入传递的前端信息一文本
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)

	} else if strings.HasPrefix(传入传递的前端信息一文本, "复选值:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		var 复选一对否, 复选值一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")

		if 复选一对否 == "对" {
			提示窗口一文本 = 返回当前英简体双语一字典["复选框的选值"] + 复选值一文本
			全局.安全锁.Lock()
			全局.多选择框选项一字典[复选值一文本] = "checked"
			全局.安全锁.Unlock()
		} else {
			提示窗口一文本 = 返回当前英简体双语一字典["复选框的取消值"] + 复选值一文本
			全局.安全锁.Lock()
			全局.多选择框选项一字典[复选值一文本] = ""
			全局.安全锁.Unlock()
		}
		返回前端一结构体.L2_str = A2模具一一返回复选选择框示例源码()

		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)

	} else if strings.HasPrefix(传入传递的前端信息一文本, "下拉列表示例:") {

		返回前端一结构体.L2_str = A2模具一一返回选择下拉表单源码()
		提示窗口一文本 = 返回当前英简体双语一字典["下拉表单示例"]
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "推理等级:") {

		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.推理等级设置一文本, 全局.推理等级介绍一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")

		返回前端一结构体.L2_str = A2模具一一返回选择下拉表单源码()
		提示窗口一文本 = 返回当前英简体双语一字典["选择推理等级"] + 全局.推理等级介绍一文本
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "多项输入表单示例:") {
		全局.输入姓名一文本, 全局.输入兴趣一文本 = "", ""

		返回前端一结构体.L2_str = A2模具一一返回多项输入表单示例源码()
		提示窗口一文本 = 返回当前英简体双语一字典["多项输入表单示例"]
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "多项输入表单提交:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		全局.输入姓名一文本, 全局.输入兴趣一文本, _ = strings.Cut(传入传递的前端信息一文本, "||")
		提示窗口一文本 = fmt.Sprintf(返回当前英简体双语一字典["多项输入表单的值"], 全局.输入姓名一文本, 全局.输入兴趣一文本)

		返回前端一结构体.L2_str = A2模具一一返回多项输入表单示例源码()

		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, strings.Replace(提示窗口一文本, "<br />", "\n", -1))

	} else if strings.HasPrefix(传入传递的前端信息一文本, "停止流式的输出:") {
		全局.安全锁.Lock()
		全局.流式终断一文本 = "流式结束"
		全局.安全锁.Unlock()

	} else if strings.HasPrefix(传入传递的前端信息一文本, "打开目录:") {

		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")

		var 文件目录一路径文本 = strings.TrimSpace(传入传递的前端信息一文本)

		提示窗口一文本 = 返回当前英简体双语一字典["打开目录"] + 文件目录一路径文本
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 提示窗口一文本)

		文件目录一路径文本 = strings.Replace(文件目录一路径文本, "/", `\`, -1)

		A0模具一一文件管理器打开文件或目录(文件目录一路径文本)
	} else if strings.HasPrefix(传入传递的前端信息一文本, "PC-Gui框架说明:") {
		_, 传入传递的前端信息一文本, _ = strings.Cut(传入传递的前端信息一文本, ":")
		返回前端一结构体.L2_str = A2模具一一返回Gui框架说明源码()

	} else if strings.HasPrefix(传入传递的前端信息一文本, "赞赏:") {

		var 赞赏说明与收款图片一源码文本 string

		if 全局.程序英汉版一文本 == "汉语" {
			赞赏说明与收款图片一源码文本 = `<br /><h2 class="tishi-model-item">&nbsp;&nbsp;赞赏&nbsp;&nbsp;<i class="fas fa-mug-hot"></i></h2><br />%s`
	
			赞赏说明与收款图片一源码文本 = fmt.Sprintf(赞赏说明与收款图片一源码文本, A2模具一一返回一般提示一纪录栏源码(`<img src="https://github.com/jiqi136/Ai-Assistant/blob/main/app/shoukuan.jpg?raw=true" width="230" height="300">`))

		} else {

			var 打赏链接一源码文本 = `<a href="https://ko-fi.com/ft890838"  target="_blank" style="display: flex; align-items: center; justify-content: space-between; background: #3b82f6; color: white; text-decoration: none; padding: 14px 20px; border-radius: 10px; font-weight: 500; transition: all 0.2s; box-shadow: 0 4px 6px -1px rgba(59, 130, 246, 0.2);">
			<span>Reward  &nbsp;&nbsp;(ko-fi) https://ko-fi.com/ft890838</span><span style="font-size: 12px; opacity: 0.9;">ko-fi →</span> </a>`

	
			赞赏说明与收款图片一源码文本 = fmt.Sprintf(`<br /><h2 class="tishi-model-item">&nbsp;&nbsp;Reward&nbsp;&nbsp;<i class="fas fa-mug-hot"></i></h2><br />%s`, 打赏链接一源码文本)
			赞赏说明与收款图片一源码文本 = A2模具一一返回一般提示一纪录栏源码(赞赏说明与收款图片一源码文本)
		}

		返回前端一结构体.L2_str = 赞赏说明与收款图片一源码文本

	} else if strings.HasPrefix(传入传递的前端信息一文本, "退出:") {

		提示窗口一文本 = A2模具一一返回一般提示一纪录栏源码(返回当前英简体双语一字典["程序已经退出提示"])
		A2模具一一返回全部的纪录栏目与提示网页("成功", 提示窗口一文本, 返回当前英简体双语一字典["程序已经退出提示"])

		go func() {
			time.Sleep(1 * time.Second)
			log.Fatal("退出程序.")
			os.Exit(0)

		}()

	} else {

	}

	return

}

func A0模具一一保存跳转网址到本地文件(传入网址一文本 string) {

	var 微博跳转网址网页源码 = `<html> <head><meta http-equiv="Content-Type" content="text/html; charset=utf8" /><title></title><meta http-equiv="refresh" content="0;URL=替换网址"></head><body></body></html>`

	微博跳转网址网页源码 = strings.Replace(微博跳转网址网页源码, "替换网址", 传入网址一文本, -1)
	A0模具一一文件保存或追加(全局.程序界面跳转网址一文件路径, 微博跳转网址网页源码, "保存")

	return

}

func A0模具一一简捷在浏览器打开网址(传入网址一文本 string) {
	time.Sleep(300 * time.Millisecond)

	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", 传入网址一文本).Run()
	if err == nil {
		return
	}

	cmd := exec.Command("cmd", "/c", "start", 传入网址一文本)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	var 错误一空值 = cmd.Start()
	if 错误一空值 != nil {

		_ = 错误一空值
	}
	return
}

func A0模具一一文件管理器打开文件或目录(传入打开的文件一文本 string) {
	传入打开的文件一文本 = strings.Replace(传入打开的文件一文本, `/`, `\`, -1)
	传入打开的文件一文本 = strings.Replace(传入打开的文件一文本, `\\`, ``, -1)

	传入打开的文件一文本 = strings.TrimSpace(传入打开的文件一文本)
	传入打开的文件一文本 = strings.Trim(传入打开的文件一文本, `\`)

	cmd := exec.Command("cmd", "/c", "start", "", 传入打开的文件一文本)

	err := cmd.Run()

	if err == nil {

		return
	}

	err = exec.Command(`cmd`, `/c`, `explorer`, 传入打开的文件一文本).Start()

	if err != nil {

		_ = err
	}

	return

}

func A0模具一一关闭本地程序并回收内存() {
	var 内存占用数 runtime.MemStats
	runtime.ReadMemStats(&内存占用数)

	runtime.GC()
	runtime.ReadMemStats(&内存占用数)

}

func A0模具一一读取本地主页网页文件并返回源码(传入本地路径一文本 string) string {
	data, err := ioutil.ReadFile(传入本地路径一文本)
	if err != nil {

		log.Fatal("读取本地主页网页文件并返回源码 ：错误,Read the local homepage webpage file and return the source code: error.", err)
		return ""
	}

	return string(data)
}

func A0模具一一文件保存或追加(传入文件路径一文本, 传入写入内容一文本, 传入保存类型一文本 string) {
	var 目录路径一文本 = filepath.Dir(传入文件路径一文本)
	var _, err = os.Stat(目录路径一文本)
	if err != nil {

		_ = os.MkdirAll(目录路径一文本, os.ModePerm)

	}

	var 文件内容一数据流 *os.File
	if 传入保存类型一文本 == "追加" {
		文件内容一数据流, err = os.OpenFile(传入文件路径一文本, os.O_APPEND|os.O_CREATE, 0666)
	} else {
		文件内容一数据流, err = os.OpenFile(传入文件路径一文本, os.O_TRUNC|os.O_CREATE, 0666)

	}

	if err != nil {

		return
	}
	defer 文件内容一数据流.Close()
	var 文件写入 = bufio.NewWriter(文件内容一数据流)
	文件写入.WriteString(传入写入内容一文本)
	文件写入.Flush()

}

func A0模具一一无效占位log库包() {
	fmt.Println("")
	time.Sleep(100 * time.Millisecond)
	log.Fatal("")
}

//===================================================================================
//===================================================================================


type RequestData struct {
	Message string `json:"message"`
}

func A1模具一一开启网页服务器() {

	var 随机端口一整数 int
	rand.Seed(time.Now().UnixNano())
	随机端口一整数 = 35658 + rand.Intn(2300)

	for {
		if A1模具一一检查端口是否可用返回对否(随机端口一整数) {

			break
		} else {

			随机端口一整数++
		}

	}

	var 端口一文本 = fmt.Sprintf(":%d", 随机端口一整数)

	http.HandleFunc("/", A1模具一设置主页路由)

	http.HandleFunc("/chat", A1模具一一接受网页前端各信息)

	http.HandleFunc("/bc", A1模具一主动向前端发收消息)

	全局.本地网页网址一文本 = fmt.Sprintf("http://localhost%s", 端口一文本)

	go func() {

		A0模具一一保存跳转网址到本地文件(全局.本地网页网址一文本)

		A0模具一一简捷在浏览器打开网址(全局.本地网页网址一文本)

	}()

	if err := http.ListenAndServe(端口一文本, nil); err != nil {

		log.Fatal(err)
	}

}

func A1模具一设置主页路由(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var 网页源码一文本 string

	网页源码一文本 = A1模具一一返回程序界面网页源码()
	//网页源码一文本 = A0模具一一读取本地主页网页文件并返回源码("./9-网页界面.html")

	网页源码一文本 = strings.Replace(网页源码一文本, "|替换换行|", "\n", -1)

	网页源码一文本 = A1模具一一返回主页的提示说明的源码(网页源码一文本)
	网页源码一文本 = A1模具一一返回主页的英汉互换后的源码(网页源码一文本)

	网页源码一文本 = strings.Replace(网页源码一文本, "|替换程序名|", 全局.程序名与版本一文本, -1)
	网页源码一文本 = strings.Replace(网页源码一文本, "|替换程序目录|", 全局.当前目录一路径文本, -1)

	网页源码一文本 = strings.Replace(网页源码一文本, "|替换CSS样式|", A1模具一一返回网页源码公用样式(), -1)

	var 当前英简体双语一字典 = make(map[string]string)

	if 全局.程序英汉版一文本 == "汉语" {

		当前英简体双语一字典["IE浏览器过时"] = "您的IE浏览器过时，推荐使用 Chrome/Firefox/Edge 等现代浏览器."
	} else {

		当前英简体双语一字典["IE浏览器过时"] = "Your IE browser is outdated. We recommend using modern browsers like Chrome, Firefox, or Edge."

	}

	网页源码一文本 = strings.Replace(网页源码一文本, "|替换IE浏览器过时|", 当前英简体双语一字典["IE浏览器过时"], -1)
	网页源码一文本 = strings.Replace(网页源码一文本, "<!--替换各国语言-->", A1模具一一返回各国语言下拉表单源码(), -1)

	网页源码一文本 = strings.Replace(网页源码一文本, "|替换程序名|", 全局.程序名与版本一文本, -1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.Write([]byte(网页源码一文本))

	return
}

func A1模具一一检查端口是否可用返回对否(传入随机端口一整数 int) bool {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 传入随机端口一整数))
	if err != nil {

		return false
	}

	_ = ln.Close()
	return true
}

func A1模具一一接受网页前端各信息(w http.ResponseWriter, r *http.Request) {
	defer func() {

	}()

	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "解析 JSON 数据失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	全局.安全锁.Lock()
	*返回前端一结构体 = 定义一前端一结构体{}
	全局.安全锁.Unlock()

	A2模具一一传递前端信息各种判断(data.Message)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(返回前端一结构体)

	time.Sleep(200 * time.Millisecond)

	r.Body.Close()

	if strings.HasPrefix(data.Message, "多次等待弹窗选择判断:") || strings.HasPrefix(data.Message, "流式:") {

	} else {
		A0模具一一关闭本地程序并回收内存()
	}

}

func A1模具一主动向前端发收消息(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	var 双向计数器一整数 = 全局.双向通信计数器一整数
	for {
		for {
			if 全局.双向通信计数器一整数 != 双向计数器一整数 {

				return
			}

			if 全局.双向再次通信序号一文本 != "" {

				全局.双向再次通信序号一文本 = ""
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		data, _ := json.Marshal(返回前端一结构体)

		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()

	}

}

func A1模具一一返回主页的英汉互换后的源码(传入主页一源码文本 string) string {
	var 当前英简体双语一字典 = make(map[string]string)
	if 全局.程序英汉版一文本 == "汉语" {

		当前英简体双语一字典["网页语言"] = "zh-CN"
		当前英简体双语一字典["请输入内容"] = "请输入内容..."
		当前英简体双语一字典["确认"] = "确认"
		当前英简体双语一字典["Ai流式输出变更按钮"] = "Ai流式输出-变更按钮"
		当前英简体双语一字典["停止流式输出"] = "停止流式输出"
		当前英简体双语一字典["对话前的多次弹窗配置"] = "对话前的多次弹窗配置"
		当前英简体双语一字典["单选选择框示例"] = "单选选择框示例"
		当前英简体双语一字典["复选选择框示例流式"] = "复选选择框示例"
		当前英简体双语一字典["下拉列表示例"] = "下拉列表示例"
		当前英简体双语一字典["多项输入表单示例"] = "多项输入表单示例"
		当前英简体双语一字典["打开程序目录"] = "打开程序目录"
		当前英简体双语一字典["退出程序"] = "退出程序"

		当前英简体双语一字典["只读多行输入框"] = "只读多行输入框. 鼠标悬停按钮时，自动显示提示信息。"

		当前英简体双语一字典["赞赏"] = "赞赏"
		当前英简体双语一字典["PC-Gui框架说明"] = "PC-Gui框架说明"
	} else {

		当前英简体双语一字典["网页语言"] = "en-US"
		当前英简体双语一字典["请输入内容"] = "Please enter content..."
		当前英简体双语一字典["确认"] = "Confirm"
		当前英简体双语一字典["只读多行输入框"] = "Read-only multiline input box.A tooltip appears when you hover over the button."

		当前英简体双语一字典["Ai流式输出变更按钮"] = "AI Streaming Output - Change Button"
		当前英简体双语一字典["停止流式输出"] = "Stop Streaming Output"
		当前英简体双语一字典["对话前的多次弹窗配置"] = "Multiple Popup Configuration Before Conversation"
		当前英简体双语一字典["单选选择框示例"] = "Example of radio selection box"
		当前英简体双语一字典["复选选择框示例流式"] = "Checkbox Example"
		当前英简体双语一字典["下拉列表示例"] = "Dropdown List Example"
		当前英简体双语一字典["多项输入表单示例"] = "Multi-input Form Example"
		当前英简体双语一字典["打开程序目录"] = "Open Program Directory"
		当前英简体双语一字典["退出程序"] = "Exit Program"
		当前英简体双语一字典["赞赏"] = "admire"
		当前英简体双语一字典["PC-Gui框架说明"] = "PC-Gui explain"

	}

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换程序名|", 全局.程序名与版本一文本, -1)

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换语言|", 当前英简体双语一字典["网页语言"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换输入框说明|", 当前英简体双语一字典["请输入内容"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换只读多行输入框|", 当前英简体双语一字典["只读多行输入框"], -1)

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换确认|", 当前英简体双语一字典["确认"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换Ai流式输出按钮|", 当前英简体双语一字典["Ai流式输出变更按钮"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换停止流式输出|", 当前英简体双语一字典["停止流式输出"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换多次弹窗配置|", 当前英简体双语一字典["对话前的多次弹窗配置"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换单选选择框示例|", 当前英简体双语一字典["单选选择框示例"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换复选选择框示例|", 当前英简体双语一字典["复选选择框示例流式"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换下拉列表示例|", 当前英简体双语一字典["下拉列表示例"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换多项输入表单示例|", 当前英简体双语一字典["多项输入表单示例"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换打开程序目录|", 当前英简体双语一字典["打开程序目录"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换退出程序|", 当前英简体双语一字典["退出程序"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换PC-Gui框架说明|", 当前英简体双语一字典["PC-Gui框架说明"], -1)

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|赞赏|", 当前英简体双语一字典["赞赏"], -1)

	return 传入主页一源码文本
}

func A1模具一一返回主页的提示说明的源码(传入主页一源码文本 string) string {
	var 当前英简体双语一字典 = make(map[string]string)
	if 全局.程序英汉版一文本 == "汉语" {

		当前英简体双语一字典["输入按钮"] = `前端html: class='Dataset-Text' data-text='输入内容:'&#10;JS:统一响应点击click元素'.Dataset-Text',同时检查 INPUT 和 TEXTAREA 元素 &#10;后端接收并响应:strings.HasPrefix('', '输入内容:')`
		当前英简体双语一字典["流式输出按钮"] = `前端html: 异步不断更新前端网页内容,减少AI返回内容等待时间,增加动态感.&#10;<textarea id='L1_str'> </textarea>&#10;JS:接收后端数据, if (data.L1_str !== '') {&#10;L1_str.innerHTML = data.L1_str;&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;后端接收并响应: (struct).L1_str='AI返回内容' &#10; (struct).M9_str='流式:'&#10; strings.HasPrefix('', '流式:')`
		当前英简体双语一字典["停止输出"] = `前端html: class='Dataset-Text' data-text='停止流式的输出:'&#10;JS:统一响应点击click元素'.Dataset-Text' &#10;后端接收并响应:strings.HasPrefix('', '停止流式的输出:')`

		当前英简体双语一字典["多次弹窗"] = `本套框架最繁杂部分,为此放弃不稳定的'SSE' 和 'WebSocket'.&#10;前端html: 多次更新弹出弹窗网页内容,等待选择.&#10;<div id='C3_str'> </div>&#10;JS:接收后端数据, if (data.C3_str !== '') {&#10;C3_str.innerHTML = data.C3_str;&#10;C2_str.style.display = 'flex';//显示&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;后端接收并响应:(struct).C3_str='等待选择的弹窗源码' &#10; (struct).M9_str='多次等待弹窗选择判断:'&#10; strings.HasPrefix('', '多次等待弹窗选择判断:')`
		当前英简体双语一字典["单选选择框"] = `前端html: <input type='radio' name='choice'  value='A'>&#10;JS:统一响应【Change事件】统一处理 单选与复选元素'change',&#10;处理单选框 else if (target.type === 'radio' && target.name === 'choice')&#10;后端接收并响应:strings.HasPrefix('', '单选值:')`
		当前英简体双语一字典["复选选择框"] = `前端html: <input type='checkbox' class='toggle-item'  value='A' >&#10;JS:统一响应【Change事件】统一处理 单选与复选元素'change',&#10;处理复选框  if (target.type === 'checkbox' && target.classList.contains('toggle-item'))&#10;后端接收并响应:strings.HasPrefix('', '复选值:')`
		当前英简体双语一字典["下拉列表"] = `前端html: <select name='format' class='form-select  change-value' >&#10;<option value='推理等级:关'></option>&#10;JS:统一响应下拉表元素 change 元素'change',检查是否是目标下拉列表&#10; if (event.target.classList.contains('change-value') && event.target.tagName === 'SELECT')&#10;后端接收并响应:strings.HasPrefix('', '推理等级:')`
		当前英简体双语一字典["输入表单"] = `前端html: <form class='ajax-form' method='POST' action='#'>&#10;<input name='fullName'>&#10;JS:统一响应【Submit事件】表单提交 元素'submit',使用 .value 获取各个字段的值&#10; const name = form.querySelector('[name='fullName']').value; &#10;后端接收并响应:strings.HasPrefix('', '多项输入表单提交:')`
		当前英简体双语一字典["打开目录提示"] = `前端html: class='Dataset-Text' data-text='打开目录:'&#10;后端接收并响应:strings.HasPrefix('', '打开目录:')&#10;测试安全无误报,打开本地电脑的目录或文件`
		当前英简体双语一字典["赞赏提示"] = "前端html: class='Dataset-Text' data-text='赞赏:'&#10;后端接收并响应:strings.HasPrefix('', '赞赏:')&#10;测试展示,打包文件的图片或网络图片"
		当前英简体双语一字典["退出程序"] = "前端html: class='Dataset-Text' data-text='退出:'&#10;后端接收并响应:strings.HasPrefix('', '退出:')&#10;先展示'退出程序'提示,程序最后才关闭退出."

	} else {

		当前英简体双语一字典["输入按钮"] = `Frontend HTML: class='Dataset-Text' data-text='输入内容:'&#10;JS: Unified response to click on element '.Dataset-Text', while checking INPUT and TEXTAREA elements &#10;Backend receives and responds: strings.HasPrefix('', '输入内容:')`
		当前英简体双语一字典["流式输出按钮"] = `Frontend HTML: Asynchronously updates frontend web content continuously, reducing waiting time for AI responses, adding dynamic feel.&#10;<textarea id='L1_str'> </textarea>&#10;JS: Receives backend data, if (data.L1_str !== '') {&#10;L1_str.innerHTML = data.L1_str;&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;Backend receives and responds: (struct).L1_str='AI response content' &#10; (struct).M9_str='流式:'&#10; strings.HasPrefix('', '流式:')`
		当前英简体双语一字典["停止输出"] = `Frontend HTML: class='Dataset-Text' data-text='停止流式的输出:'&#10;JS: Unified response to click on element '.Dataset-Text' &#10;Backend receives and responds: strings.HasPrefix('', '停止流式的输出:')`
		当前英简体双语一字典["多次弹窗"] = `The most complex part of this framework, therefore abandoning unstable 'SSE' and 'WebSocket'.&#10;Frontend HTML: Multiple updates to popup modal content, waiting for selection.&#10;<div id='C3_str'> </div>&#10;JS: Receives backend data, if (data.C3_str !== '') {&#10;C3_str.innerHTML = data.C3_str;&#10;C2_str.style.display = 'flex';//display&#10;if (data.M9_str !== '') {&#10;PostB(data.M9_str).then((data) => {&#10;Backend receives and responds: (struct).C3_str='popup source code waiting for selection' &#10; (struct).M9_str='多次等待弹窗选择判断:'&#10; strings.HasPrefix('', '多次等待弹窗选择判断:')`
		当前英简体双语一字典["单选选择框"] = `Frontend HTML: <input type='radio' name='choice'  value='A'>&#10;JS: Unified response to [Change event] uniformly handling radio and checkbox elements 'change',&#10;Handle radio button else if (target.type === 'radio' && target.name === 'choice')&#10;Backend receives and responds: strings.HasPrefix('', '单选值:')`
		当前英简体双语一字典["复选选择框"] = `Frontend HTML: <input type='checkbox' class='toggle-item'  value='A' >&#10;JS: Unified response to [Change event] uniformly handling radio and checkbox elements 'change',&#10;Handle checkbox if (target.type === 'checkbox' && target.classList.contains('toggle-item'))&#10;Backend receives and responds: strings.HasPrefix('', '复选值:')`
		当前英简体双语一字典["下拉列表"] = `Frontend HTML: <select name='format' class='form-select  change-value' >&#10;<option value='推理等级:关'></option>&#10;JS: Unified response to select element change event 'change', check if it is the target dropdown&#10; if (event.target.classList.contains('change-value') && event.target.tagName === 'SELECT')&#10;Backend receives and responds: strings.HasPrefix('', '推理等级:')`
		当前英简体双语一字典["输入表单"] = `Frontend HTML: <form class='ajax-form' method='POST' action='#'>&#10;<input name='fullName'>&#10;JS: Unified response to [Submit event] form submission element 'submit', use .value to get the values of each field&#10; const name = form.querySelector('[name='fullName']').value; &#10;Backend receives and responds: strings.HasPrefix('', '多项输入表单提交:')`
		当前英简体双语一字典["打开目录提示"] = `Frontend HTML: class='Dataset-Text' data-text='打开目录:'&#10;Backend receives and responds: strings.HasPrefix('', '打开目录:')&#10;Tested safe without false positives, open a directory or file on the local computer`
		当前英简体双语一字典["赞赏提示"] = `Frontend HTML: class='Dataset-Text' data-text='赞赏:'&#10;Backend receives and responds: strings.HasPrefix('', '赞赏:')&#10;Test display, images from packaged files or network images`
		当前英简体双语一字典["退出程序"] = `Frontend HTML: class='Dataset-Text' data-text='退出:'&#10;Backend receives and responds: strings.HasPrefix('', '退出:')&#10;First display 'Exit program' prompt, then the program finally closes and exits.`

	}

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换输入按钮提示|", 当前英简体双语一字典["输入按钮"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换Ai流式输出按钮提示|", 当前英简体双语一字典["流式输出按钮"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换停止输出按钮提示|", 当前英简体双语一字典["停止输出"], -1)

	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换多次弹窗提示|", 当前英简体双语一字典["多次弹窗"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换单选选择框提示|", 当前英简体双语一字典["单选选择框"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换复选选择框提示|", 当前英简体双语一字典["复选选择框"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换下拉列表提示|", 当前英简体双语一字典["下拉列表"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换输入表单提示|", 当前英简体双语一字典["输入表单"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换打开目录提示|", 当前英简体双语一字典["打开目录提示"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换赞赏提示|", 当前英简体双语一字典["赞赏提示"], -1)
	传入主页一源码文本 = strings.Replace(传入主页一源码文本, "|替换退出程序提示|", 当前英简体双语一字典["退出程序"], -1)

	return 传入主页一源码文本
}

//===================================================================================
//===================================================================================


func A2模具一一返回全部的纪录栏目与提示网页(传入提示类型一文本, 传入提示窗口一文本, 传入网页源码一文本 string) (返回网页源码一文本 string) {

	if strings.Contains(传入提示窗口一文本, "class=") {

		返回前端一结构体.M3_str = 传入提示窗口一文本
	} else if 传入提示类型一文本 == "成功" {
		返回前端一结构体.M3_str = A2模具一一返回一般提示一纪录栏源码(传入提示窗口一文本)
	} else if 传入提示类型一文本 == "错误" {

		返回前端一结构体.M3_str = A2模具一一返回错误提示一纪录栏源码(传入提示窗口一文本)
	} else {
		返回前端一结构体.M3_str = A2模具一一返回一般提示一纪录栏源码(传入提示窗口一文本)
	}

	返回前端一结构体.L1_str = strings.Replace(传入网页源码一文本, "|替换提示窗口|", "", -1)

	return

}

func A2模具一一返回错误提示一纪录栏源码(传入错误提示一文本 string) (返回错误提示一源码文本 string) {
	返回错误提示一源码文本 = `<div >
	<div class="tishi-note-box">
	<h3 class="tishi-note-text2">
	<h3 class="tishi-icon">✗ </h3> 
            <p  class="prompt-content">|替换错误提示|</p>
		</h3>
	</div>	
	</div>	
	`
	返回错误提示一源码文本 = strings.Replace(返回错误提示一源码文本, "|替换错误提示|", 传入错误提示一文本, -1)

	return

}

func A2模具一一返回一般提示一纪录栏源码(传入提示文本一文本 string) (返回一般提示纪录栏一源码文本 string) {
	返回一般提示纪录栏一源码文本 = `<div>
	<div class="tishi-note-box">
	<h3 class="tishi-note-text2">
	<h3 class="tishi-success-icon">✔</h3> 
            <p class="prompt-content">|替换提示|</p>
		</h3>
	</div>
	</div>
	`

	返回一般提示纪录栏一源码文本 = strings.Replace(返回一般提示纪录栏一源码文本, "|替换提示|", 传入提示文本一文本, -1)

	return

}

func A2模具一一各种判断信息并返回一简体与英文民字典(传入传递的前端信息一文本 string) (返回当前英简体双语一字典 map[string]string) {
	返回当前英简体双语一字典 = make(map[string]string)
	if 全局.程序英汉版一文本 == "汉语" {

		返回当前英简体双语一字典["打开目录"] = "正在打开目录:"
		返回当前英简体双语一字典["初始欢迎"] = "程序初始,欢迎测试演示。"

		返回当前英简体双语一字典["程序已经退出提示"] = "程序已经退出."
		返回当前英简体双语一字典["展示选择框示例"] = "展示选择框示例."
		返回当前英简体双语一字典["输入内容提示"] = "输入内容提示:<br />"

		返回当前英简体双语一字典["下拉表单示例"] = "下拉表单示例."
		返回当前英简体双语一字典["多项输入表单示例"] = "多项输入表单示例."

		返回当前英简体双语一字典["选择推理等级"] = "选择推理等级:"
		返回当前英简体双语一字典["单选框的选值"] = "单选框的选值:"
		返回当前英简体双语一字典["复选框的选值"] = "复选框的选值:"
		返回当前英简体双语一字典["复选框的取消值"] = "复选框的取消值:"
		返回当前英简体双语一字典["多项输入表单的值"] = "多项输入表单:<br />姓名:%s<br />兴趣:%s"
		返回当前英简体双语一字典["流式结束"] = "流式结束."

	} else {
		返回当前英简体双语一字典["打开目录"] = "Opening directory:"
		返回当前英简体双语一字典["初始欢迎"] = "Program initialized, welcome to the test demo."
		返回当前英简体双语一字典["程序已经退出提示"] = "Program has exited."
		返回当前英简体双语一字典["展示选择框示例"] = "Displaying selection box example."
		返回当前英简体双语一字典["输入内容提示"] = "Input prompt:<br />"
		返回当前英简体双语一字典["下拉表单示例"] = "Dropdown form example."
		返回当前英简体双语一字典["多项输入表单示例"] = "Multi-input form example."
		返回当前英简体双语一字典["选择推理等级"] = "Inference grade:"
		返回当前英简体双语一字典["单选框的选值"] = "Radio button selected value:"
		返回当前英简体双语一字典["复选框的选值"] = "Checkbox selected value:"
		返回当前英简体双语一字典["复选框的取消值"] = "Checkbox deselected value:"
		返回当前英简体双语一字典["多项输入表单的值"] = "Multi-input form:<br />Name:%s<br />Interests:%s"
		返回当前英简体双语一字典["流式结束"] = "Streaming end."

	}

	return

}

func A2模具一一流式输出(传入完成次数一整数, 传入流式输出序号数一整数 int) {
	var 计数器一整数 int
	var 时间一文本, 输出内容一文本 string

	for {
		计数器一整数++
		if 计数器一整数 == 传入完成次数一整数+1 {

			全局.安全锁.Lock()
			全局.流式终断一文本 = "流式结束"
			全局.安全锁.Unlock()

			return
		} else if 传入流式输出序号数一整数 != 全局.流式输出序号数一整数 {
			return
		}

		时间一文本 = "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "]"

		输出内容一文本 = fmt.Sprintf("%d ,%s", 计数器一整数, 时间一文本)

		全局.流式输出内容一列表 = append(全局.流式输出内容一列表, 输出内容一文本)
		time.Sleep(2 * time.Second)
		if 全局.流式终断一文本 == "流式结束" {

			return
		}

	}

	return

}

func A2模具一一弹出窗口一等待多选择完成() {
	全局.等待多选择完成一文本 = ""
	for {

		time.Sleep(300 * time.Millisecond)

		if len(全局.等待多选择完成一文本) > 2 {
			return
		}

	}

	return

}


func A2模具一一返回遮罩整个网页界面一询问流式输出次数一弹窗源码() (返回询问生成图片多选项一源码 string) {

	var 当前英简体双语一字典 = make(map[string]string)
	var 流式输出次数一文本, 流式输出次数项一源码文本 string

	if 全局.程序英汉版一文本 == "汉语" {

		当前英简体双语一字典["纪录栏源码"] =
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
		流式输出次数项一源码文本 = `<span class="chat-item-C Dataset-Text" data-text="流式输出一多项选择:|替换次数|次">|替换次数| 次</span>`

	} else {

		当前英简体双语一字典["纪录栏源码"] = `<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> <span class="tishi-highlight"> Streaming output </span > please select:</h2>
			<hr>
				<div class="dropdown-container">
			<span class="dropdown-L0" style="background: linear-gradient(135deg, #a6bed1 0%, #7d97b4 100%);color: #133025;" >Number of streaming outputs ▼ <i class="fas fa-chevron-down"></i></span>
					<div class="dropdown-menu" style="top:100%;min-height:190px;">
							|替换流式输出次数表项|
						
					</div>
				</div>
		</aside>`
		流式输出次数项一源码文本 = `<span class="chat-item-C Dataset-Text" data-text="流式输出一多项选择:|替换次数|次">|替换次数| number</span>`

	}

	返回询问生成图片多选项一源码 = 当前英简体双语一字典["纪录栏源码"]

	var 其它各类流式输出次数一列表 = []string{"5", "10", "15"}

	var 文件内容一列表 []string

	for _, 流式输出次数一文本 = range 其它各类流式输出次数一列表 {

		文件内容一列表 = append(文件内容一列表, strings.Replace(流式输出次数项一源码文本, "|替换次数|", 流式输出次数一文本, -1))
	}

	返回询问生成图片多选项一源码 = strings.Replace(返回询问生成图片多选项一源码, "|替换流式输出次数表项|", strings.Join(文件内容一列表, "\n"), -1)

	return

}

func A2模具一一返回遮罩整个网页界面一询问流式输出完成对话一弹窗源码() (返回询问生成图片多选项一源码 string) {

	var 当前英简体双语一字典 = make(map[string]string)
	if 全局.程序英汉版一文本 == "汉语" {

		当前英简体双语一字典["纪录栏源码"] =
			` 
		<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> 流式输出<span style="color:#133025;"> 对话配置完成 </span> :</h2>
			<hr>
			
			<button data-text="流式输出一多项选择:对话配置完成"  class="stream-button Dataset-Text s3" style="color: #133025; width: 100%;"><span class="tishi-highlight">对话配置完成</span></button>

			
		</aside>
		`

	} else {

		当前英简体双语一字典["纪录栏源码"] = `		<aside class="tishi-sidebar" style="height:350px; width:350px;">
			<h2> Streaming output <span style="color:#133025;"> Dialog configuration completed </span> :</h2>
			<hr>
			
			<button data-text="流式输出一多项选择:对话配置完成"  class="stream-button Dataset-Text s3" style="color: #133025; width: 100%;"><span class="tishi-highlight">Dialog configuration completed</span></button>

			
			</aside>`

	}

	返回询问生成图片多选项一源码 = 当前英简体双语一字典["纪录栏源码"]

	return

}

func A2模具一一返回复选选择框示例源码() (返回选择框示例源码一文本 string) {
	var 原多选择框选项一文本 string

	if 全局.程序英汉版一文本 == "汉语" {

		返回选择框示例源码一文本 = `<div class="action-section">
			<div class="multi-select-section">
			<div class="checkbox-group">
				<span class="checkbox-group-title">多选（点击后查看控制台输出）</span>
				<div class="checkbox-options">
				|替换选项列表|
					
				</div>
			</div>
		</div>
	 
	 </div>`
		原多选择框选项一文本 = `<div class="checkbox-item"> <input type="checkbox" class="toggle-item"  value="|替换选项|" |替换已勾选|>
	 <label>选项 |替换选项|</label></div>`
	} else {

		返回选择框示例源码一文本 = `<div class="action-section">
			<div class="multi-select-section">
			<div class="checkbox-group">
				<span class="checkbox-group-title">Multi-selection (click to view console output)</span>
				<div class="checkbox-options">
				|替换选项列表|
					
				</div>
			</div>
		</div>
	 
	 </div>`
		原多选择框选项一文本 = `<div class="checkbox-item"> <input type="checkbox" class="toggle-item"  value="|替换选项|" |替换已勾选|>
	 <label>option |替换选项|</label></div>`

	}

	var 文本一列表 = []string{"A", "B", "C"}
	var 选项一文本, 新的多选择框选项一文本 string
	var 文件内容一列表 []string

	if len(全局.多选择框选项一字典) == 0 {
		for _, 选项一文本 = range 文本一列表 {
			全局.安全锁.Lock()
			全局.多选择框选项一字典[选项一文本] = ""
			全局.安全锁.Unlock()
		}
	}

	for _, 选项一文本 = range 文本一列表 {

		新的多选择框选项一文本 = strings.Replace(原多选择框选项一文本, "|替换选项|", 选项一文本, -1)

		新的多选择框选项一文本 = strings.Replace(新的多选择框选项一文本, "|替换已勾选|", 全局.多选择框选项一字典[选项一文本], -1)
		文件内容一列表 = append(文件内容一列表, 新的多选择框选项一文本)

	}

	返回选择框示例源码一文本 = strings.Replace(返回选择框示例源码一文本, "|替换选项列表|", strings.Join(文件内容一列表, "\n"), -1)

	return

}

func A2模具一一返回单选选择框示例源码() (返回选择框示例源码一文本 string) {
	var 原单选择框选项一文本 string

	if 全局.程序英汉版一文本 == "汉语" {

		返回选择框示例源码一文本 = `<!-- 单选项 -->
        <div class="action-section">
            <div class="radio-group">
                <span class="radio-group-title">单选</span>
				<div class="radio-options">
				|替换选项列表|
                    
                </div>
            </div>
		 </div>`
		原单选择框选项一文本 = `<div class="radio-item"><input type="radio" name="choice"  value="|替换选项|" |替换已勾选|>
		 <label>选项|替换选项|</label></div>`

	} else {

		返回选择框示例源码一文本 = `<div class="action-section">
		<div class="radio-group">
			<span class="radio-group-title">Radio</span>
			<div class="radio-options">
			|替换选项列表|
			</div>
		</div>
	 </div>`
		原单选择框选项一文本 = `<div class="radio-item"><input type="radio" name="choice"  value="|替换选项|" |替换已勾选|>
	<label>option |替换选项|</label></div>`

	}

	var 文本一列表 = []string{"A", "B", "C", "D"}
	var 选项一文本, 新的单选择框选项一文本 string
	var 文件内容一列表 []string

	for _, 选项一文本 = range 文本一列表 {

		新的单选择框选项一文本 = strings.Replace(原单选择框选项一文本, "|替换选项|", 选项一文本, -1)

		if 选项一文本 == 全局.单选择框选项一文本 {
			新的单选择框选项一文本 = strings.Replace(新的单选择框选项一文本, "|替换已勾选|", "checked", -1)
		} else {
			新的单选择框选项一文本 = strings.Replace(新的单选择框选项一文本, "|替换已勾选|", "", -1)
		}
		文件内容一列表 = append(文件内容一列表, 新的单选择框选项一文本)

	}

	返回选择框示例源码一文本 = strings.Replace(返回选择框示例源码一文本, "|替换选项列表|", strings.Join(文件内容一列表, "\n"), -1)

	return

}

func A2模具一一返回选择下拉表单源码() (返回返回选择下拉表单源码一文本 string) {

	if 全局.程序英汉版一文本 == "汉语" {

		返回返回选择下拉表单源码一文本 = `<div class="action-section">
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

		返回返回选择下拉表单源码一文本 = `<div class="action-section"> Dropdown example: Select reasoning program. <br /><br /><form class="dynamic-form" action="#"> 
		<div style="display: flex; align-items: center; gap: 5px;"> 
		<select name="format" class="form-select change-value" style="background-color: #f0f7f0; color: #388e3c; flex: 1; padding: 6px; border-radius: 8px; border: 1px solid #c8e6c9;"> <option value="" selected disabled>Adjust reasoning program ▼:</option> 
		<option value="" disabled>Current reasoning program: |替换当前推理等级|</option>
		 <option value="推理等级:关:Off">Reasoning program: Off</option>
		  <option value="推理等级:低:Low">Reasoning program: Low</option>
		   <option value="推理等级:中:Medium">Reasoning program: Medium</option>
			<option value="推理等级:高:High">Reasoning program: High</option>
			 </select> </div> </form> </div>`

	}

	返回返回选择下拉表单源码一文本 = strings.Replace(返回返回选择下拉表单源码一文本, "|替换当前推理等级|", 全局.推理等级介绍一文本, -1)
	return

}

func A2模具一一返回多项输入表单示例源码() (返回多项输入表单示例源码一文本 string) {

	if 全局.程序英汉版一文本 == "汉语" {

		返回多项输入表单示例源码一文本 = ` <div class="action-section">
		<form class="ajax-form" method="POST" action="#">
			<label for="name">姓名</label>
			<input type="text"  class="inputB" name="fullName" placeholder="例如：李小萌" |替换姓名值|>
	
			<label for="interest">感兴趣的领域</label>
			<input type="text" class="inputB" name="interest" placeholder="前端 / 交互 / 设计"  |替换兴趣值|>
			<br />
			<button  class="btn btn-primary"  type="submit">📨 提交 (ajax-form) </button>
		</form></div> `
	} else {

		返回多项输入表单示例源码一文本 = `<div class="action-section">
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

	if 全局.输入姓名一文本 == "" {
		返回多项输入表单示例源码一文本 = strings.Replace(返回多项输入表单示例源码一文本, "|替换姓名值|", "", -1)
		返回多项输入表单示例源码一文本 = strings.Replace(返回多项输入表单示例源码一文本, "|替换兴趣值|", "", -1)

	} else {
		返回多项输入表单示例源码一文本 = strings.Replace(返回多项输入表单示例源码一文本, "|替换姓名值|", "value="+全局.输入姓名一文本, -1)
		返回多项输入表单示例源码一文本 = strings.Replace(返回多项输入表单示例源码一文本, "|替换兴趣值|", "value="+全局.输入兴趣一文本, -1)

	}

	return

}

//===================================================================================
//===================================================================================

func A2模具一一返回Gui框架说明源码() (返回Gui框架说明源码一文本 string) {

	defer func() {
		var 链接网址一文本 = "https://github.com/jiqi136/PC-Gui"
		返回Gui框架说明源码一文本 = strings.Replace(返回Gui框架说明源码一文本, "|替换网址|", 链接网址一文本, -1)
	}()

	if 全局.程序英汉版一文本 == "汉语" {

		返回Gui框架说明源码一文本 = `	<br /><hr><hr>		本程序采用桌面 PC-GUI 框架.	<div style="max-width: 110%;">
		<h2 style=" font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: 为 AI 而生，原生支持类 Deepseek实时打字机流式输出的轻量桌面 GUI 框架！ 🎉</h2>
		
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
		
		
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 Apache 2.0 开源授权许可</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		本项目采用 宽松型<strong style="color: #2d3748; font-weight: 700;">Apache 2.0 许可证</strong>。这意味着您可以完全自由地使用框架源码。
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		您可以：(通俗解释 )<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">商业使用</strong>：允许将本作品及其衍生品用于商业目的，并进行销售。<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">修改分发</strong>：允许修改代码，并以开源或闭源的形式重新分发。<br />
		 
		</p>
		
		<hr>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		</p>
		

		<a href="|替换网址|"  target="_blank" style="display: flex; align-items: center; justify-content: space-between; background: #3b82f6; color: white; text-decoration: none; padding: 14px 20px; border-radius: 10px; font-weight: 500; transition: all 0.2s; box-shadow: 0 4px 6px -1px rgba(59, 130, 246, 0.2);">
				<span>公开源码 |替换网址|</span>
				<span style="font-size: 12px; opacity: 0.9;">github →</span>
			  </a>
			  <h4>觉得这个项目不错？请别忘了给它点一个 ⭐！您的支持是持续维护的动力。</h4>
			</div>
			<hr><hr>
			`

	} else {

		返回Gui框架说明源码一文本 = `					<br /><hr><hr>		This program adopts a desktop PC-GUI framework.<div style="max-width: 110%;">		
				<div style="max-width: 110%;">
		<h2 style="font-size: 2.5em; font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: Designed for AI — A lightweight desktop GUI framework that natively supports Gemini-like real-time typewriter streaming output! 🎉</h2>
		
		
		
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
		
		
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #2d3748; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 Apache 2.0 Open Source License</h2>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		This project is licensed under the <strong style="color: #2d3748; font-weight: 700;">Apache 2.0 License</strong>. This means you are completely free to use the framework's source code.
		</p>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;">
		You can: (plain English explanation)<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">Commercial Use</strong>: Use the work and its derivatives for commercial purposes and sell them.<br />
		 
		-   ✅ <strong style="color: #2d3748; font-weight: 700;">Modify and Distribute</strong>: Modify the code and redistribute it under open-source or closed-source licenses.<br />
		 
		</p>
		
		<hr>
		
		<p style="line-height: 1.8; color: #2d3748; margin: 0.8em 0; text-align: justify;"> </p>
		<a href="|替换网址|"  target="_blank" style="display: flex; align-items: center; justify-content: space-between; background: #3b82f6; color: white; text-decoration: none; padding: 14px 20px; border-radius: 10px; font-weight: 500; transition: all 0.2s; box-shadow: 0 4px 6px -1px rgba(59, 130, 246, 0.2);">
		<span>Open source. |替换网址|</span>
		<span style="font-size: 12px; opacity: 0.9;">github →</span>
	  	</a>
		  <h4>Like this project? Please don't forget to give it a ⭐! Your support is the motivation for continuous maintenance.</h4>
		
			</div>
			<hr><hr>`

	}

	return

}

func A1模具一一返回各国语言下拉表单源码() (返回一源码文本 string) {

	返回一源码文本 = ` 	   <div class="dropdown-container">
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

func A1模具一一返回程序界面网页源码() (返回界面网页一源码文本 string) {
			
	defer func() {
		返回界面网页一源码文本 = strings.Replace(返回界面网页一源码文本, "|替换JS|", A1模具一一返回JS源码(), -1)
	}() 


	返回界面网页一源码文本 = `

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

	return

}

func A1模具一一返回JS源码() (返回一源码文本 string) {

	defer func() {
		返回一源码文本 = strings.Replace(返回一源码文本, "|替换点符号|", "`", -1)
	}() 


	返回一源码文本 = `   const L1_str = document.getElementById('L1_str');
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

	return
}

func A1模具一一返回网页源码公用样式() (返回网页源码公用样式一文本 string) {
	返回网页源码公用样式一文本 = `			 
		<style type='text/css'>
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
			color: #0a6e8c;
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
