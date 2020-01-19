#!/bin/bash

set -eu

client_configure() {
	sudo chmod 600 $PQSSLCERTTEST_PATH/***REMOVED***ql.key
}

pgdg_repository() {
	local sourcelist='sources.list.d/***REMOVED***ql.list'

	curl -sS 'https://www.***REMOVED***ql.org/media/keys/ACCC4CF8.asc' | sudo apt-key add -
	echo deb http://apt.***REMOVED***ql.org/pub/repos/apt/ $(lsb_release -cs)-pgdg main $PGVERSION | sudo tee "/etc/apt/$sourcelist"
	sudo apt-get -o Dir::Etc::sourcelist="$sourcelist" -o Dir::Etc::sourceparts='-' -o APT::Get::List-Cleanup='0' update
}

***REMOVED***ql_configure() {
	sudo tee /etc/***REMOVED***ql/$PGVERSION/main/pg_hba.conf > /dev/null <<-config
		local     all         all                               trust
		hostnossl all         pqgossltest 127.0.0.1/32          reject
		hostnossl all         pqgosslcert 127.0.0.1/32          reject
		hostssl   all         pqgossltest 127.0.0.1/32          trust
		hostssl   all         pqgosslcert 127.0.0.1/32          cert
		host      all         all         127.0.0.1/32          trust
		hostnossl all         pqgossltest ::1/128               reject
		hostnossl all         pqgosslcert ::1/128               reject
		hostssl   all         pqgossltest ::1/128               trust
		hostssl   all         pqgosslcert ::1/128               cert
		host      all         all         ::1/128               trust
	config

	xargs sudo install -o ***REMOVED*** -g ***REMOVED*** -m 600 -t /var/lib/***REMOVED***ql/$PGVERSION/main/ <<-certificates
		certs/root.crt
		certs/server.crt
		certs/server.key
	certificates

	sort -VCu <<-versions ||
		$PGVERSION
		9.2
	versions
	sudo tee -a /etc/***REMOVED***ql/$PGVERSION/main/***REMOVED***ql.conf > /dev/null <<-config
		ssl_ca_file   = 'root.crt'
		ssl_cert_file = 'server.crt'
		ssl_key_file  = 'server.key'
	config

	echo 127.0.0.1 ***REMOVED*** | sudo tee -a /etc/hosts > /dev/null

	sudo service ***REMOVED***ql restart
}

***REMOVED***ql_install() {
	xargs sudo apt-get -y -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confnew' install <<-packages
		***REMOVED***ql-$PGVERSION
		***REMOVED***ql-server-dev-$PGVERSION
		***REMOVED***ql-contrib-$PGVERSION
	packages
}

***REMOVED***ql_uninstall() {
	sudo service ***REMOVED***ql stop
	xargs sudo apt-get -y --purge remove <<-packages
		libpq-dev
		libpq5
		***REMOVED***ql
		***REMOVED***ql-client-common
		***REMOVED***ql-common
	packages
	sudo rm -rf /var/lib/***REMOVED***ql
}

megacheck_install() {
	# Lock megacheck version at $MEGACHECK_VERSION to prevent spontaneous
	# new error messages in old code.
	go get -d honnef.co/go/tools/...
	git -C $GOPATH/src/honnef.co/go/tools/ checkout $MEGACHECK_VERSION
	go install honnef.co/go/tools/cmd/megacheck
	megacheck --version
}

golint_install() {
	go get github.com/golang/lint/golint
}

$1
