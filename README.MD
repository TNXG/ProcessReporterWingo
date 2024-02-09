# Process Reporter For Windows Golang

Process Reporter for WinGo 是一个使用Golang开发的Windows应用程序。它旨在实时报告当前用户在PC上正在使用的前台应用程序的名称，以及正在播放的任何媒体信息，并将报告发送到 [Shiro](https://github.com/Innei/Shiro)。

正在学习 Golang 中，在实践中学习！~~（答辩代码（是）~~

~~面向ChatGPT编程（不是~~

程序基本实现逻辑沿用了[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)

本程序计划兼容新版本的Shiro实现的功能，其功能可以在[Shiro](https://github.com/Innei/Shiro)开发者Innei开发的Windows上报程序[ProcessReporterCli](https://github.com/Innei/ProcessReporterCli)中体验到 

~~被抢饭碗了，群地位Down（不是）~~

## Todo
- [x] 实现基本上报功能
- [x] 基本实现[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)的功能
- [ ] 完整实现[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)的功能
- [ ] 检测更新，自动下载`GetSmtcInfo Core`
- [ ] 兼容新版本Shiro的客端化上传
- [ ] 接入未能在Python版本中实现的[ProcessReporterWinpy/NcmNowPlay](https://github.com/TNXG/ProcessReporterWinpy/tree/master/core/NcmNowPlay)
- [ ] GUI

## 其他系统的报告程序

- [macOS](https://github.com/mx-space/ProcessReporterMac)
- [Linux](https://github.com/ttimochan/processforlinux) -- 由 ttimochan 使用 Rust 开发的 Linux 上报程序
- [ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy) -- TNXG 使用 Python 开发的 Windows 上报程序
- [ProcessReporterCli](https://github.com/Innei/ProcessReporterCli) -- Innei 使用 Nodejs 开发的 Windows 上报程序
- [ProcessReporterWin](https://github.com/ChingCdesu/ProcessReporterWin) -- ChingCdesu 使用 C# 开发的 Windows 上报程序

请根据您的需求选择适合的版本和使用方法。如果有任何问题，请随时提问。

提问前请先阅读[tvvocold/How-To-Ask-Questions-The-Smart-Way: How To Ask Questions The Smart Way 《提问的智慧》中文版 (github.com)](https://github.com/tvvocold/How-To-Ask-Questions-The-Smart-Way)