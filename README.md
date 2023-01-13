## Nextvue CLI Application

[[_TOC_]]

---

:scroll: **START**


### Introduction

Nextvue CLI is a go application that allows you to easily scaffold Next Js and Nuxt Js javascript framework with a single command.

#### - LIBRARY AND DEPENDENCIES USED
 - Corbra
 - os
 - filepath

#### - COMMANDS

    ``` nextvue scaffold nextjs 
    ```
#### - FLAGS
 ```
  -d, --directory string   Directory to scaffold into
  -h, --help              help for scaffold
  -n, --name string       Project name
 ```
#### - SCAFFOLDING A NEXT JS APPLICATION
``` nextvue scaffold nextjs --directory=Pictures/Ernest/GO/nextvue --name=my-nextjs-app```

```
    Kindly take a cofee while your nextjs app is scaffolding...
    Nextvue -):  Creating a new Next.js app in C:\Users\Ernest\Pictures\Ernest\GO\nextvue\my-nextjs-app.

    Using yarn.

    Installing dependencies:
    - react
    - react-dom
    - next

    yarn add v1.22.19
    info No lockfile found.
    [1/4] Resolving packages...
    [2/4] Fetching packages...
    [3/4] Linking dependencies...
    [4/4] Building fresh packages...
    success Saved lockfile.
    success Saved 16 new dependencies.
    info Direct dependencies
    ├─ next@13.1.2
    ├─ react-dom@18.2.0
    └─ react@18.2.0
    info All dependencies
    ├─ @next/env@13.1.2
    ├─ @next/swc-win32-x64-msvc@13.1.2
    ├─ @swc/helpers@0.4.14
    ├─ caniuse-lite@1.0.30001444
    ├─ client-only@0.0.1
    ├─ js-tokens@4.0.0
    ├─ nanoid@3.3.4
    ├─ next@13.1.2
    ├─ picocolors@1.0.0
    ├─ postcss@8.4.14
    ├─ react-dom@18.2.0
    ├─ react@18.2.0
    ├─ scheduler@0.23.0
    ├─ source-map-js@1.0.2
    ├─ styled-jsx@5.1.1
    └─ tslib@2.4.1
    Done in 70.19s.

    Installing devDependencies:
    - eslint
    - eslint-config-next

    yarn add v1.22.19
    [1/4] Resolving packages...
    [2/4] Fetching packages...
    [3/4] Linking dependencies...
    [4/4] Building fresh packages...
    success Saved lockfile.
    success Saved 177 new dependencies.
    info Direct dependencies
    ├─ eslint-config-next@13.1.2
    └─ eslint@8.31.0
    info All dependencies
    ├─ @babel/runtime@7.20.7
    ├─ @eslint/eslintrc@1.4.1
    ├─ @humanwhocodes/config-array@0.11.8
    ├─ @humanwhocodes/module-importer@1.0.1
    ├─ @humanwhocodes/object-schema@1.2.1
    ├─ @next/eslint-plugin-next@13.1.2
    ├─ @nodelib/fs.scandir@2.1.5
    ├─ @nodelib/fs.stat@2.0.5
    ├─ @nodelib/fs.walk@1.2.8
    ├─ @pkgr/utils@2.3.1
    ├─ @rushstack/eslint-patch@1.2.0
    ├─ @types/json5@0.0.29
    ├─ @typescript-eslint/parser@5.48.1
    ├─ @typescript-eslint/scope-manager@5.48.1
    ├─ @typescript-eslint/typescript-estree@5.48.1
    ├─ acorn-jsx@5.3.2
    ├─ acorn@8.8.1
    ├─ ajv@6.12.6
    ├─ ansi-regex@5.0.1
    ├─ ansi-styles@4.3.0
    ├─ argparse@2.0.1
    ├─ aria-query@5.1.3
    ├─ array-includes@3.1.6
    ├─ array-union@2.1.0
    ├─ array.prototype.flat@1.3.1
    ├─ array.prototype.flatmap@1.3.1
    ├─ array.prototype.tosorted@1.1.1
    ├─ ast-types-flow@0.0.7
    ├─ axe-core@4.6.2
    ├─ axobject-query@3.1.1
    ├─ balanced-match@1.0.2
    ├─ brace-expansion@1.1.11
    ├─ braces@3.0.2
    ├─ callsites@3.1.0
    ├─ chalk@4.1.2
    ├─ color-convert@2.0.1
    ├─ color-name@1.1.4
    ├─ concat-map@0.0.1
    ├─ cross-spawn@7.0.3
    ├─ damerau-levenshtein@1.0.8
    ├─ deep-is@0.1.4
    ├─ define-lazy-prop@2.0.0
    ├─ emoji-regex@9.2.2
    ├─ enhanced-resolve@5.12.0
    ├─ es-get-iterator@1.1.3
    ├─ es-set-tostringtag@2.0.1
    ├─ es-to-primitive@1.2.1
    ├─ escape-string-regexp@4.0.0
    ├─ eslint-config-next@13.1.2
    ├─ eslint-import-resolver-node@0.3.7
    ├─ eslint-import-resolver-typescript@3.5.3
    ├─ eslint-module-utils@2.7.4
    ├─ eslint-plugin-import@2.27.4
    ├─ eslint-plugin-jsx-a11y@6.7.1
    ├─ eslint-plugin-react-hooks@4.6.0
    ├─ eslint-plugin-react@7.32.0
    ├─ eslint-scope@7.1.1
    ├─ eslint-utils@3.0.0
    ├─ eslint@8.31.0
    ├─ esquery@1.4.0
    ├─ esrecurse@4.3.0
    ├─ estraverse@5.3.0
    ├─ fast-deep-equal@3.1.3
    ├─ fast-glob@3.2.12
    ├─ fast-json-stable-stringify@2.1.0
    ├─ fast-levenshtein@2.0.6
    ├─ fastq@1.15.0
    ├─ file-entry-cache@6.0.1
    ├─ fill-range@7.0.1
    ├─ find-up@5.0.0
    ├─ flat-cache@3.0.4
    ├─ flatted@3.2.7
    ├─ function.prototype.name@1.1.5
    ├─ get-symbol-description@1.0.0
    ├─ get-tsconfig@4.3.0
    ├─ glob-parent@6.0.2
    ├─ glob@7.1.7
    ├─ globalthis@1.0.3
    ├─ globalyzer@0.1.0
    ├─ globby@13.1.3
    ├─ globrex@0.1.2
    ├─ graceful-fs@4.2.10
    ├─ grapheme-splitter@1.0.4
    ├─ has-bigints@1.0.2
    ├─ has-flag@4.0.0
    ├─ has-proto@1.0.1
    ├─ import-fresh@3.3.0
    ├─ imurmurhash@0.1.4
    ├─ is-bigint@1.0.4
    ├─ is-boolean-object@1.1.2
    ├─ is-callable@1.2.7
    ├─ is-core-module@2.11.0
    ├─ is-date-object@1.0.5
    ├─ is-docker@2.2.1
    ├─ is-extglob@2.1.1
    ├─ is-map@2.0.2
    ├─ is-negative-zero@2.0.2
    ├─ is-number-object@1.0.7
    ├─ is-number@7.0.0
    ├─ is-path-inside@3.0.3
    ├─ is-set@2.0.2
    ├─ is-string@1.0.7
    ├─ is-symbol@1.0.4
    ├─ is-typed-array@1.1.10
    ├─ is-weakmap@2.0.1
    ├─ is-weakref@1.0.2
    ├─ is-weakset@2.0.2
    ├─ is-wsl@2.2.0
    ├─ isexe@2.0.0
    ├─ js-sdsl@4.2.0
    ├─ json-schema-traverse@0.4.1
    ├─ json-stable-stringify-without-jsonify@1.0.1
    ├─ json5@1.0.2
    ├─ jsx-ast-utils@3.3.3
    ├─ language-subtag-registry@0.3.22
    ├─ language-tags@1.0.5
    ├─ locate-path@6.0.0
    ├─ lodash.merge@4.6.2
    ├─ lru-cache@6.0.0
    ├─ micromatch@4.0.5
    ├─ minimatch@3.1.2
    ├─ minimist@1.2.7
    ├─ ms@2.1.2
    ├─ natural-compare@1.4.0
    ├─ object-assign@4.1.1
    ├─ object-inspect@1.12.3
    ├─ object-is@1.1.5
    ├─ object.assign@4.1.4
    ├─ object.hasown@1.1.2
    ├─ open@8.4.0
    ├─ optionator@0.9.1
    ├─ p-limit@3.1.0
    ├─ p-locate@5.0.0
    ├─ parent-module@1.0.1
    ├─ path-exists@4.0.0
    ├─ path-key@3.1.1
    ├─ path-type@4.0.0
    ├─ picomatch@2.3.1
    ├─ prop-types@15.8.1
    ├─ punycode@2.2.0
    ├─ queue-microtask@1.2.3
    ├─ react-is@16.13.1
    ├─ regenerator-runtime@0.13.11
    ├─ regexpp@3.2.0
    ├─ resolve-from@4.0.0
    ├─ reusify@1.0.4
    ├─ rimraf@3.0.2
    ├─ run-parallel@1.2.0
    ├─ safe-regex-test@1.0.0
    ├─ shebang-command@2.0.0
    ├─ shebang-regex@3.0.0
    ├─ slash@4.0.0
    ├─ stop-iteration-iterator@1.0.0
    ├─ string.prototype.matchall@4.0.8
    ├─ string.prototype.trimend@1.0.6
    ├─ string.prototype.trimstart@1.0.6
    ├─ strip-ansi@6.0.1
    ├─ strip-bom@3.0.0
    ├─ strip-json-comments@3.1.1
    ├─ supports-color@7.2.0
    ├─ synckit@0.8.4
    ├─ tapable@2.2.1
    ├─ text-table@0.2.0
    ├─ tiny-glob@0.2.9
    ├─ to-regex-range@5.0.1
    ├─ tsconfig-paths@3.14.1
    ├─ tsutils@3.21.0
    ├─ type-check@0.4.0
    ├─ type-fest@0.20.2
    ├─ typed-array-length@1.0.4
    ├─ unbox-primitive@1.0.2
    ├─ uri-js@4.4.1
    ├─ which-collection@1.0.1
    ├─ which@2.0.2
    ├─ word-wrap@1.2.3
    ├─ yallist@4.0.0
    └─ yocto-queue@0.1.0
    Done in 14.62s.

    Success! Created my-nextjs-app at C:\Users\Ernest\Pictures\Ernest\GO\nextvue\my-nextjs-app

    A new version of `create-next-app` is available!
    You can update by running: yarn global add create-next-app


    <nil>
    Congratulations, your nextjs scaffolded!

```

#### - SCAFFOLDING A NUXT JS APPLICATION
``` nextvue scaffold nuxtjs --directory=Pictures/Ernest/GO/nextvue --name=my-nuxtjs-app```

---
:scroll: **END**
