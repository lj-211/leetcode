package solutions

import (
	"reflect"
)

func simplifyPath(path string) string {
	parr := make([]string, 0)
	start := 1
	end := -1
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			end = i
		} else if i == len(path)-1 {
			end = len(path)
		}

		if end != -1 {
			// 处理
			p := path[start:end]
			if p == "." || p == "" {
				// do nothing
			} else if p == ".." {
				if len(parr) > 0 {
					parr = parr[0 : len(parr)-1]
				}
			} else {
				parr = append(parr, p)
			}
			start = end + 1
			end = -1
		}
	}

	ret := ""
	if len(parr) == 0 {
		ret = "/"
	} else {
		for i := len(parr) - 1; i >= 0; i-- {
			ret = "/" + parr[i] + ret
		}
	}

	return ret
}

func init() {
	desc := `
Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. For more information, see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.



Example 1:

Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
Example 4:

Input: "/a/./b/../../c/"
Output: "/c"
Example 5:

Input: "/a/../../b/../c//.//"
Output: "/c"
Example 6:

Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
	`
	sol := Solution{
		Title:  "Simplify Path",
		Desc:   desc,
		Method: reflect.ValueOf(simplifyPath),
		Tests:  make([]TestCase, 0),
	}
	a := TestCase{}
	a.Input = []interface{}{"/a//b////c/d//././/.."}
	a.Output = []interface{}{"/a/b/c"}
	sol.Tests = append(sol.Tests, a)

	SolutionMap["0071"] = sol
}
