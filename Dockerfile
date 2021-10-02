FROM golang:1.17

RUN apt-get update \
    && apt-get install --no-install-recommends -y \
        curl \
    && go get github.com/ilyubin/gotest2allure/cmd/gotest2allure \
    && wget -q -nv "https://github.com/allure-framework/allurectl/releases/download/1.19.3/allurectl_linux_amd64" \
    -O /usr/local/bin/allurectl \
    && chmod +x /usr/local/bin/allurectl || exit 11