package unidoc

import (
	"testing"
)

func TestParseDocx(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Docx
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test parse 1",
		// 	args: args{
		// 		path: "./../../../iFlytek_CN_EN.docx",
		// 	},
		// },
		{
			name: "test parse 1",
			args: args{
				path: "./../../../AI Translation Demo File - after checking.docx",
			},
		},
		// {
		// 	name: "test parse 2",
		// 	args: args{
		// 		path: "./../../../1255.docx",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDocx(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDocx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got=%+v", got)
		})
	}
}

func TestFindAndReplaceNode(t *testing.T) {
	type args struct {
		path       string
		no         int64
		newContent string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test find 1",
			args: args{
				path:       "./../../../AI Translation Demo File - after checking.docx",
				no:         5,
				newContent: "2122.概述",
			},
		},
		// {
		// 	name: "test find 2",
		// 	args: args{
		// 		path:       "./../../../1255.docx", // "AI Translation Demo File - after checking.docx"
		// 		no:         4,
		// 		newContent: "This is a new replaced content",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FindAndReplaceNode(tt.args.path, tt.args.no, tt.args.newContent); err != nil {
				t.Errorf("FindAndReplaceNode() error = %v", err)
			}
		})
	}
}
