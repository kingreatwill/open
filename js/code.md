## 密码显示
添加书签，网址内容填入以下代码 或者 在控制台执行以下代码
```
javascript:"use strict";!function(){var e,t;e=document.getElementsByTagName("input");for(var a=0;a<e.length;a++)if(t=e[a],"password"==t.type.toLowerCase())try{t.type="text"}catch(e){var r,n;r=document.createElement("input"),n=t.attributes;for(var o=0;o<n.length;o++){var i,c,d;i=n[o],c=i.nodeName,d=i.nodeValue,"type"!=c.toLowerCase()&&"height"!=c&&"width"!=c&!!d&&(r[c]=d)}t.parentNode.replaceChild(r,t)}}();
```