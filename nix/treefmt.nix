{ pkgs, ... }:
{
  projectRootFile = "flake.nix";

  programs = {
    deadnix.enable = true;
    golines.enable = true;
    nixfmt.enable = true;
    prettier.enable = true;
    shfmt.enable = true;
    templ.enable = true;
  };

  settings = {
    global.excludes = [
      "vendor/*"
    ];

    formatter = {
      golines = {
        options = [
          "--base-formatter=${pkgs.gofumpt}/bin/gofumpt"
        ];
      };

      prettier = {
        package = pkgs.nodePackages.prettier;
        excludes = [
        ];
      };
    };
  };
}
