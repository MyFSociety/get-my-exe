package core

type PackageStruct struct {
	Name    string
	URL     string
	SHAType string
	SHA     string
}

var packages = []PackageStruct{
	{"archlinux", "https://mirrors.abhy.me/archlinux/iso/2023.02.01/archlinux-x86_64.iso", "sha256", "c30718ab8e4af1a3b315ce8440f29bc3631cb67e9656cfe1e0b9fc81a5c6bf9c"},
}

func GetPackage(name string) PackageStruct {
	for _, pkg := range packages {
		if pkg.Name == name {
			return pkg
		}
	}
	return PackageStruct{}
}
