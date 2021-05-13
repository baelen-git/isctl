/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// TelemetryDruidHavingFilter - struct for TelemetryDruidHavingFilter
type TelemetryDruidHavingFilter struct {
	TelemetryDruidHavingDimensionSelectorFilter *TelemetryDruidHavingDimensionSelectorFilter
	TelemetryDruidHavingNumericFilter           *TelemetryDruidHavingNumericFilter
	TelemetryDruidHavingQueryFilter             *TelemetryDruidHavingQueryFilter
}

// TelemetryDruidHavingDimensionSelectorFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingDimensionSelectorFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingDimensionSelectorFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingDimensionSelectorFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingDimensionSelectorFilter: v}
}

// TelemetryDruidHavingNumericFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingNumericFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingNumericFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingNumericFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingNumericFilter: v}
}

// TelemetryDruidHavingQueryFilterAsTelemetryDruidHavingFilter is a convenience function that returns TelemetryDruidHavingQueryFilter wrapped in TelemetryDruidHavingFilter
func TelemetryDruidHavingQueryFilterAsTelemetryDruidHavingFilter(v *TelemetryDruidHavingQueryFilter) TelemetryDruidHavingFilter {
	return TelemetryDruidHavingFilter{TelemetryDruidHavingQueryFilter: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *TelemetryDruidHavingFilter) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["type"]; ok {
		switch v {
		case "dimSelector":
			var result *TelemetryDruidHavingDimensionSelectorFilter = &TelemetryDruidHavingDimensionSelectorFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingDimensionSelectorFilter = result
			return nil
		case "equalTo":
			var result *TelemetryDruidHavingNumericFilter = &TelemetryDruidHavingNumericFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingNumericFilter = result
			return nil
		case "greaterThan":
			var result *TelemetryDruidHavingNumericFilter = &TelemetryDruidHavingNumericFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingNumericFilter = result
			return nil
		case "having":
			var result *TelemetryDruidHavingQueryFilter = &TelemetryDruidHavingQueryFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingQueryFilter = result
			return nil
		case "lessThan":
			var result *TelemetryDruidHavingNumericFilter = &TelemetryDruidHavingNumericFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingNumericFilter = result
			return nil
		case "telemetry.DruidHavingDimensionSelectorFilter":
			var result *TelemetryDruidHavingDimensionSelectorFilter = &TelemetryDruidHavingDimensionSelectorFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingDimensionSelectorFilter = result
			return nil
		case "telemetry.DruidHavingNumericFilter":
			var result *TelemetryDruidHavingNumericFilter = &TelemetryDruidHavingNumericFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingNumericFilter = result
			return nil
		case "telemetry.DruidHavingQueryFilter":
			var result *TelemetryDruidHavingQueryFilter = &TelemetryDruidHavingQueryFilter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHavingQueryFilter = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'type' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *TelemetryDruidHavingFilter) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidHavingDimensionSelectorFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingDimensionSelectorFilter)
	}

	if src.TelemetryDruidHavingNumericFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingNumericFilter)
	}

	if src.TelemetryDruidHavingQueryFilter != nil {
		return json.Marshal(&src.TelemetryDruidHavingQueryFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *TelemetryDruidHavingFilter) MarshalYAML() ([]byte, error) {
	if src.TelemetryDruidHavingDimensionSelectorFilter != nil {
		return yaml.Marshal(&src.TelemetryDruidHavingDimensionSelectorFilter)
	}

	if src.TelemetryDruidHavingNumericFilter != nil {
		return yaml.Marshal(&src.TelemetryDruidHavingNumericFilter)
	}

	if src.TelemetryDruidHavingQueryFilter != nil {
		return yaml.Marshal(&src.TelemetryDruidHavingQueryFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidHavingFilter) GetActualInstance() interface{} {
	if obj.TelemetryDruidHavingDimensionSelectorFilter != nil {
		return obj.TelemetryDruidHavingDimensionSelectorFilter
	}

	if obj.TelemetryDruidHavingNumericFilter != nil {
		return obj.TelemetryDruidHavingNumericFilter
	}

	if obj.TelemetryDruidHavingQueryFilter != nil {
		return obj.TelemetryDruidHavingQueryFilter
	}

	// all schemas are nil
	return nil
}
