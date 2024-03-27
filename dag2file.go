package merkledag

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Hash2File 从 KVStore 中读取与给定哈希相关联的数据，并根据路径返回对应的文件内容
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 从 KVStore 中获取与哈希相关联的数据
	data, err := store.Get(hash)
	if err != nil {
		fmt.Println("Error retrieving data from KVStore:", err)
		return nil
	}

	// 解码数据，假设数据是 JSON 格式
	var treeData map[string]interface{}
	err = json.Unmarshal(data, &treeData)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return nil
	}

	// 根据路径递归查找文件内容
	parts := strings.Split(path, "/")
	currentData := treeData
	for _, part := range parts {
		if part == "" {
			continue
		}
		if node, ok := currentData[part].(map[string]interface{}); ok {
			currentData = node
		} else {
			return []byte(fmt.Sprintf("%v", currentData[part]))
		}
	}

	return nil
}
