package db

import memdb "github.com/hashicorp/go-memdb"

func InitDB() (*memdb.MemDB, error) { // Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	// Create a new data base
	return memdb.NewMemDB(schema)

}
