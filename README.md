# Go Starters

![Logo](https://github.com/hanggrian/go-starters/raw/assets/logo.png)

Common Go project templates, separated by target platform and kind of
distribution.

| | Testing | Publishing | Website | Coverage
--- | :---: | :---: | :---: | :---:
application | [testing] | &cross; | [Hugo Book] | &check;
library | [testing] | [Go Packages] | [doc2go], [Hugo Book] | &check;

## Frameworks

- Built-in testing framework with [gomock](https://github.com/uber-go/mock/)
  suite.
- Standard formatter [gofmt](https://pkg.go.dev/cmd/gofmt/).
- [Coverage profiling](https://go.dev/doc/build-cover/) for integration tests.
- Tasks defined in [Makefile](https://www.gnu.org/software/make/).

## Project layout

- Root directory:
  - GitHub [README](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-readmes/),
    [LICENSE](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository/),
    and [gitignore](https://docs.github.com/en/get-started/getting-started-with-git/ignoring-files/)
    file.
  - [EditorConfig](https://editorconfig.org/) enforces IDE settings.
- [GitHub Actions](https://docs.github.com/en/actions/) workflows:
  - Run tests, linters and push coverage to [Codecov](https://codecov.io/).
  - Activate [Renovate](https://docs.renovatebot.com/) bot to alert out-of-date
    dependencies.
- [Go Modules](https://go.dev/blog/using-go-modules/):
  - `go.mod` for dependency management.
  - `go.sum` for integrity verification.
- Website module:
  - [Hugo](https://gohugo.io/) for generating webpages displaying README and
    other content.
  - The webpages are manually deployed.

[testing]: https://pkg.go.dev/testing/
[Go Packages]: https://pkg.go.dev/
[doc2go]: https://abhinav.github.io/doc2go/
[Hugo Book]: https://book.alxs.dev/
