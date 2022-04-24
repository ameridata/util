// Copyright 2022 Ameridata.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// It creates a new MD5 hash, then passes it to the stringHasher function along with the text to be hashed
// and return the hashed text
func MD5(text string) string {
	algorithm := md5.New()
	return stringHasher(algorithm, text)
}

// It creates a new SHA256 hash, then passes it to the stringHasher function along with the text to be hashed
// and return the hashed text
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

// It creates a new SHA512 hash, then passes it to the stringHasher function along with the text to be hashed
// and return the hashed text
func SHA512(text string) string {
	algorithm := sha512.New()
	return stringHasher(algorithm, text)
}

// It takes a hash algorithm and a string, and returns the hexadecimal representation of the hash of the string
func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

// It takes a string, converts it to a byte array, and then encodes it using the base64 encoding scheme
func EncodeBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// It takes a string, decodes it from base64, and returns the decoded string
func DecodeBase64(encodedText string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedText)
	return string(decoded), err
}
