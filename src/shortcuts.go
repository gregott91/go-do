package godo

// ShortcutObservable is the observable for the shortcuts
type ShortcutObservable struct {
	Handlers map[uint64]func()
}

// Register registers a new shortcut
func (observable *ShortcutObservable) Register(key uint64, handler func()) {
	observable.Handlers[key] = handler
}

// Trigger triggers a shortcut key entry
func (observable *ShortcutObservable) Trigger(key uint64) {
	if val, ok := observable.Handlers[key]; ok {
		val()
	}
}
