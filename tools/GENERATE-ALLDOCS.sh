#!/bin/bash
#
go install github.com/mariotoffia/goasciidoc@latest

cd api
cd v1
cd api

goasciidoc -o api_d.adoc -i -t -n --nonexported  -m go.mod . -c "{\"author\": \"Constantine Vassilev\", \"email\": \"constantine@aitrailblazer.com\", \"web\": \"https://aitrailblazer.com\", \"images\": \"\", \"title\":\"ait-gcp-go-grpc\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

asciidoctor-reducer api_d.adoc -o api.adoc
rm -rf api_d.adoc
# goasciidoc -o docs.adoc -i -t -c "{\"author\": \"Constantine Vassilev\", \"email\": \"constantine@aitrailblazer.com\", \"web\": \"https://aitrailblazer.com\", \"images\": \"\", \"title\":\"ait-gcp-go-grpc\", \"toc\": \"Table of Contents\", \"toclevel\": 3}"

# cp docs.adoc api.adoc 
# rm -rf docs.adoc
cd ..
cd ..
cd ..

plantuml diagrams/main.plantuml -o .
plantuml diagrams/servers.plantuml -o .
mv diagrams/main.png  images
mv diagrams/servers.png  images