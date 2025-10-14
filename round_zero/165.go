package roundzero

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v0 := strings.Split(version1, ".")
	v1 := strings.Split(version2, ".")

	//status := 0
	for i := 0; i < max(len(v0), len(v1)); i++ {
		tmpV0 := int64(0)
		if i < len(v0) {
			tmpV0, _ = strconv.ParseInt(v0[i], 10, 64)
		}

		tmpV1 := int64(0)
		if i < len(v1) {
			tmpV1, _ = strconv.ParseInt(v1[i], 10, 64)
		}

		if tmpV0 == tmpV1 {
			continue
		}

		if tmpV0 > tmpV1 {
			return 1
		}
		return -1

	}
	return -1
}
