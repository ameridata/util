// Copyright 2022 Ameridata.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashes

// Online check tool: https://emn178.github.io/online-tools/index.html

import (
	"testing"

	"github.com/ameridata/util/strings"
	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	data := MD5("Ameridata")
	assert.Equal(t, data, "695c35a4808dbdb2c86ea17594fe1b4e")
}

func TestSHA256(t *testing.T) {
	data := SHA256("Ameridata")
	assert.Equal(t, data, "8191ea83dd4b6eeabf7df86b21c6d95c670a215b5bbc75ff691b6be9d2c51ad7")
}

func TestSHA512(t *testing.T) {
	data := SHA512("Ameridata")
	assert.Equal(t, data, "ab3603e2805c13f7f8057ed4b58cda57cf47bc7e7d56f83ca64c181286c6d2c2f74154c166ee4ec0e7b4e26a4f90cd5a71b42a7f6308d129a307720e4c31d9f8")
}

func TestEncodeBase64(t *testing.T) {
	data := EncodeBase64("Ameridata")
	assert.Equal(t, data, "QW1lcmlkYXRh")
}

func TestDecodeBase64(t *testing.T) {
	// Success case
	dataSuccess, err := DecodeBase64("QW1lcmlkYXRh")
	assert.Equal(t, dataSuccess, "Ameridata")
	assert.True(t, err == nil)

	// Fail case
	dataFail, err := DecodeBase64("Q@W@1@l@c@m@l@k@Y@X@R@h")
	assert.Equal(t, dataFail, strings.Empty)
	assert.True(t, err != nil)
}
