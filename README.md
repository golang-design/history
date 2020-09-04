# A Documentary of Go

This document collects many interesting (publiclly observable) issues,
discussions, proposals, CLs, and talks in the Go development history,
which intents to offer a comprehensive reference of the Go history.

## Sources

There are many sources for digging the documents that relate to Go's
historical design. There are some of the official sources:

- [golang.org/doc](golang.org/doc)
- [blog.golang.org](blog.golang.org)
- [golang.org/pkg](golang.org/pkg)
- [dev.golang.org](dev.golang.org)
- [github.com/golang/go](github.com/golang/go)
- [github.com/golang/talks](github.com/golang/talks)
- [github.com/golang/proposal](github.com/golang/proposal)
- [github.com/golang/go/wiki](github.com/golang/go/wiki)
- [go-review.googlesource.com](go-review.googlesource.com)
- [groups.google.com/g/golang-nuts](groups.google.com/g/golang-nuts)
- [groups.google.com/g/golang-dev](groups.google.com/g/golang-dev)
- [groups.google.com/g/golang-tools](groups.google.com/g/golang-tools)

## Committers

Go is a big project that driven by a tiny group. Here are some
core committers to the project that you might interest in follow their
excellent work.

### Core Authors

The Go was created by Rob, Robert, and Ken initially because 
they were suffered by the slow C++ compiling time.
Later on, Ian joined the project
later because he showed huge interests and wrote the [gccgo](https://github.com/golang/gofrontend).
Russ is also one of the core authors of the project in the early state.
Back then, he was a newcomer at Google, and Rob invited Russ for joining the Go team
since he knew Russ from way back because of the [Plan 9](http://plan9.bell-labs.com/plan9) project.
Now, Russ is the tech leader of the whole Go team.

- Rob Pike. [Cat-V](http://genius.cat-v.org/rob-pike/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike)
- Robert Griesemer. [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
- Ken Thompson
- Ian Lance Taylor. [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor)
- Russ Cox. [Website](https://swtch.com/~rsc/), [Blog](https://research.swtch.com/), [GitHub](https://github.com/rsc), Twitter, Reddit

### Compiler/Runtime Team

Dmitry is not from the Go team but on the dynamic testing tools team from Google.
He wrote the scalable goroutine scheduler, many other performance improvements,
synchronization primitives, race detector, and blocking profiler that
related to the Go runtime.
Austin was an intern at Google who worked on the Go project in the early days
while pursuing a Ph. D. Later, he joined the Go team after his academic career.
Now, he is leading the Compiler/Runtime team for Go.

- Dmitry Vyukov. GitHub, Twitter
- Austin Clements. GitHub,
- Than McIntosh. GitHub,
- Keith Randall. GitHub
- David Chase. GitHub
- Michael Knyszek. GitHub
- Richard L. Hudson. GitHub
- Cherry Zhang. GitHub

### Library/Tools/Security/Community

- Andrew Gerrand. GitHub, [Twitter](https://twitter.com/enneff)
- Bryan C. Mills. GitHub
- Brad Fitzpatrick. GitHub, Twitter
- Daniel Martí. GitHub, Twitter
- Nigel Tao. GitHub, Twitter
- Filippo Valsorda. GitHub,
- Michael Matloob. GitHub,
- Dave Cheney. Website, GitHub,
- Sam Boyer. [GitHub](https://github.com/sdboyer), [Twitter](https://twitter.com/sdboyer)
- ...etc

## Timeline

Timeline helps you identify the point in time about a document that reflected in Go versions.

| Version | Duration |
|:--|:--|
| Go 1    | 2012.03.28 |
| Go 1.1  | 2013.05.13 |
| Go 1.2  | 2013.12.01 |
| Go 1.3  | 2014.06.18 |
| Go 1.4  | 2014.12.10 |
| Go 1.5  | 2015.08.19 |
| Go 1.6  | 2016.02.17 |
| Go 1.7  | 2016.08.15 |
| Go 1.8  | 2017.02.16 |
| Go 1.9  | 2017.08.24 |
| Go 1.10 | 2018.02.16 |
| Go 1.11 | 2018.08.24 |
| Go 1.12 | 2019.02.25 |
| Go 1.13 | 2019.09.03 |
| Go 1.14 | 2020.02.25 |
| Go 1.15 | 2020.08.11 |

- Go Release History: https://golang.org/doc/devel/release.html
- Pre-Go 1 Release History: https://golang.org/doc/devel/pre_go1.html
- Before Go 1: https://golang.org/doc/devel/weekly.html

## Interviews

- [talk/goborn](https://changelog.com/podcast/3) The Go Programming Language from Google with Rob Pike, Principal Engineer at Google and Co-creator of Go
- [talk/goborn2](https://changelog.com/podcast/100) Go Programming with Rob Pike and Andrew Gerrand
- [talk/goborn3](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer
- [talk/goborn4](https://evrone.com/rob-pike-interview) Rob Pike Interview

## Language Design

### Go 1

- [design/go0initial](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) Rob Pike, Robert Griesemer, Ken Thompson. The Go Annotated Specification.
- [design/go0semicolon](https://golang.org/s/semicolon-proposal) Rob Pike. Semicolons in Go.
- [design/go11return](https://golang.org/s/go11return) Go “Return at End of Function” Requirements. Russ Cox
- [design/go11func](https://golang.org/s/go11func) Russ Cox. Go 1.1 Function Calls.
- [design/go12nil](https://golang.org/s/go12nil) Russ Cox. Go 1.2 Field Selectors and Nil Checks.
- [design/go12slice](https://golang.org/s/go12slice) Russ Cox. Go Slice with Cap.
- golang.org/s/go13todo
- [doc/goatgoogle](https://talks.golang.org/2012/splash.article#TOC_12.) Rob Pike. Go at Google - Language Semantics.
- [doc/makego](https://talks.golang.org/2015/how-go-was-made.slide) How Go was Made. Andrew Gerrand
- [discuss/go1preview](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) Go 1 Preview. Russ Cox

### Package (1.4, 1.5)

- [design/go14internal](https://golang.org/s/go14internal) Russ Cox. Go 1.4 “Internal” Packages.
- [design/go14nopkg](https://golang.org/s/go14nopkg) Russ Cox. Go 1.4 src/pkg → src.
- [design/go14customimport](https://golang.org/s/go14customimport) Russ Cox. Go 1.4 Custom Import Path Checking.
- [design/go15vendor](https://golang.org/s/go15vendor) Russ Cox. Go 1.5 Vendor Experiment.


### Type alias (1.9)

- [design/type-alias](https://golang.org/design/18130-type-alias) Russ Cox, Robert Griesemer. Proposal: Type Aliases
- [issue/16339](https://golang.org/issue/16339) proposal: Alias declarations for Go
- [issue/18130](https://golang.org/issue/18130) all: support gradual code repair while moving a type between packages
- [talk/type-alias](https://www.youtube.com/watch?v=t-w6MyI2qlU) GopherCon 2016 - Lightning Talk: Robert Griesemer - Alias Declarations for Go, A proposal
- [talk/refactor-video](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Russ Cox. Codebase Refactoring (with help from Go). 
- [doc/refactor](https://talks.golang.org/2016/refactor.article) Russ Cox. Codebase Refactoring (with help from Go).

### Defer (1.13)

- [issue/6980](https://golang.org/issue/6980) cmd/compile: allocate some defers in stack frames
- [issue/14939](https://golang.org/issue/14939) runtime: defer is slow #14939
- [design/open-defer](https://golang.org/design/34481-opencoded-defers) Dan Scales, Keith Randall, and Austin Clements. Proposal: Low-cost defers through inline code, and extra funcdata to manage the panic case
- Unsolved `defer recover()` edge cases.
  - [issue/10458](https://golang.org/issue/10458) spec: panicking corner-case semantics
  - [issue/23531](https://golang.org/issue/23531) spec: recover() in nested defer
  - [issue/26275](https://golang.org/issue/26275) runtime: document the behaviour of Caller and Callers when used in deferred functions
  - [issue/34530](https://golang.org/issue/34530) spec: clarify when calling recover stops a panic
  - [cl/189377](https://golang.org/cl/189377) spec: specify the behavior of recover and recursive panics in more detail

### Error values (1.13)

- [issue/40432](https://golang.org/issue/40432) language: Go 2: error handling meta issue
- [design/err-handling-overview](https://golang.org/design/go2draft-error-handling-overview) Russ Cox. Error Handling — Problem Overview.
- [design/err-values-overview](https://golang.org/design/go2draft-error-values-overview) Russ Cox. Error Values — Problem Overview. 
- [design/err-handle-check](https://golang.org/design/go2draft-error-handling) Marcel van Lohuizen. Error Handling — Draft Design.
  - [design/err-try](https://golang.org/design/32437-try-builtin) Proposal: A built-in Go error check function, "try"
  - [issue/32437](https://golang.org/issue/32437#issuecomment-512035919) Proposal: A built-in Go error check function, "try". Decision response.
- [design/err-inspect](https://golang.org/design/go2draft-error-inspection) Error Inspection — Draft Design.
  - [issue/29934](https://golang.org/issue/29934) proposal: Go 2 error values.
- [design/err-print](https://golang.org/design/go2draft-error-printing) Error Printing — Draft Design.
  - [issue/30468](https://golang.org/issue/30468) errors: performance regression in New.
- [issue/41198](https://golang.org/issue/41198) proposal: errors: add ErrUnimplemented as standard way for interface method to fail.

### Map

- [talk/map](https://www.youtube.com/watch?v=Tl7mi9QmLns) GopherCon 2016: Keith Randall - Inside the Map Implementation

### Channel/Select

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov. Go channels on steroids. Jan 28, 2014
  + [issue/8899](https://github.com/golang/go/issues/8899) runtime: lock-free channels
  + [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) update on "lock-free channels"
  + [cl/112990043](https://codereview.appspot.com/112990043/) runtime: fine-grained locking in select
  + [cl/110580043](https://codereview.appspot.com/110580043/) runtime: add fast paths to non-blocking channel operations
- [issue/8898](https://github.com/golang/go/issues/8898) runtime: special case timer channels
- [issue/8903](https://golang.org/issue/8903) runtime: make chan-based generators faster. 
- [issue/21806](https://golang.org/issue/21806) fairness in select statement.

### Generics

- [issue/15292](https://golang.org/issue/15292) proposal: spec: generic programming facilities
- [doc/generics-discuss](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/edit#heading=h.vuko0u3txoew) Summary of Go Generics Discussions
- [doc/generics-dilemma](https://research.swtch.com/generic) Russ Cox. The Generic Dilemma. December 3, 2009.
- [design/type-functions](https://golang.org/design/15292/2010-06-type-functions) Ian Lance Taylor. "Type Functions." golang/proposals, June 2010.
- [design/generalized-types](https://golang.org/design/15292/2011-03-gen) Ian Lance Taylor. "Generalized Types."golang/proposals, March 2011.
- [design/code-gen](https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM) Russ Cox. "Alternatives to Dynamic Code Generation in Go." September 2012.
- [design/generalized-types2](https://golang.org/design/15292/2013-10-gen) Ian Lance Taylor. "Generalized Types In Go." golang/proposals, October 2013.
- [design/type-parameters](https://golang.org/design/15292/2013-12-type-params) Ian Lance Taylor. "Type Parameters in Go." golang/proposals, December 2013.
- [design/compile-time-function](https://golang.org/design/15292/2016-09-compile-time-functions) Bryan C. Mills. "Compile-time Functions and First Class Types." golang/proposals, September 2016.
- [design/should-generics](https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics) Ian Lance Taylor. "Go should have generics." golang/proposals, January 2011.
- [design/should-generics2](https://golang.org/design/15292-generics) Ian Lance Taylor. "Go should have generics." golang/proposals, Updated: April 2016. 
- [design/generics-overview](https://golang.org/design/go2draft-generics-overview) Russ Cox. "Generics — Problem Overview." golang/proposals, August 27, 2018. 
- [design/contracts](https://golang.org/design/go2draft-contracts) Ian Lance Taylor, Robert Griesemer. "Contracts — Draft Design." golang/proposals, August 27, 2018, Updated: July 31, 2019.
  + [cl/187317](https://golang.org/cl/187317/) Implementation prototype.
- [design/type-parameters2](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md) Ian Lance Taylor, Robert Griesemer. "Type Parameters - Draft Design ." golang/proposals, June 16, 2020, Updated: August 28, 2020.
  + [cl/dev.go2go](https://github.com/golang/go/blob/dev.go2go/README.go2go.md) dev.go2go branch
  + [doc/type-check-readme](https://github.com/golang/go/tree/dev.go2go/src/go/types) type checking.
  + [doc/type-check-notes](https://github.com/golang/go/blob/dev.go2go/src/go/types/NOTES) This file serves as a notebook/implementation log.
- [discuss/generics-parenthesis](https://groups.google.com/g/golang-nuts/c/7t-Q2vt60J8) Robert. Generics and parenthesis.
- [discuss/generics-move](https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw) Ian, Moving forward with the generics design.
- [talk/generics-gotime](https://changelog.com/gotime/98) Generics in Go
with Ian Lance Taylor
- [talk/generics-gotime2](https://changelog.com/gotime/140) The latest on Generics
with Robert Griesemer and Ian Lance Taylor
- [talk/generics-cppcast](https://cppcast.com/ian-taylor-go) CppCast. Go with Ian Lance Taylor
- [paper/featherweight-go](https://arxiv.org/abs/2005.11710) Griesemer, Robert, et al. "Featherweight Go." arXiv preprint arXiv:2005.11710 (2020).
- [talk/featherweight-go](https://www.youtube.com/watch?v=Dq0WFigax_c) Phil Wadler: Featherweight Go. Jun 8, 2020.

## Go 2

- [talk/go2spec](https://www.youtube.com/watch?v=RIvL2ONhFBI) Sydney Golang Meetup - Rob Pike - Go 2 Draft Specifications

## Compiler Toolchain

### Compiler

- [design/go12symtab](https://golang.org/s/go12symtab) Russ Cox. Go 1.2 Runtime Symbol Information. July 2013.
- [design/go13compiler](https://golang.org/s/go13compiler) Russ Cox. Go 1.3+ Compiler Overhaul. December 2013
- [design/go14generate](https://golang.org/s/go1.4-generate) Rob Pike. Go generate: A Proposal
- [design/dev.cc](https://golang.org/s/dev.cc)  Russ Cox. dev.cc branch plan. November 2014
- [talk/gobootstrap](https://www.youtube.com/watch?v=QIE5nV5fDwA) GopherCon 2014 Go from C to Go by Russ Cox.
- [design/go15bootstrap](https://golang.org/s/go15bootstrap) Russ Cox. Go 1.5 Bootstrap Plan. January 2015.
- [design/execmodes](https://golang.org/s/execmodes) Ian Lance Taylor. Go Execution Modes. August, 2014 (updated January 2016)
- [design/go17ssa](https://golang.org/s/go17ssa) Keith Randall. New SSA Backend for the Go Compiler. 2/10/2015.
- [issue/6853](https://golang.org/issue/6853) all: binaries too big and growing.
- [issue/19348](https://golang.org/issue/19348) cmd/compile: enable mid-stack inlining.
- [issue/23109](https://golang.org/issue/23109) cmd/compile: rewrite escape analysis.
- [issue/27539](https://golang.org/issue/27539) proposal: runtime: make the ABI undefined as a step toward changing it.
- proposal: add GOEXPERIMENT=checkptr
  + [issue/22218](https://golang.org/issue/22218) proposal: add GOEXPERIMENT=checkptr
  + [issue/34964](https://golang.org/issue/34964) cmd/compile: enable -d=checkptr as part of -race and/or -msan?
  + [issue/34972](https://golang.org/issue/34972) all: get standard library building with -d=checkptr
  + [discuss/checkptr](https://groups.google.com/forum/#!msg/golang-dev/SzwDoqoRVJA/Iozu8vWdDwAJ)
- [issue/8885](https://golang.org/issue/8885) runtime: consider adding 24-byte size class. 
- [issue/16798](https://golang.org/issue/16798) proposal: cmd/compile: add tail call optimization for self-recursion only. 
- alignment
  + [issue/599](https://golang.org/issue/599) cmd/compile: make 64-bit fields 64-bit aligned on 32-bit systems
  + [issue/36606](https://golang.org/issue/36606) proposal: cmd/compile: make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed directive on structs
  + [design/64align](https://golang.org/design/36606-64-bit-field-alignment) Dan Scales. Proposal: Make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed, //go:align directives. 2020-06-08.

### Linker

- [design/go13linker](https://golang.org/s/go13linker) Russ Cox. Go 1.3 Linker Overhaul. November 2013.
- [design/go116linker](https://golang.org/s/better-linker) Austin Clements. Building a better Go linker. 2019-09-12.

### Debugger

- [design/go13heapdump](https://golang.org/s/go13heapdump) heapdump13
- [design/go14heapdump](https://github.com/golang/go/wiki/heapdump14) heapdump14
- [design/go15heapdump](https://github.com/golang/go/wiki/heapdump15-through-heapdump17) heapdump15 through heapdump17
- [design/go15trace](https://golang.org/s/go15trace) Dmitry Vyukov. Go Execution Tracer.
- https://github.com/golang/go/wiki/heapdump15-through-heapdump17

### Tracer

- [design/tracer](https://docs.google.com/document/u/1/d/1FP5apqzBgr7ahCCgFO-yoVhk4YZrNIDNf9RybngBc14/pub) Dmitry Vyukov. Go Execution Tracer. Oct 2014.
- [design/trace](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@. Trace Event Format. October 2016.

### Builder

- [design/go13nacl](golang.org/s/go13nacl) Go 1.3 Native Client Support
- [design/go14android](https://golang.org/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/go-generate](https://golang.org/s/go1.4-generate) Rob Pike. "Go Generate." January 2014.
- [issue/13560](https://golang.org/issue/13560) proposal: build: define standard way to recognize machine-generated files
- [discuss/generatedcode](golang.org/s/generatedcode) Rob Pike's Final Comments on Issue 13560
- [design/go116build](https://golang.org/design/draft-gobuild) Russ Cox. Bug-resistant build constraints — Draft Design. June 30, 2020.
- [design/go116embed](https://golang.org/design/draft-embed) Embedded files - Russ & Braid
- Windows
  + [discuss/win2000-golang-nuts](https://golang.org/s/win2000-golang-nuts) objections to removing Go support for Windows 2000 (x86-32)?
- [design/wasm](https://docs.google.com/document/d/131vjr4DH6JFnb-blm_uRdaC0_Nv3OUwjEY5qVCxCup4/edit#heading=h.mjo1bish3xni) Richard Musiol. WebAssembly architecture for Go. 28th February 2018.
- [design/wasm2](https://docs.google.com/document/d/1GRmy3rA4DiYtBlX-I1Jr_iHykbX8EixC3Mq0TCYqbKc/edit#heading=h.q4c21ihutzk0) WebAssembly assembly files

### Modules

- [design/go-dep](https://docs.google.com/document/d/18tNd8r5DV0yluCR7tPvkMTsWD_lYcRO7NhpNSDymRr8) Go Packaging Proposal Process
- [design/go-dep2](https://docs.google.com/document/d/1qnmjwfMmvSCDaY4jxPmLAccaaUI5FfySNE90gB0pTKQ/edit) Dependency Management Tool
- [doc/go-dep](6https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) Sam Boyer. The Saga of Go Dependency Management. Dec 13, 201
- [talk/go-dep](https://www.youtube.com/watch?v=5LtMb090AZI) GopherCon 2017: Sam Boyer - The New Era of Go Package Management. Jul 24, 2017
- [talk/go-dep-design](https://www.youtube.com/watch?v=wBTGd_dvnO8) dotGo 2017 - sam boyer - The Functional Design of Dep. Dec 8, 2017
- [discuss/go-dep](https://www.youtube.com/watch?v=sbrZfPgNmfw) Building Predictability into Your Pipeline :: Russ Cox; Jess Frazelle, Sam Boyer, Pete Garcin. Feb 22, 2018
- [doc/go-dep-response](https://peter.bourgon.org/blog/2018/07/27/a-response-about-dep-and-vgo.html) Peter Bourgon. A response about dep and vgo. 2018-07-27.
- [discuss/go-dep-response](https://www.reddit.com/r/golang/comments/92f3q1/peter_bourgon_a_response_about_dep_and_vgo/) Reddit discussion.
- [discuss/go-dep-response2](https://news.ycombinator.com/item?id=17628311) Hacker News discussion.
- [discuss/go-dep-twitter](https://twitter.com/_rsc/status/1022588240501661696) Russ Cox's Twitter Storm
- [talk/gotime-gomod](https://changelog.com/gotime/77) Dependencies and the future of Go with Russ Cox
- [discuss/groups-gomod](https://groups.google.com/g/golang-dev/c/a5PqQuBljF4) go modules have landed
- [doc/deps](https://research.swtch.com/deps) Russ Cox. Our Software Dependency Problem. January 23, 2019.
- [doc/vgo](https://research.swtch.com/vgo) Russ Cox. Go & Versioning
- [issue/24301](https://golang.org/issue/24301) cmd/go: add package version support to Go toolchain
- [design/lazy-gomod](https://golang.org/design/36460-lazy-module-loading) Bryan C. Mills. Proposal: Lazy Module Loading. 2020-02-20

### Testing

- Tool chain, benchseries/benchstat
- [design/fuzzing](https://golang.org/design/draft-fuzzing) Katie Hockman. Design Draft: First Class Fuzzing

## Runtime Core

### Statistics

- [issue/29696](https://github.com/golang/go/issues/29696) proposal: runtime: add way for applications to respond to GC backpressure
- [issue/16843](https://golang.org/issue/16843) runtime: mechanism for monitoring heap size
- [cl/setmaxheap](https://go-review.googlesource.com/c/go/+/46751/) Austin Clements. runtime/debug: add SetMaxHeap API. Jun 26 2017.
- [issue/19812](https://golang.org/issue/19812) runtime: cannot ReadMemStats during GC
- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek. Proposal: API for unstable runtime metrics

### Scheduler

- [design/go11sched](https://golang.org/s/go11sched) Dmitry Vyukov. Scalable Go Scheduler Design Doc, 2012]
- [design/preempt-sched](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#heading=h.3pilqarbrc9h) Dmitry Vyukov. Go Preemptive Scheduler Design Doc, 2013
- [design/go15gomaxprocs](https://golang.org/s/go15gomaxprocs) Russ Cox. Go 1.5 GOMAXPROCS Default. May 2015.
- [design/preempt-sched2](https://golang.org/design/24543-non-cooperative-preemption) Austin Clements. Proposal: Non-cooperative goroutine preemption. 2019-01-18.
  + [issue/10958](https://golang.org/issue/10958) runtime: tight loops should be preemptible
  + [issue/24543](https://golang.org/issue/24543) runtime: non-cooperative goroutine preemption
  + [issue/36365](https://github.com/golang/go/issues/36365) runtime: clean up async preemption loose ends
- [issue/18237](https://golang.org/issue/18237) runtime: scheduler is slow when goroutines are frequently woken
- [issue/27345](https://golang.org/issue/27345) runtime: use parent goroutine's stack for new goroutines
- [issue/32113](https://golang.org/issue/32113) runtime: optimization to reduce P churn.
  + [cl/deferWake](https://github.com/amscanne/go/commit/eee812b594577f71894fd30a27d9a39ba99bf590) Add deferWake to channel ops

### Execution Stack

- [design/contigstack](https://golang.org/s/contigstacks) Contiguous stacks
- [issue/26061](https://github.com/golang/go/issues/26061) runtime: g0 stack.lo is sometimes too low

### Memory Allocator

A quick history about the Go's memory allocator: Russ Cox first implements the memory allocator based on`tcmalloc` for Go 1, `mcache` are cached on M. Then he revised the allocator to allow
user code use 16GB memory and later allows 128GB.

- [issue/30333](https://golang.org/issue/30333) runtime: smarter scavenging
- [issue/34047](https://golang.org/issue/34047) runtime: potential deadlock cycle caused by scavenge.lock
- [issue/34048](https://golang.org/issue/34048) runtime: scavenger pacing fails to account for fragmentation
- [issue/35112](https://golang.org/issue/35112) runtime: make the page allocator scale
- [issue/35788](https://golang.org/issue/35788) runtime: scavenger not as effective as in previous releases
- [design/go13scavenge](https://go.googlesource.com/proposal/+/aa701aae530695d32916b779e048a3e18311a2e3/design/30333-smarter-scavenging.md) Michael Knyszek. Proposal: Smarter Scavenging. 2019-05-09.
- [design/go14pagealloc](https://go.googlesource.com/proposal/+/a078ea9d72b99dc88fdfd2cb6ee150a8ce202ea2/design/35112-scaling-the-page-allocator.md) Michael Knyszek, Austin Clements. Proposal: Scaling the Go page allocator. 2019-10-18.
- [issue/37487](https://golang.org/issue/37487) runtime: improve mcentral scalability
- [cl/221182](https://golang.org/cl/221182) runtime: add new mcentral implementation
- [issue/18155](https://golang.org/issue/18155) runtime: latency in sweep assists when there's no garbage
- [issue/19112](https://golang.org/issue/19112) runtime: deadlock involving gcControllerState.enlistWorker
- [issue/19812](https://golang.org/issue/19812) runtime: cannot ReadMemStats during GC
- [issue/29707](https://golang.org/issue/29707) cmd/trace: "failed to parse trace: no consistent ordering of events possible"
- [issue/35788](https://golang.org/issue/35788) runtime: scavenger not as effective as in previous releases
- [issue/35954](https://golang.org/issue/35954) runtime: handle hitting the top of the address space in the allocator more gracefully
- [issue/37112](https://golang.org/issue/37112) runtime: API for unstable metrics
- [issue/37487](https://golang.org/issue/37487) runtime: improve mcentral scalability
- [issue/37927](https://golang.org/issue/37927) runtime: GC pacing exhibits strange behavior with a low GOGC
- [issue/38070](https://golang.org/issue/38070) runtime: timer self-deadlock due to preemption point
- [issue/38130](https://golang.org/issue/38130) runtime: incorrect sanity checks in the page allocator
- [issue/38404](https://golang.org/issue/38404) runtime: STW GC stops working on arm64/mips64le
- [issue/38605](https://golang.org/issue/38605) runtime: pageAlloc.allocToCache updates pageAlloc.searchAddr in an invalid way
- [issue/38617](https://golang.org/issue/38617) runtime: scavenger freezes up in Go 1.14 in Windows due to coarse time granularity
- [issue/38712](https://golang.org/issue/38712) runtime: TestMemStats is flaky
- [issue/38966](https://golang.org/issue/38966) runtime: aix-ppc64 builder is failing with "bad prune", "addr range base and limit are not in the same memory segment" fatal errors
- [issue/39128](https://golang.org/issue/39128) runtime: linux-mips-rtrk builder consistently failing with "bad prune" as of CL 233497
- [issue/40191](https://golang.org/issue/40191) runtime: pageAlloc.searchAddr may point to unmapped memory in discontiguous heaps, violating its invariant
- [issue/40457](https://golang.org/issue/40457) runtime: runqputbatch does not protect its call to globrunqputbatch
- [issue/40459](https://golang.org/issue/40459) runtime: ReadMemStats called in a loop may prevent GC
- [issue/40641](https://golang.org/issue/40641) runtime: race between stack shrinking and channel send/recv leads to bad sudog values

### Garbage Collector

- [design/go13gc](https://docs.google.com/document/d/1v4Oqa0WwHunqlb8C3ObL_uNQw3DfSY-ztoA-4wWbKcg/pub) Dmitry Vyukov. Simpler and faster GC for Go. July 16, 2014
  + [cl/106260045](https://codereview.appspot.com/106260045) runtime: simpler and faster GC
- [design/go14gc](https://golang.org/s/go14gc) Richard L. Hudson. Go 1.4+ Garbage Collection (GC) Plan and Roadmap. August 6, 2014.
- [design/go15gcpacing](https://golang.org/s/go15gcpacing) Austin Clements. Go 1.5 concurrent garbage collector pacing. 2015-03-10.
- [talk/ismm-gc](https://blog.golang.org/ismmkeynote) Rick Hudson. Getting to Go: The Journey of Go's Garbage Collector. 12 July 2018.
- [discuss/ismm-gc](https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns) Garbage Collection Slides and Transcript now available
- [cl/generational-gc](https://golang.org/cl/137482) runtime: trigger generational GC
- [design/roc](https://golang.org/s/gctoc) Request Oriented Collector (ROC) Algorithm
- [cl/roc](https://golang.org/cl/25058) runtime: ROC write barrier code
- [issue/17503](https://golang.org/issue/17503) runtime: eliminate stack rescanning
- [issue/22350](https://golang.org/issue/22350) cmd/compile: compiler can unexpectedly preserve memory, 
- [design/gcscan](https://docs.google.com/document/d/1un-Jn47yByHL7I0aVIP_uVCMxjdM5mpelJhiKlIqxkE/edit#) Proposal: GC scanning of stacks
- [issue/26903](https://golang.org/issue/26903) runtime: simplify mark termination and eliminate mark 2
- [issue/27993](https://golang.org/issue/27993) runtime: error message: P has cached GC work at end of mark termination

### Memory model

Go memory model is not properly defined.

- [issue/5045](https://golang.org/issue/5045) doc: define how sync/atomic interacts with memory model
- [issue/9442](https://golang.org/issue/9442) doc: define how finalizers interact with memory model
- [issue/7948](https://golang.org/issue/7948) doc: define how sync interacts with memory model
- [issue/33815](https://golang.org/issue/33815) doc/go_mem: "hello, world" will not always be printed twice
- [talk/gomem](http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf) Russ Cox. Go’s Memory Model. February 25, 2016.
- [doc/go2017russ](https://research.swtch.com/go2017#memory) Russ Cox. My Go Resolutions for 2017 - Memory model. January 18, 2017.
- [doc/atomic-bug](https://golang.org/pkg/sync/atomic/#pkg-note-BUG) Package atomic
- [discuss/atomic-mem-order](https://groups.google.com/d/msg/golang-dev/vVkH_9fl1D8/azJa10lkAwAJ) specify the memory order guarantee provided by atomic Load/Store

### ABI

- [design/register-call](https://golang.org/design/40724-register-calling) Austin Clements, with input from Cherry Zhang, Michael Knyszek, Martin Möhrmann, Michael Pratt, David Chase, Keith Randall, Dan Scales, and Ian Lance Taylor. Proposal: Register-based Go calling convention. 2020-08-10.
- [issue/18597](https://github.com/golang/go/issues/18597) proposal: cmd/compile: define register-based calling convention
- [issue/40724](https://github.com/golang/go/issues/40724) proposal: switch to a register-based calling convention for Go functions

## Standard Library

### encoding

- [design/go12encoding](https://golang.org/s/go12encoding) Russ Cox. Go 1.2 encoding.TextMarshaler and TextUnmarshaler. July 2013.
- [design/go12xml](https://docs.google.com/document/d/1G-AMeUPoyTnbZDkuCJA89DjJwTjdWFctIZ_N9NgA9RM/pub) Russ Cox. Go 1.2 xml.Marshaler and Unmarshaler. July 2013.

### syscall

- [design/go14syscall](https://golang.org/s/go1.4-syscall) The syscall package.

### go/types

- [doc/gotypes](https://golang.org/s/types-tutorial) Alan Donovan. go/types: The Go Type Checker.
- [talk/gotypes](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view) Alan Donovan. Using go/types for
Code Comprehension and Refactoring Tools. October 2, 2015.

### sync

- [issue/22950](https://golang.org/issue/22950) sync: avoid clearing the full Pool on every GC.
- [cl/4631059](https://github.com/golang/go/commit/997c00f) runtime: replace Semacquire/Semrelease implementation.
- [issue/13086](https://golang.org/issue/13086) runtime: fall back to fair locks after repeated sleep-acquire failures. 
- [issue/21035](https://github.com/golang/go/issues/21035) sync: reduce contention between Map operations with new-but-disjoint keys

### time

- [issue/6239](https://golang.org/issue/6239) runtime: make timers faster.
- [issue/27707](https://golang.org/issue/27707) time: excessive CPU usage when using Ticker and Sleep.

### io

- [design/draft-iofs](https://golang.org/design/draft-iofs) Russ Cox, Rob Pike. File System Interfaces for Go — Draft Design
. July 2020.

### context

- [doc/context](https://blog.golang.org/context) Go Concurrency Patterns: Context.

## Unclassified yet

- https://golang.org/s/using-guru
- https://golang.org/s/gogetcmd
- https://golang.org/s/stdwhy
- https://golang.org/s/oracle-design
- https://golang.org/s/oracle-user-manual
- https://golang.org/s/cgihttpproxy
- golang.org/s/sqldrivers
- https://golang.org/s/sometext
- https://www.cflee.com/posts/golang-org-url-redirector/
- https://golang.org/s/go2designs
- https://golang.org/s/http2bug
- https://docs.google.com/document/d/19_ExiylD9MRfeAjKIfEsMU1_RGhuxB9sA0b5Zv7byVI/edit
- https://docs.google.com/spreadsheets/d/1VLxi-ac0BAtf735HyH3c1xRulbkYYUkFecKdLPH7NIQ/edit#gid=166102500


<!-- ?Late call - david chase  -->

## Contributing

PRs are welcome.

## License

[go-history](https://github.com/changkun/go-history) CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)