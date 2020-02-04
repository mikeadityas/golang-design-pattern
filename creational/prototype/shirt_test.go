package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache mustn't nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 must be different from the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Failed to assert item1 as shirt")
	}
	shirt1.SKU = "SKU-001"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Failed to assert item2 as shirt")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt1 must be different from Shirt2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: The address of the shirts are: %p and %p\n\n", &shirt1, &shirt2)
}

func TestGetPrice(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache mustn't nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 must be different from the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Failed to assert item1 as shirt")
	}
	shirt1.SKU = "SKU-001"

	whiteShirtPrice := shirt1.GetPrice()
	if whiteShirtPrice != 15.00 {
		t.Error("White shirt price should be 15.00")
	}
}

func TestColor(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache mustn't nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item1 == nil {
		t.Error("item1 mustn't nil")
	}

	item2, err := shirtCache.GetClone(Black)
	if err != nil {
		t.Error(err)
	}
	if item2 == nil {
		t.Error("item2 mustn't nil")
	}

	item3, err := shirtCache.GetClone(Blue)
	if err != nil {
		t.Error(err)
	}
	if item3 == nil {
		t.Error("item3 mustn't nil")
	}

	item4, err := shirtCache.GetClone(0)
	if err == nil {
		t.Error("Shirt with id 0 should be invalid")
	}
	if item4 != nil {
		t.Error("Shirt with id 0 should be nil")
	}
}
