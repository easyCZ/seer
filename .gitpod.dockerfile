FROM gitpod/workspace-postgres

USER root

ENV GO111MODULE=on \
    LOCAL_BIN="/usr/bin" \
    BUF_VERSION="1.0.0-rc12"

# Install buf
RUN curl -SL "https://github.com/bufbuild/buf/releases/download/v{$BUF_VERSION}/buf-Linux-x86_64" \
      -o "${LOCAL_BIN}/buf" && \
    chmod +x "${LOCAL_BIN}/buf"

USER gitpod
