package dto

func addYamlFrontAndEndMatter(data []byte) []byte {
	outData := "---\n" + string(data) + "\n---"

	return []byte(outData)
}

func addYamlFrontAndEndMatterText(data []byte, text string) []byte {
	outData := "---\n" + string(data) + "\n---\n" + text

	return []byte(outData)
}
