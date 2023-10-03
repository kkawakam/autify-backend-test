# Usage
Building the application
```bash
docker build -t katsutoshi-kawakami-autify-backend-test .
```

Running the application
```bash
docker run -v fetcher-output:/output katsutoshi-kawakami-autify-backend-test --metadata https://autify.com
```

You can look at the output html by going to the appropriate volume directory.
Each run of the container will generate a new directory suffixed by a unix timestamp
```bash
# You will most likely need to be root
$ cd /var/lib/docker/volumes/fetcher-output/_data
$ tree
.
├── results-1696374426918
│   └── autify.com.html
└── results-1696374681557
    └── autify.com.html
```