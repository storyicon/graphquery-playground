<!---
    Copyright 2018 storyicon@foxmail.com
 
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
 
        http://www.apache.org/licenses/LICENSE-2.0
 
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
-->

**We are looking for contributors**! Please check the [ROADMAP](https://github.com/storyicon/graphquery/blob/master/ROADMAP.md) to see how you can help ❤️

---

# GraphQuery-PlayGround
![GraphQuery](https://raw.githubusercontent.com/storyicon/graphquery/master/docs/screenshot/graphquery.png)   

[GraphQuery](https://github.com/storyicon/graphquery) is a query language and execution engine tied to any backend service.GraphQuery-PlayGround is a web application for practicing, learning and testing GraphQuery.

GraphQuery-PlayGround is programming language independent, we provide a binary distribution package to facilitate any platform, any language developers to use

## Install   

### 1. Use binary distribution

Go to the [Release]() page, download and unzip the corresponding binary package according to your system type, run the server, and access `127.0.0.1:8558` in the browser.

![graphquery-playground](https://raw.githubusercontent.com/storyicon/graphquery-playground/master/docs/screenshot/playground.gif)


### 2. Compile

If you don't want to use the [Release version]() and want to modify the source code, you can read the following steps to compile.

```
go get github.com/storyicon/graphquery-playground
```

Find the downloaded `storyicon/graphquery-playground` in GOPATH, 

```
go build server.go
```
Execute the obtained binary package `server`, access `127.0.0.1:8558` in the browser.