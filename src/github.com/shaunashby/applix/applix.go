//____________________________________________________________________
// File: applix.go
//____________________________________________________________________
//
// Author: Shaun Ashby <shaun@ashby.ch>
// Created: 2019-01-14 17:24:26+0100
// Revision: $Id$
// Description: Simple application for testing OpenShift deployments
//
// Copyright (C) 2019 Shaun Ashby
//
//
//--------------------------------------------------------------------
package main

import (
	"fmt"

	"github.com/shaunashby/superutil"
)

func main() {
	fmt.Println(superutil.Message())
}
