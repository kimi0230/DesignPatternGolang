package bridge

import "testing"

func TestBridge(t *testing.T) {
	brandN := new(HandsetBrandN)
	brandN.SetHandsetSoft(new(HandsetGame))
	brandN.Run()
	brandN.SetHandsetSoft(new(HandsetAddressList))
	brandN.Run()

	brandM := new(HandsetBrandM)
	brandM.SetHandsetSoft(new(HandsetGame))
	brandM.Run()
	brandM.SetHandsetSoft(new(HandsetAddressList))
	brandM.Run()
}
