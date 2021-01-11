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

// WorkflowCustomDataTypeDefinitionResponse - The response body of a HTTP GET request for the 'workflow.CustomDataTypeDefinition' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'workflow.CustomDataTypeDefinition' resources.
type WorkflowCustomDataTypeDefinitionResponse struct {
	MoAggregateTransform                 *MoAggregateTransform
	MoDocumentCount                      *MoDocumentCount
	MoTagSummary                         *MoTagSummary
	WorkflowCustomDataTypeDefinitionList *WorkflowCustomDataTypeDefinitionList
}

// MoAggregateTransformAsWorkflowCustomDataTypeDefinitionResponse is a convenience function that returns MoAggregateTransform wrapped in WorkflowCustomDataTypeDefinitionResponse
func MoAggregateTransformAsWorkflowCustomDataTypeDefinitionResponse(v *MoAggregateTransform) WorkflowCustomDataTypeDefinitionResponse {
	return WorkflowCustomDataTypeDefinitionResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsWorkflowCustomDataTypeDefinitionResponse is a convenience function that returns MoDocumentCount wrapped in WorkflowCustomDataTypeDefinitionResponse
func MoDocumentCountAsWorkflowCustomDataTypeDefinitionResponse(v *MoDocumentCount) WorkflowCustomDataTypeDefinitionResponse {
	return WorkflowCustomDataTypeDefinitionResponse{MoDocumentCount: v}
}

// MoTagSummaryAsWorkflowCustomDataTypeDefinitionResponse is a convenience function that returns MoTagSummary wrapped in WorkflowCustomDataTypeDefinitionResponse
func MoTagSummaryAsWorkflowCustomDataTypeDefinitionResponse(v *MoTagSummary) WorkflowCustomDataTypeDefinitionResponse {
	return WorkflowCustomDataTypeDefinitionResponse{MoTagSummary: v}
}

// WorkflowCustomDataTypeDefinitionListAsWorkflowCustomDataTypeDefinitionResponse is a convenience function that returns WorkflowCustomDataTypeDefinitionList wrapped in WorkflowCustomDataTypeDefinitionResponse
func WorkflowCustomDataTypeDefinitionListAsWorkflowCustomDataTypeDefinitionResponse(v *WorkflowCustomDataTypeDefinitionList) WorkflowCustomDataTypeDefinitionResponse {
	return WorkflowCustomDataTypeDefinitionResponse{WorkflowCustomDataTypeDefinitionList: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *WorkflowCustomDataTypeDefinitionResponse) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["ObjectType"]; ok {
		switch v {
		case "mo.AggregateTransform":
			var result *MoAggregateTransform = &MoAggregateTransform{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.MoAggregateTransform = result
			return nil
		case "mo.DocumentCount":
			var result *MoDocumentCount = &MoDocumentCount{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.MoDocumentCount = result
			return nil
		case "mo.TagSummary":
			var result *MoTagSummary = &MoTagSummary{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.MoTagSummary = result
			return nil
		case "workflow.CustomDataTypeDefinition.List":
			var result *WorkflowCustomDataTypeDefinitionList = &WorkflowCustomDataTypeDefinitionList{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.WorkflowCustomDataTypeDefinitionList = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'ObjectType' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'ObjectType' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *WorkflowCustomDataTypeDefinitionResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.WorkflowCustomDataTypeDefinitionList != nil {
		return json.Marshal(&src.WorkflowCustomDataTypeDefinitionList)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *WorkflowCustomDataTypeDefinitionResponse) MarshalYAML() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return yaml.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return yaml.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return yaml.Marshal(&src.MoTagSummary)
	}

	if src.WorkflowCustomDataTypeDefinitionList != nil {
		return yaml.Marshal(&src.WorkflowCustomDataTypeDefinitionList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WorkflowCustomDataTypeDefinitionResponse) GetActualInstance() interface{} {
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.WorkflowCustomDataTypeDefinitionList != nil {
		return obj.WorkflowCustomDataTypeDefinitionList
	}

	// all schemas are nil
	return nil
}
