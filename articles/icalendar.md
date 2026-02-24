## iCalendar

https://icalendar.org/

### 简述
iCalendar是"日历数据交换"的标准（RFC 5545）。此标准有时指的是“iCal”，即苹果公司出品的一款同名日历软件，这个软件也是此标准的一种实现方式。ICS (Internet Calendaring and Scheduling) 文件是以 iCalendar 标准编写的文本文件，可用于共享日历事件信息(标题、摘要、开始时间、结束时间等)，一般通过电子邮件给他人以共享会议请求或待办事项，收件人可以直接导入到自己的日历软件，通常是以 .ics 或 .ifb 为后缀的文本文件保存。

### 邮件头格式

通常 iCalendar 使用UTF-8字符编码，所以可以使用中文，也可以设置MIME中的charset参数来指定其他字符编码。iCalendar的MIME 类型是 text/calendar，Header示例如下：
```
Content-Type: text/calendar; method="REQUEST"; name="Appointment.ics"
MIME-Version: 1.0
Content-Transfer-Encoding: base64
Content-Description: attachment; filename="Appointment.ics"
Content-class: urn:content-classes:calendarmessage
Filename: Appointment.ics #文件名称
Path: Appointment.ics
```

Content-Type：标识邮件内容的格式
MIME-Version：表示MIME版本
Content-Transfer-Encoding：指定传输编码方式
Filename：文件名称

### 邮件体格式

iCalendar 邮件体由各类组件构成，组件分为事件（VEVENT）、待办事项列表（VTODO）、旅行事项（VJOURNAL）、时区信息（VZONE）、繁忙/空闲时间信息（VFREEBUSY）、警报（VALARM），其中警报可以包含于其他组件中。最常用的组件为事件（VEVENT）。
日历属性中必须包含 PRODID（软件信息）和 VERSION（icalendar版本）两个属性。

事件（VEVENT）
事件（VEVENT）是日历上一系列计划好的时间点。如果用户接受一个日历事件，则认为用户在该时间段是忙碌的，也可以应用在没有具体时间的日历事件上，如周年纪念、每日提醒。
属性说明：
```
DTSTART #开始时间，如果是循环事件，则为第一个事件的开始时间
DTEND #结束时间
SUMMARY #概要，标题
DESCRIPTION #详情描述，可用 html 语法
DTSTAMP #时间戳，表示事件的创建时间或更新时间
UID #唯一标识，取消与更新事件时用来找到唯一的日历事件
METHOD #操作，如果要取消事件，该值为 CANCEL
STATUS #事件状态，TENTATIVE - 试探、CONFIRMED - 确认、CANCELLED 取消
SEQUENCE #序列号，更新事件时需要指定新的序列号，如第一次更新指定 SEQUENCE:1
CLASS #保密类型，PRIVATE - 私有
CREATED #创建时间
LAST-MODIFIED #最后修改时间
LOCATION #地址
TRANSP #对于忙闲状态查询是否透明，OPAQUE - 不透明、TRANSPARENT - 透明
VALARM #警报对象
```

待办事项（VTODO）
不是所有软件都支持待办事项的，如 Outlook 就不能导出 VTODO 记录，导入时 VTODO 也会忽略
属性说明：
ACTION #动作，要执行的动作
TRIGGER #触发时间，数据格式与 DURATION 一致，如果是提前触发，需要在值前加负号 - ，如提前一天触发 -PT1440M
REPEAT #重复次数
DURATION #持续时间


旅行事项（VJOURNAL）
旅行事项（VJOURNAL）将一段描述文字关联到一个详细的日历日期上，可用于记录活动或成长日志，或描述待办事项的进展。旅行事项不会影响日历上的时间状态，不影响空闲和繁忙状态。只有少量程序支持VJOURNAL 。

时区信息（VTIMEZONE）
VTIMEZONE 用于存储时区在任何给定时间点（即夏令时和标准时间）的 UTC 偏移量的特定规则。由于UTC时间存在边界问题，对于周期性的规则不适用
DTSTART：#表示转换时区的开始时间，格式为YYYYMMDDTHHMMSS。
TZOFFSETFROM：#表示转换前的时区偏移量，格式为±HHMM。
TZOFFSETTO：#表示转换后的时区偏移量，格式为±HHMM。
TZNAME：#表示时区的名称

扩展属性
扩展属性 iCalendar 支持软件私有的扩展属性，这些属性只在特定软件中生效，在属性名前加 X- 前缀表示扩展属性。如：
X-MICROSOFT-CDO-BUSYSTATUS #忙碌状态
X-MICROSOFT-CDO-IMPORTANCE #重要程度
X-WR-CALNAM #通用扩展属性，表示本日历名称
X-WR-TIMEZONE #通用扩展属性，表示时区，值如：Asia/Shanghai

补充说明
Webcal 是 AFAIK 由 Apple 发明的 URI 方案，与“http”具有完全相同的语义，在一些浏览器中可直接打开日历软件，询问用户是否订阅日历主题。
`<a href='webcal://xxxx/calendar.ics' />`

### 完整示例
第一行必须是 BEGIN:VCALENDER ，最后一行必须是 END:VCALENDER ，这两行之间的数据称为 ical body。可以使用 # 或 // 添加注释，注意不能在其他元素行尾使用，只能注释整行
```
BEGIN:VCALENDAR #日历开始 
PRODID:-//Google Inc//Google Calendar 70.9054//EN #软件信息
VERSION:2.0 #遵循的 iCalendar 版本号
CALSCALE:GREGORIAN #历法：公历
METHOD:PUBLISH #方法：公开 PUBLISH - 公开，REQUEST - 请求，CANCEL - 取消
X-WR-CALNAME:yulanggong@gmail.com #这是一个通用扩展属性 表示本日历的名称
X-WR-TIMEZONE:Asia/Shanghai #通用扩展属性，表示时区
BEGIN:VEVENT #事件开始
DTSTART:20090305T112200Z #开始的日期时间
DTEND:20090305T122200Z #结束的日期时间
DTSTAMP:20140613T033914Z #有Method 属性时表示 实例创建时间，没有时表示最后修订的日期时间
UID:9r5p7q78uohmk1bbt0iedof9s4@google.com #UID
CLASS:PRIVATE #保密类型
CREATED:20090305T092105Z #创建的日期时间
DESCRIPTION:test #描述
LAST-MODIFIED:20090305T092130Z #最后修改日期时间
LOCATION:test #地址
SEQUENCE:1 #排列序号
STATUS:CONFIRMED #状态 TENTATIVE 试探 CONFIRMED 确认 CANCELLED 取消
SUMMARY: test #简介 一般是标题
TRANSP:OPAQUE #对于忙闲查询是否透明 OPAQUE 不透明 TRANSPARENT 透明
END:VEVENT #事件结束
END:VCALENDAR #日历结束
```