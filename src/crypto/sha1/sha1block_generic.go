// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 && !386 && !arm && !s390x && !arm64

package sha1

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}

func blockString(dig *digest, s string) {
	blockGeneric(dig, s)
}
