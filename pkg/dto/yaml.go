package dto

func AddYamlFrontAndEndMatter(data []byte) []byte {
	outData := "---\n" + string(data) + "\n---"

	return []byte(outData)
}

func AddYamlFrontAndEndMatterText(data []byte, text string) []byte {
	outData := "---\n" + string(data) + "\n---\n" + text

	return []byte(outData)
}
