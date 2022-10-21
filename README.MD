<div align="center" width="32" height="32">

![FullFuck Logo](/assets/iconlight.svg#gh-light-mode-only)
![FullFuck Logo](/assets/icondark.svg#gh-dark-mode-only)

</div>

# The FullFuck Language
[![Build](https://github.com/funnsam/FullFuck/actions/workflows/build.yml/badge.svg?branch=main&event=push)](https://github.com/funnsam/FullFuck/actions/workflows/build.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/funnsam/fullfuck/badge)](https://www.codefactor.io/repository/github/funnsam/fullfuck)

**As the name says, it's the most fucked programming language EVER**

## Command line usage
```
ffk input.ff output.urcl
```

## How fucked it is?
```
+047[+]>01D[+]>007[+]>>003[+]>04F[-]>057[+]>008[-]>003[+]>008[-]>008[-]>043[-]>
```
And this is just Hello world

## Syntax
```
+               will add one to the memory
-               will minus one to the memory
[]              will loop the content inside [] forever
0nn[]           will loop the content inside [] for nn times (nn is a hex number), also skips if memory is 0
>               will output the ASCII from the memory
0nn>            will output the memory to URCL port nn (nn is a hex number)
%port_name%>    will output the memory to URCL port portname
<               will copy ASCII input to the memory
0nn<            will copy data from URCL port nn to the memory (nn is a hex number)
%port_name%>    will copy data from URCL port portname to the memory
?i / ?I         will pop the thing in stack
?o / ?O         will push a thing to the stack
$               will get the current loop counter (highest number means first loop)
```


## The memory
The memory is an 8-bit unsigned integer

# Notes
- Logo licenced under [CC4.0](/assets/README.md)

# Nice Screenshots
![A RNG Test](/assets/RNGTest.svg)
![The Result](/assets/RNGTestResult.svg)
![A Graphic Test](/assets/Graphics.svg)
![The Result](/assets/GraphicsResult.svg)