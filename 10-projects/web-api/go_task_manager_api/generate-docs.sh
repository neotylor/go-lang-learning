#!/bin/bash

echo "Generating Swagger docs..."
swag init -g main.go
echo "Docs generated successfully!"


# Run this script
# bash generate-docs.sh