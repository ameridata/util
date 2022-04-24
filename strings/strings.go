// Copyright 2022 Ameridata.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// IsEmpty returns true if the text is empty.
func IsEmpty(text string) bool {
	return len(text) == 0
}

// If the text is not empty, return true.
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// It returns true if the string is empty or contains only whitespace
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// Returns true if the string is not empty and not whitespace
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}
