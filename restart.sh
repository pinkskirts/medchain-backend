#!/bin/bash

podman stop API

podman rm API

podman image rm api

podman build -t=api .

podman run -dt -p 2005:2000 --cpus=256 -m=1024m --name=API -e  api