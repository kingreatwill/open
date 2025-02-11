[Mermaid](https://mermaid-js.github.io/mermaid/#/)
https://unpkg.com/browse/mermaid@8.4.5/dist/

## 说明
第一行的graph LR中graph指定是一个图，第二个LR指定图的方向，所有的方向关键词为:
TB - top bottom
BT - bottom top
RL - right left
LR - left right
TD - same as TB


节点默认只显示标识,但也可以通过如下方法控制其显示

A[aa bb] 显示字符串aa bb的方框
B(wo) 显示字符串wo的圆角框
C((我是C)) 显示我是C字符串的圆圈
D>我是D] 显示我是D的半方框
E{我是E} 显示我是E的正方形框

连线可以选择如下形式:
A-->B 箭头
A--B 无箭头线
A--hh dd--B或者A--|hh dd|B 线之间可以添加注释
A-.->B 虚线箭头
A-. hh .->B 添加了注释的虚线箭头
A==>B 加粗的箭头
A== hh ==>B 加注释加粗的箭头

子图可以使用subgraph id定义

```mermaid
    info
```
```mermaid
    gantt
      title Exclusive end dates (Manual date should end on 3d)
      dateFormat YYYY-MM-DD
      axisFormat %d
      section Section1
       2 Days: 1, 2019-01-01,2d
       Manual Date: 2, 2019-01-01,2019-01-03
```

```mermaid
    gantt
      title Inclusive end dates (Manual date should end on 4th)
      dateFormat YYYY-MM-DD
      axisFormat %d
      inclusiveEndDates
      section Section1
       2 Days: 1, 2019-01-01,2d
       Manual Date: 2, 2019-01-01,2019-01-03
```

```mermaid
      graph LR
      sid-B3655226-6C29-4D00-B685-3D5C734DC7E1["

      提交申请
      熊大
      "];
      class sid-B3655226-6C29-4D00-B685-3D5C734DC7E1 node-executed;
      sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A["
      负责人审批
      强子
      "];
      class sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A node-executed;
      sid-E27C0367-E6D6-497F-9736-3CDC21FDE221["
      DBA审批
      强子
      "];
      class sid-E27C0367-E6D6-497F-9736-3CDC21FDE221 node-executed;
      sid-BED98281-9585-4D1B-934E-BD1AC6AC0EFD["
      SA审批
      阿美
      "];
      class sid-BED98281-9585-4D1B-934E-BD1AC6AC0EFD node-executed;
      sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7["
      主管审批
      光头强
      "];
      class sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7 node-executed;
      sid-A1B3CD96-7697-4D7C-BEAA-73D187B1BE89["
      DBA确认
      强子
      "];
      class sid-A1B3CD96-7697-4D7C-BEAA-73D187B1BE89 node-executed;
      sid-3E35A7FF-A2F4-4E07-9247-DBF884C81937["
      SA确认
      阿美
      "];
      class sid-3E35A7FF-A2F4-4E07-9247-DBF884C81937 node-executed;
      sid-4FC27B48-A6F9-460A-A675-021F5854FE22["
      结束
      "];
      class sid-4FC27B48-A6F9-460A-A675-021F5854FE22 node-executed;
      sid-19DD9E9F-98C1-44EE-B604-842AFEE76F1E["
      SA执行1
      强子
      "];
      class sid-19DD9E9F-98C1-44EE-B604-842AFEE76F1E node-executed;
      sid-6C2120F3-D940-4958-A067-0903DCE879C4["
      SA执行2
      强子
      "];
      class sid-6C2120F3-D940-4958-A067-0903DCE879C4 node-executed;
      sid-9180E2A0-5C4B-435F-B42F-0D152470A338["
      DBA执行1
      强子
      "];
      class sid-9180E2A0-5C4B-435F-B42F-0D152470A338 node-executed;
      sid-03A2C3AC-5337-48A5-B154-BB3FD0EC8DAD["
      DBA执行3
      强子
      "];
      class sid-03A2C3AC-5337-48A5-B154-BB3FD0EC8DAD node-executed;
      sid-D5E1F2F4-306C-47A2-BF74-F66E3D769756["
      DBA执行2
      强子
      "];
      class sid-D5E1F2F4-306C-47A2-BF74-F66E3D769756 node-executed;
      sid-8C3F2F1D-F014-4F99-B966-095DC1A2BD93["
      DBA执行4
      强子
      "];
      class sid-8C3F2F1D-F014-4F99-B966-095DC1A2BD93 node-executed;
      sid-1897B30A-9C5C-4D5B-B80B-76A038785070["
      负责人确认
      梁静茹
      "];
      class sid-1897B30A-9C5C-4D5B-B80B-76A038785070 node-executed;
      sid-B3655226-6C29-4D00-B685-3D5C734DC7E1-->sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7;
      sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A-->sid-1897B30A-9C5C-4D5B-B80B-76A038785070;
      sid-E27C0367-E6D6-497F-9736-3CDC21FDE221-->sid-A1B3CD96-7697-4D7C-BEAA-73D187B1BE89;
      sid-BED98281-9585-4D1B-934E-BD1AC6AC0EFD-->sid-3E35A7FF-A2F4-4E07-9247-DBF884C81937;
      sid-19DD9E9F-98C1-44EE-B604-842AFEE76F1E-->sid-6C2120F3-D940-4958-A067-0903DCE879C4;
      sid-9180E2A0-5C4B-435F-B42F-0D152470A338-->sid-D5E1F2F4-306C-47A2-BF74-F66E3D769756;
      sid-03A2C3AC-5337-48A5-B154-BB3FD0EC8DAD-->sid-8C3F2F1D-F014-4F99-B966-095DC1A2BD93;
      sid-6C2120F3-D940-4958-A067-0903DCE879C4-->sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A;
      sid-1897B30A-9C5C-4D5B-B80B-76A038785070-->sid-4FC27B48-A6F9-460A-A675-021F5854FE22;
      sid-3E35A7FF-A2F4-4E07-9247-DBF884C81937-->sid-19DD9E9F-98C1-44EE-B604-842AFEE76F1E;
      sid-A1B3CD96-7697-4D7C-BEAA-73D187B1BE89-->sid-9180E2A0-5C4B-435F-B42F-0D152470A338;
      sid-A1B3CD96-7697-4D7C-BEAA-73D187B1BE89-->sid-03A2C3AC-5337-48A5-B154-BB3FD0EC8DAD;
      sid-D5E1F2F4-306C-47A2-BF74-F66E3D769756-->sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A;
      sid-8C3F2F1D-F014-4F99-B966-095DC1A2BD93-->sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A;
      sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7-->sid-BED98281-9585-4D1B-934E-BD1AC6AC0EFD;
      sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7-->sid-E27C0367-E6D6-497F-9736-3CDC21FDE221;
      sid-3E35A7FF-A2F4-4E07-9247-DBF884C81937-->sid-6C2120F3-D940-4958-A067-0903DCE879C4;
      sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7-->sid-4DA958A0-26D9-4D47-93A7-70F39FD7D51A;
      sid-7CE72B24-E0C1-46D3-8132-8BA66BE05AA7-->sid-4FC27B48-A6F9-460A-A675-021F5854FE22;
```

```mermaid
      graph TD
      A[Christmas] -->|Get money| B(Go shopping)
      B --> C{Let me thinksssssx<br/>sssssssssssssssssssuuu<br/>tttsssssssssssssssssssssss}
      C -->|One| D[Laptop]
      C -->|Two| E[iPhone]
      C -->|Three| F[Car]
```

```mermaid
    graph TD
    A[/Christmas\]
    A -->|Get money| B[\Go shopping/]
    B --> C{Let me thinksssss<br/>ssssssssssssssssssssss<br/>sssssssssssssssssssssssssss}
    C -->|One| D[/Laptop/]
    C -->|Two| E[\iPhone\]
    C -->|Three| F[Car]
```

```mermaid
    graph LR
47(SAM.CommonFA.FMESummary)-->48(SAM.CommonFA.CommonFAFinanceBudget)
37(SAM.CommonFA.BudgetSubserviceLineVolume)-->48(SAM.CommonFA.CommonFAFinanceBudget)
35(SAM.CommonFA.PopulationFME)-->47(SAM.CommonFA.FMESummary)
41(SAM.CommonFA.MetricCost)-->47(SAM.CommonFA.FMESummary)
44(SAM.CommonFA.MetricOutliers)-->47(SAM.CommonFA.FMESummary)
46(SAM.CommonFA.MetricOpportunity)-->47(SAM.CommonFA.FMESummary)
40(SAM.CommonFA.OPVisits)-->47(SAM.CommonFA.FMESummary)
38(SAM.CommonFA.CommonFAFinanceRefund)-->47(SAM.CommonFA.FMESummary)
43(SAM.CommonFA.CommonFAFinancePicuDays)-->47(SAM.CommonFA.FMESummary)
42(SAM.CommonFA.CommonFAFinanceNurseryDays)-->47(SAM.CommonFA.FMESummary)
45(SAM.CommonFA.MetricPreOpportunity)-->46(SAM.CommonFA.MetricOpportunity)
35(SAM.CommonFA.PopulationFME)-->45(SAM.CommonFA.MetricPreOpportunity)
41(SAM.CommonFA.MetricCost)-->45(SAM.CommonFA.MetricPreOpportunity)
41(SAM.CommonFA.MetricCost)-->44(SAM.CommonFA.MetricOutliers)
39(SAM.CommonFA.ChargeDetails)-->43(SAM.CommonFA.CommonFAFinancePicuDays)
39(SAM.CommonFA.ChargeDetails)-->42(SAM.CommonFA.CommonFAFinanceNurseryDays)
39(SAM.CommonFA.ChargeDetails)-->41(SAM.CommonFA.MetricCost)
39(SAM.CommonFA.ChargeDetails)-->40(SAM.CommonFA.OPVisits)
35(SAM.CommonFA.PopulationFME)-->39(SAM.CommonFA.ChargeDetails)
36(SAM.CommonFA.PremetricCost)-->39(SAM.CommonFA.ChargeDetails)
```

```mermaid
      graph TD
      9e122290_1ec3_e711_8c5a_005056ad0002("fa:fa-creative-commons My System | Test Environment")
      82072290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Shared Business Logic Server:Service 1")
      db052290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Shared Business Logic Server:Service 2")
      4e112290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Shared Report Server:Service 1")
      30122290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Shared Report Server:Service 2")
      5e112290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Dedicated Test Business Logic Server:Service 1")
      c1112290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs Dedicated Test Business Logic Server:Service 2")
      b7042290_1ec3_e711_8c5a_005056ad0002("fa:fa-circle [DBServer\SharedDbInstance].[SupportDb]")
      8f102290_1ec3_e711_8c5a_005056ad0002("fa:fa-circle [DBServer\SharedDbInstance].[DevelopmentDb]")
      0e102290_1ec3_e711_8c5a_005056ad0002("fa:fa-circle [DBServer\SharedDbInstance].[TestDb]")
      07132290_1ec3_e711_8c5a_005056ad0002("fa:fa-circle [DBServer\SharedDbInstance].[SharedReportingDb]")
      c7072290_1ec3_e711_8c5a_005056ad0002("fa:fa-server Shared Business Logic Server")
      ca122290_1ec3_e711_8c5a_005056ad0002("fa:fa-server Shared Report Server")
      68102290_1ec3_e711_8c5a_005056ad0002("fa:fa-server Dedicated Test Business Logic Server")
      f4112290_1ec3_e711_8c5a_005056ad0002("fa:fa-database [DBServer\SharedDbInstance]")
      d6072290_1ec3_e711_8c5a_005056ad0002("fa:fa-server DBServer")
      71082290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs DBServer\:MSSQLSERVER")
      c0102290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs DBServer\:SQLAgent")
      9a072290_1ec3_e711_8c5a_005056ad0002("fa:fa-cogs DBServer\:SQLBrowser")
      1d0a2290_1ec3_e711_8c5a_005056ad0002("fa:fa-server VmHost1")
      200a2290_1ec3_e711_8c5a_005056ad0002("fa:fa-server VmHost2")
      1c0a2290_1ec3_e711_8c5a_005056ad0002("fa:fa-server VmHost3")
      9e122290_1ec3_e711_8c5a_005056ad0002-->82072290_1ec3_e711_8c5a_005056ad0002
      9e122290_1ec3_e711_8c5a_005056ad0002-->db052290_1ec3_e711_8c5a_005056ad0002
      9e122290_1ec3_e711_8c5a_005056ad0002-->4e112290_1ec3_e711_8c5a_005056ad0002
      9e122290_1ec3_e711_8c5a_005056ad0002-->30122290_1ec3_e711_8c5a_005056ad0002
      9e122290_1ec3_e711_8c5a_005056ad0002-->5e112290_1ec3_e711_8c5a_005056ad0002
      9e122290_1ec3_e711_8c5a_005056ad0002-->c1112290_1ec3_e711_8c5a_005056ad0002
      82072290_1ec3_e711_8c5a_005056ad0002-->b7042290_1ec3_e711_8c5a_005056ad0002
      82072290_1ec3_e711_8c5a_005056ad0002-->8f102290_1ec3_e711_8c5a_005056ad0002
      82072290_1ec3_e711_8c5a_005056ad0002-->0e102290_1ec3_e711_8c5a_005056ad0002
      82072290_1ec3_e711_8c5a_005056ad0002-->c7072290_1ec3_e711_8c5a_005056ad0002
      db052290_1ec3_e711_8c5a_005056ad0002-->c7072290_1ec3_e711_8c5a_005056ad0002
      db052290_1ec3_e711_8c5a_005056ad0002-->82072290_1ec3_e711_8c5a_005056ad0002
      4e112290_1ec3_e711_8c5a_005056ad0002-->b7042290_1ec3_e711_8c5a_005056ad0002
      4e112290_1ec3_e711_8c5a_005056ad0002-->8f102290_1ec3_e711_8c5a_005056ad0002
      4e112290_1ec3_e711_8c5a_005056ad0002-->0e102290_1ec3_e711_8c5a_005056ad0002
      4e112290_1ec3_e711_8c5a_005056ad0002-->07132290_1ec3_e711_8c5a_005056ad0002
      4e112290_1ec3_e711_8c5a_005056ad0002-->ca122290_1ec3_e711_8c5a_005056ad0002
      30122290_1ec3_e711_8c5a_005056ad0002-->ca122290_1ec3_e711_8c5a_005056ad0002
      30122290_1ec3_e711_8c5a_005056ad0002-->4e112290_1ec3_e711_8c5a_005056ad0002
      5e112290_1ec3_e711_8c5a_005056ad0002-->8f102290_1ec3_e711_8c5a_005056ad0002
      5e112290_1ec3_e711_8c5a_005056ad0002-->68102290_1ec3_e711_8c5a_005056ad0002
      c1112290_1ec3_e711_8c5a_005056ad0002-->68102290_1ec3_e711_8c5a_005056ad0002
      c1112290_1ec3_e711_8c5a_005056ad0002-->5e112290_1ec3_e711_8c5a_005056ad0002
      b7042290_1ec3_e711_8c5a_005056ad0002-->f4112290_1ec3_e711_8c5a_005056ad0002
      8f102290_1ec3_e711_8c5a_005056ad0002-->f4112290_1ec3_e711_8c5a_005056ad0002
      0e102290_1ec3_e711_8c5a_005056ad0002-->f4112290_1ec3_e711_8c5a_005056ad0002
      07132290_1ec3_e711_8c5a_005056ad0002-->f4112290_1ec3_e711_8c5a_005056ad0002
      c7072290_1ec3_e711_8c5a_005056ad0002-->1d0a2290_1ec3_e711_8c5a_005056ad0002
      ca122290_1ec3_e711_8c5a_005056ad0002-->200a2290_1ec3_e711_8c5a_005056ad0002
      68102290_1ec3_e711_8c5a_005056ad0002-->1c0a2290_1ec3_e711_8c5a_005056ad0002
      f4112290_1ec3_e711_8c5a_005056ad0002-->d6072290_1ec3_e711_8c5a_005056ad0002
      f4112290_1ec3_e711_8c5a_005056ad0002-->71082290_1ec3_e711_8c5a_005056ad0002
      f4112290_1ec3_e711_8c5a_005056ad0002-->c0102290_1ec3_e711_8c5a_005056ad0002
      f4112290_1ec3_e711_8c5a_005056ad0002-->9a072290_1ec3_e711_8c5a_005056ad0002
      d6072290_1ec3_e711_8c5a_005056ad0002-->1c0a2290_1ec3_e711_8c5a_005056ad0002
      71082290_1ec3_e711_8c5a_005056ad0002-->d6072290_1ec3_e711_8c5a_005056ad0002
      c0102290_1ec3_e711_8c5a_005056ad0002-->d6072290_1ec3_e711_8c5a_005056ad0002
      c0102290_1ec3_e711_8c5a_005056ad0002-->71082290_1ec3_e711_8c5a_005056ad0002
      9a072290_1ec3_e711_8c5a_005056ad0002-->d6072290_1ec3_e711_8c5a_005056ad0002
      9a072290_1ec3_e711_8c5a_005056ad0002-->71082290_1ec3_e711_8c5a_005056ad0002
```

```mermaid
graph TB
	subgraph One
		a1-->a2
	end
```

```mermaid
graph TB
         subgraph one
         a1-->a2
         end
         subgraph two
         b1-->b2
         end
         subgraph three
         c1-->c2
         end
         c1-->a2
```

```mermaid
  graph TB
  A
  B
  subgraph foo[Foo SubGraph]
    C
    D
  end
  subgraph bar[Bar SubGraph]
    E
    F
  end
  G

  A-->B
  B-->C
  C-->D
  B-->D
  D-->E
  E-->A
  E-->F
  F-->D
  F-->G
  B-->G
  G-->D

  style foo fill:#F99,stroke-width:2px,stroke:#F0F
  style bar fill:#999,stroke-width:10px,stroke:#0F0
```

```mermaid
      graph LR
      456ac9b0d15a8b7f1e71073221059886[1051 AAA fa:fa-check]
      f7f580e11d00a75814d2ded41fe8e8fe[1141 BBB fa:fa-check]
      81dc9bdb52d04dc20036dbd8313ed055[1234 CCC fa:fa-check]
      456ac9b0d15a8b7f1e71073221059886 -->|Node| f7f580e11d00a75814d2ded41fe8e8fe
      f7f580e11d00a75814d2ded41fe8e8fe -->|Node| 81dc9bdb52d04dc20036dbd8313ed055
      click 456ac9b0d15a8b7f1e71073221059886 "/admin/user/view?id=1051" "AAA
      6000"
      click f7f580e11d00a75814d2ded41fe8e8fe "/admin/user/view?id=1141" "BBB
      600"
      click 81dc9bdb52d04dc20036dbd8313ed055 "/admin/user/view?id=1234" "CCC
      3000"
      style 456ac9b0d15a8b7f1e71073221059886 fill:#f9f,stroke:#333,stroke-width:4px
```

```mermaid
graph TD
A[Christmas] -->|Get money| B(Go shopping)
B --> C(Let me think...<br/>Do I want something for work,<br/>something to spend every free second with,<br/>or something to get around?)
C -->|One| D[Laptop]
C -->|Two| E[iPhone]
C -->|Three| F[Car]
click A "index.html#link-clicked" "link test"
click B testClick "click test"
classDef someclass fill:#f96;
class A someclass;
```

```mermaid
graph TD;

subgraph Line breaks <br/> don't work in <br/> Subgraphs
    inset[Line breaks <br/>work in <br/>Insets]
end

    inset-->A

  A(Line breaks<br/>work in<br/>rounded rec nodes)
  B{Line breaks <br/>work in<br/>decision nodes}
  C[Line breaks<br/>work in<br/> rectangles]
  D((Line breaks <br/>work in <br/>circles))
  E>Line breaks <br/>work in <br/>flag nodes]


    A-->B
    B--yes-->C
    B--no-->E
    C-->D
 


style A fill:#ed6,stroke:#333,stroke-width:2px;
style B fill:#ed6,stroke:#333,stroke-width:2px;
style C fill:#ed6,stroke:#333,stroke-width:2px;
style D fill:#ed6,stroke:#333,stroke-width:2px;
style E fill:#ed6,stroke:#333,stroke-width:2px;
style inset fill:#ed6,stroke:#333,stroke-width:2px;
```

```mermaid
graph TD
    A([stadium shape test])
    A -->|Get money| B([Go shopping])
    B --> C([Let me think...<br/>Do I want something for work,<br/>something to spend every free second with,<br/>or something to get around?])
    C -->|One| D([Laptop])
    C -->|Two| E([iPhone])
    C -->|Three| F([Car<br/>wroom wroom])
    click A "index.html#link-clicked" "link test"
    click B testClick "click test"
    classDef someclass fill:#f96;
    class A someclass;
```

```mermaid
    graph LR
    A[(cylindrical<br/>shape<br/>test)]
    A -->|Get money| B1[(Go shopping 1)]
    A -->|Get money| B2[(Go shopping 2)]
    A -->|Get money| B3[(Go shopping 3)]
    C[(Let me think...<br/>Do I want something for work,<br/>something to spend every free second with,<br/>or something to get around?)]
    B1 --> C
    B2 --> C
    B3 --> C
    C -->|One| D[(Laptop)]
    C -->|Two| E[(iPhone)]
    C -->|Three| F[(Car)]
    click A "index.html#link-clicked" "link test"
    click B testClick "click test"
    classDef someclass fill:#f96;
    class A someclass;
```

```mermaid
    graph LR
    A1[Multi<br>Line] -->|Multi<br>Line| B1(Multi<br>Line)
    C1[Multi<br/>Line] -->|Multi<br/>Line| D1(Multi<br/>Line)
    E1[Multi<br/>Line] -->|Multi<br/>Line| F1(Multi<br/>Line)
    A2[Multi<br>Line] -->|Multi<br>Line| B2(Multi<br>Line)
    C2[Multi<br/>Line] -->|Multi<br/>Line| D2(Multi<br/>Line)
    E2[Multi<br/>Line] -->|Multi<br/>Line| F2(Multi<br/>Line)
    linkStyle 0 stroke:DarkGray,stroke-width:2px
    linkStyle 1 stroke:DarkGray,stroke-width:2px
    linkStyle 2 stroke:DarkGray,stroke-width:2px
```


```mermaid
sequenceDiagram
participant Alice
participant Bob
participant John as John<br/>Second Line
rect rgb(200, 220, 100)
rect rgb(200, 255, 200)
Alice ->> Bob: Hello Bob, how are you?
Bob-->>John: How about you John?
end
Bob--x Alice: I am good thanks!
Bob-x John: I am good thanks!
Note right of John: Bob thinks a long<br/>long time, so long<br/>that the text does<br/>not fit on a row.
Bob-->Alice: Checking with John...
end
alt either this
Alice->>John: Yes
else or this
Alice->>John: No
else or this will happen
Alice->John: Maybe
end
par this happens in parallel
Alice -->> Bob: Parallel message 1
and
Alice -->> John: Parallel message 2
end
```

```mermaid
    sequenceDiagram
    participant 1 as multiline<br>using #lt;br#gt;
    participant 2 as multiline<br/>using #lt;br/#gt;
    participant 3 as multiline<br/>using #lt;br /#gt;
    participant 4 as multiline<br 	/>using #lt;br 	/#gt;
    1->>2: multiline<br>using #lt;br#gt;
    note right of 2: multiline<br>using #lt;br#gt;
    2->>3: multiline<br/>using #lt;br/#gt;
    note right of 3: multiline<br/>using #lt;br/#gt;
    3->>4: multiline<br/>using #lt;br /#gt;
    note right of 4: multiline<br/>using #lt;br /#gt;
    4->>1: multiline<br 	/>using #lt;br 	/#gt;
    note right of 1: multiline<br 	/>using #lt;br 	/#gt;
```


```mermaid
gantt
dateFormat  YYYY-MM-DD
axisFormat  %d/%m
title Adding GANTT diagram to mermaid
excludes weekdays 2014-01-10

section A section
Completed task            :done,    des1, 2014-01-06,2014-01-08
Active task               :active,  des2, 2014-01-09, 3d
Future task               :         des3, after des2, 5d
Future task2               :         des4, after des3, 5d

section Critical tasks
Completed task in the critical line :crit, done, 2014-01-06,24h
Implement parser and jison          :crit, done, after des1, 2d
Create tests for parser             :crit, active, 3d
Future task in critical line        :crit, 5d
Create tests for renderer           :2d
Add to mermaid                      :1d

section Documentation
Describe gantt syntax               :active, a1, after des1, 3d
Add gantt diagram to demo page      :after a1  , 20h
Add another diagram to demo page    :doc1, after a1  , 48h

section Clickable
Visit mermaidjs               :active, cl1, 2014-01-07,2014-01-10
Calling a Callback (look at the console log) :cl2, after cl1, 3d

click cl1 href "https://mermaidjs.github.io/"
click cl2 call ganttTestClick("test", test, test)

section Last section
Describe gantt syntax               :after doc1, 3d
Add gantt diagram to demo page      : 20h
Add another diagram to demo page    : 48h
```

```mermaid
    gantt
    dateFormat  YYYY-MM-DD
    axisFormat  %d/%m
    title       GANTT diagram with multiline section titles
    excludes    weekdays 2014-01-10

    section A section<br>multiline
    Completed task            : done,    des1, 2014-01-06,2014-01-08
    Active task               : active,  des2, 2014-01-09, 3d
    Future task               :          des3, after des2, 5d
    Future task2              :          des4, after des3, 5d

    section Critical tasks<br/>multiline
    Completed task in the critical line : crit, done, 2014-01-06, 24h
    Implement parser and jison          : crit, done, after des1, 2d
    Create tests for parser             : crit, active, 3d
    Future task in critical line        : crit, 5d
    Create tests for renderer           : 2d
    Add to mermaid                      : 1d

    section Documentation<br/>multiline
    Describe gantt syntax               : active, a1, after des1, 3d
    Add gantt diagram to demo page      : after a1, 20h
    Add another diagram to demo page    : doc1, after a1, 48h

    section Last section<br/>multiline
    Describe gantt syntax               : after doc1, 3d
    Add gantt diagram to demo page      : 20h
    Add another diagram to demo page    : 48h
```


```mermaid
gitGraph:
options
{
    "nodeSpacing": 150,
    "nodeRadius": 10
}
end
commit
branch newbranch
checkout newbranch
commit
commit
checkout master
commit
commit
merge newbranch
```

```mermaid
classDiagram
Class01 <|-- AveryLongClass : Cool
&lt;&lt;interface&gt;&gt; Class01
Class03 "0" *-- "0..n" Class04
Class05 "1" o-- "many" Class06
Class07 .. Class08
Class09 "many" --> "1" C2  : Where am i?
Class09 "0" --* "1..n" C3
Class09 --|> Class07
Class07 : equals()
Class07 : Object[] elementData
Class01 : #size()
Class01 : -int chimp
Class01 : +int gorilla
Class08 <--> C2: Cool label
class Class10 {
  &lt;&lt;service&gt;&gt;
  int id
  size()
}
```

```mermaid
classDiagram
    class Class01~T~
    Class01 : #size()
    Class01 : -int chimp
    Class01 : +int gorilla
    class Class10~T~ {
      &lt;&lt;service&gt;&gt;
      int id
      size()
    }
```

```mermaid
classDiagram
        Class01~T~ <|-- AveryLongClass : Cool
        &lt;&lt;interface&gt;&gt; Class01
        Class03~T~ "0" *-- "0..n" Class04
        Class05 "1" o-- "many" Class06
        Class07~T~ .. Class08
        Class09 "many" --> "1" C2  : Where am i?
        Class09 "0" --* "1..n" C3
        Class09 --|> Class07
        Class07 : equals()
        Class07 : Object[] elementData
        Class01 : #size()
        Class01 : -int chimp
        Class01 : +int gorilla
        Class08 <--> C2: Cool label
        class Class10 {
          &lt;&lt;service&gt;&gt;
          int id
          size()
        }
```

```mermaid
    stateDiagram
      State1
```

```mermaid
        stateDiagram
        [*] --> First
        state First {
            [*] --> second
            second --> [*]
        }
```

```mermaid
    stateDiagram
        State1: The state with a note
        note right of State1
            Important information! You can write
            notes.
        end note
        State1 --> State2
        note left of State2 : This is the note to the left.
```
