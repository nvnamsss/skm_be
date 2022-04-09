#/bin/bash
image_name=skm
echo "REMOVING OLD IMAGE"
docker image rm -f $image_name
echo "BUILDING NEW IMAGE"
docker build . -t $image_name