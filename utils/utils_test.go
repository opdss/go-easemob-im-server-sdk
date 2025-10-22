package utils_test

import (
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"testing"
)

func TestPointerAny(t *testing.T) {
	v := utils.ToPointer(888)
	if *v != 888 {
		t.Errorf("PointerAny(888) = %v, want %v", *v, 888)
	}
	v1 := utils.ToPointer("888")
	if *v1 != "888" {
		t.Errorf("PointerAny(888) = %v, want %v", *v1, "")
	}
	var vv2 interface{}
	v2 := utils.ToPointer(vv2)
	if v2 != nil {
		t.Errorf("PointerAny(nil) = %v, want %v", v2, nil)
	}

}
