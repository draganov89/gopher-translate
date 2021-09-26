package pattern_translate

import (
	"sort"
	"sync"
)

type HistoryHandler struct {
	translMux    sync.RWMutex
	keysMux      sync.RWMutex
	translations map[string]string
	keys         []string
}

func (h *HistoryHandler) sortKeys() {
	h.keysMux.Lock()
	sort.Slice(h.keys, func(p, q int) bool {
		return h.keys[p] < h.keys[q]
	})
	h.keysMux.Unlock()
}
func (h *HistoryHandler) addToHistory(eng, goph string) {
	h.keysMux.Lock()
	h.keys = append(h.keys, eng)
	h.keysMux.Unlock()

	h.translMux.Lock()
	h.translations[eng] = goph
	h.translMux.Unlock()
}
