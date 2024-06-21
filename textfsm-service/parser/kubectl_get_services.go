package parser

import (
	"function/entity"
	"github.com/sirikothe/gotextfsm"
	"strings"
)

const (
	KUBECTLGETSERVICECMD = "kubectl get service"
	KUBECTLGETSVCCMD     = "kubectl get svc"
)

func ParseKubectlGetSvcOutput(output string) ([]entity.Service, error) {
	services := make([]entity.Service, 0)

	template := `Value NAME (\S+)
Value TYPE (\S+)
Value CLUSTER_IP (\S+)
Value EXTERNAL_IP (\S+)
Value PORTS (\S+)
Value AGE (\S+)

Start
 ^${NAME}\s+${TYPE}\s+${CLUSTER_IP}\s+${EXTERNAL_IP}\s+${PORTS}\s+${AGE} -> Record`
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
		service := entity.Service{
			Name:       record["NAME"].(string),
			Type:       record["TYPE"].(string),
			ClusterIP:  record["CLUSTER_IP"].(string),
			ExternalIP: record["EXTERNAL_IP"].(string),
			Ports:      record["PORTS"].(string),
			Age:        record["AGE"].(string),
		}
		services = append(services, service)
	}
	return services, nil
}

func IsKubectlGetServices(userInput string) bool {
	return strings.Contains(userInput, KUBECTLGETSERVICECMD) || strings.Contains(userInput, KUBECTLGETSVCCMD)
}
