package networking

import "testing"

func TestGetLocalNetworkAddresses(t *testing.T) {
	localAddresses, err := GetLocalNetworkAddresses()
	if err != nil {
		t.Error(err)
	}
	t.Log(localAddresses)
}
