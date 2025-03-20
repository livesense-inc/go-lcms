package lcms

import "testing"

func TestCreateTransform(t *testing.T) {
	cases := []struct {
		srcProf *Profile
		srcType CMSType
		dstProf *Profile
		dstType CMSType
		err     bool
	}{
		{nil, TYPE_CMYK_8, nil, TYPE_RGBA_8, true},
		{createTestProfile(t), TYPE_CMYK_8, createTestProfile(t), TYPE_RGBA_8, true},
		{createTestProfile(t), TYPE_RGB_8, createTestProfile(t), TYPE_RGBA_8, false},
	}
	for n, c := range cases {
		_, err := CreateTransform(c.srcProf, c.srcType, c.dstProf, c.dstType)
		if c.err && err == nil {
			t.Errorf("%d: not error", n)
		} else if !c.err && err != nil {
			t.Errorf("%d: %s", n, err)
		}
	}
}

func TestDoTransform(t *testing.T) {
	tf, err := CreateTransform(createTestProfile(t), TYPE_RGB_8, createTestProfile(t), TYPE_RGBA_8)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		in   []uint8
		out  []uint8
		size int
		err  bool
	}{
		{nil, nil, 0, true},
		{[]uint8{}, []uint8{}, 0, true},
		{[]uint8{0, 0, 0}, []uint8{0, 0, 0, 0}, 1, false},
	}
	for n, c := range cases {
		if err := tf.DoTransform(c.in, c.out, c.size); c.err && err == nil {
			t.Errorf("%d: not error", n)
		} else if !c.err && err != nil {
			t.Errorf("%d: %s", n, err)
		}
	}
}

func createTestProfile(t *testing.T) *Profile {
	t.Helper()
	p, err := CreateSRGBProfile()
	if err != nil {
		t.Fatal(err)
	}
	return p
}
