package es

import (
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func Test_query(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query()
		})
	}
}

func Test_create(t *testing.T) {
	tests := []struct {
		name string
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			create()
		})
	}
}

func Test_delete(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delete()
		})
	}
}

func Test_update(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			update()
		})
	}
}

func Test_gets(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gets()
		})
	}
}

func Test_list(t *testing.T) {
	type args struct {
		size int
		page int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list(tt.args.size, tt.args.page)
		})
	}
}

func Test_printEmployee(t *testing.T) {
	type args struct {
		res *elastic.SearchResult
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printEmployee(tt.args.res, tt.args.err)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
