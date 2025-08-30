package probe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LarkCard struct {
	MsgType string      `json:"msg_type"`
	Card    interface{} `json:"card"`
}

type NotifyInfo struct {
	DBConfig *DBConfig
	Success  bool
	ErrorMsg string
}

func NotifyFailure(notify *NotifyInfo, webhook string, actionURL string) (err error) {
	db := notify.DBConfig

	defer func() {
		if err != nil {
			fmt.Printf("Send notification failed: %s(%d) - %s\n", db.ClusterID, db.Port, err.Error())
		}
	}()

	content := fmt.Sprintf("**Cluster_ID:** %s\n**Region:** %s\n**Plan:** %s\n**Port:** %d\n**Error:** %s", db.ClusterID, db.Region, db.Plan, db.Port, notify.ErrorMsg)
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
						"content": "more infomation",
						"tag":     "lark_md",
					},
					"url":   actionURL,
					"type":  "default",
					"value": map[string]interface{}{},
				},
			},
		},
	}

	card := map[string]interface{}{
		"elements": elements,
		"header": map[string]interface{}{
			"title": map[string]interface{}{
				"content": "TiDB Cloud Probe Failure Alert",
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
