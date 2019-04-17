package kube

import (
	"context"

	"github.com/ericchiang/k8s"
	batchv1 "github.com/ericchiang/k8s/apis/batch/v1"
)

type Job struct {
	core *batchv1.Job
	Installer
}

func NewJob(name, ns string, containers ContainerList) *Job {
	return &Job{
		core: &batchv1.Job{
			Metadata: objectMeta(name, ns),
			Spec: &batchv1.JobSpec{
				Template: podTemplateSpec(name, ns, containers),
			},
		},
	}
}

func (j *Job) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, j.core)
}
