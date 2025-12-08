{
  description = "Atelier Gerber Website";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
    templ.url = "github:a-h/templ";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      templ,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        templ-cli = templ.packages.${system}.templ;
        projectName = "atelier-gerber";
      in
      {
        # Build the Go binary
        packages.default = pkgs.buildGoModule {
          pname = projectName;
          version = "1.0.1";
          src = ./.;

          vendorHash = "sha256-pcipMz5byZ55EoQUjOMc10nqhFtUh7Agcn599uSQ7Ys=";

          # Disable CGO for static binary
          CGO_ENABLED = "0";

          # Static linking flags
          ldflags = [
            "-s"
            "-w"
            "-extldflags -static"
          ];

          preBuild = ''
            ${templ-cli}/bin/templ generate


            # Generate Tailwind CSS
            ${pkgs.tailwindcss_4}/bin/tailwindcss \
              -i ./static/css/input.css \
              -o ./static/css/styles.css \
              --minify
          '';

          # Only include static files needed at runtime
          postInstall = ''
            mkdir -p $out/static/css
            cp ./static/css/styles.css $out/static/css/
          '';
        };

        # Development shell
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            templ-cli
            air
            bun
            fswatch
            nodePackages.tailwindcss
            flyctl
          ];

          shellHook = ''
            make help
            exec zsh
          '';
        };
      }
    );
}
