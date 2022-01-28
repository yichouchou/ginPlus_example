FROM 759960474/ginplus:second_with_dlv

COPY main /home/admin/main

# dlv 目录
#/root/go/bin/dlv
#
#
#GO111MODULE=""
#GOARCH="amd64"
#GOBIN=""
#GOCACHE="/root/.cache/go-build"
#GOENV="/root/.config/go/env"
#GOEXE=""
#GOEXPERIMENT=""
#GOFLAGS=""
#GOHOSTARCH="amd64"
#GOHOSTOS="linux"
#GOINSECURE=""
#GOMODCACHE="/root/go/pkg/mod"
#GONOPROXY=""
#GONOSUMDB=""
#GOOS="linux"
#GOPATH="/root/go"
#GOPRIVATE=""
#GOPROXY="https://goproxy.cn,direct"
#GOROOT="/usr/local/go"
#GOSUMDB="sum.golang.org"
#GOTMPDIR=""
#GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
#GOVCS=""
#GOVERSION="go1.17.5"
#GCCGO="gccgo"
#AR="ar"
#CC="gcc"
#CXX="g++"
#CGO_ENABLED="1"
#GOMOD="/dev/null"
#CGO_CFLAGS="-g -O2"
#CGO_CPPFLAGS=""
#CGO_CXXFLAGS="-g -O2"
#CGO_FFLAGS="-g -O2"
#CGO_LDFLAGS="-g -O2"
#PKG_CONFIG="pkg-config"


RUN  mkdir -p /home/admin/mainlogs/ && chmod 0777 /home/admin/mainlogs/

#CMD  /home/admin/main >> /home/admin/mainlogs/out.log 2>&1


# 查找gunPlus 运行的pid

#ps -ef |grep zqkd_rec|grep "/home/admin/logs/out.log"|awk  '{print $2}'
#
#ps -ef |grep "/home/admin/logs/out.log"|awk  '{print $2}'
#
## 终极版本，剔除了grep影响
#ps -ef |grep -v grep|grep "/home/admin/mainlogs"|awk  '{print $2}'

#RUN /root/go/bin/dlv --listen=:2345 --headless=true --api-version=2 --log --accept-multiclient exec /home/admin/main

CMD /bin/bash