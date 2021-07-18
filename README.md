# The Ray Tracer Challange 

This is a attempt to compete the ray tracer challange in golang


# VS Code Issues

`gopls`, the language server for go, does not currently support multiple modes in a single workspace by default.
This will cause all kinds of import errors and horrors, and if there is an import error, forget about any kind of intlliense, auto import and such.
To enable exprimental support. Add the following to your VScode `settings.json` or workspace config:

```
  "gopls": {
  
    "buildFlags": [
      "-tags=wireinject"
    ],
    "experimentalWorkspaceModule": true
  },
}
```