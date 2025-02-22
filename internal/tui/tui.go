// internal/tui/tui.go
package tui

import (
	"log"

	"github.com/koki-develop/go-fzf"
)

// InteractiveSelect displays an interactive fuzzy finder and returns the selected item.
// If no item is selected, it returns an empty string.
func InteractiveSelect(items []string) string {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}

	if len(idxs) == 0 {
		return ""
	}
	// Return the first selected item.
	return items[idxs[0]]
}
