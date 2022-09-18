package status

// HasStatus accepts an input status and will check if the flag is set using bitwise OR and return true if the flag is set.
func HasStatus[T ~int](status T, flag T) bool {
	return status|flag == status
}

// NotHasStatus accepts an input status and will check if the flag is not set.
func NotHasStatus[T ~int](status T, flag T) bool {
	return !HasStatus(status, flag)
}

// UnsetStatus accepts an input status reference and will unset the flag using bitwise AND NOT.
func UnsetStatus[T ~int](status *T, flag T) {
	*status = *status &^ flag
}

// ForceStatus accepts an input status reference and will simply overwrite the status with the flag.
func ForceStatus[T ~int](status *T, flag T) {
	*status = flag
}

// SetStatus accepts an input status reference and will set the flag to the status using bitwise OR.
func SetStatus[T ~int](status *T, flag T) {
	*status = *status | flag
}

// SetMulti allows you to set multiple flags at once
func SetMulti[T ~int](status *T, flags ...T) {
	if len(flags) == 0 {
		return
	}
	for _, flag := range flags {
		SetStatus[T](status, flag)
	}
}

// HasMulti allows you to check for many flags at once.
func HasMulti[T ~int](status T, flags ...T) bool {
	if len(flags) == 0 {
		return false
	}
	for _, flag := range flags {
		if NotHasStatus[T](status, flag) {
			return false
		}
	}
	return true
}

// UnsetMulti allows you to clear multiple flags at once.
func UnsetMulti[T ~int](status *T, flags ...T) {
	if len(flags) == 0 {
		return
	}
	for _, flag := range flags {
		UnsetStatus[T](status, flag)
	}
}
