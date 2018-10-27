TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

pack:
	docker build -t wangzuo/avatar:$(TAG) .

upload:
	docker push wangzuo/avatar:$(TAG)

deploy:
	envsubst < k8s/deployment.yaml | kubectl apply -f -

ship: pack upload deploy
