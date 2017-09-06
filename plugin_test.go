package main

import (
	"os/exec"
	"reflect"
	"testing"
)

func Test_destroyCommand(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			"default",
			args{config: Config{}},
			exec.Command("terraform", "destroy", "-force"),
		},
		{
			"with targets",
			args{config: Config{Targets: []string{"target1", "target2"}}},
			exec.Command("terraform", "destroy", "-target=target1", "-target=target2", "-force"),
		},
		{
			"with parallelism",
			args{config: Config{Parallelism: 5}},
			exec.Command("terraform", "destroy", "-parallelism=5", "-force"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destroyCommand(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("destroyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyCommand(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			"default",
			args{config: Config{}},
			exec.Command("terraform", "apply", "plan.tfout"),
		},
		{
			"with targets",
			args{config: Config{Targets: []string{"target1", "target2"}}},
			exec.Command("terraform", "apply", "--target", "target1", "--target", "target2", "plan.tfout"),
		},
		{
			"with parallelism",
			args{config: Config{Parallelism: 5}},
			exec.Command("terraform", "apply", "-parallelism=5", "plan.tfout"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyCommand(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_terraformCommand(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			"default",
			args{config: Config{}},
			exec.Command("terraform", "apply", "plan.tfout"),
		},
		{
			"destroy",
			args{config: Config{Destroy: true}},
			exec.Command("terraform", "destroy", "-force"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := terraformCommand(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("terraformCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planCommand(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		{
			"default",
			args{config: Config{}},
			exec.Command("terraform", "plan", "-out=plan.tfout"),
		},
		{
			"destroy",
			args{config: Config{Destroy: true}},
			exec.Command("terraform", "plan", "-destroy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := planCommand(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}