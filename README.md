# caseflux: text transformation tool

caseflux is a command-line tool for transforming text to various cases or
applying a Caesar cipher.

This project provides multiple text transformation options including sentence
case, lower case, upper case, capitalized case, alternating case, title case,
inverse case, and a Caesar cipher transformation.

Development is active, and new features are being added. Feel free to contribute
or suggest improvements!

## Table of Contents

  - [Usage](#usage)
  - [Support](#support)
  - [Roadmap](#roadmap)
  - [License](#license)

## Usage

Ensure you have Go installed. Clone the repository and run the following
command:

``` console
$ make
```

To transform text from a file using a specified transformation, use the
following command:

``` console
$ cflux -input=yourfile.txt -transform=lower
```

Alternatively, to read from stdin and apply a transformation:

``` console
$ echo "Your Text Here" | cflux -stdin -transform=upper
```

For the Caesar cipher transformation, you can specify a shift value:

``` console
$ cflux -input=yourfile.txt -transform=cipher -shift=3
```

## Roadmap

  - [ ] Add support for additional text transformations
  - [ ] Implement unit tests
  - [ ] Provide a more user-friendly CLI with subcommands
  - [ ] Improve performance for large text files
  - [ ] Add a configuration file for default settings

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
Before contributing, please follow these guidelines:

  - Follow the [code of conduct](CODE_OF_CONDUCT.md).
  - Keep a consistent coding style. To ensure your coding style remains the
    same, format your code with:
    ``` console
    $ go fmt path/to/source_code
    ```
  - Use the stable version of Go.
  - Prefer using the standard library over reinventing the wheel.

For support, please [open an
issue](https://github.com/walker84837/caseflux/issues).

## License

This project is licensed under the [BSD-3-Clause License](LICENSE.md).
