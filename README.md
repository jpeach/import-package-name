# import-package-name

`import-package-name` is a Go linter that checks if packages are imported
by pre-defined names.
This is useful when your project defines rules about what names packages
should be imported by, and you want the package names to be consistent.

## Usage

```bash
$ import-package-name -help
importpackagename: Check package import naming

Usage: importpackagename [-flag] [package]


Flags:
  -V	print version and exit
  -all
    	no effect (deprecated)
  -c int
    	display offending line with this many lines of context (default -1)
  -cpuprofile string
    	write CPU profile to this file
  -debug string
    	debug flags, any subset of "fpstv"
  -fix
    	apply all suggested fixes
  -flags
    	print analyzer flags in JSON
  -imports value
    	comma-separated list of Name=Path import specs
  -json
    	emit JSON output
  -memprofile string
    	write memory profile to this file
  -source
    	no effect (deprecated)
  -tags string
    	no effect (deprecated)
  -trace string
    	write trace log to this file
  -unsafeptr

  -v	no effect (deprecated)
```

To run `import-package-name` on your project, use the `-imports` flag:
```bash
$ import-package-name -imports mesh_proto=github.com/kumahq/kuma/api/mesh/v1alpha1,core_mesh=github.com/kumahq/kuma/pkg/core/resources/apis/mesh ./...

/Users/jpeach/upstream/kuma/pkg/mads/util.go:12:2: import package name "mesh_core" should be "core_mesh"
/Users/jpeach/upstream/kuma/pkg/envoy/admin/client.go:18:2: import package name "mesh_core" should be "core_mesh"
/Users/jpeach/upstream/kuma/pkg/xds/template/proxy_template.go:4:2: import package name "kuma_mesh" should be "mesh_proto"
...
```

The `-fix` flag will rewrite the package name, but see the [Bugs](#bugs) section below.

## Bugs

If `import-package-name` changes the name a package is imported by,
it doesn't fix up all the references to names in that package.
