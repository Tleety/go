[![Go](https://github.com/Tleety/go/actions/workflows/go.yml/badge.svg)](https://github.com/Tleety/go/actions/workflows/go.yml)

# Project for learning GO
This project is for me to test and implement best practices for Go developement.

## Folders

### api
API protocol definitions.

### cmd
Contains the main applications for the project. The directory name for each application should match the name of the executable you want to have.

#### main_package
For my main.go file currently.
This is currently my only application, not sure if I will have other.

### config
Configuration file templates or default configs.

### internal
Private code that you don’t want others importing in other projects. Only this project can access and use it.

### pkg
Library code that's ok to use by external applications. Other projects can import these libraries.

### scripts
Scripts to perform build, install, analysis, etc.

### tests
Additional external test apps and test data.