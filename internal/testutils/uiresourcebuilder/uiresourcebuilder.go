package uiresourcebuilder

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
)

type UIResourceBuilder struct {
	name          string
	disabledCount int
}

func New(name string) *UIResourceBuilder {
	return &UIResourceBuilder{
		name: name,
	}
}

func (u *UIResourceBuilder) WithDisabledCount(i int) *UIResourceBuilder {
	u.disabledCount = i
	return u
}

func (u *UIResourceBuilder) Build() *v1alpha1.UIResource {
	return &v1alpha1.UIResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: u.name,
		},
		Status: v1alpha1.UIResourceStatus{
			DisableStatus: v1alpha1.DisableResourceStatus{
				DisabledCount: int32(u.disabledCount),
			},
		},
	}
}