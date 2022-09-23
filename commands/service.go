package commands

import (
	"fmt"
	"log"

	. "github.com/DavidZisky/kubegen/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Generate a service manifest",
		Long:  ``,
		Run:   service,
	}
)

func service(cmd *cobra.Command, args []string) {

	if targetport == 0 {
		targetport = port
	}

	d := Service{
		TypeMeta: TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: ObjectMeta{
			Name: name + "-service",
		},
		Spec: ServiceSpec{
			Selector: map[string]string{"app": name},
			Type:     ServiceType(atype),
			Ports: []ServicePort{
				ServicePort{
					Protocol:   "TCP",
					Port:       port,
					TargetPort: targetport,
				},
			},
		},
	}

	s, err := yaml.Marshal(&d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("---\n%v", string(s))
}

func includeServiceFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "name", "n", "name", "The name of the service")
	cmd.Flags().StringVarP(&atype, "type", "t", "ClusterIP", "The type of the service")
	cmd.Flags().Int32VarP(&port, "port", "p", 80, "Which port to expose on")
	cmd.Flags().Int32VarP(&targetport, "targetport", "d", 0, "Which port container listens on")
}

func init() {
	includeServiceFlags(serviceCmd)
}
