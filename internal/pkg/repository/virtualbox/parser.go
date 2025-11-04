package virtualbox

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/neatflowcv/vesta/internal/pkg/domain"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data []byte) []*domain.Instance {
	var ret []*domain.Instance

	name := ""
	uuid := ""
	status := domain.InstanceStatusUnknown

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			ret = append(ret, domain.NewInstance(uuid, name, status))
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
			status = mapVBStateToInstanceStatus(value)
		}
	}

	if name != "" {
		ret = append(ret, domain.NewInstance(uuid, name, status))
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

func mapVBStateToInstanceStatus(item string) domain.InstanceStatus {
	idx := strings.Index(item, "(")
	if idx != -1 {
		item = item[:idx]
	}

	item = strings.TrimSpace(item)
	switch item {
	case "running":
		return domain.InstanceStatusRunning
	case "powered off", "saved":
		return domain.InstanceStatusStopped
	case "stopping", "powering off", "aborted":
		return domain.InstanceStatusStopping
	case "starting", "powering on", "booting":
		return domain.InstanceStatusBooting
	default:
		return domain.InstanceStatusUnknown
	}
}
