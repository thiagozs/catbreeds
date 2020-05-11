GOCMD=go
GOBUILD=$(GOCMD) build
GOENV=$(GOCMD) env
GREP=grep
SED=sed
FIND=find
OUTDIR=out
VERSION=1.0.0
LDFLAGS=-ldflags "-X main.version=${VERSION}"

build:
	rm -fr ${OUTDIR}
	mkdir -p ${OUTDIR}
	GOOS=linux GOARCH=arm GOARM=6 ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/acsvgen.rpi
	GOOS=linux ${GOBUILD}  ${LDFLAGS} -o ${OUTDIR}/acsvgen.lin
	GOOS=darwin ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/acsvgen.mac
	zip ${OUTDIR}/acsvgen.rpi.zip ${OUTDIR}/acsvgen.rpi
	zip ${OUTDIR}/acsvgen.lin.zip ${OUTDIR}/acsvgen.lin
	zip ${OUTDIR}/acsvgen.mac.zip ${OUTDIR}/acsvgen.mac
