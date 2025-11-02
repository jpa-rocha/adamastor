{ goBuild, filter, ... }:
let
  srcFilter = filter {
    root = ./.;
    include = [
      ./cmd
      ./go.mod
      ./go.sum
      ./internal
      ./main.go
      ./vendor
    ];
  };
in
goBuild {
  name = "github.com/jpa-rocha/adamastor";
  src = srcFilter;
  doCheck = false;
  vendorHash = null;
  ldflags = [
    "-s"
    "-w"
    "-extldflags '-static'"
  ];
  env.CGO_ENABLED = 1;
}
