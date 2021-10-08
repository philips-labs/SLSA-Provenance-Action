<div id="top"></div>

<div align="center">
  
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
  
</div>

<br />
<div align="center">
  <a href="https://github.com/philips-labs/slsa-provenance-action">
    <img src="https://slsa.dev/images/levelBadge1.svg" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">SLSA Provenance GitHub Action</h3>

  <p align="center">
    Github Action implementation of SLSA Provenance Generation of level 1
    <br>
    <a href="https://github.com/philips-labs/slsa-provenance-action/issues">Report Bug</a>
    ·
    <a href="https://github.com/philips-labs/slsa-provenance-action/issues">Request Feature</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About This Project

This GitHub action implements the level 1 requirements of the [SLSA framework](https://slsa.dev/). By using this GitHub Action it is possible to easily generate the provenance file for different artifact types.
Different artifact types include:
* Files
* Push to some API (Docker Hub, different workflow)

While there are no integrity guarantees on the produced provenance at L1,
publishing artifact provenance in a common format opens up opportunities for
automated analysis and auditing. Additionally, moving build definitions into
source control and onto well-supported, secure build systems represents a marked
improvement from the ecosystem's current state.


This is not an official GitHub Action set up and maintained by the SLSA team. This GitHub Action is built for research purposes by Philips Research. It is heavily inspired by the original [Provenance Action example](https://github.com/slsa-framework/github-actions-demo) built by SLSA.

## Background

[SLSA](https://github.com/slsa-framework/slsa) is a framework intended to codify
and promote secure software supply-chain practices. SLSA helps trace software
artifacts (e.g. binaries) back to the build and source control systems that
produced them using in-toto's
[Attestation](https://github.com/in-toto/attestation/blob/main/spec/README.md)
metadata format.

### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* [SLSA Framework](https://github.com/slsa-framework/slsa/)
* [Golang](https://golang.org/)
* [GitHub Actions](https://github.com/features/actions)

<p align="right">(<a href="#top">back to top</a>)</p>


## Getting Started

TODO Add instructions to built locally. TODO
TODO Add instructions to add it to your repository. TODO

### Prerequisites

Ensure you have the following installed:
* Golang
* Docker

#### Recommendations 

The following IDE is recommended when working on this codebase:
* [VSCode](https://code.visualstudio.com/)

### Local Installation

1. Clone the repo.
   ```sh
   git clone git@github.com:philips-labs/slsa-provenance-action.git
   ```
1. Build the binary.
   ```sh
   make build
   ```
1. Execute the binary.
   ```sh
   ./bin/slsa-provenance help
   ```

<p align="right">(<a href="#top">back to top</a>)</p>

## Usage

The easiest way to use this Action is to add the following into your workflow file. Additional configuration might be nessecary to fit your usecase.

1. Add the following part in your workflow file:
   ```yaml
   generate-provenance:
    name: Generate build provenance
    runs-on: ubuntu-latest
    steps:
      - name: Download build artifact
        uses: actions/download-artifact@v2
        with:
          path: artifact/

      - name: Generate provenance
        uses: philips-labs/SLSA-Provenance-Action@8c78a6b34703824b9561a26b1ae5893beea9a332
        with:
          artifact_path: artifact/

      - name: Upload provenance
        uses: actions/upload-artifact@v2
        with:
          path: build.provenance
   ```
### Available commands
* Generate
* Help
* Version

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please refer to the [Contributing Guidelines](/CONTRIBUTING.md) for all the guidelines.


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Brend Smits - brend.smits@philips.com


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

This project is inspired by:

* [SLSA Framework](https://slsa.dev/)
* [SLSA GitHub Action Example](https://raw.githubusercontent.com/slsa-framework/github-actions-demo)

<p align="right">(<a href="#top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/philips-labs/slsa-provenance-action.svg?style=for-the-badge
[contributors-url]: https://github.com/philips-labs/slsa-provenance-action/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/philips-labs/slsa-provenance-action.svg?style=for-the-badge
[forks-url]: https://github.com/philips-labs/slsa-provenance-action/network/members
[stars-shield]: https://img.shields.io/github/stars/philips-labs/slsa-provenance-action.svg?style=for-the-badge
[stars-url]: https://github.com/philips-labs/slsa-provenance-action/stargazers
[issues-shield]: https://img.shields.io/github/issues/philips-labs/slsa-provenance-action.svg?style=for-the-badge
[issues-url]: https://github.com/philips-labs/slsa-provenance-action/issues
[license-shield]: https://img.shields.io/github/license/philips-labs/slsa-provenance-actionsvg?style=for-the-badge
[license-url]: https://github.com/philips-labs/slsa-provenance-action/blob/main/LICENSE.txt
