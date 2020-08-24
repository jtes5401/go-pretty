package pretty

const (
	grand = 62
	less = 60
)

func PrettyXML(xmlStr string) (xml string) {

	tabCount := 0
	size := len(xmlStr)

	for i,c := range xmlStr {

		if c == '<' && xmlStr[i+1] == '/'{
			// tag foot
			tabCount -= 1

		}else if c == '<' {
			// tag head
			tabCount += 1
			xml += tabStringCount(tabCount)
		}
		xml += string(c)


		// tag 結束
		if c == '>' && i+2 < size && xmlStr[i+2] == '/' && i+1 < size && xmlStr[i+1] == '<' {

		}else if c == '>' && i+1 < size && xmlStr[i+1] == '<' {
			xml += "\n"
		}
	}
	return
}

func tabStringCount(count int) (tab string) {
	for i := 0; i < count; i++ {
		tab += "\t"
	}
	return
}