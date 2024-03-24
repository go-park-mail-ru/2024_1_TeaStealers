//go:build tools

//go:generate go build -o ../bin/mockgen github.com/golang/mock/mockgen

package tools

import (
	_ "github.com/golang/mock/mockgen"
)
