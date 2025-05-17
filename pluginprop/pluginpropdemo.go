package pluginprop

import "fmt"

type PluginProp interface {
	bool | int | string | []string
}

func ExploreluginProp[T PluginProp](entry *PluginEntry, prop T) {
	switch any(prop).(type) {
	case bool:
		fmt.Println("Found bool pluing prop: ", prop)
	case int:
		fmt.Println("Found int pluing prop: ", prop)
	case string:
		fmt.Println("Found string pluing prop: ", prop)
	case []string:
		fmt.Println("Found []string pluing prop: ", prop)
		// entry.Strings = &v
	default:
		panic("unsupported type")
	}
	entry.Val = prop
}

type PluginEntry struct {
	Name string
	Val  any
}

// NODES:
// gbxga8
// rj6oqw
