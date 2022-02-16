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

```
Please download the right release,
Put it in ~/Dev/github.com/h4ux/ (or what ever your pathe looks like) 
Add a .env as shown below
```

### .env file format
```
GH_TOKEN={GITHUB Personal token goes here}
GH_ORG={GITHUB USER NAME or ORG}
GO_FOLDERS=bin,cmd,configs,deployments,docs,internal,pkg,tests #folders that you whould like to create separated by comma
```
** Currently supports Mac OS and Linux OS (Windows can be added with very little effort)
