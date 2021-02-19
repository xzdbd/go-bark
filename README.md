# go-bark
go-bark is a Go client library for [Bark](https://github.com/Finb/Bark)

## Usage

```go
c := bark.NewClient("https://api.day.app", "key")
opts := &bark.PushOptions{
	Archive:     true,
	Copy:        "123",
	AutoCopy:    true,
	DirectedURL: "https://g.com",
}
resp, err := c.Push("title", "body", opts)
```
