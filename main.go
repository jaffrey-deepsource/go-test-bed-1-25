package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"runtime"

	"github.com/google/uuid"
)

// go-test-bed-1-25: mirror of the 1.24 test-bed but with go directive 1.25
// This file is intentionally the same as the 1.24 example to make differences
// visible if the Go toolchain introduces changes between versions.

func main() {
	fmt.Println("go runtime:", runtime.Version())

	// Generate and print a UUID from google/uuid so both projects exercise
	// the third-party dependency.
	id := uuid.New()
	fmt.Println("uuid:", id.String())

	// Linter/staticcheck test: call Sprintf and discard result (SA4006 / unused result)
	fmt.Sprintf("this formatted string is discarded: %s", id)

	// Exported function without comment (golint/staticcheck will flag)
	doSomething()

	// Unused helper below (should be detected as dead code)

	ts := httptest.NewServer(nil)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Printf("status=%d body_len=%d\n", resp.StatusCode, len(body))
}

// doSomething is intentionally left undocumented to trigger linter warnings.
func doSomething() {
	// nothing
}

func unusedHelper() string {
	return "unused"
}
