package lcms

import "testing"

func TestOpenProfileFromMem(t *testing.T) {
	cases := []struct {
		d   []byte
		err bool
	}{
		{nil, true},
		{[]byte{}, true},
		{[]byte{0, 0, 0}, true},
	}
	for n, c := range cases {
		_, err := OpenProfileFromMem(c.d)
		if c.err && err == nil {
			t.Errorf("%d: not error", n)
		} else if !c.err && err != nil {
			t.Errorf("%d: %s", n, err)
		}
	}
}
