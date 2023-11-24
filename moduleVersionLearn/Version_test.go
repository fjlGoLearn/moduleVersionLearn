package moduleVersionLearn

import "testing"

func TestGetVersion(t *testing.T) {
	if GetVersion() != Version {
		t.Errorf("GetVersion() != Version")
	}
}
