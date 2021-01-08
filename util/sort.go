package util

import "sort"

type User struct {
	Name string
}

func sortChinesebyUTF8(followingList []User) []User {
	sort.Slice(followingList, func(i, j int) bool {
		return followingList[i].Name < followingList[j].Name
	})
	return followingList
}
