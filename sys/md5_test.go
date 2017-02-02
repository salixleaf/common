package sys

import (
	//"fmt"
	"testing"
)

func Test_FileMd5(t *testing.T) {
	t_list := map[string]string{
		"md5_example": "0ac49878018936da7279a6c2c2213aa0",
	}

	for file, md5Val := range t_list {
		md5 := FileMd5(file)
		if md5 != md5Val {
			t.Errorf("File: %s, md5(%s != %s)\n", file, md5Val, md5)
		}
	}
}

func Test_StrMd5(t *testing.T) {
	t_list := map[string]string{
		"md5_example_1": "e5dd615e1996bd4a1d8e90a8de100099",
		"md5_example_2": "3c1584f00802643c5e3d2d7cbbab762d",
		"md5_example_3": "f134514bc7daec1af58273d5c43e9097",
	}

	for file, md5Val := range t_list {
		md5 := StrMd5(file)
		if md5 != md5Val {
			t.Errorf("Str: %s, md5(%s != %s)\n", file, md5Val, md5)
		}
	}
}
