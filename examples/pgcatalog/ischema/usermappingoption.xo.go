// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/JLightning/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// UserMappingOption represents a row from 'information_schema.user_mapping_options'.
type UserMappingOption struct {
	AuthorizationIdentifier pgtypes.SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
	ForeignServerCatalog    pgtypes.SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       pgtypes.SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
	OptionName              pgtypes.SQLIdentifier `json:"option_name"`              // option_name
	OptionValue             pgtypes.CharacterData `json:"option_value"`             // option_value
}
