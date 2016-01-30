package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/go-gonzo/npm/commonjs/package"
)

const pkgNotFound = "Package (%s@%s) Not Found."

const (
	NPM = "http://registry.npmjs.org/"
)

func Get(name, version string) (*pkg.Package, error) {

	npm, err := url.Parse(NPM)

	if err != nil {
		return nil, err
	}

	npm.Path = filepath.Join(npm.Path, name, version)

	resp, err := http.Get(npm.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf(pkgNotFound, name, version)
	}
	pkg := &pkg.Package{}
	err = json.NewDecoder(resp.Body).Decode(pkg)
	return pkg, err
}
