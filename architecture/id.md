## UUID：Universally Unique Identifier（通用唯一标识符）
## UDID：Unique Device Identifier（设备唯一标识符）
## ULID：Universally Unique Lexicographically Sortable Identifier（通用唯一词典分类标识符）

[ulid](https://github.com/ulid/spec)

## nanoid
[nanoid](https://github.com/ai/nanoid)
```
crypto.randomUUID         28,387,114 ops/sec
uid/secure                 8,633,795 ops/sec
@lukeed/uuid               6,888,704 ops/sec
nanoid                     6,166,399 ops/sec
customAlphabet             3,290,342 ops/sec
uuid v4                    1,662,373 ops/sec
secure-random-string         415,340 ops/sec
uid-safe.sync                400,875 ops/sec
cuid                         212,669 ops/sec
shortid                       53,453 ops/sec
```
[Other Programming Languages](https://github.com/ai/nanoid#other-programming-languages)



## 其它个人实现
[一种对数据库友好的GUID的变种使用方法](https://www.cnblogs.com/ensleep/p/17745166.html)
```csharp
public static class SeqGuid
{
    /// <summary>
    /// 生成BSV的GUID。
    /// </summary>
    /// <returns></returns>
    public static string NewGuid()
    {

        var gid = Guid.NewGuid().ToByteArray();// 获取唯一的guid，对应uuid的版本应该是v4。此处直接获取其byte数组。
        var dtvalue = DateTime.Now.Ticks;//获取当前时间到1年1月1日的总ticks数，ticks单位是100ns，即万分之一毫秒。
        var dtbytes = BitConverter.GetBytes(dtvalue);// 将ticks时间戳转换为字节数组，默认是小端。
        var bytes = new Byte[gid.Length + dtbytes.Length];// 实例化新的数字，用以存放时间值和GUID值。

        // 因为BitConverter.GetBytes获得的Byte[]是小端，不符合排序要求，所以，要逆序写入bytes数组中，形成大端的方式。
        // 将时间值放入bytes数组中。
        for (long i = 0; i < dtbytes.Length; i++)
        {
            var cvalue = dtbytes[dtbytes.Length - i - 1];
            bytes[i] = cvalue;
        }

        // 将guid的值，放入bytes数组中。
        gid.CopyTo(bytes, dtbytes.Length);
        // 将值转换为base64，主要原因是，前端、数据库比较容易处理字符串类型的数据。
        var b64 = Convert.ToBase64String(bytes);
        // 将无序的base64转换为有序的伪base64格式。

        var ss = b64.ToArray();
        for (var i = 0; i < ss.Length; i++)
        {
            ss[i] = dic[ss[i]];
        }
        return new string(ss);
    }
    /// <summary>
    /// 仿base64的有序字典，其与base64相似，使用有限的字符，表示6bit的二进制，不足的地方补=。但是，与base64的区别是，字符串是按从小到大的次序表示000000到111111的数值的。
    /// </summary>
    public static readonly Dictionary<char, char> dic = new Dictionary<char, char>()
    {
        {'A','$'},{'B','-'},{'C','0'},{'D','1'},{'E','2'},{'F','3'},{'G','4'},{'H','5'},{'I','6'},{'J','7'},{'K','8'},
        {'L','9'},{'M','A'},{'N','B'},{'O','C'},{'P','D'},{'Q','E'},{'R','F'},{'S','G'},{'T','H'},{'U','I'},{'V','J'},
        {'W','K'},{'X','L'},{'Y','M'},{'Z','N'},{'a','O'},{'b','P'},{'c','Q'},{'d','R'},{'e','S'},{'f','T'},{'g','U'},
        {'h','V'},{'i','W'},{'j','X'},{'k','Y'},{'l','Z'},{'m','a'},{'n','b'},{'o','c'},{'p','d'},{'q','e'},{'r','f'},
        {'s','g'},{'t','h'},{'u','i'},{'v','j'},{'w','k'},{'x','l'},{'y','m'},{'z','n'},{'0','o'},{'1','p'},{'2','q'},
        {'3','r'},{'4','s'},{'5','t'},{'6','u'},{'7','v'},{'8','w'},{'9','x'},{'+','y'},{'/','z'},{'=','!'}
    };
}
```

- [sql server GUID的NEWSEQUENTIALID()](https://learn.microsoft.com/zh-cn/sql/t-sql/functions/newsequentialid-transact-sql?view=sql-server-ver16)

NEWSEQUENTIALID() 和 NEWID()都可以产生uniqueidentifier类型的，GUID.NEWID()产生的GUID是无序的,随机的。
而NEWSEQUENTIALID()是SQL SERVER2005新特性,NEWSEQUENTIALID是基于硬件（一定程度上）生成的GUID以十六进制间隔递增.

