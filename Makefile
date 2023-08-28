# ==============================================================================

.PHONY: atlas-schema
atlas-schema:   ## 查看生成的SQL Schema
	atlas schema inspect \
  		-u "ent://ent/schema" \
  		--dev-url "sqlite://file?mode=memory&_fk=1" \
  		--format '{{ sql . "  " }}'


.PHONY: atlas-generate-migrate
atlas-generate-migrate:  ## 生成迁移
	atlas migrate diff migration_name \
  		--dir "file://ent/migrate/migrations" \
  		--to "ent://ent/schema" \
  		--dev-url "docker://mysql/8/ent"

.PHONY: atlas-apply-migrate
atlas-apply-migrate:	## 应用迁移
	atlas migrate apply \
  		--dir "file://ent/migrate/migrations" \
  		--url "mysql://root:admin123@localhost:3306/ent"