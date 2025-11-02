{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
    nixpkgs-stable.url = "github:nixos/nixpkgs/nixos-25.05";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    pre-commit-hooks.url = "github:cachix/git-hooks.nix";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  # ACME_EMAIL=janos.rocha@gmail.com
  # ACME_DOMAIN=jrocha.eu
  outputs =
    { self, ... }@inputs:
    inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = inputs.nixpkgs.legacyPackages.${system};
        lib = pkgs.lib;

        inherit (lib) makeBinPath;

        go = pkgs.go_1_24;
        goBuild = pkgs.buildGo124Module;

        golangci-lint = pkgs.callPackage ./nix/golangci-lint.nix {
          inherit go goBuild makeBinPath;
        };

        scripts = import ./nix/scripts.nix { inherit pkgs go; };

        devInputs = with pkgs; [
          git
          git-chglog
          gofumpt
          golines
          gopls
          tinygo
          scripts.lint
          scripts.test
          scripts.tidy
          go
          sass
          templ
          golangci-lint
          golangci-lint-langserver
          gopls
        ];

        treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs ./nix/treefmt.nix;

        preCommit = import ./nix/pre-commit.nix {
          inherit golangci-lint;
          gofumpt = pkgs.gofumpt;
          nixfmt-rfc-style = pkgs.nixfmt-rfc-style;
          prettier = pkgs.nodePackages.prettier;
          trufflehog = pkgs.trufflehog;
        };

      in
      {
        devShells.default = pkgs.mkShell {
          packages = devInputs;
          inherit (self.checks.${system}.pre-commit-check) shellHook;
        };

        checks = {
          formatting = treefmtEval.config.build.check inputs.self;
          pre-commit-check = inputs.pre-commit-hooks.lib.${system}.run preCommit;
        };

        formatter = treefmtEval.config.build.wrapper;
      }
    );
}
