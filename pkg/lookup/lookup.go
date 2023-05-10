package lookup

import (
	"github.com/geoff-hill/dw/pkg/words"
)

func Keys() int {
	return len(words.WordMap())
}


func LookupByKey(key string) []string {
	return words.WordMap()[key]
}

