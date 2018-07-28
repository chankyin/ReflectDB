/*
 * reflectdb
 *
 * Copyright (C) 2018 chankyin
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package reflectdb_json

import (
	"github.com/chankyin/reflectdb/reflectdb-go/schema"
	"time"
)

type Json struct {
	// internal reflectdb structure version
	ReflectdbVersion schema.ReflectdbVersion
	// Time of creation of this database
	Creation time.Time

	// A schema version number provided by the user. Should increment every time the version is changed.
	// reflectdb will try to detect if the user version and the database version are compatible.
	SchemaVersion uint64
	// Time of last update of the schema
	Updated time.Time

	// A sha512 hash of the schema as a safe measure to validate the schema version.
	SchemaHash [64]byte

	// The schema content
	Schema struct {
		// The class types in the schema.
		Classes []Class
		// The functions that the connector should provide.
		Functions []Function
	}
}

type Class struct {
	Name    string
	Fields  []Field
	Indices []Index
}

type Field struct {
	Name string
	Type FieldType
}

type FieldType struct {
	Value        BaseType
	LengthBounds *FieldTypeLengthBound
	LengthFixed  *FieldTypeLengthFixed
	Collection   *FieldTypeCollection
	Map          *FieldTypeMap
	Object       *FieldTypeObject
}

type FieldTypeLengthBound struct {
	Min uint16
	Max uint16
}
type FieldTypeLengthFixed struct {
	Length uint16
}
type FieldTypeCollection struct {
	CollectionType CollectionType
}
type FieldTypeMap struct {
	KeyType   FieldType
	ValueType FieldType
}
type FieldTypeObject struct {
	Class string
}

type BaseType uint16

const (
	BaseTypeBoolean BaseType = iota
	BaseTypeInt8
	BaseTypeInt16
	BaseTypeInt32
	BaseTypeInt64
	BaseTypeUint8
	BaseTypeUint16
	BaseTypeUint32
	BaseTypeUint64
	BaseTypeFloat32
	BaseTypeFloat64
	BaseTypeFixedString
	BaseTypeVarString

	BaseTypeFixedArray
	BaseTypeCollection
	BaseTypeMap

	BaseTypeObject
)

type CollectionType uint16

const (
	CollectionTypeSet CollectionType = iota
	CollectionTypeList
)

type Index struct {
	Fields []string
}

// Functions are defined to be "pure operations",
// i.e. they should not have side effects other than creating output from the input.
// However, the function is still allowed to have aggregate-local or subquery-local memory,
// e.g. to assign an incremental ID to each row in an aggregate.
// If the operations are impure, i.e. they would change the data on the database, a Procedure should be used instead.
type Function struct {
	Name    string
	Inputs  []FunctionInput
	Outputs []FunctionOutput
}

type FunctionInput struct {
	Name string
}

type FunctionOutput struct {
	Name string
}
