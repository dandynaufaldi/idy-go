package value

import (
	"reflect"
	"testing"
)

func TestNewRatingValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    *RatingValue
		wantErr bool
	}{
		{
			name: "Value 0 to 5",
			args: args{
				value: 5,
			},
			want:    &RatingValue{5},
			wantErr: false,
		},
		{
			name: "Value not in 0 to 5",
			args: args{
				value: 6,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRatingValue(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRatingValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatingValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
