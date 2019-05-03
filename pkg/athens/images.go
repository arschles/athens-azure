package athens

import (
	"github.com/spf13/pflag"
)

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

func getImages(
	flags *pflag.FlagSet,
	markRequired func(string) error,
) (*images, error) {
	imgs := images{}
	flags.StringVar(&imgs.athens, "athens", "", "The name of the Athens Image")
	if err := markRequired("athens"); err != nil {
		return nil, err
	}
	flags.StringVar(&imgs.crathens, "crathens", "", "The name of the crathens image")
	if err := markRequired("crathens"); err != nil {
		return nil, err
	}
	flags.StringVar(&imgs.lathens, "lathens", "", "The name of the crathens image")
	if err := markRequired("lathens"); err != nil {
		return nil, err
	}

	return &imgs, nil
}
