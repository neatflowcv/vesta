package virtualbox

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
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

func parseVM(data []byte) (*VM, error) {
	const size = 2

	name := ""
	uuid := ""
	status := ""

	cpu := uint64(0)
	ram := uint64(0)
	storage := uint64(0)
	var err error

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
		case "cpus":
			cpu, err = strconv.ParseUint(strings.Trim(sp[1], "\""), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse CPU: %w", err)
			}
		case "memory":
			ram, err = strconv.ParseUint(strings.Trim(sp[1], "\""), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse RAM: %w", err)
			}
			ram = ram * 1024 * 1024
		}
	}

	return NewVM(uuid, name, status, cpu, ram, storage), nil
}
