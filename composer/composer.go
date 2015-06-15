package composer

import (
	"bytes"
	"fmt"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

type Composer interface {
	Run(map[string]string) error
}

type ExecComposer struct {
}

func NewExecComposer(dockerHost string) *ExecComposer {
	return &ExecComposer{}
}

func (c *ExecComposer) DrainRequests(args []string) error {
	return nil
}

func (c *ExecComposer) Run(flags map[string]string) error {
	// TODO: implement --file functionality

	composeArgs := []string{"up", "-d"}
	for key, value := range flags {
		if key == "file" {
			composeArgs = append([]string{"-f", value}, composeArgs...)
		}
	}
	cmd := exec.Command("docker-compose", composeArgs...)
	// TODO: this port must match the proxy port, see related TODO in cluster/proxy.go
	cmd.Env = []string{"DOCKER_HOST=localhost:3000"}
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()
	fmt.Println(cmd.Stderr)
	fmt.Println(cmd.Stdout)

	log.Info("docker-compose stdout: ", out.String())
	if err != nil {
		log.Info("docker-compose stderr: ", errOut.String())
		log.Fatal(err)
	}

	return err
}