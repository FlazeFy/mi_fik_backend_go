package generator

import (
	"app/packages/helpers/typography"
)

func GenerateNormalMsg(ctx string, total int) string {
	res := typography.UcFirst(ctx)

	if total > 0 {
		return res + " found"
	} else {
		return res + " not found"
	}
}
