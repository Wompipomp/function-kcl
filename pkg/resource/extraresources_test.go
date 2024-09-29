package resource

import (
	fnv1 "github.com/crossplane/function-sdk-go/proto/v1"
	"reflect"
	"testing"
)

func TestExtraResourcesRequirement_ToResourceSelector(t *testing.T) {
	tests := []struct {
		name   string
		fields ExtraResourcesRequirement
		want   *fnv1.ResourceSelector
	}{
		{
			name: "TestWithLabels",
			fields: ExtraResourcesRequirement{
				APIVersion: "v1",
				Kind:       "Namespace",
				MatchLabels: map[string]string{
					"app": "TestApp",
				},
			},
			want: &fnv1.ResourceSelector{
				ApiVersion: "v1",
				Kind:       "Namespace",
				Match: &fnv1.ResourceSelector_MatchLabels{
					MatchLabels: &fnv1.MatchLabels{
						Labels: map[string]string{
							"app": "TestApp",
						},
					},
				},
			},
		},
		{
			name: "TestWithName",
			fields: ExtraResourcesRequirement{
				APIVersion: "v1",
				Kind:       "Namespace",
				MatchName:  "TestName",
			},
			want: &fnv1.ResourceSelector{
				ApiVersion: "v1",
				Kind:       "Namespace",
				Match: &fnv1.ResourceSelector_MatchName{
					MatchName: "TestName",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.ToResourceSelector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtraResourcesRequirement.ToResourceSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}
