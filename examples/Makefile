
ifeq ($(PHPDIR),)
	PHPDIR=/usr
endif

PHPCONFIG=$(PHPDIR)/bin/php-config
# PHPCONFIG=$(PHPDIR)/bin/php-config55

export PATH := $(PHPDIR)/bin:$(PATH)
export CFLAGS := $(shell export PATH=$(PATH) && $(PHPCONFIG) --includes)
export LDFLAGS := -L$(shell export PATH=$(PATH) && $(PHPCONFIG) --prefix)/lib/

export GOPATH := $(PWD):$(GOPATH)
export CGO_CFLAGS := $(CFLAGS) $(CGO_CFLAGS)
export CGO_LDFLAGS := $(LDFLAGS) $(CGO_LDFLAGS)

all:
	go build -v -buildmode=c-shared -o ./_out/protorpc-php.so github.com/asyou-me/protorpc-php
	php -d extension=./_out/protorpc-php.so examples/test.php

clean:
	rm -f ../../pkg/linux_amd64/zend.a
	rm -f ../../pkg/linux_amd64/phpgo.a
	rm -f ./_out/protorpc-php.so
