package model

import (
	"reflect"
	"testing"

	"github.com/dandynaufaldi/idy-go/domain/model/mocks"
)

func TestNewRatingID(t *testing.T) {
	staticUUID := "47953d94-0b8c-4122-995c-bb9e384dda3f"
	mockIDProvider := &mocks.IDProvider{}
	mockIDProvider.On("Generate").Return(staticUUID).Once()

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *RatingID
	}{
		{
			name: "Hold supplied UUID",
			args: args{
				id: "f4f1b863-cfc9-467c-a061-e9229c3b47b2",
			},
			want: &RatingID{
				id: "f4f1b863-cfc9-467c-a061-e9229c3b47b2",
			},
		},
		{
			name: "Generate UUID when needed",
			args: args{
				id: "",
			},
			want: &RatingID{
				id: staticUUID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRatingID(tt.args.id, mockIDProvider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatingID() = %v, want %v", got, tt.want)
			}
		})
	}
	mockIDProvider.AssertExpectations(t)
}

func TestRatingID_ID(t *testing.T) {
	type fields struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ratingID := &RatingID{
				id: tt.fields.id,
			}
			if got := ratingID.ID(); got != tt.want {
				t.Errorf("RatingID.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRatingID_Equals(t *testing.T) {
	type fields struct {
		id string
	}
	type args struct {
		other *RatingID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ratingID := &RatingID{
				id: tt.fields.id,
			}
			if got := ratingID.Equals(tt.args.other); got != tt.want {
				t.Errorf("RatingID.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
