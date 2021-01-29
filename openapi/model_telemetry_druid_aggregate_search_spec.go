/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-08T20:53:20Z.
 *
 * API version: 1.0.9-2908
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

// TelemetryDruidAggregateSearchSpec - Generic search specification descriptor
type TelemetryDruidAggregateSearchSpec struct {
	TelemetryDruidContainsSearchSpec            *TelemetryDruidContainsSearchSpec
	TelemetryDruidFragmentSearchSpec            *TelemetryDruidFragmentSearchSpec
	TelemetryDruidInsensitiveContainsSearchSpec *TelemetryDruidInsensitiveContainsSearchSpec
	TelemetryDruidRegexSearchSpec               *TelemetryDruidRegexSearchSpec
}

// TelemetryDruidContainsSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidContainsSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidContainsSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidContainsSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidContainsSearchSpec: v}
}

// TelemetryDruidFragmentSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidFragmentSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidFragmentSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidFragmentSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidFragmentSearchSpec: v}
}

// TelemetryDruidInsensitiveContainsSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidInsensitiveContainsSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidInsensitiveContainsSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidInsensitiveContainsSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidInsensitiveContainsSearchSpec: v}
}

// TelemetryDruidRegexSearchSpecAsTelemetryDruidAggregateSearchSpec is a convenience function that returns TelemetryDruidRegexSearchSpec wrapped in TelemetryDruidAggregateSearchSpec
func TelemetryDruidRegexSearchSpecAsTelemetryDruidAggregateSearchSpec(v *TelemetryDruidRegexSearchSpec) TelemetryDruidAggregateSearchSpec {
	return TelemetryDruidAggregateSearchSpec{TelemetryDruidRegexSearchSpec: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *TelemetryDruidAggregateSearchSpec) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["type"]; ok {
		switch v {
		case "contains":
			var result *TelemetryDruidContainsSearchSpec = &TelemetryDruidContainsSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidContainsSearchSpec = result
			return nil
		case "fragment":
			var result *TelemetryDruidFragmentSearchSpec = &TelemetryDruidFragmentSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidFragmentSearchSpec = result
			return nil
		case "insensitive_contains":
			var result *TelemetryDruidInsensitiveContainsSearchSpec = &TelemetryDruidInsensitiveContainsSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidInsensitiveContainsSearchSpec = result
			return nil
		case "regex":
			var result *TelemetryDruidRegexSearchSpec = &TelemetryDruidRegexSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidRegexSearchSpec = result
			return nil
		case "telemetry.DruidContainsSearchSpec":
			var result *TelemetryDruidContainsSearchSpec = &TelemetryDruidContainsSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidContainsSearchSpec = result
			return nil
		case "telemetry.DruidFragmentSearchSpec":
			var result *TelemetryDruidFragmentSearchSpec = &TelemetryDruidFragmentSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidFragmentSearchSpec = result
			return nil
		case "telemetry.DruidInsensitiveContainsSearchSpec":
			var result *TelemetryDruidInsensitiveContainsSearchSpec = &TelemetryDruidInsensitiveContainsSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidInsensitiveContainsSearchSpec = result
			return nil
		case "telemetry.DruidRegexSearchSpec":
			var result *TelemetryDruidRegexSearchSpec = &TelemetryDruidRegexSearchSpec{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidRegexSearchSpec = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'type' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *TelemetryDruidAggregateSearchSpec) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidContainsSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidContainsSearchSpec)
	}

	if src.TelemetryDruidFragmentSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidFragmentSearchSpec)
	}

	if src.TelemetryDruidInsensitiveContainsSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidInsensitiveContainsSearchSpec)
	}

	if src.TelemetryDruidRegexSearchSpec != nil {
		return json.Marshal(&src.TelemetryDruidRegexSearchSpec)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *TelemetryDruidAggregateSearchSpec) MarshalYAML() ([]byte, error) {
	if src.TelemetryDruidContainsSearchSpec != nil {
		return yaml.Marshal(&src.TelemetryDruidContainsSearchSpec)
	}

	if src.TelemetryDruidFragmentSearchSpec != nil {
		return yaml.Marshal(&src.TelemetryDruidFragmentSearchSpec)
	}

	if src.TelemetryDruidInsensitiveContainsSearchSpec != nil {
		return yaml.Marshal(&src.TelemetryDruidInsensitiveContainsSearchSpec)
	}

	if src.TelemetryDruidRegexSearchSpec != nil {
		return yaml.Marshal(&src.TelemetryDruidRegexSearchSpec)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidAggregateSearchSpec) GetActualInstance() interface{} {
	if obj.TelemetryDruidContainsSearchSpec != nil {
		return obj.TelemetryDruidContainsSearchSpec
	}

	if obj.TelemetryDruidFragmentSearchSpec != nil {
		return obj.TelemetryDruidFragmentSearchSpec
	}

	if obj.TelemetryDruidInsensitiveContainsSearchSpec != nil {
		return obj.TelemetryDruidInsensitiveContainsSearchSpec
	}

	if obj.TelemetryDruidRegexSearchSpec != nil {
		return obj.TelemetryDruidRegexSearchSpec
	}

	// all schemas are nil
	return nil
}