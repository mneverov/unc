package unc

func foo() {
	var err error
	if 1 == 2 {
	}

	if nil != err {
	} // want "lhs nil"

	if 1 == 2 ||
		err == nil ||
		nil == err {
	} // want "lhs nil"

	if nil == err && // want "lhs nil"
		nil == bar() {
	} // want "lhs nil"

	if 1 == 2 {
	} else if nil == err {
	} // want "lhs nil"

	if 1 == 2 || (2 == 3 && nil == err) {
	} // want "lhs nil"

	if nil != bar() {
	} // want "lhs nil"

	if nil != func() error { // want "lhs nil"
		return nil
	}() {
	}
}

func bar() error {
	return nil
}
