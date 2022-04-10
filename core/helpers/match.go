package helpers

import (
	"strconv"
	"strings"

	"github.com/riteshRcH/go-edge-device-lib/core/protocol"
)

// MultistreamSemverMatcher returns a matcher function for a given base protocol.
// The matcher function will return a boolean indicating whether a protocol ID
// matches the base protocol. A given protocol ID matches the base protocol if
// the IDs are the same and if the semantic version of the base protocol is the
// same or higher than that of the protocol ID provided.
// TODO
func MultistreamSemverMatcher(base protocol.ID) (func(string) bool, error) {
	parts := strings.Split(string(base), "/")

	// semver example
	// 1.2.3-beta.1+meta
	// 1 == major
	// 2 == minor
	// 3 == patch
	// beta.1 == pre-release
	// meta == build metadata

	semver := strings.Split(parts[len(parts)-1], ".")

	return func(check string) bool {
		chparts := strings.Split(check, "/")
		if len(chparts) != len(parts) {
			return false
		}

		for i, v := range chparts[:len(chparts)-1] {
			if parts[i] != v {
				return false
			}
		}

		checkAgainstSemver := strings.Split(chparts[len(chparts)-1], ".")

		semverMinor, err := strconv.Atoi(semver[1])
		if err != nil {
			return false
		}
		checkAgainstSemverMinor, err := strconv.Atoi(checkAgainstSemver[1])
		if err != nil {
			return false
		}

		return semver[0] == checkAgainstSemver[0] && semverMinor >= checkAgainstSemverMinor
	}, nil
}
