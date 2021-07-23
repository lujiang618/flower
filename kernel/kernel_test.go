package kernel

import (
	"reflect"
	"testing"
)

func TestNewKernel(t *testing.T) {
	tests := []struct {
		name string
		want *Kernel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKernel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKernel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKernel_AddApplication(t *testing.T) {
	type args struct {
		app Application
	}
	tests := []struct {
		name string
		k    *Kernel
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.AddApplication(tt.args.app)
		})
	}
}

func TestKernel_RemoveApplication(t *testing.T) {
	tests := []struct {
		name string
		k    *Kernel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.RemoveApplication()
		})
	}
}

func TestKernel_Run(t *testing.T) {
	tests := []struct {
		name string
		k    *Kernel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.Run()
		})
	}
}

func TestKernel_Stop(t *testing.T) {
	tests := []struct {
		name string
		k    *Kernel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.Stop()
		})
	}
}
