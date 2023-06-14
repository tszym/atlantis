// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"

	command "github.com/runatlantis/atlantis/server/events/command"
)

func AnyCommandProjectResult() command.ProjectResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(command.ProjectResult))(nil)).Elem()))
	var nullValue command.ProjectResult
	return nullValue
}

func EqCommandProjectResult(value command.ProjectResult) command.ProjectResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue command.ProjectResult
	return nullValue
}

func NotEqCommandProjectResult(value command.ProjectResult) command.ProjectResult {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue command.ProjectResult
	return nullValue
}

func CommandProjectResultThat(matcher pegomock.ArgumentMatcher) command.ProjectResult {
	pegomock.RegisterMatcher(matcher)
	var nullValue command.ProjectResult
	return nullValue
}
