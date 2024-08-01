/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// SelfSubjectReviewsGetter has a method to return a SelfSubjectReviewInterface.
// A group's client should implement this interface.
type SelfSubjectReviewsGetter interface {
	SelfSubjectReviews() SelfSubjectReviewInterface
}

// SelfSubjectReviewInterface has methods to work with SelfSubjectReview resources.
type SelfSubjectReviewInterface interface {
	Create(ctx context.Context, selfSubjectReview *v1.SelfSubjectReview, opts metav1.CreateOptions) (*v1.SelfSubjectReview, error)
	SelfSubjectReviewExpansion
}

// selfSubjectReviews implements SelfSubjectReviewInterface
type selfSubjectReviews struct {
	client rest.Interface
}

// newSelfSubjectReviews returns a SelfSubjectReviews
func newSelfSubjectReviews(c *AuthenticationV1Client) *selfSubjectReviews {
	return &selfSubjectReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a selfSubjectReview and creates it.  Returns the server's representation of the selfSubjectReview, and an error, if there is any.
func (c *selfSubjectReviews) Create(ctx context.Context, selfSubjectReview *v1.SelfSubjectReview, opts metav1.CreateOptions) (result *v1.SelfSubjectReview, err error) {
	result = &v1.SelfSubjectReview{}
	err = c.client.Post().
		Resource("selfsubjectreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectReview).
		Do(ctx).
		Into(result)
	return
}
