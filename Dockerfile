FROM golang:1.23.6-bookworm AS build

WORKDIR /build

# Install system dependencies for VDF compilation
RUN apt-get update && apt-get install -y \
    build-essential \
    libgmp-dev \
    curl \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Install Rust toolchain
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

# Copy full source
ADD . /build

# Set environment variables for GMP linking on Linux
ENV LIBRARY_PATH="/usr/lib/x86_64-linux-gnu:/usr/lib"
ENV CPATH="/usr/include"

RUN make build

FROM debian:stable-slim

# Install runtime dependencies for VDF
RUN apt-get update && apt-get install -y \
    libgmp10 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /build/bin/performer /usr/local/bin/performer

CMD ["/usr/local/bin/performer"]
