module rss

go 1.22.4

require github.com/uzairkhan98/rss/config v0.0.0

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)

replace github.com/uzairkhan98/rss/config v0.0.0 => ./internal/config

replace github.com/uzairkhan98/rss/database v0.0.0 => ./internal/database
