docker run --name skm-migration-container --rm \
            --network common-net \
            --env-file .dockerenv \
            skm-migration:latest
