package probe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LarkCard struct {
	MsgType string      `json:"msg_type"`
	Card    interface{} `json:"card"`
}

type NotifyInfo struct {
	DBConfig  *DBConfig
	Success   bool
	LatencyMs int64
	ErrorMsg  string
}

func NotifyFailure(notify *NotifyInfo, webhook string, actionURL string) (err error) {
	db := notify.DBConfig

	defer func() {
		if err != nil {
			fmt.Printf("Send notification failed: %s(%d) - %s\n", db.ClusterID, db.Port, err.Error())
		}
	}()

	contentItems := []struct {
		key   string
		value string
	}{
		{"Cluster_ID", fmt.Sprintf("%s(%s)", db.ClusterID, strings.Split(db.User, ".")[0])},
		{"Region", db.Region},
		{"Plan", db.Plan},
		{"Port", fmt.Sprintf("%d", db.Port)},
		{"Pool", db.TiDBPool},
		{"Error", notify.ErrorMsg},
	}
	var contentBuilder bytes.Buffer
	for _, item := range contentItems {
		fmt.Fprintf(&contentBuilder, "**%s:** %s\n", item.key, item.value)
	}
	content := contentBuilder.String()
	elements := []interface{}{
		map[string]interface{}{
			"tag": "div",
			"text": map[string]interface{}{
				"content": content,
				"tag":     "lark_md",
			},
		},
		map[string]interface{}{
			"tag": "action",
			"actions": []interface{}{
				map[string]interface{}{
					"tag": "button",
					"text": map[string]interface{}{
						"content": "GitHub Action URL",
						"tag":     "lark_md",
					},
					"url":   actionURL,
					"type":  "primary",
					"value": map[string]interface{}{},
				},
				map[string]interface{}{
					"tag": "button",
					"text": map[string]interface{}{
						"content": "Dashboard",
						"tag":     "lark_md",
					},
					"url":   "https://tidbcloud-connection-probe.netlify.app/",
					"type":  "default",
					"value": map[string]interface{}{},
				},
			},
		},
	}

	card := map[string]interface{}{
		"elements": elements,
		// "card_link": map[string]interface{}{
		// 	"url": "https://tidbcloud-connection-probe.netlify.app/",
		// },
		"header": map[string]interface{}{
			"title": map[string]interface{}{
				"content": "TiDB Cloud Probe Alert",
				"tag":     "plain_text",
			},
			"template": "red",
		},
	}

	payload := LarkCard{
		MsgType: "interactive",
		Card:    card,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
