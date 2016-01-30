package npm

import (
	"path/filepath"
	"strings"

	"github.com/omeid/gonzo"
	"github.com/omeid/gonzo/context"

	"github.com/go-gonzo/archive/tar"
	"github.com/go-gonzo/compress/gzip"
	"github.com/go-gonzo/fs"
	"github.com/go-gonzo/npm/commonjs/package"
	"github.com/go-gonzo/npm/commonjs/registry/client"
	"github.com/go-gonzo/path"
	"github.com/go-gonzo/util"
	"github.com/go-gonzo/web"
)

func Install(dist string, packages ...string) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		pkgs, err := Packages(packages...)
		if err != nil {
			return err
		}

		return Get(ctx, pkgs...).Then(
			fs.Dest(dist),
		)
	}
}

//An arrya of packages in "PRODUCT@tag/version/range" Format.
func Packages(packages ...string) ([]pkg.Package, error) {

	pkgs := []pkg.Package{}

	for _, p := range packages {
		pp := strings.Split(p, "@")
		var name, version string
		name = pp[0]
		if len(pp) == 1 || pp[1] == "" {
			version = "latest"
		} else {
			version = pp[1]
		}

		pkg, err := client.Get(name, version)
		if err != nil {
			return nil, err
		}


		pkgs = append(pkgs, *pkg)
	}

	return pkgs, nil
}

func Get(ctx context.Context, pkgs ...pkg.Package) gonzo.Pipe {

	var all []gonzo.Pipe
	for _, pkg := range pkgs {
		if pkg.Dist.Tarball == "" {
			ctx.Info("EMPTY", pkg.Name)
			continue
		}
		all = append(all, get(ctx, pkg))
	}
	return util.Merge(ctx, all...)
}

func get(ctx context.Context, pkg pkg.Package) gonzo.Pipe {

	return web.Get(ctx, pkg.Dist.Tarball).Pipe(
		gzip.Uncompress(),
		tar.Untar(tar.Options{
			StripComponenets: 1,
		}),
		path.Rename(func(old string) string {
			return filepath.Join(pkg.Name, old)
		}),
	)
}
