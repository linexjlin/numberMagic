// noSegs
package main

import (
	"fmt"
	"strings"
)

type SegInfo struct {
	card, province, city string
}

type Segs struct {
	segsmap map[string]SegInfo
}

func (s *Segs) importData(data string) {
	for i, v := range strings.Split(data, "\n") {
		f := strings.Split(v, ",")
		if len(f) == 4 {
			s.segsmap[f[0]] = SegInfo{f[1], f[2], f[3]}
			//fmt.Println(f,"import")
		} else {
			fmt.Println("wrong format:", i+1, f)
		}
	}
}

func (s *Segs) Init() {
	s.segsmap = make(map[string]SegInfo)
}
