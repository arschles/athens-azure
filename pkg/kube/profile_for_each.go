package kube

func forEachResource(
	res []resources.Resource,
	strat ErrorStrategy,
	fn func(resources.Resource) error,
) error {
	return forEachResourceIdx(res, strat, func(_ int, r resources.Resource) error {
		return fn(r)
	})
}

func forEachResourceIdx(
	resources []resources.Resource,
	strat ErrorStrategy,
	fn func(int, resources.Resource) error,
) error {
	errs := []error{}
	for i, prof := range resources {
		if err := fn(i, prof); err != nil {
			// TODO: error strategy
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.WithStack(errlist.Error(errs))
	}
	return nil
}

func forEachResourceReverse(
	resources []resources.Resource,
	strat ErrorStrategy,
	fn func(resources.Resource) error,
) error {
	errs := []error{}
	for i := len(resources) - 1; i >= 0; i-- {
		res := resources[i]
		if err := fn(res); err != nil {
			// TODO: error strategy
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.WithStack(errlist.Error(errs))
	}
	return nil
}
