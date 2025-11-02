{
  fetchFromGitHub,
  installShellFiles,
  makeWrapper,
  makeBinPath,
  go,
  goBuild,
  ...
}:
goBuild rec {
  pname = "golangci-lint";
  version = "2.1.6";

  src = fetchFromGitHub {
    owner = "golangci";
    repo = "golangci-lint";
    rev = "v${version}";
    hash = "sha256-L0TsVOUSU+nfxXyWsFLe+eU4ZxWbW3bHByQVatsTpXA=";
  };

  vendorHash = "sha256-tYoAUumnHgA8Al3jKjS8P/ZkUlfbmmmBcJYUR7+5u9w=";

  subPackages = [ "cmd/golangci-lint" ];

  nativeBuildInputs = [
    installShellFiles
    makeWrapper
  ];

  ldflags = [
    "-s"
    "-X main.version=${version}"
    "-X main.commit=v${version}"
    "-X main.date=19700101-00:00:00"
  ];

  postInstall = ''
    for shell in bash zsh fish; do
      HOME=$TMPDIR $out/bin/golangci-lint completion $shell > golangci-lint.$shell
      installShellCompletion golangci-lint.$shell
    done
  '';

  allowGoReference = true;

  postFixup = ''
    wrapProgram $out/bin/golangci-lint --prefix PATH : ${makeBinPath [ go ]}
  '';
}
