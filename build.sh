cd push-swap-src
go build -o push-swap
mv push-swap ../push-swap
cd ../checker-src
go build -o checker
mv checker ../checker