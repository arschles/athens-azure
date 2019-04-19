package athens

import "github.com/arschles/athens-azure/pkg/env"

const (
	athensImgEnv   = "ATHENS_IMAGE"
	crathensImgEnv = "CRATHENS_IMAGE"
	lathensImgEnv  = "LATHENS_IMAGE"
)

type images struct {
	athens   string
	crathens string
	lathens  string
}

func getImages(args []string) (*images, error) {
	ret := &images{}
	ath, err := env.CheckOrArg(athensImgEnv, args, 0)
	if err != nil {
		return nil, err
	}
	ret.athens = ath

	crath, err := env.CheckOrArg(crathensImgEnv, args, 1)
	if err != nil {
		return nil, err
	}
	ret.crathens = crath

	lath, err := env.CheckOrArg(lathensImgEnv, args, 2)
	if err != nil {
		return nil, err
	}
	ret.lathens = lath

	return ret, nil
}
