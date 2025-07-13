# Static site generator in Go

Build
```
go build -o site
```

Run
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
