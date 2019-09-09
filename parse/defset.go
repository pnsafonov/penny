package parse

import "strings"

func DefSet(arg string) (map[string]string, error) {
	defaultSet := make(map[string]string)
	if len(arg) == 0 {
		return defaultSet, nil
	}

	for _, pair := range strings.Split(arg, typeSep) {
		segs := strings.Split(pair, keyValueSep)
		if len(segs) != 2 {
			return nil, &errBadTypeArgs{Arg: arg, Message: "Type=default expected"}
		}
		key := segs[0]
		val := segs[1]
		defaultSet[key] = val
	}
	return defaultSet, nil
}
