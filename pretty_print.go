package prettygo

import "fmt"

func Print(kind string, v interface{}) {
	b, err := Encode(kind, v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

func PrintJSON(v interface{}) {
	Print("prettyjson", v)
}

func PrintYAML(v interface{}) {
	Print("yaml", v)
}

func PrintTOML(v interface{}) {
	Print("toml", v)
}
