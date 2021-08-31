TEMP_IMAGE_IIDFILE=/tmp/librootfs.tempfile
# DOCKER_BUILD_MODE="--quiet"

# build an executable named myprog from myprog.c
all: build
# enable optimzation
build: 
	go build -o librootfs.so -buildmode=c-shared rootfs_proxy.go main.go

docker-build: 
	docker build $(DOCKER_BUILD_MODE) -t build-librootfs --iidfile $(TEMP_IMAGE_IIDFILE) .
	$(eval IMAGE_NAME := librootfs-$$(shell cat $$(TEMP_IMAGE_IIDFILE)|cut -c8-))
	docker rm -f $(IMAGE_NAME)
	docker run -d --name $(IMAGE_NAME) build-librootfs
	docker cp $(IMAGE_NAME):/src/hcsshim/internal/tools/rootfs/librootfs.so .
	docker cp $(IMAGE_NAME):/src/hcsshim/internal/tools/rootfs/librootfs.h .
	docker rm -f $(IMAGE_NAME)
	$(RM) $(TEMP_IMAGE_IIDFILE)

run: 
	@time python3 rootfs_proxy.py

clean: 
	$(RM) librootfs.h librootfs.so librootfs.a $(TEMP_IMAGE_IIDFILE)