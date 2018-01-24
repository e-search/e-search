REPO=github.com/e-search/e-search
LOCALDB=e-search

local-migrate:
	cat sql/*.sql | psql -d $(LOCALDB)
