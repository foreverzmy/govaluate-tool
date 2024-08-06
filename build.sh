cd wasm && tinygo build -o govaluate.wasm -target wasm --no-debug main.go 
wasm-opt -d govaluate.wasm -o  output.wasm
