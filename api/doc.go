// Package temporalapimocks exists so go.temporal.io/api stays a real dependency in
// go.mod (and therefore visible to Dependabot and mockery's package loader).
// The generated mocks live in the subdirectories.
package temporalapimocks

import (
	_ "go.temporal.io/api/operatorservice/v1"
	_ "go.temporal.io/api/workflowservice/v1"
)
