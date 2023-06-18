package conn

import bolt "go.etcd.io/bbolt"

func InitBolt() (*bolt.DB, error) {
	db, err := bolt.Open("data/bolt.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
