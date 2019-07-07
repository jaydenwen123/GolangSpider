package es

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestCreate(t *testing.T) {
	type args struct {
		Index string
		Type  string
		Id    string
		doc   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *elastic.IndexResponse
		wantErr bool
	}{
		{
			"test1",
			args{Index:"person",
				Type:"man",
				Id:"7",
				doc: map[string]interface{}{"name":"wenxiaofei","age":12,"job":"student"}},
				nil,
				false,
		},
		{
			"test2",
			args{Index:"person",
				Type:"man",
				Id:"8",
				doc: map[string]interface{}{"name":"wenbingbing","age":10,"job":"student"}},
				nil,
				false,
		},
		{
			"test3",
			args{Index:"person",
				Type:"man",
				Id:"9",
				doc: map[string]interface{}{"name":"zhagnwen","age":32,"job":"engineer"}},
				nil,
				false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.Index, tt.args.Type, tt.args.Id, tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Create() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		Index string
		Type  string
		Id    string
	}
	tests := []struct {
		name    string
		args    args
		want    *elastic.GetResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.Index, tt.args.Type, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		Index string
		Type  string
		Id    string
	}
	tests := []struct {
		name    string
		args    args
		want    *elastic.DeleteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.args.Index, tt.args.Type, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
