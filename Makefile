IMAGE:=mamedvedkov/reply_bot
VERSION:=v1.0.0

docker-build:
	@docker build --tag ${IMAGE}:${VERSION} .
	@docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest
docker-push:
	@docker push ${IMAGE}:${VERSION}
	@docker push ${IMAGE}:latest
