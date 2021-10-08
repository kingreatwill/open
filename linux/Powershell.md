
[Powershell快速入门](https://www.cnblogs.com/lsgxeva/p/9309576.html)
[PowerShell 文档](https://docs.microsoft.com/zh-cn/powershell/)
https://github.com/PowerShell/PowerShell

## 别名

[兼容性别名](https://docs.microsoft.com/zh-cn/powershell/scripting/learn/compatibility-aliases?view=powershell-7.1)

[兼容性别名2](https://github.com/PowerShell/PowerShell/tree/master/docs/learning-powershell#map-book-for-experienced-bash-users)

获取当前会话中的所有别名`Get-Alias`

```
ac = Add-Content
asnp = Add-PSSnapin
clc = Clear-Content
cli = Clear-Item
clp = Clear-ItemProperty
clv = Clear-Variable
cpi = Copy-Item
cpp = Copy-ItemProperty
cvpa = Convert-Path
diff = Compare-Object
epal = Export-Alias
epcsv = Export-Csv
fc = Format-Custom
fl = Format-List
foreach = ForEach-Object
% = ForEach-Object
ft = Format-Table
fw = Format-Wide
gal = Get-Alias
gc = Get-Content
gci = Get-ChildItem
gcm = Get-Command
gdr = Get-PSDrive
ghy = Get-History
gi = Get-Item
gl = Get-Location
gm = Get-Member
gp = Get-ItemProperty
gps = Get-Process
group = Group-Object
gsv = Get-Service
gsnp = Get-PSSnapin
gu = Get-Unique
gv = Get-Variable
gwmi = Get-WmiObject
iex = Invoke-Expression
ihy = Invoke-History
ii = Invoke-Item
ipal = Import-Alias
ipcsv = Import-Csv
mi = Move-Item
mp = Move-ItemProperty
nal = New-Alias
ndr = New-PSDrive
ni = New-Item
nv = New-Variable
oh = Out-Host
rdr = Remove-PSDrive
ri = Remove-Item
rni = Rename-Item
rnp = Rename-ItemProperty
rp = Remove-ItemProperty
rsnp = Remove-PSSnapin
rv = Remove-Variable
rvpa = Resolve-Path
sal = Set-Alias
sasv = Start-Service
sc = Set-Content
select = Select-Object
si = Set-Item
sl = Set-Location
sleep = Start-Sleep
sort = Sort-Object
sp = Set-ItemProperty
spps = Stop-Process
spsv = Stop-Service
sv = Set-Variable
tee = Tee-Object
where = Where-Object
? = Where-Object
write = Write-Output
cat = Get-Content
cd = Set-Location
clear = Clear-Host
cp = Copy-Item
h = Get-History
history = Get-History
kill = Stop-Process
lp = Out-Printer
ls = Get-ChildItem
mount = New-PSDrive
mv = Move-Item
popd = Pop-Location
ps = Get-Process
pushd = Push-Location
pwd = Get-Location
r = Invoke-History
rm = Remove-Item
rmdir = Remove-Item
echo = Write-Output
cls = Clear-Host
chdir = Set-Location
copy = Copy-Item
del = Remove-Item
dir = Get-ChildItem
erase = Remove-Item
move = Move-Item
rd = Remove-Item
ren = Rename-Item
set = Set-Variable
type = Get-Content
```
> 参考[PowerShell命令](https://www.cnblogs.com/valin/articles/1621206.html)

Invoke-WebRequest，利用它我们可以获取网页内容、下载文件甚至是填写表单。这个命令的别名是iwr、curl和wget。
`iwr -useb https://raw.githubusercontent.com/filebrowser/get/master/get.ps1 | iex`