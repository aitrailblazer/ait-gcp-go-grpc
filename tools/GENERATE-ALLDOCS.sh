#!/bin/bash
#
go install github.com/mariotoffia/goasciidoc@latest

cd api
cd v1

goasciidoc -o api_d.adoc -i -t -n --nonexported  -m go.mod . -c "{\"author\": \"Constantine Vassilev\", \"email\": \"constantine@aitrailblazer.com\", \"web\": \"https://aitrailblazer.com\", \"images\": \"\", \"title\":\"ait-gcp-go-grpc\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

asciidoctor-reducer api_d.adoc -o api.adoc
rm -rf api_d.adoc
# goasciidoc -o docs.adoc -i -t -c "{\"author\": \"Constantine Vassilev\", \"email\": \"constantine@aitrailblazer.com\", \"web\": \"https://aitrailblazer.com\", \"images\": \"\", \"title\":\"ait-gcp-go-grpc\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

# cp docs.adoc api.adoc 
# rm -rf docs.adoc

cd ..
cd ..



# cd api
# golangci-lint run ./...

# echo goasciidoc -o api-doc.adoc -i -t --nonexported  -m go.mod . -c "{\"author\": \"Konstantin Vassilev\", \"email\": \"konstantin.vassilev@sscinc.com\", \"web\": \"https://code.ssnc.dev/cloud/cloud-registry\", \"images\": \"\", \"title\":\"cloud-registry\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

# goasciidoc -o api-d.adoc -i -t --nonexported  -m go.mod . -c "{\"author\": \"Konstantin Vassilev\", \"email\": \"konstantin.vassilev@sscinc.com\", \"web\": \"https://code.ssnc.dev/cloud/cloud-registry\", \"images\": \"\", \"title\":\"cloud-registry\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"
# asciidoctor-reducer api-d.adoc -o api-doc.adoc
# rm -rf api-d.adoc

# cd ..


# cd cmd
# cd srv
# golangci-lint run ./...

# echo goasciidoc -o cmd-srv-doc.adoc -i -t --nonexported  -m go.mod . -c "{\"author\": \"Konstantin Vassilev\", \"email\": \"konstantin.vassilev@sscinc.com\", \"web\": \"https://code.ssnc.dev/cloud/cloud-registry\", \"images\": \"\", \"title\":\"cloud-registry\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

# goasciidoc -o cmd-srv-d.adoc -i -t --nonexported  -m go.mod . -c "{\"author\": \"Konstantin Vassilev\", \"email\": \"konstantin.vassilev@sscinc.com\", \"web\": \"https://code.ssnc.dev/cloud/cloud-registry\", \"images\": \"\", \"title\":\"cloud-registry\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"
# asciidoctor-reducer cmd-srv-d.adoc -o cmd-srv-doc.adoc
# rm -rf cmd-srv-d.adoc

# cd ..
# cd ..

# asciidoctor-reducer makefile.adoc -o makefile-f.adoc
# plantuml diagrams/makefile.plantuml -o .
plantuml diagrams/main.plantuml -o .
plantuml diagrams/servers.plantuml -o .
mv diagrams/main.png  images
mv diagrams/servers.png  images