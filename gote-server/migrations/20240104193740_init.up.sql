-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS note (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"title" varchar(50) NOT NULL,
	"description" varchar(300),
	"content" varchar NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tag (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS note_tags (
	"note_id" uuid,
	"tag_id"  uuid,
	PRIMARY KEY(note_id, tag_id),
	CONSTRAINT fk_note FOREIGN KEY(note_id) REFERENCES note ("id"), 
	CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tag("id") 
);
