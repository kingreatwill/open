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