package parser

import (
	"function/entity"
	"github.com/sirikothe/gotextfsm"
	"strings"
)

const (
	IFCONFIGCMD = "ifconfig"
)

func ParseIfConfigOutput(output string) ([]entity.Interface, error) {
	networkInfo := make([]entity.Interface, 0)

	template := `Value INTERFACE (\S+)
Value MTU (\d+)
Value RX_PACKETS (\d+)
Value RX_BYTES (\d+)
Value TX_PACKETS (\d+)
Value TX_BYTES (\d+)

Start
  ^${INTERFACE}: flags=\S+  mtu ${MTU}
  ^RX packets ${RX_PACKETS}  bytes ${RX_BYTES} \(\d+\.?\d* \w+\)
  ^TX packets ${TX_PACKETS}  bytes ${TX_BYTES} \(\d+\.?\d* \w+\) -> Record
`
	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		return nil, err
	}
	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(output, fsm, true)
	if err != nil {
		return nil, err
	}

	for _, record := range parser.Dict {
		inface := entity.NewInterface(
			record["INTERFACE"].(string),
			record["MTU"].(string),
			record["RX_PACKETS"].(string),
			record["RX_BYTES"].(string),
			record["TX_PACKETS"].(string),
			record["TX_BYTES"].(string),
		)
		networkInfo = append(networkInfo, inface)
	}

	return networkInfo, nil
}

func IsIfConfig(userInput string) bool {
	return strings.Contains(userInput, IFCONFIGCMD)
}
