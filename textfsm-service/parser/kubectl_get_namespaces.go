package parser

import (
	"function/entity"
	"github.com/sirikothe/gotextfsm"
	"strings"
)

const (
	KUBECTLGETNAMESPACECMD = "kubectl get namespace"
	KUBECTLGETNSCMD        = "kubectl get ns"
)

func ParseKubectlGetNsOutput(output string) ([]entity.NameSpace, error) {
	namespaceInfo := make([]entity.NameSpace, 0)

	template := `Value NAME (\S+)
Value STATUS (\S+)
Value AGE (\S+)

Start
  ^${NAME}\s+${STATUS}\s+${AGE} -> Record`
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
		if record["NAME"] == "NAME" {
			continue
		}

		namespace := entity.NameSpace{
			Name:   record["NAME"].(string),
			Status: record["STATUS"].(string),
			Age:    record["AGE"].(string),
		}
		namespaceInfo = append(namespaceInfo, namespace)
	}
	return namespaceInfo, nil
}

func IsKubectlGetNamespaces(userInput string) bool {
	return strings.Contains(userInput, KUBECTLGETNAMESPACECMD) || strings.Contains(userInput, KUBECTLGETNSCMD)
}
