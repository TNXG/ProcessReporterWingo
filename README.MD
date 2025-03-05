> [!WARNING]  
> 此项目已停止维护，请使用现在仍处于维护状态的 [TNXG/Kizuna](https://github.com/TNXG/Kizuna)、[ChingCdesu/ShiroProcessReporter](https://github.com/ChingCdesu/ShiroProcessReporter) 或 [Innei/ProcessReporterCli](https://github.com/Innei/ProcessReporterCli)。


# Process Reporter For Windows Golang

Process Reporter for WinGo 是一个使用Golang开发的Windows应用程序。它旨在实时报告当前用户在PC上正在使用的前台应用程序的名称，以及正在播放的任何媒体信息，并将报告发送到 [Shiro](https://github.com/Innei/Shiro)。

正在学习 Golang 中，在实践中学习！~~（答辩代码（是）~~

~~面向ChatGPT编程（不是~~

程序基本实现逻辑沿用了[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)

本程序计划兼容新版本的Shiro实现的功能，其功能可以在[Shiro](https://github.com/Innei/Shiro)开发者Innei开发的Windows上报程序[ProcessReporterCli](https://github.com/Innei/ProcessReporterCli)中体验到 

~~被抢饭碗了，群地位Down（不是）~~

## ReportServer(此功能未经测试，如遇问题请issue)

### POST /api/report/media

接收并处理媒体报告，优先级大于SMTC。

**请求参数**

- `status`: 媒体状态，可以是"start"或"close"。
- `title`: 当`status`为"start"时，此参数表示媒体标题。
- `artist`: 当`status`为"start"时，此参数表示媒体艺术家。
- `SourceAppName`: 当`status`为"start"时，此参数表示源应用程序名称。

**请求示例**

```
POST /api/report/media
{
    "status": "start",
    "title": "Song Title",
    "artist": "Artist Name",
    "SourceAppName": "App Name"
}
```

**响应示例**

```json
{
    "message": "Report received and processed."
}
```

## Todo
- [x] 实现基本上报功能
- [x] 基本实现[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)的功能
- [ ] 完整实现[ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy)的功能
- [ ] 检测更新，自动下载`GetSmtcInfo Core`
- [ ] 基于网易云音乐手机端Api的播放信息获取（拿不到播放状态（
- [ ] 兼容新版本Shiro的客端化上传
- [ ] GUI

## 其他的Shiro报告程序

- [macOS](https://github.com/mx-space/ProcessReporterMac)
- [Linux](https://github.com/ttimochan/processforlinux) -- 由 ttimochan 使用 Rust 开发的 Linux 上报程序
- [ProcessReporterWinpy](https://github.com/TNXG/ProcessReporterWinpy) -- TNXG 使用 Python 开发的 Windows 上报程序
- [ProcessReporterCli](https://github.com/Innei/ProcessReporterCli) -- Innei 使用 Nodejs 开发的 Windows 上报程序
- [ProcessReporterWin](https://github.com/ChingCdesu/ProcessReporterWin) -- ChingCdesu 使用 C# 开发的 Windows 上报程序

## 其他问题

### Q：网易云音乐和QQ音乐不能上报

A：国内很多厂商不按照微软官方的媒体渠道上报媒体信息（即 Windows system media Transport control 集成）

`从 Windows 10 版本 1607 开始，默认情况下，使用 MediaPlayer 类或 AudioGraph 类播放媒体的 UWP 应用会自动与 SMTC 集成。 只需实例化 MediaPlayer 的新实例，并将 MediaSource、MediaPlaybackItem 或 MediaPlaybackList 分配给玩家的 Source 属性，然后用户将在 SMTC 中看到你的应用名称，并且可以使用 SMTC 控件播放、暂停和在播放列表中移动。  -- Windows文档`

这时需要其他方法来使本程序的media上报结构生效
- 通过插件使其通过SMTC上报信息

    QQ音乐：[FoskyM/QQMusicSMTC](https://github.com/FoskyM/QQMusicSMTC)

    网易云音乐v2：[MicroCBer/BetterNCM](https://github.com/MicroCBer/BetterNCM) 和 [BetterNCM/InfinityLink](https://github.com/BetterNCM/InfinityLink) 搭配使用
- 自行通过ReportServer上传信息
- Pr Welcome
## 有关提问

请根据您的需求选择适合的版本和使用方法。如果有任何问题，请随时提问。

提问前请先阅读[tvvocold/How-To-Ask-Questions-The-Smart-Way: How To Ask Questions The Smart Way 《提问的智慧》中文版 (github.com)](https://github.com/tvvocold/How-To-Ask-Questions-The-Smart-Way)
