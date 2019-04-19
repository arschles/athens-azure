package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/ericchiang/k8s"
	"github.com/souz9/errlist"
)

// ProfileComposer is a Profile implementation that composes multiple profiles
// together and manages them all at once. It's useful for installing more
// complex applications (like microservices etc...) in one go
//
// Use NewProfileComposer to create a new one of these
type ProfileComposer struct {
	profiles []Profile
}

// NewProfileComposer creates a new ProfileComposer from a list of profiles
func NewProfileComposer(profiles []Profile) *ProfileComposer {
	return &ProfileComposer{
		profiles: profiles,
	}
}

// String implements fmt.Stringer
func (p *ProfileComposer) String() string {
	ret := make([]string, len(p.profiles))
	for i, pr := range p.profiles {
		ret[i] = fmt.Sprintf("-----\nProfile #%d\n%s", i, pr.String())
	}
	return strings.Join(ret, "\n")
}

// Setup implements Profile
func (p *ProfileComposer) Setup(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	errs := []error{}
	for _, pr := range p.profiles {
		if err := pr.Update(ctx, cl, strat); err != nil {
			// TODO: error strategy
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}
