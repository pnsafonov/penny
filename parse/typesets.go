package parse

import (
	"strings"
	"unicode"
)

const (
	typeSep     = " "
	keyValueSep = "="
	valuesSep   = ","
	builtins    = "BUILTINS"
	numbers     = "NUMBERS"
)

func getCasePair(str string) (result string, isUpper bool, ok bool) {
	n := len(str)
	if n == 0 {
		return "", false, false
	}
	var r0 rune
	for _, r := range str {
		r0 = r
		break
	}
	if r0 > unicode.MaxASCII {
		return "", false, false
	}
	isUpper = unicode.IsUpper(r0)
	b1 := []byte(str)
	if isUpper {
		b1[0] = byte(unicode.ToLower(r0))
	} else {
		b1[0] = byte(unicode.ToUpper(r0))
	}
	result = string(b1)
	ok = true
	return
}

func changeCase(str string, isUpper bool) (result string, ok bool) {
	var r0 rune
	for _, r := range str {
		r0 = r
		break
	}
	if r0 > unicode.MaxASCII {
		return "", false
	}
	b1 := []byte(str)
	if isUpper {
		b1[0] = byte(unicode.ToLower(r0))
	} else {
		b1[0] = byte(unicode.ToUpper(r0))
	}
	result = string(b1)
	ok = true
	return
}

func isStringInSlice(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// TypeSet turns a type string into a []map[string]string
// that can be given to parse.Generics for it to do its magic.
//
// Acceptable args are:
//
//     Person=man
//     Person=man Animal=dog
//     Person=man Animal=dog Animal2=cat
//     Person=man,woman Animal=dog,cat
//     Person=man,woman,child Animal=dog,cat Place=london,paris
func TypeSet(arg string) ([]map[string]string, error) {

	types := make(map[string][]string)
	var keys []string
	for _, pair := range strings.Split(arg, typeSep) {
		segs := strings.Split(pair, keyValueSep)
		if len(segs) != 2 {
			return nil, &errBadTypeArgs{Arg: arg, Message: "Generic=Specific expected"}
		}
		key := segs[0]
		keys = append(keys, key)
		types[key] = make([]string, 0)

		key1, isUpper, ok1 := getCasePair(key)

		seg1 := strings.Split(segs[1], valuesSep)
		for _, t := range seg1 {
			if t == builtins {
				types[key] = append(types[key], Builtins...)
			} else if t == numbers {
				types[key] = append(types[key], Numbers...)
			} else {
				types[key] = append(types[key], t)

				if ok1 && !strings.Contains(t, ".") {
					t1, ok2 := changeCase(t, isUpper)
					if ok2 {
						if !isStringInSlice(keys, key1) {
							keys = append(keys, key1)
						}
						types[key1] = append(types[key1], t1)
					}
				}

			}
		}
	}

	cursors := make(map[string]int)
	for _, key := range keys {
		cursors[key] = 0
	}

	outChan := make(chan map[string]string)
	go func() {
		buildTypeSet(keys, 0, cursors, types, outChan)
		close(outChan)
	}()

	var typeSets []map[string]string
	for typeSet := range outChan {
		typeSets = append(typeSets, typeSet)
	}

	return typeSets, nil

}

func buildTypeSet(keys []string, keyI int, cursors map[string]int, types map[string][]string, out chan<- map[string]string) {
	key := keys[keyI]
	for cursors[key] < len(types[key]) {
		if keyI < len(keys)-1 {
			buildTypeSet(keys, keyI+1, copycursors(cursors), types, out)
		} else {
			// build the typeset for this combination
			ts := make(map[string]string)
			for k, vals := range types {
				ts[k] = vals[cursors[k]]
			}
			out <- ts
		}
		cursors[key]++
	}
}

func copycursors(source map[string]int) map[string]int {
	copy := make(map[string]int)
	for k, v := range source {
		copy[k] = v
	}
	return copy
}
