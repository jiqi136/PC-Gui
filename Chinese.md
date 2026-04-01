
---

# PC-Gui: 为 AI 而生，原生支持类 Deepseek实时打字机流式输出的轻量桌面 GUI 框架！ 🎉

[ [**English**](https://github.com/jiqi136/PC-Gui) ] | [ **中文版** ]




> 💡 **核心理念：极速开发 · 极致体积 · 原生性能 · 助力打造用户愿意付费的优质工具**


在桌面端，用户对高效、实用工具的需求从未减弱，并且拥有强烈的付费意愿。
PC-Gui 旨在帮助开发者快速响应这一市场需求，用最简单、最稳定的技术，构建出小巧而强大的商业级桌面应用。

---

### 核心技术栈

摒弃了复杂的依赖，臃肿的第三方 GUI 库，回归编程的本质：**用后端思维构建桌面应用**。
通过一个稳定的 Go 后端提供 Web 服务，驱动一个灵活的 Web 前端界面，实现了无与伦比的轻量化与性能。

| 组件 | 技术详情 |
| :--- | :--- |
| **后端服务** | Go 语言，基于标准库 `net/http` 提供本地 Web 服务。 |
| **前端界面** | HTML, JavaScript, CSS 标准 Web 技术。 |
| **数据存储** | 本地加密的 SQLite 数据库，轻量、可靠。 |


<br>


![框架图](https://github.com/jiqi136/PC-Gui/blob/main/add/tu1A.png?raw=true)

### 核心优势 & 多方案对比

| 类别 | ✅ PC-Gui 优势 | ⚠️ 其他方案对比 |
| :--- | :--- | :--- |
| **🚀 零依赖运行** | **后端Go 语言**极速开发，强类型易于维护；交叉编译，生成单一可执行文件，无需用户安装任何运行时或依赖库，双击即可运行。 | ⚠️需要用户预装并配置  WebView2, Python、C++, Node.js 等复杂的环境和依赖。 |
| **🎨界面技术 (HTML)** | **HTML** 前端界面可借助海量模板与 AI 工具快速生成，不仅效率极高，还能轻松打造出精美、现代的视觉风格。 | 传统 GUI 库界面通常较为陈旧，自定义难度高。 |
| **AI 流式输出** | 仅需简单的异步处理，即可实现 AI 内容的流式输出，提升用户体验。 | 实现流式输出通常需要处理复杂的回调或多线程。 |
| **Markdown 渲染** | 完美渲染 AI 返回的 Markdown 格式，并支持各类语言的语法高亮。 | Chatbox、Cherry等对 Markdown 渲染及代码高亮效果较为朴素。|
| **单文件部署** | 利用 Go 标准库中的 embed，可以将所有静态资源（如图片、CSS 等）直接打包到单一可执行文件中，并复用 HTML 服务进行直接访问。|依赖臃肿:需借助外部工具打包，产物体积庞大或文件零散，部署复杂。 |
| **📦 极致体积** | 应用打包后体积仅 **10-25MB**，分发和下载毫无压力。 | ⚠️ 基于 Electron / WebView2 的应用体积普遍在 **100MB** 以上。 |
| **🧠 超低内存占用** | 运行时内存占用仅约 **8MB**，CPU 开销近乎为零，轻快如飞。 | ⚠️ Electron / WebView2 应用内存占用轻松达到 **500MB** 以上。 |
| **代码安全性** |  Go 编译后的二进制文件,结合 garble 混淆技术，有效防止逻辑被反编译。|易泄露:Python、Node.js 脚本语言极易被反编译、扒光，毫无商业机密。 |
| **💻跨平台兼容** | Go 语言原生支持 Windows 7/10/11, Linux, macOS，覆盖最广泛的用户群体。 | Webview2 等方案不支持 Windows 7 等旧版系统。 |
| **运行稳定性** | 核心仅依赖 Go 官方标准库，可长期稳定运行不崩溃。 | 依赖大量第三方库，版本兼容性和稳定性存在风险。 |
| **💯 完全掌控** | 核心代码仅依赖 Go 官方标准库，**无任何第三方 GUI 框架黑盒**，代码完全自主可控，便于长期维护与排查问题。 | ⚠️ 依赖大型第三方框架，代码冗余，遇到疑难杂症时排查困难。 |
| **🌐 全球化支持** | 界面基于标准网页，可直接利用浏览器的**内置翻译功能**，轻松支持全球数百种语言。 | 需要内置多语言文本库，工作量巨大。 |
| **💡跨语言复用** | 框架思路清晰，任何支持 HTTP 服务的语言（如 C#, Python, Rust）均可借鉴。 | 框架与特定语言或平台深度绑定，难以迁移。 |


<br>

## 致开发者

在 AI 浪潮席卷全球、就业市场面临挑战的今天，掌握一门能够快速创造价值的技能至关重要。

希望 PC-Gui 这套轻量、高效的框架，能成为您手中的利器，帮助您快速将创意变为现实，开发出用户愿意付费的桌面实用工具，最终实现实现个人价值与商业创收。

---



---

## 📥演示程序与公开源码【下载地址】:

[蓝盘下载 https://wwbrl.lanzoum.com/b0fqa4uja](https://wwbrl.lanzoum.com/b0fqa4uja)  
密码：`epyz`

[123盘下载 https://www.123865.com/s/nKjJjv-gYg5d](https://www.123865.com/s/nKjJjv-gYg5d)


### Go 语言源码构建说明
构建项目仅需以下两个核心文件：
- **`main.go`**  
  程序入口文件，包含主要的业务逻辑实现。
- **`go.mod`**  
  Go 模块定义文件，用于管理项目依赖与版本信息。
将上述文件下载并导入项目目录即可开始构建。


SHA256：1c288811bbffe0cd39968e5d0045560271bad62e4ccbf7988e707fd1aaf98376
![在线杀软平台检测](https://github.com/jiqi136/PC-Gui/blob/main/add/tu2.png?raw=true)



---
## 觉得这个项目不错？请别忘了给它点一个 ⭐！您的支持是持续维护的动力。


---


## 📜 MIT 开源授权许可

本项目采用 **MIT 许可证**。这意味着您可以完全自由地使用框架源码。

您可以：(通俗解释 )
-   ✅ **商业使用**：允许将本作品及其衍生品用于商业目的，并进行销售。
-   ✅ **修改分发**：允许修改代码，并以开源或闭源的形式重新分发。
-   ✅ **私人使用**：允许在私人场合使用和修改。

**唯一要求**：在您的软件副本或重要部分中，包含原始的版权和许可声明。

---

### 许可原文 (Official License Text)

```text
MIT License

Copyright (c) [2026] [github.com/jiqi136]

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
```
