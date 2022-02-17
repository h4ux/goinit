# goinit
New repository created by goinit

![Release](https://github.com/h4ux/goinit/actions/workflows/release.yml/badge.svg)
[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)
[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)



This is a very simple Go Lang project structure creator and github repo.

I suggest the following directory structure for conveniance

Create a path in your working station containg your Github project structure

Ex:

```
~/Dev/github.com/h4ux/{project name / repository name}
```

## Installation via install.sh

```bash
# binary will be in $(go env GOPATH)/bin/goinit
curl -sSfL https://raw.githubusercontent.com/h4ux/goinit/main/install.sh | sh -s -- -b $(go env GOPATH)/bin

# defualt installation into ./bin/
curl -sSfL https://raw.githubusercontent.com/h4ux/goinit/main/install.sh | sh -s

```

Once you install a file called .env will be created in ~/.config/goinit/.env
Please add the relevant data as described below

You can overide the global .env file by creating a local one on the working directory

### .env file format
```
GH_TOKEN={GITHUB Personal token goes here}
GH_ORG={GITHUB USER NAME or ORG}
GO_FOLDERS=bin,cmd,configs,deployments,docs,internal,pkg,tests #folders that you whould like to create separated by comma
GO_PROJECTS_PATH={your full path for your Go projects}
```

** Currently supports Mac OS and Linux OS (Windows can be added with very little effort)
