package tt

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func maxArea(height []int) int {
	area, l := 0, len(height)-1
	i, j := 0, l
	for i < j {
		h := min(height[i], height[j])
		area = max(area, h*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Event struct {
	//ID        string            `josn:"_id"`
	EventName string                 `json:"eventName"`
	Timestamp string                 `json:"timestamp"`
	Time      int64                  `json:"time"`
	EventData map[string]interface{} `json:"eventData"`
	Info      struct {
		UserID string `json:"userId"`
		RoomId string `json:"roomId"`
		Role   string `json:"role"`
	} `json:"info"`
	IsSuccess bool `json:"isSuccess"`
}

func Sort(jsons ...string) string {
	objs := make([]Event, 0)
	for _, s := range jsons {
		arr := make([]Event, 0)
		if err := json.Unmarshal([]byte(s), &arr); err != nil {
			fmt.Printf("%v\n", err)
		}
		objs = append(objs, arr...)
	}
	for i := range objs {
		t, err := strconv.ParseInt(objs[i].Timestamp, 10, 64)
		fmt.Printf("time: %d\n", t)
		objs[i].Time = t
		if err != nil {
			fmt.Printf("%v\n", err)
			//return ""
		}
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].Time < objs[j].Time
	})
	rs, _ := json.Marshal(objs)
	return string(rs)
}

func Json(str string) Event {
	var rs Event
	if err := json.Unmarshal([]byte(str), &rs); err != nil {
		fmt.Printf("%v\n", err)
	}
	return rs
}
