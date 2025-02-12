-- +goose Up
-- +goose StatementBegin
CREATE TABLE torrents_metadata (
    id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    bytes UINTEGER NOT NULL,
    description STRING NOT NULL DEFAULT '',
    infohash STRING NOT NULL,
    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS torrents_metadata;
-- +goose StatementEnd
