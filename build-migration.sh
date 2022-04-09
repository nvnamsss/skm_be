#/bin/bash
image_name=skm-migration
echo "REMOVING OLD IMAGE"
docker image rm -f $image_name
echo "BUILDING NEW IMAGE"
docker build . -f Dockerfile-migration -t $image_name