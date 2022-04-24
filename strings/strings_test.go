// Copyright 2022 Ameridata.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.False(t, IsEmpty("text"))
	assert.False(t, IsEmpty("	"))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""))
	assert.True(t, IsNotEmpty("text"))
	assert.True(t, IsNotEmpty("	"))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""))
	assert.True(t, IsBlank("	"))
	assert.False(t, IsBlank("text"))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""))
	assert.False(t, IsNotBlank("	"))
	assert.True(t, IsNotBlank("text"))
}
