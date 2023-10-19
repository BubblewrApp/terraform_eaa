package client

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SetAttrs(d *schema.ResourceData, AttributeValues map[string]interface{}) error {
	for attr, value := range AttributeValues {
		if err := d.Set(attr, value); err != nil {
			return err
		}
	}
	return nil
}

func SetAdvancedSettings(d *schema.ResourceData, settings AdvancedSettings) error {
	advancedSettingsMap := make(map[string]interface{})

	v := reflect.ValueOf(settings)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("json")
		tagName := strings.Split(tag, ",")[0]

		// Skip fields with empty values
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		if field.Kind() == reflect.Ptr {
			// Dereference pointer fields
			advancedSettingsMap[tagName] = field.Elem().Interface()
		} else {
			advancedSettingsMap[tagName] = field.Interface()
		}
	}

	return d.Set("advanced_settings", advancedSettingsMap)
}
