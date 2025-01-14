// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	go_gitlab "github.com/xanzy/go-gitlab"
)

func AnyGoGitlabMergeCommentEvent() go_gitlab.MergeCommentEvent {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(go_gitlab.MergeCommentEvent))(nil)).Elem()))
	var nullValue go_gitlab.MergeCommentEvent
	return nullValue
}

func EqGoGitlabMergeCommentEvent(value go_gitlab.MergeCommentEvent) go_gitlab.MergeCommentEvent {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue go_gitlab.MergeCommentEvent
	return nullValue
}

func NotEqGoGitlabMergeCommentEvent(value go_gitlab.MergeCommentEvent) go_gitlab.MergeCommentEvent {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue go_gitlab.MergeCommentEvent
	return nullValue
}

func GoGitlabMergeCommentEventThat(matcher pegomock.ArgumentMatcher) go_gitlab.MergeCommentEvent {
	pegomock.RegisterMatcher(matcher)
	var nullValue go_gitlab.MergeCommentEvent
	return nullValue
}
