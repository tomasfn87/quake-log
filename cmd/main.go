package main

import quakeLog "code.repo/quakeLog/src"

func main() {

	quakeLogFilePath := "./Quake.txt"
	file := quakeLog.QuakeLogFile{Path: quakeLogFilePath}

	file.OpenQuakeLog()

	/*
		x := "a"
		y := "b"
		z := "c"
		list := []string{x, y, z}

				fmt.Println(arrayContains(list, "a"))
				fmt.Println(arrayContains(list, "b"))
				fmt.Println(arrayContains(list, "c"))
				fmt.Println(arrayContains(list, "d"))

			list = updateOldNames(list, "d")
			fmt.Println(list)
	*/
}

/*
func arrayContains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
*/

func updateOldNames(oldNames []string, currentName string) []string {
	return append(oldNames, currentName)
}
