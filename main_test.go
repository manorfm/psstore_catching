package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"net/http/httptest"
	"testing"
)

func TestShouldPanicWithNonArgs(t *testing.T) {
	os.Args = []string{`exec.exe`}

	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	main()
	assert.Equal(t, 1, got, "got wrong exit code")
}

func TestShouldPanicOnlyWithPathInArgs(t *testing.T) {
	os.Args = []string{`exec.exe`,`localhost.com`}
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	main()

	assert.Equal(t, 1, got, "got wrong exit code")
}
func TestShouldMainPanicErrorWithOnlyStringArgs(t *testing.T) {
	os.Args = []string{`exec.go`, `localhost.com`, `error`}
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	main()

	assert.Equal(t, 2, got, "got wrong exit code")
}

func TestShouldMainPanicWithInacessibleServer(t *testing.T) {
	os.Args = []string{`exec.go`, `localhost.com`, strconv.Itoa(100)}
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	main()

	assert.Equal(t, 3, got, "got wrong exit code")
}

func TestArgs(t *testing.T) {
	gamesRepository := [2]string {`{
		"age_limit": 7,
		"name": "Game test",
		"playable_platform": ["PS4"],
		"provider_name": "Activision",
		"star_rating": {
			"score": "4.0",
			"total": "2",
			"count": [{
				"star": 4,
				"count": 2
			}]
		}
	}`, 
	`{
		"age_limit": 18,
		"name": "God of war 3",
		"playable_platform": ["PS3"],
		"provider_name": "Sony",
		"star_rating": {
			"score": "5.0",
			"total": "12",
			"count": [{
				"star": 5,
				"count": 12
			}]
		}
	}`}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		index, e := strconv.ParseInt(req.FormValue("start"), 10, 32)
		if e != nil {
			t.Errorf("Error getting start value from Request: %s", e)
		}

		ioutil.NopCloser(
			bytes.NewBufferString(
				fmt.Sprintf(`{ 
					"links": [%s],
					"size": 1,
					"start": %d,
					"total_results": 2
					}`, 
					gamesRepository[index], 
					index,
				),
			),
		)

		rw.Write([]byte(
			fmt.Sprintf(`{ 
				"links": [%s],
				"size": 1,
				"start": %d,
				"total_results": 2
				}`, 
				gamesRepository[index], 
				index,
			),
		))
	}))
	// Close the server when test finishes
	defer server.Close()

	os.Args = []string{"exec.go", server.URL, strconv.Itoa(1)}

	main()
}