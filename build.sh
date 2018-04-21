cd ${HOME}/go/src/mcap
${HOME}/go/bin/gox -osarch="linux/amd64 windows/amd64 darwin/amd64"  -output="mcap_{{.OS}}"