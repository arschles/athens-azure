package kube

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/stringer"
	"github.com/ericchiang/k8s"
	batchv1 "github.com/ericchiang/k8s/apis/batch/v1"
)

// Job is a convenience wrapper around Kubernetes jobs
type Job struct {
	core *batchv1.Job
	fmt.Stringer
	// Note: Jobs implement Resource, but don't embed it here
	// because the compiler will not give errors if you don't implement
	// a method in Resource when you put it into a profile
}

// NewJob creates a new Job with sensible defaults
func NewJob(name, ns string, containers ContainerList) *Job {
	ret := &Job{
		core: &batchv1.Job{
			Metadata: objectMeta(name, ns),
			Spec: &batchv1.JobSpec{
				Template: podTemplateSpec(emptyMap(), containers),
			},
		},
	}
	ret.core.Spec.Template.Spec.RestartPolicy = k8s.String("OnFailure")
	return ret
}

// Install implements Installer
func (j *Job) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, j.core)
}

// Update implements Updater
func (j *Job) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, j.core)
}

// Name implements Namer
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

// Namespace is the implementation of Namespacer
func (j *Job) Namespace() *Namespace {
	return NewNamespace(*j.core.Metadata.Namespace)
}

// DeletedCh implements DeletedWatcher
func (j *Job) DeletedCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ch := make(chan error)
	close(ch)
	return ch
}

// Type implements Typer
func (j *Job) Type() string {
	return "Job"
}

// ReadyCh implements ReadyWatcher
func (j *Job) ReadyCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ch := make(chan error)
	close(ch)
	return ch
}

// String implements Stringer
func (j *Job) String() string {
	return stringer.ToJSON(j.core, j.Type())
}
