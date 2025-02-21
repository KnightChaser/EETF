// internal/tui/tui.go
package tui

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

// Launches an interactive search for the given items, using go-fzf.
func ShowInteractiveSearch(items []string) {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}

	// Just selected items
	for _, i := range idxs {
		fmt.Println(items[i])
	}
}
