package cachedresult

import "time"

type TimedAnswer struct {
	Answer    interface{}
	ExpiresAt time.Time
}

type CachedResult struct {
	Answers map[interface{}]TimedAnswer
}
