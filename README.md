# A Documentary of Go

**Author**: Changkun Ou <_[changkun.de](https://changkun.de)_>

This document collects many interesting (publiclly observable) issues,
discussions, proposals, CLs, and talks from the Go development process,
which intents to offer a comprehensive reference of the Go history.

## Disclaimer

- Most of the text are written as my _personal_ understanding based on public sources
- **Factual and typo errors may occur.**
Referring to the original link if some text conflicts to your understanding
- [PR](https://github.com/changkun/go-history/pulls)s are very welcome for new content, bug and typo fixes
- Use [Issues](https://github.com/changkun/go-history) for discussions

## Sources

There are many sources for digging the documents that relate to Go's
historical design. There are some of the official sources:

- [golang.org/doc](https://golang.org/doc)
- [blog.golang.org](https://blog.golang.org)
- [golang.org/pkg](https://golang.org/pkg)
- [dev.golang.org](https://dev.golang.org)
- [github.com/golang/go](https://github.com/golang/go)
- [github.com/golang/talks](https://github.com/golang/talks)
- [github.com/golang/proposal](https://github.com/golang/proposal)
- [github.com/golang/go/wiki](https://github.com/golang/go/wiki)
- [go-review.googlesource.com](https://go-review.googlesource.com)
- [groups.google.com/g/golang-nuts](https://groups.google.com/g/golang-nuts)
- [groups.google.com/g/golang-dev](https://groups.google.com/g/golang-dev)
- [groups.google.com/g/golang-tools](https://groups.google.com/g/golang-tools)

## Committers

Go is a big project that driven by a tiny group of people and massive
crowd of wisdom from community. Here are some core committers to
the project that you might interest in follow their excellent work.

By listening to the talks held by these people, you could learn more about
their oral history and fun stories behind the great work.

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
[Plan 9](http://plan9.bell-labs.com/plan9) project. Russ did many
fundamental work for the early Go compiler, runtime, as well as the leap of
Go 1.5 bootstrap.
Now, Russ is the tech leader of the Go team.

- Rob Pike. [Website](https://commandcenter.blogspot.com/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike). (Retired)
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
  + [talk/rob2019](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer. Sep 10, 2019.
  + [talk/rob2020](https://evrone.com/rob-pike-interview) A Rob Pike Interview. (Date Unclear) 2020.

- Robert Griesemer. [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
  + [talk/robert2012a](https://www.youtube.com/watch?v=on5DeUyWDqI) E2E: Erik Meijer and Robert Griesemer. Going Go. Lang.NEXT. 2012.
  + [talk/robert2012b](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Go-In-Three-Easy-Pieces) Go In Three Easy Pieces. Mar 19, 2012.
  + [talk/robert2012c](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Panel-Native-Languages) Lang.NEXT 2012 Expert Panel: Native Languages. Apr 11, 2012.
  + [talk/robert2015](https://www.youtube.com/watch?v=0ReKdcpNyQg) The Evolution of Go. GopherCon 2015. Jul 28, 2015.
  + [talk/robert2016a](https://www.youtube.com/watch?v=t-w6MyI2qlU) Lightning Talk: Alias Declarations for Gom: A proposal. GopherCon 2016. Oct 9, 2016.
  + [talk/robert2016b](https://www.youtube.com/watch?v=vLxX3yZmw5Q) Prototype your design!. dotGo 2016. Nov 29, 2016.
  + [talk/robert2017](https://www.youtube.com/watch?v=KPk1UPihWtY) Opening Keynote: Exporting Go. GopherCon SG 2017. May 29, 2017.
  + [talk/robert2017](https://www.youtube.com/watch?v=UmwJsQTSEP8) A brief overview of Go. Aug 28. 2017.
  + [talk/robert2019](https://www.youtube.com/watch?v=i0zzChzk8KE) Go is 10! Now What?. Gopherpalooza 2019. Dec 2, 2019.
  + [talk/robert2020](https://changelog.com/gotime/140) The latest on Generics with Robert Griesemer and Ian Lance Taylor. Jul 21, 2020.

- Ken Thompson (Retired)
  + [talk/ken1982a](https://www.youtube.com/watch?v=tc4ROCJYbm0) The UNIX System: Making Computers More Productive. 1982.
  + [talk/ken1982b](https://www.youtube.com/watch?v=XvDZLjaCJuw) UNIX: Making Computers Easier To Use.
  + [talk/ken1982c](https://www.youtube.com/watch?v=JoVQTPbD6UY) Ken Thompson and Dennis Ritchie Explain UNIX (Bell Labs).
  + [talk/ken1983](https://www.youtube.com/watch?v=LXZ1OL2U3lY) Ken Thompson and Dennis Ritchie. National Medal of Technology Awards. 1996.
  + [talk/ken2013](https://www.youtube.com/watch?v=dsMKJKTOte0) Systems Architecture, Design, Engineering, and Verification. Jan 17, 2013.
  + [talk/ken2019a](https://www.youtube.com/watch?v=g3jOJfrOknA) The Thompson and Ritchie Story. Feb 18, 2019.
  + [talk/ken2019b](https://www.youtube.com/watch?v=EY6q5dv_B-o) Brian Kernighan interviews Ken Thompson. VCF East 2019. May 4, 2019.


- Ian Lance Taylor. [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor), [Quora](https://www.quora.com/profile/Ian-Lance-Taylor)
  + [talk/ian2007](https://www.youtube.com/watch?v=gc78olyguqA) GCC: Current Topics and Future Directions. February 27, 2007.
  + [talk/ian2018](https://www.youtube.com/watch?v=LqKOY_pH8u0) Transition to Go 2. Gopherpalooza 2018. Oct 24, 2018
  + [talk/ian2019a](https://www.youtube.com/watch?v=WzgLqE-3IhY) Generics in Go. GopherCon 2019. Aug 27, 2019
  - [talk/ian2019b](https://changelog.com/gotime/98) Generics in Go. Aug 27, 2019.
  + [talk/ian2020](https://www.youtube.com/watch?v=yoZ05GG8aLs) Go with Ian Lance Taylor. CppCast. Aug 9, 2020.

- Russ Cox. [Website](https://swtch.com/~rsc/), [Blog](https://research.swtch.com/), [GitHub](https://github.com/rsc), [Twitter](https://twitter.com/_rsc), [Reddit](https://old.reddit.com/user/rsc)
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
  + [talk/russ2020a](https://www.youtube.com/watch?v=AgR_mdC4Rs4) `go:build` design draft. Jun 30, 2020.
  + [talk/russ2020b](https://www.youtube.com/watch?v=yx7lmuwUNv8) `io/fs` draft design. Jul 21, 2020.
  + [talk/russ2020c](https://www.youtube.com/watch?v=rmS-oWcBZaI) `//go:embed` draft design. Jul 21, 2020.

### Compiler/Runtime Team

Dmitry is not from the Go team but on the dynamic testing tools team from Google.
He wrote the scalable goroutine scheduler, many other performance improvements,
synchronization primitives, race detector, and blocking profiler that
related to the Go runtime.
Austin was an intern at Google who worked on the Go project in the early days
while pursuing a Ph. D. Later, he joined the Go team after his academic career
and work together with Rick for the Go's concurrent GC. He also the worked on the current preemptive scheduler and linker.
Now, he is leading the Compiler/Runtime team for Go.
Keith and David together focus on the Go's compiler backend,
notably the current SSA backend. Michael is a recent new comer to the Go team,
his work mainly in the runtime memory system such as the refactoring of memory allocator and runtime metrics.

- Dmitry Vyukov (Дмитрий Вьюков). [Website](http://www.1024cores.net/), [GitHub](https://github.com/dvyukov), [Twitter](https://twitter.com/dvyukov)
  + [talk/dmitry2014](https://www.youtube.com/watch?v=QEhpLb0UCfE) Writing a functional, reliable and fast web application in Go. Sep 25, 2014.
  + [talk/dmitry2015a](https://www.youtube.com/watch?v=Ef7TtSZlmlk) Chat with Dmitry Vyukov on go-fuzz, golang and IT security. Jul 3, 2015
  + [talk/dmitry2015b](https://www.youtube.com/watch?v=a9xrxRsIbSU) Go Dynamic Tools. GopherCon 2015. Jul 28, 2015.
  + [talk/dmitry2016](https://www.youtube.com/watch?v=9cpN2r22sLE) Dmitry Vyukov Interview. Jun 1, 2016.
  + [talk/dmitry2017](https://www.youtube.com/watch?v=FD30Qzd6ylk) Fuzzing: The New Unit Testing. C++ Russia 2017. May 10, 2017.
  + [talk/dmitry2018a](https://www.youtube.com/watch?v=EJVp13f_aIs) Fuzzing: new unit testing. GopherCon Russia. Apr 2, 2018.
  + [talk/dmitry2018b](https://www.youtube.com/watch?v=qrBVXxZDVQY) Syzbot and the Tale of Thousand Kernel Bugs. Sep 1, 2018.
  + [talk/dmitry2019](https://www.youtube.com/watch?v=-K11rY57K7k) Go scheduler: Implementing language with lightweight concurrency. Oct 14, 2019.
  + [talk/dmitry2020](https://www.youtube.com/watch?v=YwX4UyXnhz0) syzkaller: Adventures in Continuous Coverage-guided Kernel Fuzzing. BlueHat IL 2020. Feb 13, 2020.

- Austin Clements. [GitHub](https://github.com/aclements), [Scholar](https://scholar.google.com/citations?user=MKDtxN4AAAAJ)

- Richard L. Hudson. [GitHub](https://github.com/RLH) (Retired)
  + [talk/rick2015](https://www.youtube.com/watch?v=aiv1JOfMjm0) Go GC: Solving the Latency Problem. GopherCon 2015. Jul 8, 2015.
  + [talk/rick2015b](https://www.infoq.com/interviews/hudson-go-gc/) Rick Hudson on Garbage Collection in Go. Dec 21, 2015

- Keith Randall. [GitHub](https://github.com/randall77)
  + [talk/keith2016](https://www.youtube.com/watch?v=Tl7mi9QmLns) Inside the Map Implementation. GopherCon 2016. Jul 12, 2016.
  + [talk/keith2017](https://www.youtube.com/watch?v=uTMvKVma5ms) Generating Better Machine Code with SSA. GopherCon 2017. Jul 24, 2017.

- David Chase. [Website](https://dr2chase.wordpress.com/), [GitHub](https://github.com/dr2chase), [Twitter](https://twitter.com/dr2chase), [Scholar](https://dblp.org/pid/51/3488.html)
  + [talk/david2017](https://changelog.com/gotime/52) All About The Go Compiler. Jul 20, 2017.

- Dan Scales.

- Michael Knyszek. [Website](https://www.ocf.berkeley.edu/~mknyszek/), [GitHub](https://github.com/mknyszek)

- Than McIntosh. [GitHub](https://github.com/thanm)

- Cherry Zhang. [GitHub](https://github.com/cherrymui)

### Library/Tools/Security/Community

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

- Daniel Martí. [Website](https://mvdan.cc/), [GitHub](https://github.com/mvdan), [Twitter](https://twitter.com/mvdan_)
- Nigel Tao. [GitHub](https://github.com/nigeltao), Twitter
- Filippo Valsorda. GitHub,
- Michael Matloob. [GitHub](https://github.com/matloob), [Twitter](https://twitter.com/matloob)
- Jay Conrod. [Website](https://jayconrod.com/), [Twitter](https://twitter.com/jayconrod)
- Dave Cheney. [Website](https://dave.cheney.net/), [GitHub](https://github.com/davecheney), [Twitter](https://twitter.com/davecheney)
- Sam Boyer. [GitHub](https://github.com/sdboyer), [Twitter](https://twitter.com/sdboyer)

> more people and talks should be added...

### Group Interviews

- [talk/goteam2012](https://www.youtube.com/watch?v=sln-gJaURzk) Meet the Go Team. Google I/O 2012. Jul 2, 2012.
- [talk/goteam2013](https://www.youtube.com/watch?v=p9VUCp98ay4) Fireside Chat with the Go Team. Google I/O 2013. May 18, 2013.
- [talk/goteam2014](https://www.youtube.com/watch?v=u-kkf76TDHE) Inside the Gophers Studio with Blake Mizerany. GopherCon 2014. May 21, 2014
- [talk/goteam2019](https://www.youtube.com/watch?v=3yghHvvZQmA) Meet the Authors: Go Language. Cloud Next '19. Apr 10, 2019.

## Timeline

Timeline helps you identify the point in time about a document that reflected in Go versions.

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

The historical release notes may helpful for general informations:

- [doc/go1release](https://golang.org/doc/devel/release.html) Go Release History
- [doc/go1prerelease](https://golang.org/doc/devel/pre_go1.html) Pre-Go 1 Release History
- [doc/go0release](https://golang.org/doc/devel/weekly.html) Weekly Release History (Before Go 1)

## Language Design

### Misc

- [design/go0initial](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) Rob Pike, Robert Griesemer, Ken Thompson. The Go Annotated Specification. Mar 3, 2008.
- [design/go0semicolon](https://golang.org/s/semicolon-proposal) Rob Pike. Semicolons in Go. Dec 10, 2009.
- [design/go11func](https://golang.org/s/go11func) Russ Cox. Go 1.1 Function Calls. February 2013
- [design/go11return](https://golang.org/s/go11return) Russ Cox. Go “Return at End of Function” Requirements. March 2013
- [design/go12nil](https://golang.org/s/go12nil) Russ Cox. Go 1.2 Field Selectors and Nil Checks. July 2013.
- [doc/go13todo](https://golang.org/s/go13todo) Go 1.3 “To Do” List
- [doc/goatgoogle](https://talks.golang.org/2012/splash.article#TOC_12.) Rob Pike. Go at Google - Language Semantics. October 25, 2012.
- [doc/makego](https://talks.golang.org/2015/how-go-was-made.slide)  Andrew Gerrand. How Go was Made. 9 July 2015.
- [discuss/go1preview](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) Russ Cox. Go 1 Preview.
- [design/overlapping-interfaces](https://golang.org/design/6977-overlapping-interfaces) Robert Griesemer. Proposal: Permit embedding of interfaces with overlapping method sets. 2019-10-16
  + [issue/6977](https://golang.org/issue/6977) spec: allow embedding overlapping interfaces
- [design/struct-conversion](https://golang.org/design/16085-conversions-ignore-tags) Robert Griesemer. Proposal: Ignore tags in struct type conversions. June 16, 2016.
  + [issue/16085](https://golang.org/issue/16085) Proposal: Ignore tags in struct type conversions
- [design/go2trans](https://golang.org/design/28221-go2-transitions) Ian Lance Taylor. Proposal: Go 2 transition. October 15, 2018
- [design/signed-int-shift](https://golang.org/design/19113-signed-shift-counts) Robert Griesemer. Proposal: Permit Signed Integers as Shift Counts for Go 2. January 17, 2019
  + [issue/19113](https://golang.org/issue/19113) proposal: spec: allow signed shift counts
- [design/number-literal](https://golang.org/design/19308-number-literals) Russ Cox, Robert Griesemer. Proposal: Go 2 Number Literal Changes. March 6, 2019
  + [issue/12711](golang.org/issue/12711) proposal: different octal base literal representation
  + [issue/19308](golang.org/issue/19308) proposal: spec: binary integer literals
  + [issue/28493](golang.org/issue/28493) proposal: permit blank (_) separator in integer number literals
  + [issue/29008](golang.org/issue/29008) proposal: Go 2: hexadecimal floats
- [issue/33502](https://golang.org/issue/33502) proposal: review meeting minutes
- [issue/33892](https://golang.org/issue/33892) proposal: Go 2 review meeting minutes

### Slice (1.2)

- [design/read-only-slices](https://docs.google.com/document/d/1UKu_do3FRvfeN5Bb1RxLohV-zBOJWTzX0E8ZU1bkqX0/edit#heading=h.2wzvdd6vdi83) Brad Fitzpatrick. Read-only slices. May 13, 2013
- [design/read-only-slices-russ](https://docs.google.com/document/d/1-NzIYu0qnnsshMBpMPmuO21qd8unlimHgKjRD9qwp2A/edit) Russ Cox. Evaluation of read-only slices. May 2013.
- [design/go12slice](https://golang.org/s/go12slice) Russ Cox. Go Slice with Cap. June 2013.
- [design/multidim-slice](https://golang.org/design/6282-table-data) Brendan Tracey. Proposal: Multi-dimensional slices. November 17th, 2016.

### Package Management (1.4, 1.5, 1.7)

- [design/go14internal](https://golang.org/s/go14internal) Russ Cox. Go 1.4 “Internal” Packages. June 2014.
- [design/go14nopkg](https://golang.org/s/go14nopkg) Russ Cox. Go 1.4 src/pkg → src. June 2014.
- [design/go14customimport](https://golang.org/s/go14customimport) Russ Cox. Go 1.4 Custom Import Path Checking. June 2014.
- [design/go15vendor](https://golang.org/s/go15vendor) Russ Cox. Go 1.5 Vendor Experiment. July 2015
- [design/go17binarypkg](https://golang.org/design/2775-binary-only-packages) Russ Cox. Proposal: Binary-Only Packages. April 24, 2016
  + [issue/2775](golang.org/issue/2775) cmd/go: work when binaries are available but source is missing

### Type alias (1.9)

- [design/type-alias](https://golang.org/design/18130-type-alias) Russ Cox, Robert Griesemer. Proposal: Type Aliases. December 16, 2016
  + [talk/type-alias](https://www.youtube.com/watch?v=t-w6MyI2qlU) GopherCon 2016 - Lightning Talk: Robert Griesemer - Alias Declarations for Go, A proposal. Oct 9, 2016
  + [issue/16339](https://golang.org/issue/16339) proposal: Alias declarations for Go
  + [issue/18130](https://golang.org/issue/18130) all: support gradual code repair while moving a type between packages
- [talk/refactor-video](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Russ Cox. Codebase Refactoring (with help from Go). GothamGo 2016. November 18, 2016.
  + [doc/refactor](https://talks.golang.org/2016/refactor.article) Russ Cox. Codebase Refactoring (with help from Go).

### Defer (1.13)

- [design/open-defer](https://golang.org/design/34481-opencoded-defers) Dan Scales, Keith Randall, and Austin Clements. Proposal: Low-cost defers through inline code, and extra funcdata to manage the panic case. 2019-09-23
  + [issue/6980](https://golang.org/issue/6980) cmd/compile: allocate some defers in stack frames
  + [issue/14939](https://golang.org/issue/14939) runtime: defer is slow #14939
- Unsolved `defer recover()` edge cases.
  - [issue/10458](https://golang.org/issue/10458) spec: panicking corner-case semantics
  - [issue/23531](https://golang.org/issue/23531) spec: recover() in nested defer
  - [issue/26275](https://golang.org/issue/26275) runtime: document the behaviour of Caller and Callers when used in deferred functions
  - [issue/34530](https://golang.org/issue/34530) spec: clarify when calling recover stops a panic
  - [cl/189377](https://golang.org/cl/189377) spec: specify the behavior of recover and recursive panics in more detail

### Error values (1.13)

- [doc/err2011](https://blog.golang.org/error-handling-and-go) Andrew Gerrand. Error handling in Go. July 2011.
- [doc/err-values](https://blog.golang.org/errors-are-values) Rob Pike. Errors are values. January 2015.
- [doc/err-philosophy](https://dave.cheney.net/paste/gocon-spring-2016.pdf) Dave Cheney. My philosophy for error handling. April 2016.
- [doc/err-gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) Dave Cheney. Don’t just check errors, handle them gracefully. April 2016.
- [doc/err-stacktrace](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package) Dave Cheney. Stack traces and the errors package. June, 12 2016.
- [doc/err-upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) Rob Pike. Error handling in Upspin. December 06, 2017.
- [doc/err-work](https://blog.golang.org/go1.13-errors) Damien Neil and Jonathan Amsterdam. Working with Errors in Go 1.13. October 17, 2019.
- [design/err-handling-overview](https://golang.org/design/go2draft-error-handling-overview) Russ Cox. Error Handling — Problem Overview.
- [doc/err-value-faq](https://github.com/golang/go/wiki/ErrorValueFAQ) Jonathan Amsterdam and Bryan C. Mills. Error Values: Frequently Asked Questions. August 2019.
- [design/err-handle-check](https://golang.org/design/go2draft-error-handling) Marcel van Lohuizen. Error Handling — Draft Design. August 27, 2018.
- [design/err-try](https://golang.org/design/32437-try-builtin) Robert Griesemer. Proposal: A built-in Go error check function, "try". 2019-06-12
  - [issue/32437](https://golang.org/issue/32437#issuecomment-512035919) Proposal: A built-in Go error check function, "try". Decision response.
- [design/err-inspect](https://golang.org/design/go2draft-error-inspection) Jonathan Amsterdam, Damien Neil. Error Inspection — Draft Design. August 27, 2018.
- [design/err-values-overview](https://golang.org/design/go2draft-error-values-overview) Russ Cox. Error Values — Problem Overview. August 27, 2018.
- [design/error-values](https://golang.org/design/29934-error-values) Jonathan Amsterdam, Russ Cox, Marcel van Lohuizen, Damien Neil. Proposal: Go 2 Error Inspection. January 25, 2019
  + [issue/29934](https://golang.org/issue/29934) Jonathan Amsterdam. proposal: Go 2 error values. Jan 25, 2019.
  + [issue/29934-decision](https://golang.org/issue/29934#issuecomment-489682919) Damien Neil. Go 1.13 lunch decision about error values. May 6, 2019.
  + [issue/29934-russ](https://golang.org/issue/29934#issuecomment-490087200) Russ Cox. Response, Response regarding "proposal: Go 2 error values". May 7, 2019.
- [design/err-print](https://golang.org/design/go2draft-error-printing) Marcel van Lohuizen. Error Printing — Draft Design. August 27, 2018.
  + [issue/30468](https://golang.org/issue/30468) errors: performance regression in New.
- [issue/40432](https://golang.org/issue/40432) language: Go 2: error handling meta issue
- [issue/41198](https://golang.org/issue/41198) proposal: errors: add ErrUnimplemented as standard way for interface method to fail.

### Channel/Select

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov. Go channels on steroids. Jan 28, 2014
  + [issue/8899](https://golang.org/issue/8899) runtime: lock-free channels
  + [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) update on "lock-free channels"
  + [cl/112990043](https://codereview.appspot.com/112990043/) runtime: fine-grained locking in select
  + [cl/110580043](https://codereview.appspot.com/110580043/) runtime: add fast paths to non-blocking channel operations
- [issue/8898](https://golang.org/issue/8898) runtime: special case timer channels
- [issue/8903](https://golang.org/issue/8903) runtime: make chan-based generators faster.
- [issue/21806](https://golang.org/issue/21806) fairness in select statement.

### Generics

- [doc/generics-discuss](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/edit#heading=h.vuko0u3txoew) Summary of Go Generics Discussions
- [doc/generics-dilemma](https://research.swtch.com/generic) Russ Cox. The Generic Dilemma. December 3, 2009.
- [design/type-functions](https://golang.org/design/15292/2010-06-type-functions) Ian Lance Taylor. Type Functions. June 2010.
- [design/generalized-types](https://golang.org/design/15292/2011-03-gen) Ian Lance Taylor. Generalized Types. March 2011.
- [design/code-gen](https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM) Russ Cox. Alternatives to Dynamic Code Generation in Go. September 2012.
- [design/generalized-types2](https://golang.org/design/15292/2013-10-gen) Ian Lance Taylor. Generalized Types In Go. October 2013.
- [design/type-parameters](https://golang.org/design/15292/2013-12-type-params) Ian Lance Taylor. Type Parameters in Go. December 2013.
- [design/compile-time-function](https://golang.org/design/15292/2016-09-compile-time-functions) Bryan C. Mills. Compile-time Functions and First Class Types. September 2016.
- [design/should-generics](https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics) Ian Lance Taylor. Go should have generics. January 2011.
- [design/should-generics2](https://golang.org/design/15292-generics) Ian Lance Taylor. Go should have generics. Updated: April 2016.
  + [issue/15292](https://golang.org/issue/15292) proposal: spec: generic programming facilities
- [design/generics-overview](https://golang.org/design/go2draft-generics-overview) Russ Cox. Generics — Problem Overview. August 27, 2018.
- [design/contracts](https://golang.org/design/go2draft-contracts) Ian Lance Taylor, Robert Griesemer. Contracts — Draft Design. August 27, 2018, Updated: July 31, 2019.
  + [cl/187317](https://golang.org/cl/187317/) Implementation prototype.
  + [paper/featherweight-go](https://arxiv.org/abs/2005.11710) Griesemer, Robert, et al. Featherweight Go. arXiv preprint arXiv:2005.11710 (2020).
  + [talk/featherweight-go](https://www.youtube.com/watch?v=Dq0WFigax_c) Phil Wadler: Featherweight Go. Jun 8, 2020.
- [design/type-parameters2](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md) Ian Lance Taylor, Robert Griesemer. Type Parameters - Draft Design. June 16, 2020, Updated: August 28, 2020.
  + [cl/dev.go2go](https://github.com/golang/go/blob/dev.go2go/README.go2go.md) dev.go2go branch
  + [doc/type-check-readme](https://github.com/golang/go/tree/dev.go2go/src/go/types) type checking.
  + [doc/type-check-notes](https://github.com/golang/go/blob/dev.go2go/src/go/types/NOTES) This file serves as a notebook/implementation log.
- [discuss/generics-parenthesis](https://groups.google.com/g/golang-nuts/c/7t-Q2vt60J8) Robert. Generics and parenthesis.
- [issue/33232](https://golang.org/issue/33232) proposal: Go 2: add alias for interface {} as any
- [discuss/generics-move](https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw) Ian, Moving forward with the generics design.

## Compiler Toolchain

### Compiler

- [design/go12symtab](https://golang.org/s/go12symtab) Russ Cox. Go 1.2 Runtime Symbol Information. July 2013.
- [design/go13compiler](https://golang.org/s/go13compiler) Russ Cox. Go 1.3+ Compiler Overhaul. December 2013
- [design/go14generate](https://golang.org/s/go1.4-generate) Rob Pike. Go generate: A Proposal
- [design/dev.cc](https://golang.org/s/dev.cc)  Russ Cox. dev.cc branch plan. Nov 2014
- [design/go15bootstrap](https://golang.org/s/go15bootstrap) Russ Cox. Go 1.5 Bootstrap Plan. January 2015.
- [doc/escape-analysis](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/edit#) Dmitry Vyukov. Go Escape Analysis Flaws. Feb 10, 2015.
- [design/execmodes](https://golang.org/s/execmodes) Ian Lance Taylor. Go Execution Modes. August, 2014 (updated January 2016)
- [design/go17ssa](https://golang.org/s/go17ssa) Keith Randall. New SSA Backend for the Go Compiler. 2/10/2015.
- [doc/compiler-optimization](https://github.com/golang/go/wiki/CompilerOptimizations) Compiler And Runtime Optimizations.
- [issue/6853](https://golang.org/issue/6853) all: binaries too big and growing.
- [design/go19inlining](https://golang.org/design/19348-midstack-inlining) David Lazar, Austin Clements. Proposal: Mid-stack inlining in the Go compiler.
  + [issue/19348](https://golang.org/issue/19348) cmd/compile: enable mid-stack inlining.
  + [talk/go19inliningtalk](https://golang.org/s/go19inliningtalk) David Lazar. Mid-stack inlining in the Go compiler
- [design/dwarf-inlining](https://golang.org/design/22080-dwarf-inlining) Than McIntosh. Proposal: emit DWARF inlining info in the Go compiler. 2017-10-23.
  + [issue/22080](https://golang.org/issue/22080) cmd/compile: generated DWARF inlining info
- [issue/23109](https://golang.org/issue/23109) cmd/compile: rewrite escape analysis.
- [issue/27167](https://golang.org/issue/27167) cmd/compile: rename a bunch of things
  - [doc/renames](https://docs.google.com/document/d/19_ExiylD9MRfeAjKIfEsMU1_RGhuxB9sA0b5Zv7byVI/edit) Proposed Go 1.12 toolchain renames
- proposal: add GOEXPERIMENT=checkptr
  + [issue/22218](https://golang.org/issue/22218) proposal: add GOEXPERIMENT=checkptr
  + [issue/34964](https://golang.org/issue/34964) cmd/compile: enable -d=checkptr as part of -race and/or -msan?
  + [issue/34972](https://golang.org/issue/34972) all: get standard library building with -d=checkptr
  + [discuss/checkptr](https://groups.google.com/forum/#!msg/golang-dev/SzwDoqoRVJA/Iozu8vWdDwAJ)
- [issue/8885](https://golang.org/issue/8885) runtime: consider adding 24-byte size class.
- [issue/16798](https://golang.org/issue/16798) proposal: cmd/compile: add tail call optimization for self-recursion only.
- [design/64align](https://golang.org/design/36606-64-bit-field-alignment) Dan Scales. Proposal: Make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed, //go:align directives. 2020-06-08.
  + [issue/599](https://golang.org/issue/599) cmd/compile: make 64-bit fields 64-bit aligned on 32-bit systems
  + [issue/36606](https://golang.org/issue/36606) proposal: cmd/compile: make 64-bit fields be 64-bit aligned on 32-bit systems, add //go:packed directive on structs
- [talk/gccgo](https://www.youtube.com/watch?v=U0w9eFunkX4) Brief overview of gccgo, "the other" Go compiler. Aug 6, 2015.]

### Linker

The Go Linker was written by Ken Ken Thompson. Russ led a few more overhaul in Go 1.3. Austin led a rework to the linker together with Keith, Than,
Cheery and many other brilliant brains, which was landed
in Go 1.15 and Go 1.16.

- [design/go13linker](https://golang.org/s/go13linker) Russ Cox. Go 1.3 Linker Overhaul. Nov 2013.
- [design/go116linker](https://golang.org/s/better-linker) Austin Clements. Building a better Go linker. 2019-09-12.

### Debugger

- [doc/go13heapdump](https://golang.org/s/go13heapdump) heapdump13
- [doc/go14heapdump](https://github.com/golang/go/wiki/heapdump14) heapdump14
- [doc/go15heapdump](https://github.com/golang/go/wiki/heapdump15-through-heapdump17) heapdump15 through heapdump17
- [design/heap-viewer](https://golang.org/design/16410-heap-viewer) Michael Matloob. Proposal: Go Heap Dump Viewer. 20 July 2016
  + [issue/16410](https://golang.org/issue/16410) x/tools/cmd/heapdump: create a heap dump viewer
- [design/profiler-labels](https://golang.org/design/17280-profile-labels) Michael Matloob. Proposal: Support for pprof profiler labels. 15 May 2017.
  + [issue/17280](https://golang.org/issue/17280) pprof: add support for profiler labels

### Tracer

- [design/go15trace](https://golang.org/s/go15trace) Dmitry Vyukov. Go Execution Tracer. Oct 2014
- [design/tracefmt](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@. Trace Event Format. October 2016.

### Builder

- [design/go13nacl](https://golang.org/s/go13nacl) Russ Cox. Go 1.3 Native Client Support. October 2013.
- [design/gobuilder](http://golang.org/s/builderplan) Brad Fitzpatrick. Go Builder Plan. 2014-09-03.
  + [discuss/gobuilder](https://groups.google.com/g/golang-dev/c/sdFD0-2Ed8k) Changing how builders run
- [design/go14android](https://golang.org/s/go14android) David Crawshaw. Go support for Android. June 2014.
- [design/go-generate](https://golang.org/s/go1.4-generate) Rob Pike. Go Generate. January 2014.
- [issue/13560](https://golang.org/issue/13560) proposal: build: define standard way to recognize machine-generated files
- [discuss/generatedcode](golang.org/s/generatedcode) Rob Pike's Final Comments on Issue 13560
- [design/goenv](golang.org/design/30411-env) Russ Cox. Proposal: go command configuration file. March 1, 2019
  + [issue/30411](golang.org/issue/30411) proposal: cmd/go: add go env -w to set default env vars
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
- [design/vgo](https://golang.org/design/24301-versioned-go) Russ Cox. Proposal: Versioned Go Modules. March 20, 2018.
  + [issue/24301](https://golang.org/issue/24301) cmd/go: add package version support to Go toolchain
  + [doc/deps](https://research.swtch.com/deps) Russ Cox. Our Software Dependency Problem. January 23, 2019.
  + [doc/vgo](https://research.swtch.com/vgo) Russ Cox. Go & Versioning
  + [discuss/groups-gomod](https://groups.google.com/g/golang-dev/c/a5PqQuBljF4) go modules have landed
  + [discuss/go-dep-response](https://www.reddit.com/r/golang/comments/92f3q1/peter_bourgon_a_response_about_dep_and_vgo/) Reddit discussion.
  + [doc/go-dep-response](https://peter.bourgon.org/blog/2018/07/27/a-response-about-dep-and-vgo.html) Peter Bourgon. A response about dep and vgo. 2018-07-27.
  + [discuss/go-dep-response2](https://news.ycombinator.com/item?id=17628311) Hacker News discussion.
  + [discuss/go-dep-twitter](https://twitter.com/_rsc/status/1022588240501661696) Russ Cox's Twitter Storm
- [design/sumdb](https://golang.org/design/25530-sumdb) Russ Cox, Filippo Valsorda. Proposal: Secure the Public Go Module Ecosystem. April 24, 2019.
  + [issue/25530](https://golang.org/issue/25530) proposal: cmd/go: secure releases with transparency log
- [design/lazy-gomod](https://golang.org/design/36460-lazy-module-loading) Bryan C. Mills. Proposal: Lazy Module Loading. 2020-02-20

### gopls

- [design/gopls-workspace](https://golang.org/design/37720-gopls-workspaces) Heschi Kreinick, Rebecca Stambler. Proposal: Multi-project gopls workspaces. Mar 9, 2020.

### Testing

- [design/subtests](https://golang.org/design/12166-subtests) Marcel van Lohuizen. testing: programmatic sub-test and sub-benchmark support. September 2, 2015.
  + [issue/12166](https://golang.org/issue/12166) proposal: testing: programmatic sub-test and sub-benchmark support
- [design/gotest-bench](https://golang.org/design/14313-benchmark-format) Russ Cox, Austin Clements. Proposal: Go Benchmark Data Format. February 2016.
  + [issue/14313](golang.org/issue/14313) cmd/go: decide, document standard benchmark data format
- [design/gotest-json](https://golang.org/design/2981-go-test-json) Nodir Turakulov. Proposal: -json flag in go test. 2016-09-14.
- [design/testing-helper](https://golang.org/design/4899-testing-helper) Caleb Spare. Proposal: testing: better support test helper functions with TB.Helper. 2016-12-27
  + [issue/4899](https://golang.org/issue/4899) testing: add t.Helper to make file:line results more useful
- [design/fuzzing](https://golang.org/design/draft-fuzzing) Katie Hockman. Design Draft: First Class Fuzzing
<!-- - Tool chain, benchseries/benchstat -->

## Runtime Core

### Statistics

- [issue/16843](https://golang.org/issue/16843) runtime: mechanism for monitoring heap size
  + [cl/setmaxheap](https://go-review.googlesource.com/c/go/+/46751/) Austin Clements. runtime/debug: add SetMaxHeap API. Jun 26 2017.
- [issue/29696](https://golang.org/issue/29696) proposal: runtime: add way for applications to respond to GC backpressure
- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek. Proposal: API for unstable runtime metrics. Mar 18, 2020.
- [issue/19812](https://golang.org/issue/19812) runtime: cannot ReadMemStats during GC
- [issue/38712](https://golang.org/issue/38712) runtime: TestMemStats is flaky
- [issue/40459](https://golang.org/issue/40459) runtime: ReadMemStats called in a loop may prevent GC

### Scheduler

- [paper/work-steal](https://dl.acm.org/citation.cfm?id=324234) Robert D. Blumofe and Charles E. Leiserson. 1999. Scheduling multithreaded computations by work stealing. J. ACM 46, 5 (September 1999), 720-748.
- [cl/sched-m-1](https://github.com/golang/go/commit/96824000ed89d13665f6f24ddc10b3bf812e7f47#diff-1fe527a413d9f1c2e5e22e08e605a192) Russ Cox, Clean up scheduler. Aug 5, 2008.
- [cl/sched-m-n](https://github.com/golang/go/commit/fe1e49241c04c748d0e3f4762925241adcb8d7da) things are much better now, Nov 11, 2009.
- [design/go11sched](https://golang.org/s/go11sched) Dmitry Vyukov. Scalable Go Scheduler Design Doc, 2012
  + [cl/7314062](https://github.com/golang/go/commit/779c45a50700bda0f6ec98429720802e6c1624e8) runtime: improved scheduler
- [design/sched-preempt-dmitry](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#heading=h.3pilqarbrc9h) Dmitry Vyukov. Go Preemptive Scheduler Design Doc, 2013
- [design/sched-numa](https://docs.google.com/document/u/0/d/1d3iI2QWURgDIsSR6G2275vMeQ_X7w-qxM2Vp7iGwwuM/pub) Dmitry Vyukov, NUMA-aware scheduler for Go. Sep 2014.
- [design/go15gomaxprocs](https://golang.org/s/go15gomaxprocs) Russ Cox. Go 1.5 GOMAXPROCS Default. May 2015.
- [doc/go17sched](https://www.quora.com/How-does-the-golang-scheduler-work/answer/Ian-Lance-Taylor) Ian Lance Taylor. How does the golang scheduler work? July 16, 2016.
- [design/sched-preempt-austin](https://golang.org/design/24543-non-cooperative-preemption) Austin Clements. Proposal: Non-cooperative goroutine preemption. Jan 2019.
  + [cl/43050](https://golang.org/cl/43050) cmd/compile: loop preemption with fault branch on amd64. May 09, 2019.
  + [issue/10958](https://golang.org/issue/10958) runtime: tight loops should be preemptible
  + [issue/24543](https://golang.org/issue/24543) runtime: non-cooperative goroutine preemption
  + [issue/36365](https://golang.org/issue/36365) runtime: clean up async preemption loose ends
- [issue/14592](https://golang.org/issue/14592) runtime: let idle OS threads exit
- [issue/18237](https://golang.org/issue/18237) runtime: scheduler is slow when goroutines are frequently woken
- [issue/20395](https://golang.org/issue/20395) runtime: terminate locked OS thread if its goroutine exits
- [issue/20458](https://golang.org/issue/20458) proposal: runtime: pair LockOSThread, UnlockOSThread calls
- [issue/21827](https://golang.org/issue/21827) runtime: big performance penalty with runtime.LockOSThread
- [issue/27345](https://golang.org/issue/27345) runtime: use parent goroutine's stack for new goroutines
- [issue/28808](https://golang.org/issue/28808) runtime: scheduler work stealing slow for high GOMAXPROCS
- [issue/32113](https://golang.org/issue/32113) runtime: optimization to reduce P churn.

### Execution Stack

- [design/contigstack](https://golang.org/s/contigstacks) Contiguous stacks
- [issue/17007](https://golang.org/issue/17007) runtime: fatal error: bad g->status in ready
- [issue/26061](https://golang.org/issue/26061) runtime: g0 stack.lo is sometimes too low

### Memory Allocator

<!-- TODO: A quick history about the Go's memory allocator: Russ Cox first implements the memory allocator based on`tcmalloc` for Go 1, `mcache` are cached on M. Then he revised the allocator to allow
user code use 16GB memory and later allows 128GB. -->

- [doc/tcmalloc](http://goog-perftools.sourceforge.net/doc/tcmalloc.html) Sanjay Ghemawat, Paul Menage. TCMalloc : Thread-Caching Malloc. Google Inc., 2009
- [issue/30333](https://golang.org/issue/30333) runtime: smarter scavenging
- [issue/34047](https://golang.org/issue/34047) runtime: potential deadlock cycle caused by scavenge.lock
- [issue/34048](https://golang.org/issue/34048) runtime: scavenger pacing fails to account for fragmentation
- [issue/35112](https://golang.org/issue/35112) runtime: make the page allocator scale
- [issue/35788](https://golang.org/issue/35788) runtime: scavenger not as effective as in previous releases
- [design/go113scavenge](https://go.googlesource.com/proposal/+/aa701aae530695d32916b779e048a3e18311a2e3/design/30333-smarter-scavenging.md) Michael Knyszek. Proposal: Smarter Scavenging. 2019-05-09.
- [design/go114pagealloc](https://go.googlesource.com/proposal/+/a078ea9d72b99dc88fdfd2cb6ee150a8ce202ea2/design/35112-scaling-the-page-allocator.md) Michael Knyszek, Austin Clements. Proposal: Scaling the Go page allocator. 2019-10-18.
- [issue/37487](https://golang.org/issue/37487) runtime: improve mcentral scalability
  + [cl/221182](https://golang.org/cl/221182) runtime: add new mcentral implementation
- [issue/18155](https://golang.org/issue/18155) runtime: latency in sweep assists when there's no garbage
- [issue/19112](https://golang.org/issue/19112) runtime: deadlock involving gcControllerState.enlistWorker
- [issue/29707](https://golang.org/issue/29707) cmd/trace: failed to parse trace: no consistent ordering of events possible
- [issue/35788](https://golang.org/issue/35788) runtime: scavenger not as effective as in previous releases
- [issue/35954](https://golang.org/issue/35954) runtime: handle hitting the top of the address space in the allocator more gracefully
- [issue/37487](https://golang.org/issue/37487) runtime: improve mcentral scalability
- [issue/37927](https://golang.org/issue/37927) runtime: GC pacing exhibits strange behavior with a low GOGC
- [issue/38130](https://golang.org/issue/38130) runtime: incorrect sanity checks in the page allocator
- [issue/38404](https://golang.org/issue/38404) runtime: STW GC stops working on arm64/mips64le
- [issue/38605](https://golang.org/issue/38605) runtime: pageAlloc.allocToCache updates pageAlloc.searchAddr in an invalid way
- [issue/38617](https://golang.org/issue/38617) runtime: scavenger freezes up in Go 1.14 in Windows due to coarse time granularity
- [issue/38966](https://golang.org/issue/38966) runtime: aix-ppc64 builder is failing with "bad prune", "addr range base and limit are not in the same memory segment" fatal errors
- [issue/39128](https://golang.org/issue/39128) runtime: linux-mips-rtrk builder consistently failing with "bad prune" as of CL 233497
- [issue/40191](https://golang.org/issue/40191) runtime: pageAlloc.searchAddr may point to unmapped memory in discontiguous heaps, violating its invariant
- [issue/40457](https://golang.org/issue/40457) runtime: runqputbatch does not protect its call to globrunqputbatch
- [issue/40641](https://golang.org/issue/40641) runtime: race between stack shrinking and channel send/recv leads to bad sudog values

### Garbage Collector

- [paper/on-the-fly-gc](https://doi.org/10.1145/359642.359655) Edsger W. Dijkstra, Leslie Lamport, A. J. Martin, C. S. Scholten, and E. F. M. Steffens. 1978. On-the-fly garbage collection: an exercise in cooperation. Commun. ACM 21, 11 (November 1978), 966–975.
- [paper/yuasa-barrier](https://doi.org/10.1016/0164-1212(90)90084-Y) T. Yuasa. 1990. Real-time garbage collection on general-purpose machines. J. Syst. Softw. 11, 3 (March 1990), 181-198.
- [design/go13gc](https://docs.google.com/document/d/1v4Oqa0WwHunqlb8C3ObL_uNQw3DfSY-ztoA-4wWbKcg/pub) Dmitry Vyukov. Simpler and faster GC for Go. July 16, 2014
  + [cl/106260045](https://codereview.appspot.com/106260045) runtime: simpler and faster GC
- [design/go14gc](https://golang.org/s/go14gc) Richard L. Hudson. Go 1.4+ Garbage Collection (GC) Plan and Roadmap. August 6, 2014.
- [design/go15gcpacing](https://golang.org/s/go15gcpacing) Austin Clements. Go 1.5 concurrent garbage collector pacing. 2015-03-10.
- [discuss/gcpacing](https://groups.google.com/forum/#!topic/golang-dev/YjoG9yJktg4) Austin Clements et al. Discussion of "Proposal: Garbage collector pacing". March 10, 2015.
- [design/eliminate-rescan](https://golang.org/design/17503-eliminate-rescan) Austin Clements, Rick Hudson. Eliminate STW stack re-scanning. October 21, 2016.
- [issue/11970](https://golang.org/issue/11970) runtime: replace GC coordinator with state machine
- [design/sweep-free-alloc](https://golang.org/design/12800-sweep-free-alloc) Austin Clements. Proposal: Dense mark bits and sweep-free allocation. Sep 30, 2015.
- [issue/12800](https://golang.org/issue/12800) runtime: replace free list with direct bitmap allocation
- [design/decentralized-gc](https://golang.org/design/11970-decentralized-gc) Austin Clements. Proposal: Decentralized GC coordination. October 25, 2015.
- [issue/12967](https://golang.org/issue/12967#issuecomment-171466238) runtime: shrinkstack during mark termination significantly increases GC STW time
- [issue/14951](https://golang.org/issue/14951) runtime: mutator assists are over-aggressive, especially at high GOGC
- [issue/17503](https://golang.org/issue/17503) runtime: eliminate stack rescanning
- [design/concurrent-rescan](https://golang.org/design/17505-concurrent-rescan) Austin Clements, Rick Hudson. Proposal: Concurrent stack re-scanning. Oct 18, 2016.
- [issue/17505](https://golang.org/issue/17505) runtime: perform concurrent stack re-scanning
- [design/eliminate-rescan](https://golang.org/design/17503-eliminate-rescan) Austin Clements, Rick Hudson. Proposal: Eliminate STW stack re-scanning. Oct 21, 2016
- [design/soft-heap-limit](https://golang.org/design/14951-soft-heap-limit) Austin Clements. Proposal: Separate soft and hard heap size goal. October 21, 2017.
- [design/roc](https://golang.org/s/gctoc) Request Oriented Collector (ROC) Algorithm
  + [cl/roc](https://golang.org/cl/25058) runtime: ROC write barrier code
  + [cl/generational-gc](https://golang.org/cl/137482) runtime: trigger generational GC
- [doc/ismm-gc](https://blog.golang.org/ismmkeynote) Rick Hudson. Getting to Go: The Journey of Go's Garbage Collector. 12 July 2018.
- [discuss/ismm-gc](https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns) Garbage Collection Slides and Transcript now available
- [design/simplify-mark-termination](https://golang.org/design/26903-simplify-mark-termination) Austin Clements. Proposal: Simplify mark termination and eliminate mark 2. Aug 9, 2018.
- [issue/22350](https://golang.org/issue/22350) cmd/compile: compiler can unexpectedly preserve memory,
- [design/gcscan](https://docs.google.com/document/d/1un-Jn47yByHL7I0aVIP_uVCMxjdM5mpelJhiKlIqxkE/edit#) Proposal: GC scanning of stacks
- [issue/26903](https://golang.org/issue/26903) runtime: simplify mark termination and eliminate mark 2
- [issue/27993](https://golang.org/issue/27993) runtime: error message: P has cached GC work at end of mark termination

### Memory model

Go memory model is not well defined (yet), but atomic is likely to
guarantee sequential consistency.

- [doc/refmem](https://golang.org/ref/mem) Rob Pike and Russ Cox. The Go Memory Model. February 21, 2009.
- [issue/4947]( https://golang.org/issue/4947) cmd/cc: atomic intrinsics
- [issue/5045](https://golang.org/issue/5045) doc: define how sync/atomic interacts with memory model
- [issue/7948](https://golang.org/issue/7948) doc: define how sync interacts with memory model
- [issue/9442](https://golang.org/issue/9442) doc: define how finalizers interact with memory model
- [issue/33815](https://golang.org/issue/33815) doc/go_mem: "hello, world" will not always be printed twice
- [cl/75130045](https://codereview.appspot.com/75130045) doc: allow buffered channel as semaphore without initialization
- [doc/gomem](http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf) Russ Cox. Go’s Memory Model. February 25, 2016.
- [doc/go2017russ](https://research.swtch.com/go2017#memory) Russ Cox. My Go Resolutions for 2017 - Memory model. January 18, 2017.
- [doc/atomic-bug](https://golang.org/pkg/sync/atomic/#pkg-note-BUG) Package atomic
- [discuss/atomic-mem-order](https://groups.google.com/d/msg/golang-dev/vVkH_9fl1D8/azJa10lkAwAJ) specify the memory order guarantee provided by atomic Load/Store

### ABI

- [design/cgo-pointers](https://golang.org/design/12416-cgo-pointers) Ian Lance Taylor. Proposal: Rules for passing pointers between Go and C. October, 2015
  + [issue/12416](https://golang.org/issue/12416) cmd/cgo: specify rules for passing pointers between Go and C
- [design/internal-abi](https://golang.org/design/27539-internal-abi) Austin Clements. Proposal: Create an undefined internal calling convention. 2019-01-14.
  + [issue/27539](https://golang.org/issue/27539) proposal: runtime: make the ABI undefined as a step toward changing it.
- [design/register-call](https://golang.org/design/40724-register-calling) Austin Clements, with input from Cherry Zhang, Michael Knyszek, Martin Möhrmann, Michael Pratt, David Chase, Keith Randall, Dan Scales, and Ian Lance Taylor. Proposal: Register-based Go calling convention. 2020-08-10.
- [issue/18597](https://golang.org/issue/18597) proposal: cmd/compile: define register-based calling convention
- [issue/40724](https://golang.org/issue/40724) proposal: switch to a register-based calling convention for Go functions

## Standard Library

### syscall

- [design/go14syscall](https://golang.org/s/go1.4-syscall) The syscall package.

### io

- [design/draft-iofs](https://golang.org/design/draft-iofs) Russ Cox, Rob Pike. File System Interfaces for Go — Draft Design. July 2020.

### go/*

- [doc/gotypes](https://golang.org/s/types-tutorial) Alan Donovan. go/types: The Go Type Checker.
- [talk/gotypes](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view) Alan Donovan. Using go/types for
Code Comprehension and Refactoring Tools. October 2, 2015.
- [design/modular-interface](https://docs.google.com/document/d/1-azPLXaLgTCKeKDNg0HVMq2ovMlD-e7n1ZHzZVzOlJk/edit) Alan Donovan. Proposal: a common interface for modular static analyses for Go. 9 Sep 2018
  + [cl/134935](https://golang.org/cl/134935) go/analysis: a new API for analysis tools

### sync

- [design/percpu-sharded](https://golang.org/design/18802-percpu-sharded) Sanjay Menakuru. Proposal: percpu.Sharded, an API for reducing cache contention. Jun 12, 2018.
  + [issue/18802](https://golang.org/issue/18802) proposal: sync: support for sharded values

#### Pool

- [discuss/add-pool](https://groups.google.com/d/msg/golang-dev/kJ_R6vYVYHU/LjoGriFTYxMJ) gc-aware pool draining policy
- [issue/4720](https://golang.org/issue/4720) sync: add Pool type. Jan 28, 2013.
- [cl/46010043](https://github.com/golang/go/commit/f8e0057bb71cded5bb2d0b09c6292b13c59b5748#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: scalable Pool. Jan 24, 2014.
- [cl/86020043](https://github.com/golang/go/commit/8fc6ed4c8901d13fe1a5aa176b0ba808e2855af5#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: less agressive local caching in Pool. Apr 14, 2014.
- [cl/162980043](https://github.com/golang/go/commit/af3868f1879c7f8bef1a4ac43cfe1ab1304ad6a4#diff-491b0013c82345bf6cfa937bd78b690d) sync: release Pool memory during second and later GCs. Oct 22, 2014.
- [issue/8979](https://golang.org/issue/8979) sync: Pool does not release memory on GC
- [issue/13086](https://golang.org/issue/13086) runtime: fall back to fair locks after repeated sleep-acquire failures.
- [issue/22331](https://golang.org/issue/22331) runtime: clearpools causes excessive STW1 time
- [issue/22950](https://golang.org/issue/22950) sync: avoid clearing the full Pool on every GC.
  - [cl/166960](https://github.com/golang/go/commit/d5fd2dd6a17a816b7dfd99d4df70a85f1bf0de31) sync: use lock-free structure for Pool stealing.
  - [cl/166961](https://github.com/golang/go/commit/2dcbf8b3691e72d1b04e9376488cef3b6f93b286) 166961: sync: smooth out Pool behavior over GC with a victim cache.
- [issue/21035](https://golang.org/issue/21035) sync: reduce contention between Map operations with new-but-disjoint keys

#### Mutex

- [cl/4631059](https://github.com/golang/go/commit/997c00f) runtime: replace Semacquire/Semrelease implementation.

#### atomic

- [issue/20164](https://golang.org/issue/20164) proposal: atomic: add (*Value).Swap
- [discuss/atomic-value](https://groups.google.com/g/golang-dev/c/SBmIen68ys0/m/WGfYQQSO4nAJ)

### time

- [design/monotonic-time](https://golang.org/design/12914-monotonic) Russ Cox. Proposal: Monotonic Elapsed Time Measurements in Go. January 26, 2017.
  - [issue/12914](https://golang.org/issue/12914) time: use monotonic clock to measure elapsed time
- [issue/6239](https://golang.org/issue/6239) runtime: make timers faster.
- [issue/15133](https://golang.org/issue/15133) runtime: timer doesn't scale on multi-CPU systems with a lot of timers
  + [cl/34784](https://golang.org/cl/34784) runtime: improve timers scalability on multi-CPU systems
- [issue/18023](https://golang.org/issue/18023) runtime: unexpectedly large slowdown with runtime.LockOSThread
- [issue/27707](https://golang.org/issue/27707) time: excessive CPU usage when using Ticker and Sleep.
- [issue/38070](https://golang.org/issue/38070) runtime: timer self-deadlock due to preemption point

### context

- [issue/8082](https://golang.org/issue/8082) proposal: spec: represent interfaces by their definition and not by package and name
- [issue/14660](https://golang.org/issue/14660) proposal: context: new package for standard library
- [issue/16209](https://golang.org/issue/16209) Proposal: relaxed rules for assignability with differently-named but identical interfaces
- [issue/20280](https://golang.org/issue/20280) proposal: io: add Context parameter to Reader, etc.
- [issue/21355](https://golang.org/issue/21355) proposal: Replace Context with goroutine-local storage
- [issue/24050](https://golang.org/issue/24050) testing: test with child process sometimes hangs on 1.10; -timeout not respected
- [issue/27982] proposal: Go 2: bake a cooperative goroutine cancellation mechanism into the language
- [issue/28342](https://golang.org/issue/28342) proposal: Go 2: update context package for Go 2
- [issue/29011](https://golang.org/issue/29011) proposal: Go 2: use structured concurrency
- [doc/context-go-away](https://faiface.github.io/post/context-should-go-away-go2/) Michal Štrba. Context should go away for Go 2. 2017/08/06
- [doc/context](https://blog.golang.org/context) Go Concurrency Patterns: Context.
- [doc/context-isnt-for-cancellation](https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation) Dave Cheney. Context isn’t for cancellation. August 20, 2017.

### encoding

- [design/go12encoding](https://golang.org/s/go12encoding) Russ Cox. Go 1.2 encoding.TextMarshaler and TextUnmarshaler. July 2013.
- [design/go12xml](https://docs.google.com/document/d/1G-AMeUPoyTnbZDkuCJA89DjJwTjdWFctIZ_N9NgA9RM/pub) Russ Cox. Go 1.2 xml.Marshaler and Unmarshaler. July 2013.
- [design/natural-xml](https://golang.org/design/13504-natural-xml) Sam Whited. Proposal: Natural XML. 2016-09-27
  + [issue/13504](https://golang.org/issue/13504) encoding/xml: add generic representation of XML data
- [design/zip](https://golang.org/design/14386-zip-package-archives) Russ Cox. Proposal: Zip-based Go package archives. February 2016
  + [issue/14386](https://golang.org/issue/14386) proposal: use zip format inside .a and .o files
- [design/xmlstream](https://golang.org/design/19480-xml-stream) Sam Whited. Proposal: XML Stream. 2017-03-09
  + [issue/19480](https://golang.org/issue/19480) encoding/xml: add decode from TokenReader, to enable stream transformers
- [design/raw-xml](https://golang.org/design/26756-rawxml-token) Sam Whited. Proposal: Raw XML Token. 2018-09-01
  + [issue/26756](https://golang.org/issue/26756) proposal: encoding/xml: add RawXML token

### image, x/image

- [issue/11420](https://golang.org/issue/11420) x/image/draw: color space-correct interpolation
- [issue/20613](https://golang.org/issue/20613) image/png: don't ignore PNG gAMA chunk
- [issue/27830](https://golang.org/issue/27830) proposal: image: decoding options
- [issue/33457](https://golang.org/issue/33457) proposal: image: add generic metadata support for jpeg, gif, png
  - [cl/208559](https://golang.org/cl/208559) 
  - [cl/216799](https://golang.org/cl/216799)

### misc

- [design/mobile-audio](https://golang.org/design/13432-mobile-audio) Jaana Burcu Dogan. Proposal: Audio for Mobile. November 30, 2015.
  + [issue/13432](https://golang.org/issue/13432) proposal: x/mobile audio
- [design/localization](https://golang.org/design/12750-localization) Marcel van Lohuizen. Proposal: Localization support in Go. Jan 28, 2016.
  + [issue/12750](https://golang.org/issue/12750) x/text: localization support
- [design/unbounded-queue](https://golang.org/design/27935-unbounded-queue-package) Christian Petrin. Proposal: Built in support for high performance unbounded queue. October 2, 2018
  + [issue/27935](https://golang.org/issue/27935) proposal: add container/queue
- [design/lockfile](https://golang.org/design/33974-add-public-lockedfile-pkg) Adrien Delorme. Proposal: make the internal lockedfile package public. 2019-10-15.
  + [issue/33974](https://golang.org/issue/33974)
- [design/cidr](https://golang.org/design/16704-cidr-notation-no-proxy) Rudi Kramer, James Forrest. Proposal: Add support for CIDR notation in no_proxy variable. 2017-07-10
  + [issue/16704](https://golang.org/issue/16704) net/http: considering supporting CIDR notation in no_proxy env variable
- [design/dns](https://golang.org/design/26160-dns-based-vanity-imports) Sam Whited. Proposal: DNS Based Vanity Imports. 2018-06-30.
  + [issue/26160](https://golang.org/issue/26160) proposal: use DNS TXT records for vanity import paths

## Unclassified But Relevant Links

- https://golang.org/s/using-guru
- [Plan for Go Guru](https://docs.google.com/document/d/1UErU12vR7jTedYvKHVNRzGPmXqdMASZ6PfE7B-p6sIg/edit)
- https://golang.org/s/stdwhy
- https://golang.org/s/oracle-design
- https://golang.org/s/oracle-user-manual
- https://golang.org/s/cgihttpproxy
- https://golang.org/s/sqldrivers
- https://golang.org/s/go2designs
- http://doc.cat-v.org/bell_labs/pikestyle
- https://9p.io/plan9/
- https://www.informit.com/articles/article.aspx?p=1941206
- http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html
- https://softwareengineeringdaily.com/2017/12/28/language-design-with-brian-kernighan-holiday-repeat/
- https://www.youtube.com/watch?feature=youtu.be&v=de2Hsvxaf8M&ab_channel=Computerphile
- https://software.intel.com/content/www/us/en/develop/blogs/debugging-performance-issues-in-go-programs.html
- https://github.com/dgryski/interesting-papers
- https://github.com/golang/go/wiki/ResearchPapers


## Fun Facts

- [cl/1](https://github.com/golang/go/commit/7d7c6a97f815e9279d08cfaea7d5efb5e90695a8) Brian Kernighan. Go's first commit. Jul 19, 1972.
- [issue/9](https://golang.org/issue/9) I have already used the name for *MY* programming language

## License

[go-history](https://github.com/changkun/go-history) CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)