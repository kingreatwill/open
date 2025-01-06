
127.0.0.1 —> 127.255.255.254（去掉0和255） 的范围都是本地环回地址/回环地址
RFC 1122：互联网标准文档 RFC 1122 规定，IP 地址 127.0.0.1 用于表示本地主机（localhost）。这个地址范围 127.0.0.0/8（即 127.0.0.0 到 127.255.255.255）都被预留用于本地回环通信。
需要注意的是，环回地址范围内不包含的两个IP：127.0.0.0 和 127.255.255.255。
原因很简单：主机号全0代表「网络地址」，全1代表「广播地址」。

```go
package main

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

// IPPair 存储IP对
type IPPair struct {
	IPv6 net.IP
	IPv4 net.IP
}

// IPConverter IP地址转换器
type IPConverter struct {
	upperCase bool
}

// NewIPConverter 创建新的转换器
func NewIPConverter(upperCase bool) *IPConverter {
	return &IPConverter{
		upperCase: upperCase,
	}
}

// ToBase36 将IPv6和IPv4组合转换为36进制
func (c *IPConverter) ToBase36(ipv6, ipv4 string) (string, error) {
	// 解析IPv6
	var ip6 net.IP
	if ipv6 != "" {
		ip6 = net.ParseIP(ipv6)
		if ip6 != nil && ip6.To4() != nil {
			return "", fmt.Errorf("invalid IPv6 address: %s", ipv6)
		}
	}

	// 解析IPv4
	var ip4 net.IP
	if ipv4 != "" {
		ip4 = net.ParseIP(ipv4)
		if ip4 == nil || ip4.To4() == nil {
			return "", fmt.Errorf("invalid IPv4 address: %s", ipv4)
		}
		ip4 = ip4.To4()
	}

	if ip6 == nil && ip4 == nil {
		return "", fmt.Errorf("no valid IP address provided")
	}

	// 创建大整数存储组合结果
	result := new(big.Int)

	if ip6 != nil {
		// 存储IPv6（如果有）
		result.SetBytes(ip6.To16())
		// 左移32位，为IPv4预留空间
		result.Lsh(result, 32)
	}

	if ip4 != nil {
		// 将IPv4添加到结果中
		ipv4Int := new(big.Int).SetBytes(ip4)
		result.Or(result, ipv4Int)
	}

	// 转换为36进制
	base36 := result.Text(36)
	if c.upperCase {
		base36 = strings.ToUpper(base36)
	}

	return base36, nil
}

// FromBase36 从36进制解析回IPv6和IPv4
func (c *IPConverter) FromBase36(s string) (*IPPair, error) {
	// 解析36进制字符串
	ipInt := new(big.Int)
	_, ok := ipInt.SetString(s, 36)
	if !ok {
		return nil, fmt.Errorf("invalid base36 string: %s", s)
	}

	// 创建返回结果
	pair := &IPPair{}

	// 提取IPv4部分（最后32位）
	ipv4Int := new(big.Int).And(ipInt, new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 32), big.NewInt(1)))
	if ipv4Int.BitLen() > 0 {
		ipv4Bytes := ipv4Int.Bytes()
		pair.IPv4 = make(net.IP, 4)
		copy(pair.IPv4[4-len(ipv4Bytes):], ipv4Bytes)
	}

	// 提取IPv6部分（去除最后32位）
	ipv6Int := new(big.Int).Rsh(ipInt, 32)
	if ipv6Int.BitLen() > 0 {
		ipv6Bytes := ipv6Int.Bytes()
		if len(ipv6Bytes) > 0 {
			pair.IPv6 = make(net.IP, 16)
			copy(pair.IPv6[16-len(ipv6Bytes):], ipv6Bytes)
		}
	}

	return pair, nil
}

func main() {
	converter := NewIPConverter(false)

	// 测试场景
	testCases := []struct {
		name string
		ipv6 string
		ipv4 string
	}{
		{
			name: "IPv6 + IPv4",
			ipv6: "2001:db8::1",
			ipv4: "192.168.1.1",
		},
		{
			name: "只有IPv4",
			ipv6: "",
			ipv4: "192.168.1.1",
		},
		{
			name: "只有IPv6",
			ipv6: "2001:db8::1",
			ipv4: "",
		},
		{
			name: "特殊地址组合",
			ipv6: "fe80::1",
			ipv4: "127.0.0.1",
		},
		{
			name: "最大值组合",
			ipv6: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
			ipv4: "255.255.255.255",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试场景: %s\n", tc.name)
		fmt.Printf("原始 IPv6: %s\n", tc.ipv6)
		fmt.Printf("原始 IPv4: %s\n", tc.ipv4)

		// 转换为36进制
		base36, err := converter.ToBase36(tc.ipv6, tc.ipv4)
		if err != nil {
			fmt.Printf("转换错误: %v\n", err)
			continue
		}
		fmt.Printf("Base36: %s\n", base36)

		// 转换回IP对
		pair, err := converter.FromBase36(base36)
		if err != nil {
			fmt.Printf("反向转换错误: %v\n", err)
			continue
		}

		fmt.Printf("解析后 IPv6: %s\n", pair.IPv6)
		fmt.Printf("解析后 IPv4: %s\n", pair.IPv4)

		// 验证结果
		if tc.ipv6 != "" && pair.IPv6.String() != net.ParseIP(tc.ipv6).String() {
			fmt.Printf("警告: IPv6不匹配!\n")
		}
		if tc.ipv4 != "" && pair.IPv4.String() != net.ParseIP(tc.ipv4).String() {
			fmt.Printf("警告: IPv4不匹配!\n")
		}
	}
}

```

校验
```go
package main

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

// IPPair 存储IP对
type IPPair struct {
	IPv6 net.IP
	IPv4 net.IP
}

// IPConverter IP地址转换器
type IPConverter struct {
	upperCase bool
	maxIPv6   *big.Int // IPv6最大值 (2^128 - 1)
	maxIPv4   *big.Int // IPv4最大值 (2^32 - 1)
}

// NewIPConverter 创建新的转换器
func NewIPConverter(upperCase bool) *IPConverter {
	// 计算IPv6最大值 (2^128 - 1)
	maxIPv6 := new(big.Int).Lsh(big.NewInt(1), 128)
	maxIPv6.Sub(maxIPv6, big.NewInt(1))

	// 计算IPv4最大值 (2^32 - 1)
	maxIPv4 := new(big.Int).Lsh(big.NewInt(1), 32)
	maxIPv4.Sub(maxIPv4, big.NewInt(1))

	return &IPConverter{
		upperCase: upperCase,
		maxIPv6:   maxIPv6,
		maxIPv4:   maxIPv4,
	}
}

// validateIPv4 验证IPv4地址
func validateIPv4(ip string) (net.IP, error) {
	if ip == "" {
		return nil, nil
	}

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return nil, fmt.Errorf("invalid IPv4 address: %s", ip)
	}

	ipv4 := parsed.To4()
	if ipv4 == nil {
		return nil, fmt.Errorf("not an IPv4 address: %s", ip)
	}

	return ipv4, nil
}

// validateIPv6 验证IPv6地址
func validateIPv6(ip string) (net.IP, error) {
	if ip == "" {
		return nil, nil
	}

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return nil, fmt.Errorf("invalid IPv6 address: %s", ip)
	}

	if parsed.To4() != nil {
		return nil, fmt.Errorf("not an IPv6 address: %s", ip)
	}

	return parsed.To16(), nil
}

// ToBase36 将IPv6和IPv4组合转换为36进制
func (c *IPConverter) ToBase36(ipv6, ipv4 string) (string, error) {
	// 验证IP地址
	ip6, err := validateIPv6(ipv6)
	if err != nil {
		return "", err
	}

	ip4, err := validateIPv4(ipv4)
	if err != nil {
		return "", err
	}

	if ip6 == nil && ip4 == nil {
		return "", fmt.Errorf("at least one IP address must be provided")
	}

	// 创建大整数存储组合结果
	result := new(big.Int)

	if ip6 != nil {
		// 存储IPv6
		ipv6Int := new(big.Int).SetBytes(ip6)
		// 验证IPv6不超过最大值
		if ipv6Int.Cmp(c.maxIPv6) > 0 {
			return "", fmt.Errorf("IPv6 value exceeds maximum")
		}
		result.Set(ipv6Int)
		// 左移32位，为IPv4预留空间
		result.Lsh(result, 32)
	}

	if ip4 != nil {
		// 验证IPv4不超过最大值
		ipv4Int := new(big.Int).SetBytes(ip4)
		if ipv4Int.Cmp(c.maxIPv4) > 0 {
			return "", fmt.Errorf("IPv4 value exceeds maximum")
		}
		// 将IPv4添加到结果中
		result.Or(result, ipv4Int)
	}

	// 转换为36进制
	base36 := result.Text(36)
	if c.upperCase {
		base36 = strings.ToUpper(base36)
	}

	return base36, nil
}

// FromBase36 从36进制解析回IPv6和IPv4
func (c *IPConverter) FromBase36(s string) (*IPPair, error) {
	// 解析36进制字符串
	ipInt := new(big.Int)
	_, ok := ipInt.SetString(s, 36)
	if !ok {
		return nil, fmt.Errorf("invalid base36 string: %s", s)
	}

	// 验证总值不超过最大允许值
	maxCombined := new(big.Int).Set(c.maxIPv6)
	maxCombined.Lsh(maxCombined, 32)
	maxCombined.Or(maxCombined, c.maxIPv4)

	if ipInt.Cmp(maxCombined) > 0 {
		return nil, fmt.Errorf("combined value exceeds maximum allowed")
	}

	// 创建返回结果
	pair := &IPPair{}

	// 提取IPv4部分（最后32位）
	ipv4Int := new(big.Int).And(ipInt, c.maxIPv4)
	if ipv4Int.BitLen() > 0 {
		// 验证IPv4值
		if ipv4Int.Cmp(c.maxIPv4) > 0 {
			return nil, fmt.Errorf("IPv4 portion exceeds maximum")
		}
		ipv4Bytes := ipv4Int.Bytes()
		pair.IPv4 = make(net.IP, 4)
		copy(pair.IPv4[4-len(ipv4Bytes):], ipv4Bytes)
	}

	// 提取IPv6部分（去除最后32位）
	ipv6Int := new(big.Int).Rsh(ipInt, 32)
	if ipv6Int.BitLen() > 0 {
		// 验证IPv6值
		if ipv6Int.Cmp(c.maxIPv6) > 0 {
			return nil, fmt.Errorf("IPv6 portion exceeds maximum")
		}
		ipv6Bytes := ipv6Int.Bytes()
		if len(ipv6Bytes) > 0 {
			pair.IPv6 = make(net.IP, 16)
			copy(pair.IPv6[16-len(ipv6Bytes):], ipv6Bytes)
		}
	}

	return pair, nil
}

func main() {
	converter := NewIPConverter(false)

	// 测试场景
	testCases := []struct {
		name string
		ipv6 string
		ipv4 string
	}{
		{
			name: "IPv6 + IPv4",
			ipv6: "2001:db8::1",
			ipv4: "192.168.1.1",
		},
		{
			name: "只有IPv4",
			ipv6: "",
			ipv4: "192.168.1.1",
		},
		{
			name: "只有IPv6",
			ipv6: "2001:db8::1",
			ipv4: "",
		},
		{
			name: "最大值测试",
			ipv6: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
			ipv4: "255.255.255.255",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试场景: %s\n", tc.name)
		fmt.Printf("原始 IPv6: %s\n", tc.ipv6)
		fmt.Printf("原始 IPv4: %s\n", tc.ipv4)

		// 转换为36进制
		base36, err := converter.ToBase36(tc.ipv6, tc.ipv4)
		if err != nil {
			fmt.Printf("转换错误: %v\n", err)
			continue
		}
		fmt.Printf("Base36: %s\n", base36)

		// 转换回IP对
		pair, err := converter.FromBase36(base36)
		if err != nil {
			fmt.Printf("反向转换错误: %v\n", err)
			continue
		}

		fmt.Printf("解析后 IPv6: %s\n", pair.IPv6)
		fmt.Printf("解析后 IPv4: %s\n", pair.IPv4)

		// 验证结果
		if tc.ipv6 != "" && pair.IPv6.String() != net.ParseIP(tc.ipv6).String() {
			fmt.Printf("警告: IPv6不匹配!\n")
		}
		if tc.ipv4 != "" && pair.IPv4.String() != net.ParseIP(tc.ipv4).String() {
			fmt.Printf("警告: IPv4不匹配!\n")
		}
		fmt.Printf("----------\n")
	}
}


```

检查IPv4
```go
// IPv4Checker IPv4地址检查器
type IPv4Checker struct {
    privateIPRanges []*net.IPNet
}

// NewIPv4Checker 创建新的IPv4检查器
func NewIPv4Checker() *IPv4Checker {
    checker := &IPv4Checker{}
    // 初始化内网IP范围
    privateRanges := []string{
        "10.0.0.0/8",     // RFC1918
        "172.16.0.0/12",  // RFC1918
        "192.168.0.0/16", // RFC1918
        "169.254.0.0/16", // RFC3927 Link-Local
        "127.0.0.0/8",    // RFC1122 Loopback
        "100.64.0.0/10",  // RFC6598 Shared Address Space
        "192.0.0.0/24",   // RFC6890
        "192.0.2.0/24",   // RFC5737 TEST-NET-1
        "198.18.0.0/15",  // RFC2544 Benchmark
        "198.51.100.0/24",// RFC5737 TEST-NET-2
        "203.0.113.0/24", // RFC5737 TEST-NET-3
        "224.0.0.0/4",    // RFC5771 Multicast
        "240.0.0.0/4",    // RFC1112 Reserved
    }
    
    for _, cidr := range privateRanges {
        _, network, _ := net.ParseCIDR(cidr)
        if network != nil {
            checker.privateIPRanges = append(checker.privateIPRanges, network)
        }
    }
    return checker
}

// IsPrivateIP 判断是否为内网IP
func (c *IPv4Checker) IsPrivateIP(ipStr string) bool {
    ip := net.ParseIP(ipStr)
    if ip == nil {
        return false
    }
    
    // 检查是否在内网范围内
    for _, privateRange := range c.privateIPRanges {
        if privateRange.Contains(ip) {
            return true
        }
    }
    return false
}
```

```
// 定义所有内网IPv4网段
	networks := []struct {
		CIDR string
		Type string
	}{
		// RFC 1918 私有网络
		{"10.0.0.0/8", "Private-A"},
		{"172.16.0.0/12", "Private-B"},
		{"192.168.0.0/16", "Private-C"},
		// 回环地址
		{"127.0.0.0/8", "Loopback"},
		// Link-Local
		{"169.254.0.0/16", "Link-Local"},
		// 组播地址
		{"224.0.0.0/4", "Multicast"},
		// 广播地址
		{"255.255.255.255/32", "Broadcast"},
		// 保留地址
		{"0.0.0.0/8", "Reserved"},
		{"100.64.0.0/10", "Shared-Address-Space"},
		{"192.0.0.0/24", "IETF-Protocol"},
		{"192.0.2.0/24", "TEST-NET-1"},
		{"198.18.0.0/15", "Benchmark"},
		{"198.51.100.0/24", "TEST-NET-2"},
		{"203.0.113.0/24", "TEST-NET-3"},
		{"240.0.0.0/4", "Reserved"},
	}


// 定义所有内网IPv6网段
	networks := []struct {
		CIDR string
		Type string
	}{
		// Unique Local Address (ULA)
		{"fc00::/7", "ULA"},
		// Link-Local
		{"fe80::/10", "Link-Local"},
		// Site-Local (已弃用，但仍需检查)
		{"fec0::/10", "Site-Local-Deprecated"},
		// Loopback
		{"::1/128", "Loopback"},
		// 组播地址
		{"ff00::/8", "Multicast"},
		// 本地组播
		{"ff01::/16", "Interface-Local-Multicast"},
		{"ff02::/16", "Link-Local-Multicast"},
		{"ff03::/16", "Realm-Local-Multicast"},
		{"ff04::/16", "Admin-Local-Multicast"},
		{"ff05::/16", "Site-Local-Multicast"},
		// 文档示例地址
		{"2001:db8::/32", "Documentation"},
		// 6to4地址
		{"2002::/16", "6to4"},
		// Teredo隧道
		{"2001::/32", "Teredo"},
		// 保留地址
		{"::/128", "Unspecified"},
		{"100::/64", "Discard-Only"},
		{"64:ff9b::/96", "IPv4-IPv6-Translation"},
	}	
```