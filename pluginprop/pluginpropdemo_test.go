package pluginprop

import (
	"reflect"
	"testing"
)

func TestExploreluginProp(t *testing.T) {
	tests := []struct {
		name  string
		prop  any
		entry *PluginEntry
	}{
		{
			name:  "Test with bool",
			prop:  true,
			entry: &PluginEntry{Name: "BoolTest"},
		},
		{
			name:  "Test with int",
			prop:  42,
			entry: &PluginEntry{Name: "IntTest"},
		},
		{
			name:  "Test with string",
			prop:  "test string",
			entry: &PluginEntry{Name: "StringTest"},
		},
		{
			name:  "Test with []string",
			prop:  []string{"one", "two", "three"},
			entry: &PluginEntry{Name: "StringSliceTest"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Test %s panicked: %v", tt.name, r)
				}
			}()

			sliceType := reflect.TypeOf([]string{})
			typ := reflect.TypeOf(tt.prop)
			switch typ.Kind() {
			case reflect.Bool:
				ExploreluginProp(tt.entry, tt.prop.(bool))
			case sliceType.Kind():
				ExploreluginProp(tt.entry, tt.prop.([]string))
			case reflect.Int:
				ExploreluginProp(tt.entry, tt.prop.(int))

			case reflect.String:
				ExploreluginProp(tt.entry, tt.prop.(string))
			}

			if tt.entry.Val != tt.prop {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.prop, tt.entry.Val)
			}
		})
	}

	//	t.Run("Test with unsupported type", func(t *testing.T) {
	//		defer func() {
	//			if r := recover(); r == nil {
	//				t.Errorf("Expected panic for unsupported type, but did not panic")
	//			}
	//		}()
	//		ExploreluginProp(&PluginEntry{Name: "UnsupportedTest"}, 3.14)
	//	})
}
