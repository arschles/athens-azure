package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/arschles/athens-azure/pkg/kube/resources"
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
	if err := forEachProfileIdx(
		p.profiles,
		ErrorStrategyStop,
		func(i int, pr Profile) error {
			ret = append(ret, fmt.Sprintf("-----\nProfile #%d\n%s", i, pr.String()))
			return nil
		}); err != nil {

		return ""
	}
	return strings.Join(ret, "\n")
}

// Setup implements Profile
func (p *ProfileComposer) Setup(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	if err := forEachProfile(p.profiles, strat, func(pr Profile) error {
		return pr.Update(ctx, cl, strat)
	}); err != nil {
		return err
	}
	return nil
}

// AllResources implements Profile. Since profiles themselves have resource
// lists, this function flattens each of the lists into one list, in the
// order of each list, and then in the order of the profiles in p
func (p *ProfileComposer) AllResources() []resources.Resource {
	ret := []resources.Resource{}
	if err := forEachProfile(p.profiles, ErrorStrategyStop, func(pr Profile) error {
		resources := pr.AllResources()
		ret = append(ret, resources...)
		return nil
	}); err != nil {
		// there won't be an error, because we always return nil in the closure
		// passed to forEachProfile
		return nil
	}
	return ret
}

// Install implements Profile
func (p *ProfileComposer) Install(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	if err := forEachProfile(p.profiles, strat, func(pr Profile) error {
		return pr.Install(ctx, cl, strat)
	}); err != nil {
		return err
	}
	return nil
}

// Uninstall implements Profile
func (p *ProfileComposer) Uninstall(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	if err := forEachProfile(p.profiles, strat, func(pr Profile) error {
		return pr.Uninstall(ctx, cl, strat)
	}); err != nil {
		return err
	}
	return nil
}

func forEachProfile(
	profs []Profile,
	strat ErrorStrategy,
	fn func(p Profile) error,
) error {
	return forEachProfileIdx(profs, strat, func(_ int, pr Profile) error {
		return fn(pr)
	})
}

func forEachProfileIdx(
	profs []Profile,
	strat ErrorStrategy,
	fn func(i int, p Profile) error,
) error {
	errs := []error{}
	for i, prof := range profs {
		if err := fn(i, prof); err != nil {
			// TODO: error strategy
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}
