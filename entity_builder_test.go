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
	spaceName := "test_space_upsert"
	sessionPool, err := newSessionPool(spaceName)
	assert.Nil(t, err)
	defer sessionPool.Close()

	schemaManager := nebula.NewSchemaManager(sessionPool)

	_, err = schemaManager.ApplyTag(testTagSchema01)
	assert.Nil(t, err)

	_, err = schemaManager.ApplyEdge(testEdgeSchema01)
	assert.Nil(t, err)

	// waiting for the schema to be propagated
	time.Sleep(5 * time.Second)

	b := NewEntityBuilder("user").
		SetProp("name", "Bob").
		SetProp("age", 18).
		UpsertVertex("test1")
	assert.Equal(t, `UPSERT VERTEX ON user "test1" SET name = "Bob", age = 18;`, b.String())

	_, err = b.Exec(sessionPool)
	assert.Nil(t, err)

	_, err = NewEntityBuilder("user").
		SetProp("name", `Lily "Double quotation"`).
		SetProp("age", 20).
		UpsertVertex("test2").
		Exec(sessionPool)
	assert.Nil(t, err)

	now := time.Now().Unix()
	b = NewEntityBuilder("friend").
		SetProp("created_at", now).
		UpsertEdge("test1", "test2")
	assert.Equal(t, fmt.Sprintf(`UPSERT EDGE ON friend "test1" -> "test2" SET created_at = %d;`, now), b.String())

	_, err = b.Exec(sessionPool)
	assert.Nil(t, err)

	rs, err := sessionPool.ExecuteAndCheck(`FETCH PROP ON friend "test1" -> "test2" YIELD edge AS e;`)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(rs.GetRows()))
}

func TestInsert(t *testing.T) {
	spaceName := "test_space_insert"
	sessionPool, err := newSessionPool(spaceName)
	assert.Nil(t, err)
	defer sessionPool.Close()

	schemaManager := nebula.NewSchemaManager(sessionPool)

	_, err = schemaManager.ApplyTag(testTagSchema01)
	assert.Nil(t, err)

	_, err = schemaManager.ApplyEdge(testEdgeSchema01)
	assert.Nil(t, err)

	// waiting for the schema to be propagated
	time.Sleep(5 * time.Second)

	b := NewEntityBuilder("user").
		InsertVertex("test0")
	assert.Equal(t, `INSERT VERTEX user () VALUES "test0":();`, b.String())

	b = NewEntityBuilder("user").
		SetProp("name", "Bob").
		SetProp("age", 18).
		InsertVertex("test1")
	assert.Equal(t, `INSERT VERTEX user (name, age) VALUES "test1":("Bob", 18);`, b.String())

	_, err = b.Exec(sessionPool)
	assert.Nil(t, err)

	_, err = NewEntityBuilder("user").
		SetProp("name", `Lily "Double quotation"`).
		SetProp("age", 20).
		InsertVertex("test2").
		Exec(sessionPool)
	assert.Nil(t, err)

	now := time.Now().Unix()
	b = NewEntityBuilder("friend").
		SetProp("created_at", now).
		InsertEdge("test1", "test2")
	assert.Equal(t, fmt.Sprintf(`INSERT EDGE friend (created_at) VALUES "test1" -> "test2":(%d);`, now), b.String())

	_, err = b.Exec(sessionPool)
	assert.Nil(t, err)

	rs, err := sessionPool.ExecuteAndCheck(`FETCH PROP ON friend "test1" -> "test2" YIELD edge AS e;`)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(rs.GetRows()))
}
