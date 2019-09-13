package main

import "bytes"

/*
	/. 无视 /.. 上一级
*/

func simplifyPath(path string) string {
	now := 0
	var res []string
	var resBuf bytes.Buffer
	for now < len(path) {
		bef := now
		for now < len(path) && path[now] != '/' {
			now++
		}
		if bef != now {
			if path[bef] == '.' {
				if bef < len(path)-1 && path[bef+1] == '.' {
					if len(res) != 0 {
						res = res[:len(res)-1]
					}
				}
				// ignore one dot
			} else {
				res = append(res, path[bef:now])
			}
		} else {
			now++
			continue
		}

		for now < len(path) && path[now] == '/' { // remove duplicate '/'
			now++
		}
	}

	if res == nil || len(res) == 0 {
		return "/"
	}

	for idx := range res {
		resBuf.WriteString("/" + res[idx])
	}

	return resBuf.String()
}
