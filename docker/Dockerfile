From golang:1.14.2
RUN set -eux; \
        apt-get update && apt-get install vim -y && apt-get install git -y

COPY client-go /go/src/k8s.io/client-go
RUN set -eux; \
				echo "192.30.253.112 github.com" >> /etc/hosts && \
				echo "151.101.185.194 github.global.ssl.fastly.net" >> /etc/hosts && \
				echo "export GOPROXY=https://goproxy.cn,direct" >> /root/.bash_profile && \
				echo "export GO111MODULE=on" >> /root/.bash_profile && \
				/bin/bash -c "source /root/.bash_profile" && \
				mkdir -p /go/src/k8s.io && \
				cd /go/src/k8s.io/client-go && \
        go mod tidy 

COPY Entrypoint.sh /usr/local/bin/
ENTRYPOINT ["Entrypoint.sh"]
CMD ["sleep", "360000"]