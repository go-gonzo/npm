package pkg


type Version string

type Person struct {
	Name string
	Email string
	Web string //URL?
}

type People []Person

type Bugs struct {
	Name string
	Web string
}

type License struct {
	Type string
	Web string //url?
}

type Repository struct {
	Type string
	URL string //URL?
}

type Repositories []Repository

type Dependency struct {
	Version
	Options map[string]Dependency
}

type Dependencies map[string]Dependency

type Dist struct {
	Shasum string //[]byte?
	Tarball string //url?
}


type Package struct {
	Name string
	Description string
	Version Version

	Maintainers People
	Contributors People

	Bugs Bugs
	Repositories Repositories

	Implements []string
	OS []string
	CPU []string
	Engines []string 
	Scripts map[string]string

	Dist Dist
}


type PackageRoot struct {
	Package
	Versions []Package
}

