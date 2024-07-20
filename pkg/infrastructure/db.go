package infrastructure

type DB struct {
	store map[string]interface{}
}

func newDB() *DB {
	return &DB{
		store: make(map[string]interface{}),
	}
}

func (db *DB) Save(id string, value interface{}) {
	db.store[id] = value
}

func (db *DB) Get(id string) interface{} {
	return db.store[id]
}

func (db *DB) GetAll() map[string]interface{} {
	return db.store
}
