package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestFindSubmatch(t *testing.T) {
	var re = regexp.MustCompile(`foo(.*?)`)
	result := re.FindAllSubmatch([]byte("cood"), -1)

	fmt.Println(len(result))
}
