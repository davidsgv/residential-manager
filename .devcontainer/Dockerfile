FROM golang:1.21.6-bullseye

#Herramientas de desarrollo
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
#migration tool 
# https://github.com/golang-migrate/migrate
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

#Intall nodejs
# https://docs.docker.com/engine/reference/builder/#shell
# https://copyprogramming.com/howto/how-to-install-nvm-in-docker
# SHELL ["/bin/bash", "--login", "-c"]
# RUN apt-get update && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash && source ~/.profile && nvm install 21.4.0 && npm install --global yarn

# install git
# RUN apt update && apt upgrade
# RUN apk add git

#config git user and email
# RUN git config --global user.name "${git_user}"
# RUN git config --global user.email "${git_mail}"


# EXPOSE 80