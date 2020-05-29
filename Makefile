setup: setup-buildx

setup-buildx:
	docker buildx create --name multi-arch-tilt
	docker buildx use multi-arch-tilt
	docker buildx inspect --bootstrap

