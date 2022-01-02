# NOTE: I can't write makefiles

.PHONY: clean

all: cmd/simple-test/simple-test honggfuzz/honggfuzz libs/libhfuzz.a libs/libhfcommon.a

cmd/simple-test/simple-test: libs/libhfuzz.a libs/libhfcommon.a
	cd cmd/simple-test && go build

libs/libhfuzz.a: honggfuzz/honggfuzz
	cp ./honggfuzz/libhfuzz/libhfuzz.a ./libs/

libs/libhfcommon.a: honggfuzz/honggfuzz
	cp ./honggfuzz/libhfcommon/libhfcommon.a ./libs/

honggfuzz/honggfuzz:
	git submodule update \
	&& cd honggfuzz \
	&& $(MAKE) -j

clean:
	rm -rf honggfuzz
	rm -f ./libs/*
	rm -f ./cmd/simple-test/simple-test
