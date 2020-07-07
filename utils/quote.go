// Copyright (c) 2017 pavelblossom. All rights reserved.

package utils

import "strings"

func QuoteArg(s string) string {
	return "'" + strings.Replace(s, "'", "\"", -1) + "'"
}
