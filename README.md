# [Go: A Documentary](https://golang.design/history)

_by Changkun Ou <[changkun.de](https://changkun.de)>_ (and many inputs from [contributors](https://github.com/golang-design/history/graphs/contributors))

This document collects many interesting (publicly observable) issues,
discussions, proposals, CLs, and talks from the Go development process,
which intends to offer a comprehensive reference of the Go history.

<a id="top"></a>

**Table of Contents**

- [Go: A Documentary](#go-a-documentary)
  - [Disclaimer](#disclaimer)
  - [Sources](#sources)
  - [Origin](#origin)
  - [Committers](#committers)
    - [Core Authors](#core-authors)
    - [Compiler/Runtime Team](#compilerruntime-team)
    - [Library/Tools/Release/Security/Community](#librarytoolsreleasesecuritycommunity)
    - [Group Interviews](#group-interviews)
  - [Timeline](#timeline)
  - [Language Design](#language-design)
    - [General](#general)
    - [Slice (1.2)](#slice-12)
    - [Package Management (1.4, 1.5, 1.7)](#package-management-14-15-17)
    - [Type alias (1.9)](#type-alias-19)
    - [Defer (1.14)](#defer-114)
    - [Error values](#error-values)
    - [Channel/Select](#channelselect)
    - [Generics](#generics)
    - [Iterator](#iterator)
    - [Compatibility](#compatibility)
  - [Compiler Toolchain](#compiler-toolchain)
    - [Office Hours](#office-hours)
    - [Compiler](#compiler)
    - [Linker](#linker)
    - [Debugger](#debugger)
    - [Race Detector](#race-detector)
    - [Tracer](#tracer)
    - [Lock Analysis](#lock-analysis)
    - [Builder](#builder)
    - [Modules](#modules)
    - [gopls](#gopls)
    - [Testing, x/perf](#testing-xperf)
  - [Runtime Core](#runtime-core)
    - [Scheduler](#scheduler)
    - [Execution Stack](#execution-stack)
    - [Memory Management](#memory-management)
      - [Allocator](#allocator)
      - [Collector](#collector)
    - [Statistics](#statistics)
    - [Memory model](#memory-model)
    - [ABI](#abi)
    - [Misc](#misc)
  - [Standard Library](#standard-library)
    - [syscall](#syscall)
    - [os, io, io/fs, embed](#os-io-iofs-embed)
    - [go/\*](#go)
    - [sync](#sync)
      - [Map](#map)
      - [Pool](#pool)
      - [Mutex, RWMutex](#mutex-rwmutex)
      - [Groups](#groups)
      - [atomic](#atomic)
    - [time](#time)
    - [context](#context)
    - [encoding](#encoding)
    - [image, x/image](#image-ximage)
    - [math](#math)
    - [Mobile](#mobile)
    - [log](#log)
    - [misc](#misc-1)
  - [Unclassified But Relevant Links](#unclassified-but-relevant-links)
  - [Fun Facts](#fun-facts)
  - [Acknowledgements](#acknowledgements)
  - [License](#license)

---

## Disclaimer

- Most of the texts are written as _subjective_ understanding based on public sources
- **Factual and typo errors may occur.**
Referring to the original link if some text conflicts to your understanding
- [PR](https://github.com/golang-design/history/pulls)s are very welcome for new content submission, bug fixes, and typo fixes
- Use [Issues](https://github.com/golang-design/history) for discussions

[Back To Top](#top)

## Sources

<!--
  + [issue/](https://go.dev/issue/)
  + [issue/](https://go.dev/issue/)
  + [issue/](https://go.dev/issue/)
  + [design/](https://go.dev/design/)
  + [cl/](https://go.dev/cl/)
-->

There are many sources for digging the documents that relate to Go's
historical design, and here are some of the official sources:

- [go.dev/blog](https://go.dev/blog)
- [dev.golang.org](https://dev.golang.org)
- [build.golang.org](https://build.golang.org/)
- [farmer.golang.org](https://farmer.golang.org/)
- [go.dev/talks](https://go.dev/talks)
- [go.dev/doc](https://go.dev/doc)
- [go.dev/pkg](https://go.dev/pkg)
- [github.com/golang/go](https://github.com/golang/go)
- [github.com/golang/proposal](https://github.com/golang/proposal)
- [github.com/golang/go/wiki](https://github.com/golang/go/wiki)
- [go-review.googlesource.com](https://go-review.googlesource.com)
- [groups.google.com/g/golang-nuts](https://groups.google.com/g/golang-nuts)
- [groups.google.com/g/golang-dev](https://groups.google.com/g/golang-dev)
- [groups.google.com/g/golang-tools](https://groups.google.com/g/golang-tools)

[Back To Top](#top)

## Origin

Go is a big project driven by a tiny group of people and the crowd of
wisdom from the language user community. Here are some core committers
to the project if you are interested in following their excellent work.

Go's origin is attractive without doubts. By listening to the talks held
by these people, you could learn more about their oral history and fun
stories behind the great work.

For instance, here are some exciting talks regarding the historical feats
in Go might be your starting points (as a _subjective_ recommendation):
`talk/rob2007`, `talk/rob2009`, `talk/rob2010b`, `talk/rob2010d`,
`talk/rob2011a`, `talk/rob2013a`, `talk/rob2019`, `talk/robert2015`,
`talk/russ2014`, `steve2019b`, etc.

[Back To Top](#top)

## Committers

### Core Authors

The Go language was created by Rob, Robert, and Ken initially because
Rob was bothered by slow C++ compiling times, talked to Robert, and luckily
Ken was in the next office.
Later, Ian joined the project after showing interest, and
wrote the [gccgo](https://github.com/golang/gofrontend).
Rob and Ken are retired. Robert and Ian worked together for years on adding
generics to Go, which shipped in Go 1.18 in 2022; Robert still works on the
language and its type checker. Ian [left Google](https://www.airs.com/blog/archives/670)
in May 2025 and has kept contributing to Go since.
Russ is also one of the core authors of the project in the early stage.
Back then, he was a newcomer at Google, and Rob invited Russ for joining
the Go team since he knew Russ from way back because of the
[Plan 9](https://9p.io/plan9/) project. Russ worked on the early Go compiler, runtime, as well as the leap of
Go 1.5 bootstrap.
Russ led the Go project for over 12 years. In August 2024 he
[passed the torch](https://groups.google.com/g/golang-dev/c/0OqBkS2RzWw):
since September 1, 2024, Austin Clements has been the tech lead of Go, and
Cherry Mui leads Go core (the compiler toolchain, runtime, and releases).
Russ stayed on the project, turning to Gaby and Oscar, tools that help
maintainers work through the issue tracker.

- Rob Pike. (Robert C. Pike, M. Sc.) [Website](http://herpolhode.com/rob/), [Blog](https://commandcenter.blogspot.com/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike). (Retired)
  + Alma mater: University of Toronto (BS), California Institute of Technology
  + [talk/rob2007](https://www.youtube.com/watch?v=hB05UFqOtFA) Advanced Topics in Programming Languages: Concurrency/message passing Newsqueak. May 9, 2007
  + [talk/rob2009](https://changelog.com/podcast/3) The Go Programming Language. Nov 27, 2009.
  + [talk/rob2010a](https://www.youtube.com/watch?v=jgVhBThJdXc) Go Programming. Google I/O 2010. May 20, 2010
  + [talk/rob2010b](https://www.youtube.com/watch?v=-_DKfAn4pFA) Origins of Go Concurrency style. Emerging Languages Camp 2010. July 21, 2010.
  + [talk/rob2010c](https://www.youtube.com/watch?v=5kj5ApnhPAE) Public Static Void. OSCON 2010. Jul 22, 2010.
  + [talk/rob2010d](https://www.youtube.com/watch?v=7VcArS4Wpqk) Another Go at Language Design. Aug 27, 2010.
  + [talk/rob2011a](https://www.infoq.com/interviews/pike-concurrency/) Parallelism and Concurrency in Programming Languages. Feb 17, 2011.
  + [talk/rob2011b](https://www.infoq.com/interviews/pike-google-go/) Google Go: Concurrency, Type System, Memory Management and GC. Feb 25, 2011.
  + [talk/rob2011c](https://www.youtube.com/watch?v=HxaD_trXwRE) Lexical Scanning in Go. Aug 30, 2011.
  + [talk/rob2012a](https://www.youtube.com/watch?v=f6kdp27TYZs) Go Concurrency Patterns. Google I/O 2012. Jul 2, 2012.
  + [talk/rob2012b](https://www.youtube.com/watch?v=FTl0tl9BGdc) Why Learn Go?. Sep 12, 2012.
  + [talk/rob2013a](https://www.youtube.com/watch?v=bj9T2c2Xk_s) The path to Go 1. Mar 14, 2013.
  + [talk/rob2013b](https://www.infoq.com/presentations/Go-Google/) Go at Google. Apr 13, 2013.
  + [talk/rob2013c](https://changelog.com/podcast/100) Go Programming with Rob Pike and Andrew Gerrand. Aug 14, 2013.
  + [talk/rob2013d](https://www.youtube.com/watch?v=qmg1CF3gZQ0) Concurrency Is Not Parallelism. Oct 20, 2013.
  + [talk/rob2014a](https://www.youtube.com/watch?v=VoS7DsT1rdM) Hello Gophers! Gophercon Opening Keynote. GopherCon 2014. Apr 24, 2014.
  + [talk/rob2014b](https://www.youtube.com/watch?v=PXoG0WX0r_E) Implementing a bignum calculator. Sydney Go meeup. Nov, 19 2014.
  + [talk/rob2015a](https://www.youtube.com/watch?v=cF1zJYkBW4A) The move from C to Go in the toolchain. GopherFest 2015. Jun 10, 2015.
  + [talk/rob2015b](https://www.youtube.com/watch?v=PAAkCSZUG1c) Go Proverbs. Gopherfest. Nov 18, 2015.
  + [talk/rob2015c](https://www.youtube.com/watch?v=rFejpH_tAHM) Simplicity is Complicated. dotGo 2015. Dec 2, 2015.
  + [talk/rob2016a](https://www.youtube.com/watch?v=KINIAgRpkDA) The Design of the Go Assembler. GopherCon 2016. Aug 18, 2016.
  + [talk/rob2016b](https://www.youtube.com/watch?v=sDTGhIqyMjo) Stacks of Tokens: A study in interfaces. Sydney Go Meetup. September 19, 2016.
  + [talk/rob2017](https://www.youtube.com/watch?v=ENLWEfi0Tkg) Upspin. Gopherfest 2017. Jun 22, 2017.
  + [talk/rob2018a](https://www.youtube.com/watch?v=_2NI6t2r_Hs) The History of Unix. Nov 7, 2018.
  + [talk/rob2018b](https://www.youtube.com/watch?v=RIvL2ONhFBI) Go 2 Draft Specifications. Sydney Golang Meetup. Nov 13, 2018.
  + [talk/rob2019a](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer. Sep 10, 2019.
  + [talk/rob2019b](https://www.youtube.com/watch?v=oU9cfQCxjpM) Rob Pike. A Brief History of Go. Dec 12, 2019.
  + [talk/rob2020](https://evrone.com/rob-pike-interview) A Rob Pike Interview. (Date Unclear) Apr 30, 2020.
  + [talk/rob2021](https://www.youtube.com/watch?v=YXV7sa4oM4I) The Go Programming Language and Environment. John Lions Distinguished Lecture, UNSW, May 27, 2021.
  + [talk/rob2023] What We Got Right, What We Got Wrong. GopherCon Australia 2023. Nov 10, 2023.
- Robert Griesemer. (Dr. Robert Griesemer) [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
  + Alma mater: ETH Zürich
  + [paper/robert1993](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.127.5290&rep=rep1&type=pdf) A Programming Language for Vector Computers. Doctor Dissertation. 1993.
  + [talk/robert2012a](https://www.youtube.com/watch?v=on5DeUyWDqI) E2E: Erik Meijer and Robert Griesemer. Going Go. Lang.NEXT. Mar 16, 2012.
  + [talk/robert2012b](https://www.youtube.com/watch?v=Ns3csOBiIuE) Go In Three Easy Pieces. Mar 19, 2012.
  + [talk/robert2012c](https://www.youtube.com/watch?v=MrY6f4siVA4) Lang.NEXT 2012 Expert Panel: Native Languages. Apr 11, 2012.
  + [talk/robert2015](https://www.youtube.com/watch?v=0ReKdcpNyQg) The Evolution of Go. GopherCon 2015. Jul 28, 2015.
  + [talk/robert2016a](https://www.youtube.com/watch?v=t-w6MyI2qlU) Lightning Talk: Alias Declarations for Gom: A proposal. GopherCon 2016. Oct 9, 2016.
  + [talk/robert2016b](https://www.youtube.com/watch?v=vLxX3yZmw5Q) Prototype your design!. dotGo 2016. Nov 29, 2016.
  + [talk/robert2017](https://www.youtube.com/watch?v=KPk1UPihWtY) Opening Keynote: Exporting Go. GopherCon SG 2017. May 29, 2017.
  + [talk/robert2017](https://www.youtube.com/watch?v=UmwJsQTSEP8) A brief overview of Go. Aug 28. 2017.
  + [talk/robert2019](https://www.youtube.com/watch?v=i0zzChzk8KE) Go is 10! Now What?. Gopherpalooza 2019. Dec 2, 2019.
  + [talk/robert2020a](https://changelog.com/gotime/140) The latest on Generics with Robert Griesemer and Ian Lance Taylor. Jul 21, 2020.
  + [talk/robert2020b](https://www.youtube.com/watch?v=TborQFPY2IM) Typing [Generic] Go. Nov 11, 2020.
  + [talk/robert2021](https://www.youtube.com/watch?v=Pa_e9EeCdy8) Generics! Dec 17, 2021.
  + [talk/robert2023](https://www.youtube.com/watch?v=2putv2jOqdg) Everything You Always Wanted to Know About Type Inference. GopherCon 2023, San Diego. September 2023.
- Ken Thompson. (Kenneth Lane Thompson, M. Sc.) (Retired)
  + Alma mater: UC Berkeley
  + [talk/ken1982a](https://www.youtube.com/watch?v=tc4ROCJYbm0) The UNIX System: Making Computers More Productive. 1982.
  + [talk/ken1982b](https://www.youtube.com/watch?v=XvDZLjaCJuw) UNIX: Making Computers Easier To Use.
  + [talk/ken1982c](https://www.youtube.com/watch?v=JoVQTPbD6UY) Ken Thompson and Dennis Ritchie Explain UNIX (Bell Labs).
  + [talk/ken1998](https://www.youtube.com/watch?v=LXZ1OL2U3lY) Ken Thompson and Dennis Ritchie. National Medal of Technology Awards. 1998.
  + [talk/ken2013](https://www.youtube.com/watch?v=dsMKJKTOte0) Systems Architecture, Design, Engineering, and Verification. Jan 17, 2013.
  + [talk/ken2019a](https://www.youtube.com/watch?v=g3jOJfrOknA) The Thompson and Ritchie Story. Feb 18, 2019.
  + [talk/ken2019b](https://www.youtube.com/watch?v=EY6q5dv_B-o) Brian Kernighan interviews Ken Thompson. VCF East 2019. May 4, 2019.
  + [talk/ken2023](https://www.youtube.com/watch?v=kaandEt_pKw&t=643s). Closing Keynote. SCaLE 20x. Mar 12, 2023. https://www.socallinuxexpo.org/scale/20x/presentations/keynote-ken-thompson
- Ian Taylor. (Ian Lance Taylor, B. Sc.) [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor), [Quora](https://www.quora.com/profile/Ian-Lance-Taylor)
  + Alma mater: Yale University
  + [talk/ian2007](https://www.youtube.com/watch?v=gc78olyguqA) GCC: Current Topics and Future Directions. February 27, 2007.
  + [talk/ian2018](https://www.youtube.com/watch?v=LqKOY_pH8u0) Transition to Go 2. Gopherpalooza 2018. Oct 24, 2018
  + [talk/ian2019a](https://www.youtube.com/watch?v=WzgLqE-3IhY) Generics in Go. GopherCon 2019. Aug 27, 2019
  - [talk/ian2019b](https://changelog.com/gotime/98) Generics in Go. Aug 27, 2019.
  + [talk/ian2020](https://www.youtube.com/watch?v=yoZ05GG8aLs) Go with Ian Lance Taylor. CppCast. Aug 9, 2020.
  + [talk/ian2021a](https://www.youtube.com/watch?v=nr8EpUO9jhw) Using Generics in Go. Dec 6, 2021.
  + [talk/ian2021b](https://www.youtube.com/watch?v=Pa_e9EeCdy8) Generics! Dec 17, 2021.
  + [talk/ian2024](https://www.youtube.com/watch?v=i9zwUT9dlVc) Range Over Function Types. GopherCon 2024.
  + [doc/ian2025](https://www.airs.com/blog/archives/670) Ian Lance Taylor. Leaving Google. May 10, 2025.

- Russ Cox. (Dr. Russell Stensby Cox) [Website](https://swtch.com/~rsc/), [Blog](https://research.swtch.com/), [GitHub](https://github.com/rsc), [Twitter](https://twitter.com/_rsc), [Reddit](https://old.reddit.com/user/rsc), [YouTube](https://www.youtube.com/channel/UC3P6PrEBAVH1UaiPOzZ-u-w)
  + Alma mater: MIT
  + [paper/russ2008](https://pdos.csail.mit.edu/~rsc/) An Extension-Oriented Compiler. Doctor Dissertation. Aug 20, 2008.
  + [talk/russ2009](https://www.youtube.com/watch?v=wwoWei-GAPo) The Go Programming Language Promo. Nov 10, 2009.
  + [talk/russ2012a](https://www.youtube.com/watch?v=MzYZhh6gpI0) A Tour of the Go Programming Language. Jun 24, 2012.
  + [talk/russ2012b](https://www.youtube.com/watch?v=dP1xVpMPn8M) A Tour of the Acme Editor. Sep 17, 2012.
  + [talk/russ2014](https://www.youtube.com/watch?v=QIE5nV5fDwA) Go from C to Go. GopherCon 2014.  May 18, 2014
  + [talk/russ2015](https://www.youtube.com/watch?v=XvZOdpd_9tc) Go, Open Source, Community. GopherCon 2015. Jul 28, 2015.
  + [talk/russ2016](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Codebase Refactoring (with help from Go). Dec 5, 2016.
  + [talk/russ2017](https://www.youtube.com/watch?v=0Zbh_vmAKvk) The Future of Go. GopherCon 2017. Jul 24, 2017.
  + [talk/russ2018a](https://changelog.com/gotime/77) Dependencies and the future of Go. Apr 19, 2018.
  + [talk/russ2018b](https://changelog.com/gotime/bonus-77) Go and WebAssembly (Wasm). Apr 19, 2018
  + [talk/russ2018c](https://www.youtube.com/watch?v=F8nrpe0XWRg) Opening keynote: Go with Versions. GopherConSG 2018. May 5, 2018
  + [talk/russ2018d](https://www.youtube.com/watch?v=6wIP3rO6On8) Go 2 Drafts Announcement. Aug 28, 2018.
  + [talk/russ2019](https://www.youtube.com/watch?v=kNHo788oO5Y) On the Path to Go 2. GopherCon 2019. Aug 27, 2019.
  + [talk/russ2020a](https://go.dev/s/go-build-video) `go:build` design draft. Jun 30, 2020.
  + [talk/russ2020b](https://go.dev/s/draft-iofs-video) `io/fs` draft design. Jul 21, 2020.
  + [talk/russ2020c](https://go.dev/s/draft-embed-video) `//go:embed` draft design. Jul 21, 2020.
  + [talk/russ2021](https://archive.org/details/PLTalk/%23PLTalk+-+12+Years+of+Go+with+Russ+Cox+%5Bv1203523364%5D.mp4) #PLTalk: 12 Years of Go with Russ Cox. Nov. 12, 2021.
  + [talk/russ2022](https://youtube.com/watch?v=v24wrd3RwGo) Compatibility: How Go Programs Keep Working. GopherCon 2022. Oct 28, 2022.
  + [talk/russ2023a](https://www.youtube.com/watch?v=BNmxtp26I5s) Go Changes. GopherCon 2023. Sep 27, 2023.
  + [talk/russ2023b](https://www.youtube.com/watch?v=X4rxi9jStLo) Go Testing By Example. GopherCon Australia 2023. Nov 9, 2023.
  + [talk/russ2023c](https://www.youtube.com/watch?v=6H-V-0oQvCA) Open Source Supply Chain Security at Google. ACM SCORED. Nov 30, 2023.
  + [talk/russ2024](https://www.youtube.com/watch?v=3QL062vGbjQ) Go Telemetry Wins. GoLab. Dec 12, 2024.
  + [doc/13years](https://go.dev/blog/13years) Russ Cox. Thirteen Years of Go. November 10, 2022.
  + [doc/14years](https://go.dev/blog/14years) Russ Cox. Fourteen Years of Go. November 10, 2023.
  + [doc/torches](https://groups.google.com/g/golang-dev/c/0OqBkS2RzWw) Russ Cox. passing torches to Austin and Cherry. golang-dev. August 1, 2024.
  + [discuss/67901](https://github.com/golang/go/discussions/67901) Russ Cox. gabyhelp feedback. June 8, 2024.
  + [discuss/68490](https://github.com/golang/go/discussions/68490) Russ Cox. x/oscar feedback. July 17, 2024.

[Back To Top](#top)

### Compiler/Runtime Team

Dmitry is not from the Go team but on the dynamic testing tools team from Google.
He wrote the scalable goroutine scheduler, many other performance improvements,
synchronization primitives, race detector, and blocking profiler that
related to the Go runtime.
Austin was an intern at Google who worked on the Go project in the early days
while pursuing a Ph. D. Later, he joined the Go team after his academic career
and work together with Rick for the Go's concurrent GC. He also worked on
the current preemptive scheduler and linker.
Now, he is leading the Compiler/Runtime team for Go.
Keith and David together focus on the Go's compiler backend,
notably the current SSA backend. Michael is a recent newcomer to the Go team,
his work is mainly in the runtime memory system such as the refactoring of memory allocator and runtime metrics.

- Dmitry Vyukov. (Дмитрий Вьюков, M. Sc.) [Website](http://www.1024cores.net/), [GitHub](https://github.com/dvyukov), [Twitter](https://twitter.com/dvyukov)
  + Alma mater: Bauman Moscow State Technical University
  + [talk/dmitry2014](https://www.youtube.com/watch?v=QEhpLb0UCfE) Writing a functional, reliable and fast web application in Go. Sep 25, 2014.
  + [talk/dmitry2015a](https://www.youtube.com/watch?v=Ef7TtSZlmlk) Chat with Dmitry Vyukov on go-fuzz, golang and IT security. Jul 3, 2015
  + [talk/dmitry2015b](https://www.youtube.com/watch?v=a9xrxRsIbSU) Go Dynamic Tools. GopherCon 2015. Jul 28, 2015.
  + [talk/dmitry2016](https://www.youtube.com/watch?v=9cpN2r22sLE) Dmitry Vyukov Interview. Jun 1, 2016.
  + [talk/dmitry2017](https://www.youtube.com/watch?v=FD30Qzd6ylk) Fuzzing: The New Unit Testing. C++ Russia 2017. May 10, 2017.
  + [talk/dmitry2018a](https://www.youtube.com/watch?v=EJVp13f_aIs) Fuzzing: new unit testing. GopherCon Russia. Apr 2, 2018.
  + [talk/dmitry2018b](https://www.youtube.com/watch?v=qrBVXxZDVQY) Syzbot and the Tale of Thousand Kernel Bugs. Sep 1, 2018.
  + [talk/dmitry2019](https://www.youtube.com/watch?v=-K11rY57K7k) Go scheduler: Implementing language with lightweight concurrency. Oct 14, 2019.
  + [talk/dmitry2020](https://www.youtube.com/watch?v=YwX4UyXnhz0) syzkaller: Adventures in Continuous Coverage-guided Kernel Fuzzing. BlueHat IL 2020. Feb 13, 2020.
- Austin Clements. (Dr. Austin T. Clements) (Go tech lead since September 2024) [GitHub](https://github.com/aclements), [Scholar](https://scholar.google.com/citations?user=MKDtxN4AAAAJ)
  + Alma mater: MIT
  + [paper/autin2014](https://pdos.csail.mit.edu/papers/aclements-phd.pdf) The Scalable Commutativity Rule: Designing Scalable Software for Multicore Processors. Doctor Dissertation. May 21, 2014.
  + [talk/austin2020](https://www.gophercon.com/agenda/session/233441) Pardon the Interruption: Loop Preemption in Go 1.14. Nov 12, 2020.
  + [doc/15years](https://go.dev/blog/15years) Austin Clements. Go Turns 15. November 11, 2024.
  + [doc/16years](https://go.dev/blog/16years) Austin Clements. Go’s Sweet 16. November 14, 2025.
- Richard Hudson. (Richard L. Hudson, M. Sc.) [GitHub](https://github.com/RLH) (Retired)
  + Alma mater: University of Massachusetts Amherst
  + [paper/rick](https://dl.acm.org/profile/81100566849/publications?Role=author) Research List
  + [talk/rick2015a](https://www.youtube.com/watch?v=aiv1JOfMjm0) Go GC: Solving the Latency Problem. GopherCon 2015. Jul 8, 2015.
  + [talk/rick2015b](https://www.infoq.com/interviews/hudson-go-gc/) Rick Hudson on Garbage Collection in Go. Dec 21, 2015.
- Keith Randall. (Dr. Keith H. Randall.) [GitHub](https://github.com/randall77)
  + Alma mater: MIT
  + [paper/1998keith](http://supertech.csail.mit.edu/papers/randall-phdthesis.pdf) Cilk: Efficient Multithreaded Computing. Doctor Dissertation. May 21, 1998.
  + [talk/keith2016](https://www.youtube.com/watch?v=Tl7mi9QmLns) Inside the Map Implementation. GopherCon 2016. Jul 12, 2016.
  + [talk/keith2017](https://www.youtube.com/watch?v=uTMvKVma5ms) Generating Better Machine Code with SSA. GopherCon 2017. Jul 24, 2017.
  + [talk/keith2024](https://www.youtube.com/watch?v=7_h9iT672HQ) Interface Internals. GopherCon 2024

- David Chase. (Dr. David Chase) [Website](http://chasewoerner.org/resume.html), [Block](https://dr2chase.wordpress.com/), [GitHub](https://github.com/dr2chase), [Twitter](https://twitter.com/dr2chase), [Scholar](https://dblp.org/pid/51/3488.html)
  + Alma mater: Rice University
  + [paper/1987david](http://www.chasewoerner.org/dissertation.pdf) Garbage Collection and Other Optimizations. Doctor Dissertation. Aug 1987.
  + [talk/david2017](https://changelog.com/gotime/52) All About The Go Compiler. Jul 20, 2017.
- Dan Scales.
  + [talk/dan2020](https://www.gophercon.com/agenda/session/233397) Implementing Faster Defers. Nov 11, 2020
- Michael Knyszek. [Website](https://www.ocf.berkeley.edu/~mknyszek/), [GitHub](https://github.com/mknyszek)
  + [talk/michael2020](https://www.gophercon.com/agenda/session/233086) Evolving the Go Memory Manager's RAM and CPU Efficiency. Nov 11, 2020
  + [talk/michael2025](https://www.youtube.com/watch?v=gPJkM95KpKo) Advancing Go Garbage Collection with Green Tea. GopherCon 2025, New York City. August 2025.
- Than McIntosh. [GitHub](https://github.com/thanm)
- Cherry Mui (formerly Cherry Zhang). (Go core lead since September 2024) [GitHub](https://github.com/cherrymui)

[Back To Top](#top)

### Library/Tools/Release/Security/Community

- Andrew Gerrand. [GitHub](https://github.com/adg), [Twitter](https://twitter.com/enneff)
- Brad Fitzpatrick. [Website](https://bradfitz.com/), [GitHub](https://github.com/bradfitz), [Twitter](https://twitter.com/bradfitz)
  + [talk/brad2010](https://www.youtube.com/watch?v=c4znvD-7VDA) Writing zippy Android apps. Google I/O 2010. May 27, 2010.
  + [talk/brad2011](https://www.youtube.com/watch?v=OExxMUVVjRM) Palestra Brad FitzPatrick. Perl Workshop. Parte 1. May 12, 2011.
  + [talk/brad2012](https://www.youtube.com/watch?v=zHXoDB07Iwg) Behind the Scenes at LiveJournal: Scaling Storytime. Brad Fitzpatrick. May 21, 2012.
  + [talk/brad2013a](https://www.youtube.com/watch?v=sYukPc0y_Ro) Software I'm excited about. dotScale 2013. Jul 30, 2013.
  + [talk/brad2013b](https://www.youtube.com/watch?v=kJ68OWnlycQ) LiveJournal's Moscow office. Brad Fitzpatrick. Dec 15, 2013.
  + [talk/brad2014a](https://www.youtube.com/watch?v=Ds0KXj8ohR8) Camlistore. Go meetup. Jan 2014.
  + [talk/brad2014b](https://www.youtube.com/watch?v=D6okO8Qzusk) Camlistore & The Standard Library. GopherCon 2014. May 18, 2014.
  + [talk/brad2014c](https://www.youtube.com/watch?v=4KFTacxqkcQ) The State of the Gopher. dotGo 2014. October 10, 2014.
  + [talk/brad2014d](https://www.youtube.com/watch?v=VLciNxKAKyc) Interview. dotGo 2014. Nov 6, 2014.
  + [talk/brad2015a](https://www.youtube.com/watch?v=yG-UaBJXZ80) Hacking with Andrew and Brad: an HTTP/2 client. Feb 25, 2015.
  + [talk/brad2015b](https://www.youtube.com/watch?v=gukAZO1fqZQ) HTTP/2 and Go. London Go Gathering 2015. Mar 4, 2015.
  + [talk/brad2015c](https://www.youtube.com/watch?v=gAfoLwog_MA) HTTP/2 for Go. Go Devroom FOSDEM 2015. Mar 4, 2015.
  + [talk/brad2015d](https://www.youtube.com/watch?v=rHBbqjWCGq8) The state of the Gopher. Munich Gophers. Apr 18, 2015
  + [talk/brad2015e](https://www.youtube.com/watch?v=stram5J144s) The Go Language. devoxx poland 2015. June, 2015
  + [talk/brad2015f](https://www.youtube.com/watch?v=1rZ-JorHJEY) Hacking with Andrew and Brad: tip.golang.org. Jan 8, 2015.
  + [talk/brad2015g](https://www.youtube.com/watch?v=mj-1wscEQO8) Lightning Talk: HTTP/2 in Go. GopherCon 2015. Aug 19, 2015.
  + [talk/brad2015h](https://www.youtube.com/watch?v=xxDZuPEgbBU) Profiling & Optimizing in Go. Aug 27, 2015.
  + [talk/brad2016a](https://www.youtube.com/watch?v=4Dr8FXs9aJM) Introducing Go 1.6: asymptotically approaching boring. Gophercon India 2016. Mar 29, 2016.
  + [talk/brad2016b](https://www.youtube.com/watch?v=4yFb-b5GYWc) Lightning Talk: My video & security system. GopherCon 2016. July, 2016.
  + [talk/brad2016c](https://www.youtube.com/watch?v=8cQcPnzfkLk) Go 1.7. Seattle Go Meetup. Aug 25, 2016.
  + [talk/brad2016d](https://www.youtube.com/watch?v=18kmlJvR6Bk) The Go Open Source Project's Innards. GothamGo 2016. Dec 15, 2016.
  + [talk/brad2017](https://www.youtube.com/watch?v=4fWqcOubYQ0) Half my life spent in open source. May 19, 2017.
  + [talk/brad2018a](https://www.youtube.com/watch?v=ZCB-g2B4Y5A) Go: looking back and looking forward. Apr 2, 2018.
  + [talk/brad2018b](https://www.youtube.com/watch?v=rWJHbh6qO_Y) Go 1 11 and beyond. Aug 26, 2018.
  + [talk/brad2018c](https://www.youtube.com/watch?v=69Zy77O-BUM) Lightning Talk: The nuclear option, go test -run=InQemu. GopherCon 2018. Sep 11, 2018.
  + [talk/brad2019](https://www.youtube.com/watch?v=BRSam0xQJKY) Brad Fitzpatrick likes Go better than C, C++, Rust, Perl, Python, Ruby, JavaScript and Java. Nov 28, 2019.

- Bryan C. Mills. [GitHub](https://github.com/bcmills)
  + [talk/bryan2017](https://www.youtube.com/watch?v=C1EtfDnsdDs) Lightning Talk: An overview of sync.Map. GopherCon 2017. Jul 24, 2017.
  + [talk/bryan2018](https://www.youtube.com/watch?v=5zXAHh5tJqQ) Rethinking Classical Concurrency Patterns. GopherCon 2018. Sep 14, 2018

- Dmitri Shuralyov. (Dmitri Shuralyov, M. Sc.) [Website](https://dmitri.shuralyov.com/), [GitHub](https://github.com/dmitshur), [YouTube](https://www.youtube.com/c/DmitriShuralyov)
  + [talk/dmitri2016](https://www.youtube.com/watch?v=9XTl1d4nwdY) Go in the browser. dotGo. Oct 10, 2016.

- Steve Francia. [Website](https://spf13.com/), [GitHub](https://github.com/spf13), [Twitter](https://twitter.com/spf13).
  + [talk/steve2019a](https://spf13.com/presentation/what-should-a-modern-practical-programming-language-look-like/) What Should A Modern Practical Programming Language Look Like. Landing Festival. April 4, 2019.
  + [talk/steve2019b](https://spf13.com/presentation/the-legacy-of-go/) The Legacy Of Go. Go Lab. Oct 22, 2019.
  + [talk/steve2022](https://www.youtube.com/watch?v=atv47M6kBYc) Go: Our Journey from Early Adopter to Mainstream and Beyond. GopherCon 2022, Chicago. October 2022.
- Jonathan Amsterdam.
  + [talk/jonathan2020](https://www.gophercon.com/agenda/session/233432) Working with Errors. Nov 13, 2020.
  + [talk/jonathan2023](https://www.youtube.com/watch?v=rJfvv_c9mYU) Structured Logging for the Standard Library. GopherCon 2023, San Diego. September 2023.
- Daniel Martí. [Website](https://mvdan.cc/), [GitHub](https://github.com/mvdan), [Twitter](https://twitter.com/mvdan_)
- Nigel Tao. [GitHub](https://github.com/nigeltao), Twitter
- Michael Matloob. [GitHub](https://github.com/matloob), [Twitter](https://twitter.com/matloob)
- Jay Conrod. [Website](https://jayconrod.com/), [Twitter](https://twitter.com/jayconrod)
- Dave Cheney. [Website](https://dave.cheney.net/), [GitHub](https://github.com/davecheney), [Twitter](https://twitter.com/davecheney)
- Sam Boyer. [GitHub](https://github.com/sdboyer), [Twitter](https://twitter.com/sdboyer)
- Fillippo Valsorda. [Website](https://filippo.io/), [GitHub](https://github.com/FiloSottile), [Twitter](https://twitter.com/FiloSottile)
  + [talk/filo2016](https://www.youtube.com/watch?v=lhMhApWQp2E) From cgo back to Go. July 12, 2016
  + [talk/filo2017](https://speakerdeck.com/filosottile/calling-rust-from-go-without-cgo-at-gothamgo-2017) From Go to Rust without cgo.
  + [talk/filo2018](https://speakerdeck.com/filosottile/why-cgo-is-slow-at-capitalgo-2018) Why cgo is slow. CapitalGo 2018.
  + [talk/speakerdeck](https://speakerdeck.com/filosottile?page=1)
  + [talk/filo2023](https://www.youtube.com/watch?v=lkEH3V3PkS0) High-Assurance Go Cryptography. GopherCon 2023, San Diego. September 2023.
  + [talk/filo2024](https://www.youtube.com/watch?v=F3iDp6LZtyk) The Go Cryptography State of the Union. GopherCon 2024, Chicago. July 2024.
  + [talk/filo2025](https://www.youtube.com/watch?v=YnyeAQblUyA) The Go Cryptography State of the Union. GopherCon 2025, New York City. August 2025.
- Alan Donovan. [GitHub](https://github.com/adonovan)
  + [talk/alan2025](https://www.youtube.com/watch?v=_VePjjjV9JU) Analysis and Transformation Tools for Go Codebase Modernization. GopherCon 2025, New York City. August 2025.
- Cameron Balahan. [GitHub](https://github.com/cameronbalahan)
  + [talk/cameron2022](https://www.youtube.com/watch?v=b7l0Qt8m6OQ) A Holistic Go Experience. GopherCon 2022, Chicago. October 2022.
  + [talk/cameron2025](https://www.youtube.com/watch?v=M2gduDM-MT0) Go's Next Frontier. GopherCon 2025, New York City. August 2025.
- Katie Hockman. [GitHub](https://github.com/katiehockman)
  + [talk/katie2022](https://www.youtube.com/watch?v=7KWPiRq3ZYI) Fuzz Testing Made Easy. GopherCon 2022, Chicago. October 2022.
- Julie Qiu. [GitHub](https://github.com/julieqiu)
  + [talk/julie2023](https://www.youtube.com/watch?v=nJm2GP6kMJo) Vulnerability Management for Go. GopherCon 2023, San Diego. September 2023.

> more people and talks should be added...

[Back To Top](#top)

### Group Interviews

- [talk/goteam2012](https://www.youtube.com/watch?v=sln-gJaURzk) Meet the Go Team. Google I/O 2012. Jul 2, 2012.
- [talk/goteam2013](https://www.youtube.com/watch?v=p9VUCp98ay4) Fireside Chat with the Go Team. Google I/O 2013. May 18, 2013.
- [talk/goteam2014](https://www.youtube.com/watch?v=u-kkf76TDHE) Inside the Gophers Studio with Blake Mizerany. GopherCon 2014. May 21, 2014
- [talk/goteam2019](https://www.youtube.com/watch?v=3yghHvvZQmA) Meet the Authors: Go Language. Cloud Next '19. Apr 10, 2019.
- [talk/goteam2020a](https://www.youtube.com/watch?v=gJxvkOHpTSM) GoLab 2020: Go Team AMA. Oct 22, 2020.
- [talk/goteam2020b](https://www.youtube.com/watch?v=BNHwHLNLjLs) GopherCon 2020: Go Team AMA. Nov 16, 2020.
- [talk/goteam2021](https://www.youtube.com/watch?v=hu9spj-GJ1U&) GopherCon 2021: Go Time Live - Go Team "Ask Me Anything", Dec 28, 2021.

[Back To Top](#top)

## Timeline

A timeline helps you identify the point in time about a document that is
reflected in Go versions.

For what changed in the runtime and compiler in each of these releases,
including internal changes that release notes leave out, see the
[runtime and compiler history reconstructed from commits](https://golang.design/history/runtime-compiler/)
(Go 1.0 to 1.28). Its [paths not taken](https://golang.design/history/runtime-compiler/paths.html) page collects work that was built or designed but never landed in that form, such as the request-oriented collector and lock-free channels, with the reasons each one stopped.

| Version | Release Expected/Actual Date | Days Since Last Release (+Delay) |
|:--|:--|:--|
| Go 1    | - / 2012.03.28 | - |
| Go 1.1  | - / 2013.05.13 | 440 |
| Go 1.2  | - / 2013.12.01 | 202 |
| Go 1.3  | - / 2014.06.18 | 199 |
| Go 1.4  | - / 2014.12.10 | 175 |
| Go 1.5  | 2015.07.31 / 2015.08.19 | 252 (+19) |
| Go 1.6  | 2016.01.31 / 2016.02.17 | 182 (+17) |
| Go 1.7  | 2016.07.31 / 2016.08.15 | 180 (+15) |
| Go 1.8  | 2017.01.31 / 2017.02.16 | 185 (+16) |
| Go 1.9  | 2017.07.31 / 2017.08.24 | 189 (+24) |
| Go 1.10 | 2018.01.31 / 2018.02.16 | 176 (+16) |
| Go 1.11 | 2018.07.31 / 2018.08.24 | 189 (+24) |
| Go 1.12 | 2019.01.31 / 2019.02.25 | 185 (+25) |
| Go 1.13 | 2019.07.31 / 2019.09.03 | 190 (+34) |
| Go 1.14 | 2020.01.31 / 2020.02.25 | 175 (+25) |
| Go 1.15 | 2020.07.31 / 2020.08.11 | 168 (+11) |
| Go 1.16 | 2021.01.31 / 2021.02.16 | 189 (+16) |
| Go 1.17 | 2021.07.31 / 2021.08.16 | 181 (+16) |
| Go 1.18 | 2022.01.31 / 2022.03.15 | 196 (+43) |
| Go 1.19 | 2022.07.31 / 2022.08.02 | 140 (+2)  |
| Go 1.20 | 2023.01.31 / 2023.02.01 | 183 (+1)  |
| Go 1.21 | 2023.07.31 / 2023.08.08 | 188 (+8)  |
| Go 1.22 | 2024.01.31 / 2024.02.07 | 184 (+7)  |
| Go 1.23 | 2024.07.31 / 2024.08.14 | 189 (+14)  |
| Go 1.24 | 2025.01.31 / 2025.02.12 | 182 (+12)  |
| Go 1.25 | 2025.07.31 / 2025.08.12 | 181 (+12)  |
| Go 1.26 | 2026.01.31 / 2026.02.10 | 182 (+10)  |
| Go 1.27 | 2026.07.31 / 2026.08.19 | 190 (+19)  |

The historical release notes may helpful for general information:

- [doc/go1release](https://go.dev/doc/devel/release.html) Go Release History
- [doc/go1prerelease](https://go.dev/doc/devel/pre_go1.html) Pre-Go 1 Release History
- [doc/go0release](https://go.dev/doc/devel/weekly.html) Weekly Release History (Before Go 1)

[Back To Top](#top)

## Language Design

### General

- [design/go0initial](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) Rob Pike, Robert Griesemer, Ken Thompson. The Go Annotated Specification. Mar 3, 2008.
- [design/go0spec0](https://github.com/golang/go/blob/e6626dafa8de8a0efae351e85cf96f0c683e0a4f/doc/go_lang.txt) The Go Programming Language. Language Specification. Mar 7, 2008.
- [design/go0semicolon](https://go.dev/s/semicolon-proposal) Rob Pike. Semicolons in Go. Dec 10, 2009.
- [design/go11func](https://go.dev/s/go11func) Russ Cox. Go 1.1 Function Calls. February 2013
- [design/go11return](https://go.dev/s/go11return) Russ Cox. Go “Return at End of Function” Requirements. March 2013
- [design/go12nil](https://go.dev/s/go12nil) Russ Cox. Go 1.2 Field Selectors and Nil Checks. July 2013.
- [doc/go13todo](https://go.dev/s/go13todo) Go 1.3 “To Do” List. Jun 1, 2014.
- [doc/goatgoogle](https://go.dev/talks/2012/splash.article#TOC_12.) Rob Pike. Go at Google - Language Semantics. October 25, 2012.
- [doc/makego](https://go.dev/talks/2015/how-go-was-made.slide)  Andrew Gerrand. How Go was Made. July 9, 2015.
- [discuss/go1preview](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) Russ Cox. Go 1 Preview. 2011.
- [design/overlapping-interfaces](https://go.dev/design/6977-overlapping-interfaces) Robert Griesemer. Proposal: Permit embedding of interfaces with overlapping method sets. Oct 16, 2019.
  + [issue/6977](https://go.dev/issue/6977) spec: allow embedding overlapping interfaces
- [design/struct-conversion](https://go.dev/design/16085-conversions-ignore-tags) Robert Griesemer. Proposal: Ignore tags in struct type conversions. June 16, 2016.
  + [issue/16085](https://go.dev/issue/16085) Proposal: Ignore tags in struct type conversions
- [design/go2trans](https://go.dev/design/28221-go2-transitions) Ian Lance Taylor. Proposal: Go 2 transition. October 15, 2018
- [design/signed-int-shift](https://go.dev/design/19113-signed-shift-counts) Robert Griesemer. Proposal: Permit Signed Integers as Shift Counts for Go 2. January 17, 2019
  + [issue/19113](https://go.dev/issue/19113) proposal: spec: allow signed shift counts
- [design/number-literal](https://go.dev/design/19308-number-literals) Russ Cox, Robert Griesemer. Proposal: Go 2 Number Literal Changes. March 6, 2019
  + [issue/12711](https://go.dev/issue/12711) proposal: different octal base literal representation
  + [issue/19308](https://go.dev/issue/19308) proposal: spec: binary integer literals
  + [issue/28493](https://go.dev/issue/28493) proposal: permit blank (_) separator in integer number literals
  + [issue/29008](https://go.dev/issue/29008) proposal: Go 2: hexadecimal floats
- [discuss/delete-return](https://www.reddit.com/r/golang/comments/5tfx7i/why_delete_doesnt_return_a_bool/ddmo4ug/?utm_source=share&utm_medium=web2x&context=3) why delete() doesn't return a bool ?
  + [issue/5147](https://go.dev/issue/5147) runtime: consider caching previous map hash value
  + [issue/51405](https://go.dev/issue/51405) proposal: Go 2: builtin: delete should return the deleted map item
- [issue/33502](https://go.dev/issue/33502) proposal: review meeting minutes
- [issue/33892](https://go.dev/issue/33892) proposal: Go 2 review meeting minutes
- [issue/19623](https://go.dev/issue/19623) proposal: spec: change int to be arbitrary precision
- [design/unsafearithmetic](https://docs.google.com/document/d/1yyCMzE4YPfsXvnZNjhszaYNqavxHhvbY-OWPqdzZK30/pub) Matthew Dempsky. Go 1.4: unsafe.Pointer arithmetic. Aug 2014.
- [issue/19367](https://go.dev/issue/19367) unsafe: add Slice(ptr *T, len anyIntegerType) []T
- [issue/40481](https://go.dev/issue/40481) unsafe: add Add function
- [issue/53003](https://go.dev/issue/53003) unsafe: add StringData, String, SliceData
- [issue/43615](https://go.dev/issue/43615) proposal: weak reference maps
- [issue/48105](https://go.dev/issue/48105) spec: clarify sequencing of function calls within expressions
- [issue/56351](https://go.dev/issue/56351) proposal: spec: add delete(m) to clear map
- [issue/58233](https://go.dev/issue/58233) proposal: spec: define return statement's result assignment order?
- [issue/57616](https://go.dev/issue/57616) proposal: spec: add simple string interpolation similar to Swift
- [issue/57411](https://go.dev/issue/57411) proposal: spec: define initialization order more precisely
- [issue/56103](https://go.dev/issue/56103) spec: disallow anonymous interface cycles
- [issue/25448](https://go.dev/issue/25448) spec: guarantee non-nil return value from recover
- [issue/61372](https://go.dev/issue/61372) proposal: spec: add untyped builtin zero
- [issue/56351](https://go.dev/issue/56351) spec: add clear(x) builtin, to clear map, zero content of slice
- [issue/58625](https://go.dev/issue/58625) unsafe: allow conversion of uintptr to unsafe.Pointer when it points to non-Go memory
- [discuss/48305](https://go.dev/issue/48305) doc comment revisions: headings, lists, and links
- [issue/12445](https://go.dev/issue/12445) spec: clarify how unsafe.Pointers may be used to determine offsets between addresses
- [issue/45624](https://go.dev/issue/45624) spec: expression to create pointer to simple types (new(expr) in Go 1.26)
- [issue/46505](https://go.dev/issue/46505) spec: allow conversion from slice to array
- [issue/9859](https://go.dev/issue/9859) spec: direct reference to embedded fields in struct literals

[Back To Top](#top)

### Slice (1.2)

- [design/read-only-slices](https://docs.google.com/document/d/1UKu_do3FRvfeN5Bb1RxLohV-zBOJWTzX0E8ZU1bkqX0/edit#heading=h.2wzvdd6vdi83) Brad Fitzpatrick. Read-only slices. May 13, 2013
- [design/read-only-slices-russ](https://docs.google.com/document/d/1-NzIYu0qnnsshMBpMPmuO21qd8unlimHgKjRD9qwp2A/edit) Russ Cox. Evaluation of read-only slices. May 2013.
- [design/go12slice](https://go.dev/s/go12slice) Russ Cox. Go Slice with Cap. June 2013.
- [design/multidim-slice](https://go.dev/design/6282-table-data) Brendan Tracey. Proposal: Multi-dimensional slices. November 17th, 2016.
- [issue/41239](https://go.dev/issue/41239) runtime: append growth sort of inconsistent and surprising
- [discuss/why-slice-grow](https://groups.google.com/g/golang-nuts/c/UaVlMQ8Nz3o) slices grow at 25% after 1024 but why 1024?
- [cl/347917](https://go.dev/cl/347917) runtime: make slice growth formula a bit smoother
  + [doc/cl-347917-graph](https://docs.google.com/document/d/1JQvV6vyAYdHhIboY-zAwK06OXZjxHrUhOFeG38MuJ94/edit?resourcekey=0-L5OsHqwZZBxvjfK0dwsyVQ) Graphs for CL 347917
- [cl/418554](https://go.dev/cl/418554) cmd/compile,runtime: redo growslice calling convention

[Back To Top](#top)

### Package Management (1.4, 1.5, 1.7)

- [design/go14internal](https://go.dev/s/go14internal) Russ Cox. Go 1.4 “Internal” Packages. June 2014.
- [design/go14nopkg](https://go.dev/s/go14nopkg) Russ Cox. Go 1.4 src/pkg → src. June 2014.
- [design/go14customimport](https://go.dev/s/go14customimport) Russ Cox. Go 1.4 Custom Import Path Checking. June 2014.
- [design/go15vendor](https://go.dev/s/go15vendor) Russ Cox. Go 1.5 Vendor Experiment. July 2015
- [design/go17binarypkg](https://go.dev/design/2775-binary-only-packages) Russ Cox. Proposal: Binary-Only Packages. April 24, 2016
  + [issue/2775](https://go.dev/issue/2775) cmd/go: work when binaries are available but source is missing

[Back To Top](#top)

### Type alias (1.9)

- [design/type-alias](https://go.dev/design/18130-type-alias) Russ Cox, Robert Griesemer. Proposal: Type Aliases. December 16, 2016
  + [talk/type-alias](https://www.youtube.com/watch?v=t-w6MyI2qlU) GopherCon 2016 - Lightning Talk: Robert Griesemer - Alias Declarations for Go, A proposal. Oct 9, 2016
  + [issue/16339](https://go.dev/issue/16339) proposal: Alias declarations for Go
  + [issue/18130](https://go.dev/issue/18130) all: support gradual code repair while moving a type between packages
- [talk/refactor-video](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Russ Cox. Codebase Refactoring (with help from Go). GothamGo 2016. November 18, 2016.
  + [doc/refactor](https://go.dev/talks/2016/refactor.article) Russ Cox. Codebase Refactoring (with help from Go).
- [design/alias-decls](https://go.dev/design/16339-alias-decls) Robert Griesemer & Rob Pike. Proposal: Alias declarations for Go. July 18, 2016.
- [issue/76472](https://go.dev/issue/76472) remove gotypesalias, asynctimerchan GODEBUG flag for 1.27
- [issue/17487](https://go.dev/issue/17487) cmd/compile: support alias
- [issue/67547](https://go.dev/issue/67547) go/types, types2: audit type assertions for potential Alias bugs
- [issue/65294](https://go.dev/issue/65294) x/tools: get tests passing with GODEBUG=gotypesalias=1
- [doc/alias-names](https://go.dev/blog/alias-names) Robert Griesemer. What's in an (Alias) Name? September 17, 2024.

[Back To Top](#top)

### Defer (1.14)

- [design/open-defer](https://go.dev/design/34481-opencoded-defers) Dan Scales, Keith Randall, and Austin Clements. Proposal: Low-cost defers through inline code, and extra funcdata to manage the panic case. 2019-09-23
  + [issue/6980](https://go.dev/issue/6980) cmd/compile: allocate some defers in stack frames
  + [issue/14939](https://go.dev/issue/14939) runtime: defer is slow #14939
- Unsolved `defer recover()` edge cases.
  - [issue/10458](https://go.dev/issue/10458) spec: panicking corner-case semantics
  - [issue/23531](https://go.dev/issue/23531) spec: recover() in nested defer
  - [issue/26275](https://go.dev/issue/26275) runtime: document the behaviour of Caller and Callers when used in deferred functions
  - [issue/34530](https://go.dev/issue/34530) spec: clarify when calling recover stops a panic
  - [cl/189377](https://go.dev/cl/189377) spec: specify the behavior of recover and recursive panics in more detail
- [issue/34481](https://go.dev/issue/34481) cmd/compile, runtime: a design for low-cost defers via inline code

[Back To Top](#top)

### Error values

Error handling includes two separate works: error values and error formatting. Historically, there was a try proposal that had been investigated for a year, however, due to its lack of simplicity and impact to stack tracing, it was rejected. In 1.13, error values received a large revision, and the package `errors` includes different APIs, such as `errors.Is`. Until recently (1.21), #53435 allows wrapping multiple errors, and there is a community repository that tries to implement the original try proposal.

- [doc/err2011](https://go.dev/blog/error-handling-and-go) Andrew Gerrand. Error handling in Go. July 2011.
- [doc/err-values](https://go.dev/blog/errors-are-values) Rob Pike. Errors are values. January 2015.
- [doc/err-philosophy](https://dave.cheney.net/paste/gocon-spring-2016.pdf) Dave Cheney. My philosophy for error handling. April 2016.
- [doc/err-gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) Dave Cheney. Don’t just check errors, handle them gracefully. April 2016.
- [doc/err-stacktrace](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package) Dave Cheney. Stack traces and the errors package. June, 12 2016.
- [doc/err-upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) Rob Pike. Error handling in Upspin. December 06, 2017.
- [doc/err-work](https://go.dev/blog/go1.13-errors) Damien Neil and Jonathan Amsterdam. Working with Errors in Go 1.13. October 17, 2019.
- [design/err-handling-overview](https://go.dev/design/go2draft-error-handling-overview) Russ Cox. Error Handling — Problem Overview.
- [doc/err-value-faq](https://github.com/golang/go/wiki/ErrorValueFAQ) Jonathan Amsterdam and Bryan C. Mills. Error Values: Frequently Asked Questions. August 2019.
- [design/err-handle-check](https://go.dev/design/go2draft-error-handling) Marcel van Lohuizen. Error Handling — Draft Design. August 27, 2018.
- [design/err-try](https://go.dev/design/32437-try-builtin) Robert Griesemer. Proposal: A built-in Go error check function, "try". Jun 12, 2019.
  - [issue/32437](https://go.dev/issue/32437#issuecomment-512035919) Proposal: A built-in Go error check function, "try". Decision response.
- [issue/71203](https://go.dev/issue/71203) proposal: spec: reduce error handling boilerplate using ?
- [design/err-inspect](https://go.dev/design/go2draft-error-inspection) Jonathan Amsterdam, Damien Neil. Error Inspection — Draft Design. August 27, 2018.
- [design/err-values-overview](https://go.dev/design/go2draft-error-values-overview) Russ Cox. Error Values — Problem Overview. August 27, 2018.
- [design/error-values](https://go.dev/design/29934-error-values) Jonathan Amsterdam, Russ Cox, Marcel van Lohuizen, Damien Neil. Proposal: Go 2 Error Inspection. January 25, 2019
  + [issue/29934](https://go.dev/issue/29934) Jonathan Amsterdam. proposal: Go 2 error values. Jan 25, 2019.
  + [issue/29934-decision](https://go.dev/issue/29934#issuecomment-489682919) Damien Neil. Go 1.13 lunch decision about error values. May 6, 2019.
  + [issue/29934-russ](https://go.dev/issue/29934#issuecomment-490087200) Russ Cox. Response, Response regarding "proposal: Go 2 error values". May 7, 2019.
- [design/err-print](https://go.dev/design/go2draft-error-printing) Marcel van Lohuizen. Error Printing — Draft Design. August 27, 2018.
  + [issue/30468](https://go.dev/issue/30468) errors: performance regression in New.
- [issue/40432](https://go.dev/issue/40432) language: Go 2: error handling meta issue
- [issue/40776](https://go.dev/issue/40776) proposal: dynamic ignored error detector
- [issue/41198](https://go.dev/issue/41198) proposal: errors: add ErrUnimplemented as standard way for interface method to fail.
- [issue/47811](https://go.dev/issue/47811) proposal: errors: add Errors as a standard way to represent multiple errors as a single error
- [issue/53435](https://go.dev/issue/53435) proposal: wrapping multiple errors
- [repo/try](https://github.com/dsnet/try) Try: Simplified Error Handling in Go
- [discuss/71460](https://github.com/golang/go/discussions/71460) Ian Lance Taylor. discussion: spec: reduce error handling boilerplate using ? January 27, 2025.

[Back To Top](#top)

### Channel/Select

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov. Go channels on steroids. Jan 28, 2014
  + [issue/8899](https://go.dev/issue/8899) runtime: lock-free channels
  + [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) update on "lock-free channels"
  + [cl/112990043](https://codereview.appspot.com/112990043/) runtime: fine-grained locking in select
  + [cl/110580043](https://codereview.appspot.com/110580043/) runtime: add fast paths to non-blocking channel operations
- [issue/8898](https://go.dev/issue/8898) runtime: special case timer channels
  + [issue/61542](https://go.dev/issue/61542) time: stop requiring Timer/Ticker.Stop for prompt GC
- [issue/9120](https://go.dev/issue/9120) runtime: remove implementation restriction on channel element size
- [issue/37196](https://go.dev/issue/37196) time: make Timer/Ticker channels not receivable with old values after Stop or Reset returns
- [issue/8903](https://go.dev/issue/8903) runtime: make chan-based generators faster.
- [issue/21806](https://go.dev/issue/21806) fairness in select statement.
- [issue/40410](https://go.dev/issue/40410) runtime: trim unnecessary fields from scase
- [issue/40641](https://go.dev/issue/40641) runtime: race between stack shrinking and channel send/recv leads to bad sudog values
- [issue/37350](https://go.dev/issue/37350) reflect: Select panics if array length greater than 1<<16
- [doc/no-err-syntax](https://go.dev/blog/error-syntax) Robert Griesemer. [ On | No ] syntactic support for error handling. June 2025.

[Back To Top](#top)

### Generics

- [doc/generics-discuss](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/edit#heading=h.vuko0u3txoew) Summary of Go Generics Discussions
- [doc/generics-dilemma](https://research.swtch.com/generic) Russ Cox. The Generic Dilemma. December 3, 2009.
- [design/type-functions](https://go.dev/design/15292/2010-06-type-functions) Ian Lance Taylor. Type Functions. June 2010.
- [design/generalized-types](https://go.dev/design/15292/2011-03-gen) Ian Lance Taylor. Generalized Types. March 2011.
- [design/code-gen](https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM) Russ Cox. Alternatives to Dynamic Code Generation in Go. September 2012.
- [design/generalized-types2](https://go.dev/design/15292/2013-10-gen) Ian Lance Taylor. Generalized Types In Go. October 2013.
- [design/type-parameters](https://go.dev/design/15292/2013-12-type-params) Ian Lance Taylor. Type Parameters in Go. December 2013.
- [design/compile-time-function](https://go.dev/design/15292/2016-09-compile-time-functions) Bryan C. Mills. Compile-time Functions and First Class Types. September 2016.
- [design/should-generics](https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics) Ian Lance Taylor. Go should have generics. January 2011.
- [design/should-generics2](https://go.dev/design/15292-generics) Ian Lance Taylor. Go should have generics. Updated: April 2016.
  + [issue/15292](https://go.dev/issue/15292) proposal: spec: generic programming facilities
- [design/generics-overview](https://go.dev/design/go2draft-generics-overview) Russ Cox. Generics — Problem Overview. August 27, 2018.
- [design/contracts](https://go.dev/design/go2draft-contracts) Ian Lance Taylor, Robert Griesemer. Contracts — Draft Design. August 27, 2018, Updated: July 31, 2019.
  + [cl/187317](https://go.dev/cl/187317/) Implementation prototype.
  + [paper/featherweight-go](https://arxiv.org/abs/2005.11710) Griesemer, Robert, et al. Featherweight Go. arXiv preprint arXiv:2005.11710 (2020).
  + [talk/featherweight-go](https://www.youtube.com/watch?v=Dq0WFigax_c) Phil Wadler: Featherweight Go. Jun 8, 2020.
- [design/type-parameters2](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md) Ian Lance Taylor, Robert Griesemer. Type Parameters - Draft Design. June 16, 2020, Updated: August 28, 2020.
  + [cl/dev.go2go](https://github.com/golang/go/blob/dev.go2go/README.go2go.md) dev.go2go branch
  + [doc/type-check-readme](https://github.com/golang/go/tree/dev.go2go/src/go/types) type checking.
  + [doc/type-check-notes](https://github.com/golang/go/blob/dev.go2go/src/go/types/NOTES) This file serves as a notebook/implementation log.
- [discuss/generics-parenthesis](https://groups.google.com/g/golang-nuts/c/7t-Q2vt60J8) Robert. Generics and parenthesis.
- [discuss/generics-move](https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw) Ian, Moving forward with the generics design.
- [discuss/generics-implementation](https://groups.google.com/g/golang-dev/c/OcW0ATRS4oM) Implementing Generics
- [design/generics-implementation-stenciling](https://go.dev/design/generics-implementation-stenciling) Generics implementation - Stenciling
- [design/generics-implementation-dictionaries](https://go.dev/design/generics-implementation-dictionaries) Generics implementation - Dictionaries
- [design/generics-implementation-gcshape](https://go.dev/design/generics-implementation-gcshape) Generics implementation - GC Shape Stenciling
- [design/generics-implementation-dictionaries-go1.18](https://go.dev/design/generics-implementation-dictionaries-go1.18) Go 1.18 Implementation of Generics via Dictionaries and Gcshape Stenciling
- [issue/43651](https://go.dev/issue/43651) proposal: spec: add generic programming using type parameters
- [design/type-parameters3](https://go.dev/design/43651-type-parameters) Type Parameters Proposal
- [issue/45346](https://go.dev/issue/45346) proposal: spec: generics: use type sets to remove type keyword in constraints
- [issue/46477](https://go.dev/issue/46477) proposal: spec: generics: type parameters on aliases
- [issue/38907](https://go.dev/issue/38907) proposal: spec: disallow impossible interface-interface type assertions
- [discussion/48287](https://github.com/golang/go/discussions/48287) how to update APIs for generics
- Generics behavior
  + [issue/48274](https://go.dev/issue/48274) doc: update FAQ for the arrival of parametric polymorphism
  + [issue/50202](https://go.dev/issue/50202) spec: document that type parameters are not permitted as constant declaration types
  + [issue/50272](https://go.dev/issue/50272) spec: function type inference ignores type parameter constraints
  + [issue/50954](https://go.dev/issue/50954) go/types, types2: review error messages
  + [issue/51110](https://go.dev/issue/51110) spec: document behavior of type switches containing type parameter cases
  + any/comparable
    * [issue/33232](https://go.dev/issue/33232) spec: allow 'any' for 'interface{}' in non-constraint contexts
    * [issue/46746](https://go.dev/issue/46746) reflect: add Value.Equal, Value.Comparable
    * [issue/49587](https://go.dev/issue/49587) proposal: spec: add comparable w/o interfaces
    * [issue/49927](https://go.dev/issue/49927) builtin: add documentation for any and comparable to pseudo-package builtin
    * [issue/49884](https://go.dev/issue/49884) all: rewrite interface{} to any
    * [issue/50646](https://go.dev/issue/50646) spec: document/explain which interfaces implement comparable
    * [issue/50646](https://go.dev/issue/50791) spec: document definition of comparable
    * [issue/51257](https://go.dev/issue/51257) spec: any no longer implements comparable
    * [issue/51338](https://go.dev/issue/51338) proposal: spec: permit values to have type "comparable"
    * [issue/52474](https://go.dev/issue/52474) proposal: spec: permit non-interface types that support == to satisfy the comparable constraint
    * [issue/52509](https://go.dev/issue/52509) proposal: spec: allow interface types to instantiate comparable type parameters
    * [issue/52531](https://go.dev/issue/52531) proposal: spec: add new constraint kind satisfied by types that support == (including interface types)
    * [issue/52614](https://go.dev/issue/52614) proposal: type parameters are comparable unless they exclude comparable types
    * [issue/52624](https://go.dev/issue/52624) proposal: the comparable interface represents the comparable subset of run-time values
- Standard packages using generics
  + [issue/45458](https://go.dev/issue/45458) constraints: new package to define standard type parameter constraints
  + [discuss/47319](https://github.com/golang/go/discussions/47319) proposal: constraints: new package to define standard type parameter constraints (discussion)
  + [issue/45955](https://go.dev/issue/45955) slices: new package to provide generic slice functions
  + [discuss/47203](https://github.com/golang/go/discussions/47203) proposal: slices: new package to provide generic slice functions (discussion)
  + [discuss/47331](https://github.com/golang/go/discussions/47331) proposal: container/set: new package to provide a generic set type (discussion)
  + [discuss/47330](https://github.com/golang/go/discussions/47330) proposal: maps: new package to provide generic map functions (discussion)
  + [issue/47649](https://go.dev/issue/47649) maps: new package to provide generic map functions
  + [issue/57436](https://go.dev/issue/57436) maps: new standard library package based on x/exp/maps
  + [issue/47657](https://go.dev/issue/47657) proposal: sync, sync/atomic: add PoolOf, MapOf, ValueOf
  + [issue/50792](https://go.dev/issue/50792) constraints: move to x/exp for Go 1.18
  + [issue/53427](https://go.dev/issue/53427) proposal: x/sync/singleflight: add generic version
- Related packages
  + [issue/47781](https://go.dev/issue/47781) go/ast, go/token: additions to support type parameters
  + [issue/47916](https://go.dev/issue/47916) go/types: additions to support type parameters
  + [issue/48525](https://go.dev/issue/48525) x/tools/go/ssa: generics support
  + [issue/50887](https://go.dev/issue/50887) go/types, types2: document predicates on generic types
- [issue/52654](https://go.dev/issue/52654) spec: provide some way to refer to a generic function without instantiation
- [issue/59488](https://go.dev/issue/59488) cmp: new package with Ordered, Compare, Less; min, max as builtins
- [issue/60648](https://go.dev/issue/60648) cmp: document ordering for floating point values in cmp.Ordered
- [issue/59531](https://go.dev/issue/59531) all: ordering numeric values, including NaNs, and the consequences thereof
- [issue/60519](https://go.dev/issue/60519) math: remove Compare, Compare32
- [issue/58650](https://go.dev/issue/58650) spec: a general approach to type inference
- [issue/59338](https://go.dev/issue/59338) spec: infer type arguments from assignments of generic functions (reverse type inference)
- [issue/58671](https://go.dev/issue/58671) spec: type inference should be more lenient about untyped numeric literals
- [issue/57644](https://go.dev/issue/57644) proposal: spec: sum types based on general interfaces
- [issue/70128](https://go.dev/issue/70128) spec: remove notion of core types
  + [issue/60994](https://go.dev/issue/60994) proposal: go/types: add func CommonUnderlying(Type) Type, (formerly CoreType)
- [issue/56548](https://go.dev/issue/56548) spec: allow basic interface types to instantiate comparable type parameters
- [design/parameterized-go-types](https://go.dev/design/47916-parameterized-go-types) Rob Findley, Robert Griesemer. Additions to go/types to support type parameters. 2021-08-17.
- [issue/12854](https://go.dev/issue/12854) spec: type inferred composite literals
  + [issue/77245](https://go.dev/issue/77245) spec: function type inference should work in all assignment contexts
- [issue/61035](https://go.dev/issue/61035) go/types, types2: fix SizesFor("gc", ...) to match actual gc behavior
- [issue/52728](https://go.dev/issue/52728) go/types: guarantee that a finite number of distinct instances are reachable via the API
- [issue/57192](https://go.dev/issue/57192) spec: interfaces that have different-size method sets never unify (but they possibly could)
- [issue/63974](https://go.dev/issue/63974) go/types, types2: switch to using go/version once we can use it in the compiler
- [issue/51343](https://go.dev/issue/51343) go/types, types2: simplify type-checking of method receivers with type parameters
- [issue/55326](https://go.dev/issue/55326) go/types, types2: remove discrepancy between type-checker and compiler error messages
- [issue/54511](https://go.dev/issue/54511) types2, go/types: share tests
- [issue/69673](https://go.dev/issue/69673) go/types: position-independent type checking
- [issue/57410](https://go.dev/issue/57410) cmd/compile: remove GOEXPERIMENT=nounified
- [issue/68778](https://go.dev/issue/68778) exportdata: support type parameterized aliases

[Back To Top](#top)
- [issue/60353](https://go.dev/issue/60353) spec: when unifying against interface types, consider common methods
- [issue/57433](https://go.dev/issue/57433) slices: new standard library package based on x/exp/slices
- [issue/51259](https://go.dev/issue/51259) proposal: spec: support for struct members in interface/constraint syntax
- [issue/50285](https://go.dev/issue/50285) proposal: generic should infer type from variable definition
- [issue/66651](https://go.dev/issue/66651) proposal: spec: variadic type parameters
- [issue/47796](https://go.dev/issue/47796) types2, go/types: self-referential type parameter list is not accepted (allowed in Go 1.26)
  + [issue/52190](https://go.dev/issue/52190) proposal: generics: self referencing type constraints
- [issue/75526](https://go.dev/issue/75526) proposal: allow type parameters on methods (generic methods)
  + [issue/77273](https://go.dev/issue/77273) spec: generic methods for Go
- [doc/comparable](https://go.dev/blog/comparable) Robert Griesemer. All your comparable types. February 17, 2023.
- [doc/deconstructing-type-parameters](https://go.dev/blog/deconstructing-type-parameters) Ian Lance Taylor. Deconstructing Type Parameters. September 26, 2023.
- [doc/type-inference](https://go.dev/blog/type-inference) Robert Griesemer. Everything You Always Wanted to Know About Type Inference - And a Little Bit More. October 9, 2023.
- [doc/generic-slice-functions](https://go.dev/blog/generic-slice-functions) Valentin Deleplace. Robust generic functions on slices. February 22, 2024.
- [doc/coretypes](https://go.dev/blog/coretypes) Robert Griesemer. Goodbye core types - Hello Go as we know and love it! March 26, 2025.
- [doc/generic-interfaces](https://go.dev/blog/generic-interfaces) Axel Wagner. Generic interfaces. July 7, 2025.
- [issue/75883](https://go.dev/issue/75883) spec: remove cycle restriction for type parameters
- [doc/type-construction](https://go.dev/blog/type-construction-and-cycle-detection) Mark Freeman. Type Construction and Cycle Detection. March 24, 2026.
- [doc/generic-methods](https://go.dev/blog/generic-methods) Mark Freeman. Generic Methods. August 26, 2026.
- [talk/axel2024](https://www.youtube.com/watch?v=dab3I-HcTVk) Axel Wagner. Advanced Generics Patterns. GopherCon 2024, Chicago. July 2024.

### Iterator

- [issue/24282](https://go.dev/issue/24282) proposal: Go 2: improve for-loop ergonomics
- [issue/40605](https://go.dev/issue/40605) proposal: Go 2: iterators
- [issue/43557](https://go.dev/issue/43557) proposal: Go 2: function values as iterators
- [issue/47707](https://go.dev/issue/47707) proposal: Go 2: spec: for range with defined types
- [issue/48567](https://go.dev/issue/48567) proposal: use channels as iterators
- [issue/50112](https://go.dev/issue/50112) proposal: package collection and iterator design for go
- [issue/54047](https://go.dev/issue/54047) proposal: Go 2: add a new iterator syntax, package, interfaces
- [discuss/54245](https://go.dev/issue/54245) discussion: standard iterator interface
- [issue/20733](https://go.dev/issue/20733) proposal: spec: redefine range loop variables in each iteration
- [issue/56010](https://go.dev/issue/56010) redefining for loop variable semantics
- [issue/60078](https://go.dev/issue/60078) spec: less error-prone loop variable scoping
- [discuss/56413](https://go.dev/issue/56413) user-defined iteration using range over func values
- [issue/61405](https://go.dev/issue/61405) proposal: spec: add range over int, range over func
- [doc/coro](https://research.swtch.com/coro) Coroutines for Go
- [doc/undefined-behavior](https://research.swtch.com/ub) C and C++ Prioritize Performance over Correctness
- [issue/54650](https://go.dev/issue/54650) x/exp/slices: memory leak on Delete
- [issue/66455](https://go.dev/issue/66455) proposal: errors: add All and AllAs iterators
- [issue/61898](https://go.dev/issue/61898) proposal: x/exp/xiter: new package with iterator adapters
- [issue/65236](https://go.dev/issue/65236) spec: allow range-over-func to omit iteration variables
- [issue/61899](https://go.dev/issue/61899) slices: add iterator-related functions
- [issue/61897](https://go.dev/issue/61897) iter: new package for iterators
- [issue/61900](https://go.dev/issue/61900) maps: add iterator-related functions
- [issue/61901](https://go.dev/issue/61901) bytes, strings: add iterator forms of existing functions
- [issue/65237](https://go.dev/issue/65237) spec: document range-over-func
- [design/loopvar](https://go.dev/design/60078-loopvar) Proposal: Less Error-Prone Loop Variable Scoping.
- [issue/57969](https://go.dev/issue/57969) cmd/compile: add GOEXPERIMENT=loopvar
- [issue/61717](https://go.dev/issue/61717) cmd/compile: add new range behind GOEXPERIMENT=range
- [doc/loopvar-preview](https://go.dev/blog/loopvar-preview) David Chase and Russ Cox. Fixing For Loops in Go 1.22. September 19, 2023.
- [doc/range-functions](https://go.dev/blog/range-functions) Ian Lance Taylor. Range Over Function Types. August 20, 2024.
- [issue/61902](https://go.dev/issue/61902) regexp: add iterator forms of matching methods

[Back To Top](#top)

### Compatibility

- [doc/go1compat](https://go.dev/doc/go1compat) Go 1 and the Future of Go Programs
- [discussion/55090](https://go.dev/issue/55090) extending Go backward compatibility
  + [issue/56986](https://go.dev/issue/56986) extended backwards compatibility for Go
  + [design/godebug](https://go.dev/design/56986-godebug) Proposal: Extended backwards compatibility for Go
- [discussion/55092](https://go.dev/issue/55092) extending Go forward compatibility
  + [issue/57001](https://go.dev/issue/57001) extended forwards compatibility for Go
  + [design/gotoolchain](https://go.dev/design/57001-gotoolchain) Proposal: Extended forwards compatibility in Go
- [issue/33791](https://go.dev/issue/33791) proposal: add explicit decision doc for large changes
- [issue/28221](https://go.dev/issue/28221) cmd/go, cmd/compile: record language version to support language transitions
- [issue/31793](https://go.dev/issue/31793) go/types, types2: provide configuration support for different language versions
- [doc/compat](https://go.dev/blog/compat) Russ Cox. Backward Compatibility, Go 1.21, and Go 2. August 14, 2023.
- [doc/toolchain](https://go.dev/blog/toolchain) Russ Cox. Forward Compatibility and Toolchain Management in Go 1.21. August 14, 2023.
- [issue/65573](https://go.dev/issue/65573) cmd/go: separate default GODEBUGs from go language version
- [issue/76163](https://go.dev/issue/76163) policy for removing GODEBUG flags
- [discuss/80580](https://github.com/golang/go/discussions/80580) Austin Clements. a more robust proposal process. July 27, 2026.




## Compiler Toolchain

### Office Hours

- [discuss/runtime-office-hours](https://docs.google.com/document/d/17YYCLhsyoGx7wLVGUET2VvHog9oJE4xLCljUGmJCucU/edit) compiler & runtime office hours
  + [talk/runtime-office-meeting1](https://www.youtube.com/watch?v=p69c_R-YdFk&ab_channel=TheGoProgrammingLanguage) Go compiler & runtime office hours 2022-12-19
  + [talk/runtime-office-meeting2](https://www.youtube.com/watch?v=x7-5qlTW0J8&ab_channel=TheGoProgrammingLanguage) Go compiler & runtime office hours 2023-01-19

### Compiler

- [code/gc0initial](https://github.com/golang/go/tree/cb87526ce3531557ccf69969de4c8018956b10b5/src/c) Ken Thompson. Go0 compiler initial version. 28 Mar 2008.
- [code/6g](https://github.com/golang/go/commit/0cafb9ea3d3d34627e8f492ccafa6ba9b633a213) Rob Pike. the first complete Go0 compiler. 4 Jun 2008.
- [design/go12symtab](https://go.dev/s/go12symtab) Russ Cox. Go 1.2 Runtime Symbol Information. July 2013.
- [design/go13compiler](https://go.dev/s/go13compiler) Russ Cox. Go 1.3+ Compiler Overhaul. December 2013
- [design/go14generate](https://go.dev/s/go1.4-generate) Rob Pike. Go generate: A Proposal
- [design/dev.cc](https://go.dev/s/dev.cc)  Russ Cox. dev.cc branch plan. Nov 2014
- [design/go15bootstrap](https://go.dev/s/go15bootstrap) Russ Cox. Go 1.5 Bootstrap Plan. January 2015.
- [doc/escape-analysis](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/edit#) Dmitry Vyukov. Go Escape Analysis Flaws. Feb 10, 2015.
- [design/execmodes](https://go.dev/s/execmodes) Ian Lance Taylor. Go Execution Modes. August, 2014 (updated January 2016)
- [design/go17ssa](https://go.dev/s/go17ssa) Keith Randall. New SSA Backend for the Go Compiler. 2/10/2015.
- [doc/compiler-optimization](https://github.com/golang/go/wiki/CompilerOptimizations) Compiler And Runtime Optimizations.
- [issue/6853](https://go.dev/issue/6853) all: binaries too big and growing.
- [design/go19inlining](https://go.dev/design/19348-midstack-inlining) David Lazar, Austin Clements. Proposal: Mid-stack inlining in the Go compiler.
  + [issue/19348](https://go.dev/issue/19348) cmd/compile: enable mid-stack inlining.
  + [talk/go19inliningtalk](https://go.dev/s/go19inliningtalk) David Lazar. Mid-stack inlining in the Go compiler
- [design/go122inlining](https://docs.google.com/document/d/1a6p7-nbk5PVyM1S2tmccFrrIuGzCyzclstBtaciHxVw/edit#heading=h.o82v9r2rjlej) Go 1.22 inlining overhaul
- [design/dwarf-inlining](https://go.dev/design/22080-dwarf-inlining) Than McIntosh. Proposal: emit DWARF inlining info in the Go compiler. 2017-10-23.
  + [issue/22080](https://go.dev/issue/22080) cmd/compile: generated DWARF inlining info
- [issue/23109](https://go.dev/issue/23109) cmd/compile: rewrite escape analysis.
- [issue/27167](https://go.dev/issue/27167) cmd/compile: rename a bunch of things
  - [doc/renames](https://docs.google.com/document/d/19_ExiylD9MRfeAjKIfEsMU1_RGhuxB9sA0b5Zv7byVI/edit) Proposed Go 1.12 toolchain renames
- GOEXPERIMENT=checkptr
  + [issue/22218](https://go.dev/issue/22218) proposal: add GOEXPERIMENT=checkptr
  + [issue/34964](https://go.dev/issue/34964) cmd/compile: enable -d=checkptr as part of -race and/or -msan?
  + [issue/34972](https://go.dev/issue/34972) all: get standard library building with -d=checkptr
  + [discuss/checkptr](https://groups.google.com/forum/#!msg/golang-dev/SzwDoqoRVJA/Iozu8vWdDwAJ)
- [issue/37121](https://go.dev/issue/37121) runtime: unaligned jumps causing performance regression on Intel
- [issue/16798](https://go.dev/issue/16798) proposal: cmd/compile: add tail call optimization for self-recursion only. (declined)
- [issue/22624](https://go.dev/issue/22624) proposal: Go 2: add become statement to support tail calls (declined)
- [design/64align](https://go.dev/design/36606-64-bit-field-alignment) Dan Scales. Proposal: Make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed, //go:align directives. 2020-06-08.
  + [issue/599](https://go.dev/issue/599) cmd/compile: make 64-bit fields 64-bit aligned on 32-bit systems
  + [issue/36606](https://go.dev/issue/36606) proposal: cmd/compile: make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed directive on structs
  + [issue/19057](https://go.dev/issue/19057) proposal: runtime: add AlignedN types that can be used to increase alignment
- [talk/gccgo](https://www.youtube.com/watch?v=U0w9eFunkX4) Brief overview of gccgo, "the other" Go compiler. Aug 6, 2015.
- [issue/54534](https://go.dev/issue/54534) cmd/compile: design doc explaining unified IR implementation
- [discuss/53060](https://go.dev/issue/53060) discussion: clarify Go support policy for secondary ports
- [issue/53383](https://go.dev/issue/53383) clarify Go support policy for secondary ports
- **Profile-guided Optimization**
  - [paper/autofdo](https://research.google/pubs/pub45290/) Chen, et al. AutoFDO: Automatic Feedback-Directed Optimization for Warehouse-Scale Applications. 2016.
  - [doc/pgo](https://go.dev/doc/pgo) Profile-guided optimization
  - [issue/28262](https://go.dev/issue/28262) cmd/compile: feedback-guided optimization
  - [issue/55022](https://go.dev/issue/55022) proposal: cmd/compile: profile-guided optimization
  - [issue/55025](https://go.dev/issue/55025) proposal: design and implementation of Profile-Guided Optimization (PGO)
  - [design/pgo](https://go.dev/design/55022-pgo) Proposal: profile-guided optimization
  - [design/pgo-implementation](https://go.dev/design/55022-pgo-implementation) Design and Implementation of Profile-Guided Optimization (PGO) for Go
  - [issue/62463](https://go.dev/issue/62463) cmd/compile: PGO opportunities umbrella issue
  - [talk/pgo](https://www.youtube.com/watch?v=0GkRJQ1aUOI) Michael Pratt, Boost application performance with profile guided optimization, GopherConUK 2024
- **Toolchain Telemetry**
  - [issue/58409](https://go.dev/issue/58409) telemetry in the Go toolchain
  - [issue/58894](https://go.dev/issue/58894) all: add opt-in transparent telemetry to Go toolchain
  - [doc/telemetry-sample](https://research.swtch.com/sample) The Magic of Sampling, and its Limitations
  - [doc/telemetry-intro](https://research.swtch.com/telemetry-intro) Transparent Telemetry for Open-Source Projects.
  - [doc/telemetry-design](https://research.swtch.com/telemetry-design) The Design of Transparent Telemetry.
  - [doc/telemetry-uses](https://research.swtch.com/telemetry-uses) Use Cases for Transparent Telemetry.
  - [doc/telemetry-optin](https://research.swtch.com/telemetry-opt-in) Opting In to Transparent Telemetry.
- [discuss/reducing-gc-peak-memory](https://groups.google.com/g/golang-dev/c/iOF44IL_i1M/m/CWI9U8XFKsEJ) Reducing gc peak memory use (golang-dev)
- [discuss/bug-in-new-amd64](https://groups.google.com/g/golang-dev/c/2oHi8t1s0Z8) Bug in new amd64 ABI (golang-dev)
- [issue/51225](https://go.dev/issue/51225) cmd/compile: require build systems to specify -importcfg
- [issue/18616](https://go.dev/issue/18616) math/bits: an integer bit twiddling library
- [issue/74299](https://go.dev/issue/74299) runtime, cmd/compile: add runtime.freegc and runtime.freegcTracked to reduce GC work (GOEXPERIMENT=runtimefreegc)
- [issue/62737](https://go.dev/issue/62737) cmd/compile: merge temporaries and named variables by GC shape
- [discuss/jump-tables-in-6a](https://groups.google.com/g/golang-nuts/c/6M0BLgrKBRk) Jump Tables in 6a (golang-nuts)
- [issue/61515](https://go.dev/issue/61515) testing: add testing.B.Loop for iteration
  + [discuss/go-1-24-blooper](https://groups.google.com/g/golang-nuts/c/Z017c0f59vY) Go 1.24 BLooper (golang-nuts)
- [discuss/get-rid-of-stack](https://groups.google.com/g/golang-nuts/c/DbmqfDlAR0U/m/elFowQ_z_UsJ) get rid of stack check prolog (golang-nuts)
- [discuss/fused-multiply-add](https://groups.google.com/g/golang-dev/c/cVnE1K08Aks) Fused multiply-add? (golang-dev)
- [discuss/rfc-go-assembly-for](https://groups.google.com/g/golang-dev/c/rWgDxCrL4GU) [RFC] Go assembly for ARM64 (golang-dev)
- [discuss/x86-asm-source-of](https://groups.google.com/g/golang-dev/c/5fmxhPD3MwM/m/IMqWPRXJBAAJ) x86 asm source of truth (golang-dev)
- [issue/71105](https://go.dev/issue/71105) cmd/compile: change GORISCV64=rva20u64 to include compressed instructions
  + [issue/61476](https://go.dev/issue/61476) all: add GORISCV64 environment variable
- [issue/46229](https://go.dev/issue/46229) all: port to linux/loong64
- [design/wasmexport](https://go.dev/design/42372-wasmexport) Francesco Guardiani. Proposal: go:wasmexport directive. 2020-12-17.
  + [issue/42372](https://go.dev/issue/42372) proposal: cmd/compile: go:wasmexport directive
- [issue/60072](https://go.dev/issue/60072) cmd/go: add GOMIPS32, GOMIPS64 ISA levels (iii, r1, r2, r5, r6)
- [issue/73787](https://go.dev/issue/73787) simd/archsimd: architecture-specific SIMD intrinsics under a GOEXPERIMENT
  + [issue/79781](https://go.dev/issue/79781) simd/archsimd: support ARM64 SVE SIMD intrinsics under a GOEXPERIMENT
- [issue/78979](https://go.dev/issue/78979) proposal: simd/archsimd: enable AMD64 architecture-specific SIMD by default
  + [issue/76473](https://go.dev/issue/76473) simd: rename low-level SIMD package to simd/archsimd (GOEXPERIMENT)
- [issue/13241](https://go.dev/issue/13241) cmd/compile: switch to new export format
- [issue/14473](https://go.dev/issue/14473) cmd/compile: investigate replacing NodeList with slice
- [issue/15756](https://go.dev/issue/15756) cmd/compile: parallelize compilation
- [issue/22326](https://go.dev/issue/22326) cmd/compile: incorrect package initialization order for spec example
- [issue/19683](https://go.dev/issue/19683) cmd/compile: eliminate base.Pos and ir.CurFunc
- [issue/51691](https://go.dev/issue/51691) cmd/compile: remove support for "broken" IR and types
- [issue/60234](https://go.dev/issue/60234) cmd/compile: update compiler's PGO profile
- [issue/15734](https://go.dev/issue/15734) cmd/go: build dependent packages as soon as export data is ready
- [issue/51734](https://go.dev/issue/51734) cmd/compile: cleanup code paths that aren't needed now that -p is required
- [issue/17583](https://go.dev/issue/17583) cmd/compile: insert write barriers in SSA
- [issue/19361](https://go.dev/issue/19361) cmd/compile: Devirtualize calls when concrete type behind interface is statically known
- [issue/33981](https://go.dev/issue/33981) cmd/compile: improve escape analysis fidelity
- [issue/15561](https://go.dev/issue/15561) cmd/compile: support inlining non-escaping closures
- [issue/30810](https://go.dev/issue/30810) cmd/compile: rewrite rulegen to not generate plaintext Go code directly
- [issue/71359](https://go.dev/issue/71359) cmd/compile: avoid allocs by better tracking of literals for interface conversions and make
- [issue/8618](https://go.dev/issue/8618) runtime: don't allocate for non-escaping conversions to interface{}
- [issue/61502](https://go.dev/issue/61502) cmd/compile: overhaul inliner
- [issue/78633](https://go.dev/issue/78633) cmd/compile: do bitwise constant folding
- [issue/5496](https://go.dev/issue/5496) cmd/compile: compile dense, pure-integer "switch" into jump table
- [issue/34381](https://go.dev/issue/34381) cmd/compile: use hashing for switch statements
- [issue/73137](https://go.dev/issue/73137) testing: implement b.Loop keepalive without preventing inlining, which can lead to heap allocs
- [issue/71268](https://go.dev/issue/71268) cmd/compile: opt: generate conditional comparisons for || and && conditions
- [issue/75203](https://go.dev/issue/75203) cmd/compile: intrinsify reflect.TypeFor
- [issue/80409](https://go.dev/issue/80409) cmd/compile/internal/ssa: split into subpackages
- [issue/10587](https://go.dev/issue/10587) cmd/gc: stack split check does not follow static branch prediction rules of modern CPUs
- [issue/15365](https://go.dev/issue/15365) cmd/compile: ARM SSA support
- [issue/16358](https://go.dev/issue/16358) cmd/compile: 386 SSA support
- [issue/16010](https://go.dev/issue/16010) cmd/compile: PowerPC (LE) SSA support
- [issue/16359](https://go.dev/issue/16359) cmd/compile: SSA backend for MIPS64
- [issue/18892](https://go.dev/issue/18892) all: WebAssembly ("wasm") support
- [issue/16357](https://go.dev/issue/16357) cmd/compile: delete the legacy compiler back end
- [issue/27532](https://go.dev/issue/27532) all: port to RISC-V
- [issue/44734](https://go.dev/issue/44734) cmd/asm: refactor the framework of the arm64 assembler
- [issue/44549](https://go.dev/issue/44549) cmd/asm,cmd/compile,cmd/link,cmd/objdump: ppc64 ISA 3.1 support
- [issue/20104](https://go.dev/issue/20104) cmd/compile: move arch-specific rewrite rules and ops into arch-specific packages
- [issue/81819](https://go.dev/issue/81819) simd: unify API generation around executable spec
- [design/bce](https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g) Bounds Checking Elimination.
- [discuss/checkptr](https://groups.google.com/g/golang-dev/c/SzwDoqoRVJA) Matthew Dempsky. Dynamic unsafe pointer checking available at master. golang-dev. October 17, 2019.
- [doc/pgo-preview](https://go.dev/blog/pgo-preview) Michael Pratt. Profile-guided optimization preview. February 8, 2023.
- [doc/pgo](https://go.dev/blog/pgo) Michael Pratt. Profile-guided optimization in Go 1.21. September 5, 2023.
- [doc/gotelemetry](https://go.dev/blog/gotelemetry) Robert Findley. Telemetry in Go 1.23 and beyond. September 3, 2024.
- [issue/67111](https://go.dev/issue/67111) cmd/go: add go telemetry subcommand
- [doc/allocation-optimizations](https://go.dev/blog/allocation-optimizations) Keith Randall. Allocating on the Stack. February 27, 2026.
- [doc/simd-experiment](https://go.dev/blog/simd-experiment) David Chase and Junyang Shao. Platform-independent SIMD in Go. September 24, 2026.
- [doc/archsimd](https://go.dev/blog/archsimd) Junyang Shao and David Chase. Arch-specific SIMD in Go. October 2, 2026.
- [talk/jin2022](https://www.youtube.com/watch?v=fYB5uhyDi7c) Jin Lin. Profile-Guided Optimization in the Go Compiler. Lightning talk, GopherCon 2022, Chicago. October 2022.
- [talk/jin2024](https://www.youtube.com/watch?v=58xVeeuXbKQ) Jin Lin. Automating Efficiency Improvement by Profile Guided Optimization. GopherCon 2024, Chicago. July 2024.

[Back To Top](#top)

### Linker

The Go Linker was written by Ken Thompson. Russ led a few more overhaul in Go 1.3. Austin led a rework to the linker together with Keith, Than,
Cheery and many other brilliant brains, which was landed
in Go 1.15 and Go 1.16.

- [design/go13linker](https://go.dev/s/go13linker) Russ Cox. Go 1.3 Linker Overhaul. Nov 2013.
- [design/go116linker](https://go.dev/s/better-linker) Austin Clements. Building a better Go linker. 2019-09-12.
- [issue/52209](https://go.dev/issue/52209) cmd/compile, cmd/link: generate dwarf type info in compiler
- [design/go-dynamic-tls](https://go.dev/design/71953-go-dynamic-tls) Alexander Musman. Proposal: Go general dynamic TLS. 2025-01-28.
  + [issue/71953](https://go.dev/issue/71953) proposal: runtime: support general dynamic thread local storage model
- [issue/42681](https://go.dev/issue/42681) build: move GOEXPERIMENT knob from make.bash to cmd/go
- [discuss/do-we-document-the](https://groups.google.com/g/golang-nuts/c/n5XSgDP5q2M) Do we document the history of .gosymtab and .gopclntab, and the differences with .symtab and .debug_line, .debug_frame? (golang-nuts)
- [discuss/presence-of-pcln-table](https://groups.google.com/g/golang-nuts/c/oxseUnhRtXM) Presence of PCLN Table in CGo binaries (golang-nuts)
- [issue/62577](https://go.dev/issue/62577) cmd/link, cmd/go: emit split DWARF on darwin
- [issue/4069](https://go.dev/issue/4069) cmd/ld: invoke host linker for Windows cgo build
- [issue/11799](https://go.dev/issue/11799) cmd/link: compress debug info
- [issue/38485](https://go.dev/issue/38485) runtime: tracking bug for ARM-based macOS and GOOS/GOARCH values
- [issue/36439](https://go.dev/issue/36439) all: port to Windows/arm64
- [issue/25893](https://go.dev/issue/25893) all: port to AIX/ppc64
- [issue/76038](https://go.dev/issue/76038) cmd/link: rearrange compiler/linker generated data for clarity
- [issue/67401](https://go.dev/issue/67401) cmd/link: lock down future uses of linkname
- [issue/26379](https://go.dev/issue/26379) cmd/compile: consider using DWARF 5
- [issue/41004](https://go.dev/issue/41004) cmd/link: add support for automatically generating .note.gnu.build-id
- [issue/68678](https://go.dev/issue/68678) cmd/link: LC_UUID not generated by go linker, resulting in failure to access local network on macOS 15
- [issue/69536](https://go.dev/issue/69536) crypto: obtain a FIPS 140-3 validation

[Back To Top](#top)

### Debugger

- [doc/go13heapdump](https://go.dev/s/go13heapdump) heapdump13
- [doc/go14heapdump](https://github.com/golang/go/wiki/heapdump14) heapdump14
- [doc/go15heapdump](https://github.com/golang/go/wiki/heapdump15-through-heapdump17) heapdump15 through heapdump17
- [design/heap-viewer](https://go.dev/design/16410-heap-viewer) Michael Matloob. Proposal: Go Heap Dump Viewer. 20 July 2016
  + [issue/16410](https://go.dev/issue/16410) x/tools/cmd/heapdump: create a heap dump viewer
- [design/profiler-labels](https://go.dev/design/17280-profile-labels) Michael Matloob. Proposal: Support for pprof profiler labels. May 15, 2017.
  + [issue/17280](https://go.dev/issue/17280) pprof: add support for profiler labels
- [issue/36821](https://go.dev/issue/36821) proposal: runtime/pprof: add PMU-based profiles
- [issue/53286](https://go.dev/issue/53286) proposal: runtime/pprof: extend profile configuration to support perf events
- [talk/derek2024](https://www.youtube.com/watch?v=ZnrXDnhSmH0) Derek Parker. A Decade of Debugging. GopherCon 2024, Chicago. July 2024.

[Back To Top](#top)

### Race Detector

- [issue/42598](https://go.dev/issue/42598) runtime: apparent false-positive race report for a buffered channel after CL 220419
- [issue/44853](https://go.dev/issue/44853) runtime: enable address sanitizer in Go
  + [discuss/enable-address-sanitizer-in](https://groups.google.com/g/golang-dev/c/ACKO07YeeW8) Enable address sanitizer in Go (golang-dev)
- [issue/64611](https://go.dev/issue/64611) all: add internal/asan and internal/msan packages

[Back To Top](#top)

### Tracer

- [design/go15trace](https://go.dev/s/go15trace) Dmitry Vyukov. Go Execution Tracer. Oct 2014
- [design/tracefmt](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@. Trace Event Format. October 2016.
- [issue/54466](https://go.dev/issue/54466) runtime: rewrite gentraceback as an iterator API
- [issue/57175](https://go.dev/issue/57175) runtime: diagnostics improvements tracking issue
- [design/execution-tracer-overhaul](https://go.dev/design/60773-execution-tracer-overhaul) Execution tracer overhaul.
  + [issue/60773](https://go.dev/issue/60773) runtime/trace: execution tracer overhaul
- [issue/63185](https://go.dev/issue/63185) runtime/trace: flight recording
- [issue/66703](https://go.dev/issue/66703) runtime/trace: set the v2 tracer as the default tracer
- [doc/execution-traces-2024](https://go.dev/blog/execution-traces-2024) Michael Knyszek. More powerful Go execution traces. March 14, 2024.
- [doc/flight-recorder](https://go.dev/blog/flight-recorder) Carlos Amedee and Michael Knyszek. Flight Recorder in Go 1.25. September 26, 2025.
- [talk/felix2025](https://www.youtube.com/watch?v=BayZ3k-QkFw) Felix Geisendörfer. Profiling Request Latency with Critical Path Analysis. GopherCon 2025, New York City. August 2025.

[Back To Top](#top)

### Lock Analysis

- [issue/38029](https://go.dev/issue/38029) x/build: linux-amd64-staticlockranking consistently failing
  - [cl/192704](https://go.dev/cl/192704) runtime: lock operation logging support
  - [cl/207348](https://go.dev/cl/207348) runtime: release timersLock while running timer
  - [cl/207846](https://go.dev/cl/207846) runtime: avoid potential deadlock when tracing memory code
  - [cl/207619](https://go.dev/cl/207619) runtime: static lock ranking for the runtime (enabled by GOEXPERIMENT)
  - [cl/222925](https://go.dev/cl/222925) cmd/go: define a build tag for any GOEXPERIMENT which is enabled
  - [cl/228417](https://go.dev/cl/228417) runtime: incorporate Gscan acquire/release into lock ranking order
  - [cl/229480](https://go.dev/cl/229480) runtime: added several new lock-rank partial order edges
  - [cl/231463](https://go.dev/cl/231463) runtime: add one extra lock ranking partial edge
  - [cl/233599](https://go.dev/cl/233599) runtime: add a lock partial order edge (assistQueue -> mspanSpecial)
  - [cl/236137](https://go.dev/cl/236137) runtime: add three new partial orders for lock ranking
- [issue/40677](https://go.dev/issue/40677) runtime: lock held checking

### Builder

- [design/go13nacl](https://go.dev/s/go13nacl) Russ Cox. Go 1.3 Native Client Support. October 2013.
- [design/gobuilder](http://go.dev/s/builderplan) Brad Fitzpatrick. Go Builder Plan. 2014-09-03.
  + [discuss/gobuilder](https://groups.google.com/g/golang-dev/c/sdFD0-2Ed8k) Changing how builders run
- [design/go14android](https://go.dev/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/go-generate](https://go.dev/s/go1.4-generate) Rob Pike. Go Generate. January 2014.
- [issue/13560](https://go.dev/issue/13560) proposal: build: define standard way to recognize machine-generated files
- [discuss/generatedcode](https://go.dev/s/generatedcode) Rob Pike's Final Comments on Issue 13560
- [design/goenv](https://go.dev/design/30411-env) Russ Cox. Proposal: go command configuration file. March 1, 2019
  + [issue/30411](https://go.dev/issue/30411) proposal: cmd/go: add go env -w to set default env vars
- [design/go116build](https://go.dev/s/go-build-design​) Russ Cox. Bug-resistant build constraints — Draft Design. June 30, 2020.
  + [issue/41184](https://go.dev/issue/41184) cmd/go: continue conversion to bug-resistant //go:build constraints
- Windows
  - [issue/41191](https://go.dev/issue/41191#issuecomment-690887303) toolchain directives
  + [discuss/win2000-golang-nuts](https://go.dev/s/win2000-golang-nuts) objections to removing Go support for Windows 2000 (x86-32)?
- [design/wasm](https://docs.google.com/document/d/131vjr4DH6JFnb-blm_uRdaC0_Nv3OUwjEY5qVCxCup4/edit#heading=h.mjo1bish3xni) Richard Musiol. WebAssembly architecture for Go. 28th February 2018.
- [design/wasm2](https://docs.google.com/document/d/1GRmy3rA4DiYtBlX-I1Jr_iHykbX8EixC3Mq0TCYqbKc/edit#heading=h.q4c21ihutzk0) WebAssembly assembly files
- [issue/57631](https://go.dev/issue/57631) build: use a zero for third digit for major release, such as 'go1.21.0'
- [issue/54265](https://go.dev/issue/54265) build: adopt Go 1.20 as bootstrap toolchain for Go 1.22
- [issue/69315](https://go.dev/issue/69315) build: adopt Go 1.24 as bootstrap toolchain for Go 1.26
- [doc/rebuild](https://go.dev/blog/rebuild) Russ Cox. Perfectly Reproducible, Verified Go Toolchains. August 28, 2023.
- [issue/57179](https://go.dev/issue/57179) cmd/go: load $GOROOT/go.env before loading user go.env
- [issue/59719](https://go.dev/issue/59719) cmd/go: add GOEXPERIMENT=cacheprog to let a child process implement the internal action/output cache
- [issue/64876](https://go.dev/issue/64876) cmd/go: enable GOCACHEPROG by default
- [issue/62067](https://go.dev/issue/62067) cmd/go: go build -json
- [issue/38248](https://go.dev/issue/38248) cmd/compile: replace CallImport with go:wasmimport directive
- [issue/58141](https://go.dev/issue/58141) all: add GOOS=wasip1 GOARCH=wasm port
- [doc/wasi](https://go.dev/blog/wasi) Johan Brandhorst-Satzkorn, Julien Fabre, Damian Gryski, Evan Phoenix, and Achille Roussel. WASI support in Go. September 13, 2023.
- [issue/65199](https://go.dev/issue/65199) cmd/compile: add go:wasmexport directive
- [design/wasmexport-notes](https://docs.google.com/document/d/1R14wRHv_UjeFnemvV5hbp75cCRNzVH9ZZS-Yo_zITxg) Achille Roussel. go:wasmexport - implementation notes.
- [doc/wasmexport](https://go.dev/blog/wasmexport) Cherry Mui. Extensible Wasm Applications with Go. February 13, 2025.
- [talk/johan2023](https://www.youtube.com/watch?v=BAFKruuvUN4) Johan Brandhorst-Satzkorn. Go Support for WASI/Wasm. Lightning talk, GopherCon 2023, San Diego. September 2023.

[Back To Top](#top)

### Modules

- [design/go-dep](https://docs.google.com/document/d/18tNd8r5DV0yluCR7tPvkMTsWD_lYcRO7NhpNSDymRr8) Go Packaging Proposal Process
- [design/go-dep2](https://docs.google.com/document/d/1qnmjwfMmvSCDaY4jxPmLAccaaUI5FfySNE90gB0pTKQ/edit) Dependency Management Tool
  - [doc/go-dep](https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) Sam Boyer. The Saga of Go Dependency Management. Dec 13, 201
  - [talk/go-dep](https://www.youtube.com/watch?v=5LtMb090AZI) GopherCon 2017: Sam Boyer - The New Era of Go Package Management. Jul 24, 2017
  - [talk/go-dep-design](https://www.youtube.com/watch?v=wBTGd_dvnO8) dotGo 2017 - sam boyer - The Functional Design of Dep. Dec 8, 2017
  - [discuss/go-dep](https://www.youtube.com/watch?v=sbrZfPgNmfw) Building Predictability into Your Pipeline :: Russ Cox; Jess Frazelle, Sam Boyer, Pete Garcin. Feb 22, 2018
- [design/vgo](https://go.dev/design/24301-versioned-go) Russ Cox. Proposal: Versioned Go Modules. March 20, 2018.
  + [issue/24301](https://go.dev/issue/24301) cmd/go: add package version support to Go toolchain
  + [doc/deps](https://research.swtch.com/deps) Russ Cox. Our Software Dependency Problem. January 23, 2019.
  + [doc/vgo](https://research.swtch.com/vgo) Russ Cox. Go & Versioning
  + [discuss/groups-gomod](https://groups.google.com/g/golang-dev/c/a5PqQuBljF4) go modules have landed
  + [discuss/go-dep-response](https://www.reddit.com/r/golang/comments/92f3q1/peter_bourgon_a_response_about_dep_and_vgo/) Reddit discussion.
  + [doc/go-dep-response](https://peter.bourgon.org/blog/2018/07/27/a-response-about-dep-and-vgo.html) Peter Bourgon. A response about dep and vgo. 2018-07-27.
  + [discuss/go-dep-response2](https://news.ycombinator.com/item?id=17628311) Hacker News discussion.
  + [discuss/go-dep-twitter](https://twitter.com/_rsc/status/1022588240501661696) Russ Cox's Twitter Storm
- [design/sumdb](https://go.dev/design/25530-sumdb) Russ Cox, Filippo Valsorda. Proposal: Secure the Public Go Module Ecosystem. April 24, 2019.
  + [issue/25530](https://go.dev/issue/25530) proposal: cmd/go: secure releases with transparency log
- [issue/23966](https://go.dev/issue/23966#issuecomment-377997161) why go.mod has its own bespoke syntax?
- [design/lazy-gomod](https://go.dev/design/36460-lazy-module-loading) Bryan C. Mills. Proposal: Lazy Module Loading. Feb 20, 2020
- [issue/45713](https://go.dev/issue/45713) proposal: cmd/go: add a workspace mode
  + [design/workspace](https://go.dev/design/45713-workspace) Proposal: Multi-Module Workspaces in cmd/go
- [doc/vuln](https://go.dev/blog/vuln) Julie Qiu. Vulnerability Management for Go. September 6, 2022.
- [doc/govulncheck](https://go.dev/blog/govulncheck) Julie Qiu. Govulncheck v1.0.0 is released! July 13, 2023.
- [discuss/60696](https://github.com/golang/go/discussions/60696) Michael Matloob. cmd/go: add package forwarding. June 8, 2023.
- [discuss/61669](https://github.com/golang/go/discussions/61669) Russ Cox. gonew: templates for new modules. July 31, 2023.
- [doc/gonew](https://go.dev/blog/gonew) Cameron Balahan. Experimenting with project templates. July 31, 2023.
- [issue/48429](https://go.dev/issue/48429) cmd/go: track tool dependencies in go.mod
- [design/x-repo-continuous-go](https://go.dev/design/69095-x-repo-continuous-go) Dmitri Shuralyov. Proposal: all, x/build/cmd/relui: automate go directive maintenance in golang.org/x repositories. August 27, 2024.
- [doc/pkgsite-api](https://go.dev/blog/pkgsite-api) Ethan Lee, Hana Kim, and Jonathan Amsterdam. Introducing the pkg.go.dev API. May 21, 2026.

[Back To Top](#top)

### gopls

- [design/gopls-workspace](https://go.dev/design/37720-gopls-workspaces) Heschi Kreinick, Rebecca Stambler. Proposal: Multi-project gopls workspaces. Mar 9, 2020.
- [doc/gopls-scalability](https://go.dev/blog/gopls-scalability) Robert Findley and Alan Donovan. Scaling gopls for the growing Go ecosystem. September 8, 2023.
- [issue/63501](https://go.dev/issue/63501) x/tools/cmd/deadcode: add new command
- [doc/deadcode](https://go.dev/blog/deadcode) Alan Donovan. Finding unreachable functions with deadcode. December 12, 2023.
- [issue/61324](https://go.dev/issue/61324) x/tools/go/analysis/checker: an importable analysis driver based on go/packages
- [issue/32816](https://go.dev/issue/32816) cmd/fix: automate migrations for simple deprecations
- [issue/73605](https://go.dev/issue/73605) cmd/fix: remove all functionality
- [issue/71859](https://go.dev/issue/71859) cmd/go: fix: apply fixes from modernizers, inline, and other analyzers
- [doc/gofix](https://go.dev/blog/gofix) Alan Donovan. Using go fix to modernize Go code. February 17, 2026.
- [doc/inliner](https://go.dev/blog/inliner) Alan Donovan. //go:fix inline and the source-level inliner. March 10, 2026.
- [design/lsp-interactive-refactoring](https://go.dev/design/76331-lsp-interactive-refactoring) Hongxiang Jiang. Proposal: Language Server Protocol (LSP) Interactive Code Actions. July 16, 2026.
- [doc/interactive-refactoring](https://go.dev/blog/interactive-refactoring) Hongxiang Jiang. Interactive refactoring in gopls. October 16, 2026.

[Back To Top](#top)

### Testing, x/perf

- [design/subtests](https://go.dev/design/12166-subtests) Marcel van Lohuizen. testing: programmatic sub-test and sub-benchmark support. September 2, 2015.
  + [issue/12166](https://go.dev/issue/12166) proposal: testing: programmatic sub-test and sub-benchmark support
- [design/gotest-bench](https://go.dev/design/14313-benchmark-format) Russ Cox, Austin Clements. Proposal: Go Benchmark Data Format. February 2016.
  + [issue/14313](https://go.dev/issue/14313) cmd/go: decide, document standard benchmark data format
- [issue/20875](https://go.dev/issue/20875) testing: consider calling ReadMemStats less during benchmarking
- [issue/27217](https://go.dev/issue/27217) testing: tiny benchmark with StopTimer runs forever
- [issue/41637](https://go.dev/issue/41637) testing: benchmark iteration reports incorrectly
- [issue/41641](https://go.dev/issue/41641) testing: inconsistent benchmark measurements when interrupts timer
- [design/gotest-json](https://go.dev/design/2981-go-test-json) Nodir Turakulov. Proposal: -json flag in go test. 2016-09-14.
- [design/testing-helper](https://go.dev/design/4899-testing-helper) Caleb Spare. Proposal: testing: better support test helper functions with TB.Helper. 2016-12-27
  + [issue/4899](https://go.dev/issue/4899) testing: add t.Helper to make file:line results more useful
- [design/fuzzing](https://go.dev/s/draft-fuzzing-design) Katie Hockman. Design Draft: First Class Fuzzing
- [issue/43744](https://go.dev/issue/43744) testing: benchmark unit properties
- [issue/48803](https://go.dev/issue/48803) all: Go compiler/runtime performance monitoring system
- [issue/49121](https://go.dev/issue/49121) x/perf/storage: support postgres for db
- [issue/61179](https://go.dev/issue/61179) proposal: testing: add identity function that forces evaluation for benchmarks
- [issue/67434](https://go.dev/issue/67434) proposal: testing/synctest: new package for testing concurrent code

<!-- - Tool chain, benchseries/benchstat -->
- [design/revamp-code-coverage](https://go.dev/design/51430-revamp-code-coverage) Than McIntosh. Proposal: extend code coverage testing to include applications. 2022-03-02.
  + [issue/51430](https://go.dev/issue/51430) cmd/cover: extend coverage testing to include applications
- [doc/integration-test-coverage](https://go.dev/blog/integration-test-coverage) Than McIntosh. Code coverage for Go integration tests. March 8, 2023.
- [doc/synctest](https://go.dev/blog/synctest) Damien Neil. Testing concurrent code with testing/synctest. February 19, 2025.
- [doc/testing-b-loop](https://go.dev/blog/testing-b-loop) Junyang Shao. More predictable benchmarking with testing.B.Loop. April 2, 2025.
- [issue/73567](https://go.dev/issue/73567) testing/synctest: replace Run with Test
- [doc/testing-time](https://go.dev/blog/testing-time) Damien Neil. Testing Time (and other asynchronicities). August 26, 2025.
- [issue/71287](https://go.dev/issue/71287) testing: store test artifacts

[Back To Top](#top)

## Runtime Core

### Scheduler

- [paper/work-steal](https://dl.acm.org/citation.cfm?id=324234) Robert D. Blumofe and Charles E. Leiserson. 1999. Scheduling multithreaded computations by work stealing. J. ACM 46, 5 (September 1999), 720-748.
- [cl/sched-m-1](https://github.com/golang/go/commit/96824000ed89d13665f6f24ddc10b3bf812e7f47#diff-1fe527a413d9f1c2e5e22e08e605a192) Russ Cox, Clean up scheduler. Aug 5, 2008.
- [cl/sched-m-n](https://github.com/golang/go/commit/fe1e49241c04c748d0e3f4762925241adcb8d7da) things are much better now, Nov 11, 2009.
- [design/go11sched](https://go.dev/s/go11sched) Dmitry Vyukov. Scalable Go Scheduler Design Doc, 2012
  + [cl/7314062](https://github.com/golang/go/commit/779c45a50700bda0f6ec98429720802e6c1624e8) runtime: improved scheduler
- [design/sched-preempt-dmitry](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#heading=h.3pilqarbrc9h) Dmitry Vyukov. Go Preemptive Scheduler Design Doc, 2013
- [design/sched-numa](https://docs.google.com/document/u/0/d/1d3iI2QWURgDIsSR6G2275vMeQ_X7w-qxM2Vp7iGwwuM/pub) Dmitry Vyukov, NUMA-aware scheduler for Go. Sep 2014.
- [design/go15gomaxprocs](https://go.dev/s/go15gomaxprocs) Russ Cox. Go 1.5 GOMAXPROCS Default. May 2015.
- [issue/73193](https://go.dev/issue/73193) runtime: CPU limit-aware GOMAXPROCS default (container/cgroup-aware, Go 1.25)
  + [issue/33803](https://go.dev/issue/33803) runtime: make GOMAXPROCS cfs-aware on GOOS=linux
- [doc/go17sched](https://www.quora.com/How-does-the-golang-scheduler-work/answer/Ian-Lance-Taylor) Ian Lance Taylor. How does the golang scheduler work? July 16, 2016.
- [design/sched-preempt-austin](https://go.dev/design/24543-non-cooperative-preemption) Austin Clements. Proposal: Non-cooperative goroutine preemption. Jan 2019.
  + [cl/43050](https://go.dev/cl/43050) cmd/compile: loop preemption with fault branch on amd64. May 09, 2019.
  + [issue/10958](https://go.dev/issue/10958) runtime: tight loops should be preemptible
  + [issue/24543](https://go.dev/issue/24543) runtime: non-cooperative goroutine preemption
  + [issue/36365](https://go.dev/issue/36365) runtime: clean up async preemption loose ends
- [issue/14592](https://go.dev/issue/14592) runtime: let idle OS threads exit
- [issue/18237](https://go.dev/issue/18237) runtime: scheduler is slow when goroutines are frequently woken
- [issue/20395](https://go.dev/issue/20395) runtime: terminate locked OS thread if its goroutine exits
- [issue/20458](https://go.dev/issue/20458) proposal: runtime: pair LockOSThread, UnlockOSThread calls
- [issue/21827](https://go.dev/issue/21827) runtime: big performance penalty with runtime.LockOSThread
- [issue/27345](https://go.dev/issue/27345) runtime: use parent goroutine's stack for new goroutines
- [issue/28808](https://go.dev/issue/28808) runtime: scheduler work stealing slow for high GOMAXPROCS
- [issue/32113](https://go.dev/issue/32113) runtime: optimization to reduce P churn.
- [issue/43997](https://go.dev/issue/43997) runtime: non-spinning Ms spin uselessly when work exists
- [issue/44313](https://go.dev/issue/44313) runtime: stopped Ms can't become dedicated or fractional GC workers
- [issue/51071](https://go.dev/issue/51071) runtime: investigate possible Go scheduler improvements inspired by Linux Kernel's CFS
- [discuss/periodic-task-when-time](https://groups.google.com/g/golang-nuts/c/_XEuq69kUOY) Periodic task when time.Ticker and time.Sleep are pretty expensive (golang-nuts)
- [discuss/why-there-is-a](https://groups.google.com/g/golang-nuts/c/WqVEThZnQkM) Why there is a so great performance penalty when the usage combines LockOSThread and channel? (golang-nuts)
- [design/mutex-spinbit](https://go.dev/design/68578-mutex-spinbit) Rhys Hiltner. Proposal: Improve scalability of runtime.lock2. 2024-10-16.
- [issue/61015](https://go.dev/issue/61015) runtime: change mutex profile to scale contention by number of blocked goroutines
- [discuss/runtime-setting-gomaxprocs-in](https://groups.google.com/g/golang-nuts/c/he617hSABGo) runtime: Setting GOMAXPROCS in over-scheduled Contanier Environments (golang-nuts)
- [discuss/concurrency-and-task-scheduler](https://groups.google.com/g/golang-nuts/c/Q4fgggATcTo) Concurrency and task scheduler scalability (golang-nuts)
- [issue/69687](https://go.dev/issue/69687) testing/synctest: experimental package for testing concurrent code
- [issue/45435](https://go.dev/issue/45435) sync: add Mutex.TryLock
  + [issue/63796](https://go.dev/issue/63796) sync: add WaitGroup.Go
  + [issue/61696](https://go.dev/issue/61696) sync: add Map.Clear method
- [design/goroutine-leak-detection-gc](https://go.dev/design/74609-goroutine-leak-detection-gc) Georgian-Vlad Saioc, Milind Chabbi. Proposal: Goroutine leak detection via garbage collection. 14 Aug 2025.
  + [issue/74609](https://go.dev/issue/74609) runtime/pprof,runtime: new goroutine leak profile
  + [issue/38651](https://go.dev/issue/38651) runtime: add 'created by goroutine number' to stack traces
- [issue/543](https://go.dev/issue/543) runtime: preemptive scheduling
- [issue/68578](https://go.dev/issue/68578) runtime: improve scaling of lock2
- [issue/66999](https://go.dev/issue/66999) runtime: correct attribution of internal mutex contention
- [issue/57071](https://go.dev/issue/57071) runtime/metrics: add a way to measure runtime-internal lock contention
- [issue/49881](https://go.dev/issue/49881) runtime/metrics: add contention metrics
- [issue/81657](https://go.dev/issue/81657) runtime: speed up GC phase transitions
- [issue/65694](https://go.dev/issue/65694) runtime: stronger affinity between G ↔ P ↔ M ↔ CPU?
- [issue/53789](https://go.dev/issue/53789) runtime: express lock rank graph as a DAG like go/build/deps_test.go
- [issue/13759](https://go.dev/issue/13759) proposal: Add a Partial Deadlock Detector For Go
- [issue/21678](https://go.dev/issue/21678) runtime: provide a mechanism for debuggers to execute a function in the program being debugged
- [issue/70089](https://go.dev/issue/70089) proposal: runtime/mainthread: add mainthread.Do for mediating access to the main thread
- [doc/container-aware-gomaxprocs](https://go.dev/blog/container-aware-gomaxprocs) Michael Pratt and Carlos Amedee. Container-aware GOMAXPROCS. August 20, 2025.

[Back To Top](#top)

### Execution Stack

- [design/contigstack](https://go.dev/s/contigstacks) Contiguous stacks
- [doc/stackroots](https://docs.google.com/document/d/13v_u3UrN2pgUtPnH4y-qfmlXwEEryikFu0SQiwk35SA/pub) precise stack roots
  + [discussion/stackroots](https://groups.google.com/g/golang-dev/c/5cw0mjxRB_o/m/h0L1GmnY_HAJ) Precise garbage collection of stack roots
- [issue/17007](https://go.dev/issue/17007) runtime: fatal error: bad g->status in ready
- [issue/18138](https://go.dev/issue/18138) runtime: new goroutines can spend excessive time in morestack
  + [design/predict-stack-size](https://docs.google.com/document/d/1YDlGIdVTPnmUiTAavlZxBI1d9pwGQgZT7IKFKlIXohQ/edit#) Keith Randall. Determining a good starting stack size. 2021/08/18.
  + [cl/341990](https://go.dev/cl/341990) runtime: predict stack sizing
  + [cl/345889](https://go.dev/cl/345889) runtime: measure stack usage; start stacks larger if needed
- [issue/26061](https://go.dev/issue/26061) runtime: g0 stack.lo is sometimes too low
- [discuss/proposal-shared-libraries-in](https://groups.google.com/g/golang-nuts/c/zmjXkGrEx6Q) Proposal: Shared libraries in Go, take 2 (+linux/amd64 implementation) (golang-nuts)
- [discuss/re-patch-rfc-go](https://groups.google.com/g/golang-nuts/c/P05BDjLcQ5k) Re: [PATCH] RFC: go shared libraries support (Linux/x64) (golang-nuts)
- [issue/15403](https://go.dev/issue/15403) internal/cpu: new package to expose processor capabilities
  + [issue/27218](https://go.dev/issue/27218) runtime: document (and support) cpu options in GODEBUG
- [issue/30439](https://go.dev/issue/30439) all: remove nacl port
- [issue/21865](https://go.dev/issue/21865) runtime/secret: add new package
  + [issue/66450](https://go.dev/issue/66450) crypto/subtle: add WithDataIndependentTiming
- [issue/45453](https://go.dev/issue/45453) all: add GOAMD64 environment variable
  + [issue/60905](https://go.dev/issue/60905) all: add GOARM64=v8.1 and so on
- [issue/36435](https://go.dev/issue/36435) all: stop using non-libc system calls on OpenBSD
- [issue/17490](https://go.dev/issue/17490) cmd/link: use libsystem_kernel.dylib or libSystem.dylib for syscalls on macOS
- [issue/58378](https://go.dev/issue/58378) runtime: support unwinding logical cgo stack transitions using frame pointers
- [issue/16638](https://go.dev/issue/16638) runtime: use frame pointers for callers
- [issue/69577](https://go.dev/issue/69577) crypto/rand: use new Linux vDSO

[Back To Top](#top)

### Memory Management

#### Allocator

A quick history about the Go's memory allocator: Russ Cox first implements
the memory allocator based on `tcmalloc` for Go 1, `mcache` is cached on M.
Then he revised the allocator to allow user code to use 16GB memory and later allows 128GB.
However, the allocator (including scavenger) was suffered from massive lock contentions and
does not scale. After Dmitry's scalable runtime scheduler, the allocator can allocate directly
from P (with much less) lock contentions. In the meantime, the scavenger is migrated from
an independent thread into the system monitor thread. Now, Michael is actively working on
improving the memory allocator's scalability, such as migrating scavenger
to user threads, bitmap-based page allocator, scalable mcentral.

- [doc/tcmalloc](http://goog-perftools.sourceforge.net/doc/tcmalloc.html) Sanjay Ghemawat, Paul Menage. TCMalloc : Thread-Caching Malloc. Google Inc., 2009
- [design/go113scavenge](https://go.googlesource.com/proposal/+/aa701aae530695d32916b779e048a3e18311a2e3/design/30333-smarter-scavenging.md) Michael Knyszek. Proposal: Smarter Scavenging. 2019-05-09.
  + [issue/30333](https://go.dev/issue/30333) runtime: smarter scavenging
  + [issue/32012](https://go.dev/issue/32012) runtime: background scavenger is overzealous with small heaps
  + [issue/31966](https://go.dev/issue/31966) runtime: background scavenger can delay deadlock detection significantly
  + [issue/34047](https://go.dev/issue/34047) runtime: potential deadlock cycle caused by scavenge.lock
  + [issue/34048](https://go.dev/issue/34048) runtime: scavenger pacing fails to account for fragmentation
  + [issue/35788](https://go.dev/issue/35788) runtime: scavenger not as effective as in previous releases
  + [issue/36521](https://go.dev/issue/36521) runtime: performance degradation in go 1.12
  + [issue/36603](https://go.dev/issue/36603) runtime: sysUsed often called on non-scavenged memory
- [design/go114pagealloc](https://go.googlesource.com/proposal/+/a078ea9d72b99dc88fdfd2cb6ee150a8ce202ea2/design/35112-scaling-the-page-allocator.md) Michael Knyszek, Austin Clements. Proposal: Scaling the Go page allocator. 2019-10-18.
  + [issue/35112](https://go.dev/issue/35112) runtime: make the page allocator scale
  + [cl/200439](https://go.dev/cl/200439) runtime: place lower limit on trigger ratio
- [issue/8885](https://go.dev/issue/8885) runtime: consider adding 24-byte size class.
- [issue/37487](https://go.dev/issue/37487) runtime: improve mcentral scalability
  + [cl/221182](https://go.dev/cl/221182) runtime: add new mcentral implementation
- [issue/18155](https://go.dev/issue/18155) runtime: latency in sweep assists when there's no garbage
- [issue/19112](https://go.dev/issue/19112) runtime: deadlock involving gcControllerState.enlistWorker
- [issue/23687](https://go.dev/issue/23687) runtime: use MADV_FREE on linux as well
  + [cl/135395](https://go.dev/cl/135395) runtime: use MADV_FREE on Linux if available
- [issue/29707](https://go.dev/issue/29707) cmd/trace: failed to parse trace: no consistent ordering of events possible
- [issue/35954](https://go.dev/issue/35954) runtime: handle hitting the top of the address space in the allocator more gracefully
- [issue/37927](https://go.dev/issue/37927) runtime: GC pacing exhibits strange behavior with a low GOGC
- [issue/38130](https://go.dev/issue/38130) runtime: incorrect sanity checks in the page allocator
- [issue/38404](https://go.dev/issue/38404) runtime: STW GC stops working on arm64/mips64le
- [issue/38605](https://go.dev/issue/38605) runtime: pageAlloc.allocToCache updates pageAlloc.searchAddr in an invalid way
- [issue/38617](https://go.dev/issue/38617) runtime: scavenger freezes up in Go 1.14 in Windows due to coarse time granularity
- [issue/38966](https://go.dev/issue/38966) runtime: aix-ppc64 builder is failing with "bad prune", "addr range base and limit are not in the same memory segment" fatal errors
- [issue/39128](https://go.dev/issue/39128) runtime: linux-mips-rtrk builder consistently failing with "bad prune" as of CL 233497
- [issue/40191](https://go.dev/issue/40191) runtime: pageAlloc.searchAddr may point to unmapped memory in discontiguous heaps, violating its invariant
- [issue/40457](https://go.dev/issue/40457) runtime: runqputbatch does not protect its call to globrunqputbatch
- [issue/40641](https://go.dev/issue/40641) runtime: race between stack shrinking and channel send/recv leads to bad sudog values
- [issue/42330](https://go.dev/issue/42330) runtime: default to MADV_DONTNEED on Linux
  + [cl/267100](https://go.dev/cl/267100) runtime: default to MADV_DONTNEED on Linux
- [issue/46787](https://go.dev/issue/46787) runtime: provide Pinner API for object pinning
- [issue/51317](https://go.dev/issue/51317) proposal: arena: new package providing memory arenas
- [issue/59960](https://go.dev/issue/59960) runtime: improve heap hugepage utilization
- [issue/62483](https://go.dev/issue/62483) unique: new package with unique.Handle
  - [repo/unsafe-assuming-no-moving-gc](https://github.com/go4org/unsafe-assume-no-moving-gc) go4.org/unsafe/assume-no-moving-gc
  - [cl/498121](https://go.dev/cl/498121) runtime: add heapObjectsCanMove
  - [cl/498261](https://go.dev/cl/498261) runtime: change heapObjectsCanMove to a func
- [design/scaling-the-page-allocator](https://go.dev/design/35112-scaling-the-page-allocator) Michael Knyszek, Austin Clements. Proposal: Scaling the Go page allocator. 2019-10-18.
- [design/smarter-scavenging](https://go.dev/design/30333-smarter-scavenging) Michael Knyszek. Proposal: Smarter Scavenging.
- [design/heap-hugepage-util](https://go.dev/design/59960-heap-hugepage-util) Michael Knyszek, Michael Pratt. A more hugepage-aware Go heap.
- [design/memory-regions](https://go.dev/design/70257-memory-regions) Memory regions.
  + [discuss/70257](https://github.com/golang/go/discussions/70257) memory regions
- [design/runtime-freegc](https://go.dev/design/74299-runtime-freegc) Directly freeing user memory to reduce GC work.
  + [discuss/a-safe-free-of](https://groups.google.com/g/golang-nuts/c/cmpiArv10f4) a safe free of memory proposal, runtime.FreeMemory() (golang-nuts)
- [issue/10460](https://go.dev/issue/10460) runtime: 512GB memory limitation
- [issue/79286](https://go.dev/issue/79286) runtime: sizespecializedmalloc experiment
- [issue/40431](https://go.dev/issue/40431) proposal: runtime.Pin() to protect against hypothetical moving GC in the future
- [issue/69327](https://go.dev/issue/69327) runtime: non-preemptible zeroing in clear, append, bytes.growSlice, etc
- [design/go-pmem](https://go.dev/design/43810-go-pmem) Jerrin Shaji George, Mohit Verma, Rajesh Venkatasubramanian, Pratap Subrahmanyam. Proposal: Add support for Persistent Memory in Go. January 20, 2021.
  + [issue/43810](https://go.dev/issue/43810) proposal: spec: add native support for programming persistent memory in Go (declined)
- [doc/unique](https://go.dev/blog/unique) Michael Knyszek. New unique package. August 27, 2024.
- [doc/size-specialized-allocations](https://go.dev/blog/size-specialized-allocations) Michael Matloob. Size-Specialized Memory Allocation. September 16, 2026.


[Back To Top](#top)

#### Collector

- [paper/on-the-fly-gc](https://doi.org/10.1145/359642.359655) Edsger W. Dijkstra, Leslie Lamport, A. J. Martin, C. S. Scholten, and E. F. M. Steffens. 1978. On-the-fly garbage collection: An exercise in cooperation. Commun. ACM 21, 11 (November 1978), 966–975.
- [paper/yuasa-barrier](https://doi.org/10.1016/0164-1212(90)90084-Y) T. Yuasa. 1990. Real-time garbage collection on general-purpose machines. J. Syst. Softw. 11, 3 (March 1990), 181-198.
- [design/bettergc](https://docs.google.com/document/d/1HCPu3WKyCX3ZRYxmIMKTk0Ik1dePxKW1p02k3uhcft4/edit) Dmitry Vyukov. Better GC and Memory Allocator for Go. May 20, 2013
  + [discuss/bettergc](https://groups.google.com/g/golang-dev/c/pwUh0BVFpY0) Better GC and Memory Allocator for Go
- [design/go13gc](https://docs.google.com/document/d/1v4Oqa0WwHunqlb8C3ObL_uNQw3DfSY-ztoA-4wWbKcg/pub) Dmitry Vyukov. Simpler and faster GC for Go. July 16, 2014
  + [cl/106260045](https://codereview.appspot.com/106260045) runtime: simpler and faster GC
- [design/go14gc](https://go.dev/s/go14gc) Richard L. Hudson. Go 1.4+ Garbage Collection (GC) Plan and Roadmap. August 6, 2014.
- [design/go15gcpacing](https://go.dev/s/go15gcpacing) Austin Clements. Go 1.5 concurrent garbage collector pacing. Mar 10, 2015.
- [discuss/gcpacing](https://groups.google.com/forum/#!topic/golang-dev/YjoG9yJktg4) Austin Clements et al. Discussion of "Proposal: Garbage collector pacing". March 10, 2015.
- [issue/11970](https://go.dev/issue/11970) runtime: replace GC coordinator with state machine
- [design/sweep-free-alloc](https://go.dev/design/12800-sweep-free-alloc) Austin Clements. Proposal: Dense mark bits and sweep-free allocation. Sep 30, 2015.
- [issue/12800](https://go.dev/issue/12800) runtime: replace free list with direct bitmap allocation
- [design/decentralized-gc](https://go.dev/design/11970-decentralized-gc) Austin Clements. Proposal: Decentralized GC coordination. October 25, 2015.
- [issue/12967](https://go.dev/issue/12967#issuecomment-171466238) runtime: shrinkstack during mark termination significantly increases GC STW time
- [issue/14951](https://go.dev/issue/14951) runtime: mutator assists are over-aggressive, especially at high GOGC
- [design/eliminate-rescan](https://go.dev/design/17503-eliminate-rescan) Austin Clements, Rick Hudson. Eliminate STW stack re-scanning. October 21, 2016.
  + [issue/17503](https://go.dev/issue/17503) runtime: eliminate stack rescanning
- [design/concurrent-rescan](https://go.dev/design/17505-concurrent-rescan) Austin Clements, Rick Hudson. Proposal: Concurrent stack re-scanning. Oct 18, 2016.
  + [issue/17505](https://go.dev/issue/17505) runtime: perform concurrent stack re-scanning
- [design/soft-heap-limit](https://go.dev/design/14951-soft-heap-limit) Austin Clements. Proposal: Separate soft and hard heap size goal. October 21, 2017.
- [issue/22460](https://go.dev/issue/22460) runtime: optimize write barrier
- [design/roc](https://go.dev/s/gctoc) Request Oriented Collector (ROC) Algorithm
  + [cl/roc](https://go.dev/cl/25058) runtime: ROC write barrier code
  + [cl/generational-gc](https://go.dev/cl/137482) runtime: trigger generational GC
- [doc/ismm-gc](https://go.dev/blog/ismmkeynote) Rick Hudson. Getting to Go: The Journey of Go's Garbage Collector. 12 July 2018.
- [discuss/ismm-gc](https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns) Garbage Collection Slides and Transcript now available
- [design/simplify-mark-termination](https://go.dev/design/26903-simplify-mark-termination) Austin Clements. Proposal: Simplify mark termination and eliminate mark 2. Aug 9, 2018.
  + [issue/26903](https://go.dev/issue/26903) runtime: simplify mark termination and eliminate mark 2
- [design/gcscan](https://docs.google.com/document/d/1un-Jn47yByHL7I0aVIP_uVCMxjdM5mpelJhiKlIqxkE/edit#) Proposal: GC scanning of stacks
  + [issue/22350](https://go.dev/issue/22350) cmd/compile: compiler can unexpectedly preserve memory,
- [issue/27993](https://go.dev/issue/27993) runtime: error message: P has cached GC work at end of mark termination
- [issue/37116](https://go.dev/issue/37116) runtime: 10ms-26ms latency from GC in go1.14rc1, possibly due to 'GC (idle)' work
- [issue/42430](https://go.dev/issue/42430) runtime: GC pacer problems meta-issue
  + [issue/39983](https://go.dev/issue/39983) runtime: idle GC interferes with auto-scaling
  + [issue/17969](https://go.dev/issue/17969) runtime: aggressive GC completion is disruptive to co-tenants
  + [issue/14812](https://go.dev/issue/14812) runtime: GC causes latency spikes
  + [issue/40460](https://go.dev/issue/40460) runtime: goroutines may allocate beyond the hard heap goal
  + [issue/29696](https://go.dev/issue/29696) proposal: runtime: add way for applications to respond to GC backpressure
  + [issue/23044](https://go.dev/issue/23044) proposal: runtime: add a mechanism for specifying a minimum target heap size
- [issue/42642](https://go.dev/issue/42642) runtime: multi-ms sweep termination pauses (second edition)
- [issue/44163](https://go.dev/issue/44163) runtime: remove idle GC workers
- [issue/44167](https://go.dev/issue/44167) runtime: GC pacer redesign
  + [design/gc-pacer-redesign](https://go.dev/design/44167-gc-pacer-redesign) GC Pacer Redesign
- [issue/44309](https://go.dev/issue/44309) proposal: runtime/debug: user-configurable memory target
- [issue/48409](https://go.dev/issue/48409) proposal: runtime/debug: soft memory limit
  + [design/user-configurable-memory-target](https://go.dev/design/44309-user-configurable-memory-target) User-configurable memory target
  + [design/soft-memory-limit](https://go.dev/design/48409-soft-memory-limit) Soft memory limit
- [issue/45894](https://go.dev/issue/45894) runtime: mark termination is slow to restart mutator
- [issue/45315](https://go.dev/issue/45315) runtime: runtime.GC can return without finishing sweep
- [issue/49075](https://go.dev/issue/49075) runtime: possible memory corruption
- [issue/52433](https://go.dev/issue/52433) runtime: heap goal overrun due to scheduler delays in mark termination
- [doc/gc-guide](https://go.dev/doc/gc-guide) A Guide to the Go Garbage Collector
- [issue/12234](https://go.dev/issue/12234) runtime: revisit non-constant assist ratio
- [issue/27732](https://go.dev/issue/27732) runtime: mark assist blocks GC microbenchmark for 7ms
- [design/batch-wb](https://docs.google.com/document/d/1uXH_HKo2QlU2N2nE-0tm1ikibDT1stR56_-Ki7N-LUg/edit#heading=h.n0os25ndb49z) Keith Randall. Batching Write Barriers. 2022 Oct, 25.
- [issue/31222](https://go.dev/issue/31222) runtime: long pauses STW (sweep termination) on massive block allocation
- [doc/greenteagc](https://go.dev/blog/greenteagc) Michael Knyszek and Austin Clements. The Green Tea Garbage Collector. October 29, 2025. (experimental in Go 1.25, default in Go 1.26)
  + [issue/73581](https://go.dev/issue/73581) runtime: green tea garbage collector
- [discuss/sub-millisecond-gc-pauses](https://groups.google.com/g/golang-dev/c/Ab1sFeoZg_8/m/_DaL0E8fAwAJ) Sub-millisecond GC pauses (golang-dev)
- [issue/67535](https://go.dev/issue/67535) runtime: add AddCleanup
  + [issue/71825](https://go.dev/issue/71825) runtime: allow cleanups to run concurrently
- [issue/67552](https://go.dev/issue/67552) weak: new package providing weak pointers
- [discuss/secure-go-binaries](https://groups.google.com/g/golang-nuts/c/Jd9tlNc6jUE/m/qp2oyfEEfjQJ) Secure Go binaries (golang-nuts)
- [issue/11911](https://go.dev/issue/11911) runtime: unbounded heap growth with concurrent GC
- [issue/12041](https://go.dev/issue/12041) runtime: redesign GC synchronization
- [issue/27583](https://go.dev/issue/27583) runtime: feature request: Support for ASLR with the heap
- [doc/go119runtime](https://go.dev/blog/go119runtime) Michael Knyszek. Go runtime: 4 years later. September 26, 2022.
- [doc/cleanups-and-weak](https://go.dev/blog/cleanups-and-weak) Michael Knyszek. From unique to cleanups and weak: new low-level tools for efficiency. March 6, 2025.
- [talk/madhav2022](https://www.youtube.com/watch?v=We-8RSk4eZA) Madhav Jivrajani. Control Theory and Concurrent Garbage Collection Deep Dive. GopherCon 2022, Chicago. October 2022.
- [talk/tarun2022](https://www.youtube.com/watch?v=iy0yzxAtuwY) Tarun Pothulapati. Soft Memory Limit in Go 1.19. Lightning talk, GopherCon 2022, Chicago. October 2022.
- [talk/yifei2023](https://www.youtube.com/watch?v=DpQgJ06ZjGc) Yifei Zhang. Balanced GC: A Copying Garbage Collector for Golang. GopherCon 2023, San Diego. September 2023.

[Back To Top](#top)

### Statistics

- [issue/16843](https://go.dev/issue/16843) runtime: mechanism for monitoring heap size
  + [cl/setmaxheap](https://go-review.googlesource.com/c/go/+/46751/) Austin Clements. runtime/debug: add SetMaxHeap API. Jun 26 2017.
- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek. Proposal: API for unstable runtime metrics. Mar 18, 2020.
- [issue/19812](https://go.dev/issue/19812) runtime: cannot ReadMemStats during GC
- [issue/38712](https://go.dev/issue/38712) runtime: TestMemStats is flaky
- [issue/40459](https://go.dev/issue/40459) runtime: ReadMemStats called in a loop may prevent GC
- [issue/50285](https://go.dev/issue/50285) proposal: runtime/metrics: define a recommended set of metrics
- [issue/15490](https://go.dev/issue/15490) runtime/metrics: add goroutine state counts, total goroutines created, total threads
- [design/go-runtime-stats](https://docs.google.com/spreadsheets/d/1PTJjoWdgKJvY2S63V2Lh3VgLJrXXW3wO_IXxy2rNA8s) Go runtime stats.
- [issue/37112](https://go.dev/issue/37112) runtime: API for unstable metrics
- [issue/47216](https://go.dev/issue/47216) runtime/metrics: add GOMAXPROCS, NumCgoCall, GC CPU time
  + [issue/56857](https://go.dev/issue/56857) runtime/metrics: add /gc/heap/live:bytes
- [design/memlimits](https://docs.google.com/document/d/1zn4f3-XWmoHNj702mCCNvHqaS7p9rzqQGa74uOwOBKM) Memory Limits for Go (Public). Proposed in issue/16843.

[Back To Top](#top)

### Memory model

The Go memory model consists the following parts:

- Memory order regarding atomic operations
- Memory order regarding the `sync` package APIs
- Memory order regarding runtime mechanism (i.e. Object finalizer)

- [doc/refmem](https://go.dev/ref/mem) Rob Pike and Russ Cox. The Go Memory Model. February 21, 2009.
- [issue/4947]( https://go.dev/issue/4947) cmd/cc: atomic intrinsics
- [issue/5045](https://go.dev/issue/5045) doc: define how sync/atomic interacts with memory model
- [issue/7948](https://go.dev/issue/7948) doc: define how sync interacts with memory model
- [issue/9442](https://go.dev/issue/9442) doc: define how finalizers interact with memory model
- [issue/33815](https://go.dev/issue/33815) doc/go_mem: "hello, world" will not always be printed twice
- [cl/75130045](https://codereview.appspot.com/75130045) doc: allow buffered channel as semaphore without initialization
- [doc/gomem](http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf) Russ Cox. Go’s Memory Model. February 25, 2016.
- [doc/go2017russ](https://research.swtch.com/go2017#memory) Russ Cox. My Go Resolutions for 2017: Memory model. January 18, 2017.
- [doc/atomic-bug](https://go.dev/pkg/sync/atomic/#pkg-note-BUG) Package atomic
- [discuss/atomic-mem-order](https://groups.google.com/d/msg/golang-dev/vVkH_9fl1D8/azJa10lkAwAJ) specify the memory order guarantee provided by atomic Load/Store
- [issue/37355](https://go.dev/issue/37355) runtime/race: running with -race misses races (mismatch with memory model)
  - [cl/220419](https://go.dev/cl/220419) runtime: swap the order of raceacquire() and racerelease
  - [issue/42598](https://go.dev/issue/42598) runtime: apparent false-positive race report for a buffered channel after CL 220419
  - [cl/271987](https://go.dev/cl/271987) runtime: check channel's elemsize before calling race detector
  - [paper/fava2020fix] Fava, Daniel Schnetzer. "Finding and Fixing a Mismatch Between the Go Memory Model and Data-Race Detector." International Conference on Software Engineering and Formal Methods. Springer, Cham, 2020.
- [doc/mm](https://research.swtch.com/mm) Russ Cox. Memory Models. June, 2021.
  + [doc/hwmm](https://research.swtch.com/hwmm) Russ Cox. Hardware Memory Models. June 29, 2021.
  + [doc/plmm](https://research.swtch.com/plmm) Russ Cox. Programming Language Memory Models. July 6, 2021.
  + [doc/gomm](https://research.swtch.com/gomm) Russ Cox. Updating the Go Memory Model. July 12, 2021.
  + [discuss/47141](https://github.com/golang/go/discussions/47141) Updating the Go memory model.
  + [issue/50860](https://go.dev/issue/50860) proposal: sync/atomic: add typed atomic values
  + [issue/50859](https://go.dev/issue/50859) proposal: doc: update Go memory model
  + [doc/out-of-thin-air](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1217r2.html) Hans-J. Boehm. P1217R2: Out-of-thin-air, revisited, again. 2019-06-16.
  + [doc/cpp-memory-model](https://www.hpl.hp.com/techreports/2008/HPL-2008-56.pdf) Hans-J. Boehm, Sarita V. Adve. Foundations of the C++ Concurrency Memory Model. May 21, 2008.
  + [cl/381315](https://github.com/golang/go/commit/865911424d509184d95d3f9fc6a8301927117fdc) doc: update Go memory model
  + [cl/381316](https://github.com/golang/go/commit/a71ca3dfbd32faf351ff68bcc26a4d5abd9b06d7) runtime, sync, sync/atomic: document happens-before guarantees
- [discuss/align-64-bit-integers](https://groups.google.com/g/golang-dev/c/cUju6W3nWEU) align 64-bit integers on 32-bit platforms (golang-dev)
- [issue/61395](https://go.dev/issue/61395) sync/atomic: add OR/AND operators for unsigned types
- [issue/53821](https://go.dev/issue/53821) runtime: convert atomic values to runtime/internal/atomic types


[Back To Top](#top)

### ABI

- [design/cgo-pointers](https://go.dev/design/12416-cgo-pointers) Ian Lance Taylor. Proposal: Rules for passing pointers between Go and C. October, 2015
  + [issue/12416](https://go.dev/issue/12416) cmd/cgo: specify rules for passing pointers between Go and C
- [design/internal-abi](https://go.dev/design/27539-internal-abi) Austin Clements. Proposal: Create an undefined internal calling convention. Jan 14, 2019.
  + [issue/27539](https://go.dev/issue/27539) proposal: runtime: make the ABI undefined as a step toward changing it.
  + [issue/31193](https://go.dev/issue/31193) cmd/compile, runtime: reduce function prologue overhead
- [design/register-call](https://go.dev/design/40724-register-calling) Austin Clements, with input from Cherry Zhang, Michael Knyszek, Martin Möhrmann, Michael Pratt, David Chase, Keith Randall, Dan Scales, and Ian Lance Taylor. Proposal: Register-based Go calling convention. Aug 10, 2020.
- [issue/18597](https://go.dev/issue/18597) proposal: cmd/compile: define register-based calling convention
- [issue/40724](https://go.dev/issue/40724) proposal: switch to a register-based calling convention for Go functions
  + [cl/266638](https://go.dev/cl/266638) reflect,runtime: use internal ABI for selected ASM routines, attempt 2
  + [cl/259445](https://go.dev/cl/259445) cmd/compile,cmd/link: initial support for ABI wrappers
- [design/internal-abi-spec](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md) Go internal ABI specification.
- [issue/78902](https://go.dev/issue/78902) simd: architecture and vector-size agnostic SIMD intrinsics under a GOEXPERIMENT
- [issue/44816](https://go.dev/issue/44816) cmd/compile: register ABI bringup testcase failures (tracking issue)

[Back To Top](#top)

### Misc

- [issue/20135](https://go.dev/issue/20135) runtime: shrink map as elements are deleted
- [issue/48687](https://go.dev/issue/48687) runtime: enhance map cacheline efficiency
- [issue/54766](https://go.dev/issue/54766) runtime: use SwissTable
- [discuss/design-discuss-switch-hashmap](https://groups.google.com/g/golang-dev/c/R2O25Xz11zU/m/92mOve0eGy4J) Design discuss: Switch hashmap implementation to SipHash-2-4 (golang-dev)
- [issue/14874](https://go.dev/issue/14874) proposal: cmd/compile: compiler-emitted itab symbols
- [issue/61716](https://go.dev/issue/61716) math/rand/v2: revised API for math/rand
  + [discuss/60751](https://github.com/golang/go/discussions/60751) math/rand/v2: a new API for math/rand and a first v2 for std
- [issue/16093](https://go.dev/issue/16093) runtime/pprof: generate profiles in gzipped proto format
- [issue/76349](https://go.dev/issue/76349) runtime: add runtime/pprof goroutine labels to tracebacks
- [issue/42888](https://go.dev/issue/42888) runtime/debug: add SetCrashOutput(file *os.File)
  + [issue/37475](https://go.dev/issue/37475) cmd/go: stamp git/vcs current HEAD hash/commit hash/dirty bit in binaries
  + [issue/67182](https://go.dev/issue/67182) runtime/debug: add SetCrashOutput options
- [issue/3885](https://go.dev/issue/3885) runtime: profile and tune map code
- [issue/25237](https://go.dev/issue/25237) runtime: opportunistically rotate map key seed
- [issue/2630](https://go.dev/issue/2630) runtime: hash function should be randomized to prevent DOS attacks
- [issue/20505](https://go.dev/issue/20505) runtime: need a better itab table
- [issue/75280](https://go.dev/issue/75280) runtime/pprof,runtime: add new goroutine leak profile as a GOEXPERIMENT
- [doc/randv2](https://go.dev/blog/randv2) Russ Cox. Evolving the Go Standard Library with math/rand/v2. May 1, 2024.
- [doc/chacha8rand](https://go.dev/blog/chacha8rand) Russ Cox and Filippo Valsorda. Secure Randomness in Go 1.22. May 2, 2024.
- [doc/swisstable](https://go.dev/blog/swisstable) Michael Pratt. Faster Go maps with Swiss Tables. February 26, 2025.
- [doc/goroutine-leak-profiles](https://go.dev/blog/goroutine-leak-profiles) Vlad Saioc. Goroutine Leak Profiles. September 2, 2026.

## Standard Library

### syscall

- [design/go14syscall](https://go.dev/s/go1.4-syscall) The syscall package.
- [issue/51087](https://go.dev/issue/51087) runtime: clean up system call code
- [issue/68595](https://go.dev/issue/68595) syscall: support process sandboxing using Landlock on Linux

[Back To Top](#top)

### os, io, io/fs, embed

In Go 1.16, tons of major rework and improvements surround the new `os/fs` package.

- [design/draft-iofs](https://go.dev/design/draft-iofs) Russ Cox, Rob Pike. File System Interfaces for Go — Draft Design. July 2020.
  + [issue/13473](https://go.dev/issue/13473) proposal: os: Stdin, Stdout and Stderr should be interfaces
  + [issue/14106](https://go.dev/issue/14106) proposal: os: File should be an interface
  + [issue/19660](https://go.dev/issue/19660) proposal: io/ioutil: rename to io/fileio or similar
  + [issue/40025](https://go.dev/issue/40025) proposal: io/ioutil: move Discard, NopCloser, ReadAll to io
  + [issue/42027](https://go.dev/issue/40027) proposal: path/filepath: add WalkDir (Walk using DirEntry)
  + [issue/41190](https://go.dev/issue/41190) io/fs: add file system interfaces
  + [issue/41467](https://go.dev/issue/41467) os: add ReadDir method for lightweight directory reading
  + [issue/41974](https://go.dev/issue/41974) proposal: io/fs, filepath: add more efficient Walk alternative
  + [issue/42026](https://go.dev/issue/42026) proposal: os: add ReadDir, ReadFile, WriteFile, CreateTemp, MkdirTemp & deprecate io/ioutil
  + [issue/43223](https://go.dev/issue/43223) proposal: io/fs, net/http: define interface for automatic ETag serving
- [design/go116embed](https://go.dev/s/draft-embed-design) Embedded files - Russ & Braid
  + [issue/41191](https://go.dev/issue/41191) embed, cmd/go: add support for embedded files
  + [issue/42321](https://go.dev/issue/42321) embed: warn about dotfiles in embed.FS documentation
  + [issue/42328](https://go.dev/issue/42328) proposal: cmd/go: avoid surprising inclusion of "hidden" files when using //go:embed
  + [issue/43216](https://go.dev/issue/43216) proposal: embed: remove support for embedding directives on local variables
  + [issue/43217](https://go.dev/issue/43217) proposal: embed: define String and Bytes type aliases that must be used with //go:embed
  + [issue/43218](https://go.dev/issue/43218) embed: resolve string, []byte issues
  + [issue/44166](https://go.dev/issue/44166) io/fs,os: fs.ReadDir with an os.DirFS can produce invalid paths
  + [issue/42322](https://go.dev/issue/42322) io/fs: add func Sub(fsys FS, dir string) FS
- [issue/45757](https://go.dev/issue/45757) proposal: io/fs: add writable interfaces
- [issue/9051](https://go.dev/issue/9051) io: no easy way to fan out to multiple readers
- [issue/56219](https://go.dev/issue/56219) path/filepath: add IsLocal
- [issue/57151](https://go.dev/issue/57151) path/filepath: add Localize to safely convert a slash-separated path into an operating system path
- [issue/54898](https://go.dev/issue/54898) archive/zip: add (*Writer).AddFS
- [issue/62484](https://go.dev/issue/62484) os: add CopyFS
- [issue/67002](https://go.dev/issue/67002) os: safer file open functions
- [doc/osroot](https://go.dev/blog/osroot) Damien Neil. Traversal-resistant file APIs. March 12, 2025.

[Back To Top](#top)

### go/*

- [doc/gotypes](https://go.dev/s/types-tutorial) Alan Donovan. go/types: The Go Type Checker.
- [talk/gotypes](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view) Alan Donovan. Using go/types for
Code Comprehension and Refactoring Tools. October 2, 2015.
- [design/modular-interface](https://docs.google.com/document/d/1-azPLXaLgTCKeKDNg0HVMq2ovMlD-e7n1ZHzZVzOlJk/edit) Alan Donovan. Proposal: a common interface for modular static analyses for Go. 9 Sep 2018
  + [cl/134935](https://go.dev/cl/134935) go/analysis: a new API for analysis tools
- [issue/62039](https://go.dev/issue/62039) go/version: package for handling Go version strings

[Back To Top](#top)

### sync

- [design/percpu-sharded](https://go.dev/design/18802-percpu-sharded) Sanjay Menakuru. Proposal: percpu.Sharded, an API for reducing cache contention. Jun 12, 2018.
  + [issue/18802](https://go.dev/issue/18802) proposal: sync: support for sharded values
  + [cl/4850045](https://codereview.appspot.com/4850045/) co: new package
- [issue/37142](https://go.dev/issue/37142) sync: shrink types in sync package
- [issue/56102](https://go.dev/issue/56102) sync: add OnceFunc, OnceValue, OnceValues

[Back To Top](#top)

#### Map

- [issue/21031](https://go.dev/issue/21031) sync: reduce pointer overhead in Map
- [issue/21032](https://go.dev/issue/21032) sync: reduce (*Map).Load penalty for Stores with new keys
- [issue/21035](https://go.dev/issue/21035) sync: reduce contention between Map operations with new-but-disjoint keys
- [issue/37033](https://go.dev/issue/37033) runtime: provide centralized facility for managing (c)go pointer handles
- [issue/51972](https://go.dev/issue/51972) sync: add new Map methods CompareAndSwap, CompareAndDelete, Swap

[Back To Top](#top)

#### Pool

- [discuss/add-pool](https://groups.google.com/d/msg/golang-dev/kJ_R6vYVYHU/LjoGriFTYxMJ) gc-aware pool draining policy
- [issue/4720](https://go.dev/issue/4720) sync: add Pool type. Jan 28, 2013.
- [cl/46010043](https://github.com/golang/go/commit/f8e0057bb71cded5bb2d0b09c6292b13c59b5748#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: scalable Pool. Jan 24, 2014.
- [cl/86020043](https://github.com/golang/go/commit/8fc6ed4c8901d13fe1a5aa176b0ba808e2855af5#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: less agressive local caching in Pool. Apr 14, 2014.
- [cl/162980043](https://github.com/golang/go/commit/af3868f1879c7f8bef1a4ac43cfe1ab1304ad6a4#diff-491b0013c82345bf6cfa937bd78b690d) sync: release Pool memory during second and later GCs. Oct 22, 2014.
- [issue/8979](https://go.dev/issue/8979) sync: Pool does not release memory on GC
- [issue/22331](https://go.dev/issue/22331) runtime: clearpools causes excessive STW1 time
- [issue/22950](https://go.dev/issue/22950) sync: avoid clearing the full Pool on every GC.
  - [cl/166960](https://github.com/golang/go/commit/d5fd2dd6a17a816b7dfd99d4df70a85f1bf0de31) sync: use lock-free structure for Pool stealing.
  - [cl/166961](https://github.com/golang/go/commit/2dcbf8b3691e72d1b04e9376488cef3b6f93b286) 166961: sync: smooth out Pool behavior over GC with a victim cache.
- [issue/24479](https://go.dev/issue/24479) sync: eliminate global Mutex in Pool operations
- [discuss/pool-draining](https://groups.google.com/g/golang-dev/c/kJ_R6vYVYHU/m/LjoGriFTYxMJ) Russ Cox. gc-aware pool draining policy. golang-dev. November 27, 2013.

[Back To Top](#top)

#### Mutex, RWMutex

- [cl/4631059](https://go.dev/cl/4631059) runtime: replace Semacquire/Semrelease implementation.
- [issue/9201](https://go.dev/issue/9201) proposal: sync: prohibit unlocking mutex in a different goroutine
- [issue/13086](https://go.dev/issue/13086) runtime: fall back to fair locks after repeated sleep-acquire failures.
  + [cl/34310](https://go.dev/cl/34310) sync: make Mutex more fair
- [issue/17973](https://go.dev/issue/17973) sync: RWMutex scales poorly with CPU count
  + [cl/215361](https://go.dev/cl/215361) sync: Implement a version of RWMutex that can avoid cache contention

[Back To Top](#top)

#### Groups

- [cl/134395](https://go.dev/cl/134395) errgroup: rethink concurrency patterns
  + [cl/131815](https://go.dev/cl/131815) errgroup: handle runtime.Goexit from child goroutines
  + [issue/15758](https://go.dev/issue/15758) testing: complain loudly during concurrent use of T.FatalX and T.SkipX
  + [issue/25448](https://go.dev/issue/25448) proposal: promote panic(nil) to non-nil panic value
  + [issue/53757](https://go.dev/issue/53757) x/sync/errgroup: propagate panics and Goexits through Wait

#### atomic

- [issue/8739](https://go.dev/issue/8739) runtime,sync/atomic: unify API for runtime/internal/atomic and sync/atomic
- [issue/20164](https://go.dev/issue/20164) proposal: atomic: add (*Value).Swap
- [discuss/atomic-value](https://groups.google.com/g/golang-dev/c/SBmIen68ys0/m/WGfYQQSO4nAJ)
- [issue/36606](https://go.dev/issue/36606) proposal: cmd/compile: make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed directive on structs
- [issue/37262](https://go.dev/issue/37262) runtime: can't atomic access of first word of tiny-allocated struct on 32-bit architecture

[Back To Top](#top)

### time

- [design/monotonic-time](https://go.dev/design/12914-monotonic) Russ Cox. Proposal: Monotonic Elapsed Time Measurements in Go. January 26, 2017.
  + [issue/12914](https://go.dev/issue/12914) time: use monotonic clock to measure elapsed time
- Scalable Timer
  + [cl/34784](https://go.dev/cl/34784) runtime: improve timers scalability on multi-CPU systems
  + [issue/6239](https://go.dev/issue/6239) runtime: make timers faster.
  + [issue/15133](https://go.dev/issue/15133) runtime: timer doesn't scale on multi-CPU systems with a lot of timers
  + [issue/27707](https://go.dev/issue/27707) time: excessive CPU usage when using Ticker and Sleep.
- Followup latency issues
  - [issue/18023](https://go.dev/issue/18023) runtime: unexpectedly large slowdown with runtime.LockOSThread
  - [issue/25471](https://go.dev/issue/25471) time: Sleep requires ~7 syscalls
  - [issue/38070](https://go.dev/issue/38070) runtime: timer self-deadlock due to preemption point
  - [issue/36298](https://go.dev/issue/36298) net: 1.14 performance regression on mac
  - [issue/38860](https://go.dev/issue/38860) runtime: CPU bound goroutines cause unnecessary timer latency
    + [cl/216198](https://go.dev/cl/216198) runtime: add goroutines returned by poller to local run queue
    + [cl/232199](https://go.dev/cl/232199) runtime: steal timers from running P's
    + [cl/232298](https://go.dev/cl/232298) runtime: reduce timer latency
  - [issue/44343](https://go.dev/issue/44343) runtime: time.Sleep takes more time than expected
  - [issue/44868](https://go.dev/issue/44868) time, runtime: zero duration timer takes 2 minutes to fire

[Back To Top](#top)

### context

- [issue/8082](https://go.dev/issue/8082) proposal: spec: represent interfaces by their definition and not by package and name
- [issue/14660](https://go.dev/issue/14660) proposal: context: new package for standard library
- [issue/16209](https://go.dev/issue/16209) Proposal: relaxed rules for assignability with differently-named but identical interfaces
- [issue/20280](https://go.dev/issue/20280) proposal: io: add Context parameter to Reader, etc.
- [issue/21355](https://go.dev/issue/21355) proposal: Replace Context with goroutine-local storage
- [issue/24050](https://go.dev/issue/24050) testing: test with child process sometimes hangs on 1.10; -timeout not respected
- [issue/27982](https://go.dev/issue/27982) proposal: Go 2: bake a cooperative goroutine cancellation mechanism into the language
- [issue/28342](https://go.dev/issue/28342) proposal: Go 2: update context package for Go 2
- [issue/29011](https://go.dev/issue/29011) proposal: Go 2: use structured concurrency
- [doc/context-go-away](https://faiface.github.io/post/context-should-go-away-go2/) Michal Štrba. Context should go away for Go 2. 2017/08/06
- [doc/context](https://go.dev/blog/context) Go Concurrency Patterns: Context.
- [doc/context-isnt-for-cancellation](https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation) Dave Cheney. Context isn’t for cancellation. August 20, 2017.
- [issue/36503](https://go.dev/issue/36503) proposal: context: add Merge
- [issue/42564](https://go.dev/issue/42564) context: cancelCtx exclusive lock causes extreme contention
- [issue/51365](https://go.dev/issue/51365) proposal: context: add APIs for writing and reading cancelation cause
- [issue/56661](https://go.dev/issue/56661) context: add APIs for setting a cancelation cause when deadline or timer expires
- [issue/57928](https://go.dev/issue/57928) context: add AfterFunc

[Back To Top](#top)

### encoding

- [design/go12encoding](https://go.dev/s/go12encoding) Russ Cox. Go 1.2 encoding.TextMarshaler and TextUnmarshaler. July 2013.
- [design/go12xml](https://docs.google.com/document/d/1G-AMeUPoyTnbZDkuCJA89DjJwTjdWFctIZ_N9NgA9RM/pub) Russ Cox. Go 1.2 xml.Marshaler and Unmarshaler. July 2013.
- [design/natural-xml](https://go.dev/design/13504-natural-xml) Sam Whited. Proposal: Natural XML. 2016-09-27
  + [issue/13504](https://go.dev/issue/13504) encoding/xml: add generic representation of XML data
- [design/zip](https://go.dev/design/14386-zip-package-archives) Russ Cox. Proposal: Zip-based Go package archives. February 2016
  + [issue/14386](https://go.dev/issue/14386) proposal: use zip format inside .a and .o files
- [design/xmlstream](https://go.dev/design/19480-xml-stream) Sam Whited. Proposal: XML Stream. Mar 9, 2017.
  + [issue/19480](https://go.dev/issue/19480) encoding/xml: add decode from TokenReader, to enable stream transformers
- [design/raw-xml](https://go.dev/design/26756-rawxml-token) Sam Whited. Proposal: Raw XML Token. Sep 1, 2018.
  + [issue/26756](https://go.dev/issue/26756) proposal: encoding/xml: add RawXML token
- [discuss/63397](https://github.com/golang/go/discussions/63397) Joe Tsai. encoding/json/v2. October 2023.
  + [issue/71497](https://go.dev/issue/71497) encoding/json/v2: new API for encoding/json
  + [issue/71845](https://go.dev/issue/71845) encoding/json/v2: add new JSON API behind a GOEXPERIMENT=jsonv2 guard (Go 1.25)
  + [repo/json-experiment](https://github.com/go-json-experiment/json) Experimental implementation of a proposed v2 encoding/json package
- [doc/jsonv2-exp](https://go.dev/blog/jsonv2-exp) Joe Tsai, Daniel Martí, Johan Brandhorst-Satzkorn, Roger Peppe, Chris Hines, and Damien Neil. A new experimental Go API for JSON. September 9, 2025.
- [talk/joe2023](https://www.youtube.com/watch?v=avilmOcHKHE) Joe Tsai. The Future of JSON in Go. GopherCon 2023, San Diego. September 2023.

[Back To Top](#top)

### image, x/image

The following issues are surrounding by the color management of the `image` standard library.
At the moment, Go's `image` library doesn't read or write meta information from an image in
the encoding or decoding phase. Therefore the color information could go wrong while processing
an image such as scaling in a non-linear sRGB space. A universal solution is to design image
metadata APIs to aware the color profile in an encoded image file.

- [issue/11420](https://go.dev/issue/11420) x/image/draw: color space-correct interpolation
  + [issue/20613](https://go.dev/issue/20613) image/png: don't ignore PNG gAMA chunk
  + [issue/27830](https://go.dev/issue/27830) proposal: image: decoding options
  + [cl/253497](https://go.dev/cl/253497) x/image/draw: gamma corrected non linear interpolation
  + [issue/37188](https://go.dev/issue/37188) image/color: documentation doesn't include links to relevant color theory resources
- [issue/33457](https://go.dev/issue/33457) proposal: image: add generic metadata support for jpeg, gif, png
  + [discuss/image-metadata](https://groups.google.com/g/golang-dev/c/aRvnYIcaIaA/m/9GVKL7mIAgAJ) image.Metadata interface type
  + [issue/18365](https://go.dev/issue/18365) image/png: no support for setting and retrieving the PPI/DPI
  + [cl/208559](https://go.dev/cl/208559) image: New metadata-aware read/write API
  + [cl/216799](https://go.dev/cl/216799) image: metadata API sketch
- [issue/44808](https://go.dev/issue/44808) image, image/draw: add interfaces for using RGBA64 directly
- [issue/46395](https://go.dev/issue/46395) image/draw: increase performances by applying special case if mask is *image.Alpha
- [issue/8055](https://go.dev/issue/8055) image: decode / resize into an existing buffer
- [issue/11793](https://go.dev/issue/11793) image/color: NRGBA(64).RGBA() optimization
- [issue/15759](https://go.dev/issue/15759) image: optimize Image.At().RGBA()
- [issue/20851](https://go.dev/issue/20851) image: Decode drops interfaces
- [issue/24499](https://go.dev/issue/24499) image/jpeg: Decode is slow

<!--
TODO: read all of these!
These issues are discussion the current performance issue that exist in the current implementation.
- [issue/10447](https://go.dev/issue/10447) image/jpeg: add options to partially decode or tolerantly decode invalid images?
- [issue/13614](https://go.dev/issue/13614) image/jpeg: add a jpeg option to disable chroma subsampling
- [issue/22170](https://go.dev/issue/22170) image/jpeg: Unable to decode concatenated JPEGs (MIME-less "MJPEG")
- [issue/23936](https://go.dev/issue/23936) image/jpeg: encoding with RGB profile causing loss of image saturation
- [issue/29512](https://go.dev/issue/29512) image/jpeg: support for yuvj444p jpeg images
- [issue/40130](https://go.dev/issue/40130) image/jpeg: "bad RST marker" error when decoding
- [issue/6635](https://go.dev/issue/6635) image/gif: encoder does not honor image bounds.
- [issue/5050](https://go.dev/issue/5050) image/gif: decoding untrusted (very large) images can cause huge memory allocations
- [issue/26108](https://go.dev/issue/26108) image/gif: encoded images incompatible with some viewers
- [issue/20694](https://go.dev/issue/20694) image/gif: Mention the uselessness of BackgroundIndex in the docs Documentation
- [issue/20804](https://go.dev/issue/20804) image/gif: decoding gif returns `unknown block type: 0x01` error
- [issue/20856](https://go.dev/issue/20856) image/gif: decoding gif returns `frame bounds larger than image bounds` error
- [issue/33748](https://go.dev/issue/33748) image/gif: generated image cannot be opened in xv and crashes OmniWeb 3.x web browser
- [issue/35166](https://go.dev/issue/35166) image/gif: TestDecodeMemoryConsumption flake on dragonfly-amd64
- [issue/35503](https://go.dev/issue/35503) image/gif: decode fails with "gif: too much image data"
- [issue/38958](https://go.dev/issue/38958) image/gif: "not enough image data" on gif that works in browser
- [issue/38853](https://go.dev/issue/38853) image/gif: GIF files with extraneous 0x00 bytes cause "gif: unknown block type: 0x00"
- [issue/41142](https://go.dev/issue/41142) image/gif: Decode reads the entire animated gif image, even though it returns only the first frame (while DecodeAll exists to read and return all frames)-->

[Back To Top](#top)

### math

- [issue/8037](https://go.dev/issue/8037) cmd/internal/obj/x86: add fma
- [issue/17895](https://go.dev/issue/17895) spec: allow the use of fused multiply-add floating point instructions
  + [cl/40391](https://go.dev/cl/40391) spec: clarify use of fused-floating point operations
- [issue/24813](https://go.dev/issue/24813) math/bits: add extended precision Add, Sub, Mul, Div
- [issue/25819](https://go.dev/issue/25819) math: add guaranteed FMA

[Back To Top](#top)

### Mobile

- [design/go14android](https://go.dev/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/gobind](https://go.dev/s/gobind) David Crawshaw. Binding Go and Java. July 2014

[Back To Top](#top)

### log

- [discussion/54763](https://go.dev/issue/54763) discussion: structured, leveled logging
- [design/structured-logging](https://go.dev/design/56345-structured-logging) Proposal: Structured Logging
- [issue/56345](https://go.dev/issue/56345) proposal: log/slog: structured, leveled logging
- [issue/58243](https://go.dev/issue/58243) log/slog: make the current context easily available to loggers
- [issue/59928](https://go.dev/issue/59928) proposal: testing: Add log/slog support
- [doc/slog](https://go.dev/blog/slog) Jonathan Amsterdam. Structured Logging with slog. August 22, 2023.
- [issue/62005](https://go.dev/issue/62005) log/slog: add `slog.DiscardHandler`

### net, net/http

- [discuss/47323](https://github.com/golang/go/discussions/47323) Russ Cox. proposal: net/netaddr: add new IP address type, netaddr package (discussion). July 21, 2021.
- [issue/56539](https://go.dev/issue/56539) net: add support for MPTCP
- [discuss/60227](https://github.com/golang/go/discussions/60227) Jonathan Amsterdam. net/http: add methods and path variables to ServeMux patterns. May 16, 2023.
  + [issue/61410](https://go.dev/issue/61410) net/http: enhanced ServeMux routing
- [doc/routing-enhancements](https://go.dev/blog/routing-enhancements) Jonathan Amsterdam. Routing Enhancements for Go 1.22. February 13, 2024.
- [issue/66405](https://go.dev/issue/66405) net/http: expose matched pattern in Request
- [discuss/60746](https://github.com/golang/go/discussions/60746) Damien Neil. net/http: move HTTP/2 into the standard library. June 12, 2023.
- [issue/67813](https://go.dev/issue/67813) net/http: HTTP/2 configuration API
- [issue/73626](https://go.dev/issue/73626) net/http: add CrossOriginForgeryHandler
- [issue/75772](https://go.dev/issue/75772) net/http: client connection API
- [doc/http-connection-reuse](https://go.dev/blog/http-connection-reuse) Nicholas Husin and Damien Neil. Better HTTP/1 connection reuse. October 14, 2026.
- [issue/58547](https://go.dev/issue/58547) x/net/quic: add QUIC implementation

[Back To Top](#top)

### crypto

- [issue/66821](https://go.dev/issue/66821) crypto/rand: crash process on error reading randomness
- [issue/67057](https://go.dev/issue/67057) crypto/rand: add func Text
- [issue/61477](https://go.dev/issue/61477) crypto/hkdf: add package
- [issue/69488](https://go.dev/issue/69488) crypto/pbkdf2: add package
- [issue/63369](https://go.dev/issue/63369) crypto/tls: support Encrypted Client Hello in clients
- [issue/70122](https://go.dev/issue/70122) crypto/mlkem: new package
- [issue/75300](https://go.dev/issue/75300) crypto/hpke: new package
- [issue/77626](https://go.dev/issue/77626) crypto/mldsa: new package
- [issue/70123](https://go.dev/issue/70123) crypto: mechanism to enable FIPS mode
- [issue/70200](https://go.dev/issue/70200) cmd/go: add fips140 module selection mechanism
- [doc/tob-crypto-audit](https://go.dev/blog/tob-crypto-audit) Roland Shoemaker and Filippo Valsorda. Go Cryptography Security Audit. May 19, 2025.
- [doc/fips140](https://go.dev/blog/fips140) Filippo Valsorda (Geomys), Daniel McCarney (Geomys), and Roland Shoemaker (Google). The FIPS 140-3 Go Cryptographic Module. July 15, 2025.
- [design/crypto-ssh-v2](https://go.dev/design/68723-crypto-ssh-v2) Nicola Murino, Filippo Valsorda. x/crypto/ssh/v2. August 19, 2024.

[Back To Top](#top)

### misc

- [design/mobile-audio](https://go.dev/design/13432-mobile-audio) Jaana Burcu Dogan. Proposal: Audio for Mobile. November 30, 2015.
  + [issue/13432](https://go.dev/issue/13432) proposal: x/mobile audio
- [design/localization](https://go.dev/design/12750-localization) Marcel van Lohuizen. Proposal: Localization support in Go. Jan 28, 2016.
  + [issue/12750](https://go.dev/issue/12750) x/text: localization support
- [design/unbounded-queue](https://go.dev/design/27935-unbounded-queue-package) Christian Petrin. Proposal: Built in support for high performance unbounded queue. October 2, 2018
  + [issue/27935](https://go.dev/issue/27935) proposal: add container/queue
- [design/lockfile](https://go.dev/design/33974-add-public-lockedfile-pkg) Adrien Delorme. Proposal: make the internal lockedfile package public. 2019-10-15.
  + [issue/33974](https://go.dev/issue/33974)
- [design/cidr](https://go.dev/design/16704-cidr-notation-no-proxy) Rudi Kramer, James Forrest. Proposal: Add support for CIDR notation in no_proxy variable. 2017-07-10
  + [issue/16704](https://go.dev/issue/16704) net/http: considering supporting CIDR notation in no_proxy env variable
- [design/dns](https://go.dev/design/26160-dns-based-vanity-imports) Sam Whited. Proposal: DNS Based Vanity Imports. 2018-06-30.
  + [issue/26160](https://go.dev/issue/26160) proposal: use DNS TXT records for vanity import paths
- [issue/46518](https://go.dev/issue/46518) net/netip: add new IP address package, use in net
- [issue/53171](https://go.dev/issue/53171) proposal: add package for using SIMD instructions
- [issue/54880](https://go.dev/issue/54880) math/rand: seed global generator randomly
- [issue/60204](https://go.dev/issue/60204) cmp: add Or
- [issue/60088](https://go.dev/issue/60088) reflect: add TypeFor
- [issue/60370](https://go.dev/issue/60370) database/sql: add generic Null[T]
- [issue/66408](https://go.dev/issue/66408) structs: add HostLayout "directive" type
- [issue/54670](https://go.dev/issue/54670) hash/maphash: add func Comparable
- [issue/70471](https://go.dev/issue/70471) hash/maphash: standardize the hash function
- [issue/69264](https://go.dev/issue/69264) x/exp/stats: new package with Mean, Median, more
- [issue/62026](https://go.dev/issue/62026) uuid: add API to generate and parse UUID

[Back To Top](#top)

## Unclassified But Relevant Links

- [Using Go Guru: an editor-integrated tool for navigating Go code By Alan Donovan](https://go.dev/s/using-guru)
- [Plan for Go Guru (née Oracle) By Alan Donovan](https://docs.google.com/document/d/1UErU12vR7jTedYvKHVNRzGPmXqdMASZ6PfE7B-p6sIg/edit)
- [(Abandoned WIP) Why is X in the standard library?](https://go.dev/s/stdwhy)
- [go oracle: design By Alan Donovan](https://go.dev/s/oracle-design)
- [go oracle: user manual By Alan Donovan](https://go.dev/s/oracle-user-manual)
- [cgihttpproxy](https://go.dev/s/cgihttpproxy)
- [sqldrivers](https://go.dev/s/sqldrivers)
- [go2designs](https://go.dev/s/go2designs)
- [Notes on Programming in C By Rob Pike](http://doc.cat-v.org/bell_labs/pikestyle)
- ["The Best Programming Advice I Ever Got" with Rob Pike](https://www.informit.com/articles/article.aspx?p=1941206)
- [An Interview with Brian Kernighan By Mihai Budiu](http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html)
- [Language Design with Brian Kernighan Holiday Repeat By SE Daily](https://softwareengineeringdaily.com/2017/12/28/language-design-with-brian-kernighan-holiday-repeat/)
- ["C" Programming Language: Brian Kernighan - Computerphile](https://youtu.be/de2Hsvxaf8M) "C" Programming Language: Brian Kernighan - Computerphile
- [Debugging performance issues in Go* programs By Dmitriy Vyukov](https://software.intel.com/content/www/us/en/develop/blogs/debugging-performance-issues-in-go-programs.html)
- [Interesting papers I'd like to implement (or at least have implementations of)](https://github.com/dgryski/interesting-papers)
- [golang/wiki/ResearchPapers](https://github.com/golang/go/wiki/ResearchPapers)
- John Ousterhout. Open Decision-Making. June, 2021. https://web.stanford.edu/~ouster/cgi-bin/decisions.php.
- [Go’s Version Control History](https://research.swtch.com/govcs)
- [talk/jake2025](https://www.youtube.com/watch?v=PZm_YbE3fcA) Jake Bailey. Porting the TypeScript Compiler to Go for a 10x Speedup. GopherCon 2025, New York City. August 2025.

[Back To Top](#top)

## Fun Facts

- [cl/1](https://github.com/golang/go/commit/7d7c6a97f815e9279d08cfaea7d5efb5e90695a8) Brian Kernighan. Go's first commit. Jul 19, 1972.
- [issue/9](https://go.dev/issue/9) I have already used the name for *MY* programming language. Nov 11, 2009
- [issue/2870](https://go.dev/issue/2870) 9 is prime if it's a hot day. Feb 3, 2012
- [doc/gophercount](https://research.swtch.com/gophercount) How Many Go Developers Are There?. November 1, 2019.
- [discuss/google-owns-go](https://groups.google.com/forum/#!msg/golang-nuts/6dKNSN0M_kg/Y1yDJRwQBgAJ) Russ Cox's response on "Go is Google's language, not ours"

[Back To Top](#top)

## Acknowledgements

The document author would like to first thank the [TalkGo](https://github.com/talkgo)
community creator [Mai Yang](https://github.com/yangwenmai)'s champion sponsorship
for the [golang.design](https://golang.design) initiative. His creation of
the TalkGo significantly changed the Go community in China. He is also a great person
that is actively contributing to all kinds of Go related projects.

It is also important to thank the continuing, inspiring discussion and sharing with
the TalkGo community core members [qcrao](https://github.com/qcrao),
[aofei](https://github.com/aofei), and [eddycjy](https://github.com/eddycjy).

The document would not be organized without all of the supports from them.

![](https://changkun.de/urlstat?mode=github&repo=golang-design/history)

[Back To Top](#top)

## License

[golang.design/history](https://github.com/golang-design/history) | CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)
