# PC-Gui: A lightweight desktop GUI framework built for AI, natively supporting real-time typewriter streaming output like Gemini! 🎉

[ **English** ] | [ [**中文版**](https://github.com/jiqi136/PC-Gui/blob/main/Chinese.md) ][远古洛书-无限数值幻方](https://github.com/jiqi136/Magic-square/blob/main/%E8%BF%9C%E5%8F%A4%E6%B4%9B%E4%B9%A6-%E6%97%A0%E9%99%90%E6%95%B0%E5%80%BC%E5%B9%BB%E6%96%B9-%E4%B8%AD%E6%96%87.md)
[Infinite Numerical Magic Square](https://github.com/jiqi136/Magic-square)

---


> 💡 **Core Philosophy: Rapid Development · Minimal Size · Native Performance · Empowering the Creation of Premium Tools Users Are Willing to Pay For**


On the desktop, the demand for efficient, practical tools from users has never waned, coupled with a strong willingness to pay.
PC-Gui aims to help developers quickly respond to this market need, using the simplest and most stable technologies to build compact yet powerful commercial-grade desktop applications.

---

### Core Technology Stack

Abandon complex dependencies and bloated third-party GUI libraries, returning to the essence of programming: **building desktop applications with a backend mindset**.
By using a stable Go backend to provide web services and drive a flexible web frontend, unparalleled lightweight design and performance are achieved.

| Component | Technical Details |
| :--- | :--- |
| **Backend Service** | Go language, providing local web services based on the standard library `net/http`. |
| **Frontend Interface** | Standard web technologies: HTML, JavaScript, CSS. |
| **Data Storage** | Locally encrypted SQLite database, lightweight and reliable. |


<br>


![Framework Diagram](https://github.com/jiqi136/PC-Gui/blob/main/add/tu1C.png?raw=true)


### Core Advantages & Multi-Solution Comparison

| Category | ✅ PC-Gui Advantages | ⚠️ Comparison with Other Solutions |
| :--- | :--- | :--- |
| **🚀 Zero-Dependency Runtime** | **Go backend** enables rapid development with strong typing for easy maintenance; cross-compiles into a single executable file, requiring no runtime or dependency installation from users—just double-click to run. | ⚠️ Requires users to pre-install and configure complex environments and dependencies like WebView2, Python, C++, Node.js. |
| **🎨 Interface Technology (HTML)** | **HTML** frontend can be quickly generated using a wealth of templates and AI tools, offering high efficiency and making it easy to create beautiful, modern visual styles. | Traditional GUI libraries often have outdated interfaces and high customization difficulty. |
| **AI Streaming Output** | Simple asynchronous processing is all that's needed to achieve streaming output of AI content, enhancing user experience. | Implementing streaming output typically requires handling complex callbacks or multi-threading. |
| **Markdown Rendering** | Perfectly renders Markdown format returned by AI, with syntax highlighting support for various languages. | Tools like Chatbox and Cherry often have relatively basic Markdown rendering and code highlighting effects. |
| **Single-File Deployment** | Using Go's standard library `embed`, all static resources (like images, CSS) can be directly bundled into a single executable file, served via the same HTML service. | Relies on bloated dependencies: requires external packaging tools, resulting in large binaries or scattered files, complicating deployment. |
| **📦 Minimal Size** | The packaged application size is only **10-25MB**, making distribution and download effortless. | ⚠️ Applications based on Electron / WebView2 typically exceed **100MB**. |
| **🧠 Ultra-Low Memory Usage** | Runtime memory consumption is approximately **8MB**, with near-zero CPU overhead, feeling light and fast. | ⚠️ Electron / WebView2 applications easily consume over **500MB** of memory. |
| **Code Security** | Go's compiled binaries, combined with garble obfuscation techniques, effectively prevent logic from being reverse-engineered. | Easily exposed: Scripting languages like Python and Node.js are extremely easy to decompile and expose, offering no protection for trade secrets. |
| **💻 Cross-Platform Compatibility** | Go natively supports Windows 7/10/11, Linux, and macOS, covering the broadest user base. | Solutions like Webview2 do not support older systems like Windows 7. |
| **Runtime Stability** | The core relies only on Go's official standard library, ensuring long-term stable operation without crashes. | Relies on numerous third-party libraries, posing risks related to version compatibility and stability. |
| **💯 Complete Control** | The core code depends only on Go's official standard library, **with no third-party GUI framework black boxes**, giving you complete control over the code for easy long-term maintenance and troubleshooting. | ⚠️ Depends on large third-party frameworks, resulting in code redundancy and difficulty troubleshooting complex issues. |
| **🌐 Globalization Support** | The interface is based on standard web pages, allowing direct use of the browser's **built-in translation feature** to easily support hundreds of languages worldwide. | Requires embedding multilingual text libraries, incurring significant workload. |
| **💡 Cross-Language Reusability** | The framework concept is clear, making it adaptable for any language that supports HTTP services (e.g., C#, Python, Rust). | Frameworks deeply coupled with specific languages or platforms are difficult to migrate. |


<br>

## To Developers

In today's world, where the AI wave is sweeping the globe and the job market faces challenges, mastering a skill that enables rapid value creation is crucial.

I hope PC-Gui, this lightweight and efficient framework, can become a powerful tool in your hands, helping you quickly turn ideas into reality, develop desktop utilities that users are willing to pay for, and ultimately achieve personal value and commercial success.

---



---

## 📥 Demo Program & Open Source Code [Download Link]:

[GitHub Repository Download https://github.com/jiqi136/PC-Gui/releases](https://github.com/jiqi136/PC-Gui/releases)


[LanDisk Download https://wwbrl.lanzoum.com/b0fqa7hxe](https://wwbrl.lanzoum.com/b0fqa7hxe)  
Password: `1122`

### Go Language Source Code Build Instructions
Building the project requires only the following two core files:
- **`main.go`**  
  The program entry file, containing the main business logic implementation.
- **`go.mod`**  
  The Go module definition file, used for managing project dependencies and version information.
Download the above files and import them into the project directory to start building.



SHA256：317100e77d438f10c94663ab97fed47bd76e04b01a3350508980e87e9ef941d8
![Online Antivirus Platform Scan](https://github.com/jiqi136/PC-Gui/blob/main/add/tu2C.png?raw=true)


---
## Like this project? Don't forget to give it a ⭐! Your support is the motivation for continued maintenance.

---
##  Projects using desktop PC-GUI frameworks :

[Ai-Assistant https://github.com/jiqi136/Ai-Assistant](https://github.com/jiqi136/Ai-Assistant) 

[syntax-query https://github.com/jiqi136/syntax-query](https://github.com/jiqi136/syntax-query) 

[English-study https://github.com/jiqi136/English-study](https://github.com/jiqi136/English-study) 

[Magic-square https://github.com/jiqi136/Magic-square](https://github.com/jiqi136/Magic-square) 

---


## 📜 Apache 2.0 Open Source License

This project is licensed under the **Apache 2.0** License. This means you are completely free to use the framework's source code.

In simple terms, you can:
-   ✅ **Commercial Use**: Use the work and its derivatives for commercial purposes and sell them.
-   ✅ **Modify and Distribute**: Modify the code and redistribute it under open-source or closed-source terms.


---
