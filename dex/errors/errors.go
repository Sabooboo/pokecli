package dex

import (
	"github.com/Sabooboo/pokecli/ui/typdef"
)

const (
	FileNotFound typdef.Error = "file not found"
	FetchFailed  typdef.Error = "failed to fetch from web"
	DexNotFound  typdef.Error = "no pokedex found at id"
)
