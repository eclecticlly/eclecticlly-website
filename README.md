<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Template" />

  &#xa0;

  <!-- <a href="https://template.netlify.app">Demo</a> -->
</div>

<h1 align="center">Eclecticlly Website</h1>

<p align="center">
  <img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/eclecticlly/eclecticlly-website/Eclecticlly%20workflow">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/eclecticlly/eclecticlly-website?color=56BEB8">
  <img alt="Github language count" src="https://img.shields.io/github/languages/count/eclecticlly/eclecticlly-website?color=56BEB8">
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/eclecticlly/eclecticlly-website?color=56BEB8">
  <img alt="License" src="https://img.shields.io/github/license/eclecticlly/eclecticlly-website?color=56BEB8">
</p>

<!-- Status -->

<hr>

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/eclecticlly" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

Eclecticlly Website based to [Eclecticlly Web-Stack](https://github.com/eclecticlly/web-stack)

## :sparkles: Features ##

:heavy_check_mark: Easy development workflow with Air;\
:heavy_check_mark: Authentication out of the box;\
:heavy_check_mark: Postgres and GORM;\
:heavy_check_mark: BDD with Goconvey;\
:heavy_check_mark: Jet Template and Tailwind;

## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/index.html)
- [Jet Template Engine](https://github.com/CloudyKit/jet/)
- [Tailwind](https://tailwindcss.com/)
- [xtmx](https://htmx.org/)

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com), [Tailwind standalone CLI](), [Make (pay attantion on Windows)](http://gnuwin32.sourceforge.net/), [Goconvey](http://goconvey.co/) and [Air](https://github.com/cosmtrek/air) installed.

## :checkered_flag: Starting ##

```bash
# Access
$ cd [yourepository]

# Create a postgres database
createdb [databasename]

# Create the env file and fill in the blanks
$ mv .env-example .env

# Start in Development mode (Air, Tailwind and Goconvey)
$ make dev

```

## :memo: License ##

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.


Made with :heart: by <a href="https://github.com/eclecticlly" target="_blank">Eclecticlly</a>

&#xa0;

<a href="#top">Back to top</a>
