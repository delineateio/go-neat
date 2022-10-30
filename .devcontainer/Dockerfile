ARG VARIANT="buster"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# init ----------------------------------------------------------------------------
RUN echo "source /home/vscode/.devcontainer.env" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc
RUN echo "task init" | tee -a /home/vscode/.bashrc >> /home/vscode/.zshrc
