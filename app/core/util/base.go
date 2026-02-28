package util

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func ParseDate(value string) (time.Time, error) {
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("无法解析时间: %s", value)
}

func UIDWithContext(ctx context.Context) uint64 {
	if uid, ok := ctx.Value("userID").(string); ok {
		id, _ := strconv.ParseUint(uid, 10, 64)
		return id
	}
	return 0
}
