CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE nodes (
	node_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
	name TEXT NOT NULL,
	type TEXT CHECK (type IN ('country', 'state', 'city', 'region', 'suburb')),
	parent_node_id UUID REFERENCES nodes (node_id),
	latitude REAL,
	longitude REAL,
	has_seaport BOOLEAN DEFAULT FALSE,
	has_airport BOOLEAN DEFAULT FALSE,
	capacity INTEGER,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE node_metadata (node_id UUID PRIMARY KEY REFERENCES nodes (node_id), metadata JSONB);

CREATE TABLE edges (
    edge_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_node_id UUID REFERENCES nodes(node_id),
    to_node_id UUID REFERENCES nodes(node_id),
    distance_km FLOAT,
    travel_time_sec INTEGER,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
