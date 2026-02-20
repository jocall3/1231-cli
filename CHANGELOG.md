# Changelog

## 0.1.0 (2026-02-20)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/jocall3/1231-cli/compare/v0.0.1...v0.1.0)

### ⚠ BREAKING CHANGES

* add support for passing files as parameters

### Features

* add readme documentation for passing files as arguments ([1b3c8c4](https://github.com/jocall3/1231-cli/commit/1b3c8c457fcef9c50e878edebdcc746e7cbacadd))
* add support for passing files as parameters ([af78eef](https://github.com/jocall3/1231-cli/commit/af78eefdcaad92b35a8bcd415874d1c5f91708e8))
* added mock server tests ([a52c71e](https://github.com/jocall3/1231-cli/commit/a52c71e07b248641fbaffa2769e01eb23ce20f13))
* added support for --foo.baz inner field flags ([188dcce](https://github.com/jocall3/1231-cli/commit/188dcce7359fb40322c96bf65eeaa37a749d2da5))
* **api:** manual updates ([2d51c7f](https://github.com/jocall3/1231-cli/commit/2d51c7f921146dcd3a3780dcb982971221abb53b))
* **api:** manual updates ([fb95a99](https://github.com/jocall3/1231-cli/commit/fb95a994d85fb2835b955b1100c924692dd0c68f))
* **api:** manual updates ([29b73ba](https://github.com/jocall3/1231-cli/commit/29b73badf86f5aebdc7b305b0124778d3d35e221))
* **api:** manual updates ([30289ab](https://github.com/jocall3/1231-cli/commit/30289ab5b873376a9bc15754d03e4d74ce5cc57f))
* **api:** manual updates ([093c8c3](https://github.com/jocall3/1231-cli/commit/093c8c3fdae1f7640aec9d2ec26c0e2a3cdd4f7c))
* **api:** manual updates ([fbd0580](https://github.com/jocall3/1231-cli/commit/fbd05802d673f3adc9b7f7f39018119aa7b2c80f))
* **client:** provide file completions when using file embed syntax ([7e5c55a](https://github.com/jocall3/1231-cli/commit/7e5c55afab2c58300ecb6d0bcb01284e6fea7ef6))
* **cli:** improve shell completions for namespaced commands and flags ([28160fd](https://github.com/jocall3/1231-cli/commit/28160fd8d77d5e5553bc9a8c5aabe9e52f612c46))
* enable CI tests ([e9ce4ec](https://github.com/jocall3/1231-cli/commit/e9ce4eccb1839c0ea9fbd2bcca58e4db90611ac8))
* improved behavior for exploring paginated/streamed endpoints ([5d7fcd5](https://github.com/jocall3/1231-cli/commit/5d7fcd57ffa4ab4421a8865dd59a696c7593bb07))
* improved support for passing files for `any`-typed arguments ([be83eee](https://github.com/jocall3/1231-cli/commit/be83eee107661cf8b760565909879f1e5684d9b5))


### Bug Fixes

* check required arguments ([bb2718e](https://github.com/jocall3/1231-cli/commit/bb2718ebe9446116c017f87c88da7e9e37416226))
* fix for file uploads to octet stream and form encoding endpoints ([6e80fac](https://github.com/jocall3/1231-cli/commit/6e80facaad8e52ef9c0fbe7104e0ea34c0b68f06))
* fix for generated tests for some array flags ([fb1d1b0](https://github.com/jocall3/1231-cli/commit/fb1d1b05b6af0aef5ec475b739702e3e38c48f75))
* fix for nullable arguments ([64e4150](https://github.com/jocall3/1231-cli/commit/64e4150e02b9db7786bf10b80396bde201b3856c))
* fix for when terminal width is not available ([e116c5f](https://github.com/jocall3/1231-cli/commit/e116c5f837ee2fd978559bc26e5b6ce4313f6908))
* fix generated flag types and value wrapping ([a47b840](https://github.com/jocall3/1231-cli/commit/a47b8400bcd8849786f5fc1bc94879a74c36d6c0))
* fix mock tests with inner fields that have underscores ([ad29f15](https://github.com/jocall3/1231-cli/commit/ad29f1501ef8790169b57bcd7e6913a953eaf233))
* fix terminal height issues causing test failures ([8f7d266](https://github.com/jocall3/1231-cli/commit/8f7d26698a442a6edb19a3d4a56c53d540166939))
* fixed placeholders for date/time arguments ([fb40a5a](https://github.com/jocall3/1231-cli/commit/fb40a5a6c842062c22731a044dfdaa6dbb4fa6c4))
* flag defaults ([24f1e1e](https://github.com/jocall3/1231-cli/commit/24f1e1ec42d29bd8e0fa2992b1b1b822e657dd14))
* overly broad redaction of Authorization ([4ffbe9d](https://github.com/jocall3/1231-cli/commit/4ffbe9d06e7a8c96a2affd8f17e812e981a8b0ba))
* preserve filename in content-disposition for file uploads ([babfbf3](https://github.com/jocall3/1231-cli/commit/babfbf30677599bde99f65f332e15529e0a05a78))
* prevent flag duplication ([6ea35ee](https://github.com/jocall3/1231-cli/commit/6ea35ee76c2dda21bfe24ab55af034825affbb53))
* prevent tests from hanging on streaming/paginated endpoints ([2e96a62](https://github.com/jocall3/1231-cli/commit/2e96a620fa99c2e5e30ec85051f97a046b64e620))
* remove unsupported methods ([500fcad](https://github.com/jocall3/1231-cli/commit/500fcad877ca4c8c7c3fa0c463a663d29d684a13))
* restore support for void endpoints ([a89ebea](https://github.com/jocall3/1231-cli/commit/a89ebea05a95fd0df8622b8f825f188bb5ed5b49))
* use RawJSON for iterated values instead of re-marshalling ([c7c6107](https://github.com/jocall3/1231-cli/commit/c7c6107fe1674da0ac15be7224f7bad316e9f47d))


### Chores

* add build step to ci ([94df75a](https://github.com/jocall3/1231-cli/commit/94df75a85024e10295e5b5e766ac24f7298a3407))
* **internal:** codegen related update ([ef1cd0c](https://github.com/jocall3/1231-cli/commit/ef1cd0c7ce68f5f7520738e6136a39dcb366b389))
* **internal:** codegen related update ([bcb9333](https://github.com/jocall3/1231-cli/commit/bcb9333e10612898f15ff3bd6b9461e0b02d4f5f))
* **internal:** codegen related update ([cf383bb](https://github.com/jocall3/1231-cli/commit/cf383bbd471f0ef8fe05e856797e46a144da03d4))
* **internal:** codegen related update ([bce50c2](https://github.com/jocall3/1231-cli/commit/bce50c2e488ebe98d2353e189e9d9b6806043a58))
* **internal:** codegen related update ([c80a340](https://github.com/jocall3/1231-cli/commit/c80a340eb7448636adfe5e04147dad27bbaae080))
* **internal:** codegen related update ([c190531](https://github.com/jocall3/1231-cli/commit/c1905313a95e6d8baaf1a591efef6d78e53b6b1e))
* **internal:** remove mock server code ([b327a8e](https://github.com/jocall3/1231-cli/commit/b327a8e77573bd5628479a611f64d894df20bf78))
* **internal:** update `actions/checkout` version ([ffb810c](https://github.com/jocall3/1231-cli/commit/ffb810cc6fdf4ebea4008986a3d300784625412b))
* update documentation in readme ([7f39075](https://github.com/jocall3/1231-cli/commit/7f390758705296e0253996fe07c09cf6007b7aa8))
* update internal comment ([4ffbe9d](https://github.com/jocall3/1231-cli/commit/4ffbe9d06e7a8c96a2affd8f17e812e981a8b0ba))
* update mock server docs ([af58cc5](https://github.com/jocall3/1231-cli/commit/af58cc59684552df24109faf367c4fcd3f539f1c))
* update SDK settings ([19bcc25](https://github.com/jocall3/1231-cli/commit/19bcc25bea07552804dc1687938d3db5ea4a0c94))
* update SDK settings ([70f127b](https://github.com/jocall3/1231-cli/commit/70f127b922f3aeac2c8a83cb9d8edcff96ac9214))
* update SDK settings ([aeda910](https://github.com/jocall3/1231-cli/commit/aeda910debab762f2836f2363260398647c647e4))
* update SDK settings ([ea539ba](https://github.com/jocall3/1231-cli/commit/ea539ba12334781982fe21b08ee9d3e9f9566d5c))
