# The Go History

This document includes many interesting issues, discussions, proposals, CLs, and talks in the Go development history, which intents to offer 
a comprehensive reference of the Go history.

## Sources

- [golang.org/doc](golang.org/doc)
- [blog.golang.org](blog.golang.org)
- [golang.org/pkg](golang.org/pkg)
- [dev.golang.org](dev.golang.org)
- [github.com/golang/go](github.com/golang/go)
- [github.com/golang/proposal](github.com/golang/proposal)
- [github.com/golang/go/wiki](github.com/golang/go/wiki)
- [go-review.googlesource.com](go-review.googlesource.com)
- [groups.google.com/g/golang-nuts](groups.google.com/g/golang-nuts)
- [groups.google.com/g/golang-dev](groups.google.com/g/golang-dev)
- [groups.google.com/g/golang-tools](groups.google.com/g/golang-tools)

## Committers

### Core Authors

- Rob Pike. [Cat-V](http://genius.cat-v.org/rob-pike/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike)
- Robert Griesemer. [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
- Ken Thompson
- Russ Cox. GitHub, Twitter, Reddit
- Ian Lance Taylor. [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor)

### Compiler/Runtime Team

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
- Filippo Valsorda
- Michael Matloob
- Dave Cheney
- ...etc

## Interviews


- [talk/goborn](https://changelog.com/podcast/3) The Go Programming Language from Google with Rob Pike, Principal Engineer at Google and Co-creator of Go
- [talk/goborn2](https://changelog.com/podcast/100) Go Programming with Rob Pike and Andrew Gerrand
- [talk/goborn3](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer
- [talk/goborn4](https://evrone.com/rob-pike-interview) Rob Pike Interview

## Language Design

### Before Go 1

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

- [design/type-alias](https://go.googlesource.com/proposal/+/master/design/18130-type-alias.md) Russ Cox, Robert Griesemer. Proposal: Type Aliases
- [issue/16339](https://golang.org/issue/16339) proposal: Alias declarations for Go
- [issue/18130](https://golang.org/issue/18130) all: support gradual code repair while moving a type between packages
- [talk/type-alias](https://www.youtube.com/watch?v=t-w6MyI2qlU) GopherCon 2016 - Lightning Talk: Robert Griesemer - Alias Declarations for Go, A proposal
- [talk/refactor-video](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Russ Cox. Codebase Refactoring (with help from Go). 
- [doc/refactor](https://talks.golang.org/2016/refactor.article) Russ Cox. Codebase Refactoring (with help from Go).

### Defer (1.13)

- [issue/6980](https://golang.org/issue/6980) cmd/compile: allocate some defers in stack frames
- [issue/14939](https://golang.org/issue/14939) runtime: defer is slow #14939
- [design/open-defer](https://github.com/golang/proposal/blob/master/design/34481-opencoded-defers.md) Dan Scales, Keith Randall, and Austin Clements. Proposal: Low-cost defers through inline code, and extra funcdata to manage the panic case
- Unsolved `defer recover()` edge cases.
  - [issue/10458](https://golang.org/issue/10458) spec: panicking corner-case semantics
  - [issue/23531](https://golang.org/issue/23531) spec: recover() in nested defer
  - [issue/26275](https://golang.org/issue/26275) runtime: document the behaviour of Caller and Callers when used in deferred functions
  - [issue/34530](https://golang.org/issue/34530) spec: clarify when calling recover stops a panic
  - [cl/189377](https://golang.org/cl/189377) spec: specify the behavior of recover and recursive panics in more detail

### Error values (1.13)

- [issue/40432](https://golang.org/issue/40432) language: Go 2: error handling meta issue
- [design/err-handling-overview](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling-overview.md) Russ Cox. Error Handling — Problem Overview.
- [design/err-values-overview](https://github.com/golang/proposal/blob/master/design/go2draft-error-values-overview.md) Russ Cox. Error Values — Problem Overview. 
- [design/err-handle-check](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling.md) Marcel van Lohuizen. Error Handling — Draft Design.
  - [design/err-try](https://github.com/golang/proposal/blob/master/design/32437-try-builtin.md) Proposal: A built-in Go error check function, "try"
  - [issue/32437](https://golang.org/issue/32437#issuecomment-512035919) Proposal: A built-in Go error check function, "try". Decision response.
- [design/err-inspect](https://github.com/golang/proposal/blob/master/design/go2draft-error-inspection.md) Error Inspection — Draft Design.
  - [issue/29934](https://golang.org/issue/29934) proposal: Go 2 error values.
- [design/err-print](https://github.com/golang/proposal/blob/master/design/go2draft-error-printing.md) Error Printing — Draft Design.
  - [issue/30468](https://golang.org/issue/30468) errors: performance regression in New.
- [issue/41198](https://golang.org/issue/41198) proposal: errors: add ErrUnimplemented as standard way for interface method to fail. 

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
- [design/type-functions](https://github.com/golang/proposal/blob/master/design/15292/2010-06-type-functions.md) Ian Lance Taylor. "Type Functions." golang/proposals, June 2010.
- [design/generalized-types](https://github.com/golang/proposal/blob/master/design/15292/2011-03-gen.md) Ian Lance Taylor. "Generalized Types."golang/proposals, March 2011.
- [design/code-gen](https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM) Russ Cox. "Alternatives to Dynamic Code Generation in Go." September 2012.
- [design/generalized-types2](https://github.com/golang/proposal/blob/master/design/15292/2013-10-gen.md) Ian Lance Taylor. "Generalized Types In Go." golang/proposals, October 2013.
- [design/type-parameters](https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md) Ian Lance Taylor. "Type Parameters in Go." golang/proposals, December 2013.
- [design/go-generate](http://golang.org/s/go1.4-generate) Rob Pike. "Go Generate." January 2014.
- [design/compile-time-function](https://github.com/golang/proposal/blob/master/design/15292/2016-09-compile-time-functions.md) Bryan C. Mills. "Compile-time Functions and First Class Types." golang/proposals, September 2016.
- [design/should-generics](https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics.md) Ian Lance Taylor. "Go should have generics." golang/proposals, January 2011.
- [design/should-generics2](https://github.com/golang/proposal/blob/master/design/15292-generics.md) Ian Lance Taylor. "Go should have generics." golang/proposals, Updated: April 2016. 
- [design/generics-overview](https://github.com/golang/proposal/blob/master/design/go2draft-generics-overview.md) Russ Cox. "Generics — Problem Overview." golang/proposals, August 27, 2018. 
- [design/contracts](https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md) Ian Lance Taylor, Robert Griesemer. "Contracts — Draft Design." golang/proposals, August 27, 2018, Updated: July 31, 2019.
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

## Compiler Toolchain

### Compiler


- https://golang.org/s/go12symtab
- https://golang.org/s/go13compiler
- https://golang.org/s/go1.4-generate
- https://golang.org/s/go15bootstrap
  - C to Go bootstraping, https://www.youtube.com/watch?v=QIE5nV5fDwA
- https://golang.org/s/execmodes
- https://golang.org/s/go17ssa
- all: binaries too big and growing. [issue/6853](https://golang.org/issue/6853)
- cmd/compile: enable mid-stack inlining. [issue/19348](https://golang.org/issue/19348)
- cmd/compile: rewrite escape analysis. [issue/23109](https://golang.org/issue/23109)
- proposal: runtime: make the ABI undefined as a step toward changing it. [issue/27539](https://golang.org/issue/27539)
- proposal: add GOEXPERIMENT=checkptr [issue/22218](https://golang.org/issue/22218), [issue/34964](https://golang.org/issue/34964), [issue/34972](https://golang.org/issue/34972), [discuss/checkptr](https://groups.google.com/forum/#!msg/golang-dev/SzwDoqoRVJA/Iozu8vWdDwAJ)
- runtime: consider adding 24-byte size class. [issue/8885](https://golang.org/issue/8885)
- proposal: cmd/compile: add tail call optimization for self-recursion only. [issue/16798](https://golang.org/issue/16798)
- alignment. [issue/599](https://golang.org/issue/599), [issue/36606](https://golang.org/issue/36606), [design/64align](https://github.com/golang/proposal/blob/master/design/36606-64-bit-field-alignment.md)

### Linker

- [design/go13linker](https://golang.org/s/go13linker) Russ Cox. Go 1.3 Linker Overhaul. November 2013.
- [design/go116linker](https://golang.org/s/better-linker) Austin Clements. Building a better Go linker. 2019-09-12.

### Debugger

- [design/go15trace](https://golang.org/s/go15trace) Dmitry Vyukov. Go Execution Tracer.
- https://github.com/golang/go/wiki/heapdump15-through-heapdump17

### Builder

- [design/go13nacl](golang.org/s/go13nacl) Go 1.3 Native Client Support
- [design/go14android](https://golang.org/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/go116build](https://github.com/golang/proposal/blob/master/design/draft-gobuild.md) Russ Cox. Bug-resistant build constraints — Draft Design. June 30, 2020.
- [design/go116embed](https://github.oom/golang/proposal/blob/master/design/draft-embed.md) Embedded files - Russ & Braid

### Modules

- https://semver.org/
- https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/
- https://www.youtube.com/watch?v=5LtMb090AZI
- https://www.youtube.com/watch?v=wBTGd_dvnO8
- https://www.youtube.com/watch?v=sbrZfPgNmfw
- https://zhuanlan.zhihu.com/p/41627929
- https://peter.bourgon.org/blog/2018/07/27/a-response-about-dep-and-vgo.html
- https://news.ycombinator.com/item?id=17628311
- https://www.reddit.com/r/golang/comments/92f3q1/peter_bourgon_a_response_about_dep_and_vgo/
- https://twitter.com/_rsc/status/1022588240501661696
- https://changelog.com/gotime/77
- https://groups.google.com/forum/#!msg/golang-dev/a5PqQuBljF4/scQU-TfXBwAJ
- https://research.swtch.com/deps
- Lazy module load - Bryan https://github.com/golang/proposal/blob/master/design/36460-lazy-module-loading.md

### Testing

- Tool chain, benchseries/benchstat
- Fuzzing https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md

## Runtime Core

### Statistics

- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek. Proposal: API for unstable runtime metrics

### Scheduler

- [VYUKOV, 2012] [Vyukov, Dmitry. Scalable Go Scheduler Design Doc, 2012](https://golang.org/s/go11sched)
- [VYUKOV, 2013] [Vyukov, Dmitry. Go Preemptive Scheduler Design Doc, 2013](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#heading=h.3pilqarbrc9h)
- Go 1.5 GOMAXPROCS Default, https://golang.org/s/go15gomaxprocs
- runtime: tight loops should be preemptible, [#10958](https://golang.org/issue/10958)
- runtime: non-cooperative goroutine preemption, [#24543](https://golang.org/issue/24543)
- [issue/18237](https://golang.org/issue/18237) runtime: scheduler is slow when goroutines are frequently woken
- [issue/27345](https://golang.org/issue/27345) runtime: use parent goroutine's stack for new goroutines
- [issue/32113](https://golang.org/issue/32113) runtime: optimization to reduce P churn.
  + [cl/deferWake](https://github.com/amscanne/go/commit/eee812b594577f71894fd30a27d9a39ba99bf590) Add deferWake to channel ops

### Execution Stack

- Contiguous stacks, https://golang.org/s/contigstacks

### Memory Allocator

- runtime: smarter scavenging, [#30333](https://golang.org/issue/30333)
- runtime: scavenger pacing fails to account for fragmentation [#34048](https://golang.org/issue/34048)
- runtime: potential deadlock cycle caused by scavenge.lock [#34047](https://golang.org/issue/34047)
- runtime: make the page allocator scale [#35112](https://golang.org/issue/35112)
- runtime: scavenger not as effective as in previous releases [#35788](https://golang.org/issue/35788)
- runtime: mechanism for monitoring heap size [#16843](https://golang.org/issue/16843)
- Scavenging https://go.googlesource.com/proposal/+/aa701aae530695d32916b779e048a3e18311a2e3/design/30333-smarter-scavenging.md
- Page allocator https://go.googlesource.com/proposal/+/a078ea9d72b99dc88fdfd2cb6ee150a8ce202ea2/design/35112-scaling-the-page-allocator.md
- Mcentral https://golang.org/issue/37487, https://golang.org/cl/221182

### Garbage Collector

- https://golang.org/s/go14gc
- https://golang.org/s/go15gcpacing
- runtime: eliminate stack rescanning, [#17503](https://golang.org/issue/17503)
- cmd/compile: compiler can unexpectedly preserve memory, [#22350](https://golang.org/issue/22350)
    + Proposal: [GC scanning of stacks](https://docs.google.com/document/d/1un-Jn47yByHL7I0aVIP_uVCMxjdM5mpelJhiKlIqxkE/edit#)
- runtime: simplify mark termination and eliminate mark 2, [#26903](https://golang.org/issue/26903)
- runtime: error message: P has cached GC work at end of mark termination [#27993](https://golang.org/issue/27993)
- runtime: cannot ReadMemStats during GC [#19812](https://golang.org/issue/19812)
- https://blog.golang.org/ismmkeynote
- https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns
- generational gc: https://golang.org/cl/137482
- ROC: https://golang.org/cl/25058

### Memory model

Go memory model is not properly defined.

- https://golang.org/issue/5045
- https://golang.org/issue/9442
- https://golang.org/issue/7948
- https://golang.org/issue/28306
- https://golang.org/issue/33815
- http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf
- https://research.swtch.com/go2017#memory
- https://golang.org/pkg/sync/atomic/#pkg-note-BUG
- https://groups.google.com/d/msg/golang-dev/vVkH_9fl1D8/azJa10lkAwAJ

### ABI

- register based ABI - Austin - Late call - david chase https://github.com/golang/proposal/blob/master/design/40724-register-calling.md

## Standard Library

### encoding

- Go 1.2 encoding.TextMarshaler and TextUnmarshaler. [design/go12encoding](https://golang.org/s/go12encoding)

### syscall

- The syscall package. [design/go14syscall](https://golang.org/s/go1.4-syscall)

### go/types

- go/types: The Go Type Checker. [doc/types](https://golang.org/s/types-tutorial)

### sync

- sync: avoid clearing the full Pool on every GC. [issue/22950](https://golang.org/issue/22950)
- runtime: replace Semacquire/Semrelease implementation. [cl/4631059-97c00f](https://github.com/golang/go/commit/997c00f)
- runtime: fall back to fair locks after repeated sleep-acquire failures. [issue/13086](https://golang.org/issue/13086)

### time

- runtime: make timers faster. [issue/6239](https://golang.org/issue/6239)
- time: excessive CPU usage when using Ticker and Sleep. [issue/27707](https://golang.org/issue/27707)

### io

- io/fs abstraction, Russ & Rob. [design/draft-iofs](https://github.com/golang/proposal/blob/master/design/draft-iofs.md)

### context

- Go Concurrency Patterns: Context. [doc/context](https://blog.golang.org/context)

## Contributing

PRs are welcome.

## License

[go-history](https://github.com/changkun/go-history) CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)