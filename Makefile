migrate:
	scripts/migrate_db.sh

rollback:
	scripts/migrate_db.sh down

recreate_db:
	dropdb --if-exists ${DATABASE_NAME}
	createdb ${DATABASE_NAME}
	make migrate

debug: FORCE
	dlv debug main.go

install_tools:
	scripts/install_tools.sh

FORCE: ;
