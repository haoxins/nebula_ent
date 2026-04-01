package ent

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	nebula "github.com/vesoft-inc/nebula-go/v3"
)

var testTagSchema01 = nebula.LabelSchema{
	Name: "user",
	Fields: []nebula.LabelFieldSchema{
		{
			Field:    "name",
			Type:     "string",
			Nullable: false,
		},
		{
			Field:    "age",
			Type:     "int",
			Nullable: true,
		},
	},
}

var testEdgeSchema01 = nebula.LabelSchema{
	Name: "friend",
	Fields: []nebula.LabelFieldSchema{
		{
			Field:    "created_at",
			Type:     "int64",
			Nullable: true,
		},
	},
}

func TestUpsert(t *testing.T) {
	b := NewEntityBuilder("user").
		SetProp("name", "Bob").
		SetProp("age", 18).
		UpsertVertex("test1")
	assert.Equal(t, `UPSERT VERTEX ON user "test1" SET name = "Bob", age = 18;`, b.String())

	now := time.Now().Unix()
	b = NewEntityBuilder("friend").
		SetProp("created_at", now).
		UpsertEdge("test1", "test2")
	assert.Equal(t, fmt.Sprintf(`UPSERT EDGE ON friend "test1" -> "test2" SET created_at = %d;`, now), b.String())
}

func TestInsert(t *testing.T) {
	b := NewEntityBuilder("user").
		InsertVertex("test0")
	assert.Equal(t, `INSERT VERTEX user () VALUES "test0":();`, b.String())

	b = NewEntityBuilder("user").
		SetProp("name", "Bob").
		SetProp("age", 18).
		InsertVertex("test1")
	assert.Equal(t, `INSERT VERTEX user (name, age) VALUES "test1":("Bob", 18);`, b.String())

	now := time.Now().Unix()
	b = NewEntityBuilder("friend").
		SetProp("created_at", now).
		InsertEdge("test1", "test2")
	assert.Equal(t, fmt.Sprintf(`INSERT EDGE friend (created_at) VALUES "test1" -> "test2":(%d);`, now), b.String())
}
