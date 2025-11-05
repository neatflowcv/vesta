package virtualbox

import (
	"bufio"
	"bytes"
	"strings"
)

func parseVMs(data []byte) []*VM {
	var ret []*VM

	name := ""
	uuid := ""
	status := ""

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			ret = append(ret, NewVM(uuid, name, status))
			name = ""

			continue
		}

		key, value := split(line)
		value = strings.TrimSpace(value)

		switch key {
		case "Name":
			name = value
		case "UUID":
			uuid = value
		case "State":
			status = value
		}
	}

	if name != "" {
		ret = append(ret, NewVM(uuid, name, status))
	}

	return ret
}

func split(s string) (string, string) {
	const size = 2

	parts := strings.SplitN(s, ":", size)
	if len(parts) != size {
		return s, ""
	}

	return parts[0], parts[1]
}

func parseVM(data []byte) *VM {
	const size = 2

	name := ""
	uuid := ""
	status := ""

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.SplitN(line, "=", size)

		switch sp[0] {
		case "name":
			name = strings.Trim(sp[1], "\"")
		case "UUID":
			uuid = strings.Trim(sp[1], "\"")
		case "VMState":
			status = strings.Trim(sp[1], "\"")
		}
	}

	return NewVM(uuid, name, status)
}
