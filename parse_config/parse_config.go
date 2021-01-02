package parse_config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"taskmaster/job_control"
)

func ParseConfig(filePath string) (jobs job_control.Jobs, err error) {
	data, err := ioutil.ReadFile(filePath)
	err = yaml.Unmarshal(data, &jobs)
	return
}
