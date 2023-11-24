# Accelerator Log

## Options
```json
{
  "bsGitBranch" : "main",
  "projectName" : "tanzu-accelerator-go-hellodb-main",
  "dbClassClaim" : "db-pgsql-bitnami",
  "goVersion" : "1.20",
  "bsGitRepository" : "github.com?owner=zahooruk2022&repo=tanzu-go-hellodb"
}
```
## Log
```
┏ engine (Chain)
┃  Info Running Chain(GeneratorValidationTransform, UniquePath)
┃ ┏ ┏ engine.transformations[0].validated (Combo)
┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ engine.transformations[0].validated.delegate (Chain)
┃ ┃ ┃  Info Running Chain(Merge, UniquePath)
┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0] (Merge)
┃ ┃ ┃ ┃  Info Running Merge(Combo, Combo, Combo, Combo, Combo, Combo)
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[0] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[0].delegate (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, Exclude)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[0].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**/*]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[0].delegate.transformations[1] (Exclude)
┃ ┃ ┃ ┃ ┃ ┃  Info Will exclude [config/*.yaml, go.mod, .*, **/*.go, README.md]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go matched [config/*.yaml, go.mod, .*, **/*.go, README.md] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┗ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [config/*.yaml, go.mod, .*, **/*.go, README.md] -> included
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[1] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[1].delegate (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[1].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃  Info Will include [catalog/*.yaml]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [catalog/*.yaml] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[1].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗  Info Will replace [tanzu-go-hellodb->tanzu-accelerator-go...(truncated)]
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[1].delegate.transformations[2] (ReplaceText)
┃ ┃ ┃ ┃ ┗ ┗  Info Will replace [{{CLASS_CLAIM}}->db-pgsql-bitnami]
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[2] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[2].delegate (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[2].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**/*.go]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go matched [**/*.go] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [**/*.go] -> excluded
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[2].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┗ ┗  Info Will replace [tanzu-go-hellodb->tanzu-accelerator-go...(truncated)]
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText, ReplaceText, Combo)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃  Info Will include [config/*.yaml]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [config/*.yaml] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [config/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗  Info Will replace [tanzu-go-hellodb->tanzu-accelerator-go...(truncated)]
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[2] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗  Info Will replace [{{CLASS_CLAIM}}->db-pgsql-bitnami]
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3] (Combo)
┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Merge, UniquePath)
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0] (Merge)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Running Merge(InvokeFragment, Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0] (InvokeFragment)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0].validated (Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Condition (#bsGitRepository != null) evaluated to true
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Let
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0].validated.delegate (Let)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Adding symbol repoUrl with value 'https://github.com?owner=zahooruk2022&repo=tanzu-go-hellodb'
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0].validated.delegate.in (Chain)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(OpenRewriteRecipe, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ╺ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0].validated.delegate.in.transformations[0] (OpenRewriteRecipe)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[0].validated.delegate.in.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┗ ┗ ┗ ┗  Info Will replace regex '(?<beforeBranch>[\s\S]+)(?<branch>branch: [\S]+)(?<rest>[\S\s]*)' with '${beforeBranch}branc...(truncated)'
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[1] (Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Include
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[0].sources[1].delegate (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**]
┃ ┃ ┃ ┃ ┃ ┃ ┗ ┗ Debug config/workload.yaml matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[3].delegate.transformations[3].delegate.transformations[1] (UniquePath)
┃ ┃ ┃ ┃ ┗ ┗ ┗ Debug Multiple representations for path 'config/workload.yaml', will use the one appearing first 
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[4] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[4].delegate (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[4].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃  Info Will include [go.mod]
┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug go.mod matched [go.mod] -> included
┃ ┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/visits.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[4].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗  Info Will replace regex '1\.20' with '1.20'
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[4].delegate.transformations[2] (ReplaceText)
┃ ┃ ┃ ┃ ┗ ┗  Info Will replace regex 'tanzu-go-hellodb' with 'tanzu-accelerator-go...(truncated)'
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[0].sources[5] (Combo)
┃ ┃ ┃ ┃ ┃  Info Combo running as Include
┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.transformations[0].sources[5].delegate (Include)
┃ ┃ ┃ ┃ ┃  Info Will include [README.md]
┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug README.md matched [README.md] -> included
┃ ┃ ┃ ┃ ┃ Debug go.mod didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug visits/visits.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.down.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug visits/migrations/1_initial_schema.up.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.codecov.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.gitignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/.travis.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/CHANGELOG.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/LICENSE.txt didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/Makefile didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/bool_ext.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/duration_ext.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/error_ext.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/float64_ext.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/gen.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int32.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/int64.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/nocmp.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/string_ext.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint32.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/uint64.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/go.uber.org/atomic/value.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/errwrap/errwrap.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/Makefile didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/append.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/flatten.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/format.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/group.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/multierror.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/prefix.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/hashicorp/go-multierror/sort.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/.travis.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgpassfile/pgpass.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/.travis.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgservicefile/pgservicefile.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.gitignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.sh didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/.travis.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/LICENSE.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/TESTS.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/array.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/buf.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/conn_go18.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/connector.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/copy.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/encode.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/error.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/krb.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notice.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/notify.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/rows.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_permissions.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/ssl_windows.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/url.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_other.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_posix.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/user_windows.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/uuid.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/PATENTS didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/PATENTS didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.dockerignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.gitignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.golangci.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.goreleaser.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/.travis.yml didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/CONTRIBUTING.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.circleci didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Dockerfile.github-actions didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/FAQ.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/GETTING_STARTED.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/MIGRATIONS.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/Makefile didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/SECURITY.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/docker-deploy.sh didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/log.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migrate.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/migration.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/util.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/.gitignore didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CHANGELOG.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/CONTRIBUTING.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/LICENSE didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/Rakefile didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/batch.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/conn.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/copy_from.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/extended_query_builder.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/large_objects.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/named_args.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/rows.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tracer.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/tx.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/values.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/oid/types.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/lib/pq/scram/scram.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/crypto/pbkdf2/pbkdf2.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/cases.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/context.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/fold.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/icu.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/info.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/map.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables11.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables12.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables13.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/tables9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/cases/trieval.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/internal.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/match.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/coverage.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/language.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/match.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/parse.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tables.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/language/tags.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/cond.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/runes/runes.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/transform/transform.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/kind_string.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables11.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables12.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables13.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/tables9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/transform.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/trieval.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/width/width.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/driver.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/error.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/util.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/driver.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/errors.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/migration.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/parse.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/auth_scram.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/config.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/defaults_windows.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/errors.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/krb5.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/pgconn.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_cleartext_password.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_gss_continue.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_md5_password.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_ok.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_continue.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/authentication_sasl_final.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/backend_key_data.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/big_endian.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/bind_complete.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/cancel_request.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/chunkreader.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/close_complete.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/command_complete.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_both_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_data.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_done.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_fail.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_in_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/copy_out_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/data_row.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/describe.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/empty_query_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/error_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/execute.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/flush.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/frontend.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/function_call_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_enc_request.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/gss_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/no_data.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notice_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/notification_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_description.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parameter_status.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/parse_complete.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/password_message.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/pgproto3.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/portal_suspended.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/query.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ready_for_query.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/row_description.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_initial_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sasl_response.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/ssl_request.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/startup_message.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/sync.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/terminate.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgproto3/trace.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/array_codec.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bits.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bool.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/box.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/builtin_wrappers.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/bytea.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/circle.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/composite.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/convert.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/date.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/enum_codec.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float4.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/float8.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/hstore.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/inet.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int.go.erb didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/int_test.go.erb didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test.go.erb didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/integration_benchmark_test_gen.sh didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/interval.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/json.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/jsonb.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/line.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/lseg.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/macaddr.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/multirange.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/numeric.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/path.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/pgtype.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/point.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/polygon.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/qchar.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/range_codec.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/record_codec.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/register_default_pg_types_disabled.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/text_format_only_codec.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/tid.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/time.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamp.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/timestamptz.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uint32.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgtype/uuid.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/stdlib/sql.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/common.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compose.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/coverage.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/language.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/lookup.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/match.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/parse.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tables.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/tags.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/tag/tag.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/bidirule/bidirule9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/class.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/context.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/nickname.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/options.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profile.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/profiles.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables11.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables12.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables13.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/tables9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/transformer.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/secure/precis/trieval.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bidi.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/bracket.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/core.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/prop.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables11.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables12.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables13.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/tables9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/bidi/trieval.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/composition.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/forminfo.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/input.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/iter.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/normalize.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/readwriter.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables10.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables11.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables12.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables13.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/tables9.0.0.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/transform.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/unicode/norm/trie.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/multistmt/parse.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/TUTORIAL.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/database/postgres/postgres.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/internal/url/url.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/iofs.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/anynil/anynil.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/iobufpool/iobufpool.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/bufferqueue.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_fake_non_block.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/nbconn/nbconn_real_non_block.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/README.md didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/doc.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/pgio/write.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/sanitize/sanitize.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/lru_cache.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/stmtcache.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/internal/stmtcache/unlimited_cache.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/compact.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/language.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/parents.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tables.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/golang.org/x/text/internal/language/compact/tags.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/jackc/pgx/v5/pgconn/internal/ctxwatch/context_watcher.go didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.down.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/1_foobar.up.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/3_foobar.up.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.down.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/4_foobar.up.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/5_foobar.down.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.down.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┗ ┗ Debug vendor/github.com/golang-migrate/migrate/v4/source/iofs/testdata/migrations/7_foobar.up.sql didn't match [README.md] -> excluded
┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.transformations[1] (UniquePath)
┃ ┃ ┃ ┃ Debug Multiple representations for path 'catalog/catalog-info.yaml', will use the one appearing last 
┃ ┗ ┗ ┗ Debug Multiple representations for path 'README.md', will use the one appearing last 
┗ ╺ engine.transformations[1] (UniquePath)
```
