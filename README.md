# Static site generator in Go

Build
```
go build -o site
```

Install globally
```
go install .
```

This places the binary in `~/go/bin/`. Make sure it's in your PATH by adding this to your `~/.bashrc` or `~/.zshrc`:
```
export PATH=$PATH:$(go env GOPATH)/bin
```

Then you can run `website-generator-go` from any directory.

Run locally
```
./site
```

Test
```
go build -o site
cp site test/
cd test
find . -name "index.html" -exec rm  {} \;
./site
```

Ensure that all index.html files were created and open them in the browser.
