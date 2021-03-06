// Copyright 2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// From Linux header: /include/uapi/linux/kdev_t.h
const (
	minorBits = 8
	minorMask = (1 << minorBits) - 1
)

func major(dev uint64) uint64 {
	return dev >> minorBits
}

func minor(dev uint64) uint64 {
	return dev & minorMask
}
