package net

import "regexp"

var ArpTableParser = regexp.MustCompile("^[^\\d\\.]+([\\d\\.]+).+\\s+([a-f0-9:]{17}).+\\s+(.+)$")
var ArpTableTokens = 4
