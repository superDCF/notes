package unidoc

import (
	"github.com/unidoc/unioffice/common/license"
)

func init() {
	// err := license.SetMeteredKey(`6df3c2bf651920c7c7d723758c2f4fe7a29059b1e2eec0dd5b675880227d6772`)
	err := license.SetMeteredKey(`28f19a9d582862968eee8d276a0b80fcf051a1bacb08b1f0285be7215261236a`)
	if err != nil {
		panic(err)
	}
}
