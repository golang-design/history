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
Rob and Ken are retired. Robert and Ian currently work on adding generics
to Go. Russ is also one of the core authors of the project in the early stage.
Back then, he was a newcomer at Google, and Rob invited Russ for joining
the Go team since he knew Russ from way back because of the
[Plan 9](https://9p.io/plan9/) project. Russ worked on the early Go compiler, runtime, as well as the leap of
Go 1.5 bootstrap.
Now, Russ is the tech leader of the Go team.

- Rob Pike. (Robert C. Pike, M. Sc.) [Website](http://herpolhode.com/rob/), [Blog](https://commandcenter.blogspot.com/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike). (Retired)
  + Alma mater: University of Toronto (BS), California Institute of Technology
  + [talk/rob2007](https://www.youtube.com/watch?v=hB05UFqOtFA) <details><summary>Advanced Topics in Programming Languages: Concurrency/message passing Newsqueak. May 9, 2007</summary>
    <p><b>Summary:</b> <b>Summary:</b> In this Google Tech Talk from May 2007, Rob Pike, creator of UTF-8 and Plan 9 for Bell Labs, discusses his programming language, New Squeak, which features interesting mechanisms for concurrency and message passing. Pike argues that there are two ways to deal with the mismatch between the concurrent world we live in and the sequential computers we use to write programs: make the world look synchronous and sequential, or make software concurrent. Pike discusses the limitations of low-level concepts like threads, shared memory, and locks when it comes to programming for concurrency. New Squeak is a programming language that was created in 1988 to address issues that arise when writing software for a concurrent world. It has lambdas as functions, process management software, and channels as first-class citizens. The language inspired other languages and was used as a systems language at Bell Labs for a few years to build interesting tools. The talk focuses on processes and channels in New Squeak and how to write programs using these concepts. Channels are a communication tool used in programming that reduce the number of tokens needed for communication. Channels in CSP are introduced by the keyword Chan in New Squeak and are unbuffered synchronous communication ports that synchronize and communicate values between processes. New Squeak's channels and select statements enable potential communications and allow processes to block until one or more communications can proceed. The article introduces a simple language with n-way muxing and presents a program that prints sequential integers to a channel. It also explains the concept of a prime sieve and filter process to find all the primes that come out of a calculation. The system can also be used to manipulate power series. The article discusses the use of interfaces and channels in programming, with an example from the Newspeak language. The author discusses a system model used in building window systems in Plan Nine, which involves defining components as interfaces that capture communication and data flow through channels. This approach was used in all major user-level services in Plan Nine and was made easier by the author's previous experience with concurrency models in programming. The model allows for complex interactions between components and allows for composing interfaces themselves, rather than just state machines. The author also discusses the challenges of debugging concurrent programs and the limitations of CSP-like libraries for implementing select operations. They propose the idea of a shim interface that guarantees deadlock-free operation and can isolate and kill misbehaving clients. The Squint interpreter, a code from 1988, lacks many libraries for graphics and other functions. The idea of implementing a communication channel model on top of mutexes in C++ is discussed, and the cost of a channel versus a callback is debated. Optimizations can be made to generate more compact code for different forms of select.</p>
  </details>

  + [talk/rob2009](https://changelog.com/podcast/3) The Go Programming Language. Nov 27, 2009.
  + [talk/rob2010a](https://www.youtube.com/watch?v=jgVhBThJdXc) <details><summary>Go Programming. Google I/O 2010. May 20, 2010</summary>
    <p><b>Summary:</b> In the Google I/O 2010 podcast episode titled "Go Programming, " Rob Pike and Russ Cox discuss the unique features and principles of the Go programming language. They emphasize the importance of using core concepts and idioms specific to Go, rather than translating code from other languages. One key feature of Go is its different types, including basic and composite types. Pike and Cox explain that Go is object-oriented but not type-oriented, meaning it does not have classes or inheritance. They also highlight the implicit nature of Go, where type declarations can be omitted. Concurrency is another significant aspect of Go programming. Pike and Cox explain that Go focuses on concurrent programming rather than parallel programming. They discuss the benefits of concurrency and how it allows for the construction of well-structured programs that can effectively utilize multiple cores. The podcast episode also delves into the use of interfaces in Go programming. Interfaces are used to define common methods that different types can implement, allowing for code reuse and flexibility. Pike and Cox provide examples of how interfaces are used in the block cipher package in Go. The speakers compare Go to other languages, such as Java, and highlight the advantages of using Go for implementing interfaces. They mention the flexibility of multiple wrappers and satisfying multiple interfaces. The concept of chaining readers together to decrypt and decompress data is also discussed in the episode. Pike and Cox explain the concept of load balancing in distributed systems and describe a simple model for distributing tasks to workers. They discuss the role of a load balancer in handling requests and the use of request and response channels. Throughout the episode, Pike and Cox emphasize the unique principles and features of Go, such as closures and channels, which make it a preferred choice for concurrent programming. They also mention the suitability of Go for different environments, such as server, desktop, and mobile. In conclusion, the "Go Programming" episode of the Google I/O 2010 podcast provides an in-depth overview of the Go programming language. Pike and Cox highlight its unique features, such as different types, concurrency, and the use of interfaces. They compare Go to other languages and discuss its advantages for concurrent programming. Overall, the episode showcases the power and efficiency of Go and its potential for various programming environments.</p></details>
  + [talk/rob2010b](https://www.youtube.com/watch?v=-_DKfAn4pFA) <details><summary>Origins of Go Concurrency style. Emerging Languages Camp 2010. July 21, 2010.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Rob Pike. Origins of Go Concurrency Style, " Rob Pike discusses the origins and concepts of concurrency in the Go programming language. He highlights the influence of Tony Hoare's paper on communicating sequential processes (CSP) in 1978, which emphasized communication and parallel composition of sequential processes. Pike describes CSP as a mathematical and elegant language that focuses on communication and synchronization between sender and receiver. Pike explains how processes can be combined in parallel, resembling a pipeline, and mentions the limitations of CSP, such as the inability to dynamically create processes or use send as a guard, and the lack of support for threads or mutexes. He also discusses the development of concurrent programming models and the emergence of the Occam language, which laid out the foundations for programming multi-processors and concurrent algorithms without the need for locks. The podcast episode also touches on the development of the Limbo language into Go, highlighting the power of Go's CSP model and its use of channels for communication and concurrency. Channels are described as first-class values that enable the transmission of data and capabilities between processes. Pike mentions the development of the lf language, which faced challenges due to its lack of garbage collection, but Limbo addressed these issues and became more successful in its limited domain. The benefits of using concurrency in programming are discussed, particularly in the context of the CSP model. Pike emphasizes how concurrency allows for efficient computation in areas such as cryptography and graphics, while also providing the opportunity for code reusability. He recommends reading the original CSP paper for a deeper understanding of the language and discusses the concepts of process control and communication, emphasizing the importance of understanding the power of co-routines. Pike also highlights the ease and safety of writing concurrent programs in Go, emphasizing that Go's concurrency features are natural and easy to use. He mentions the safety of Go's memory system and type safety, cautioning against using the unsafe package. The language has been successfully used in various applications, including high-traffic websites, and offers powerful channel capabilities for communication. The podcast episode concludes by explaining the concept of channels in a systems language, where types determine what can be sent through a channel. It discusses the efficiency and flexibility of this approach compared to traditional sharing and locking models. Pike also mentions the concept of passing pointers in programming and how it relates to efficiency and ownership, emphasizing the importance of understanding that once a pointer is passed, it is no longer the concern of the original owner. The episode highlights the idea that making things easy can lead to increased efficiency, which is seen as a positive point.</p></details>
  + [talk/rob2010c](https://www.youtube.com/watch?v=5kj5ApnhPAE) <details><summary>Public Static Void. OSCON 2010. Jul 22, 2010.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Public Static Void" from OSCON 2010, Rob Pike, a software engineer from Google, Inc. , discusses various topics related to programming languages and their complexities. Pike emphasizes the value of early programming languages, highlighting their simplicity and efficiency compared to the complexity and noise of modern software development. One of the challenges Pike addresses is the intricacy involved in calling functions in C++, as well as the absence of garbage collection. He also mentions the difficulty for non-expert programmers in choosing from different Boost templated pointer types for memory management. Pike acknowledges the complexity of sophisticated programming languages like C++ and the issues that arise from this. The significance of C++ and Java in programming, particularly in education and industry, is also discussed. Pike questions why these languages have become the standard and provides a simplified history of their development. He disagrees with the trend of using C++ and Java for teaching, arguing that both languages are too complex and verbose to be user-friendly. The limitations of using patterns in software development are highlighted, with the author suggesting that as programming languages improve, the need for patterns may decrease. Object-oriented programming languages are criticized for being bureaucratic and repetitive, and the text encourages awareness of alternative programming models. The issue of repetitive and nonsensical code is addressed, with an example given to illustrate the importance of avoiding such code in programming. The problems with using ambiguous integers and booleans in code are also discussed, emphasizing the need for clear data declarations. The bureaucratic nature of programming languages is explored, with popular languages like Python, Ruby, and JavaScript being criticized for becoming cumbersome. The emergence of new programming languages, such as Haskell and Scala, is seen as a response to the frustrations of working with older languages. The misconception that dynamic, interpreted languages are superior to compiled, static languages is challenged. The limitations of both old and new programming languages are discussed, and the need for a language that combines the benefits of both is argued for. In conclusion, the podcast episode delves into the complexities and challenges of programming languages. It highlights the importance of simplicity, efficiency, and user-friendliness in programming languages, and suggests that better solutions are needed to reduce the burden on developers. The Go programming language is mentioned as an attempt to fulfill these requirements, combining the safety and performance of a statically typed compiled language with the expressiveness and convenience of a dynamically typed interpreted language.</p></details>
  + [talk/rob2010d](https://www.youtube.com/watch?v=7VcArS4Wpqk) <details><summary>Another Go at Language Design. Aug 27, 2010.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Another Go at Language Design, " Rob Pike, a Principal Engineer at Google, Inc. , discusses the development of the Go programming language. Pike explains that the creation of Go was driven by the need for a language that could address the challenges faced by large code bases. He emphasizes that the development of Go is a collaborative effort involving many talented individuals, and expresses gratitude for the talented individuals he is working with. Pike discusses the concept of a "one true way" in computer science and how it depends on tools, problems, and beliefs. He mentions an upcoming talk about dynamic languages and static techniques. Pike also shares a quote about the simplicity and effectiveness of early programming languages. The podcast episode highlights the challenges of managing dependencies in C and C++ programs, which can lead to slower compilation times and hinder scalability. It introduces Go as an alternative, highlighting its ability to handle packages and dependencies effectively. The benefits of Go are discussed, including its speed and the ability to compile and run programs faster. The author suggests including the compiler in the runtime system to make the process even more efficient. The value of a new and clean compiler is emphasized, as it can significantly reduce build times and improve overall performance. The episode also explains methods and interfaces in Go, covering the syntax and examples of method declarations. The use of interfaces to define behavior is discussed, as well as the concept of empty interfaces, which can be satisfied by any type. The power of empty interfaces in controlling printing behavior is explored, mentioning the implementation of the printf function. Implicit interfaces are introduced, emphasizing the importance of type safety and convention. The flexibility and abstraction provided by interfaces are discussed, using the example of the 'reader' interface. The use of mutex and channels in worker pools to manage data sharing and communication is explained. The importance of garbage collection in concurrent server software is discussed, and how Go provides intrinsic safety and simplifies interface design. The use of the 'unsafe' library in low-level programming is mentioned, along with its potential risks. The improved performance of bounds check loops in modern machines and the reduced risk of buffer overflow exploits in Go are highlighted. The advantages of using Go for systems development are discussed, including its features such as unsigned types, bit-level operations, and control over memory layout. The design principles of Go are emphasized, including its control, safety, simplicity, and clarity. The episode explains how Go provides control over memory allocation and usage while ensuring safety without sacrificing performance. The visibility rules and the concept of constants in Go are also mentioned. Overall, the podcast episode praises the simplicity, efficiency, and productivity benefits of the Go programming language. It highlights its use in large-scale software development, including within Google, and its favorable license and open-source development. The episode also discusses various concepts in programming, including package namespaces, exceptions, classes and inheritance, and channels in the concurrency model. It emphasizes the importance of teamwork and unanimous decision-making in the collaborative design process of Go.</p></details>
  + [talk/rob2011a](https://www.infoq.com/interviews/pike-concurrency/) Parallelism and Concurrency in Programming Languages. Feb 17, 2011.
  + [talk/rob2011b](https://www.infoq.com/interviews/pike-google-go/) Google Go: Concurrency, Type System, Memory Management and GC. Feb 25, 2011.
  + [talk/rob2011c](https://www.youtube.com/watch?v=HxaD_trXwRE) <details><summary>Lexical Scanning in Go. Aug 30, 2011.</summary>
    <p><b>Summary:</b> In this episode of the Google Technology User Group, Rob Pike delivers a talk on Lexical Scanning in Go. The talk, given on Tuesday, 30 August 2011, focuses on the Go programming language and its relevance to solving structural mismatch problems in computing. Pike begins by explaining the concept of lexemes and tokens in programming languages. He discusses the challenges of tokenization and the advantages and disadvantages of using lexical analysis tools. Pike argues that it is often more efficient to write a custom lexer and highlights the importance of adaptability across programming languages. The process of writing a lexer is then discussed, which involves defining states and actions. Pike proposes a better approach to the current state machine model by continuously moving to the next state instead of discarding the current state. He introduces the concept of state functions, which are functions that return another state function and can be used in a lexer to loop through different states. Pike also explains how Go channels can enable communication between a lexer and a parser. He provides an explanation of the purpose of variables such as 'start' and 'pause' in a lexer and describes the role of Lex Text in scanning the input. The structure of an action block and the process of transitioning to a new state are also discussed. The talk then delves into the process of parsing characters in a template system, including how different characters are handled and the use of helper functions. Pike explains the concept of acceptors in a lexer, which are helpful for scanning complicated input. He also discusses how to lex numbers in a string, including different number formats and the use of tools to find the end of a number. Pike emphasizes the importance of error checking in parsing and validating numeric input. He mentions the use of a parser library to convert the input into a number, relieving the programmer from manual conversion. The concept of error functions in state machines is also explained, along with their role in creating formatted error messages. The talk concludes with a discussion on the challenges of running Go routines to completion during initialization in Go programming. Pike suggests a solution by changing the input and using a traditional lexing API without channels. The process of transforming existing code to use a traditional lexing API is explained, along with the concept of a run loop and select statement in programming. Overall, Pike's talk provides valuable insights into lexical scanning in Go and offers practical solutions to common challenges faced by programmers.</p></details>
  + [talk/rob2012a](https://www.youtube.com/watch?v=f6kdp27TYZs) <details><summary>Go Concurrency Patterns. Google I/O 2012. Jul 2, 2012.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Google I/O 2012 - Go Concurrency Patterns, " Rob Pike discusses the importance of concurrency in designing high-performance network services. Pike, a renowned expert in programming, highlights Go's concurrency primitives, such as goroutines and channels, which provide a simple and efficient way to express concurrent execution. Pike begins by explaining the concept of concurrency and its practical uses. He clarifies that Go is a concurrent language, not a parallel one, and discusses how concurrent code can still have a useful structure even when running on a single processor. He also discusses the origins of concurrent programming ideas and languages, leading up to the development of Go. Pike emphasizes the unique features and intellectual depth of these languages. The podcast delves into the concept of goroutines in Go programming and how they allow for concurrent execution of multiple functions. Pike emphasizes the importance of communication and synchronization in creating proper concurrent programs. He provides a comprehensive explanation of channels in Go, which are essential for concurrent programming. Pike covers the declaration, initialization, sending, and receiving of values on channels. He highlights that channel operations are blocking and serve as a synchronization mechanism between goroutines. Pike also mentions the use of buffered channels and the Go approach of using channels for communication and synchronization instead of sharing memory. The episode explores the use of concurrency and the select statement in Go programming. The select statement is highlighted as a key feature of Go's concurrency model, allowing for easier control of program behavior based on communication. Pike discusses non-blocking communication, timeouts, and the use of a quit channel to signal the end of a process. He emphasizes the importance of proper communication to avoid premature program shutdown and the need for sophisticated communication between programs for synchronization. Pike also discusses the concept of goroutines, which are lightweight elements in programming. He uses the example of running 100, 000 gophers to illustrate the speed and efficiency of goroutines. Pike explains the process of how a Google search works, including the use of independently executing backends to find and deliver search results. He describes a process of testing the speed of a program and measures the time it takes for the search results to be obtained. Pike explains how launching goroutines for each backend makes the search process concurrent and parallel, reducing waiting time. The episode concludes by discussing the advantages of using concurrency in Go programming. Pike highlights the simplicity and reliability of using Go compared to other approaches. He also discusses the use of concurrency tools in software construction and the importance of caution when experimenting with concurrent programs. Pike addresses questions about channel locking, garbage collection, stack allocation, and the select control structure in Go. Overall, the podcast episode provides a comprehensive exploration of Go's concurrency features and their practical applications in designing high-performance network services.</p></details>
  + [talk/rob2012b](https://www.youtube.com/watch?v=FTl0tl9BGdc) <details><summary>Why Learn Go?. Sep 12, 2012.</summary>
    <p><b>Summary:</b> In this podcast episode titled "Why Learn Go?", Rob Pike, co-creator of the Go programming language, discusses the need for a new compiled language that can effectively handle the demands of modern computing. Pike argues that while languages like C, C++, and Java have been sufficient for server software development, they do not directly address the properties of the modern computing environment. With the rise of networking, cluster computing, and big data, there is a growing need for a language that is efficient and can run on multiple machines. Pike highlights the importance of dependency management in Go programming and how it contributes to faster build times. Unlike other languages, Go's import mechanism and clean dependency hierarchy prevent redundant imports and unnecessary recompilation. This results in significantly faster build times, with Go programs being built in seconds compared to minutes or hours for other languages. The recent release of Go version 1, which offers stability and locked-down APIs, has led to increased adoption of the language. Pike emphasizes that the Go community prioritizes effective use of the language rather than constantly releasing new versions. This focus on stability and usability has contributed to the language's popularity and widespread adoption. Pike concludes by expressing his appreciation for the opportunity to address the importance of utilizing an efficient and swift language for present tasks, rather than solely focusing on its development. He believes that Go is a next-generation language that is well-suited for today's modern computer environment, with its fluidity, ease of construction, and efficiency for building large programs. In summary, Rob Pike's podcast episode "Why Learn Go?" discusses the need for a new compiled language that can handle the demands of modern computing. He highlights the importance of dependency management in Go programming, which contributes to faster build times. The recent release of Go version 1, with its stability and locked-down APIs, has led to increased adoption of the language. Pike emphasizes the significance of utilizing an efficient and swift language for present tasks, rather than solely focusing on its development. Overall, Go is presented as a next-generation language for today's modern computer environment.</p></details>
  + [talk/rob2013a](https://www.youtube.com/watch?v=bj9T2c2Xk_s) <details><summary>The path to Go 1. Mar 14, 2013.</summary>
    <p><b>Summary:</b> In the podcast episode titled "The path to Go 1, " Rob Pike and Andrew Gerrand discuss the development and release of Go 1 at OSCON 2012. The episode begins with a discussion on the development of the Go programming language, which was designed to solve software writing problems at Google. Go is a statically typed and compiled language that focuses on composing programs using interfaces and native concurrency support. The speakers then delve into the growth and development of the Go project, including the use of the mercurial version control system and the Rietveld code review plugin. They highlight the challenges faced in maintaining stability and the implementation of weekly snapshots to ensure stability. However, this caused confusion among contributors and users. To address version skew issues, a formal release process was implemented, but users still struggled to stay up to date. This led to the development of the powerful tool called "go fix, " which parses and rewrites Go code, making it easier for users to update their code to the latest version of the language. While beneficial for the Go project, the tool has drawbacks, such as an increase in code churn and the perception of Go as an unstable language, hindering adoption by some companies. The development process of Go version one is then discussed, with the goal of addressing concerns about its perceived instability and providing a reliable version for companies to depend on. The process involved addressing concerns, improving the language and its libraries, and building a new build toolset. The engagement of developers in the open-source community, particularly in Windows, was crucial. The release of Go 1 marked a shift in development approach, with a focus on long-term compatibility and improvements to Windows support. Language changes included the introduction of a new rune type and improvements to APIs. The introduction of a new "go" tool eliminated the need for make files and other build scripts, improving dependency management and the development workflow. Go 1 also brought improvements in documentation and testing, with a redesigned website providing streamlined installation instructions and comprehensive documentation. The release had a positive impact on the Go programming community, as developers redirected their efforts towards improving performance, stability, and bug fixing. The development team behind Go 1 has shifted their focus to using Go themselves and gathering feedback for future versions. Active development is ongoing, with a focus on stability, bug fixes, and efficiency improvements. Major improvements have been made to the compiler's code generation and the garbage collector. The team is also working on portability and developing new libraries. In conclusion, the podcast episode provides insights into the development and release of Go 1, highlighting the challenges faced, the improvements made, and the positive impact it had on the Go programming community. The speakers emphasize the importance of clean and rigorous dependency management and the need to reach a wider audience through diverse talks and engaging tutorials.</p></details>
  + [talk/rob2013b](https://www.infoq.com/presentations/Go-Google/) Go at Google. Apr 13, 2013.
  + [talk/rob2013c](https://changelog.com/podcast/100) Go Programming with Rob Pike and Andrew Gerrand. Aug 14, 2013.
  + [talk/rob2013d](https://www.youtube.com/watch?v=qmg1CF3gZQ0) <details><summary>Concurrency Is Not Parallelism. Oct 20, 2013.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Concurrency Is Not Parallelism, " Rob Pike speaks at Heroku Waza 2012. He explores the difference between concurrency and parallelism in programming languages, with a focus on the Go programming language. Pike argues that concurrency is superior to parallelism and clarifies the misconceptions surrounding these concepts. Pike emphasizes the importance of communication in coordinating concurrent tasks and references Tony Hoare's paper on communicating sequential processes as a highly regarded resource. He also discusses the select statement in Go, which serves as a multi-way concurrent control switch. The episode delves into the practical problem of getting rid of obsolete manuals using gophers and the need for efficient execution tools. Pike introduces the idea of using multiple gophers to move books more efficiently, highlighting the concepts of concurrency and parallelism. By coordinating the actions of multiple gophers, the process can be sped up. The episode explores different design patterns and strategies for achieving higher throughput, such as introducing staging dumps and increasing the number of gophers involved. However, it acknowledges that the example used is simplistic and lacks real-world relevance. Pike also discusses the similarities between the design of a book pile and a web serving architecture. He introduces the concept of go routines in Go programming language, which are similar to threads but more efficient and easier to create. Go routines allow for parallel execution and increased efficiency in programming. The episode explains the use of channels in Go for communication between go routines and introduces the concept of select, which allows the program to listen to multiple channels at once. Pike highlights the advantages of go routines over traditional threads in terms of efficiency and cost-effectiveness. Furthermore, the episode discusses the use of closures and channels in concurrent Go programming. It explains how closures can be used to wrap background operations and perform tasks concurrently, highlighting their simplicity and efficiency. The episode also demonstrates the use of channels in building a load balancer, showcasing the ease of expressing concurrent operations and the benefits of using channels in Go. Overall, the episode explores the advantages of concurrency and its role in building efficient algorithms. It discusses load balancing in a busy system and the use of channels in the Go programming language for communication. The episode distinguishes between concurrency and parallelism and provides additional resources for further understanding. Pike expresses appreciation to Hiroko for the invitation. In conclusion, Rob Pike's podcast episode on "Concurrency Is Not Parallelism" provides valuable insights into the differences between concurrency and parallelism in programming languages, with a focus on the Go programming language. Pike emphasizes the importance of communication, introduces various concepts and tools in Go, and discusses the advantages of concurrency in building efficient algorithms.</p></details>
  + [talk/rob2014a](https://www.youtube.com/watch?v=VoS7DsT1rdM) <details><summary>Hello Gophers! Gophercon Opening Keynote. GopherCon 2014. Apr 24, 2014.</summary>
    <p><b>Summary:</b> In the GopherCon 2014 Opening Keynote by Rob Pike, the history and development of the Go programming language are explored. Pike discusses the initial discussions and drafting of the Go specification, as well as the challenges faced in creating the first Go compiler. He highlights the changes in syntax and functionality of the program over time, including the introduction of the 'print' keyword and the modification of the 'main' function. The significance of packages in the design of the Go programming language is emphasized. Pike explains the importance of wrapping up code into libraries, controlling dependencies, and ensuring fast build times. He also discusses the concept of package imports, the benefits of using linear compilation, and the use of the 'export' keyword. The importance of initialization in Go programming is also explored. Pike mentions the challenges faced in implementing controlled initialization in C and the importance of proper program initialization in Go. Various aspects of the Go programming language are discussed, including syntax, import and namespace, formatted printing, UTF-8 handling, and the history of semicolons. Pike explains the reflection-driven approach to formatted printing, the significance of UTF-8 in Go, and the unique handling of non-ASCII characters. The development and evolution of the Go programming language are highlighted, emphasizing the importance of collaboration and a small team in its design process. Pike mentions the first concurrent program written in Go, called the prime sieve, and its structure. He also discusses the evolution of the CSP programming language and the changes in the syntax of Go over time. The stability of the Go programming language is emphasized, as it has remained largely unchanged for almost five years. Pike mentions that this stability has instilled trust in users and attracted beginners with its fast compiler and binary execution. He also discusses the development process and the contributions of the open-source community. The 'select' feature in Go is highlighted as crucial for implementing concurrency and creating complex structures. Pike discusses the challenges of implementing a debugger for Go language and using Go as an embedded language. He mentions ongoing efforts in the open-source community to address these challenges. Overall, the GopherCon 2014 Opening Keynote by Rob Pike provides a comprehensive overview of the history, development, and unique features of the Go programming language.</p></details>
  + [talk/rob2014b](https://www.youtube.com/watch?v=PXoG0WX0r_E) <details><summary>Implementing a bignum calculator. Sydney Go meeup. Nov, 19 2014.</summary>
    <p><b>Summary:</b> In the Golang-syd podcast episode titled "Implementing a bignum calculator with Rob Pike, " Rob Pike, a renowned programmer, discusses his experience in developing an APL-like calculator language. The episode covers various topics related to the implementation of the calculator and the challenges faced during the process. Pike begins by reflecting on the history of the calculator, called Hawk, which was initially developed for educational purposes. He acknowledges the limitations of the calculator, such as lack of precision and floating point issues. Pike also mentions his involvement in publishing a book on the Unix programming environment. The discussion then moves on to the issues with the 30-year-old calculator, including inaccuracies in calculations and the absence of support for hexadecimal numbers. Pike expresses his desire for a calculator that can handle larger numbers and explains his decision to implement an APL interpreter in Go. APL, a programming language developed by Ken Iverson in the 1950s and 60s, is introduced as a groundbreaking language with a simple kernel based on linear algebra. Pike highlights the uniqueness of APL, which uses special characters instead of keywords, making it an interesting language to learn and use. He mentions the rarity of seeing APL in action and refers to a video showcasing the development of an APL expression. The episode also briefly touches upon IV, a programming language named after Ken Iverson, which is still in its early stages but has interesting features such as exact arithmetic with rationals and support for large numbers. Pike mentions IV's potential for certain cryptographic calculations. The podcast then delves into various mathematical concepts and operations in programming. It covers vectors, matrices, random numbers, and sorting, emphasizing the use of APL for calculations and its potential for creating complex programs. The limitations of the APL implementation, particularly in base 16, are also discussed. Pike further explains the process of scanning tokens, parsing them into a parse tree, and evaluating the result in the context of implementing a numerical processor in Go. He shares his experience in designing a lexical scanner based on concurrency and discusses bugs encountered and fixed during the development process. The episode concludes with Pike discussing the parsing of expressions in an arithmetic grammar and a programming language. He explains the rules for operands, binary operators, and statement lists, highlighting the simplicity of parsing in APL and the influence of recursive descent parsers. Overall, the episode provides insights into the implementation of an APL-like calculator language and the challenges faced during the process. Pike's expertise and experience in programming shine through as he shares his knowledge and enthusiasm for different programming languages.</p></details>
  + [talk/rob2015a](https://www.youtube.com/watch?v=cF1zJYkBW4A) <details><summary>The move from C to Go in the toolchain. GopherFest 2015. Jun 10, 2015.</summary>
    <p><b>Summary:</b> In this episode of GopherFest 2015, Rob Pike discusses the transition from using C code to Go code in the Go programming language. The decision to move the compiler from C to Go was driven by practical reasons such as easier writing and debugging, better modularity and tooling, and support for parallel execution. The benefits of the transition are already being seen, including simplifying the management of two co-existing languages, improving testing and profiling, and overall maintenance of the codebase. The challenges of implementing a concurrent garbage collector in C are discussed, including type ambiguity and aliasing issues. The use of segmented stacks and imprecise stack data collection in GCC Go is highlighted as a solution to these challenges. The development process involved transitioning from segmented to contiguous stacks for C code and converting the runtime to a type-safe language. The decision to translate the Go compiler from C to Go was made for correctness and to avoid introducing new bugs. A machine-generated Go compiler was created using a custom translator to convert code from C to Go. The resulting code is not optimal but can be improved using various tools. The use of a parser written in Yak and the need for manual configuration are mentioned. The text also explains how a source-to-source translator in Go works and how it was used to fix slow code generated by a previous compiler. The differences between the compiler and garbage collector in the Go programming language are discussed. The Go compiler does not free memory, leading to overhead. The Go compiler team has made optimizations to improve performance and memory usage, such as using the math big package, reducing memory usage, improving escape analysis, and hand tuning. Recent changes have increased compiler speed by 15%, including better escape analysis and unification of architectures. The Go Build tool simplifies the compilation of Go programs into different architectures and operating systems. It makes code more portable, reducing the need for duplicated code in C. A new portable assembler has been introduced, allowing for easier code development and compatibility across architectures. The tool also includes a linker with a translation tool and a library. The improvements and future plans for the Go programming language are discussed, with a focus on enhancing performance. The release of Go 1. 5 includes updates such as a new assembler, garbage collector, and scheduler, resulting in cleaner and faster code. The tool chain and runtime have also been improved, making the code easier to test and maintain. The goal is to make the language more portable and flexible, although challenges with different instruction sets remain.</p></details>
  + [talk/rob2015b](https://www.youtube.com/watch?v=PAAkCSZUG1c) <details><summary>Go Proverbs. Gopherfest. Nov 18, 2015.</summary>
    <p><b>Summary:</b> In this episode of Gopherfest 2015, Rob Pike discusses various topics related to the game of Go and programming in the Go language. He begins by introducing the book 'Go Proverbs Illustrated', which provides valuable insights into the principles of the game of Go and how they can be applied to programming. Pike emphasizes the difficulty that Westerners face in learning and excelling in the game of Go, highlighting the unique gameplay and strategic thinking required. Moving on to programming, Pike emphasizes the significance of code formatting and the benefits of adhering to Go's formatting guidelines, specifically gofmt. He stresses the importance of consistency and readability in code and discusses the advantages of using small interfaces and structuring APIs effectively. Pike also mentions the cultural aspect surrounding interfaces in the Go ecosystem and the need to make the zero value of all types in a package useful. Pike then delves into the topic of dependency trees in programming and advocates for keeping them small. He suggests that copying small portions of code instead of importing entire libraries can lead to faster compilation, easier maintenance, and simplicity. He also emphasizes the use of build tags in code to make it more compact and less dependent on other pieces, particularly when it comes to guarding syscall and cisco uses for portability and compatibility. The drawbacks of using the unsafe package in Go are discussed, with Pike discouraging its use due to the potential for crashes and instability. He emphasizes the importance of writing clear and simple code, avoiding clever coding techniques, and limiting the use of reflection in Go. Pike also highlights the significance of designing the architecture and naming components in building a big system in Go. He stresses the need for good names that aid in understanding the design and make programming feel natural. Additionally, Pike discusses the importance of user-focused documentation that explains the purpose and usage of functions. He suggests the use of proverbs in code documentation to effectively convey information and resolve arguments. In conclusion, this episode of Gopherfest 2015 provides valuable insights into the game of Go and programming in the Go language. Rob Pike emphasizes the importance of understanding key concepts in programming, such as code formatting, small interfaces, and effective API structuring. He also discusses the significance of keeping dependency trees small, avoiding the use of unsafe packages, and designing architecture and naming components in building big systems. Overall, this episode offers valuable advice and principles for programmers in the Go ecosystem.</p></details>
  + [talk/rob2015c](https://www.youtube.com/watch?v=rFejpH_tAHM) <details><summary>Simplicity is Complicated. dotGo 2015. Dec 2, 2015.</summary>
    <p><b>Summary:</b> In the podcast episode titled "dotGo 2015 - Rob Pike - Simplicity is Complicated, " Rob Pike discusses the success and complexity of the programming language Go. Pike, a renowned expert in the field, explains that while Go is often described as a simple language, it is not as straightforward as it seems. Pike begins by highlighting the importance of simplicity in Go's success. Unlike other languages that incorporate features from various sources, Go stands out for its straightforwardness. However, Pike expresses concern over the trend of languages converging towards a single language by adopting features from others, as it limits diversity in problem-solving approaches. He emphasizes the importance of having different languages optimized for different domains and ways of thinking. The podcast also delves into the challenges of balancing conciseness and readability in code. Pike emphasizes that readable code is easier to understand, work on, extend, and fix. He uses an example in the APL dialect called dialogue to illustrate how a concise program can be difficult to read. Pike also discusses the trade-off between using more expressive programming features and the potential decrease in efficiency. He emphasizes the importance of finding the right balance and selecting the appropriate features to ensure both conciseness and expressiveness in programming. The complexity hidden behind Go's simplicity is explored in detail. Pike discusses various aspects of Go, such as data types, functions, interfaces, and concurrency. He emphasizes the importance of a good implementation and effective tooling. Pike also highlights the complexity of garbage collection in Go, despite its lack of user interface. He mentions the simplicity of Go's concurrency model, particularly go routines, which allow for lightweight sub-processes. However, he explains that go routines involve complex management behind the scenes. The podcast also touches on Go's support for Unicode UTF-8, magic packages like net/http, and concurrency. Pike emphasizes the simplicity and ease of use of Go, as well as its popularity. He explores how Go handles numeric types and constants, highlighting the complexities involved in designing the language. Pike discusses the challenges of working with constants and interfaces, and the importance of packages in scoping and compiling. Despite their simplicity in usage, Go packages hide a lot of complexity. In conclusion, the podcast episode provides a comprehensive overview of the features and advantages of the Go programming language. Pike's insights shed light on the complexity hidden behind Go's simplicity and emphasize the importance of finding the right balance between simplicity and expressiveness in programming.</p></details>
  + [talk/rob2016a](https://www.youtube.com/watch?v=KINIAgRpkDA) <details><summary>The Design of the Go Assembler. GopherCon 2016. Aug 18, 2016.</summary>
    <p><b>Summary:</b> In the GopherCon 2016 podcast episode titled "The Design of the Go Assembler, " Rob Pike discusses the importance of assembly language for programmers. He explains that assembly language allows programmers to access system functions at the lowest level and optimize performance. It is necessary for tasks like bootstrapping environments and taking advantage of hardware features. Pike emphasizes that understanding assembly language provides a deeper understanding of how computers work. Pike goes on to discuss the structure of assembly language, highlighting common features such as labels, instructions, operands, and comments. He explains that most CPUs have a similar structure, allowing for a common grammar. Pike also mentions the development of a common grammar for all machines and the work of Ken Thompson in developing a C compiler. The podcast episode delves into the evolution of assemblers and compilers in the Go programming language. Pike explains that the transition from C code to a Go implementation began with the creation of a library called Liblink, resulting in faster builds. The Go compiler and linker have undergone significant changes, with the obsolescence structure selection being moved to the back end of the compiler and assembler. The old C source has been translated into Go programs and the labeling has been rewritten into a suite of libraries. Pike suggests replacing all assemblers with a single program written in Go. He discusses the use of the -S flag in a compiler to display assembler instructions, explaining that the instructions produced are pseudo instructions from the stat phase of the compiler. Pike highlights the advantages of using a common assembly language for easier programming and portability. The podcast episode also covers the development of a new assembler program that can assemble any machine by parsing the input language into binary form. Pike explains the process of text processing in assembler and the validation and testing methods used. He describes the development of a universal assembler that eliminates the need for hardware manuals and replaces multiple programs with a single one. The assembler is compatible with old ones and can handle multiple architectures. Pike concludes by mentioning the use of machine-generated disassemblers and the challenges of reverse engineering. The goal is to have a machine-generated assembler that can be used with different architectures, making it easier for developers. He describes the process of developing the assembler using the C programming language and finds it exciting and relatively easy. Overall, the podcast episode provides valuable insights into the design and evolution of the Go assembler, highlighting the importance of assembly language and the benefits of a common assembly language for programming and portability.</p></details>
  + [talk/rob2016b](https://www.youtube.com/watch?v=sDTGhIqyMjo) <details><summary>Stacks of Tokens: A study in interfaces. Sydney Go Meetup. September 19, 2016.</summary>
    <p><b>Summary:</b> In this episode of the Sydney Go Meetup, Rob Pike, a Distinguished Engineer at Google, discusses the use of interfaces to solve a specific problem. He emphasizes the importance of design, portability, and generating manuals. Pike explores the role of a lexer in extracting tokens from input and the challenges of converting assembler language code written in C to a new assembler. He also discusses the implementation of a new C compiler and its impact on the lecture, highlighting its ability to use assembly language programs and define constants, macros, and instructions. Pike then delves into the features and functionality of the C preprocessor, including enabling and disabling blocks of code, as well as the concept of a token reader for analyzing macro definitions in C programming. He provides an overview of token readers and stacks in the Go programming language, explaining how token readers are used to process files and includes, and how they are implemented as a stack. Pike discusses the process of pushing and popping token readers from the stack, as well as retrieving the top of the stack. He also mentions the implementation of a stack in the text and the overall process of building a text processor. Additionally, Pike explains the concept of an input in a parser and the need for additional pieces to handle specific tasks at the top level of the parser. He provides an overview of the C preprocessor and its various components, including parsing and invoking macros, preventing infinite recursion, and using hash functions to determine token types. Pike emphasizes the benefits of breaking down problems into simple components and combining them to solve complex problems. He also explains the process of file inclusion, macro definition, and the importance of input. next in the preprocessor. Throughout the episode, Pike reflects on his experience designing and implementing a parser and assembler. He discusses the benefits of structuring the program to implement one interface through all the pieces and using Graco interfaces to break down the problem. Pike praises the clean structure of the stack and the role of the tech scanner package in facilitating the implementation process. He also discusses the challenges of working with assembly code and the importance of thorough testing. Pike mentions the use of a concurrent scanner for lexical processing and his experience in writing a C preprocessor. Overall, this episode provides valuable insights into the use of interfaces, lexer, token readers, and stacks in solving complex programming problems. Pike's expertise and experience shine through as he shares his knowledge and practical tips for designing and implementing efficient parsers and assemblers.</p></details>
  + [talk/rob2017](https://www.youtube.com/watch?v=ENLWEfi0Tkg) <details><summary>Upspin. Gopherfest 2017. Jun 22, 2017.</summary>
    <p><b>Summary:</b> In this episode of Gopherfest 2017, Rob Pike discusses Upspin, an experimental project aimed at creating a secure and uniform framework for naming and sharing files and data globally. The podcast begins by highlighting the challenges of managing personal data in the modern world, where downloading and sharing digital media often means renting and potentially losing access if the account is lost. The history of data management is explored, with a focus on the shift towards cloud storage. The dissatisfaction with the current state of data ownership is expressed, and the need for a system like Upspin is emphasized. Upspin is described as a global space for storing and organizing data, prioritizing personal privacy, security, and data ownership. Unlike platforms like Dropbox, Upspin is not meant for corporations and aims to provide a secure data storage system that allows access to specific groups of people. The podcast delves into the technical aspects of Upspin, explaining that it uses the Go programming language and provides a brief overview of the infrastructure involved. The system uses email addresses as usernames for verification purposes and utilizes end-to-end encryption to ensure only authorized individuals can access the data. Sharing within the Upspin tree is done in a way that makes it easy to understand who can access shared content. The Go Centric Model, consisting of a key server, storage server, and directory server, is introduced as the foundation of Upspin. The key server stores user data, the storage server allows users to retrieve and store data based on references, and the directory server handles second-order lookups. The Upspin file system, its design decisions, and components are discussed, including the storage server, directory server, and client library. The podcast also explores the potential applications of Upspin, such as securely accessing data from a nursery camera or sharing an iTunes library. The importance of secure and convenient information access in the modern world is emphasized. Overall, Upspin aims to provide users with fine-grained control over their data and a unified computing experience across multiple devices. The project is in its early development stages and encourages user input to improve functionality. While there are still areas that need improvement, such as documentation and design, Upspin differentiates itself from other global file systems by allowing the directory server and storage server to be separate machines. The podcast concludes by mentioning Keybase, a cloud provider that offers secure data access and sharing, and its integration with various platforms for consistent access, privacy, and security.</p></details>
  + [talk/rob2018a](https://www.youtube.com/watch?v=_2NI6t2r_Hs) <details><summary>The History of Unix. Nov 7, 2018.</summary>
    <p><b>Summary:</b> In this episode of the podcast, titled "The History of Unix, " Rob Pike takes us on a journey through the evolution of computing and the development of Unix. As an insider in the field, Pike provides a personal account of the key pieces that shaped the modern computing world. Pike begins by discussing the early days of computing, where punch cards were used and the development of Unix was still in its infancy. He shares his experiences working with IBM computers and his fascination with ray tracing and designing optical systems. Pike reflects on the reliability of the technology and his determination to access more resources. Moving forward, Pike delves into the PDP-11 machine, highlighting its various components and features such as tape drives, disc racks, and the graphic wonder. He also mentions the challenges of graphics processing and the significance of a frame buffer. Pike provides insights into the visible components of the machines and the early days of computing. Pike then recounts his experience of bringing a Unix system to grad school in California and using it to run Voyager ground stations. He reflects on the challenges of importing software across the US border and the significance of translating fan fold to a photo phone. Pike also discusses his early work at Google, focusing on graphics and sound projects, and running the lab on a small amount of memory. The podcast episode also touches on Pike's time at Bell Labs in the 1980s, where he worked on interesting projects such as the Cardiac simulation and the development of the Multix operating system. Pike stumbled upon a PDP seven computer and created the Space War program, providing insight into early computer gaming. Throughout the episode, Pike emphasizes the importance of Unix and its impact on the computing world. He discusses the evolution of the operating system and his experience learning to program on Unix. Pike also highlights the significance of graphics and networking in the development of computing technology. Overall, this episode offers a fascinating glimpse into the history of Unix and the journey of Rob Pike as a key figure in the field of computer programming.</p></details>
  + [talk/rob2018b](https://www.youtube.com/watch?v=RIvL2ONhFBI) <details><summary>Go 2 Draft Specifications. Sydney Golang Meetup. Nov 13, 2018.</summary>
    <p><b>Summary:</b> In the Sydney Golang Meetup podcast episode featuring Rob Pike, the focus was on the draft specifications for Go 2. The episode covered a wide range of topics related to the Go programming language, including its popularity, the need for evolution, and proposals for improvements. The podcast highlighted the importance of Go's stability and compatibility, which allows developers to focus on writing code rather than the language itself. However, in order to reach a larger audience and address certain issues, the language needs to evolve without compromising compatibility. The Go team has been engaging in discussions and design drafts to explore potential improvements and changes to the language. These design drafts serve as a way to generate ideas and gather feedback, with the aim of ensuring compatibility and minimizing disruptions to existing code while making enhancements. Specific proposals were discussed in the episode, such as the use of a new keyword called 'check' to simplify error handling, the simplification of variable declarations, and the introduction of a function local error handling mechanism. The episode also emphasized the importance of better semantics and standardization in error handling, as well as the significance of having a standard error formatting in software packages. The concept of parametric polymorphism, also known as generics, in Java was also explored in the episode. The challenges faced in implementing this feature and the ongoing efforts to find a solution were discussed. The episode highlighted the importance of being able to sort and communicate between different types of data in programming, and the drawbacks of dynamic type checking and reflection. The episode also touched on the significance of contracts in programming, the role of the debugging tool Delve in the Go community, and the concepts of covariance and contravariance in object-oriented languages. Throughout the episode, the importance of collaboration and open-mindedness in language design and software engineering was emphasized. The collaborative effort involved in designing the Go programming language was discussed, highlighting the contributions of multiple individuals and the importance of reaching consensus. In conclusion, the Sydney Golang Meetup podcast episode featuring Rob Pike provided a comprehensive overview of the recent discussions and proposals surrounding the Go programming language, specifically focusing on the concept of Go 2. The episode highlighted the need for the language to evolve while maintaining compatibility, and discussed specific proposals and challenges related to error handling, parametric polymorphism, and contracts. The importance of collaboration and diverse perspectives in language design and software engineering was also emphasized.</p></details>
  + [talk/rob2019a](https://changelog.com/gotime/100) Creating the Go programming language with Rob Pike & Robert Griesemer. Sep 10, 2019.
  + [talk/rob2019b](https://www.youtube.com/watch?v=oU9cfQCxjpM) <details><summary>Rob Pike. A Brief History of Go. Dec 12, 2019.</summary>
    <p><b>Summary:</b> In this episode of the GolangSyd Meetup, Rob Pike delivers a talk titled "A Brief History of Go" on December 12, 2019. The video of the talk can be found on YouTube, and the audio has been enhanced to provide a better listening experience. Unfortunately, the slides from the talk are not available. During his speech, Pike reflects on the origins and development of the Go programming language. He discusses the challenges faced in creating the language and highlights its success and ongoing updates. The name "Go" is derived from the first two letters of "Google, " emphasizing the language's connection to the company. Pike also explores the collaborative efforts of individuals in organizing conferences and distributing software, showcasing the language's progress over the course of two years. He mentions a significant conference in April 2013 that played a crucial role in promoting Go. The evolution and popularity of gopher figurines are also discussed. Pike mentions their distribution and their particular appeal in China. He highlights a well-designed figurine by Renee and a strong figurine released in 2016. Additionally, he mentions the launch of a language-related project. The importance of the go-to process in management and deployment is emphasized. Pike explains the benefits of faster testing with a cache version and introduces three proposals for language changes. He also mentions an unexpected surprise related to these proposals. The episode celebrates the 10th anniversary of Open Source Police, a community with 115, 000 members. Pike highlights the proposal-making process, the scaling back of proposals, and the various community projects that have been undertaken. He expresses gratitude for the support received and reflects on the challenges faced in the early years of the project. Pike acknowledges the difficulties in managing a community at the start of a project but expresses optimism about the progress that has been made. He emphasizes the importance of effective community management from the beginning. Overall, this episode provides a comprehensive overview of the development and improvement of the Go programming language. Pike's insights and reflections offer valuable insights into the challenges and successes of creating and managing a programming language, as well as the importance of community involvement and collaboration.</p></details>
  + [talk/rob2020](https://evrone.com/rob-pike-interview) A Rob Pike Interview. (Date Unclear) Apr 30, 2020.
  + [talk/rob2021](https://www.youtube.com/watch?v=YXV7sa4oM4I) <details><summary>The Go Programming Language and Environment. John Lions Distinguished Lecture, UNSW, May 27, 2021.</summary>
    <p><b>Summary:</b> In this podcast episode titled "The Go Programming Language and Environment, " the last scheduled speaker was unable to attend due to illness. However, Rob Pike, a Google engineer and co-creator of the Go programming language, stepped in to speak. Pike expressed his gratitude for being able to attend virtually and shared his admiration for John Lyons, who had a significant influence on his career. He also mentioned how encounters with Lyons and an Australian named Ian Johnston led him to Bell Labs and eventually to Australia. Pike then delved into the success of the Go programming language. Initially disliked, Go has gained popularity in cloud computing due to its ability to address challenges faced by other languages. It offers scalability, security, performance, and automation. The language's success can be attributed to its reliable libraries, long-term stability, and better software writing capabilities. Go was designed with scalability in mind, focusing on concurrency and multi-core CPUs. One of the challenges in software development is the lack of multi-core languages and coordinating massive compute clusters. Building tools for languages like C++ and Java can be difficult due to their complexity and slow compilation. Pike discussed how Google's transition to a new build system reduced binary size by specifying dependencies more precisely. The podcast episode provides a comprehensive overview of Go programming, covering topics such as web server structure, interfaces, concurrency, and building a rate limiter. It explains how Go utilizes simple mechanisms like maps and arrays to compensate for the lack of parametric polymorphism, allowing for compatibility among different implementations. The significance of interfaces in Go is also highlighted, emphasizing their simplicity and flexibility. Pike further discussed how Go introduced goroutines for efficient parallelism and concurrency, and how channels and the select statement are used for communication and synchronization. The simplicity and safety of Go programming were emphasized, along with the importance of building libraries with built-in capacity. Go's safety features, including no pointer arithmetic and indexing checks, contribute to its reputation as a secure language. The podcast episode also touched on the evolution of Go's project structure, known as GOPATH, into the current module system, and its impact on program reasoning. Pike expressed his desire for additional language features like channel packages and multiplexers, as well as the difficulties faced in trying to make networking and channels work together. Overall, the episode highlighted the success and advantages of the Go programming language, including its compatibility, performance, and strong library support. Pike emphasized the importance of collaboration and code sharing in creating a thriving ecosystem in software development. The episode concluded with Pike expressing appreciation for a participant who couldn't attend and mentioning upcoming talks.</p></details>
- Robert Griesemer. (Dr. Robert Griesemer) [GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
  + Alma mater: ETH Zürich
  + [paper/robert1993](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.127.5290&rep=rep1&type=pdf) A Programming Language for Vector Computers. Doctor Dissertation. 1993.
  + [talk/robert2012a](https://www.youtube.com/watch?v=on5DeUyWDqI) <details><summary>E2E: Erik Meijer and Robert Griesemer. Going Go. Lang.NEXT. Mar 16, 2012.</summary>
    <p><b>Summary:</b> In this episode of E2E, Erik Meijer interviews Robert Griesemer, the designer of the Go programming language. Go is a concurrent, garbage-collected systems programming language with fast compilation. The conversation covers various topics related to Go and its suitability as a native language. The discussion begins with the question of whether Go is a native language. Griesemer explains that Go can be compiled to native code and has scripting capabilities, making it a versatile language for native development. He highlights the advantages of using Go, such as fast compilation and execution, as well as a new build tool called 'bill' that simplifies the development process. The conversation then delves into the significance of compilers in programming. Griesemer explains that faster startup times and code optimization are key benefits of using compilers. He also explores techniques used to optimize startup time in the JVM and the benefits of using native languages and natively compiled binaries. The philosophy of error handling in Go is another topic of discussion. Griesemer emphasizes the importance of handling errors rather than ignoring them and explains how Go simplifies error handling with its multiple return values feature. He also discusses different approaches to error handling, including the use of special error types and panics, and highlights the need to treat exceptional situations as such and not use exception handling as a control flow mechanism. The significance of interfaces in object-oriented programming, particularly in Go, is also explored. Griesemer explains how Go allows methods to be attached to various types and the use of interface types as specifications of methods. He suggests that implementation inheritance may be overrated and that embedding another type can achieve similar results. The Go programming language provides primitives for forwarding and delegation, allowing developers to build these concepts themselves. The debate surrounding the inclusion of generics in programming languages, particularly in the context of C++, is another topic discussed. Griesemer acknowledges the benefits of generics, such as improved code reusability and type safety, but also acknowledges the challenges they bring, including integration with existing language features. He discusses the trade-offs and challenges of using templates and explores alternative approaches to generics. The conversation concludes with a discussion of Go routines, which are lightweight threads in Go that offer strong support for concurrency. Griesemer explains their implementation and use case, as well as the benefits of using channels for communication between Go routines. Overall, the conversation provides valuable insights into the design and features of the Go programming language, making it a must-listen for developers interested in Go and native language development.</p></details>
  + [talk/robert2012b](https://www.youtube.com/watch?v=Ns3csOBiIuE) <details><summary>Go In Three Easy Pieces. Mar 19, 2012.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Lang NEXT 2012 Go In Three Easy Pieces, " the speakers discuss various aspects of the Go programming language. They highlight its simplicity, compact size, and strong concurrency support, which have contributed to its growing popularity among developers. The episode begins by explaining the unique features of Go, such as its efficient implementation, powerful standard library, and a range of development tools. The language emphasizes comprehensibility and simplicity, making it easier for developers to understand and write code. It also has a different approach to exception handling and object-oriented programming, allowing methods to be defined without classes and providing built-in support for concurrency. The speakers then delve into the topic of string methods and type compatibility in Go. They explain how to use string methods using dot notation and static dispatch, and discuss the advantages of using them. They also discuss explicit conversion, particularly in the context of numeric types. The episode introduces interfaces and their ability to compose software effectively, with examples of dynamic dispatch and plug and play compatibility. The benefits of using interfaces in programming, such as flexible code reuse and the ability to interact with code without modification, are highlighted. The challenges of designing object hierarchies and implementing interfaces are also discussed. Next, the speakers discuss the use of file handlers in an HTTP application and the use of Ingo to present mandible images. They explain the implementation of a struct with various properties for the image and the methods used to generate it. The episode also mentions the use of a library function to encode the image as a PNG file. Additionally, the speakers discuss the use of closure generators to analyze the performance of a Mandelbrot server by measuring task execution time. The episode then focuses on the speed and efficiency of the Go programming language. It explains how Go can run code quickly and safely with features like garbage collection. The use of go routines and channels for lightweight and flexible code implementation is discussed. Channels are used for communication and synchronization between go routines, eliminating the need for manual synchronization. The speakers also discuss a typical Go concurrency pattern using channels and closures, as well as the concept of a work scheduler for efficient task execution. The process of transforming sequential code into concurrent code is explained, highlighting the benefits of faster computation and speedup. The episode concludes by discussing deferred statements in Go, which allow functions to be launched and suspended before executing the body. The panic and recovery mechanisms in Go for exception handling are also explained. The launch of Go is mentioned, highlighting its stability and backward compatibility. The use of Go in startups and its availability on App Engine are also mentioned. The episode ends by emphasizing the significance of language support in providing effective concurrency support in Go.</p></details>
  + [talk/robert2012c](https://www.youtube.com/watch?v=MrY6f4siVA4) <details><summary>Lang.NEXT 2012 Expert Panel: Native Languages. Apr 11, 2012.</summary>
    <p><b>Summary:</b> In the podcast episode titled "Lang NEXT 2012 Expert Panel: Native Languages, " a panel of experts discusses the advantages and limitations of native code programming. The panelists share their experiences and preferences, with some favoring native code for its speed, efficiency, and control. They emphasize the importance of using the most advantageous mechanism for each situation, comparing it to the choice of programming languages for specific tasks. The panelists discuss the benefits of native code execution, such as optimized user experiences and the ability to work closely with hardware. They also highlight the ability to directly call C code as a major advantage of native code. However, they acknowledge the potential issues with virtual machine (VM) implementation. The panelists also explore the advantages and disadvantages of native code and managed code. Native code offers speed and control, while managed code allows for advanced processing during load time. The choice between the two depends on the biggest cost for the application. The panelists discuss the challenges faced by native languages in achieving interoperability and language design. They mention the development of protocols for interoperability by companies like Google and Facebook. The panel expresses optimism about research advancements in this area and the significance of a common interface for plug-and-play compatibility. The podcast also highlights the shift in Facebook's backend development from PHP to C++. This change has been driven by the need for more efficient power consumption and improved performance. The use of C++ allows for a more efficient use of power per user and enables the development of high-performance ads. The panelists discuss the trade-off between program productivity and performance in software design. They emphasize the importance of prevention over cure and the different approaches taken by managed and native languages. The significance of maintaining consistent databases in the web world is also emphasized. The podcast concludes with a discussion on the complexity of programming languages and the challenges of memory management. The panelists mention the need for research to determine the most complex language specification and highlight the trade-off between safety and efficiency in memory management. Overall, the episode emphasizes the importance of using the appropriate programming language for specific tasks and staying updated with advancements in the field. The panelists provide insights into the advantages and limitations of native code programming and the ongoing developments in the industry.</p></details>
  + [talk/robert2015](https://www.youtube.com/watch?v=0ReKdcpNyQg) <details><summary>The Evolution of Go. GopherCon 2015. Jul 28, 2015.</summary>
    <p><b>Summary:</b> In the podcast episode titled "GopherCon 2015: Robert Griesemer - The Evolution of Go, " Robert Griesemer discusses the evolution of the Go programming language. Griesemer begins by sharing his background in programming languages, including studying under the creator of Pascal. He reflects on the challenges of creating a good programming language and the difficulty of maintaining productivity in his career. After programming in C++ for 15 years, Griesemer decides to start a new project to create a better language. He emphasizes the importance of simplicity, safety, and readability in language design. Griesemer believes that these principles are crucial in creating an easy-to-use and understandable language. He also discusses the decision-making process involved in language design and the importance of making certain aspects of programming easy to write. Griesemer mentions the lack of literature on language design but recommends two papers by Tony Hoare. He also discusses the process of brainstorming and the importance of direction and collaboration in language design. The podcast episode explores the evolution of programming languages, focusing on successors of Pascal and the programming language Oberon. Griesemer highlights the similarities and differences between Pascal and C, as well as the introduction of new features in languages like Modula-2 and Oberon. He also discusses the similarities between the programming languages Go and Oberon, emphasizing their shared origins in Oberon and the influence of C on Go. Griesemer then delves into the development of the Go programming language, highlighting its roots in Oberon and the incorporation of object orientation. He emphasizes the importance of object orientation and interfaces in Go, drawing inspiration from Smalltalk. Griesemer also discusses dynamic dispatch and the absence of generics in Go, explaining the challenges and the Go team's decision to prioritize established mechanisms. The team is still evaluating the implications of introducing generics. The development of the Go programming language involved a thorough testing process and emotional discussions among the small team of creators. The addition of Ross Cox improved the development process, and multiple independent implementations revealed bugs in both the compilers and the specification. The inclusion of the go types package further validated the code. The language has undergone significant changes, resulting in a more robust language with tools for easy adjustments and backward compatibility. The future of Go looks promising based on clear targets, available libraries and tools, market readiness, technological breakthroughs, and unique features. The podcast episode concludes by questioning whether Go will become mainstream and mentioning the need to unify the Go community. It also discusses pitfalls in language design, the history of programming languages, and how they reflect their creators. The episode wonders if the new language Co can be successful.</p></details>
  + [talk/robert2016a](https://www.youtube.com/watch?v=t-w6MyI2qlU) <details><summary>Lightning Talk: Alias Declarations for Gom: A proposal. GopherCon 2016. Oct 9, 2016.</summary>
    <p><b>Summary:</b> </p></details>
  + [talk/robert2016b](https://www.youtube.com/watch?v=vLxX3yZmw5Q) <details><summary>Prototype your design!. dotGo 2016. Nov 29, 2016.</summary>
    <p><b>Summary:</b> In the podcast episode titled "dotGo 2016 - Robert Griesemer - Prototype your design!", Robert Griesemer discusses the importance of prototyping in software design. He explains how prototyping can inform the existing design and lead to a better final design. Griesemer starts by emphasizing the significance of the design phase in the software development process. He mentions Stanford University's five-step design process as an example of how designing and prototyping are important in various industries. He highlights that designers are not just thinkers, but also doers. The main focus of the episode is on language design, specifically multi-dimensional slice support in the Go programming language. Griesemer explains that the Go community has been working on finding a solution for two-dimensional indexing, which is still a challenge. He proposes a solution to improve readability and performance in Go through multi-dimensional indexing. To implement this proposal, Griesemer emphasizes the importance of a rewriter prototype. This prototype allows for the analysis of design options and avoids making drastic changes to the compiler. He explains how index expressions can be manually or automatically rewritten into method calls using special method names and the plus operator for clarity. Griesemer discusses the process of rewriting and modifying code in the Go programming language. He explains how new index expressions can be given meaning by calling specific methods. He also mentions the importance of rewriting the syntax tree instead of the source code, as it allows for changes to be made to existing libraries with minimal code modifications. The episode also touches on the use of a Go type checker to improve accuracy and efficiency in the language. Griesemer explains how rewriting binary additions can address missing types and serve as the core of the prototype for experimentation and refinement. Throughout the episode, Griesemer emphasizes the importance of concrete implementation for testing and refining design. He highlights the surprise discovery of effective index operators and questions if they are sufficient for certain problems. He concludes by reiterating the value of prototyping in programming. In summary, the podcast episode explores the importance of prototyping in software design, specifically in the context of multi-dimensional slice support in the Go programming language. Griesemer discusses the process of rewriting and modifying code, as well as the use of a Go type checker for accuracy and efficiency. He emphasizes the significance of concrete implementation and the role of prototyping in informing and improving the final design.</p></details>
  + [talk/robert2017](https://www.youtube.com/watch?v=KPk1UPihWtY) <details><summary>Opening Keynote: Exporting Go. GopherCon SG 2017. May 29, 2017.</summary>
    <p><b>Summary:</b> In the opening keynote of GopherCon SG 2017, Robert Griesemer discusses the Go programming language and its use of packages to divide large projects. He compares the export and import mechanism of packages to the simpler times of the 1970s when C was invented. Griesemer highlights the importance of using header files in creating a library in C, as they allow for the explicit declaration of library interfaces, making it easier for the C compiler to compile the program. However, he also acknowledges the limitations and challenges of using header files, such as information leakage and potential performance issues in large systems. Griesemer explains that package serialization is important in programming languages as it allows the compiler to understand and communicate imported packages. He discusses the process of compiling Go packages and the need for serialization, as well as the use of dot o and dot a files for exporting and importing data. He mentions that object files can be read using a text editor and explains the structure and content of a pre-1. 7 compiler object file, which includes redundant and unnecessary information. To address these issues, Griesemer explains that a binary export format was adopted, which is more compact and efficient. The new format also includes additional information and allows for easy extension without a full-fledged parser. He discusses the serialization of package interfaces in Go, which involves representing the internal data structure as a graph. Griesemer explains a serialization algorithm in Go that can be used for any graph data structure, involving assigning unique integer values to nodes and writing out their content. Griesemer emphasizes the need to efficiently process export data by re-engineering it for indexing, reducing processing waste. He discusses the differences between textual and binary export formats, with the binary format being more space-efficient. However, challenges remain in processing only required fields. Griesemer suggests re-engineering export data to make it indexable for more efficient access, especially for large data structures like protocol buffers. Overall, Griesemer's keynote provides insights into the inner workings of import and export mechanisms in Go, and how these mechanisms enable robust separate compilation of packages. He discusses the historical context of these mechanisms and the design trade-offs that affect scalability. The talk highlights the importance of package serialization and the challenges involved in creating efficient and compact export formats.</p></details>
  + [talk/robert2017](https://www.youtube.com/watch?v=UmwJsQTSEP8) <details><summary>A brief overview of Go. Aug 28. 2017.</summary>
    <p><b>Summary:</b> In this episode of the podcast, Robert Griesemer, one of the designers of the Go programming language, gives a brief overview of Go and its key features. He starts by explaining that Go was developed by Google as an alternative to C++ and has gained significant success and growth since its release. It is highly regarded by developers and has between 500, 000 to a million users. Griesemer highlights the compactness, readability, and conciseness of Go, as well as its garbage collection feature and support for concurrency. He mentions that big tech companies like Google, IBM, Microsoft, and Uber, as well as companies like the New York Times, BBC, Amazon, and Walmart, use Go. It is also one of the fastest-growing languages in China. The unique features of Go are discussed, including constant declarations and powerful composite literal constructors. Griesemer explains the simplified syntax and powerful constructors of a scripted language, as well as the use of functions as first-class objects and closures. He showcases a mathematical matrix multiplication example to demonstrate the simplicity and efficiency of Go. The concept of methods and interfaces in an object-oriented language without classes or inheritance is also explained. Griesemer addresses the lack of support for multi-dimensional matrices and provides a workaround. He also mentions dynamic dispatch in programming. The use of goroutines and channels in concurrent programming is introduced. Channels in Go are used for communication and synchronization between goroutines, allowing for efficient communication without sharing memory. Griesemer discusses the use of channels in Go for efficient communication between goroutines, highlighting their cost-effectiveness and efficiency in terms of memory usage and communication. The author then explores the use of concurrency in matrix multiplication. They discuss their experiment of rewriting code using concurrent concurrency and explain the use of goroutines and wait groups. The performance differences between traditional and concurrent multiplication are also discussed. The features and benefits of Go are further discussed, including its standard library, platform independence, and tools for manipulating the language. The ease of writing an HTTP server in Go and its popularity for containerization and cloud environments are highlighted. The text concludes by mentioning the tools and features of the Co programming language, as well as the challenges of using C++ in scientific computation with Go. The Go programming language has an active community and extensive libraries, with increasing popularity and regular conferences and meetups worldwide. The community is working on improving dependency management, user experience, and inclusivity, with the goal of making Go the language of choice for cloud systems.</p></details>
  + [talk/robert2019](https://www.youtube.com/watch?v=i0zzChzk8KE) <details><summary>Go is 10! Now What?. Gopherpalooza 2019. Dec 2, 2019.</summary>
    <p><b>Summary:</b> </p></details>
  + [talk/robert2020a](https://changelog.com/gotime/140) The latest on Generics with Robert Griesemer and Ian Lance Taylor. Jul 21, 2020.
  + [talk/robert2020b](https://www.youtube.com/watch?v=TborQFPY2IM) <details><summary>Typing [Generic] Go. Nov 11, 2020.</summary>
    <p><b>Summary:</b> </p></details>
  + [talk/robert2021](https://www.youtube.com/watch?v=Pa_e9EeCdy8) <details><summary>Generics! Dec 17, 2021.</summary>
    <p><b>Summary:</b> </p></details>
- Ken Thompson. (Kenneth Lane Thompson, M. Sc.) (Retired)
  + Alma mater: UC Berkeley
  + [talk/ken1982a](https://www.youtube.com/watch?v=tc4ROCJYbm0) The UNIX System: Making Computers More Productive. 1982.
  + [talk/ken1982b](https://www.youtube.com/watch?v=XvDZLjaCJuw) UNIX: Making Computers Easier To Use.
  + [talk/ken1982c](https://www.youtube.com/watch?v=JoVQTPbD6UY) Ken Thompson and Dennis Ritchie Explain UNIX (Bell Labs).
  + [talk/ken1998](https://www.youtube.com/watch?v=LXZ1OL2U3lY) Ken Thompson and Dennis Ritchie. National Medal of Technology Awards. 1998.
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
  + [talk/ian2021a](https://www.youtube.com/watch?v=nr8EpUO9jhw) Using Generics in Go. Dec 6, 2021.
  + [talk/ian2021b](https://www.youtube.com/watch?v=Pa_e9EeCdy8) Generics! Dec 17, 2021.
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
- Austin Clements. (Dr. Austin T. Clements) [GitHub](https://github.com/aclements), [Scholar](https://scholar.google.com/citations?user=MKDtxN4AAAAJ)
  + Alma mater: MIT
  + [paper/autin2014](https://pdos.csail.mit.edu/papers/aclements-phd.pdf) The Scalable Commutativity Rule: Designing Scalable Software for Multicore Processors. Doctor Dissertation. May 21, 2014.
  + [talk/austin2020](https://www.gophercon.com/agenda/session/233441) Pardon the Interruption: Loop Preemption in Go 1.14. Nov 12, 2020.
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
- David Chase. (Dr. David Chase) [Website](http://chasewoerner.org/resume.html), [Block](https://dr2chase.wordpress.com/), [GitHub](https://github.com/dr2chase), [Twitter](https://twitter.com/dr2chase), [Scholar](https://dblp.org/pid/51/3488.html)
  + Alma mater: Rice University
  + [paper/1987david](http://www.chasewoerner.org/dissertation.pdf) Garbage Collection and Other Optimizations. Doctor Dissertation. Aug 1987.
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
- [talk/goteam2021](https://www.youtube.com/watch?v=hu9spj-GJ1U&) GopherCon 2021: Go Time Live - Go Team "Ask Me Anything", Dec 28, 2021.

[Back To Top](#top)

## Timeline

A timeline helps you identify the point in time about a document that is
reflected in Go versions.

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

[Back To Top](#top)

### Channel/Select

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov. Go channels on steroids. Jan 28, 2014
  + [issue/8899](https://go.dev/issue/8899) runtime: lock-free channels
  + [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) update on "lock-free channels"
  + [cl/112990043](https://codereview.appspot.com/112990043/) runtime: fine-grained locking in select
  + [cl/110580043](https://codereview.appspot.com/110580043/) runtime: add fast paths to non-blocking channel operations
- [issue/8898](https://go.dev/issue/8898) runtime: special case timer channels
- [issue/9120](https://go.dev/issue/9120) runtime: remove implementation restriction on channel element size
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
[Back To Top](#top)
- [issue/60353](https://go.dev/issue/60353) spec: when unifying against interface types, consider common methods
- [issue/57433](https://go.dev/issue/57433) slices: new standard library package based on x/exp/slices
- [issue/51259](https://go.dev/issue/51259) proposal: spec: support for struct members in interface/constraint syntax
- [issue/50285](https://go.dev/issue/50285) proposal: generic should infer type from variable definition

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
- [issue/54650](https://go.dev/issue/54650) x/exp/slices: memory leak on Delete

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
- [issue/55022](https://go.dev/issue/55022) proposal: cmd/compile: profile-guided optimization
- [issue/55025](https://go.dev/issue/55025) proposal: design and implementation of Profile-Guided Optimization (PGO)
- [design/pgo](https://go.dev/design/55022-pgo) Proposal: profile-guided optimization
- [design/pgo-implementation](https://go.dev/design/55022-pgo-implementation) Design and Implementation of Profile-Guided Optimization (PGO) for Go
- [issue/58409](https://go.dev/issue/58409) telemetry in the Go toolchain
- [issue/58894](https://go.dev/issue/58894) all: add opt-in transparent telemetry to Go toolchain
- [issue/54534](https://go.dev/issue/54534) cmd/compile: design doc explaining unified IR implementation
- [discuss/53060](https://go.dev/issue/53060) discussion: clarify Go support policy for secondary ports
- [issue/53383](https://go.dev/issue/53383) clarify Go support policy for secondary ports

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
- [design/profiler-labels](https://go.dev/design/17280-profile-labels) Michael Matloob. Proposal: Support for pprof profiler labels. May 15, 2017.
  + [issue/17280](https://go.dev/issue/17280) pprof: add support for profiler labels

[Back To Top](#top)

### Race Detector

- [issue/42598](https://go.dev/issue/42598) runtime: apparent false-positive race report for a buffered channel after CL 220419

[Back To Top](#top)

### Tracer

- [design/go15trace](https://go.dev/s/go15trace) Dmitry Vyukov. Go Execution Tracer. Oct 2014
- [design/tracefmt](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@. Trace Event Format. October 2016.
- [issue/54466](https://go.dev/issue/54466) runtime: rewrite gentraceback as an iterator API
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
- [design/lazy-gomod](https://go.dev/design/36460-lazy-module-loading) Bryan C. Mills. Proposal: Lazy Module Loading. Feb 20, 2020
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
- [issue/61179](https://go.dev/issue/61179) proposal: testing: add identity function that forces evaluation for benchmarks

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
- [issue/43997](https://go.dev/issue/43997) runtime: non-spinning Ms spin uselessly when work exists
- [issue/44313](https://go.dev/issue/44313) runtime: stopped Ms can't become dedicated or fractional GC workers
- [issue/51071](https://go.dev/issue/51071) runtime: investigate possible Go scheduler improvements inspired by Linux Kernel's CFS

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
  + [issue/50860](https://go.dev/issue/50860) proposal: sync/atomic: add typed atomic values
  + [issue/50859](https://go.dev/issue/50859) proposal: doc: update Go memory model
  + [doc/out-of-thin-air](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2019/p1217r2.html) Hans-J. Boehm. P1217R2: Out-of-thin-air, revisited, again. 2019-06-16.
  + [doc/cpp-memory-model](https://www.hpl.hp.com/techreports/2008/HPL-2008-56.pdf) Hans-J. Boehm, Sarita V. Adve. Foundations of the C++ Concurrency Memory Model. May 21, 2008.
  + [cl/381315](https://github.com/golang/go/commit/865911424d509184d95d3f9fc6a8301927117fdc) doc: update Go memory model
  + [cl/381316](https://github.com/golang/go/commit/a71ca3dfbd32faf351ff68bcc26a4d5abd9b06d7) runtime, sync, sync/atomic: document happens-before guarantees


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

[Back To Top](#top)

### Misc

- [issue/20135](https://go.dev/issue/20135) runtime: shrink map as elements are deleted
- [issue/48687](https://go.dev/issue/48687) runtime: enhance map cacheline efficiency
- [issue/54766](https://go.dev/issue/54766) runtime: use SwissTable

## Standard Library

### syscall

- [design/go14syscall](https://go.dev/s/go1.4-syscall) The syscall package.
- [issue/51087](https://go.dev/issue/51087) runtime: clean up system call code

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

It is also important to thank the continuing, inspiring discussion and sharing with
the TalkGo community core members [qcrao](https://github.com/qcrao),
[aofei](https://github.com/aofei), and [eddycjy](https://github.com/eddycjy).

The document would not be organized without all of the supports from them.

![](https://changkun.de/urlstat?mode=github&repo=golang-design/history)

[Back To Top](#top)

## License

[golang.design/history](https://github.com/golang-design/history) | CC-BY-NC-ND 4.0 &copy; [changkun](https://changkun.de)
