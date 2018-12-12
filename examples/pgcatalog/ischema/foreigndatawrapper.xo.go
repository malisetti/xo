// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/JLightning/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ForeignDataWrapper represents a row from 'information_schema.foreign_data_wrappers'.
type ForeignDataWrapper struct {
	ForeignDataWrapperCatalog  pgtypes.SQLIdentifier `json:"foreign_data_wrapper_catalog"`  // foreign_data_wrapper_catalog
	ForeignDataWrapperName     pgtypes.SQLIdentifier `json:"foreign_data_wrapper_name"`     // foreign_data_wrapper_name
	AuthorizationIdentifier    pgtypes.SQLIdentifier `json:"authorization_identifier"`      // authorization_identifier
	LibraryName                pgtypes.CharacterData `json:"library_name"`                  // library_name
	ForeignDataWrapperLanguage pgtypes.CharacterData `json:"foreign_data_wrapper_language"` // foreign_data_wrapper_language
}
