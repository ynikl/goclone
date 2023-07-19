package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_splitURLIntoFragments(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name             string
		args             args
		wantUrlFragments map[string]string
	}{
		{
			"",
			args{url: "https://github.com/gkeaisin/goclone"},
			map[string]string{
				"host_name": "github.com",
				"org":       "gkeaisin",
				"repo":      "goclone"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUrlFragments := splitURLIntoFragments(tt.args.url); !reflect.DeepEqual(gotUrlFragments, tt.wantUrlFragments) {
				t.Errorf("splitURLIntoFragments() = %v, want %v", gotUrlFragments, tt.wantUrlFragments)
			}
		})
	}
}

func Test_getTargetDir(t *testing.T) {

	got, err := getTargetDir()

	fmt.Println(got, err)

}
