# [Go: A Documentary](https://golang.design/history)

_by Changkun Ou <[changkun.de](https://changkun.de)>_ (and many inputs from [contributors](https://github.com/golang-design/history/graphs/contributors))

This document collects many interesting (publicly observable) issues,
discussions, proposals, CLs, and talks from the Go development process,
which intents to offer a comprehensive reference of the Go history.

<a id="top"></a>

**Table of Contents**

- [Go: A Documentary](#go-a-documentary)
  - [Disclaimer](#disclaimer)
  - [Sources](#sources)
  - [Origin](#origin)
  - [Committers](#committers)
    - [Core Authors](#core-authors)
    - [Compiler/Runtime Team](#compilerruntime-team)
    - [Library/Tools/Security/Community](#librarytoolssecuritycommunity)
    - [Group Interviews](#group-interviews)
  - [Timeline](#timeline)
  - [Language Design](#language-design)
    - [Misc](#misc)
    - [Slice (1.2)](#slice-12)
    - [Package Management (1.4, 1.5, 1.7)](#package-management-14-15-17)
    - [Type alias (1.9)](#type-alias-19)
    - [Defer (1.14)](#defer-114)
    - [Error values (1.13)](#error-values-113)
    - [Channel/Select](#channelselect)
    - [Generics](#generics)
  - [Compiler Toolchain](#compiler-toolchain)
    - [Compiler](#compiler)
    - [Linker](#linker)
    - [Debugger](#debugger)
    - [Race Detector](#race-detector)
    - [Tracer](#tracer)
    - [Lock Analysis](#lock-analysis)
    - [Builder](#builder)
    - [Modules](#modules)
    - [gopls](#gopls)
    - [Testing](#testing-xperf)
  - [Runtime Core](#runtime-core)
    - [Scheduler](#scheduler)
    - [Execution Stack](#execution-stack)
    - [Memory Allocator](#memory-allocator)
    - [Garbage Collector](#garbage-collector)
    - [Statistics](#statistics)
    - [Memory model](#memory-model)
    - [ABI](#abi)
    - [Misc](#misc-1)
  - [Standard Library](#standard-library)
    - [syscall](#syscall)
    - [os, io, io/fs, embed](#os-io-iofs-embed)
    - [go/*](#go)
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
    - [Mobile](#mobile)
    - [misc](#misc-2)
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
to the project that you might interest in following their excellent work.

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

The Go was created by Rob, Robert, and Ken initially because
Rob were suffered by the slow C++ compiling time, talked to Robert, and luckily
Ken was in the next office.
Later, Ian joined the project since he showed huge interests and
wrote the [gccgo](https://github.com/golang/gofrontend).
Rob and Ken are retired. Robert and Ian currently work on adding generics
to Go. Russ is also one of the core authors of the project in the early stage.
Back then, he was a newcomer at Google, and Rob invited Russ for joining
the Go team since he knew Russ from way back because of the
[Plan 9](https://9p.io/plan9/) project. Russ did many
fundamental work for the early Go compiler, runtime, as well as the leap of
Go 1.5 bootstrap.
Now, Russ is the tech leader of the Go team.

- Rob Pike. (Robert C. Pike, M. Sc.) [Website](http://herpolhode.com/rob/), [Blog](https://commandcenter.blogspot.com/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike). (Retired)
  + Alma mater: University of Toronto (BS), California Institute of Technology
  + [talk/rob2007](https://www.youtube.com/watch?v=hB05UFqOtFA) Advanced Topics in Programming Languages: Concurrency/message passing Newsqueak. May 9, 2007
  + [talk/rob2009](https://changelog.com/podcast/3) The Go Programming Language. Nov 27, 2009.
  + [talk/rob2010a](https://www.youtube.com/watch?v=jgVhBThJdXc) Go Programming. Google I/O 2010. May 20, 2010
  + [talk/rob2010b](https://www.youtube.com/watch?v=-_DKfAn4pFA) Origins of Go Concurrency style. Emerging Languages Camp 2010. July 21, 2010.
  + [talk/rob2010c](https://www.youtube.com/watch?v=5kj5ApnhPAE) Public Static Void. OSCON 2010. Jul 22, 2010.
  + [talk/rob2010d](https://www.youtube.com/watch?v=7VcArS4Wpqk) Another Go at Language Design. Aug 27, 2010
  + [talk/rob2011a](https://www.infoq.com/interviews/pike-concurrency/) Parallelism and Concurrency in Programming Languages. Feb 17, 2011.
  + [talk/rob2011b](https://www.infoq.com/interviews/pike-google-go/) Google Go: Concurrency, Type System, Memory Management and GC. Feb 25, 2011.
  + [talk/rob2011c](https://www.youtube.com/watch?v=HxaD_trXwRE) Lexical Scanning in Go. Aug 30, 2011
  + [talk/rob2012a](https://www.youtube.com/watch?v=f6kdp27TYZs) Go Concurrency Patterns. Google I/O 2012. Jul 2, 2012.
  + [talk/rob2012b](https://www.youtube.com/watch?v=FTl0tl9BGdc) Why Learn Go?. Sep 12, 2012.
  + [talk/rob2013a](https://www.youtube.com/watch?v=bj9T2c2Xk_s) The path to Go 1. Mar 14, 2013.
  + [talk/rob2013b](https://www.infoq.com/presentations/Go-Google/) Go at Google. Apr 13, 2013.
  + [talk/rob2013c](https://changelog.com/podcast/100) Go Programming with Rob Pike and Andrew Gerrand. Aug 14, 2013.
  + [talk/rob2013d](https://www.youtube.com/watch?v=cN_DpYBzKso) Concurrency Is Not Parallelism. Oct 20, 2013
  + [talk/rob2014a](https://www.youtube.com/watch?v=VoS7DsT1rdM) Hello Gophers! Gophercon Opening Keynote. GopherCon 2014. Apr 24, 2014
  + [talk/rob2014b](https://www.youtube.com/watch?v=PXoG0WX0r_E) Implementing a bignum calculator. Sydney Go meeup. Nov, 19 2014.
  + [talk/rob2015a](https://www.youtube.com/watch?v=cF1zJYkBW4A) The move from C to Go in the toolchain. GopherFest 2015. Jun 10, 2015.
  + [talk/rob2015b](https://www.youtube.com/watch?v=PAAkCSZUG1c) Go Proverbs. Gopherfest. Nov 18, 2015
  + [talk/rob2015c](https://www.youtube.com/watch?v=rFejpH_tAHM) Simplicity is Complicated. dotGo 2015. Dec 2, 2015
  + [talk/rob2016a](https://www.youtube.com/watch?v=KINIAgRpkDA) The Design of the Go Assembler. GopherCon 2016. Aug 18, 2016.
  + [talk/rob2016b](https://www.youtube.com/watch?v=sDTGhIqyMjo) Stacks of Tokens: A study in interfaces. Sydney Go Meetup. September 2016.
  + [talk/rob2017](https://www.youtube.com/watch?v=ENLWEfi0Tkg) Upspin. Gopherfest 2017. Jun 22, 2017.
  + [talk/rob2018a](https://www.youtube.com/watch?v=_2NI6t2r_Hs) The History of Unix. Nov 7, 2018.
  + [talk/rob2018b](https://www.youtube.com/watch?v=RIvL2ONhFBI) Go 2 Draft Specifications. Sydney Golang Meetup. Nov 13, 2018
  + [talk/rob2019a](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer. Sep 10, 2019.
  + [talk/rob2019b](https://www.youtube.com/watch?v=oU9cfQCxjpM) Rob Pike. A Brief History of Go. Dec 12, 2019.
  + [talk/rob2020](https://evrone.com/rob-pike-interview) A Rob Pike Interview. (Date Unclear) 2020.
  + [talk/rob2021](https://www.youtube.com/watch?v=YXV7sa4oM4I) The Go Programming Language and Environment. John Lions Distinguished Lecture, UNSW, May 27, 2021.
- Robert Griesemer. (Dr. Robert Griesemer) [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
  + Alma mater: ETH Zürich
  + [paper/robert1993](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.127.5290&rep=rep1&type=pdf) A Programming Language for Vector Computers. Doctor Dissertation.
  + [talk/robert2012a](https://www.youtube.com/watch?v=on5DeUyWDqI) E2E: Erik Meijer and Robert Griesemer. Going Go. Lang.NEXT. 2012.
  + [talk/robert2012b](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Go-In-Three-Easy-Pieces) Go In Three Easy Pieces. Mar 19, 2012.
  + [talk/robert2012c](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Panel-Native-Languages) Lang.NEXT 2012 Expert Panel: Native Languages. Apr 11, 2012.
  + [talk/robert2015](https://www.youtube.com/watch?v=0ReKdcpNyQg) The Evolution of Go. GopherCon 2015. Jul 28, 2015.
  + [talk/robert2016a](https://www.youtube.com/watch?v=t-w6MyI2qlU) Lightning Talk: Alias Declarations for Gom: A proposal. GopherCon 2016. Oct 9, 2016.
  + [talk/robert2016b](https://www.youtube.com/watch?v=vLxX3yZmw5Q) Prototype your design!. dotGo 2016. Nov 29, 2016.
  + [talk/robert2017](https://www.youtube.com/watch?v=KPk1UPihWtY) Opening Keynote: Exporting Go. GopherCon SG 2017. May 29, 2017.
  + [talk/robert2017](https://www.youtube.com/watch?v=UmwJsQTSEP8) A brief overview of Go. Aug 28. 2017.
  + [talk/robert2019](https://www.youtube.com/watch?v=i0zzChzk8KE) Go is 10! Now What?. Gopherpalooza 2019. Dec 2, 2019.
  + [talk/robert2020a](https://changelog.com/gotime/140) The latest on Generics with Robert Griesemer and Ian Lance Taylor. Jul 21, 2020.
  + [talk/robert2020b](https://www.gophercon.com/agenda/session/233094) Typing [Generic] Go. Nov 11, 2020.
- Ken Thompson. (Kenneth Lane Thompson, M. Sc.) (Retired)
  + Alma mater: UC Berkeley
  + [talk/ken1982a](https://www.youtube.com/watch?v=tc4ROCJYbm0) The UNIX System: Making Computers More Productive. 1982.
  + [talk/ken1982b](https://www.youtube.com/watch?v=XvDZLjaCJuw) UNIX: Making Computers Easier To Use.
  + [talk/ken1982c](https://www.youtube.com/watch?v=JoVQTPbD6UY) Ken Thompson and Dennis Ritchie Explain UNIX (Bell Labs).
  + [talk/ken1983](https://www.youtube.com/watch?v=LXZ1OL2U3lY) Ken Thompson and Dennis Ritchie. National Medal of Technology Awards. 1996.
  + [talk/ken2013](https://www.youtube.com/watch?v=dsMKJKTOte0) Systems Architecture, Design, Engineering, and Verification. Jan 17, 2013.
  + [talk/ken2019a](https://www.youtube.com/watch?v=g3jOJfrOknA) The Thompson and Ritchie Story. Feb 18, 2019.
  + [talk/ken2019b](https://www.youtube.com/watch?v=EY6q5dv_B-o) Brian Kernighan interviews Ken Thompson. VCF East 2019. May 4, 2019.
- Ian Taylor. (Ian Lance Taylor, B. Sc.) [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor), [Quora](https://www.quora.com/profile/Ian-Lance-Taylor)
  + Alma mater: Yale University
  + [talk/ian2007](https://www.youtube.com/watch?v=gc78olyguqA) GCC: Current Topics and Future Directions. February 27, 2007.
  + [talk/ian2018](https://www.youtube.com/watch?v=LqKOY_pH8u0) Transition to Go 2. Gopherpalooza 2018. Oct 24, 2018
  + [talk/ian2019a](https://www.youtube.com/watch?v=WzgLqE-3IhY) Generics in Go. GopherCon 2019. Aug 27, 2019
  - [talk/ian2019b](https://changelog.com/gotime/98) Generics in Go. Aug 27, 2019.
  + [talk/ian2020](https://www.youtube.com/watch?v=yoZ05GG8aLs) Go with Ian Lance Taylor. CppCast. Aug 9, 2020.
- Russ Cox. (Dr. Russell Stensby Cox) [Website](https://swtch.com/~rsc/), [Blog](https://research.swtch.com/), [GitHub](https://github.com/rsc), [Twitter](https://twitter.com/_rsc), [Reddit](https://old.reddit.com/user/rsc), [YouTube](https://www.youtube.com/channel/UC3P6PrEBAVH1UaiPOzZ-u-w)
  + Alma mater: MIT
  + [paper/russ2008](https://pdos.csail.mit.edu/~rsc/) An Extension-Oriented Compiler. Doctor Dissertation.
  + [talk/russ2009](https://www.youtube.com/watch?v=wwoWei-GAPo) The Go Programming Language Promo. Nov 10, 2009.
  + [talk/russ2012](https://www.youtube.com/watch?v=MzYZhh6gpI0) A Tour of the Go Programming Language. Jun 24, 2012.
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
  + [talk/russ2021](https://www.twitch.tv/videos/1203523364) #PLTalk: 12 Years of Go with Russ Cox. Nov. 12, 2021.

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
his work mainly in the runtime memory system such as the refactoring of memory allocator and runtime metrics.

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
- Austin Clements. (Dr. Austin T. Clements) [GitHub](https://github.com/aclements), [Scholar](https://scholar.google.com/citations?user=MKDtxN4AAAAJ)
  + Alma mater: MIT
  + [paper/autin2014](https://pdos.csail.mit.edu/papers/aclements-phd.pdf) The Scalable Commutativity Rule: Designing Scalable Software for Multicore Processors. Doctor Dissertation
  + [talk/austin2020](https://www.gophercon.com/agenda/session/233441) Pardon the Interruption: Loop Preemption in Go 1.14. Nov 12, 2020.
- Richard Hudson. (Richard L. Hudson, M. Sc.) [GitHub](https://github.com/RLH) (Retired)
  + Alma mater: University of Massachusetts Amherst
  + [paper/rick](https://dl.acm.org/profile/81100566849/publications?Role=author) Research List
  + [talk/rick2015](https://www.youtube.com/watch?v=aiv1JOfMjm0) Go GC: Solving the Latency Problem. GopherCon 2015. Jul 8, 2015.
  + [talk/rick2015b](https://www.infoq.com/interviews/hudson-go-gc/) Rick Hudson on Garbage Collection in Go. Dec 21, 2015
- Keith Randall. (Dr. Keith H. Randall.) [GitHub](https://github.com/randall77)
  + Alma mater: MIT
  + [paper/1998keith](http://supertech.csail.mit.edu/papers/PPoPP95.pdf) Cilk: Efficient Multithreaded Computing.  Doctor Dissertation.
  + [talk/keith2016](https://www.youtube.com/watch?v=Tl7mi9QmLns) Inside the Map Implementation. GopherCon 2016. Jul 12, 2016.
  + [talk/keith2017](https://www.youtube.com/watch?v=uTMvKVma5ms) Generating Better Machine Code with SSA. GopherCon 2017. Jul 24, 2017.
- David Chase. (Dr. David Chase) [Website](http://chasewoerner.org/resume.html), [Block](https://dr2chase.wordpress.com/), [GitHub](https://github.com/dr2chase), [Twitter](https://twitter.com/dr2chase), [Scholar](https://dblp.org/pid/51/3488.html)
  + Alma mater: Rice University
  + [paper/1987david](http://www.chasewoerner.org/dissertation.pdf) Garbage Collection and Other Optimizations. Doctor Dissertation.
  + [talk/david2017](https://changelog.com/gotime/52) All About The Go Compiler. Jul 20, 2017.
- Dan Scales.
  + [talk/dan2020](https://www.gophercon.com/agenda/session/233397) Implementing Faster Defers. Nov 11, 2020
- Michael Knyszek. [Website](https://www.ocf.berkeley.edu/~mknyszek/), [GitHub](https://github.com/mknyszek)
  + [talk/michael2020](https://www.gophercon.com/agenda/session/233086) Evolving the Go Memory Manager's RAM and CPU Efficiency. Nov 11, 2020
- Than McIntosh. [GitHub](https://github.com/thanm)
- Cherry Zhang. [GitHub](https://github.com/cherrymui)

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
- Jonathan Amsterdam.
  + [talk/jonathan2020](https://www.gophercon.com/agenda/session/233432) Working with Errors. Nov 13, 2020.
- Daniel Martí. [Website](https://mvdan.cc/), [GitHub](https://github.com/mvdan), [Twitter](https://twitter.com/mvdan_)
- Nigel Tao. [GitHub](https://github.com/nigeltao), Twitter
- Filippo Valsorda. GitHub,
- Michael Matloob. [GitHub](https://github.com/matloob), [Twitter](https://twitter.com/matloob)
- Jay Conrod. [Website](https://jayconrod.com/), [Twitter](https://twitter.com/jayconrod)
- Dave Cheney. [Website](https://dave.cheney.net/), [GitHub](https://github.com/davecheney), [Twitter](https://twitter.com/davecheney)
- Sam Boyer. [GitHub](https://github.com/sdboyer), [Twitter](https://twitter.com/sdboyer)
- Fillippo Valsorda. [Website](https://filippo.io/), [GitHub](https://github.com/FiloSottile), [Twitter](https://twitter.com/FiloSottile)
  + [talk/filo2016](https://www.youtube.com/watch?v=lhMhApWQp2E) From cgo back to Go. July 12, 2016
  + [talk/filo2017](https://speakerdeck.com/filosottile/calling-rust-from-go-without-cgo-at-gothamgo-2017) From Go to Rust without cgo. 
  + [talk/filo2018](https://speakerdeck.com/filosottile/why-cgo-is-slow-at-capitalgo-2018) Why cgo is slow. CapitalGo 2018.
  + [talk/speakerdeck](https://speakerdeck.com/filosottile?page=1)

> more people and talks should be added...

[Back To Top](#top)

### Group Interviews

- [talk/goteam2012](https://www.youtube.com/watch?v=sln-gJaURzk) Meet the Go Team. Google I/O 2012. Jul 2, 2012.
- [talk/goteam2013](https://www.youtube.com/watch?v=p9VUCp98ay4) Fireside Chat with the Go Team. Google I/O 2013. May 18, 2013.
- [talk/goteam2014](https://www.youtube.com/watch?v=u-kkf76TDHE) Inside the Gophers Studio with Blake Mizerany. GopherCon 2014. May 21, 2014
- [talk/goteam2019](https://www.youtube.com/watch?v=3yghHvvZQmA) Meet the Authors: Go Language. Cloud Next '19. Apr 10, 2019.
- [talk/goteam2020a](https://www.youtube.com/watch?v=gJxvkOHpTSM) GoLab 2020: Go Team AMA. Oct 22, 2020.
- [talk/goteam2020b](https://www.youtube.com/watch?v=BNHwHLNLjLs) GopherCon 2020: Go Team AMA. Nov 16, 2020.

[Back To Top](#top)

## Timeline

A timeline helps you identify the point in time about a document that is
reflected in Go versions.

| Version | Release Date |
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
| Go 1.16 | 2021.02.16 |
| Go 1.17 | 2021.08.16 |

The historical release notes may helpful for general information:

- [doc/go1release](https://go.dev/doc/devel/release.html) Go Release History
- [doc/go1prerelease](https://go.dev/doc/devel/pre_go1.html) Pre-Go 1 Release History
- [doc/go0release](https://go.dev/doc/devel/weekly.html) Weekly Release History (Before Go 1)

[Back To Top](#top)

## Language Design

### Misc

- [design/go0initial](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) Rob Pike, Robert Griesemer, Ken Thompson. The Go Annotated Specification. Mar 3, 2008.
- [design/go0spec0](https://github.com/golang/go/blob/e6626dafa8de8a0efae351e85cf96f0c683e0a4f/doc/go_lang.txt) The Go Programming Language. Language Specification (March 7, 2008)
- [design/go0semicolon](https://go.dev/s/semicolon-proposal) Rob Pike. Semicolons in Go. Dec 10, 2009.
- [design/go11func](https://go.dev/s/go11func) Russ Cox. Go 1.1 Function Calls. February 2013
- [design/go11return](https://go.dev/s/go11return) Russ Cox. Go “Return at End of Function” Requirements. March 2013
- [design/go12nil](https://go.dev/s/go12nil) Russ Cox. Go 1.2 Field Selectors and Nil Checks. July 2013.
- [doc/go13todo](https://go.dev/s/go13todo) Go 1.3 “To Do” List
- [doc/goatgoogle](https://go.dev/talks/2012/splash.article#TOC_12.) Rob Pike. Go at Google - Language Semantics. October 25, 2012.
- [doc/makego](https://go.dev/talks/2015/how-go-was-made.slide)  Andrew Gerrand. How Go was Made. 9 July 2015.
- [discuss/go1preview](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) Russ Cox. Go 1 Preview.
- [design/overlapping-interfaces](https://go.dev/design/6977-overlapping-interfaces) Robert Griesemer. Proposal: Permit embedding of interfaces with overlapping method sets. 2019-10-16
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
- [issue/33502](https://go.dev/issue/33502) proposal: review meeting minutes
- [issue/33892](https://go.dev/issue/33892) proposal: Go 2 review meeting minutes
- [issue/19623](https://go.dev/issue/19623) proposal: spec: change int to be arbitrary precision
- [issue/19367](https://go.dev/issue/19367) unsafe: add Slice(ptr *T, len anyIntegerType) []T
- [issue/40481](https://go.dev/issue/40481) unsafe: add Add function

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

[Back To Top](#top)

### Error values (1.13)

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
- [design/err-try](https://go.dev/design/32437-try-builtin) Robert Griesemer. Proposal: A built-in Go error check function, "try". 2019-06-12
  - [issue/32437](https://go.dev/issue/32437#issuecomment-512035919) Proposal: A built-in Go error check function, "try". Decision response.
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

[Back To Top](#top)

### Channel/Select

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov. Go channels on steroids. Jan 28, 2014
  + [issue/8899](https://go.dev/issue/8899) runtime: lock-free channels
  + [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) update on "lock-free channels"
  + [cl/112990043](https://codereview.appspot.com/112990043/) runtime: fine-grained locking in select
  + [cl/110580043](https://codereview.appspot.com/110580043/) runtime: add fast paths to non-blocking channel operations
- [issue/8898](https://go.dev/issue/8898) runtime: special case timer channels
- [issue/37196](https://go.dev/issue/37196) time: make Timer/Ticker channels not receivable with old values after Stop or Reset returns
- [issue/8903](https://go.dev/issue/8903) runtime: make chan-based generators faster.
- [issue/21806](https://go.dev/issue/21806) fairness in select statement.
- [issue/40410](https://go.dev/issue/40410) runtime: trim unnecessary fields from scase
- [issue/40641](https://go.dev/issue/40641) runtime: race between stack shrinking and channel send/recv leads to bad sudog values
- [issue/37350](https://go.dev/issue/37350) reflect: Select panics if array length greater than 1<<16

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
- [issue/33232](https://go.dev/issue/33232) proposal: Go 2: add alias for interface {} as any
- [discuss/generics-move](https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw) Ian, Moving forward with the generics design.
- [discuss/generics-implementation](https://groups.google.com/g/golang-dev/c/OcW0ATRS4oM) Implementing Generics
- [design/generics-implementation-stenciling](https://go.dev/design/generics-implementation-stenciling) Generics implementation - Stenciling
- [design/generics-implementation-dictionaries](https://go.dev/design/generics-implementation-dictionaries) Generics implementation - Dictionaries
- [issue/43651](https://go.dev/issue/43651) proposal: spec: add generic programming using type parameters
- [design/type-parameters3](https://go.dev/design/43651-type-parameters) Type Parameters Proposal
- [issue/45346](https://go.dev/issue/45346) proposal: spec: generics: use type sets to remove type keyword in constraints
- [issue/46477](https://go.dev/issue/46477) proposal: spec: generics: type parameters on aliases
- Standard packages using generics
  + [issue/45458](https://go.dev/issue/45458) proposal: constraints: new package to define standard type parameter constraints
  + [discuss/47319](https://github.com/golang/go/discussions/47319) proposal: constraints: new package to define standard type parameter constraints (discussion)
  + [issue/45955](https://go.dev/issue/45955) proposal: slices: new package to provide generic slice functions
  + [discuss/47203](https://github.com/golang/go/discussions/47203) proposal: slices: new package to provide generic slice functions (discussion)
  + [discuss/47331](https://github.com/golang/go/discussions/47331) proposal: container/set: new package to provide a generic set type (discussion)
  + [discuss/47330](https://github.com/golang/go/discussions/47330) proposal: maps: new package to provide generic map functions (discussion)
  + [issue/47657](https://go.dev/issue/47657) proposal: sync, sync/atomic: add PoolOf, MapOf, ValueOf

[Back To Top](#top)

## Compiler Toolchain

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
- [talk/gccgo](https://www.youtube.com/watch?v=U0w9eFunkX4) Brief overview of gccgo, "the other" Go compiler. Aug 6, 2015.
- [issue/28262](https://go.dev/issue/28262) cmd/compile: feedback-guided optimization

[Back To Top](#top)

### Linker

The Go Linker was written by Ken Thompson. Russ led a few more overhaul in Go 1.3. Austin led a rework to the linker together with Keith, Than,
Cheery and many other brilliant brains, which was landed
in Go 1.15 and Go 1.16.

- [design/go13linker](https://go.dev/s/go13linker) Russ Cox. Go 1.3 Linker Overhaul. Nov 2013.
- [design/go116linker](https://go.dev/s/better-linker) Austin Clements. Building a better Go linker. 2019-09-12.

[Back To Top](#top)

### Debugger

- [doc/go13heapdump](https://go.dev/s/go13heapdump) heapdump13
- [doc/go14heapdump](https://github.com/golang/go/wiki/heapdump14) heapdump14
- [doc/go15heapdump](https://github.com/golang/go/wiki/heapdump15-through-heapdump17) heapdump15 through heapdump17
- [design/heap-viewer](https://go.dev/design/16410-heap-viewer) Michael Matloob. Proposal: Go Heap Dump Viewer. 20 July 2016
  + [issue/16410](https://go.dev/issue/16410) x/tools/cmd/heapdump: create a heap dump viewer
- [design/profiler-labels](https://go.dev/design/17280-profile-labels) Michael Matloob. Proposal: Support for pprof profiler labels. 15 May 2017.
  + [issue/17280](https://go.dev/issue/17280) pprof: add support for profiler labels

[Back To Top](#top)

### Race Detector

- [issue/42598](https://go.dev/issue/42598) runtime: apparent false-positive race report for a buffered channel after CL 220419

[Back To Top](#top)

### Tracer

- [design/go15trace](https://go.dev/s/go15trace) Dmitry Vyukov. Go Execution Tracer. Oct 2014
- [design/tracefmt](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@. Trace Event Format. October 2016.

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

[Back To Top](#top)

### Modules

- [design/go-dep](https://docs.google.com/document/d/18tNd8r5DV0yluCR7tPvkMTsWD_lYcRO7NhpNSDymRr8) Go Packaging Proposal Process
- [design/go-dep2](https://docs.google.com/document/d/1qnmjwfMmvSCDaY4jxPmLAccaaUI5FfySNE90gB0pTKQ/edit) Dependency Management Tool
  - [doc/go-dep](6https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) Sam Boyer. The Saga of Go Dependency Management. Dec 13, 201
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
- [design/lazy-gomod](https://go.dev/design/36460-lazy-module-loading) Bryan C. Mills. Proposal: Lazy Module Loading. 2020-02-20
- [issue/45713](https://go.dev/issue/45713) proposal: cmd/go: add a workspace mode
  + [design/workspace](https://go.dev/design/45713-workspace) Proposal: Multi-Module Workspaces in cmd/go

[Back To Top](#top)

### gopls

- [design/gopls-workspace](https://go.dev/design/37720-gopls-workspaces) Heschi Kreinick, Rebecca Stambler. Proposal: Multi-project gopls workspaces. Mar 9, 2020.

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

<!-- - Tool chain, benchseries/benchstat -->

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
- [issue/44313](https://go.dev/issue/44313) runtime: stopped Ms can't become dedicated or fractional GC workers



[Back To Top](#top)

### Execution Stack

- [design/contigstack](https://go.dev/s/contigstacks) Contiguous stacks
- [issue/17007](https://go.dev/issue/17007) runtime: fatal error: bad g->status in ready
- [issue/18138](https://go.dev/issue/18138) runtime: new goroutines can spend excessive time in morestack
  + [design/predict-stack-size](https://docs.google.com/document/d/1YDlGIdVTPnmUiTAavlZxBI1d9pwGQgZT7IKFKlIXohQ/edit#) Keith Randall. Determining a good starting stack size. 2021/08/18.
  + [cl/341990](https://go.dev/cl/341990) runtime: predict stack sizing
  + [cl/345889](https://go.dev/cl/345889) runtime: measure stack usage; start stacks larger if needed
- [issue/26061](https://go.dev/issue/26061) runtime: g0 stack.lo is sometimes too low

[Back To Top](#top)

### Memory Allocator

A quick history about the Go's memory allocator: Russ Cox first implements
the memory allocator based on `tcmalloc` for Go 1, `mcache` is cached on M.
Then he revised the allocator to allow user code to use 16GB memory and later allows 128GB. However, the allocator (including scavenger) was
suffered from massive lock contentions and does not scale. After Dmitry's scalable runtime
scheduler, the allocator can allocate directly from P (with much less)
lock contentions. In the meantime, the scavenger is migrated from an independent
thread into the system monitor thread. Now, Michael is actively working on
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

[Back To Top](#top)

### Garbage Collector

- [paper/on-the-fly-gc](https://doi.org/10.1145/359642.359655) Edsger W. Dijkstra, Leslie Lamport, A. J. Martin, C. S. Scholten, and E. F. M. Steffens. 1978. On-the-fly garbage collection: An exercise in cooperation. Commun. ACM 21, 11 (November 1978), 966–975.
- [paper/yuasa-barrier](https://doi.org/10.1016/0164-1212(90)90084-Y) T. Yuasa. 1990. Real-time garbage collection on general-purpose machines. J. Syst. Softw. 11, 3 (March 1990), 181-198.
- [design/go13gc](https://docs.google.com/document/d/1v4Oqa0WwHunqlb8C3ObL_uNQw3DfSY-ztoA-4wWbKcg/pub) Dmitry Vyukov. Simpler and faster GC for Go. July 16, 2014
  + [cl/106260045](https://codereview.appspot.com/106260045) runtime: simpler and faster GC
- [design/go14gc](https://go.dev/s/go14gc) Richard L. Hudson. Go 1.4+ Garbage Collection (GC) Plan and Roadmap. August 6, 2014.
- [design/go15gcpacing](https://go.dev/s/go15gcpacing) Austin Clements. Go 1.5 concurrent garbage collector pacing. 2015-03-10.
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


[Back To Top](#top)

### Statistics

- [issue/16843](https://go.dev/issue/16843) runtime: mechanism for monitoring heap size
  + [cl/setmaxheap](https://go-review.googlesource.com/c/go/+/46751/) Austin Clements. runtime/debug: add SetMaxHeap API. Jun 26 2017.
- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek. Proposal: API for unstable runtime metrics. Mar 18, 2020.
- [issue/19812](https://go.dev/issue/19812) runtime: cannot ReadMemStats during GC
- [issue/38712](https://go.dev/issue/38712) runtime: TestMemStats is flaky
- [issue/40459](https://go.dev/issue/40459) runtime: ReadMemStats called in a loop may prevent GC

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

[Back To Top](#top)

### ABI

- [design/cgo-pointers](https://go.dev/design/12416-cgo-pointers) Ian Lance Taylor. Proposal: Rules for passing pointers between Go and C. October, 2015
  + [issue/12416](https://go.dev/issue/12416) cmd/cgo: specify rules for passing pointers between Go and C
- [design/internal-abi](https://go.dev/design/27539-internal-abi) Austin Clements. Proposal: Create an undefined internal calling convention. 2019-01-14.
  + [issue/27539](https://go.dev/issue/27539) proposal: runtime: make the ABI undefined as a step toward changing it.
  + [issue/31193](https://go.dev/issue/31193) cmd/compile, runtime: reduce function prologue overhead
- [design/register-call](https://go.dev/design/40724-register-calling) Austin Clements, with input from Cherry Zhang, Michael Knyszek, Martin Möhrmann, Michael Pratt, David Chase, Keith Randall, Dan Scales, and Ian Lance Taylor. Proposal: Register-based Go calling convention. 2020-08-10.
- [issue/18597](https://go.dev/issue/18597) proposal: cmd/compile: define register-based calling convention
- [issue/40724](https://go.dev/issue/40724) proposal: switch to a register-based calling convention for Go functions
  + [cl/266638](https://go.dev/cl/266638) reflect,runtime: use internal ABI for selected ASM routines, attempt 2
  + [cl/259445](https://go.dev/cl/259445) cmd/compile,cmd/link: initial support for ABI wrappers
- [design/internal-abi-spec](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md) Go internal ABI specification.

[Back To Top](#top)

### Misc

- [issue/20135](https://go.dev/issue/20135) runtime: shrink map as elements are deleted

## Standard Library

### syscall

- [design/go14syscall](https://go.dev/s/go1.4-syscall) The syscall package.

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


[Back To Top](#top)

### go/*

- [doc/gotypes](https://go.dev/s/types-tutorial) Alan Donovan. go/types: The Go Type Checker.
- [talk/gotypes](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view) Alan Donovan. Using go/types for
Code Comprehension and Refactoring Tools. October 2, 2015.
- [design/modular-interface](https://docs.google.com/document/d/1-azPLXaLgTCKeKDNg0HVMq2ovMlD-e7n1ZHzZVzOlJk/edit) Alan Donovan. Proposal: a common interface for modular static analyses for Go. 9 Sep 2018
  + [cl/134935](https://go.dev/cl/134935) go/analysis: a new API for analysis tools

[Back To Top](#top)

### sync

- [design/percpu-sharded](https://go.dev/design/18802-percpu-sharded) Sanjay Menakuru. Proposal: percpu.Sharded, an API for reducing cache contention. Jun 12, 2018.
  + [issue/18802](https://go.dev/issue/18802) proposal: sync: support for sharded values
- [issue/37142](https://go.dev/issue/37142) sync: shrink types in sync package

[Back To Top](#top)

#### Map

- [issue/21031](https://go.dev/issue/21031) sync: reduce pointer overhead in Map
- [issue/21032](https://go.dev/issue/21032) sync: reduce (*Map).Load penalty for Stores with new keys
- [issue/21035](https://go.dev/issue/21035) sync: reduce contention between Map operations with new-but-disjoint keys
- [issue/37033](https://go.dev/issue/37033) runtime: provide centralized facility for managing (c)go pointer handles

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
- [issue/42564](https://go.dev/issue/42564) context: cancelCtx exclusive lock causes extreme contention

[Back To Top](#top)

### encoding

- [design/go12encoding](https://go.dev/s/go12encoding) Russ Cox. Go 1.2 encoding.TextMarshaler and TextUnmarshaler. July 2013.
- [design/go12xml](https://docs.google.com/document/d/1G-AMeUPoyTnbZDkuCJA89DjJwTjdWFctIZ_N9NgA9RM/pub) Russ Cox. Go 1.2 xml.Marshaler and Unmarshaler. July 2013.
- [design/natural-xml](https://go.dev/design/13504-natural-xml) Sam Whited. Proposal: Natural XML. 2016-09-27
  + [issue/13504](https://go.dev/issue/13504) encoding/xml: add generic representation of XML data
- [design/zip](https://go.dev/design/14386-zip-package-archives) Russ Cox. Proposal: Zip-based Go package archives. February 2016
  + [issue/14386](https://go.dev/issue/14386) proposal: use zip format inside .a and .o files
- [design/xmlstream](https://go.dev/design/19480-xml-stream) Sam Whited. Proposal: XML Stream. 2017-03-09
  + [issue/19480](https://go.dev/issue/19480) encoding/xml: add decode from TokenReader, to enable stream transformers
- [design/raw-xml](https://go.dev/design/26756-rawxml-token) Sam Whited. Proposal: Raw XML Token. 2018-09-01
  + [issue/26756](https://go.dev/issue/26756) proposal: encoding/xml: add RawXML token

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

<!--
TODO: read all of these!
These issues are discussion the current performance issue that exist in the current implementation.

- [issue/8055](https://go.dev/issue/8055) image: decode / resize into an existing buffer
- [issue/11793](https://go.dev/issue/11793) image/color: NRGBA(64).RGBA() optimization
- [issue/15759](https://go.dev/issue/15759) image: optimize Image.At().RGBA()
- [issue/20851](https://go.dev/issue/20851) image: Decode drops interfaces
- [issue/24499](https://go.dev/issue/24499) image/jpeg: Decode is slow


- [issue/22535](https://go.dev/issue/22535) image: support LJPEG
- [issue/18098](https://go.dev/issue/18098) proposal: add Validate functions to image/jpeg, image/png etc.
- [issue/2362](https://go.dev/issue/2362) image/jpeg: chroma downsampling ratios are restricted
- [issue/4341](https://go.dev/issue/4341) image/jpeg: correct for EXIF orientation?
- [issue/10447](https://go.dev/issue/10447) image/jpeg: add options to partially decode or tolerantly decode invalid images?
- [issue/12202](https://go.dev/issue/12202) image/jpeg: specify APP1 segment for outputting EXIF data in jpeg.Encode()?
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
- [issue/41142](https://go.dev/issue/41142) image/gif: Decode reads the entire animated gif image, even though it returns only the first frame (while DecodeAll exists to read and return all frames)

x/image:

- [issue/40173](https://go.dev/issue/40173) x/image: WebP decode contrast issue
- [issue/39705](https://go.dev/issue/39705) x/image: CCITT reader EOF error for tiff image
- [issue/25657](https://go.dev/issue/25657) x/image: vector.go rasterizer shifts alpha mask and is slow when target is offset and small relative image size
- [issue/39900](https://go.dev/issue/39900) x/image/tiff: Missing raw stream read/write
- [issue/38252](https://go.dev/issue/38252) x/image/tiff: add 32bit float grayscale support
- [issue/36121](https://go.dev/issue/36121) x/image/tiff: grayscale tiled images are not decoded correctly
- [issue/33708](https://go.dev/issue/33708) x/image/tiff: sony .arw files decode as a 0x0 image.Gray
- [issue/30827](https://go.dev/issue/30827) x/image/tiff: unexpected EOF
- [issue/26450](https://go.dev/issue/26450) x/image/tiff: implement a generic tiff parser
- [issue/26360](https://go.dev/issue/26360) x/image/tiff: compressed tiffs are invalid (at least on Mac OS X)
- [issue/23115](https://go.dev/issue/23115) x/image/tiff: no support for cJPEG or cJPEGOld
- [issue/20742](https://go.dev/issue/20742) x/image/tiff: package does not support image resolution
- [issue/11413](https://go.dev/issue/11413) x/image/tiff: invalid format: wrong number of samples for RGB
- [issue/11389](https://go.dev/issue/11389) x/image/tiff: excessive memory consumption
- [issue/11386](https://go.dev/issue/11386) x/image/tiff: index out of range
- [issue/19672](https://go.dev/issue/19672) x/image/webp: issue with colors contrast when converting to jpeg/png
- [issue/30902](https://go.dev/issue/30902) x/image/riff: Implement write functionality
- [issue/29711](https://go.dev/issue/29711) x/image/bmp: support 1-bit format
- [issue/37532](https://go.dev/issue/37532) x/image/font/gofont: Go Mono font readability for users
- [issue/37441](https://go.dev/issue/37441) x/image/font/gofont: Go fonts not representative of OpenType state of the art
- [issue/30699](https://go.dev/issue/30699) x/image/font/sfnt: read more glyph metrics
- [issue/28932](https://go.dev/issue/28932) x/image/font: wrong rendering of intersecting paths
- [issue/28380](https://go.dev/issue/28380) x/image/font/sfnt: support trimmed table mapping cmap format
- [issue/27281](https://go.dev/issue/27281) x/image/font: rendering texts in Arabic
- [issue/23497](https://go.dev/issue/23497) x/image/font/gofont/gomedium: wrong shape for "l" letter
- [issue/22451](https://go.dev/issue/22451) x/image/font/sfnt: implement font.Face
- [issue/20208](https://go.dev/issue/20208) x/image/font: Tool for running Unicode’s text rendering tests
- [issue/14436](https://go.dev/issue/14436) x/image/font: make it easier to measure a string's bounds and draw it in a bounding box
- [issue/33990](https://go.dev/issue/33990) x/image/font/sfnt: GlyphName returns empty string on OpenType font
- [issue/16904](https://go.dev/issue/16904) proposal: x/image packages to render TrueType fonts -->

[Back To Top](#top)

### Mobile

- [design/go14android](https://go.dev/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/gobind](https://go.dev/s/gobind) David Crawshaw. Binding Go and Java. July 2014

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
- ["C" Programming Language: Brian Kernighan - Computerphile](https://youtu.be/de2Hsvxaf8M)
- [Debugging performance issues in Go* programs By Dmitriy Vyukov](https://software.intel.com/content/www/us/en/develop/blogs/debugging-performance-issues-in-go-programs.html)
- [Interesting papers I'd like to implement (or at least have implementations of)](https://github.com/dgryski/interesting-papers)
- [golang/wiki/ResearchPapers](https://github.com/golang/go/wiki/ResearchPapers)

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

It is also important to thank the continuing, inspiring discussion and sharing with the TalkGo community core members [qcrao](https://github.com/qcrao), and [eddycjy](https://github.com/eddycjy).

The document would not be organized without all of the supports from them.

![](https://changkun.de/urlstat?mode=github&repo=golang-design/history)

[Back To Top](#top)

## License

[golang.design/history](https://github.com/golang-design/history) | CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)
