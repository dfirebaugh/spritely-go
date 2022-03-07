package idempotency

import (
	"crypto/sha256"
	"fmt"
)

type IdempotencyManager struct {
	queue []string
}

func (i IdempotencyManager) String() string {
	var str string
	for index, v := range i.queue {
		str = str + "\n" + fmt.Sprintf("%d - ", index) + v
	}
	return str + "\n"
}

func (i IdempotencyManager) Hash(values string) string {
	h := sha256.New()
	h.Write([]byte(values))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (i *IdempotencyManager) LimitSize() {
	if len(i.queue) <= 20 {
		return
	}

	i.queue = i.queue[1:20]
}

// not true idempotency, but we check if we've
//   seen these values in the last 20 requests
func (i *IdempotencyManager) AlreadyProcessed(values string) bool {
	hash := i.Hash(values)

	// this could be more efficient
	for _, token := range i.queue {
		if token == hash {
			// println(i.String())
			return true
		}
	}
	i.LimitSize()
	i.queue = append(i.queue, hash)
	return false
}
