package conf

// Athens returns the configuration for a full Athens deployment
//
// TODO: build a decoder from HCL and put this into an HCL file!
func Athens() (*Root, error) {
	return &Root{
		Name: "athens",
		Webs: []Web{athensWeb(), lathensWeb()},
		Jobs: []Job{crathensJob()},
	}, nil
}

func athensWeb() Web {
	return Web{
		// NOTE: no domain because lathens is a frontend cache
		Replicas: 3,
		Name:     "athens",
		Image:    "gomods/athens:v0.3.1",
		Port:     8080,
	}
}

func lathensWeb() Web {
	return Web{
		Domain:   "athens.azurefd.net",
		Replicas: 3,
		Name:     "lathens",
		Image:    "quay.io/arschles/lathens:canary",
		Port:     8080,
	}
}

func crathensJob() Job {
	return Job{
		Name:  "crathens",
		Image: "quay.io/arschles/crathens:canary",
	}
}
