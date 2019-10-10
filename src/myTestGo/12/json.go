package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func getJsonData() []byte {
	// 初始化数据
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 10180,
		},

		Battery: Battery{
			2910,
		},

		HasTouchID: true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {
	jsonData := getJsonData()

	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)

	fmt.Printf("%+v\n", screenAndTouch)
  
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &batteryAndTouch)

	fmt.Printf("%+v\n", batteryAndTouch)
}
