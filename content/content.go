package content

import (
	"bufio"
	"fmt"
	"github.com/t0of/qbang/db"
	"io"
	"os"
	"strconv"
	"strings"
)

type Content struct {
	Path string
}

func NewContent(Path string) *Content {
	return &Content{
		Path: Path,
	}
}

func (c Content) Do() {
	ReadFile(c.Path, func(s string) {
		fmt.Println(s)
	})
}
func ReadFile(filePath string, handle func(string)) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	for i := 0; i < 1000; i++ {
		data := " "
		for j := i * 10000; j < i*10000+10000; j++ {
			line, err := buf.ReadString('\n')
			sli := splitStr(line)
			if j < i*10000+9999 {
				id := strconv.Itoa(j)
				onedata := "(" + id + ", " + sli[0] + ", " + sli[1] + "), "
				data = data + onedata
			} else {
				id := strconv.Itoa(j)
				onedata := "(" + id + ", " + sli[0] + ", " + sli[1] + ")"
				data = data + onedata
			}
			if err == io.EOF {
				handle(line)
				return
			}
			handle(line)
		}
		err := db.InsertRows(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func splitStr(s string) []string {
	sli := strings.Split(s, "----")
	if len(sli) != 2 {
		sli = append(sli, "''")
	}
	return sli
}
