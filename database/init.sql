CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    ip VARCHAR(15) NOT NULL,
    wireguard_config TEXT NOT NULL,
    v2ray_config TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_active TIMESTAMP
);

CREATE INDEX idx_username ON clients(username);
CREATE INDEX idx_ip ON clients(ip);