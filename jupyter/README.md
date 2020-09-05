```yaml
version: "3.4"
services:
    jupyter:
        restart: always
        build: ./tools/jupyter
        ports:
          - 8888:8888
        volumes:
          - ./tools/jupyter/volume/jupyter:/opt/python/jupyter
          - ./tools/jupyter/volume/jupyterlab:/opt/python/jupyterlab
          - ./tools/jupyter/volume/library:/opt/python/library
```


Password: `password`
