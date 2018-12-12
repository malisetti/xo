// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/JLightning/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ForeignServerOption represents a row from 'information_schema.foreign_server_options'.
type ForeignServerOption struct {
	ForeignServerCatalog pgtypes.SQLIdentifier `json:"foreign_server_catalog"` // foreign_server_catalog
	ForeignServerName    pgtypes.SQLIdentifier `json:"foreign_server_name"`    // foreign_server_name
	OptionName           pgtypes.SQLIdentifier `json:"option_name"`            // option_name
	OptionValue          pgtypes.CharacterData `json:"option_value"`           // option_value
}
