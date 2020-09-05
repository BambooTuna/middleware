```yaml
version: "3.4"
services:
    jupyter:
        restart: always
        build: ./tools/jupyter
        ports:
          - 8888:8888
        volumes:
          - ./tools/jupyter/volume/tmp:/tmp
          - ./tools/jupyter/volume/jupyter:/opt/python/jupyter
          - ./tools/jupyter/volume/jupyterlab:/opt/python/jupyterlab
          - ./tools/jupyter/volume/library:/opt/python/library
```

```bash
$ docker build . -t jupyter
$ docker run \
    --restart=always -d \
    -v $PWD/volume/tmp:/tmp \
    -v $PWD/volume/jupyter:/opt/python/jupyter \
    -v $PWD/volume/jupyterlab:/opt/python/jupyterlab \
    -v $PWD/volume/library:/opt/python/library \
    -p 8888:8888 \
    jupyter
```

Password: `password`
