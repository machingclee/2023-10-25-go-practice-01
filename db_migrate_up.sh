DB_URL=postgresql://pguser:pguser@127.0.0.1:5432/rssagg

cd sql/schema
goose postgres $DB_URL up