package blog_list

import (
	"testing"
)

// const BLOG = "/home/critter/programming/CrispinStichart.github.io/_posts/"

func testCheck(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// func TestGetPostsFromDir(t *testing.T) {
// 	for _, entry := range getPostsFromDir(BLOG) {
// 		fmt.Println(entry)
// 	}
// }
