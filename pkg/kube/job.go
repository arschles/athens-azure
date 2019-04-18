package kube

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ericchiang/k8s"
	batchv1 "github.com/ericchiang/k8s/apis/batch/v1"
)

type Job struct {
	core *batchv1.Job
	fmt.Stringer
	Resource
}

func NewJob(name, ns string, containers ContainerList) *Job {
	ret := &Job{
		core: &batchv1.Job{
			Metadata: objectMeta(name, ns),
			Spec: &batchv1.JobSpec{
				Template: podTemplateSpec(name, ns, containers),
			},
		},
	}
	ret.core.Spec.Template.Spec.RestartPolicy = k8s.String("OnFailure")
	return ret
}

func (j *Job) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, j.core)
}

func (j *Job) Name() string {
	return *j.core.Metadata.Name
}

func (j *Job) Get(ctx context.Context, cl *k8s.Client, name, ns string) error {
	return cl.Get(ctx, ns, name, j.core)
}

func (j *Job) GetImage(idx int) (string, error) {
	con := containerFromPodTemplateSpec(j.core.Spec.Template, idx)
	if con == nil {
		return "", fmt.Errorf("container %d doesn't exist", idx)
	}
	return con.GetImage(), nil
}

func (j *Job) Delete(ctx context.Context, cl *k8s.Client) error {
	return cl.Delete(ctx, j.core)
}

func (j *Job) Namespace() *Namespace {
	return NewNamespace(*j.core.Metadata.Namespace)
}

func (j *Job) DeletedCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ch := make(chan error)
	close(ch)
	return ch
}

func (j *Job) Type() string {
	return "Job"
}

func (j *Job) ReadyCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ch := make(chan error)
	close(ch)
	return ch
}

func (j *Job) String() string {
	b, err := json.Marshal(j.core)
	if err != nil {
		return fmt.Sprintf("error marshaling Job %s", j.Name())
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, "", "    "); err != nil {
		return fmt.Sprintf("error indenting JSON for job %s", j.Name())
	}
	return string(buf.Bytes())
}
