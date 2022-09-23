package commands

import (
	"fmt"
	"log"

	. "github.com/DavidZisky/kubegen/types"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	ingressCmd = &cobra.Command{
		Use:   "ingress",
		Short: "Generate a ingress manifest",
		Long:  ``,
		Run:   ingress,
	}
)

func ingress(cmd *cobra.Command, args []string) {

	if servicename == "servicename" {
		servicename = name + "-service"
	}

	d := Ingress{
		TypeMeta: TypeMeta{
			APIVersion: "networking.k8s.io/v1",
			Kind:       "Ingress",
		},
		ObjectMeta: ObjectMeta{
			Name: name + "-ingress",
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/",
			},
		},
		Spec: IngressSpec{
			/*
			TLS: []IngressTLS{
				IngressTLS{
					Hosts:      []string{"foo.bar.com"},
					SecretName: name + "-tls-secret",
				},
			},
			*/
			IngressClass: ingressclass,
			Rules: []IngressRule{
				IngressRule{
					//Host: "foo.bar.com",
					IngressRuleValue: IngressRuleValue{
						HTTP: &HTTPIngressRuleValue{
							Paths: []HTTPIngressPath{
								HTTPIngressPath{
									Path: "/",
									PathType: "Prefix",
									Backend: IngressBackend{
											Service: BackendService{
												ServiceName: servicename,
												Port: IngressServicePort{
													Number: port,
												},
											},
									},
								},
							},
						},
					},
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

func includeIngressFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&name, "name", "n", "name", "The name of the ingress")
	cmd.Flags().StringVarP(&ingressclass, "ingressclass", "c", "", "IngressClassName")
	cmd.Flags().StringVarP(&servicename, "servicename", "s", "servicename", "The name of the destination service")
	cmd.Flags().Int32VarP(&port, "port", "p", 80, "The main container port")
}

func init() {
	includeIngressFlags(ingressCmd)
}
