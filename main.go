package main

import (
	"os"
	"path/filepath"

	demo "github.com/saschagrunert/demo"

	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()
	d.Add(fosdem22Demo(), "fosdem-22", "Fosdem 22 demo")
	d.Setup(setup)
	d.Run()
}

func setup(*cli.Context) error {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	os.RemoveAll(filepath.Join(userHomeDir, ".cache", "kubewarden"))

	return nil
}

func fosdem22Demo() *demo.Run {
	r := demo.NewRun(
		"Kubewarden FOSDEM 22 demo",
	)

	r.Step(demo.S(
		"List policies",
	), demo.S(
		"kwctl policies",
	))

	r.Step(demo.S(
		"Pull a policy",
	), demo.S(
		"kwctl pull registry://ghcr.io/kubewarden/policies/disallow-service-loadbalancer:v0.1.2",
	))

	r.Step(demo.S(
		"List policies",
	), demo.S(
		"kwctl policies",
	))

	r.Step(demo.S(
		"Inspect a policy",
	), demo.S(
		"kwctl inspect registry://ghcr.io/kubewarden/policies/disallow-service-loadbalancer:v0.1.2",
	))

	r.Step(demo.S(
		"Request that should be allowed",
	), demo.S(
		"jq -r .object.spec.type service-lb/service_clusterip.json",
	))

	r.Step(demo.S(
		"Policy accepts a request",
	), demo.S(
		"kwctl run --request-path service-lb/service_clusterip.json registry://ghcr.io/kubewarden/policies/disallow-service-loadbalancer:v0.1.2 | jq",
	))

	r.Step(demo.S(
		"Request that should be rejected",
	), demo.S(
		"jq -r .object.spec.type service-lb/service_loadbalancer.json",
	))

	r.Step(demo.S(
		"Policy rejects a request",
	), demo.S(
		"kwctl run --request-path service-lb/service_loadbalancer.json registry://ghcr.io/kubewarden/policies/disallow-service-loadbalancer:v0.1.2 | jq",
	))

	return r
}
