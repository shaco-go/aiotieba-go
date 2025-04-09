package core

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// 模拟 _sign 函数，使用 MD5 哈希计算签名
func _sign(data url.Values) string {
	var builder strings.Builder

	// 预分配缓冲区（根据参数数量估算）
	builder.Grow(len(data)*20 + 15) // 假设每个键值对约20字符，加15字符后缀

	// 获取排序后的键
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按排序后的键拼接参数
	for _, key := range keys {
		values := data[key]
		if len(values) == 0 {
			continue
		}

		builder.WriteString(key)
		builder.WriteByte('=')
		builder.WriteString(values[0]) // 默认取第一个值
	}

	// 添加签名后缀
	builder.WriteString("tiebaclient!!!")

	// 计算 MD5 哈希
	hash := md5.Sum([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}

// Sign 函数为参数添加签名
func Sign(data url.Values) url.Values {
	// 创建拷贝避免修改原始数据
	newData := make(url.Values, len(data)+1)
	for k, v := range data {
		newData[k] = v
	}

	newData.Set("sign", _sign(data))
	return newData
}
