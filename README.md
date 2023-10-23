# rssagg
Experiment with go


- `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- `go install github.com/pressly/goose/v3/cmd/goose@latest`
- For `Windows`, instead of `sqlc generate` we run 
  - `docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate` in `cmd`