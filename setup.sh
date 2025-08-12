# Create bin directory if it doesn't exist
mkdir -p ./bin

os=$(uname -s)
arch=$(uname -m)

if [[ "$os" == "Linux" ]]; then
    if [[ "$arch" == "aarch64" || "$arch" == "arm64" ]]; then
        download_url="https://github.com/use-ink/ink-node/releases/download/v0.44.0/ink-node-linux-arm64.tar.gz"
    else
        download_url="https://github.com/use-ink/ink-node/releases/download/v0.44.0/ink-node-linux.tar.gz"
    fi
elif [[ "$os" == "Darwin" ]]; then
    download_url="https://github.com/use-ink/ink-node/releases/download/v0.44.0/ink-node-mac-universal.tar.gz"
else
    echo "unsupported"
fi


curl -L "$download_url" | tar -xz -C "./bin"
mv ./bin/*/* ./bin
rm -rf ./bin/*/
chmod +x ./bin/ink-node