package repository

import sq "github.com/Masterminds/squirrel"

// psql — общий билдер для Postgres с плейсхолдерами $1..$n
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
