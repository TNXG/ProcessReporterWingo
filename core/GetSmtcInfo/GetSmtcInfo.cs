using Windows.Media.Control;

class Program
{
    static string GetMediaInfoSync()
    {
        var sessionManager = GlobalSystemMediaTransportControlsSessionManager.RequestAsync().GetAwaiter().GetResult();
        var currentSession = sessionManager.GetCurrentSession();
        var mediaProperties = currentSession.TryGetMediaPropertiesAsync().GetAwaiter().GetResult();
        string sourceAppName = currentSession.SourceAppUserModelId;
        string title = mediaProperties.Title ?? "<Null>";
        string artist = mediaProperties.Artist ?? "<Null>";
        sourceAppName = sourceAppName ?? "<Null>";
        return $"[ProcessReporterWingo.GetSmtcInfo.Title]{title}[/ProcessReporterWingo.GetSmtcInfo.Title][ProcessReporterWingo.GetSmtcInfo.Artist]{artist}[/ProcessReporterWingo.GetSmtcInfo.Artist][ProcessReporterWingo.GetSmtcInfo.SourceAppName]{sourceAppName}[/ProcessReporterWingo.GetSmtcInfo.SourceAppName]";
    }

    static void Main()
    {
        string result = GetMediaInfoSync();
        Console.OutputEncoding = System.Text.Encoding.UTF8;
        Console.WriteLine(result);
    }
}