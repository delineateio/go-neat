ARG VARIANT="buster"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

ENV INSTALL_DIR /installs

# installs ------------------------------------------------------------------------

RUN apt update && apt -y install gcc musl-dev acl

# go-task -------------------------------------------------------------------------

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b $INSTALL_DIR
RUN ln -sf $INSTALL_DIR/task /usr/bin/task

# init ----------------------------------------------------------------------------

RUN echo "task init" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc
