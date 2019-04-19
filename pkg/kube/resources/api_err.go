package resources

import (
	"github.com/ericchiang/k8s"
)

func errToAPIErr(e error) *k8s.APIError {
	apiErrPtr, ok := e.(*k8s.APIError)
	if ok {
		return apiErrPtr
	}
	return nil

}
