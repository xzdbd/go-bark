package integration

import (
	"testing"

	"xzdbd.com/go-bark/bark"
)

func TestPush(t *testing.T) {
	c := bark.NewClient("https://api.day.app", "invalid")
	opts := &bark.PushOptions{
		Archive:     true,
		Copy:        "123",
		AutoCopy:    true,
		DirectedURL: "https://g.com",
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
