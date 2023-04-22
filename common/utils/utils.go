package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
)

// UnmarshalYamlToMap
// desc: 反序列化 yaml 为 map[interface{}]interface{}
func UnmarshalYamlToMap(yamlStr string) (map[interface{}]interface{}, error) {
	var m map[interface{}]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &m); err != nil {
		return nil, err
	}
	return m, nil
}

// MergeMap
// desc: 合并源 map 到目标 map
func MergeMap(dest, src map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(dest))
	for k, v := range dest {
		out[k] = v
	}
	for k, v := range src {
		value := v
		if av, ok := out[k]; ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
				if av, ok := av.(map[interface{}]interface{}); ok {
					out[k] = MergeMap(av, v)
				} else {
					out[k] = v
				}
			} else {
				out[k] = value
			}
		} else {
			out[k] = v
		}
	}
	return out
}

// MarshalObjectToYamlString
// desc: 序列化对象为 yaml 字符串
func MarshalObjectToYamlString(obj interface{}) ([]byte, error) {
	bs, err := yaml.Marshal(obj)
	if err != nil {
		return []byte{}, fmt.Errorf("marshal obj [%#v] faild, err: %v", obj, err)
	}
	return bs, nil
}

func MD5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
