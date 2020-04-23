D:\Android\SDK\tools\emulator.exe -list-avds
set /p ver=Input AVD(default Pixel_10.0):
if "%ver%"=="" (
    set ver="Pixel_10.0"
)
D:\Android\SDK\tools\emulator.exe -netdelay none -netspeed full -avd %ver%
