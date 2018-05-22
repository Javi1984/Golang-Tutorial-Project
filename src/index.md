---
title: 2017.07.19 - Angular2,4 and beyond
author: kcs
sort: 108
template: blog-article-2-level-nav.jade
group: blog
activeBlog: true
activeBlogArticle: true
links: [
       ]
---


### Links

  1) [Angular Blog](https://blog.angular.io)


### Angular

We are now on Angular version 4.  Finally with semantic versioning we have kept up to date with newer Angular versions.


### Team Training

The team will be expanding over the next few weeks.  We should take this opportunity to appreciate what the existin team
has learned, and prioritize how a new developer with no Angular experience should focus on training.

The good news is that the Enable2020 application code is not very deep.  We have very few of our own frameworks.  Therefore,
new developers are mainly concentrating on industry standard Angular development.  I remember early on talking to other
developers new to Angular struggling to see an overall application architecture.  Even needing to create one.  The truth is,
Angular is the application architecture.  Follow the documentation from Google and open source contributors and one will
do well in this framework.

Here's my recommended order of how things should be learned for this stack.

1. Javascript
    a. What is Asynchronous Programming
    b. Functional programming and how it helps with Asynchronous
    d. Truthfulness: if (null){], if( false){}, if( 0){}, if( ''), if([]){} ) are all false.
    c. JSON as a data structure standard

2. Node and NPM
    a. Command line, server side system
    b. NPM (Node Package Manager) all 3rd party library management done through some form of NPM.

3. RxJS
    a. An answer to Asynchronous programming adopted by Angular
    b. Subscribe/Unsubscribe, ReplaySubjects and Observables.
    c. You can test and learn from the nodejs command line.

4. ExpressJS
    a. A NodeJS based Web Server.
    b. Very simple and powerful URL patter routes making API exptremely simple to develop.

5. Typesript
    a. Transpiled Language.  Instead of compiled, it is converted to Browser compatible Javascript.
    b. Why not just Javascript???  Trust us!
    c. Type safety when you need it, or just use the "any" type
    d. Classes, Interfaces, etc.  Helps make Javascript and "Enterprise" language.

7. Angular
    a. Know the difference between Component, Service, Module, and Routes
    b. Modules control how Services are injected into Components
    c. Components model something visual, only alive while they are active
        i. Components are a class with functions and properties
        ii. Additionally, a component has a partial HTML called the template.
        iii. Component's values are sprayed into the template at runtime.

    d. Services are kept alive depending on where they are declared in the module tree.
        i. We currently use only a single module "app.module.ts"
        ii. Therefore all services are Singletons.  Never destroyed once first created.
        iii. This is where you store values that need to stay alive in a user's browser session.






