// +build linux

package fs

import (
	"testing"
)

const (
	classidBefore = "0x100002"
	classidAfter  = "0x100001"
)

func TestNetClsSetClassid(t *testing.T) {
	helper := NewCgroupTestUtil("net_cls", t)
	defer helper.cleanup()

	helper.writeFileContents(map[string]string{
		"net_cls.classid": classidBefore,
	})

	helper.CgroupData.c.NetClsClassid = classidAfter
	netcls := &NetClsGroup{}
	if err := netcls.Set(helper.CgroupPath, helper.CgroupData.c); err != nil {
		t.Fatal(err)
	}

	// As we are in mock environment, we can't get correct value of classid from
	// net_cls.classid.
	// So. we just judge if we successfully write classid into file
	value, err := getCgroupParamString(helper.CgroupPath, "net_cls.classid")
	if err != nil {
		t.Fatalf("Failed to parse net_cls.classid - %s", err)
	}
	if value != classidAfter {
		t.Fatal("Got the wrong value, set net_cls.classid failed.")
	}
}
