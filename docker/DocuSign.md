## DocuSign/文档签署平台/电子签签/数字签名/文档签名
### documenso
https://github.com/documenso/documenso
### DocuSeal
https://github.com/docusealco/docuseal

```
docker run -d --name docuseal -e DATABASE_URL="postgres://postgres:pwd@postgresql/docuseal?sslmode=disable" -p 10004:3000 -v /data/dockerv/docuseal/data:/data --link postgresql --restart always docuseal/docuseal:1.3.2
```
默认用户名: admin@docuseal.co
密码: admin

集成的我们的程序中
docuseal 提供了 JS、Vue 和 React 三种语言的兼容方式。了不起这里列举VUE的添加方式：
```vue
<template>
    <DocusealForm
      :src="'https://docuseal.co/d/LEVGR9rhZYf86M'"
      :email="'signer@example.com'"
    />
  </template>

  <script>
  import { DocusealForm } from '@docuseal/vue'

  export default {
    name: 'App',
    components: {
      DocusealForm
    }
  }
  </script>
```
### OpenSign
https://github.com/OpenSignLabs/OpenSign
