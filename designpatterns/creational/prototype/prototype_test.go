package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Not implemented yet")
	}
	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("Item1 cannot be equal to the whitePrototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assetion for shirt1 couldnt be done successfully")
	}
	shirt1.SKU = "abbcc"
	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assetion for shirt1 couldnt be done successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKUs of shirt1 and shirt2 must be different")
	}
	if shirt1 == shirt2 {
		t.Error("Shirt1 and shirt2 cannot be equal")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The memory position of the shirts are different %p !=%p\n\n", &shirt1, &shirt2)
}
