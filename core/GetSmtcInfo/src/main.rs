use std::error::Error;
use windows::core::HSTRING;
use windows::Foundation::IAsyncOperation;
use windows::Media::Control::GlobalSystemMediaTransportControlsSessionManager;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    if let Err(e) = async_main().await {
        eprintln!("Error: {}", e);
    }
    Ok(())
}

async fn async_main() -> Result<(), Box<dyn Error>> {
    let session_manager_operation: IAsyncOperation<
        GlobalSystemMediaTransportControlsSessionManager,
    > = GlobalSystemMediaTransportControlsSessionManager::RequestAsync()?;
    let session_manager = session_manager_operation.get()?;
    let current_session = session_manager.GetCurrentSession()?;
    let media_properties_operation = current_session.TryGetMediaPropertiesAsync()?;
    let media_properties = media_properties_operation.get()?;

    // 将String转换为HSTRING
    let source_app_name_hstring: HSTRING = current_session.SourceAppUserModelId()?.into();
    let title_hstring: HSTRING = media_properties.Title()?.into();
    let artist_hstring: HSTRING = media_properties.Artist()?.into();

    // 使用HSTRING类型的变量
    let result = format!(
        "[ProcessReporterWingo.GetSmtcInfo.Title]{}[/ProcessReporterWingo.GetSmtcInfo.Title][ProcessReporterWingo.GetSmtcInfo.Artist]{}[/ProcessReporterWingo.GetSmtcInfo.Artist][ProcessReporterWingo.GetSmtcInfo.SourceAppName]{}[/ProcessReporterWingo.GetSmtcInfo.SourceAppName]",
        title_hstring.to_string_lossy().to_owned(),
        artist_hstring.to_string_lossy().to_owned(),
        source_app_name_hstring.to_string_lossy().to_owned()
    );
    println!("{}", result);

    Ok(())
}
