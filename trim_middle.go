package xstrings

import (
	"strings"
)

const (
	dots = "..."
)

func TrimMiddleWithPlacehodler(origin string, totalLen int, placehodler string) string {
	if len(origin) <= totalLen {
		return origin
	}

	if len(placehodler) >= totalLen {
		return placehodler
	}

	var (
		rightLen = (totalLen - len(placehodler)) / 2
		leftLen  = rightLen + (totalLen-len(placehodler))%2
		builder  strings.Builder
	)

	builder.Grow(totalLen)

	builder.WriteString(origin[:leftLen])
	builder.WriteString(placehodler)
	builder.WriteString(origin[len(origin)-rightLen:])

	return builder.String()
}

func TrimMiddle(s string, totalLen int) string {
	return TrimMiddleWithPlacehodler(s, totalLen, dots)
}
