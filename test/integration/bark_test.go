package integration

import (
	"testing"

	"xzdbd.com/go-bark/bark"
)

func TestPush(t *testing.T) {
	c := bark.NewClient("https://api.day.app/push", "invalid")
	opts := &bark.Options{
		Archive: "1",
		Copy:    "123",
		Group:   "group1",
		Badge:   "1",
	}
	resp, err := c.Push("title", "body", opts)
	t.Log(resp, err)
	if err == nil {
		t.Fatalf("Expect error\n")
	}
	if resp == nil {
		t.Fatalf("Expect resp\n")
	}
	if resp.Code != 400 {
		t.Fatalf("Error code not equal to 400\n")
	}

}
