// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	logging "github.com/runatlantis/atlantis/server/logging"
)

func EqPtrToLoggingSimpleLogger(value logging.SimpleLogging) logging.SimpleLogging {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue logging.SimpleLogging
	return nullValue
}
