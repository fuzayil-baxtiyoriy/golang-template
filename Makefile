.PHONY: run
run:
	go run main.go


.PHONY: devup
devup:
	docker-compose -f compose-dev.yaml up -d


.PHONY: devdown
devdown:
	docker-compose -f compose-dev.yaml down


.PHONY: devlogs
devlogs:
	docker-compose -f compose-dev.yaml logs -f --tail=150 ${name}