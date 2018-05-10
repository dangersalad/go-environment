package environment // import "github.com/dangersalad/go-environment"

import (
	"fmt"
	"os"
	"strings"
)

// Options are the options being looked for in the environment
type Options map[string]string

// ReadOptions reads options from the environment. If the value of a
// key is an empty string and the key is not in the environment, an
// error is returned. Otherwise the supplied value is returned.
func ReadOptions(params Options) (Options, error) {

	missingKeys := []string{}
	for envVar := range params {
		if param := os.Getenv(envVar); param == "" {
			// if no default is supplied for the missing key, return an error
			if params[envVar] == "" {
				missingKeys = append(missingKeys, envVar)
			}
		} else {
			params[envVar] = param
		}
	}

	missingCount := len(missingKeys)
	if missingCount > 0 {
		if missingCount == 1 {
			return nil, fmt.Errorf("missing %s from environment", missingKeys[0])
		} else if missingCount == 2 {
			return nil, fmt.Errorf("missing %s and %s from environment", missingKeys[0], missingKeys[1])
		}
		return nil, fmt.Errorf("missing %s and %s from environment", strings.Join(missingKeys[:missingCount-1], ", "), missingKeys[missingCount-1:][0])
	}

	return params, nil

}
