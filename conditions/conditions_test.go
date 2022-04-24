// Copyright 2022 Ameridata.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conditions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThen(t *testing.T) {
	assert.Equal(t, IfThen(1 == 1, "Yes"), "Yes")
	assert.Equal(t, IfThen(1 != 1, "Woo"), nil)
	assert.Equal(t, IfThen(1 < 2, "Less"), "Less")
}

func TestIfThenElse(t *testing.T) {
	assert.Equal(t, IfThenElse(1 == 1, "Yes", false), "Yes")
	assert.Equal(t, IfThenElse(1 != 1, nil, 1), 1)
	assert.Equal(t, IfThenElse(1 < 2, nil, "No"), nil)
}

func TestDefaultIfNil(t *testing.T) {
	assert.Equal(t, DefaultIfNil(nil, nil), nil)
	assert.Equal(t, DefaultIfNil(nil, ""), "")
	assert.Equal(t, DefaultIfNil("A", "B"), "A")
	assert.Equal(t, DefaultIfNil(true, "B"), true)
	assert.Equal(t, DefaultIfNil(1, false), 1)
}
