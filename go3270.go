// Copyright 2009 Richard Thorp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package go3270 is a Golang interface to the s3270 binary which will allow you to communicate with an IBM mainframe session programmatically.
*/
package go3270

// Session is a new s3270 session
type Session struct {
	url string
}