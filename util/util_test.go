package util

import "testing" 
 
func TestUpdatePathPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20"
	UpdatePathPagination(&path, 21, 50)
	
	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=21&size=50", "Error while trying to replae path", t)
}

func TestAddPathPaginationWithOutPaginationOnPath(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES"
	UpdatePathPagination(&path, 0, 20)
	
	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20", "Error while trying to replae path", t)
}
func TestAddPathSizeOnPaginationData(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0"
	UpdatePathPagination(&path, 0, 20)
	
	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=20", "Error while trying to replae path", t)
}
func TestAddPathStartOnPaginationData(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?size=0"
	UpdatePathPagination(&path, 0, 20)
	
	assert(path, "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?size=20&start=0", "Error while trying to replae path", t)
}

func TestInitialPathWithStartPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=10"
	start := FindStartPagination(path)

	if start != 10 {
		t.Errorf("%s, got: %d, want: %d.", "Pagination parameter not found", start, 10)
	}
}
func TestInitialPathWithoutStartPagination(t *testing.T) {
	path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES"
	start := FindStartPagination(path)

	if start != 0 {
		t.Errorf("%s, got: %d, want: %d.", "Pagination parameter not found", start, 0)
	}
}

func assert(value, compare string, message string, t *testing.T) {
	if value != compare {
		t.Errorf("%s, got: %s, want: %s.", message, value, compare)
	}
}