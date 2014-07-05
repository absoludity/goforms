package fields

// XXX see reflect.DeepEqual
func ErrorsEquivalent(e1, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if e1 == nil || e2 == nil {
		return false
	}

	if e1.Error() == e2.Error() {
		return true
	}

	return false
}
