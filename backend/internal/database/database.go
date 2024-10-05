package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type Client struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	IP              string `json:"ip"`
	WireGuardConfig string `json:"wireguard_config"`
	V2RayConfig     string `json:"v2ray_config"`
}

func Connect() (*DB, error) {
	connStr := "user=postgres dbname=hasbowleh sslmode=disable password=your_password"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) GetClients() ([]Client, error) {
	rows, err := db.Query("SELECT id, username, ip FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Username, &c.IP); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}

	return clients, nil
}

func (db *DB) AddClient(client *Client) error {
	_, err := db.Exec("INSERT INTO clients (username, ip, wireguard_config, v2ray_config) VALUES ($1, $2, $3, $4)",
		client.Username, client.IP, client.WireGuardConfig, client.V2RayConfig)
	return err
}

func (db *DB) DeleteClient(id string) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = $1", id)
	return err
}

func (db *DB) GetClient(id string) (*Client, error) {
	var c Client
	err := db.QueryRow("SELECT id, username, ip, wireguard_config, v2ray_config FROM clients WHERE id = $1", id).
		Scan(&c.ID, &c.Username, &c.IP, &c.WireGuardConfig, &c.V2RayConfig)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
