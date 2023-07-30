package run

import (
	"encoding/json"
	"log"
	"os/exec"
)

type Output struct {
        Schema []Schema `json:""`
	Job    string `json:""`
}

type Schema struct {
	Results map[string]float64 `json:"results"`
	Labels map[string]string `json:"labels"`
}

func (o *Output) RunJob(p *Params) {
	if p.UseWg {
		defer p.Wg.Done()
	}
	o.RunExec(p.Path)
}

func (o *Output) RunExec(path *string) {

	out, err := exec.Command(*path).CombinedOutput()
	if err != nil {
		log.Fatalf("Exception from command execution",err)
	}

	err = json.Unmarshal(out, &o.Schema)
	if err != nil {
		log.Fatalf("Fatal exception then reading json",err)
	}

}
