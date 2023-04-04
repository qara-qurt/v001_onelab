run:
	docker run -dp 8080:8080 --name sad_gould  --rm app

stop:
	docker stop sad_gould