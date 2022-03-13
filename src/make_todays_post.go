package main

import (
	"io"
	"os"
	"strconv"
	"time"
)

func make_filepath(currentTimeStr string, i_int int) string {
	var res string
	i := strconv.Itoa(i_int)
	if i == "0" {
		res = "_posts/" + currentTimeStr + "-post.md"
	} else {
		res = "_posts/" + currentTimeStr + "-post" + i + ".md"
	}
	return res
}

func main() {
	currentTime := time.Now()

	timeStr := currentTime.Format("2006-01-02")

make_file:
	for i := 0; ; i++ {
		_, err := os.Open(make_filepath(timeStr, i))
		if err != nil {
			dstFile, _ := os.Create(make_filepath(timeStr, i))
			srcFile, _ := os.Open("post_template/post_template.md")
			_, _ = io.Copy(dstFile, srcFile)
			break make_file
		}
	}
}
