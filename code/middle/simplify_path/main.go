package simplify_path

import "strings"

func simplifyPath(path string) string {
	result := make([]string, 0)

	for _, v := range strings.Split(path, "/") {
		if v == "" || v == "." {
			continue
		}

		if v == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, v)
		}
	}

	return "/" + strings.Join(result, "/")
}
