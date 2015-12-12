package jobs

import (
	"github.com/exherb/Dashing"
	"math/rand"
	"time"
)

type SampleJob struct{}

func random(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func (j SampleJob) Work(send chan *dashing.Event) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			send <- &dashing.Event{"sample", "widget", "synergy", map[string]interface{}{
				"value": random(0, 100),
			}}

			var doorStatus string
			if random(0, 100) > 50 {
				doorStatus = "打开"
			} else {
				doorStatus = "关闭"
			}
			send <- &dashing.Event{"sample", "widget", "door", map[string]interface{}{
				"text": doorStatus,
			}}

			send <- &dashing.Event{"sample", "widget", "elevator", map[string]interface{}{
				"value": random(1, 7),
			}}

			send <- &dashing.Event{"sample", "widget", "wifi", map[string]interface{}{
				"value": random(0, 300),
			}}

			now := time.Now()
			send <- &dashing.Event{"sample", "widget", "line", map[string]interface{}{
				"value": []int{random(0, 100), random(0, 100)},
				"label": now.Format("04:05"),
			}}

			send <- &dashing.Event{"sample", "widget", "segment", map[string]interface{}{
				"value": []int{random(0, 100), random(0, 100), random(0, 100), random(0, 100), random(0, 100)},
			}}

			send <- &dashing.Event{"map", "widget", "map", map[string]interface{}{
				"coordinate": []int{random(73*1000, 135*1000) / 1000.0, random(20*1000, 55*1000) / 1000.0},
				"value":      random(0, 50),
			}}
		}
	}
}

func init() {
	dashing.Register(SampleJob{})
}
