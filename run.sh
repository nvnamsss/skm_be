#/bin/bash
echo "STOP RUNNING API CONTAINER"
docker stop -t 30 skm-service-container 
docker rm -f skm-service-container

echo "DONE STOPPING"
docker run --name skm-service-container -d \
            --network common-net \
            --env-file .dockerenv \
            -p 8080:8080 \
            skm:latest
