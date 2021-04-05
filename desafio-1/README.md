# Instructions

#### **Imagem:**
elbesh/imersao-fsfc2-desafio-1

#### **Docker Build**

```
docker build \
    -f Dockerfile.prod \
    -t elbesh/imersao-fsfc2-desafio-1 \
    .
```

#### **Docker Run (Production)**

```
docker run \
    --name desafio-1 \
    -p 8000:8080 \
    -d \
    --rm \
    elbesh/imersao-fsfc2-desafio-1
```

#### **Docker Compose (Dev)**

```
docker-compose up -d
```