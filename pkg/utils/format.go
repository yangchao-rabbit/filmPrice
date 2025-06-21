package utils

import (
	"encoding/json"
	"filmPrice/internal/models"
	"time"
)

// StringValue 将String指针转为值
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntValue(s *int) int {
	if s == nil {
		return 0
	}

	return *s
}

func Int64Value(s *int64) int64 {
	if s == nil {
		return 0
	}

	return *s
}

func Uint64Value(s *uint64) uint64 {
	if s == nil {
		return 0
	}

	return *s
}

// ToString 将任何对象转为字符串
func ToString(v interface{}) string {
	marshal, _ := json.Marshal(v)
	return string(marshal)
}

// TimeValue 将时间戳转换为时间
func TimeValue(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func Map2CustomMap(m map[string]string) models.CustomMap {
	if m == nil {
		return nil
	}

	cm := make(models.CustomMap)
	for k, v := range m {
		cm[k] = v
	}

	return cm
}

func List2CustomList[T any](list []T) models.CustomList {
	if list == nil {
		return nil
	}

	cl := make(models.CustomList, 0, len(list))
	for _, v := range list {
		cl = append(cl, v)
	}

	return cl
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
