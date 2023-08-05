package openweathermap

import "testing"

func TestGeoLocateDecodeZipcode(t *testing.T) {
	test_data := []byte("{\"zip\": \"95125\",\"name\": \"San Jose\",\"lat\": 37.296,\"lon\": -121.8939,\"country\": \"US\"}")
	geolocation, err := decodeGeoLocationData(test_data)
	if err != nil {
		t.Fatal(err)
	}

	if geolocation.Country != "US" {
		t.Fatal("Incorrect country")
	}
	if geolocation.Latitude != 37.296 {
		t.Fatal("Incorrect latitude")
	}
	if geolocation.Longitude != -121.8939 {
		t.Fatal("Incorrect longitude")
	}
	if geolocation.Name != "San Jose" {
		t.Fatal("Incorrect city")
	}
	if geolocation.Zip != "95125" {
		t.Fatal("Incorrect zipcode")
	}
}

func TestGeoLocateDecodeAddress(t *testing.T) {
	test_data := []byte("[{\"name\": \"San Jose\",\"local_names\": {\"ar\": \"سان خوسيه\",\"eo\": \"San-Joseo\",\"zh\": \"聖荷西\",\"pt\": \"San José\",\"oc\": \"San José\",\"en\": \"San Jose\",\"vi\": \"San Jose\",\"gl\": \"San Xosé\",\"es\": \"San José\",\"ru\": \"Сан-Хосе\",\"uk\": \"Сан-Хосе\",\"am\": \"ሳን ሆዜ\"},\"lat\": 37.3361663,\"lon\": -121.890591,\"country\": \"US\",\"state\": \"California\"}]")
	geolocation, err := decodeGeoLocationData(test_data)
	if err != nil {
		t.Fatal(err)
	}

	if geolocation.Country != "US" {
		t.Fatal("Incorrect country")
	}
	if geolocation.Latitude != 37.296 {
		t.Fatal("Incorrect latitude")
	}
	if geolocation.Longitude != -121.8939 {
		t.Fatal("Incorrect longitude")
	}
	if geolocation.Name != "San Jose" {
		t.Fatal("Incorrect city")
	}
	if geolocation.Zip != "95125" {
		t.Fatal("Incorrect zipcode")
	}
}
