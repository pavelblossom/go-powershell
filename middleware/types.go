// Copyright (c) 2017 pavelblossom. All rights reserved.

package middleware

type Middleware interface {
	Execute(cmd string) (string, string, error)
	Exit()
}
