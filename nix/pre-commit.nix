{
  gofumpt,
  golangci-lint,
  nixfmt-rfc-style,
  prettier,
  trufflehog,
  ...
}:
{
  src = ./.;
  excludes = [
    "flake.lock"
    "vendor/.+"
  ];
  hooks = {
    check-yaml.enable = true;
    convco.enable = true;
    deadnix.enable = true;
    ripsecrets.enable = true;
    shellcheck.enable = true;

    nixfmt-rfc-style = {
      enable = true;
      package = nixfmt-rfc-style;
    };

    trufflehog = {
      enable = true;
      package = trufflehog;
    };

    prettier = {
      enable = true;
      package = prettier;
    };

    golines = {
      enable = true;
      settings.flags = "--base-formatter=${gofumpt}/bin/gofumpt";
    };

    golangci-lint = {
      enable = true;
      package = golangci-lint;
    };
  };
}
