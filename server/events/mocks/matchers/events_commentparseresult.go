// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	events "github.com/runatlantis/atlantis/server/events"
)

func AnyEventsCommentParseResult() events.CommentParseResult {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.CommentParseResult))(nil)).Elem()))
	var nullValue events.CommentParseResult
	return nullValue
}

func EqEventsCommentParseResult(value events.CommentParseResult) events.CommentParseResult {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.CommentParseResult
	return nullValue
}

func NotEqEventsCommentParseResult(value events.CommentParseResult) events.CommentParseResult {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue events.CommentParseResult
	return nullValue
}

func EventsCommentParseResultThat(matcher pegomock.ArgumentMatcher) events.CommentParseResult {
	pegomock.RegisterMatcher(matcher)
	var nullValue events.CommentParseResult
	return nullValue
}
