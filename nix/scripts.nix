{ pkgs, go, ... }:
{
  tidy = pkgs.writeShellScriptBin "run-tidy" ''
    set -euo pipefail
    ${go}/bin/go mod tidy
    ${go}/bin/go mod vendor
  '';

  test = pkgs.writeShellScriptBin "run-test" ''
    set -euo pipefail
    ${go}/bin/go test -v -race -failfast -count 1 ./internal/...
  '';

  lint = pkgs.writeShellScriptBin "run-lint" ''
    set -euo pipefail
    ${pkgs.nix}/bin/nix flake check
  '';

  generate = pkgs.writeShellScriptBin "gen-templates" ''
    set -euo pipefail
    ${pkgs.templ}/bin/templ generate --path ./internal/server/components/
    ${pkgs.templ}/bin/templ generate --path ./internal/server/templates/
  '';

}
